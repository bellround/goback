package logmgr

import (
	logger "github.com/mccomack/simple-go-logger"
)

var manager *logger.Manager = nil

func new() (*logger.Manager, error) {
	manager, err := logger.Open(logger.Config{
		Directory:   "./logs",
		ProjectName: "goback",
	})

	if err != nil {
		return nil, err
	}

	return manager, nil
}

func Get() (*logger.Manager, error) {
	if manager != nil {
		return manager, nil
	}

	return new()
}
