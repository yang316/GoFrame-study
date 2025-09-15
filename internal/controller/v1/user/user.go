package user

import (
	"context"
	"encoding/json"
	"fmt"
	v1 "gf_study/api/v1"
	"gf_study/api/v1/user"
	"gf_study/internal/dao"
	"gf_study/internal/model/do"
	"gf_study/internal/model/entity"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"golang.org/x/crypto/bcrypt"
)

// ControllerUser is the controller for user-related operations
type ControllerUser struct{}

func NewUser() v1.IV1User {
	return &ControllerUser{}
}

func (c *ControllerUser) Login(ctx context.Context, req *user.LoginReq) (res *user.LoginRes, err error) {
	var userData do.GfUsers

	//查看用户名是否注册
	err = dao.GfUsers.Ctx(ctx).Where("username", req.Username).Scan(&userData)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "用户信息有误或未注册")
	}
	if gconv.Uint(userData.Id) == 0 {
		return nil, gerror.New("用户名未注册")
	}
	//查看密码是否正确
	// 解密查看密码
	checkPassword, err := DecryptPassword(gconv.String(userData.Password), req.Password)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeInternalError, err, "密码解密失败")
	}
	if req.Password != checkPassword {
		return nil, gerror.New("密码错误")
	}
	userEntity := &entity.GfUsers{
		Id:       gconv.Uint(userData.Id),
		Username: gconv.String(userData.Username),
		Nickname: gconv.String(userData.Nickname),
	}

	// 构建返回结果（注意：不要返回原始密码）
	return &user.LoginRes{
		Info: userEntity,
	}, nil
}

func (c *ControllerUser) Profile(ctx context.Context, req *user.ProfileReq) (res *user.ProfileRes, err error) {

	result, err := dao.GfUsers.Ctx(ctx).Data(do.GfUsers{
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
	}).Where("nickname", req.Nickname).Update()
	if err != nil {
		return nil, err
	}

	// 2. 检查 result 是否为 nil
	if result == nil {
		return nil, gerror.New("更新操作未返回有效结果")
	}

	// 3. 获取影响行数
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}

	// 4. 检查是否成功更新
	if rowsAffected < 1 {
		return nil, gerror.New("修改失败")
	}

	// 5. 返回成功结果
	return &user.ProfileRes{
		Result: true,
	}, nil
}

func (c *ControllerUser) Register(ctx context.Context, req *user.RegisterReq) (res *user.RegisterRes, err error) {

	// 检查用户是否已存在
	isExist, err := dao.GfUsers.Ctx(ctx).Where("username", req.Username).Exist()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "检查用户是否存在失败")
	}
	if isExist {
		return nil, gerror.Newf("用户名 %s 已存在", req.Username)
	}

	// 密码加密处理（实际项目中必须添加）
	encryptedPassword, err := EncryptPassword(req.Password)
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeOperationFailed, err, "密码加密失败")
	}

	// 插入新用户
	result, err := dao.GfUsers.Ctx(ctx).Insert(do.GfUsers{
		Username:      req.Username,
		Nickname:      req.Nickname,
		Password:      encryptedPassword, // 使用加密后的密码
		RegistTime:    gtime.Now(),
		LastLoginTime: gtime.Now(),
	})
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "用户注册失败")
	}

	// 获取插入的用户ID
	userId, err := result.LastInsertId()
	if err != nil {
		return nil, gerror.WrapCode(gcode.CodeDbOperationError, err, "获取用户ID失败")
	}

	// 查询刚创建的用户信息
	userEntity := &entity.GfUsers{
		Id:       uint(userId),
		Username: req.Username,
		Nickname: req.Nickname,
	}

	// 构建返回结果（注意：不要返回原始密码）
	return &user.RegisterRes{
		Info: userEntity,
	}, nil
}
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

func (c *ControllerUser) Upload(ctx context.Context, req *user.UploadReq) (res *user.UploadRes, err error) {
	if req.File == nil {
		return nil, gerror.New("no file uploaded")
	}
	filename, err := req.File.Save("./resource/uploads/")
	if err != nil {
		return nil, err
	}
	url := fmt.Sprintf("/uploads/%s", filename)
	return &user.UploadRes{
		FileName: filename,
		Url:      url,
	}, nil
}

// 加密
func EncryptPassword(password string) (string, error) {
	// 使用 bcrypt 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// 解密密码
func DecryptPassword(encryptedPassword string, password string) (string, error) {

	// 使用 bcrypt 解密密码
	err := bcrypt.CompareHashAndPassword([]byte(encryptedPassword), []byte(password))
	if err != nil {
		return "", err
	}
	return password, nil
}
