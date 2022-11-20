package controllers

import (
	"encoding/json"
	"erinmaguire/WaterPump/util"
	"fmt"
	"net/http"

	"github.com/HouzuoGuo/tiedot/db"
)

func GetHumidityLogs(w http.ResponseWriter, r *http.Request, apiKey string, localDbClient *db.DB) {
	authErr := util.Authorize(w, r, apiKey)

	if authErr != nil {
		return
	}

	spaPasses := util.GetHumidityLogDocuments(localDbClient, "humidity_logs")

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

	humidityStruct := util.GetCurrentHumidity()

	jsonSpaPasses, jsonErr := json.Marshal(humidityStruct)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "server error")
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(jsonSpaPasses)
}
