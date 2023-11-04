package cloudinary

import (
	"MiniProject/configs"
	"context"
	"time"

	"github.com/cloudinary/cloudinary-go/v2"
	"github.com/cloudinary/cloudinary-go/v2/api/uploader"
)

type CloudinaryInterface interface {
	ImageUploadHelper(input interface{}) (string, error)
	FileUploadHelper(input interface{}) (string, error)
}

type Cloudinary struct {
	c configs.ProgramConfig
}

func InitCloud(cfg configs.ProgramConfig) CloudinaryInterface {
	return &Cloudinary{
		c: cfg,
	}
}

func (cl *Cloudinary) ImageUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(cl.c.CloudName, cl.c.CloudAPIKey, cl.c.CloudAPISecret)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: cl.c.CloudFolderName})
	if err != nil {
		return "", nil
	}

	return uploadParam.SecureURL, nil
}

func (cl *Cloudinary) FileUploadHelper(input interface{}) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cld, err := cloudinary.NewFromParams(cl.c.CloudName, cl.c.CloudAPIKey, cl.c.CloudAPISecret)
	if err != nil {
		return "", err
	}

	uploadParam, err := cld.Upload.Upload(ctx, input, uploader.UploadParams{Folder: cl.c.CloudFolderName, Format: "png"})
	if err != nil {
		return "", nil
	}

	return uploadParam.SecureURL, nil
}
