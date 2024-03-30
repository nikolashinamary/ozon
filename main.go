package main

import (
	"Ozon/Storages"
	"Ozon/executors"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func main() {
	port := ":8080"
	err := godotenv.Load()
	log1 := logrus.New()
	if err != nil {
		log1.Fatalln(err)
	}
	var stor Storages.Storage

	typeOfStorage := os.Getenv("STORAGE")
	if typeOfStorage == "DB" {
		database := Storages.DBConnection(log1)
		err = database.AutoMigrate(&Storages.Model{})
		stor = Storages.DatabaseConstr(database)
		defer Storages.DBDisconnection(database, log1)
	} else {
		stor = Storages.InMemoryStorageConstr()
	}
	http.HandleFunc("/shortLink", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			http.Error(w, "Method is nor supported", http.StatusMethodNotAllowed)
			log1.Errorln("Method is not supported EndPoint: /shortLink")
			return
		}
		log1.Infoln("Method: POST EndPoint: /shortLink")
		executors.ShortenURL(w, r, log1, stor)
	})
	http.HandleFunc("/longURL", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			http.Error(w, "Method is nor supported", http.StatusMethodNotAllowed)
			log1.Errorln("Method is not supported EndPoint: /shortLink")
			return
		}
		log1.Infoln("Method: POST EndPoint: /longURL")
		executors.LongURL(w, r, log1, stor)
	})
	log1.Fatalln(http.ListenAndServe(port, nil))

}
