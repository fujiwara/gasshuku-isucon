package benchmark

import (
	"github.com/isucon/isucandar"
	"github.com/logica0419/gasshuku-isucon/bench/logger"
	"github.com/logica0419/gasshuku-isucon/bench/model"
)

func registerErrorHandler(b *Benchmark) {
	b.ib.OnError(func(err error, step *isucandar.BenchmarkStep) {
		if model.IsErrCritical(err) {
			logger.Contestant.Printf("critical error - %v", err)
			logger.Admin.Printf("critical error - %v", err)
			logger.Contestant.Print("--------- stop benchmarking ---------")
			logger.Admin.Print("--------- stop benchmarking ---------")
			step.Cancel()
			return
		}

		logger.Contestant.Printf("error - %v", err)
		logger.Admin.Printf("error - %v", err)
	})
}
