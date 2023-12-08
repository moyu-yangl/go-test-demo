package log

//
//import (
//	"go.uber.org/zap"
//	"go.uber.org/zap/zapcore"
//)
//
//var logger *zap.Logger
//var sugar *zap.SugaredLogger
//
//func init() {
//	newLogger, _ := NewLogger(Option{
//		Level:       "info",
//		OutPuts:     []string{"stdout", "test.log"},
//		ErrorOutPus: []string{"stdout", "error.log"},
//	})
//	logger = newLogger
//	sugar = logger.Sugar()
//}
//
//func NewLogger(o Option) (*zap.Logger, error) {
//	cfg := zap.Config{
//		Level:       zap.NewAtomicLevelAt(zap.DebugLevel),
//		Development: false,
//		Encoding:    "json",
//		EncoderConfig: zapcore.EncoderConfig{
//			TimeKey:        "log_time",
//			LevelKey:       "level",
//			NameKey:        "logger",
//			CallerKey:      "log_caller",
//			FunctionKey:    zapcore.OmitKey,
//			MessageKey:     "message",
//			LineEnding:     zapcore.DefaultLineEnding,
//			EncodeLevel:    zapcore.CapitalLevelEncoder,
//			EncodeTime:     zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05.000"),
//			EncodeDuration: zapcore.MillisDurationEncoder,
//			EncodeCaller:   zapcore.ShortCallerEncoder,
//		},
//		OutputPaths:      o.OutPuts,
//		ErrorOutputPaths: o.ErrorOutPus,
//	}
//	return cfg.Build()
//}
//
//type Option struct {
//	Level       string
//	OutPuts     []string
//	ErrorOutPus []string
//}
//
//func U_Info(msg string) {
//	defer sugar.Sync()
//	sugar.Infoln(msg)
//}
//
//func U_Error(msg string) {
//	defer sugar.Sync()
//	sugar.Errorln(msg)
//}
