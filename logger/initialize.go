package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Initialize() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
}
