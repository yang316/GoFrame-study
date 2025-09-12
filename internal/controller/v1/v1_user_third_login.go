package v1

import (
	"context"
	"encoding/json"
	"fmt"
	"gf_study/api/v1/user"
	"gf_study/internal/model/entity"
	"github.com/gogf/gf/v2/errors/gerror"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
)

func (c *ControllerUser) ThirdLogin(ctx context.Context, req *user.ThirdLoginReq) (res *user.ThirdLoginRes, err error) {

	switch req.Platform {
	case "wechat":
		res, err = wechatMPLogin(req.Code)
		return
	default:
		res, err = wechatMPLogin(req.Code)
		return
	}
}

// 微信授权登录
func wechatMPLogin(code string) (*user.ThirdLoginRes, error) {
	// 添加请求参数
	params := url.Values{}
	params.Add("appid", "your_appid")   // 替换为你的微信小程序appid
	params.Add("secret", "your_secret") // 替换为你的微信小程序secret
	params.Add("js_code", code)
	params.Add("grant_type", "authorization_code")
	fullURL := fmt.Sprintf("%s?%s", "https://api.weixin.qq.com/sns/jscode2session", params.Encode())

	response, err := http.Get(fullURL)
	if err != nil {
		fmt.Printf("请求微信API失败: %v\n", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("关闭响应体失败: %v\n", err)
		}
	}(response.Body)

	// 读取响应内容
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("读取响应失败: %v\n", err)
		return nil, err
	}

	var result map[string]interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Printf("JSON解析失败: %v\n", err)
		return nil, err
	}

	if result["errcode"] != nil {
		errMsg := fmt.Sprintf("微信API错误: %v", result["errmsg"])
		//fmt.Println(errMsg)
		return nil, gerror.New(errMsg)
	}

	fmt.Println("解析后的JSON内容:", result)
	// 根据 ThirdLoginRes 的结构定义调整返回值
	gfUser := &entity.GfUsers{
		// 假设 GfUsers 有 OpenID 和 SessionKey 字段
		OpenID: result["openid"].(string),
		//SessionKey: result["session_key"].(string),
	}

	return &user.ThirdLoginRes{
		Info: gfUser,
	}, nil
}
