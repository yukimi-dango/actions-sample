package util

import (
	"context"
	"fmt"
	"io"
	"os"
	"time"

	"cloud.google.com/go/storage"
)

type UploadRequest struct {
	Object string
	Data   io.Reader
}

func Upload(ctx context.Context, req UploadRequest) error {
	bucket := os.Getenv("MY_BUCKET")
	if bucket == "" {
		return fmt.Errorf("no bucket")
	}

	client, err := storage.NewClient(ctx)
	if err != nil {
		return fmt.Errorf("storage.NewClient: %v", err)
	}
	defer client.Close()

	ctx, cancel := context.WithTimeout(ctx, time.Second*10)
	defer cancel()

	o := client.Bucket(bucket).Object(req.Object)
	w := o.NewWriter(ctx)

	if _, err = io.Copy(w, req.Data); err != nil {
		return fmt.Errorf("io.Copy: %v", err)
	}

	return w.Close()
}
