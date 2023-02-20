package websocket

import (
	"app/util"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rabbitmq/amqp091-go"
	"nhooyr.io/websocket"
	"time"
)

var server *Server          //单例server
var queueName = "websocket" //广播队列名

// Server websocket服务端
type Server struct {
	Client     map[string]map[string]*Client //客户端map
	UserClient map[string][]string           //用户id到客户端id映射关系
	Register   chan *Client                  //注册客户端chan
	UnRegister chan *Client                  //注销客户端chan
	Message    chan *ServerMessage           //消息chan
}

// Client websocket客户端
type Client struct {
	Id      string              //连接id
	UserId  string              //用户id
	Group   string              //分组
	Conn    *websocket.Conn     //websocket连接
	Context context.Context     //连接context
	Message chan *ClientMessage //消息chan
}

// ServerMessage 用于多节点广播消息
type ServerMessage struct {
	Id       string         `json:"id"`      //发送到的连接id
	UserId   string         `json:"user_id"` //发送到的用户id
	Group    string         `json:"group"`   //发送到的分组
	Type     string         `json:"type"`    //group user
	Message  *ClientMessage `json:"message"` //发送的消息
	JsonByte []byte         `json:"-"`       //消息json后的byte数组
}

// ClientMessage 单个消息
type ClientMessage struct {
	Type     string      `json:"type"` //自定义的消息类型
	Data     interface{} `json:"data"` //消息内容
	JsonByte []byte      `json:"-"`    //消息json后的byte数组
}

// NewServer 实例化server并返回一个单例server
func NewServer() *Server {
	if server == nil {
		server = &Server{
			Client:     make(map[string]map[string]*Client, 0),
			UserClient: make(map[string][]string, 0),
			Register:   make(chan *Client, 16),
			UnRegister: make(chan *Client, 16),
			Message:    make(chan *ServerMessage, 128),
		}
		go server.ClientHandle()
		go server.MessageConsumeHandle()
		go server.MessagePublishHandle()
	}
	return server
}

// ClientHandle 客户端注册与注销处理
func (s *Server) ClientHandle() {
	var (
		ok     bool
		client *Client
	)
	for {
		select {
		//注册
		case client = <-s.Register:
			//维护客户端列表
			if _, ok = s.Client[client.Group]; !ok {
				s.Client[client.Group] = make(map[string]*Client, 0)
			}
			s.Client[client.Group][client.Id] = client
			//维护用户到客户端映射列表
			if client.UserId != "" {
				if _, ok = s.UserClient[client.UserClientKey()]; !ok {
					s.UserClient[client.UserClientKey()] = make([]string, 0)
				}
				s.UserClient[client.UserClientKey()] = append(s.UserClient[client.UserClientKey()], client.Id)
			}
		//取消注册
		case client = <-s.UnRegister:
			//维护用户到客户端映射列表
			if client.UserId != "" {
				if _, ok = s.UserClient[client.UserClientKey()]; ok {
					for k, v := range s.UserClient[client.UserClientKey()] {
						if v == client.Id {
							s.UserClient[client.UserClientKey()] = append(s.UserClient[client.UserClientKey()][:k], s.UserClient[client.UserClientKey()][k+1:]...)
							break
						}
					}
				}
			}
			//维护客户端列表
			if _, ok = s.Client[client.Group]; ok {
				if _, ok = s.Client[client.Group][client.Id]; ok {
					delete(s.Client[client.Group], client.Id)
				}
			}
		}
	}
}

// MessageConsumeHandle 订阅消息 在本节点客户端连接中发送消息
func (s *Server) MessageConsumeHandle() {
	var (
		mq      *util.Rabbitmq
		err     error
		message *ServerMessage
	)
	mq, err = util.NewRabbitmq()
	if err != nil {
		util.NewLogger().Error("websocket", util.Map{
			"error": err.Error(),
		})
	}
	mq.Consume(queueName, "fanout", func(delivery amqp091.Delivery, rabbitmq *util.Rabbitmq) {
		message = &ServerMessage{}
		if err = json.Unmarshal(delivery.Body, message); err != nil {
			util.NewLogger().Error("websocket", util.Map{
				"error": err.Error(),
			})
			return
		}
		if message.Type == "group" {
			go s.sendGroupMessage(message)
		} else if message.Type == "user" {
			go s.sendUserMessage(message)
		} else if message.Type == "client" {
			go s.sendClientMessage(message)
		}
	})
}

// MessagePublishHandle 发布消息 将本节点收到的消息进行发布
func (s *Server) MessagePublishHandle() {
	var (
		mq  *util.Rabbitmq
		err error
	)
	//死循环 当rabbitmq发送消息时做重新连接操作
	for {
		mq, err = util.NewRabbitmq()
		if err != nil {
			time.Sleep(time.Second * 5)
			mq.Close()
			err = mq.InitConn()
			if err != nil {
				util.NewLogger().Error("websocket", util.Map{
					"error": err.Error(),
				})
				continue
			}
		}
		//正常情况 启动后一直在这里等待消息 当rabbitmq发送消息失败才会进行重连
		for message := range s.Message {
			err = mq.Publish(queueName, "fanout", message.Json(), 0)
			if err != nil {
				//发送失败时 将消息重新放入chan 并跳出chan循环 进行重启生产者操作
				s.Message <- message
				mq.Close()
				break
			}
		}
	}
}

