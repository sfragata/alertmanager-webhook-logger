package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com.sfragata.alertmanager-webhook/handler"
	"github.com/integrii/flaggy"
)

// These variables will be replaced by real values when do gorelease
var (
	version = "none"
	date    string
	commit  string
)

func main() {

	info := fmt.Sprintf(
		"%s\nDate: %s\nCommit: %s\nOS: %s\nArch: %s",
		version,
		date,
		commit,
		runtime.GOOS,
		runtime.GOARCH,
	)
	flaggy.SetName("Alermanager Webhook file")
	flaggy.SetDescription("It receives alertmanager requests and log it into a file")
	flaggy.SetVersion(info)

	var webhookPort = "5001"
	flaggy.String(&webhookPort, "p", "port", "Listen port")

	var targetFile string
	flaggy.String(&targetFile, "t", "target", "target file")

	flaggy.Parse()

	if len(targetFile) == 0 {
		flaggy.ShowHelpAndExit("missing parameter target")
	}

	handler := handler.AlertHandler{TargetFile: targetFile}
	healthCheck := func(rw http.ResponseWriter, r *http.Request) {
		log.Printf("receiving healthcheck request from %v", r)
	}
	http.Handle("/", handler)
	http.HandleFunc("/health", healthCheck)

	log.Printf("Application started in port %s", webhookPort)
	err := http.ListenAndServe(fmt.Sprintf(":%s", webhookPort), nil)
	if err != nil {
		log.Fatalf("Can't start server %d", 5001)
	}

}
