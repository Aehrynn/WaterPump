package structs

import "time"

type WaterPumpLog struct {
	Action    string    `json:"action"`
	Timestamp time.Time `json:"timestamp"`
}
