// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	comment "doushen_by_liujun/service/content/api/internal/handler/comment"
	favorite "doushen_by_liujun/service/content/api/internal/handler/favorite"
	video "doushen_by_liujun/service/content/api/internal/handler/video"
	"doushen_by_liujun/service/content/api/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/feed",
				Handler: video.FeedHandler(serverCtx),
			},
		},
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodGet,
					Path:    "/publish/list",
					Handler: video.PublishListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: favorite.FavoriteActionHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: favorite.FavoriteListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/favorite"),
	)

	server.AddRoutes(
		rest.WithMiddlewares(
			[]rest.Middleware{serverCtx.JwtAuthMiddleware},
			[]rest.Route{
				{
					Method:  http.MethodPost,
					Path:    "/action",
					Handler: comment.CommentActionHandler(serverCtx),
				},
				{
					Method:  http.MethodGet,
					Path:    "/list",
					Handler: comment.CommentListHandler(serverCtx),
				},
			}...,
		),
		rest.WithPrefix("/douyin/comment"),
	)
}
