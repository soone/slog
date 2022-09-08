package slog

import (
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/sirupsen/logrus"
)

const ProduceMode = "produce"

var rotateWriter *rotatelogs.RotateLogs

func Default(path string) io.Writer {
	logrus.SetLevel(logrus.InfoLevel)

	mode := os.Getenv("mode")
	if mode == ProduceMode {

		logrus.SetFormatter(&logrus.JSONFormatter{})

		rotateOptions := []rotatelogs.Option{
			rotatelogs.WithRotationTime(time.Hour * 24),
		}

		var err error
		rotateWriter, err = rotatelogs.New(path, rotateOptions...)
		if err != nil {
			panic(err)
		}

		logrus.SetOutput(rotateWriter)
		return rotateWriter
	} else {
		logrus.SetOutput(os.Stdout)
	}

	return nil
}
