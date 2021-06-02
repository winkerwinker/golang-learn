package logger

import (
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

// Logger record all logs
var Logger = logrus.New()

func init() {


		Logger.SetLevel(logrus.DebugLevel)

		infoFile := "../logs/warning.log"
		infoDur := "720h"
		infoSplit := "24h"

		infoWriter, infoErr := initLogWriter(infoFile, infoDur, infoSplit)

		if infoErr != nil {
			panic(infoErr)
		}


		Logger.Hooks.Add(lfshook.NewHook(
			lfshook.WriterMap{
				logrus.InfoLevel:  infoWriter,
			},
			new(LogFormatter),
		))



}

func initLogWriter(filePath string, dur string, split string) (res *rotatelogs.RotateLogs, err error) {

	maxAge, _ := time.ParseDuration(dur)
	rotationTime, _ := time.ParseDuration(split)

	writer, err := rotatelogs.New(
		filePath+".%Y%m%d",
		rotatelogs.WithLinkName(filePath),
		rotatelogs.WithMaxAge(maxAge),
		rotatelogs.WithRotationTime(rotationTime),
	)

	if err != nil {
		Logger.Infof("failed to create rotatelogs: %s", err)
		return
	}

	return writer, nil

}