// RegisterClient 注册客户端到server
func (s *Server) RegisterClient(c *gin.Context, userId string, group string) error {
	var (
		conn *websocket.Conn
		err  error
	)
	conn, err = websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true, //允许跨域请求
	})
	if err != nil {
		return errors.New("request upgrade websocket error:" + err.Error())
	}
	client := NewClient(c, conn, userId, group)
	defer func() {
		client.Close()
		s.UnRegister <- client
	}()
	s.Register <- client
	client.ReadHandle()
	client.WriteHandle()
	return nil
}

// sendGroupMessage 将消息发送到组
func (s *Server) sendGroupMessage(message *ServerMessage) {
	if groupClient, ok := s.Client[message.Group]; ok {
		for _, client := range groupClient {
			client.SendMessage(message.Message)
		}
	}
}

// sendClientMessage 将消息发送给指定客户端
func (s *Server) sendClientMessage(message *ServerMessage) {
	if _, ok := s.Client[message.Group]; ok {
		if client, ok := s.Client[message.Group][message.Id]; ok {
			client.SendMessage(message.Message)
		}
	}
}

// sendUserMessage 将消息发送给指定用户的所有客户端
func (s *Server) sendUserMessage(message *ServerMessage) {
	if _, ok := s.UserClient[message.UserClientKey()]; ok {
		for _, id := range s.UserClient[message.UserClientKey()] {
			message.Id = id
			s.sendClientMessage(message)
		}
	}
}

// SendGroupMessage 将需要发送到组的消息发送到server
func (s *Server) SendGroupMessage(msg *ClientMessage, group string) {
	message := &ServerMessage{
		Id:      "",
		UserId:  "",
		Group:   group,
		Type:    "group",
		Message: msg,
	}
	s.SendMessage(message)
}

// SendUserMessage 将需要发送到指定用户的所有客户端的消息发送到server
func (s *Server) SendUserMessage(msg *ClientMessage, group string, userId string) {
	message := &ServerMessage{
		Id:      "",
		UserId:  userId,
		Group:   group,
		Type:    "user",
		Message: msg,
	}
	s.SendMessage(message)
}

// SendMessage 将消息发送到server 如果server chan堵塞 则丢弃本条消息
func (s *Server) SendMessage(message *ServerMessage) {
	select {
	case s.Message <- message:
	default:
	}
}

// NewClient 实例化客户端
func NewClient(c *gin.Context, conn *websocket.Conn, userId string, group string) *Client {
	return &Client{
		Id:      util.Uuid(),
		UserId:  userId,
		Group:   group,
		Conn:    conn,
		Context: c.Request.Context(),
		Message: make(chan *ClientMessage, 16),
	}
}

// Close 关闭客户端websocket连接
func (c *Client) Close() {
	_ = c.Conn.Close(websocket.StatusPolicyViolation, "close")
}

// ReadHandle 读取websocket客户端发送的消息 这里只处理连接错误和关闭连接 不处理消息
func (c *Client) ReadHandle() {
	var cancel context.CancelFunc
	c.Context, cancel = context.WithCancel(c.Context)
	go func() {
		defer cancel()
		for {
			messageType, _, err := c.Conn.Read(c.Context)
			//err或者消息类型为关闭连接 则退出
			if err != nil || messageType == 0 {
				return
			}
		}
	}()
	go c.Ping()
}

// WriteHandle 将客户端消息chan的消息发送到客户端
func (c *Client) WriteHandle() {
	var (
		message *ClientMessage
		err     error
	)
	for {
		select {
		case message = <-c.Message:
			err = c.WriteMessage(message)
			if err != nil {
				c.Close()
				util.NewLogger().Error("websocket", util.Map{
					"error": "write message error:" + err.Error(),
				})
				return
			}
		case <-c.Context.Done():
			c.Close()
			return
		}
	}
}

// Ping 向客户端发送ping 并等待pong 使用心跳包保持连接
func (c *Client) Ping() {
	ticker := time.NewTicker(time.Second * 5)
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			err := c.Conn.Ping(c.Context)
			if err != nil {
				c.Close()
				return
			}
		case <-c.Context.Done():
			return
		}
	}
}

// SendMessage 发送消息到客户端chan 如果客户端消息chan已满 则关闭客户端连接
func (c *Client) SendMessage(msg *ClientMessage) {
	select {
	case c.Message <- msg:
	default:
		go c.Close()
	}
}

// WriteMessage 将消息写入到客户端 5秒超时时间
func (c *Client) WriteMessage(message *ClientMessage) error {
	ctx, cancel := context.WithTimeout(c.Context, time.Second*5)
	defer cancel()
	return c.Conn.Write(ctx, websocket.MessageText, message.Json())
}

// UserClientKey 拼接并返回UserClient的key
func (c *Client) UserClientKey() string {
	return fmt.Sprintf("%s_%s", c.Group, c.UserId)
}

// Json json自身 返回json后的字节数组 并缓存
func (m *ServerMessage) Json() []byte {
	if m.JsonByte == nil {
		m.JsonByte, _ = json.Marshal(m)
	}
	return m.JsonByte
}

// UserClientKey 拼接并返回UserClient的key
func (m *ServerMessage) UserClientKey() string {
	return fmt.Sprintf("%s_%s", m.Group, m.UserId)
}

// NewClientMessage 实例化客户端消息
func NewClientMessage(t string, data interface{}) *ClientMessage {
	return &ClientMessage{
		Type:     t,
		Data:     data,
		JsonByte: nil,
	}
}

// Json json自身 返回json后的字节数组 并缓存
func (m *ClientMessage) Json() []byte {
	if m.JsonByte == nil {
		m.JsonByte, _ = json.Marshal(m)
	}
	return m.JsonByte
}
