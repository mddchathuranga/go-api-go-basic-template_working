package log

import (
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"runtime"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gopkg.in/natefinch/lumberjack.v2"
)

var logger *logrus.Logger

// InitializeLogger initializes the logger with the specified log level and log file path.
func InitializeLogger() error {

	logFilePath, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	viper.SetConfigName("config") // Name of the configuration file without extension
	viper.AddConfigPath(".")      // Search the current directory for the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return fmt.Errorf("error reading config file: %v", err)
	}

	logLevel := viper.GetString("DTE.Alarm.Level")
	msName := viper.GetString("ProjectSettings.MsName")
	fileSizeLimitMegaBytes := viper.GetInt("DTE.Alarm.FileSizeLimitBytes")
	MaxAgeFromFile := viper.GetInt("DTE.Alarm.MaxAge")
	MaxBackupsFromFile := viper.GetInt("DTE.Alarm.MaxBackups")
	LocalTimeFromFile := viper.GetBool("DTE.Alarm.LocalTime")
	CompressStatusFromFile := viper.GetBool("DTE.Alarm.Compress")

	logger = logrus.New()

	// Set the log level based on the configuration
	switch strings.ToUpper(logLevel) {
	case "DEBUG":
		logger.SetLevel(logrus.DebugLevel)
	case "INFO":
		logger.SetLevel(logrus.InfoLevel)
	case "WARN":
		logger.SetLevel(logrus.WarnLevel)
	case "ERROR":
		logger.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		logger.SetLevel(logrus.FatalLevel)
	default:
		logger.SetLevel(logrus.InfoLevel)
	}
	// Customize the log format with a custom formatter
	logger.SetFormatter(&customFormatter{})
	logger.SetOutput(os.Stdout) // Console output
	relativeLogsFolderPath := "logs"
	logmainfolderpath := filepath.Join(logFilePath, relativeLogsFolderPath)

	if _, err := os.Stat(logmainfolderpath); os.IsNotExist(err) {
		if err := os.Mkdir(logmainfolderpath, 0755); err != nil {
			logrus.Fatalf("Fail to create logmainfolderpath folder:%s", err)
			return err
		}

	}
	logfoldername := "logs"
	logfolderpath := filepath.Join(logmainfolderpath, logfoldername)

	if _, err := os.Stat(logfolderpath); os.IsNotExist(err) {
		if err := os.Mkdir(logfolderpath, 0755); err != nil {
			logrus.Fatalf("Fail to create logfolderpath folder:%s", err)
			return err
		}

	}

	logsFolderPath := filepath.Join(logfolderpath, viper.GetString("ProjectSettings.AppName"))
	fmt.Println("logsFolderPath:", logsFolderPath)
	if _, err := os.Stat(logsFolderPath); os.IsNotExist(err) {
		if err := os.Mkdir(logsFolderPath, 0755); err != nil {
			logrus.Fatalf("Fail to create logs folder:%s", err)
			return err
		}

	}
	dateStamp := time.Now().Format("2006-01-02")
	// Set the log output to both Console and File
	logFileName := fmt.Sprintf("%s-app-%s.log", msName, dateStamp)

	logFile := &lumberjack.Logger{
		Filename:   filepath.Join(logsFolderPath, logFileName),
		MaxSize:    fileSizeLimitMegaBytes, //MaxSize is the maximum size in megabytes of the log file before it gets rotated. It defaults to 100 megabytes.
		MaxAge:     MaxAgeFromFile,         //MaxAge is the maximum number of days to retain old log files based on the timestamp encoded in their filename. Note that a day is defined as 24 hours and may not exactly correspond to calendar days due to daylight savings, leap seconds, etc. The default is not to remove old log files based on age.
		MaxBackups: MaxBackupsFromFile,     // MaxBackups is the maximum number of old log files to retain. The default is to retain all old log files (though MaxAge may still cause them to get deleted.)
		LocalTime:  LocalTimeFromFile,      //LocalTime determines if the time used for formatting the timestamps in backup files is the computer's local time. The default is to use UTC time.
		Compress:   CompressStatusFromFile, //Compress determines if the rotated log files should be compressed using gzip. The default is not to perform compression.
	}

	multiWriter := io.MultiWriter(logFile, os.Stdout)
	logger.SetOutput(multiWriter) // File output

	return nil
}

// customFormatter is a custom logrus formatter to match the desired log output format
type customFormatter struct{}

// Format formats the log entry to match the desired output format
func (f *customFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	viper.SetConfigName("config") // Name of the configuration file without extension
	viper.AddConfigPath(".")      // Search the current directory for the configuration file
	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("error reading config file: %v", err)
	}

	timestamp := time.Now().Format("2006-01-02 15:04:05.000")
	level := strings.ToUpper(entry.Level.String())
	message := entry.Message
	// Generate a UUID
	uuid, err := uuid.NewRandom()
	if err != nil {
		return nil, fmt.Errorf("failed to generate UUID: %v", err)
	}
	thread := fmt.Sprintf("Thread-%d", getGoroutineID())
	hostName, err := os.Hostname()
	if err != nil {
		return nil, fmt.Errorf("failed to get host name: %v", err)
	}
	logAppender := viper.GetString("ProjectSettings.LogAppender")
	appName := viper.GetString("ProjectSettings.AppName")
	msName := viper.GetString("ProjectSettings.MsName")
	messageType := viper.GetString("ProjectSettings.MessageType")

	// Create the formatted log message
	formattedLog := fmt.Sprintf("%s|[%s]|%s| %s|%s|%s|%s|%s|%s|%s\n",
		timestamp, thread, hostName, level, logAppender, appName, msName, uuid, messageType, message)

	return []byte(formattedLog), nil
}

func GetLogger() *logrus.Logger {
	return logger
}
func getGoroutineID() int {
	var buf [64]byte
	n := runtime.Stack(buf[:], false)
	idStr := strings.Fields(strings.TrimPrefix(string(buf[:n]), "goroutine "))[0]
	id, _ := strconv.Atoi(idStr)
	return id
}
