package logger

import "go.uber.org/zap"

var log *zap.SugaredLogger

func Init(level string) error {
	cfg := zap.NewProductionConfig()

	if level != "" {
		// Игнорируем ошибку и остаёмся на уровне по умолчанию (info),
		// если уровень указан некорректно.
		_ = cfg.Level.UnmarshalText([]byte(level))
	}

	l, err := cfg.Build()
	if err != nil {
		return err
	}

	log = l.Sugar()
	return nil
}

func L() *zap.SugaredLogger {
	if log == nil {
		l, _ := zap.NewProduction()
		log = l.Sugar()
	}

	return log
}

func Sync() {
	if log != nil {
		_ = log.Sync()
	}
}
