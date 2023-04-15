package adapter

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

const DownloadCountTagKey = "DownloadCount"

type S3Adapter interface {
	IncrementDownloadCount(bucket, key string) (int64, int64, error)
}

type s3Adapter struct {
	s3Client *s3.S3
}

func NewS3Adapter(s3Client *s3.S3) S3Adapter {
	return &s3Adapter{s3Client: s3Client}
}

func (a *s3Adapter) IncrementDownloadCount(bucket, key string) (int64, int64, error) {
	// Get the existing download count from the object's tags
	taggingOutput, err := a.s3Client.GetObjectTagging(&s3.GetObjectTaggingInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return 0, 0, err
	}

	downloadCount := int64(0)

	for _, tag := range taggingOutput.TagSet {
		if aws.StringValue(tag.Key) == DownloadCountTagKey {
			downloadCount, _ = strconv.ParseInt(aws.StringValue(tag.Value), 10, 64)
			break
		}
	}

	// Increment the download count
	downloadCount++

	// Update the object's tags with the new download count
	_, err = a.s3Client.PutObjectTagging(&s3.PutObjectTaggingInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Tagging: &s3.Tagging{
			TagSet: []*s3.Tag{
				{
					Key:   aws.String(DownloadCountTagKey),
					Value: aws.String(strconv.FormatInt(downloadCount, 10)),
				},
			},
		},
	})
	if err != nil {
		return 0, 0, err
	}

	// Get the size of the object
	headObjectOutput, err := a.s3Client.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return 0, 0, err
	}

	size := aws.Int64Value(headObjectOutput.ContentLength)

	return downloadCount, size, nil
}
