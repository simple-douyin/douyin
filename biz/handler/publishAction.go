package handler

import (
	"context"
	"douyin/biz/pb"
	"douyin/biz/service"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/hlog"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

func PublishAction(ctx context.Context, c *app.RequestContext) {
	//TODO 鉴权
	var req pb.DouyinPublishActionRequest
	userid := c.Query("user_id")
	req.Token = c.Query("token")
	req.Title = c.Query("title")
	req.Data = []byte(c.Query("data"))

	hlog.Infof("handler:\nuser_id:%s,token:%s,title:%s", userid, req.Token, req.Title)
	resp, err := service.Svr.CreateVideo(ctx, &req)
	if err != nil {
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	c.JSON(consts.StatusOK, resp)
}
