package main

import (
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/oldfritter/HealthManager/initializers"
	. "github.com/oldfritter/HealthManager/models"
	"github.com/oldfritter/HealthManager/publishes"
	"github.com/oldfritter/HealthManager/subscribes"
)

func main() {
	initialize()
	subscribe()
	publish()

	err := ioutil.WriteFile("pids/healthManager.pid", []byte(strconv.Itoa(os.Getpid())), 0644)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("Health Manager started.")
	}

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	closeResource()
}

func initialize() {
	initializers.InitializeAmqpConfig()
	initializers.InitializeAmqpConnection("main")

	InitializeRabbitmqs()
	InitializeMysqlDBs()
	InitializeRedisDBs()
}

func subscribe() {
	subscribes.HeartBeat()
}

func publish() {
	publishes.InitSchedule()
}

func closeResource() {
	initializers.CloseAmqpConnection("main")
}
