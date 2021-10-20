package kit

import (
	"net/http"
	"net/url"

	"github.com/ahmek/kit/types"
)

// HTTPContext 请求信息封装
type HTTPContext struct {
	w http.ResponseWriter
	r *http.Request
	// uid 用户登录态
	user interface{}
	// 登录态验证类型 [0.不验证; 1.验证登录; 2.验证且必须为管理员]
	authType int8
}

func NewHTTPContext(w http.ResponseWriter, r *http.Request) *HTTPContext {
	return &HTTPContext{
		w: w,
		r: r,
	}
}

func (c *HTTPContext) PrintErr(err error) {
	c.w.Write(types.GetErrorJSON(err, nil))
}

// GetRequest 获取请求头请求数据
func (ctx *HTTPContext) GetRequest() *http.Request {
	return ctx.r
}

// GetResponse 获取请求头返回结果
func (ctx *HTTPContext) GetResponse() http.ResponseWriter {
	return ctx.w
}

// GetURLPath 获取当前路径
func (ctx *HTTPContext) GetURLPath() string {
	return ctx.r.URL.Path
}

// GetURL 获取当前url
func (ctx *HTTPContext) GetURL() *url.URL {
	return ctx.r.URL
}

// Auth 验证登录情况
func (c *HTTPContext) Auth(token tokenCBK, isAdmin adminCBK) error {
	user := token(c.r)
	if user == nil || c.authType == 1 {
		return types.ErrorTokenInvalid
	}
	c.user = user
	if c.authType == 2 {
		if err := isAdmin(c.r); err != nil {
			return err
		}
	}
	return nil
}

// SetAuthType 设置登录权限
// [0.不验证登录; 1.验证; 2.验证且必须是管理员]
func (ctx *HTTPContext) SetAuthType(t int8) {
	ctx.authType = t
}

// SetAuthType 获取登录态类型
func (ctx *HTTPContext) GetAuthType() int8 {
	return ctx.authType
}

// GetUser 获取登录态当前用户id
func (ctx *HTTPContext) GetUser() interface{} {
	return ctx.user
}

// GetUid 获取登录态当前用户id
func (ctx *HTTPContext) GetUid() int64 {
	// return ctx.user.Id
	return 0
}
