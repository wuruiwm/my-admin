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

func (r *Rabbitmq) Publish(queueName string, data string) (err error) {
	//设置队列 不存在则创建
	q, err := r.Channel.QueueDeclare(queueName, true, false, false, false, nil)
	if err != nil {
		return errors.New("rabbitmq create queue error: " + err.Error())
	}
	//向队列插入消息
	err = r.Channel.Publish("", q.Name, false, false, amqp.Publishing{
		ContentType:  "application/json",
		Body:         []byte(data),
		DeliveryMode: 2,
	})
	if err != nil {
		return errors.New("rabbitmq publish message error: " + err.Error())
	}
	return nil
}

func (r *Rabbitmq) Consume(queueName string, handle func(amqp.Delivery, *Rabbitmq)) {
	//死循环 如果异常退出 则进行重连操作 重新启动消费者
	for {
		err := func() error {
			//设置队列 不存在则创建
			q, err := r.Channel.QueueDeclare(queueName, true, false, false, false, nil)
			if err != nil {
				return errors.New("rabbitmq create queue error: " + err.Error())
			}

			//开始消费 获取消息队列chan
			msgChan, err := r.Channel.Consume(q.Name, "", false, false, false, false, nil)
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
		//异常重连逻辑 保证网络抖动或rabbitmq宕机等情况下 重启消费者
		if err != nil {
			global.Logger.Error("rabbitmq",
				zap.String("error", err.Error()),
			)
			for {
				time.Sleep(time.Second * 5)
				err := r.InitConn()
				if err == nil {
					break
				}
				global.Logger.Error("rabbitmq",
					zap.String("error", err.Error()),
				)
			}
		}
	}
}

func (r *Rabbitmq) Close() {
	_ = r.Channel.Close()
	_ = r.Conn.Close()
}
