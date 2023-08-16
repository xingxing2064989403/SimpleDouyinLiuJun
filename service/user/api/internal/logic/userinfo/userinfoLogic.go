package userinfo

import (
	"context"
	"doushen_by_liujun/internal/common"
	"doushen_by_liujun/internal/util"
	"doushen_by_liujun/service/user/api/internal/svc"
	"doushen_by_liujun/service/user/api/internal/types"
	"doushen_by_liujun/service/user/rpc/pb"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"strconv"
)

type UserinfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserinfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserinfoLogic {
	return &UserinfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserinfoLogic) Userinfo(req *types.UserinfoReq) (resp *types.UserinfoResp, err error) {
	logger, e := util.ParseToken(req.Token)
	fmt.Println(logger.Username)
	fmt.Println(logger.UserID)
	fmt.Println(e)
	if e != nil {
		return &types.UserinfoResp{
			StatusCode: common.TOKEN_EXPIRE_ERROR,
			StatusMsg:  "无效token",
			User:       types.User{},
		}, err
	}
	IntUserId, _ := strconv.Atoi(logger.ID)
	info, e := l.svcCtx.UserRpcClient.GetUserinfoById(l.ctx, &pb.GetUserinfoByIdReq{
		Id:     req.UserId,
		UserID: int64(IntUserId),
	})
	var user types.User
	if e != nil {
		return &types.UserinfoResp{
			StatusCode: common.DB_ERROR,
			StatusMsg:  e.Error(),
			User:       user,
		}, err
	}
	user = types.User{
		UserId:          info.Userinfo.Id,
		Name:            info.Userinfo.Username,
		FollowCount:     info.Userinfo.FollowCount,
		FollowerCount:   info.Userinfo.FollowerCount,
		IsFollow:        info.Userinfo.IsFollow, //我对这个的理解就是当前用户对这条数据的用户是否关注
		Avatar:          info.Userinfo.Avatar,
		BackgroundImage: info.Userinfo.BackgroundImage,
		Signature:       info.Userinfo.Signature,
		WorkCount:       0, //查表
		FavoriteCount:   0, //查表
		TotalFavorited:  0, //查表
	}
	return &types.UserinfoResp{
		StatusCode: common.OK,
		StatusMsg:  "查询成功",
		User:       user,
	}, nil
}
