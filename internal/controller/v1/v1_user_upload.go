package v1

import (
	"context"
	"fmt"
	"gf_study/api/v1/user"
	"github.com/gogf/gf/v2/errors/gerror"
)

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
