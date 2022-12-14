package controllers

import (
	"encoding/json"
	"erinmaguire/WaterPump/structs"
	"erinmaguire/WaterPump/util"
	"fmt"
	"net/http"
	"time"

	"github.com/HouzuoGuo/tiedot/db"
	"github.com/stianeikeland/go-rpio/v4"
)

func GetWaterPumpLogs(w http.ResponseWriter, r *http.Request, apiKey string, localDbClient *db.DB) {
	authErr := util.Authorize(w, r, apiKey)

	if authErr != nil {
		return
	}

	spaPasses := util.GetLogDocuments(localDbClient, "logs")

	jsonSpaPasses, jsonErr := json.Marshal(spaPasses)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "server error")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonSpaPasses)
}

func PostToggleWaterPump(w http.ResponseWriter, r *http.Request, apiKey string, pin rpio.Pin, localDbClient *db.DB) {
	authErr := util.Authorize(w, r, apiKey)

	if authErr != nil {
		return
	}

	waterpumpLog := structs.WaterPumpLog{
		Action:    "Toggle",
		Timestamp: time.Now(),
	}
	var myMap map[string]interface{}
	data, _ := json.Marshal(&waterpumpLog)
	json.Unmarshal(data, &myMap)

	_, insertErr := localDbClient.Use("logs").Insert(myMap)

	if insertErr != nil {
		fmt.Println(insertErr)
	}
	// Toggle pin 20 times
	for x := 0; x < 20; x++ {
		pin.Toggle()
		time.Sleep(time.Second / 5)
	}
}
