package structs

import "time"

type Humidity struct {
	Temperature float32   `json:"temperature"`
	Humidity    float32   `json:"humidity"`
	Retried     int       `json:"retried"`
	Timestamp   time.Time `json:"timestamp"`
}
