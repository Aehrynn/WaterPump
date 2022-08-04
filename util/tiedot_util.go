package util

import (
	"encoding/json"
	"erinmaguire/WaterPump/structs"

	"github.com/HouzuoGuo/tiedot/db"
)

func GetLogDocuments(localDbClient *db.DB) []structs.WaterPumpLog {
	spaPasses := []structs.WaterPumpLog{}

	localDbClient.Use("spapasses").ForEachDoc(func(id int, docContent []byte) (moveOn bool) {
		spaPass := structs.WaterPumpLog{}
		jsonSpaPass, _ := json.Marshal(docContent)
		json.Unmarshal(jsonSpaPass, &spaPass)
		spaPasses = append(spaPasses, spaPass)
		return true // move on to the next document OR
		// do not move on to the next document
	})

	return spaPasses
}
