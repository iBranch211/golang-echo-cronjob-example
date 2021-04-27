package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/robfig/cron/v3"
	"golang-echo-cronjob-example/config"
	"golang-echo-cronjob-example/handler"
	"golang-echo-cronjob-example/scheduler"
	"log"
)

func main() {
	c := cron.New()
	log.Println("Create new scheduler")
	scheduler.CreateScheduledJobs(c)

	c.Start()

	e := echo.New()
	e.HTTPErrorHandler = handler.ErrorHandler
	config.CORSConfig(e)

	// echo server 9000 de başlatıldı.
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%s", config.ServerPort)))
}
