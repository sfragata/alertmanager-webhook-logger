package handler

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
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
	rollingFilename := rollingFile(file)
	f, err := os.OpenFile(rollingFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(text); err != nil {
		log.Println(err)
	}
}

func rollingFile(filename string) string {
	stringDate := time.Now().Format("2006-01-02") //golang time formatting
	var rollingFilename string
	index := strings.LastIndex(filename, ".")
	if index > -1 {
		rollingFilename = fmt.Sprintf("%s-%s%s", filename[0:index], stringDate, filename[index:])
	} else {
		rollingFilename = fmt.Sprintf("%s-%s", filename, stringDate)
	}
	return rollingFilename

}
