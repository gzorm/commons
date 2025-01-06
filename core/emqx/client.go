package emqx

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type EMQXClient struct {
	client mqtt.Client
}

// NewEMQXClient 初始化 EMQX 客户端
func NewEMQXClient(broker string, clientID string, username string, password string) *EMQXClient {
	opts := mqtt.NewClientOptions().
		AddBroker(broker).
		SetClientID(clientID).
		SetUsername(username).
		SetPassword(password).
		SetKeepAlive(60 * time.Second).
		SetDefaultPublishHandler(func(client mqtt.Client, msg mqtt.Message) {
			fmt.Printf("Received message on topic: %s, message: %s\n", msg.Topic(), msg.Payload())
		}).
		SetPingTimeout(1 * time.Second)

	client := mqtt.NewClient(opts)
	return &EMQXClient{client: client}
}

// Connect 连接到 EMQX
func (c *EMQXClient) Connect() error {
	token := c.client.Connect()
	if token.Wait() && token.Error() != nil {
		return token.Error()
	}
	fmt.Println("Connected to EMQX broker")
	return nil
}

// Publish 发布消息
func (c *EMQXClient) Publish(topic string, qos byte, retained bool, payload interface{}) error {
	token := c.client.Publish(topic, qos, retained, payload)
	token.Wait()
	return token.Error()
}

// Subscribe 订阅主题
func (c *EMQXClient) Subscribe(topic string, qos byte, callback mqtt.MessageHandler) error {
	token := c.client.Subscribe(topic, qos, callback)
	token.Wait()
	return token.Error()
}

// Unsubscribe 取消订阅
func (c *EMQXClient) Unsubscribe(topics ...string) error {
	token := c.client.Unsubscribe(topics...)
	token.Wait()
	return token.Error()
}

// Disconnect 断开连接
func (c *EMQXClient) Disconnect() {
	c.client.Disconnect(250)
	fmt.Println("Disconnected from EMQX broker")
}
