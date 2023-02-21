package grader

import (
	"log"

	"github.com/isucon/isucandar"
	"github.com/logica0419/gasshuku-isucon/bench/logger"
	"github.com/logica0419/gasshuku-isucon/bench/model"
)

func CalcResult(result *isucandar.BenchmarkResult, finish bool) bool {
	logger.Admin.Print("")
	logger.Admin.Print("---------Bench Result---------")

	passed := true
	status := "pass"
	errors := result.Errors.All()

	setScore(result)
	scoreRaw := result.Score.Sum()

	logger.Admin.Printf("breakdown:")
	for tag, count := range result.Score.Breakdown() {
		logger.Admin.Printf("  %s: %d", tag, count)
	}

	errorCount := int64(0)
	timeoutCount := int64(0)
	for _, err := range errors {
		switch {
		case model.IsErrCanceled(err):
			continue
		case model.IsErrCritical(err):
			passed = false
			status = "fail: critical"
		case model.IsErrTimeout(err):
			timeoutCount++
		default:
			errorCount += 1
		}
	}
	deductionTotal := errorCount*10 + timeoutCount/10

	score := scoreRaw - deductionTotal
	if score <= 0 && passed {
		passed = false
		status = "fail: score"
	}

	var scoreLogger *log.Logger
	if finish {
		scoreLogger = logger.Contestant
	} else {
		scoreLogger = logger.Admin
	}

	scoreLogger.Print("")
	scoreLogger.Printf("status:    %s", status)
	scoreLogger.Printf("raw score: %d", scoreRaw)
	scoreLogger.Printf("deduction: %d (error: %d / timeout: %d)", deductionTotal, errorCount, timeoutCount)
	scoreLogger.Printf("score:     %d - %d = %d", scoreRaw, deductionTotal, score)

	return passed
}
