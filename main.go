package main

import (
	"fmt"
	"os"

	"plants/internal/pkg/dht"
	"plants/internal/pkg/healpers"
	"plants/internal/pkg/mqtt"
	"time"

	"github.com/akamensky/argparse"
)

func argparser() (string, string, string, string, int, bool, int) {
	parser := argparse.NewParser("flags", "Example of required vs optional arguments")
	user := parser.String("u", "user", &argparse.Options{Required: false, Help: "mqtt user"})
	password := parser.String("p", "password", &argparse.Options{Required: false, Help: "mqtt password"})
	deviceID := parser.String("d", "device_id", &argparse.Options{Default: "testing device", Help: "device id"})
	mqttHostname := parser.String("H", "mqtt_hostname", &argparse.Options{Default: "localhost", Help: "mqtt hostname"})
	mqttPort := parser.Int("P", "mqtt_port", &argparse.Options{Default: 1883, Help: "mqtt port"})
	testingMode := parser.Flag("t", "testing_mode", &argparse.Options{Help: "testing mode"})
	sleepTime := parser.Int("s", "sleep", &argparse.Options{Default: 15, Help: "sleep time beetwen executing"})
	err := parser.Parse(os.Args)
	if err != nil {
		fmt.Print(parser.Usage(err))
		os.Exit(1)
	}
	return *user, *password, *deviceID, *mqttHostname, *mqttPort, *testingMode, *sleepTime
}
func main() {

	fmt.Printf("Start %s\n", healpers.GetDate())
	user, password, deviceID, mqttHostname, mqttPort, testingMode, sleepTime := argparser()
	dht := dht.Sensor{
		Pin:          4,
		RetriesLimit: 4}

	mqttClient := mqtt.Client{
		Hostname: mqttHostname,
		Port:     mqttPort,
		Name:     deviceID,
		Username: user,
		Password: password}
	// go mqtt.Listen("test1"1)

	var i int
	for {
		temp, mois := dht.GetData()
		message := mqttClient.PrepareData(fmt.Sprintf("message_%d", 1), map[string]string{"temperature": fmt.Sprintf("%f", temp), "moisture": fmt.Sprintf("%f", mois)})
		mqttClient.Publish("office", message)

		if testingMode {
			time.Sleep(2 * time.Second)
			os.Exit(1)
		}
		i++
		time.Sleep(time.Duration(sleepTime) * time.Minute)
	}
}
