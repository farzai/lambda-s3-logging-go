package usecase

import (
	"github.com/farzai/lambda-s3-logging-go/adapter"
	"github.com/farzai/lambda-s3-logging-go/domain"
)

type LogUsecase interface {
	LogDataTransfer(logEntry *domain.LogEntry) error
}

type logUsecase struct {
	cwLogsAdapter adapter.CloudWatchLogsAdapter
}

func NewLogUsecase(cwLogsAdapter adapter.CloudWatchLogsAdapter) LogUsecase {
	return &logUsecase{cwLogsAdapter: cwLogsAdapter}
}

func (u *logUsecase) LogDataTransfer(logEntry *domain.LogEntry) error {
	return u.cwLogsAdapter.PutLogEntry(logEntry)
}
