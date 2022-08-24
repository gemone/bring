package bring

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
)

// Logger interface used by this package. It is compatible with Logrus,
// but anything implementing this interface can be used
type Logger interface {
	Tracef(format string, args ...interface{})
	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// DefaultLogger Simple console logger
type DefaultLogger struct {
	Quiet bool
	Zap   *zap.Logger
}

func (l *DefaultLogger) Tracef(format string, args ...interface{}) {
	if l.Quiet {
		return
	}

	if l.Zap != nil {
		undo := zap.RedirectStdLog(l.Zap)
		defer undo()
	}

	log.Printf("TRAC: "+format, args...)
}

func (l *DefaultLogger) Debugf(format string, args ...interface{}) {
	if l.Quiet {
		return
	}

	if l.Zap != nil {
		undo, err := zap.RedirectStdLogAt(l.Zap, zapcore.DebugLevel)
		if err != nil {
			l.Zap.Error("err in debug log", zap.Error(err))
		}
		defer undo()
	}

	log.Printf("DEBU: "+format, args...)
}

func (l *DefaultLogger) Infof(format string, args ...interface{}) {
	if l.Quiet {
		return
	}

	if l.Zap != nil {
		undo := zap.RedirectStdLog(l.Zap)
		defer undo()
	}

	log.Printf("INFO: "+format, args...)
}

func (l *DefaultLogger) Warnf(format string, args ...interface{}) {
	if l.Quiet {
		return
	}

	if l.Zap != nil {
		undo, err := zap.RedirectStdLogAt(l.Zap, zapcore.WarnLevel)
		if err != nil {
			l.Zap.Error("err in waring log", zap.Error(err))
		}
		defer undo()
	}

	log.Printf("WARN: "+format, args...)
}

func (l *DefaultLogger) Errorf(format string, args ...interface{}) {
	if l.Quiet {
		return
	}

	if l.Zap != nil {
		undo, err := zap.RedirectStdLogAt(l.Zap, zapcore.ErrorLevel)
		if err != nil {
			l.Zap.Error("err in error log", zap.Error(err))
		}
		defer undo()
		defer undo()
	}

	log.Printf("ERRO: "+format, args...)
}
