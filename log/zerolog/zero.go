package logutil

import (
	"os"
	"io"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

const (
	defaultLevel = "info"
)

var Log zerolog.Logger

Log = zerolog.New(zerolog.ConsoleWriter{
	// Out: os.Stdout,
	Out:     file,
	NoColor: true,
}).With().Timestamp().Str("", "").Logger()

// NewZeroLog creates a new zerolog logger
func NewZeroLog(writer io.Writer) *zerolog.Logger {
	zl := zerolog.New(writer).Output(zerolog.ConsoleWriter{Out: writer}).With().Timestamp().Logger()

	errorlogFileHandler, err := os.OpenFile(config.Log.Path,		os.O_RDWR|os.O_CREATE|os.O_APPEND, 0660)
	if err != nil {
		return err
	}
	application.Logger = zerolog.New(errorlogFileHandler).With().
		Timestamp().
		Str("host", config.HTTP.Hostname).
		Int("port", config.HTTP.Port).
		Logger()
	return &zl
}
// SetupLogPath 设置日志文件的路径
func SetupLogPath(logFile string) {
	zerolog.TimeFieldFormat = "2006-01-02 15:04:05 MST"

	log.Logger = zerolog.New(output).With().Timestamp().Logger()
}

// Prepare will create logger instance.
func Prepare(output string) error {
	switch output {
	case "stdout":
		log.Logger = log.
			Output(os.Stdout).
			Level(parseLevel("debug"))
	default:
		f, err := file.NewFileWriter(l.Output)
		if err != nil {
			log.Fatal().Msgf("Failed to open logger file, %v.", err)
			return err
		}

		log.Logger = log.Output(f).Level(parseLevel("debug"))
	}

	return nil
}
// SetupLogLevel 设置日志文件的输出级别
func SetupLogLevel(level string) string {
	switch level {
	case "debug":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		return "debug"
	case "info":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		return defaultLevel
	case "warn":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
		return "warn"
	case "error":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		return "error"
	case "fatal":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
		return "fatal"
	case "panic":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
		return "panic"
	default:
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		return defaultLevel
	}
}


// parseLevel will parse the log level.
func parseLevel(l string) zerolog.Level {
	switch l {
	case "debug":
		return zerolog.DebugLevel
	case "info":
		return zerolog.InfoLevel
	case "warn":
		return zerolog.WarnLevel
	case "error":
		return zerolog.ErrorLevel
	case "fatal":
		return zerolog.FatalLevel
	case "panic":
		return zerolog.PanicLevel
	case "disable":
		return zerolog.Disabled
	}

	log.Panic().Msgf("Parse logger level failed.")
	return zerolog.Disabled
}
