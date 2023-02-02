package service

import (
	"context"
	"douyin/biz/pb"
)

func (s *Service) GetPublishList(ctx context.Context, req *pb. DouyinPublishListRequest ) (resp *pb. DouyinPublishListResponse, err error) {

	var videos []*pb.Video
	
	//todo

	resp = &pb. DouyinPublishListResponse{}
	resp.StatusCode = 0
	resp.VideoList = videos
	
	return 
}