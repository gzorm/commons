package emqx

import (
	"fmt"
	mqtt "github.com/eclipse/paho.mqtt.golang"
	"testing"
	"time"
)

func TestMQTT(t *testing.T) {
	// 创建 EMQX 客户端实例
	client := NewEMQXClient("tcp://broker.emqx.io:1883", "test-client", "", "")

	// 连接到 EMQX
	if err := client.Connect(); err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	defer client.Disconnect()

	// 订阅主题
	client.Subscribe("test/topic", 1, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Message received: %s\n", msg.Payload())
	})

	// 发布消息
	client.Publish("test/topic", 1, false, "Hello from Golang!")

	// 保持连接一段时间以接收消息
	time.Sleep(10 * time.Second)
}

/*长连接的场景 DEMO

func main() {
	client := emqxclient.NewEMQXClient("tcp://broker.emqx.io:1883", "test-client", "", "")

	if err := client.Connect(); err != nil {
		fmt.Printf("Failed to connect: %v\n", err)
		return
	}
	// 使用 defer 确保程序退出时断开连接
	defer client.Disconnect()

	// 订阅主题
	client.Subscribe("test/topic", 1, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("Received message: %s\n", msg.Payload())
	})

    // 发布消息
	client.Publish("test/topic", 1, false, "Hello from Golang!")

	// 捕获退出信号
	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, os.Interrupt, syscall.SIGTERM)

	fmt.Println("Running... Press Ctrl+C to exit.")
	<-signalChan
	fmt.Println("Exiting...")
}



*/

/*JS 调用如下
import mqtt from 'mqtt';

const brokerURL = 'ws://broker.emqx.io:8083/mqtt'; // WebSocket 连接地址
const clientID = `mqttjs_${Math.random().toString(16).substr(2, 8)}`;

// 创建客户端
const client = mqtt.connect(brokerURL, {
  clientId: clientID,
  username: '', // 如需要认证，填入用户名
  password: '', // 如需要认证，填入密码
});

// 连接事件
client.on('connect', () => {
  console.log('Connected to EMQX');

  // 订阅主题
  client.subscribe('test/topic', (err) => {
    if (!err) {
      console.log('Subscribed to topic: test/topic');
    } else {
      console.error('Failed to subscribe:', err);
    }
  });

  // 发布消息
  client.publish('test/topic', 'Hello from JavaScript!');
});

// 接收消息事件
client.on('message', (topic, message) => {
  console.log(`Received message on topic "${topic}": ${message.toString()}`);
});

// 错误事件
client.on('error', (err) => {
  console.error('Connection error:', err);
});







*/
