package router

import (
	"bytes"
	"crypto/md5"
	"fmt"
	cache "github.com/chenyahui/gin-cache"
	"github.com/chenyahui/gin-cache/persist"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/controller"
	"github.com/ljxsteam/coinside-backend-kratos/app/bff/internal/middleware"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
	"io"
	"net/http"
	"time"
)

func NewApiRouter(conf *config.Config, userController *controller.UserController, cardController *controller.CardController, attachmentController *controller.AttachmentController, teamController *controller.TeamController, voteController *controller.VoteController, todoController *controller.TodoController, markdownController *controller.MarkdownController, objectController *controller.ObjectController) *gin.Engine {
	// 初始化Controllers
	router := gin.Default()

	// 配置跨域中间件
	router.Use(middleware.Cors)

	// 配置缓存中间件
	redisStore := persist.NewRedisStore(redis.NewClient(&redis.Options{
		Network: conf.GetString("redis.network"),
		Addr:    conf.GetString("redis.addr"),
	}))

	router.Use(cache.Cache(
		redisStore,
		time.Minute*10,
		cache.WithCacheStrategyByRequest(func(c *gin.Context) (bool, cache.Strategy) {
			hashStr := ""
			if c.Request.Method != http.MethodGet && c.Request.Method != http.MethodDelete {
				body, err := io.ReadAll(c.Request.Body)
				if err != nil {
					return false, cache.Strategy{}
				}

				if len(body) > 0 {
					hash := md5.Sum(body)
					hashStr = fmt.Sprintf("%x", hash)

					c.Request.Body = io.NopCloser(bytes.NewReader(body))
				}
			}

			return true, cache.Strategy{
				CacheKey:      c.Request.RequestURI + c.Request.Method + c.Request.Header.Get("Authorization") + hashStr,
				CacheDuration: time.Second * 20,
			}
		})))

	// 配置路由
	router = UserRouter(router, userController)
	router = CardRouter(router, cardController)
	router = TeamRouter(router, teamController)
	router = MarkdownRouter(router, markdownController)
	router = ObjectRouter(router, objectController)
	router = TodoRouter(router, todoController)
	router = VoteRouter(router, voteController)
	router = AttachmentRouter(router, attachmentController)

	return router
}
