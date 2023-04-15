package adapter

import (
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/farzai/lambda-s3-logging-go/domain"
)

type CloudWatchLogsAdapter interface {
	PutLogEntry(logEntry *domain.LogEntry) error
}

type cloudWatchLogsAdapter struct {
	cwLogsClient  *cloudwatchlogs.CloudWatchLogs
	logGroupName  string
	logStreamName string
}

func NewCloudWatchLogsAdapter(cwLogsClient *cloudwatchlogs.CloudWatchLogs, logGroupName, logStreamName string) CloudWatchLogsAdapter {
	return &cloudWatchLogsAdapter{
		cwLogsClient:  cwLogsClient,
		logGroupName:  logGroupName,
		logStreamName: logStreamName,
	}
}

func (a *cloudWatchLogsAdapter) PutLogEntry(logEntry *domain.LogEntry) error {
	logMessage := formatLogEntry(logEntry)
	input := &cloudwatchlogs.PutLogEventsInput{
		LogGroupName:  aws.String(a.logGroupName),
		LogStreamName: aws.String(a.logStreamName),
		LogEvents: []*cloudwatchlogs.InputLogEvent{
			{
				Message:   aws.String(logMessage),
				Timestamp: aws.Int64(logEntry.Timestamp.UnixNano() / int64(time.Millisecond)),
			},
		},
	}

	_, err := a.cwLogsClient.PutLogEvents(input)
	return err
}

func formatLogEntry(logEntry *domain.LogEntry) string {
	return fmt.Sprintf("File Path: %s | Bucket: %s | Total Downloads: %d | Total Data Transfer: %d bytes",
		logEntry.FilePath, logEntry.Bucket, logEntry.DownloadCount, logEntry.DataTransfer)
}
