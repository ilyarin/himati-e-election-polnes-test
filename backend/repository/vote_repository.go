package repository

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/proyek2-git/HIMA-TI-e-Election/backend/model/domain"
)

type VoteRepository interface {
	GetByCandidateId(ctx context.Context, tx pgx.Tx, candidateId int) ([]domain.Vote, error)
	IsUserVotedInPeriod(ctx context.Context, tx pgx.Tx, hashedNIM string) (bool, error)
	SaveVoteRecord(ctx context.Context, tx pgx.Tx, vote domain.Vote) (domain.Vote, error)
	GetTotalVotesByCandidateId(ctx context.Context, tx pgx.Tx, candidateId int) (int, error)
	IsUserEverVoted(ctx context.Context, tx pgx.Tx, userId int) (bool, error)
}
