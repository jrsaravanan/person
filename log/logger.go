//Package log Wrapper class for logrus
// logrus doesn't provide an option log rolling option
// file rolling handled externally using 'logrotate'
//
package log

import (
	"flag"
	"fmt"
	"os"

	"github.com/Sirupsen/logrus"
)

var (
	//TODO : get the log location from config
	//f, err = os.OpenFile("iris.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	//f *os.File

	//log level , default: error
	level string

	//logfile name and location
	location string

	//file location and name
)

//Logger to log
var Logger = logrus.New()

//InitLog initate log parameter
func init() {
	Logger.Formatter = new(logrus.JSONFormatter)

	flag.StringVar(&level, "level", "debug", "application log level")
	flag.StringVar(&location, "location", "security.log", "application log path and name")

	InitLog()
}

//InitLog initalize logger
func InitLog() {
	f, err := os.OpenFile(location, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0777)
	if err != nil {
		fmt.Printf("error opening file: %v , swithing to to default file iris.log", err.Error())
		f, err = os.OpenFile("iris.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
		if err != nil {
			fmt.Printf("error opening file: %v ", err.Error())
		}

	}

	if f != nil {
		//logrus.SetOutput(f)
		Logger.Out = f
	} else {
		Logger.Out = os.Stdout
	}

	lvl, err := logrus.ParseLevel(level)
	if err != nil {
		lvl = logrus.ErrorLevel
	}
	Logger.Level = lvl
	fmt.Println("logger level > ", level)
}
