package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MateSousa/aegis/config"
	"github.com/MateSousa/aegis/internal/app"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func main() {
	go sigterm()

	if err := loadConfig(".env"); err != nil {
		logrus.Error("Error to load environment variables...")
	}

	app.StartApi()

	sigterm()
}

func sigterm() {
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	<-sigs
	fmt.Println("Shutting down...")
	os.Exit(0)
}

func loadConfig(filename string) error {
	if err := config.ReadConfig(filename); err != nil {
		return errors.Wrap(err, "read container")
	}
	return nil
}
