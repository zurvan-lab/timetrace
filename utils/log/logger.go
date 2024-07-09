package log

import (
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"reflect"
	"slices"
	"strings"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/zurvan-lab/TimeTrace/config"
	"gopkg.in/natefinch/lumberjack.v2"
)

var globalInst *logger

type logger struct {
	writer io.Writer
}

func InitGlobalLogger(cfg *config.Log) {
	if globalInst == nil {
		writers := []io.Writer{}

		if slices.Contains(cfg.Targets, "file") {
			// File writer.
			fw := &lumberjack.Logger{
				Filename:   cfg.Path,
				MaxSize:    cfg.MaxLogSize,
				MaxBackups: cfg.MaxBackups,
				Compress:   cfg.Compress,
			}
			writers = append(writers, fw)
		}

		if slices.Contains(cfg.Targets, "console") {
			// Console writer.
			writers = append(writers, zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: "15:04:05"})
		}

		globalInst = &logger{
			writer: io.MultiWriter(writers...),
		}

		level, err := zerolog.ParseLevel(strings.ToLower(cfg.Level))
		if err != nil {
			level = zerolog.InfoLevel
		}

		zerolog.SetGlobalLevel(level)
		log.Logger = zerolog.New(globalInst.writer).With().Timestamp().Logger()
	}
}

func addFields(event *zerolog.Event, keyvals ...interface{}) *zerolog.Event {
	if len(keyvals)%2 != 0 {
		keyvals = append(keyvals, "!MISSING-VALUE!")
	}

	for i := 0; i < len(keyvals); i += 2 {
		key, ok := keyvals[i].(string)
		if !ok {
			key = "!INVALID-KEY!"
		}
		///
		value := keyvals[i+1]
		switch v := value.(type) {
		case fmt.Stringer:
			if isNil(v) {
				event.Any(key, v)
			} else {
				event.Stringer(key, v)
			}
		case error:
			event.AnErr(key, v)
		case []byte:
			event.Str(key, hex.EncodeToString(v))
		default:
			event.Any(key, v)
		}
	}

	return event
}

func Trace(msg string, keyvals ...interface{}) {
	addFields(log.Trace(), keyvals...).Msg(msg)
}

func Debug(msg string, keyvals ...interface{}) {
	addFields(log.Debug(), keyvals...).Msg(msg)
}

func Info(msg string, keyvals ...interface{}) {
	addFields(log.Info(), keyvals...).Msg(msg)
}

func Warn(msg string, keyvals ...interface{}) {
	addFields(log.Warn(), keyvals...).Msg(msg)
}

func Error(msg string, keyvals ...interface{}) {
	addFields(log.Error(), keyvals...).Msg(msg)
}

func Fatal(msg string, keyvals ...interface{}) {
	addFields(log.Fatal(), keyvals...).Msg(msg)
}

func Panic(msg string, keyvals ...interface{}) {
	addFields(log.Panic(), keyvals...).Msg(msg)
}

func isNil(i interface{}) bool {
	if i == nil {
		return true
	}

	if reflect.TypeOf(i).Kind() == reflect.Ptr {
		return reflect.ValueOf(i).IsNil()
	}

	return false
}
