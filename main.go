package main

import (
	"flag"
	"github.com/jasonlvhit/gocron"
	"log"
	"monitoring_service/service"
)

func runApi() {
	log.Println("run api begin")

	service.Monitor()

	log.Println("run api end")
}

var (
	timeStr string
)

func init() {

	flag.StringVar(&timeStr, "time", "10:30", "time to run heart interval")

	flag.Parse()
}

func main() {

	log.Printf("the job run at %s every day", timeStr)
	gocron.Every(1).Day().At(timeStr).Do(runApi)

	log.Println("main begin")
	<-gocron.Start()

	log.Println("main end")
}
