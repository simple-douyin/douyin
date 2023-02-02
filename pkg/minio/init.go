package minio

import (
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// 初始化minio

var (
	minioClient          *minio.Client
	MinioEndpoint        = "192.168.0.247:9000" // 根据自己本地环境进行修改
	MinioAccessKeyId     = "minioadmin"         //运行minio后会提供
	MinioSecretAccessKey = "minioadmin"         //运行minio后会提供
	MinioUseSSL          = false
	MinioVideoBucketName = "douyin-video" //创建存放视频的目录
)

func init() {
	client, err := minio.New(MinioEndpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(MinioAccessKeyId, MinioSecretAccessKey, ""),
		Secure: MinioUseSSL,
	})

	if err != nil {
		hlog.Debug("minio client init failed:%v", err)
	}

	hlog.Debug("minio client init successfully")
	minioClient = client
	if err := CreateBucket(MinioVideoBucketName); err != nil {
		hlog.Error("minio client init failed:%v", err)
	}
}
