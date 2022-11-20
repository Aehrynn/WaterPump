package util

import (
	"encoding/json"
	"erinmaguire/WaterPump/structs"

	"github.com/HouzuoGuo/tiedot/db"
)

func GetLogDocuments(localDbClient *db.DB, collection string) []structs.WaterPumpLog {
	spaPasses := []structs.WaterPumpLog{}

	localDbClient.Use(collection).ForEachDoc(func(id int, docContent []byte) (moveOn bool) {
		spaPass := structs.WaterPumpLog{}
		json.Unmarshal(docContent, &spaPass)
		spaPasses = append(spaPasses, spaPass)
		return true // move on to the next document OR
		// do not move on to the next document
	})

	return spaPasses
}

func GetHumidityLogDocuments(localDbClient *db.DB, collection string) []structs.Humidity {
	spaPasses := []structs.Humidity{}

	localDbClient.Use(collection).ForEachDoc(func(id int, docContent []byte) (moveOn bool) {

		spaPass := structs.Humidity{}

		json.Unmarshal(docContent, &spaPass)
		spaPasses = append(spaPasses, spaPass)
		return true  // move on to the next document OR
		return false // do not move on to the next document
	})

	return spaPasses
}
