package lib

import (
	"os"
	"redis_publisher/config"
	"time"
)

type Logger struct {
	DateTime      string
	Service       string
	Type          string
	Msg           string
	FileLine      string
	MessageString string
}

const (
	ERROR    = "ERROR"
	INFO     = "INFO"
	WARNING  = "WARNING"
	NOTICE   = "NOTICE"
	CRITICAL = "CRITICAL"
	UNKNOWN  = "UNKNOWN"
)

func (logger *Logger) LogError(msg string) {
	logger.Log(msg, ERROR)
}

func (logger *Logger) LogINFO(msg string) {
	logger.Log(msg, INFO)
}

func (logger *Logger) LogWARNING(msg string) {
	logger.Log(msg, WARNING)
}

func (logger *Logger) LogNOTICE(msg string) {
	logger.Log(msg, NOTICE)
}

func (logger *Logger) LogCRITICAL(msg string) {
	logger.Log(msg, CRITICAL)
}

func (logger *Logger) LogUNKNOWN(msg string) {
	logger.Log(msg, UNKNOWN)
}

var cnf = config.Values

// Format:
// 2021-10-01 20:29:30,000 _| accounts.php.cli.consumer _| ERROR _| Redis server accounts-redis:6371 went away _| /src/kernel/system/classes/MessageQueue/Consumer.php _| 144
func (logger *Logger) Log(msg string, _type string) {

	current_time := time.Now()
	logger.DateTime = current_time.Format("2006-01-02 15:04:05") // time.Now().Format(time.RFC3339)

	if logger.Service == "" {
		logger.Service = cnf.Service
	}

	logger.Type = _type

	if logger.Type == "" {
		logger.Type = UNKNOWN
	}

	logger.Msg = msg

	logger.MessageString = logger.DateTime + " _| " + logger.Service + " _| " + logger.Type + " _| " + logger.Msg

	if logger.FileLine != "" {
		logger.MessageString = logger.MessageString + " _| " + logger.FileLine
	}
	logger.MessageString = logger.MessageString + "\n"

	//f, err := os.Create(logFilePath)
	f, err := os.OpenFile(cnf.LogPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	//logfile, err := os.OpenFile(logFilePath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {
		panic(err)
	}

	//defer logfile.Close()
	defer f.Close()
	f.WriteString(logger.MessageString)
	// log.SetOutput(logfile)
	// log.Print(logger.MessageString)

}
