package handler

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type AlertHandler struct {
	TargetFile string
}

func (hl AlertHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {

	log.Printf("receiving request from %v", request)
	defer request.Body.Close()
	body, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("Error getting request body %v", err)
	} else {
		print(string(body), hl.TargetFile)
	}
}

func print(text string, file string) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text); err != nil {
		log.Println(err)
	}
}
