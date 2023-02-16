package services

import (
	"github.com/juliotorresmoreno/SpecialistTalk/configs"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

func NewClient() (*minio.Client, error) {
	conf := configs.GetConfig().Minio
	endpoint := conf.Endpoint
	accessKeyID := conf.AccessKeyID
	secretAccessKey := conf.SecretAccessKey
	useSSL := conf.UseSSL

	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL,
	})

	return minioClient, err
}
