package mqtt

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"

	"os"
	"plants/internal/pkg/healpers"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

// Client will do something
type Client struct {
	Name     string
	Hostname string
	Port     int
	Username string
	Password string
}

// ReturnURL will do sometihng
func (mqtt_client Client) ReturnURL() string {
	return "mqtt://internal:internal@localhost:1883/testing"
}

// Connect will do sometihng
func (mqtt_client Client) Connect(clientID string) mqtt.Client {
	// opts := createClientOptions(clientId, uri)
	opts := mqtt.NewClientOptions()
	// opts.AddBroker(fmt.Sprintf("tcp://%s:%d", mqtt_client.Hostname,mqtt_client.Port))
	opts.AddBroker(fmt.Sprintf("tcp://%s:%d", mqtt_client.Hostname, mqtt_client.Port))
	opts.SetUsername(mqtt_client.Username)
	opts.SetPassword(mqtt_client.Password)
	opts.SetClientID(clientID)
	client := mqtt.NewClient(opts)
	token := client.Connect()
	for !token.WaitTimeout(3 * time.Second) {
	}
	if err := token.Error(); err != nil {
		log.Fatal(err)
	}
	return client
}

// Listen will do sometihng
func (mqtt_client Client) Listen(topic string) {
	fmt.Println("Listen")
	client := mqtt_client.Connect("sub")
	client.Subscribe(topic, 0, func(client mqtt.Client, msg mqtt.Message) {
		fmt.Printf("* [%s] %s\n", msg.Topic(), string(msg.Payload()))
	})
}

// Publish will do sometihng
func (mqtt_client Client) Publish(topic string, message string) {
	fmt.Println("Publish")
	client := mqtt_client.Connect("pub")
	client.Publish(topic, 2, false, message)
}

type deviceInfo struct {
	Address  string `json:"address"`
	DeviceID string `json:"deviceid"`
	Hostname string `json:"hostname"`
	Mac      string `json:"mac"`
}
type mqttMessage struct {
	Time       string `json:"time"`
	ID         string `json:"id"`
	Data       string `json:"data"`
	DeviceInfo string `json:"device_info"`
}

// PrepareData ...
// var dataobj = `{"arr":` + data + `}`
// {"time": "2017-03-01 15:01",
//  "id": "message_001",
// 	"data": {"temperature": 21.5,
// 			 "humidity": 15.3,},
//  "device_info": {"address": "192.168.123.1",
//                  "device_id": "",
//                  "mac": ""}
// }
func (mqtt_client Client) PrepareData(messageID string, inputData map[string]string) string {

	hostname, _ := os.Hostname()
	ip := "unknown"
	ipData, _ := net.LookupHost(hostname)
	if len(ipData) > 0 {
		ip = fmt.Sprintf("%v", ipData[0])
	}
	res2D := &deviceInfo{
		Address:  ip,
		DeviceID: mqtt_client.Name,
		Mac:      healpers.GetMacAddr(),
		Hostname: hostname}
	res2B, err := json.Marshal(res2D)
	jsonInput, err := json.Marshal(inputData)
	if err != nil {
		fmt.Printf("%s\n", err)
	}

	res1D := &mqttMessage{
		Time:       healpers.GetDate(),
		ID:         messageID,
		Data:       string(jsonInput),
		DeviceInfo: string(res2B)}
	return strings.Replace(string(healpers.GetJSON(res1D)), "\\\"", "\"", -1)
}
