package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/proyek2-git/HIMA-TI-e-Election/backend/model/domain"
)

type AuthRepository interface {
	Create(ctx context.Context, tx pgx.Tx, session domain.Session) (domain.Session, error)
	GetSessionById(ctx context.Context, tx pgx.Tx, sessionId string) (domain.Session, error)
	Delete(ctx context.Context, tx pgx.Tx, sessionId string) error
}
