package service

import (
	"context"

	"github.com/proyek2-git/HIMA-TI-e-Election/backend/model/web"
)

type DownloadService interface {
	CreatePresignedURL(ctx context.Context, fileName string) (web.PresignedURLResponse, error)
}
