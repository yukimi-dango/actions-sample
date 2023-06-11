package util_test

import (
	"bytes"
	"context"
	"image"
	"image/jpeg"
	"os"
	"testing"

	util "github.com/yukimi-dango/actions-sample/lib"
)

func TestUpload(t *testing.T) {
	ctx := context.Background()

	baseURL := os.Getenv("GCS_URL")
	if baseURL == "" {
		t.Fatal("no gcs url")
	}

	os.Setenv("STORAGE_EMULATOR_HOST", baseURL)

	img := image.NewRGBA(image.Rect(0, 0, 256, 256))
	buf := new(bytes.Buffer)
	if err := jpeg.Encode(buf, img, nil); err != nil {
		t.Fatal(err)
	}

	if err := util.Upload(ctx, util.UploadRequest{
		Object: "sample/image.jpeg",
		Data:   buf,
	}); err != nil {
		t.Fatal(err)
	}
}
