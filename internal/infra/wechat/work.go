package wechat

import (
	"context"
	"fmt"
	"github.com/haysons/eino-bot/config"
	"github.com/haysons/eino-bot/internal/utils"
	"github.com/haysons/gokit/errors"
	"github.com/redis/go-redis/v9"
	"time"
)

const (
	// urlGetToken 企微获取access_token，使用get请求
	urlGetToken = "https://qyapi.weixin.qq.com/cgi-bin/gettoken?corpid=%s&corpsecret=%s"
)

// Work 企业微信相关接口
type Work struct {
	conf  *config.WechatConfig
	cache *redis.Client
}

func NewWork(conf *config.WechatConfig, cache *redis.Client) *Work {
	return &Work{
		conf:  conf,
		cache: cache,
	}
}

// GetAccessToken 获取企微接口所需的access_token
func (w *Work) GetAccessToken(ctx context.Context) (string, error) {
	// 自缓存中获取
	key := fmt.Sprintf("eino-bot:access_token:%s:%s", w.conf.CorpID, w.conf.AgentID)
	tokenCache, err := w.cache.Get(ctx, key).Result()
	if err != nil && !errors.Is(err, redis.Nil) {
		return "", errors.Wrap(err, "get access token from cache failed")
	}
	if err == nil && tokenCache != "" {
		return tokenCache, nil
	}
	// 调用接口获取
	tokenRes, err := utils.HTTPGet[AccessTokenResp](fmt.Sprintf(urlGetToken, w.conf.CorpID, w.conf.CorpSecret))
	if err != nil {
		return "", err
	}
	if tokenRes.ErrCode != 0 {
		return "", errors.Newf("get access_token failed, %d, %s", tokenRes.ErrCode, tokenRes.ErrMsg)
	}
	// 缓存access_token
	expireSecs := tokenRes.ExpiresIn - 600
	if err = w.cache.Set(ctx, key, tokenRes.AccessToken, time.Duration(expireSecs)*time.Second).Err(); err != nil {
		return "", errors.Wrap(err, "set access_token to cache failed")
	}
	return tokenRes.AccessToken, nil
}

// SendMsg 发送应用消息
func (w *Work) SendMsg(ctx context.Context) error {
	return nil
}
