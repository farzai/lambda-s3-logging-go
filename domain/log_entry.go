package domain

import "time"

type LogEntry struct {
	FilePath      string
	Bucket        string
	DownloadCount int64
	DataTransfer  int64
	Timestamp     time.Time
}
