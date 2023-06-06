package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"projectONE/internal/factory"
	httpProjectONE "projectONE/internal/http"
	"syscall"
	"time"

	db "projectONE/database"
	"projectONE/internal/config"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

var PASSWORD string

func init() {
	err := godotenv.Load()
	if err != nil {
		logrus.WithFields(logrus.Fields{
			"cause": err,
		}).Fatal("Load .env file error")
	}

	err = config.Load(PASSWORD)
	if err != nil {
		logrus.Fatal(err)
	}

	logrus.SetReportCaller(true)
	logrus.SetOutput(os.Stdout)
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		ForceColors:     true,
		FullTimestamp:   true,
		TimestampFormat: "2006-01-02 15:04:05",
	})

	var level logrus.Level = logrus.InfoLevel
	if config.Get().Logging.LogrusLevel != 0 {
		switch config.Get().Logging.LogrusLevel {
		case 1:
			level = 1
		case 2:
			level = 2
		case 3:
			level = 3
		case 4:
			level = 4
		case 5:
			level = 5
		case 6:
			level = 6
		}
	}
	logrus.SetLevel(level)

}

func main() {

	PORT := config.Get().Server.Port
	if PORT == "0" {
		PORT = "8080"
	}

	db.Init()

	e := echo.New()

	f := factory.NewFactory()

	httpProjectONE.Init(e, f)

	ch := make(chan os.Signal, 1)

	signal.Notify(ch, os.Interrupt, syscall.SIGTERM)

	e.Static("/", "assets")

	go func() {
		err := e.Start(":" + PORT)
		if err != nil {
			if err != http.ErrServerClosed {
				logrus.Fatal(err)
			}
		}
	}()

	<-ch

	logrus.Println("Shutting down server...")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel2()
	e.Shutdown(ctx2)
	logrus.Println("Server gracefully stopped")
}
