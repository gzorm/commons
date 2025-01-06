package ws

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/google/uuid"
	"github.com/gzorm/commons/core/logx"
	"time"
)

type ServiceContext struct {
	MQTTClient mqtt.Client
}

func NewServiceContext() *ServiceContext {
	return &ServiceContext{
		MQTTClient: NewMQTT(),
	}

}
func NewMQTT() mqtt.Client {
	var mqttClient mqtt.Client

	var broker = "132222222"
	var port = 1883
	opts := mqtt.NewClientOptions()
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", broker, port))
	opts.SetClientID(fmt.Sprintf("zero-game-rpc-%s", uuid.New().String()))
	opts.SetUsername("test001")
	opts.SetPassword("123456")
	opts.SetAutoReconnect(true)
	opts.SetMaxReconnectInterval(10 * time.Second)

	opts.OnConnect = func(client mqtt.Client) {
		logx.Info("connect mqtt")
	}
	opts.OnConnectionLost = func(client mqtt.Client, err error) {
		logx.Info("OnConnectionLost mqtt")
	}

	opts.SetConnectionLostHandler(func(c mqtt.Client, err error) {
		logx.Error("mqtt connection lost error:" + err.Error())
	})

	opts.SetReconnectingHandler(func(c mqtt.Client, options *mqtt.ClientOptions) {
		logx.Debug("mqtt reconnect")
	})
	opts.SetKeepAlive(time.Second * 20)
	mqttClient = mqtt.NewClient(opts)
	if token := mqttClient.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}
	return mqttClient
}
