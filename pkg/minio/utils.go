package minio

import (
	"context"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/minio/minio-go/v7"
	"io"
	"net/url"
	"time"
)

// CreateBucket 在服务器创建视频桶用于保存视频
//
//	参照文档 https://github.com/minio/minio-go
//
// /**
func CreateBucket(bucketName string) error {
	if len(bucketName) <= 0 {
		hlog.Debug("bucketName invalid")
	}

	location := "guangdong"
	ctx := context.Background()

	err := minioClient.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{Region: location})
	if err != nil {
		exists, errBucketExists := minioClient.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			hlog.Debugf("bucket %s already exists", bucketName)
			return nil
		} else {
			return err
		}
	} else {
		hlog.Infof("bucket %s create successfully", bucketName)
	}
	return nil
}

// UploadLocalFile 将filePath指定的文件保存到视频桶中，并命名为objectName
//
// /**
func UploadLocalFile(bucketName string, objectName string, filePath string, contentType string) (int64, error) {
	ctx := context.Background()
	info, err := minioClient.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		hlog.Errorf("localfile upload failed,%s", err)
		return 0, err
	}
	hlog.Infof("upload %s of size %d successfully", objectName, info.Size)
	return info.Size, nil
}

// UploadFile 将io流中包含的数据保存到本地，并命名为objectName
//
// /**
func UploadFile(bucketName string, objectName string, reader io.Reader, objectSize int64) error {
	ctx := context.Background()
	n, err := minioClient.PutObject(ctx, bucketName, objectName, reader, objectSize, minio.PutObjectOptions{ContentType: "application/octet-stream"})
	if err != nil {
		hlog.Errorf("upload %s of size %d failed,%s", bucketName, objectSize, err)
		return err
	}
	hlog.Infof("upload %s of bytes %d successfully", objectName, n.Size)
	return nil
}

// GetFileUrl 根据fileName获取视频桶中保存的文件路径
//
// /**
func GetFileUrl(bucketName string, fileName string, expires time.Duration) (*url.URL, error) {
	ctx := context.Background()
	reqParams := make(url.Values)
	if expires <= 0 {
		expires = time.Second * 60 * 60 * 24
	}
	presignedUrl, err := minioClient.PresignedGetObject(ctx, bucketName, fileName, expires, reqParams)
	if err != nil {
		hlog.Errorf("get url of file %s from bucket %s failed,%s", fileName, bucketName, err)
		return nil, err
	}

	return presignedUrl, nil
}
