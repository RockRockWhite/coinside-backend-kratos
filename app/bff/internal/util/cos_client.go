package util

import (
	"fmt"
	"github.com/ljxsteam/coinside-backend-kratos/pkg/config"
	"github.com/tencentyun/cos-go-sdk-v5"
	"github.com/tencentyun/cos-go-sdk-v5/debug"
	"net/http"
	"net/url"
)

func NewCOSClinet(config *config.Config) *cos.Client {
	r := fmt.Sprintf("https://%s.cos.ap-%s.myqcloud.com", config.GetString("cos.bucket"), config.GetString("cos.region"))
	u, _ := url.Parse(r)
	b := &cos.BaseURL{BucketURL: u}
	return cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  config.GetString("cos.secret_id"),
			SecretKey: config.GetString("cos.secret_key"),
			Transport: &debug.DebugRequestTransport{
				RequestHeader:  true,
				RequestBody:    false,
				ResponseHeader: true,
				ResponseBody:   false,
			},
		},
	})
}
