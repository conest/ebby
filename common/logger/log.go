package logger

import (
	"ebby/common/cfgloader"
	"os"

	"github.com/sirupsen/logrus"
)

// New : logger
func New() *logrus.Logger {

	log := logrus.New()

	log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})

	config := cfgloader.Init()

	if config.GetString("logger.level") == "debug" {
		log.SetLevel(logrus.DebugLevel)
	} else {
		log.SetLevel(logrus.WarnLevel)
	}

	log.SetOutput(os.Stdout)

	if config.GetBool("logger.logToFile") {
		file, err := os.OpenFile(
			config.GetString("logger.filePath"),
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0666)
		if err == nil {
			log.SetOutput(file)
		} else {
			log.Warn("Failed to log to file, using default stderr")
		}
	}

	return log
}
