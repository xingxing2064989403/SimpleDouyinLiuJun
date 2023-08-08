// Code generated by goctl. DO NOT EDIT.
// Source: user.proto

package server

import (
	"context"

	"doushen_by_liujun/service/user/rpc/internal/logic"
	"doushen_by_liujun/service/user/rpc/internal/svc"
	"doushen_by_liujun/service/user/rpc/pb"
)

type UserServer struct {
	svcCtx *svc.ServiceContext
	pb.UnimplementedUserServer
}

func NewUserServer(svcCtx *svc.ServiceContext) *UserServer {
	return &UserServer{
		svcCtx: svcCtx,
	}
}

// -----------------------鐢ㄦ埛鍩烘湰淇℃伅-----------------------
func (s *UserServer) AddFollows(ctx context.Context, in *pb.AddFollowsReq) (*pb.AddFollowsResp, error) {
	l := logic.NewAddFollowsLogic(ctx, s.svcCtx)
	return l.AddFollows(in)
}

func (s *UserServer) UpdateFollows(ctx context.Context, in *pb.UpdateFollowsReq) (*pb.UpdateFollowsResp, error) {
	l := logic.NewUpdateFollowsLogic(ctx, s.svcCtx)
	return l.UpdateFollows(in)
}

func (s *UserServer) DelFollows(ctx context.Context, in *pb.DelFollowsReq) (*pb.DelFollowsResp, error) {
	l := logic.NewDelFollowsLogic(ctx, s.svcCtx)
	return l.DelFollows(in)
}

func (s *UserServer) GetFollowsById(ctx context.Context, in *pb.GetFollowsByIdReq) (*pb.GetFollowsByIdResp, error) {
	l := logic.NewGetFollowsByIdLogic(ctx, s.svcCtx)
	return l.GetFollowsById(in)
}

func (s *UserServer) GetFollowersById(ctx context.Context, in *pb.GetFollowersByIdReq) (*pb.GetFollowersByIdResp, error) {
	l := logic.NewGetFollowersByIdLogic(ctx, s.svcCtx)
	return l.GetFollowersById(in)
}

func (s *UserServer) GetFollowsCountById(ctx context.Context, in *pb.GetFollowsCountByIdReq) (*pb.GetFollowsCountByIdResp, error) {
	l := logic.NewGetFollowsCountByIdLogic(ctx, s.svcCtx)
	return l.GetFollowsCountById(in)
}

func (s *UserServer) GetFollowersCountById(ctx context.Context, in *pb.GetFollowersCountByIdReq) (*pb.GetFollowersCountByIdResp, error) {
	l := logic.NewGetFollowersCountByIdLogic(ctx, s.svcCtx)
	return l.GetFollowersCountById(in)
}

func (s *UserServer) CheckIsFollow(ctx context.Context, in *pb.CheckIsFollowReq) (*pb.CheckIsFollowResp, error) {
	l := logic.NewCheckIsFollowLogic(ctx, s.svcCtx)
	return l.CheckIsFollow(in)
}

func (s *UserServer) GetFriendsById(ctx context.Context, in *pb.GetFriendsByIdReq) (*pb.GetFriendsByIdResp, error) {
	l := logic.NewGetFriendsByIdLogic(ctx, s.svcCtx)
	return l.GetFriendsById(in)
}

func (s *UserServer) SearchFollows(ctx context.Context, in *pb.SearchFollowsReq) (*pb.SearchFollowsResp, error) {
	l := logic.NewSearchFollowsLogic(ctx, s.svcCtx)
	return l.SearchFollows(in)
}

// -----------------------鐢ㄦ埛鍩烘湰淇℃伅-----------------------
func (s *UserServer) AddUserinfo(ctx context.Context, in *pb.AddUserinfoReq) (*pb.AddUserinfoResp, error) {
	l := logic.NewAddUserinfoLogic(ctx, s.svcCtx)
	return l.AddUserinfo(in)
}

func (s *UserServer) UpdateUserinfo(ctx context.Context, in *pb.UpdateUserinfoReq) (*pb.UpdateUserinfoResp, error) {
	l := logic.NewUpdateUserinfoLogic(ctx, s.svcCtx)
	return l.UpdateUserinfo(in)
}

func (s *UserServer) DelUserinfo(ctx context.Context, in *pb.DelUserinfoReq) (*pb.DelUserinfoResp, error) {
	l := logic.NewDelUserinfoLogic(ctx, s.svcCtx)
	return l.DelUserinfo(in)
}

func (s *UserServer) GetUserinfoById(ctx context.Context, in *pb.GetUserinfoByIdReq) (*pb.GetUserinfoByIdResp, error) {
	l := logic.NewGetUserinfoByIdLogic(ctx, s.svcCtx)
	return l.GetUserinfoById(in)
}

func (s *UserServer) SearchUserinfo(ctx context.Context, in *pb.SearchUserinfoReq) (*pb.SearchUserinfoResp, error) {
	l := logic.NewSearchUserinfoLogic(ctx, s.svcCtx)
	return l.SearchUserinfo(in)
}
