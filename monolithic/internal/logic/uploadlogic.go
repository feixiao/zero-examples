package logic

import (
	"context"
	"io"
	"net/http"
	"os"
	"path"

	"monolithic/internal/svc"

	"github.com/zeromicro/go-zero/core/logx"
)

const maxFileSize = 10 << 20 // 10 MB

type UploadLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUploadLogic(ctx context.Context, svcCtx *svc.ServiceContext) UploadLogic {
	return UploadLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UploadLogic) Upload(svc *svc.ServiceContext, w http.ResponseWriter, r *http.Request) {
	r.ParseMultipartForm(maxFileSize)
	file, handler, err := r.FormFile("myFile")
	if err != nil {
		logx.Error(err)
		return
	}
	defer file.Close()

	logx.Infof("upload file: %+v, file size: %d, MIME header: %+v",
		handler.Filename, handler.Size, handler.Header)

	tempFile, err := os.Create(path.Join(l.svcCtx.Config.Path, handler.Filename))
	if err != nil {
		logx.Error(err)
		return
	}
	defer tempFile.Close()
	io.Copy(tempFile, file)

	// return &types.UploadResponse{
	// 	Code: 0,
	// }, nil
}
