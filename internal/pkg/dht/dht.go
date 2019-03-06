package dht

import (
	"fmt"

	"github.com/d2r2/go-dht"
)

// Sensor ...
type Sensor struct {
	boostPerfFlag bool
	Pin           int
	Type          dht.SensorType
	RetriesLimit  int
	Temperature   float32
	Humidity      float32
}

// GetData ...
func (sensor Sensor) GetData() (float32, float32) {
	sensor.Type = dht.DHT11
	temperature, humidity, retried, err := dht.ReadDHTxxWithRetry(sensor.Type, sensor.Pin, sensor.boostPerfFlag, sensor.RetriesLimit)
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	// get rid of variable as is not used
	_ = retried
	return temperature, humidity
}
