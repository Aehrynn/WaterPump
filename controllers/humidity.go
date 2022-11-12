package controllers

import (
	"encoding/json"
	"erinmaguire/WaterPump/util"
	"fmt"
	"log"
	"net/http"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/d2r2/go-dht"
)

func GetHumidityLogs(w http.ResponseWriter, r *http.Request, apiKey string, localDbClient *db.DB) {
	authErr := util.Authorize(w, r, apiKey)

	if authErr != nil {
		return
	}

	spaPasses := util.GetLogDocuments(localDbClient)

	jsonSpaPasses, jsonErr := json.Marshal(spaPasses)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "server error")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonSpaPasses)
}

func GetCurrentHumidity(w http.ResponseWriter, r *http.Request, apiKey string) {
	authErr := util.Authorize(w, r, apiKey)

	if authErr != nil {
		return
	}
	temperature, humidity, retried, err :=
		dht.ReadDHTxxWithRetry(dht.DHT11, 27, false, 10)
	if err != nil {
		log.Fatal(err)
	}
	// Print temperature and humidity
	fmt.Printf("Temperature = %v*C, Humidity = %v%% (retried %d times)\n",
		temperature, humidity, retried)
}
