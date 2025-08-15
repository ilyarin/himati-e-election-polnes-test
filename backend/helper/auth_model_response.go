package helper

import (
	"github.com/proyek2-git/HIMA-TI-e-Election/backend/model/domain"
	"github.com/proyek2-git/HIMA-TI-e-Election/backend/model/web"
)

func ToLoginResponse(user domain.User) web.LoginResponse {
	return web.LoginResponse{
		ID:           user.Id,
		NIM:          user.NIM,
		FullName:     user.FullName,
		StudyProgram: user.StudyProgram,
		Role:         user.Role,
		PhoneNumber:  user.PhoneNumber,
		CreatedAt:    user.CreatedAt,
		UpdatedAt:    user.UpdatedAt,
	}
}
