package executors

import (
	"Ozon/Storages"
	"Ozon/payload"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func LongURL(w http.ResponseWriter, r *http.Request, log1 *logrus.Logger, stor Storages.Storage) {
	request := r.URL.Query().Get("Short")
	response := stor.GetByHash(request, log1)
	var responseDec payload.ResponseDecode
	responseDec.Url = response
	jsonResponse, err := json.Marshal(responseDec)
	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(jsonResponse)
	if err != nil {
		http.Error(w, "Unable to send response", http.StatusInternalServerError)
		log1.Errorln("Unable to send response EndPoint: /shortLink")
		return
	} else {
		log1.Infof("Successfully completed: %s", responseDec.Url)
	}
}
