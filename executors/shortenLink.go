package executors

import (
	"Ozon/Storages"
	"Ozon/generating"
	"Ozon/payload"
	"encoding/json"
	"github.com/sirupsen/logrus"
	"net/http"
)

func ShortenURL(w http.ResponseWriter, r *http.Request, log1 *logrus.Logger, stor Storages.Storage) {

	var request payload.RequestEncode
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		log1.Errorln("Invalid request body EndPoint: /shortLink ")
		return
	}
	log1.Infof("Received info: %s", request.Url)

	var response payload.ResponseEncode
	k := stor.ContainsByUrl(request.Url, log1)
	if k {
		elem := stor.GetByUrl(request.Url, log1)
		response.Hash = elem
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Unable to send response", http.StatusInternalServerError)
			log1.Errorln("Unable to send response EndPoint: /shortLink")
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonResponse)
		if err != nil {
			http.Error(w, "Unable to send response", http.StatusInternalServerError)
			log1.Errorln("Unable to send response EndPoint: /shortLink")
			return
		} else {
			log1.Infof("Successfully completed: %s", response.Hash)
		}
		return
	}

	flag, hashUrl := generating.GenerateShortenURL(request, stor, log1)
	if flag {
		response.Hash = hashUrl
		jsonResponse, err := json.Marshal(response)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			log1.Errorln("Internal Server Error EndPoint: /shortLink")
			return
		}

		stor.WriteByUrl(hashUrl, request.Url, log1)
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonResponse)
		if err != nil {
			http.Error(w, "Unable to send response", http.StatusInternalServerError)
			log1.Errorln("Unable to send response EndPoint: /shortLink")
			return
		} else {
			log1.Infof("Successfully completed: %s", response.Hash)
		}

	} else {
		http.Error(w, "Unable to generate hash", http.StatusInternalServerError)
		log1.Errorln("Unable to generate hash EndPoint: /shortLink")
		return
	}

}
