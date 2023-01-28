package handler

import (
	"context"
	"douyin/biz/pb"
	"douyin/biz/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"log"
	"strconv"
)

func PublishList(ctx context.Context, c *app.RequestContext) {
	var req pb.DouyinPublishListRequest

	id, err := strconv.Atoi(c.Query("user_id"))
	if err != nil {
		log.Println("get parameter error: ", err)
		return
	}
	req.UserId = int64(id)
	req.Token = c.Query("token")

	resp, err := service.Svr.GetPublishList(ctx, &req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}
