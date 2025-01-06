package ws

import (
	"fmt"
	"testing"
)

func TestPub(test *testing.T) {
	server := NewServiceContext()
	topic := fmt.Sprintf("%s/userprofileRefresh", "aa") // aa/userprofileRefresh
	token := server.MQTTClient.Publish(topic, 0, false, "okssss")
	token.Wait()
	if token.Error() != nil {
		fmt.Errorf("public mqtt eror:%v", token.Error())
	}
	fmt.Println("success")
}
