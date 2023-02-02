package service

import (
	"bytes"
	"context"
	"douyin/biz/pb"
	"douyin/pkg/minio"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"strconv"
	"strings"
	"time"
)

// CreateVideo 使用minio将传输视频保存到本地
// /
func (s *Service) CreateVideo(ctx context.Context, req *pb.DouyinPublishActionRequest) (resp *pb.DouyinPublishActionResponse, err error) {
	resp = &pb.DouyinPublishActionResponse{}
	MinioVideoBucketName := minio.MinioVideoBucketName
	hlog.Infof("bucket name:%s", MinioVideoBucketName)

	videoData := req.Data

	reader := bytes.NewReader(videoData)

	fileName := req.Title + strconv.FormatInt(time.Now().Unix(), 10) + "." + "mp4"

	err = minio.UploadFile(MinioVideoBucketName, fileName, reader, int64(len(videoData)))
	if err != nil {
		msg := err.Error()
		resp.StatusCode = 1
		resp.StatusMsg = &msg
	}

	url, err := minio.GetFileUrl(MinioVideoBucketName, fileName, 0)

	hlog.Infof("url:%s", url)
	playUrl := strings.Split(url.String(), "?")[0]
	if err != nil {
		msg := err.Error()
		resp.StatusCode = 1
		resp.StatusMsg = &msg
	}

	msg := "title:" + req.Title + " ,filename:" + fileName + " ,playurl:" + playUrl

	resp.StatusCode = 0
	resp.StatusMsg = &msg
	//TODO 保存视频路径和文件名到数据库

	return

}
