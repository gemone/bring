package bring

import (
	"go.uber.org/zap"
	"testing"
)

func TestDefaultLogger(t *testing.T) {
	logger := &DefaultLogger{}

	logger.Debugf("debug logger %d", 1)
	logger.Tracef("trace logger %d", 2)
	logger.Infof("info logger %d", 3)
	logger.Warnf("warn logger %d", 4)
	logger.Errorf("error logger %d", 5)
}

func TestDefaultLoggerWithZap(t *testing.T) {
	log, _ := zap.NewProduction()
	defer log.Sync()

	logger := &DefaultLogger{
		Zap: log,
	}

	logger.Debugf("debug logger %d", 1)
	logger.Tracef("trace logger %d", 2)
	logger.Infof("info logger %d", 3)
	logger.Warnf("warn logger %d", 4)
	logger.Errorf("error logger %d", 5)
}

func BenchmarkDefaultLogger(b *testing.B) {
	/*
		Benchmark Result
		goos: linux
		goarch: amd64
		pkg: github.com/gemone/bring
		cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
		136533          9682 ns/op
	*/
	for i := 0; i < b.N; i++ {
		logger := &DefaultLogger{}

		logger.Debugf("debug logger %d", 1)
		logger.Tracef("trace logger %d", 2)
		logger.Infof("info logger %d", 3)
		logger.Warnf("warn logger %d", 4)
		logger.Errorf("error logger %d", 5)
	}
}

func BenchmarkDefaultLoggerWithZap(b *testing.B) {
	/*
		Benchmark Result
		goos: linux
		goarch: amd64
		pkg: github.com/gemone/bring
		cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
		175021          6525 ns/op
	*/
	log, _ := zap.NewProduction()
	defer log.Sync()
	for i := 0; i < b.N; i++ {
		logger := &DefaultLogger{
			Zap: log,
		}

		logger.Debugf("debug logger %d", 1)
		logger.Tracef("trace logger %d", 2)
		logger.Infof("info logger %d", 3)
		logger.Warnf("warn logger %d", 4)
		logger.Errorf("error logger %d", 5)
	}
}

func BenchmarkDefaultLoggerWithZapAllocRuntime(b *testing.B) {
	/*
		Benchmark Result
		goos: linux
		goarch: amd64
		pkg: github.com/gemone/bring
		cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
		5122        225812 ns/op
	*/
	for i := 0; i < b.N; i++ {
		log, _ := zap.NewProduction()
		logger := &DefaultLogger{
			Zap: log,
		}

		logger.Debugf("debug logger %d", 1)
		logger.Tracef("trace logger %d", 2)
		logger.Infof("info logger %d", 3)
		logger.Warnf("warn logger %d", 4)
		logger.Errorf("error logger %d", 5)
		log.Sync()
	}
}

func BenchmarkLoggerWithZap(b *testing.B) {
	/*
		Benchmark Result
		goos: linux
		goarch: amd64
		pkg: github.com/gemone/bring
		cpu: Intel(R) Core(TM) i5-7200U CPU @ 2.50GHz
		501127	      2079 ns/op
	*/
	log, _ := zap.NewProduction()
	defer log.Sync()
	for i := 0; i < b.N; i++ {
		log.Debug("debug logger", zap.Int("var", 1))
		log.Info("trace logger", zap.Int("var", 2))
		log.Info("info logger", zap.Int("var", 3))
		log.Warn("warn logger", zap.Int("var", 4))
		log.Error("error logger", zap.Int("var", 5))
	}
}
