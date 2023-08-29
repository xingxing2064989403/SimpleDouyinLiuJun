// Code generated by goctl. DO NOT EDIT.

package model

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/zeromicro/go-zero/core/stores/builder"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/core/stores/sqlc"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/core/stringx"
)

var (
	favoriteFieldNames          = builder.RawFieldNames(&Favorite{})
	favoriteRows                = strings.Join(favoriteFieldNames, ",")
	favoriteRowsExpectAutoSet   = strings.Join(stringx.Remove(favoriteFieldNames, "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), ",")
	favoriteRowsWithPlaceHolder = strings.Join(stringx.Remove(favoriteFieldNames, "`id`", "`create_at`", "`create_time`", "`created_at`", "`update_at`", "`update_time`", "`updated_at`"), "=?,") + "=?"

	cacheLiujunContentFavoriteIdPrefix = "cache:liujunContent:favorite:id:"
	cacheLiujunContentFavoriteUserIdVideoIdPrefix = "cache:liujunContent:favorite:userid:videoid:"
	cacheLiujunContentFavoriteListUserIdPrefix = "cache:liujunContent:favoriteList:userid:"
)

type (
	favoriteModel interface {
		Insert(ctx context.Context, data *Favorite) (sql.Result, error)
		FindOne(ctx context.Context, id int64) (*Favorite, error)
		Update(ctx context.Context, data *Favorite) error
		Delete(ctx context.Context, id int64) error
		FindFavoriteByUserIdVideoId(ctx context.Context, userid int64, videoid int64) (*Favorite, error)
		FindFavoriteListByUserId(ctx context.Context, userid int64) (*[]*Favorite, error)
		FindFavoriteListByVideoId(ctx context.Context, videoid int64) (*[]*Favorite, error)
		FindFavoritedCntByVideoIdList(ctx context.Context, videoIdList *[]int64) (int64, error)
		GetFavoriteCountByUserId(ctx context.Context, user_id int64) (*Count, error)
	}

	defaultFavoriteModel struct {
		sqlc.CachedConn
		table string
	}

	Favorite struct {
		Id         int64     `db:"id"`          // 主键
		VideoId    int64     `db:"video_id"`    // 视频id
		UserId     int64     `db:"user_id"`     // 视频id
		CreateTime time.Time `db:"create_time"` // 该条记录创建时间
		UpdateTime time.Time `db:"update_time"` // 该条最后一次更新时间
		IsDelete   int64     `db:"is_delete"`   // 逻辑删除
	}
)

func newFavoriteModel(conn sqlx.SqlConn, c cache.CacheConf, opts ...cache.Option) *defaultFavoriteModel {
	return &defaultFavoriteModel{
		CachedConn: sqlc.NewConn(conn, c, opts...),
		table:      "`favorite`",
	}
}

func (m *defaultFavoriteModel) withSession(session sqlx.Session) *defaultFavoriteModel {
	return &defaultFavoriteModel{
		CachedConn: m.CachedConn.WithSession(session),
		table:      "`favorite`",
	}
}

func (m *defaultFavoriteModel) Delete(ctx context.Context, id int64) error {
	liujunContentFavoriteIdKey := fmt.Sprintf("%s%v", cacheLiujunContentFavoriteIdPrefix, id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("delete from %s where `id` = ?", m.table)
		return conn.ExecCtx(ctx, query, id)
	}, liujunContentFavoriteIdKey)
	return err
}

func (m *defaultFavoriteModel) FindOne(ctx context.Context, id int64) (*Favorite, error) {
	liujunContentFavoriteIdKey := fmt.Sprintf("%s%v", cacheLiujunContentFavoriteIdPrefix, id)
	var resp Favorite
	err := m.QueryRowCtx(ctx, &resp, liujunContentFavoriteIdKey, func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", favoriteRows, m.table)
		return conn.QueryRowCtx(ctx, v, query, id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFavoriteModel) Insert(ctx context.Context, data *Favorite) (sql.Result, error) {
	liujunContentFavoriteIdKey := fmt.Sprintf("%s%v", cacheLiujunContentFavoriteIdPrefix, data.Id)
	ret, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("insert into %s (%s) values (?, ?, ?, ?)", m.table, favoriteRowsExpectAutoSet)
		return conn.ExecCtx(ctx, query, data.Id, data.VideoId, data.UserId, data.IsDelete)
	}, liujunContentFavoriteIdKey)
	return ret, err
}

func (m *defaultFavoriteModel) Update(ctx context.Context, data *Favorite) error {
	liujunContentFavoriteIdKey := fmt.Sprintf("%s%v", cacheLiujunContentFavoriteIdPrefix, data.Id)
	_, err := m.ExecCtx(ctx, func(ctx context.Context, conn sqlx.SqlConn) (result sql.Result, err error) {
		query := fmt.Sprintf("update %s set %s where `id` = ?", m.table, favoriteRowsWithPlaceHolder)
		return conn.ExecCtx(ctx, query, data.VideoId, data.UserId, data.IsDelete, data.Id)
	}, liujunContentFavoriteIdKey)
	return err
}

//根据（用户id，视频id）字段查找favorite表
func (m *defaultFavoriteModel) FindFavoriteByUserIdVideoId(ctx context.Context, userid int64, videoid int64) (*Favorite, error) {
	var resp Favorite
	query := fmt.Sprintf("select %s from %s where `user_id` = ? and `video_id` = ? limit 1", favoriteRows, m.table)
	err := m.QueryRowNoCacheCtx(ctx, &resp, query, userid, videoid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFavoriteModel) FindFavoriteListByUserId(ctx context.Context, userid int64) (*[]*Favorite, error) {
	var resp []*Favorite
	query := fmt.Sprintf("select %s from %s where `user_id` = ? ", favoriteRows,  m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, userid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFavoriteModel) FindFavoriteListByVideoId(ctx context.Context, videoid int64) (*[]*Favorite, error) {
	var resp []*Favorite
	query := fmt.Sprintf("select %s from %s where `video_id` = ? ", favoriteRows,  m.table)
	err := m.QueryRowsNoCacheCtx(ctx, &resp, query, videoid)
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}

func (m *defaultFavoriteModel) FindFavoritedCntByVideoIdList(ctx context.Context, videoIdList *[]int64) (int64, error) {
	var favoritedCnt int64
	// []int64 转换为 “,” 分隔的 string
	var str_arr = make([]string, len(*videoIdList))
	for k, v := range *videoIdList {
		str_arr[k] = fmt.Sprintf("%d", v)
	}
	var IdListStr = strings.Join(str_arr, ",")
	query := fmt.Sprintf("select count(*) from %s where `video_id` in (%s) and `is_delete`!= '1'", m.table, IdListStr)
	err := m.QueryRowNoCacheCtx(ctx, &favoritedCnt, query)
	switch err {
	case nil:
		return favoritedCnt, nil
	case sqlc.ErrNotFound:
		return 0, ErrNotFound
	default:
		return 0, err
	}
}


func (m *defaultFavoriteModel) formatPrimary(primary any) string {
	return fmt.Sprintf("%s%v", cacheLiujunContentFavoriteIdPrefix, primary)
}

func (m *defaultFavoriteModel) queryPrimary(ctx context.Context, conn sqlx.SqlConn, v, primary any) error {
	query := fmt.Sprintf("select %s from %s where `id` = ? limit 1", favoriteRows, m.table)
	return conn.QueryRowCtx(ctx, v, query, primary)
}

func (m *defaultFavoriteModel) tableName() string {
	return m.table
}

func (m *defaultFavoriteModel) GetFavoriteCountByUserId(ctx context.Context, user_id int64) (*Count, error) {
	var resp Count
	err := m.QueryRowCtx(ctx, &resp, "GetFavoriteCountByUserId--"+strconv.Itoa(int(user_id)), func(ctx context.Context, conn sqlx.SqlConn, v any) error {
		query := fmt.Sprintf("select COUNT(*) as count from favorite where `user_id` = ?")
		return conn.QueryRowCtx(ctx, v, query, user_id)
	})
	switch err {
	case nil:
		return &resp, nil
	case sqlc.ErrNotFound:
		return nil, ErrNotFound
	default:
		return nil, err
	}
}