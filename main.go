package main

import (
	"context"
	"log"
	"os"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/farzai/lambda-s3-logging-go/adapter"
	"github.com/farzai/lambda-s3-logging-go/domain"
	"github.com/farzai/lambda-s3-logging-go/usecase"
)

func handleRequest(ctx context.Context, s3Event events.S3Event) {
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String(os.Getenv("AWS_REGION")),
	})

	if err != nil {
		log.Fatalf("Failed to create a session: %v", err)
	}

	s3Client := s3.New(sess)
	s3Adapter := adapter.NewS3Adapter(s3Client)

	cwLogsClient := cloudwatchlogs.New(sess)
	cwLogsAdapter := adapter.NewCloudWatchLogsAdapter(cwLogsClient, "<your-log-group-name>", "<your-log-stream-name>")

	logUsecase := usecase.NewLogUsecase(cwLogsAdapter)

	for _, record := range s3Event.Records {
		bucket := record.S3.Bucket.Name
		key := record.S3.Object.Key

		downloadCount, size, err := s3Adapter.IncrementDownloadCount(bucket, key)
		if err != nil {
			log.Printf("Error processing record: %v", err)
			continue
		}

		logEntry := domain.LogEntry{
			FilePath:      key,
			Bucket:        bucket,
			DownloadCount: downloadCount,
			DataTransfer:  size,
		}

		err = logUsecase.LogDataTransfer(&logEntry)
		if err != nil {
			log.Printf("Error logging data transfer: %v", err)
		}
	}
}

func main() {
	lambda.Start(handleRequest)
}
