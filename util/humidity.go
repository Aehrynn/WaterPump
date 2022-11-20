package util

import (
	"encoding/json"
	"erinmaguire/WaterPump/structs"
	"fmt"
	"log"
	"time"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/d2r2/go-dht"
)

func GetCurrentHumidity() structs.Humidity {
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 27, false, 10)
	if err != nil {
		log.Fatal(err)
	}

	// // Print temperature and humidity
	// fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
	// 	temperature, humidity, retried)

	return structs.Humidity{
		Humidity:    humidity,
		Temperature: temperature,
		Retried:     retried,
		Timestamp:   time.Now().UTC(),
	}
}

func LogCurrentHumidity(localDbClient *db.DB) {
	currentHumidity := GetCurrentHumidity()
	fmt.Printf("%+v\n", currentHumidity)
	var myMap map[string]interface{}
	data, _ := json.Marshal(currentHumidity)
	fmt.Println(data)
	json.Unmarshal(data, &myMap)
	fmt.Println(myMap)
	_, insertErr := localDbClient.Use("humidity_logs").Insert(myMap)

	if insertErr != nil {
		fmt.Println(insertErr)
	}
}
