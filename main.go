package main

import (
	"github.com/mathesmonte/localDeviceAPI/pkg/models"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request){
	w.WriteHeader(200)
	w.Write([]byte("Local Device API - Local IoT management api"))
}

func main() {
	models.ConnectDatabase()

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}