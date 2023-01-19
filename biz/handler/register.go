package handler

import (
	"context"
	"douyin/biz/pb"
	"douyin/biz/service"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func Register(ctx context.Context, c *app.RequestContext) {
	//todo
	var req pb.DouyinUserRegisterRequest
	req.Username = c.Query("username")
	req.Password = c.Query("password")
	fmt.Println("test")
	if len(req.Username) == 0 || len(req.Password) == 0 {
		c.JSON(consts.StatusBadRequest, "1 ")
		return
	}

	resp, err := service.Svr.RegisterUser(ctx, &pb.DouyinUserRegisterRequest{
		Username: req.Username,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(consts.StatusBadRequest, resp)
		return
	}
	c.JSON(consts.StatusOK, resp)

}
