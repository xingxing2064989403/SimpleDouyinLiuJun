package common

const DefaultPass = "liujun"

const JwtSecret = "liujun"
const JwtExpire = 7200
const OrganizationName = " SixSteeds"
const DefaultBackImage = " https://staraway.love/02c05856-e60c-44a0-9e85-6c2ba7ae9e0942d98a82be40d551ce6e123abbd0718292f80a71_raw.jpg"
const DefaultAvatar = "https://staraway.love/031ec513-8976-45b5-80fd-8c725bc7ada7u%3D2169083367%2C64951360%26fm%3D253%26fmt%3Dauto%26app%3D138%26f%3DJPEG.webp"

/*
redis缓存前缀
*/

// 计数服务缓存
const CntCacheUserLikePrefix = "CntCache:user_like_cnt:"
const CntCacheVideoLikedPrefix = "CntCache:video_liked_cnt:"
const CntCacheVideoCommentedPrefix = "CntCache:video_commented_cnt:"

// 点赞记录缓存
const LikeCacheVideoLikedPrefix = "LikeCache:video_liked:"
