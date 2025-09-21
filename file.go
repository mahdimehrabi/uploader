package uploader

import (
	"context"
	"io"
)

type FileRepository interface {
	UploadPublicFile(ctx context.Context, bucketName, objectName, contentType string, file io.Reader, metadata map[string]string) error
	DeleteFile(ctx context.Context, bucketName string, objectName string) error
}
