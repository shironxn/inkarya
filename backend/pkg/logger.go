package pkg

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type LoggerService interface {
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
}

type Logger struct {
	log *zap.Logger
}

type LoggerConfig struct {
	Level       string
	Development bool
}

func NewLogger(cfg LoggerConfig) LoggerService {
	// Set default values if not provided
	if cfg.Level == "" {
		cfg.Level = "info"
	}

	// Create encoder config
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "ts",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	// Parse log level
	level, err := zapcore.ParseLevel(cfg.Level)
	if err != nil {
		level = zapcore.InfoLevel
	}

	// Create console core
	core := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		level,
	)

	// Create logger
	logger := zap.New(core,
		zap.AddCaller(),
		zap.AddStacktrace(zapcore.ErrorLevel),
	)

	if cfg.Development {
		logger = logger.WithOptions(zap.Development())
	}

	return &Logger{
		log: logger,
	}
}

func (l *Logger) Debug(msg string, fields ...zap.Field) {
	l.log.Debug(msg, fields...)
}

func (l *Logger) Info(msg string, fields ...zap.Field) {
	l.log.Info(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...zap.Field) {
	l.log.Warn(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...zap.Field) {
	l.log.Error(msg, fields...)
}

func (l *Logger) Fatal(msg string, fields ...zap.Field) {
	l.log.Fatal(msg, fields...)
}
