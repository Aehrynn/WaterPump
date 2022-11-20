package structs

type Configuration struct {
	WaterPumpRelayGpio int
	DatabaseLocation   string
	ApiKey             string
	HumidityCron       string
}
