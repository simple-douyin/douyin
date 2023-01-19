package service

import (
	"context"
	"douyin/biz/pb"
)

func (s *Service) RegisterUser(ctx context.Context, req *pb.DouyinUserRegisterRequest) (resp *pb.DouyinUserRegisterResponse, err error) {

	resp = &pb.DouyinUserRegisterResponse{}
	resp.UserId = 1
	resp.StatusCode = 0
	resp.Token = "todo"
	return
}
