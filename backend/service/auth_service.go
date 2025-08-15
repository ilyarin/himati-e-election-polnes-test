package service

import (
	"context"

	"github.com/proyek2-git/HIMA-TI-e-Election/backend/model/web"
)

type AuthService interface {
	LoginUser(ctx context.Context, maxAge int, request web.LoginRequest) (web.LoginResponse, string, error)
	LoginAdmin(ctx context.Context, maxAge int, request web.LoginRequest) (web.LoginResponse, string, error)
	Logout(ctx context.Context, sessionId string) error
	UserValidateSession(ctx context.Context, sessionId string) (web.SessionResponse, error)
	AdminValidateSession(ctx context.Context, sessionId string) (web.SessionResponse, error)
}
