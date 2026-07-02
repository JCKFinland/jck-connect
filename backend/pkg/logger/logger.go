package logger

import (
	"go.uber.org/zap"
)

type Logger struct {
	*zap.Logger
}

func New(environment string) (*Logger, error) {
	var (
		log *zap.Logger
		err error
	)

	if environment == "production" {
		log, err = zap.NewProduction()
	} else {
		log, err = zap.NewDevelopment()
	}

	if err != nil {
		return nil, err
	}

	return &Logger{
		Logger: log,
	}, nil
}