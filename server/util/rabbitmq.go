package util

import (
	"app/global"
	"errors"
	"fmt"
	"github.com/streadway/amqp"
	"go.uber.org/zap"
	"time"
)

type Rabbitmq struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
}

func (r *Rabbitmq) conn() (err error) {
	r.Conn, err = amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%d/", global.Config.Rabbitmq.Username, global.Config.Rabbitmq.Password, global.Config.Rabbitmq.Host, global.Config.Rabbitmq.Port))
	return
}

func (r *Rabbitmq) channel() (err error) {
	if r.Conn == nil {
		return errors.New("rabbitmq not conn")
	}
	r.Channel, err = r.Conn.Channel()
	return
}

func NewRabbitmq() (*Rabbitmq, error) {
	rabbitmq := &Rabbitmq{}
	err := rabbitmq.InitConn()
	if err != nil {
		return nil, err
	}
	return rabbitmq, nil
}

func (r *Rabbitmq) InitConn() (err error) {
	err = r.conn()
	if err != nil {
		return errors.New("rabbitmq conn error: " + err.Error())
	}
	err = r.channel()
	if err != nil {
		return err
	}
	return nil
}

func (r *Rabbitmq) CheckMode(mode string) (err error) {
	if !InArray(mode, []string{"direct", "fanout", "delay"}) {
		return errors.New("rabbitmq mode error: " + mode)
	}
	return nil
}

func (r *Rabbitmq) GetDelayQueueName(queueName string) string {
	return queueName + "-" + "delay"
}

/*
Publish 推送消息
queueName 队列名
mode 模式 可选值 direct生产消费  fanout订阅发布  delay延迟队列
data 消息内容
delayMilliSecond 延迟队列消息延迟时间(毫秒) mode为delay有效
*/
func (r *Rabbitmq) Publish(queueName string, mode string, data string, delayMilliSecond int) (err error) {
	if err = r.CheckMode(mode); err != nil {
		panic(err)
	}
	exchangeName := queueName
	queueType := mode
	if mode == "delay" {
		//设置队列 不存在则创建
		_, err = r.Channel.QueueDeclare(r.GetDelayQueueName(queueName), true, false, false, false, amqp.Table{
			"x-dead-letter-exchange": exchangeName,
		})
		if err != nil {
			return errors.New("rabbitmq create queue error: " + err.Error())
		}
		queueType = "direct"
	}
	err = r.Channel.ExchangeDeclare(exchangeName, queueType, true, false, false, false, nil)
	if err != nil {
		return errors.New("rabbitmq create exchange error: " + err.Error())
	}
	//向交换机插入消息
	if mode == "delay" {
		expiration := ""
		if delayMilliSecond > 0 {
			expiration = fmt.Sprintf("%d", delayMilliSecond)
		}
		err = r.Channel.Publish("", r.GetDelayQueueName(queueName), false, false, amqp.Publishing{
			ContentType:  "application/json",
			Body:         []byte(data),
			DeliveryMode: 2,
			Expiration:   expiration,
		})
	} else {
		err = r.Channel.Publish(exchangeName, "", false, false, amqp.Publishing{
			ContentType:  "application/json",
			Body:         []byte(data),
			DeliveryMode: 2,
		})
	}
	if err != nil {
		return errors.New("rabbitmq publish message error: " + err.Error())
	}
	return nil
}

/*
Consume 消费消息
queueName 队列名
mode 模式 可选值 direct生产消费  fanout订阅发布  delay延迟队列
handle 消息处理函数
*/
func (r *Rabbitmq) Consume(queueName string, mode string, handle func(amqp.Delivery, *Rabbitmq)) {
	var (
		autoAck      bool        //是否自动ack
		autoDelete   bool        //是否自动删除队列
		exchangeName = queueName //交换机名
		key          string
		err          error
	)
	if err = r.CheckMode(mode); err != nil {
		panic(err)
	}
	//direct 生产消费模型 不自动ack 不自动删除队列
	//fanout 订阅发布模型 自动ack 连接断开自动删除队列
	//delay  延迟队列模型 不自动ack 不自动删除队列 延迟一段时间后 放入队列等待执行
	if mode == "fanout" {
		autoAck = true
		autoDelete = true
		queueName = queueName + "-" + Uuid()
	} else if mode == "delay" {
		key = r.GetDelayQueueName(queueName)
	}
	//死循环 如果异常退出 则进行重连操作 重新启动消费者
	for {
		err = func() error {
			//设置队列 不存在则创建
			_, err = r.Channel.QueueDeclare(queueName, true, autoDelete, false, false, nil)
			if err != nil {
				return errors.New("rabbitmq create queue error: " + err.Error())
			}
			//绑定队列和交换机
			err = r.Channel.QueueBind(queueName, key, exchangeName, false, nil)
			if err != nil {
				return errors.New("rabbitmq queue bind error: " + err.Error())
			}
			//开始消费 获取消息队列chan
			msgChan, err := r.Channel.Consume(queueName, "", autoAck, false, false, false, nil)
			if err != nil {
				return errors.New("rabbitmq get queue chan error: " + err.Error())
			}
			//正常情况下 第一次执行后 会一直在这里堵塞 等待消息过来 进行消费
			//如果循环结束 则说明因为网络抖动或rabbitmq宕机等异常情况 chan会被close
			for d := range msgChan {
				handle(d, r)
			}
			return errors.New("rabbitmq consume exit")
		}()
		//异常重连逻辑 保证网络异常或rabbitmq宕机等情况下 重启消费者
		if err != nil {
			global.Logger.Error("rabbitmq", zap.String("error", err.Error()))
			for {
				time.Sleep(time.Second * 5)
				r.Close()
				err = r.InitConn()
				if err == nil {
					break
				}
				global.Logger.Error("rabbitmq", zap.String("error", err.Error()))
			}
		}
	}
}

func (r *Rabbitmq) Close() {
	_ = r.Channel.Close()
	_ = r.Conn.Close()
}
