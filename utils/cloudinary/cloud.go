package cloudinary

import (
	c "MiniProject/configs"
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryInterface interface {
	ImageUploadHelper(input interface{}) (string, error)
	FileUploadHelper(input interface{}) (string, error)
}

func ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(c.EnvCloudName(), c.EnvCloudAPIKey(), c.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: c.EnvCloudUploadFolder()})
	if err != nil {
		return "", nil
	}

	return uploadParam.SecureURL, nil
}

func FileUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(c.EnvCloudName(), c.EnvCloudAPIKey(), c.EnvCloudAPISecret())
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: c.EnvCloudUploadFolder(), Format: "png"})
	if err != nil {
		return "", nil
	}

	return uploadParam.SecureURL, nil
}
