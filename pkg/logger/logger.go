package logger

import (
	"fmt"
	"os"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

const (
	// SHORTDATE format
	SHORTDATE = "02-Jan-2006"

	// LONGDATE format
	LONGDATE = "02-Jan-2006 15:04:05"
)

var (
	// Log zap logger
	Log *zap.Logger
)

// InitLog initializes logger
func InitLog() error {
	os.Mkdir("logs", os.ModePerm)

	highPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl >= zapcore.ErrorLevel
	})
	lowPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
		return lvl < zapcore.ErrorLevel
	})

	logConfig := zap.NewProductionEncoderConfig()
	logConfig.EncodeTime = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(LONGDATE))
	}

	consoleEncoder := zapcore.NewJSONEncoder(logConfig)

	debugLogFile, _ := os.OpenFile(fmt.Sprintf("./logs/Debug%s.log", time.Now().Format(SHORTDATE)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)
	errorLogFile, _ := os.OpenFile(fmt.Sprintf("./logs/Error%s.log", time.Now().Format(SHORTDATE)), os.O_APPEND|os.O_CREATE|os.O_WRONLY, os.ModePerm)

	consoleDebugging := zapcore.AddSync(debugLogFile)
	consoleError := zapcore.AddSync(errorLogFile)

	core := zapcore.NewTee(
		zapcore.NewCore(consoleEncoder, consoleDebugging, lowPriority),
		zapcore.NewCore(consoleEncoder, consoleError, highPriority),
	)

	Log = zap.New(core)

	return nil
}
