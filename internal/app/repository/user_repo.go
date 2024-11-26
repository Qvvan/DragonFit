package repository

import (
	"log/slog"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/qvvan/dragonfit/internal/app/models"
	errors "github.com/qvvan/dragonfit/pkg/client/postgresql/utils"
)

const (
	UsersTName = "users"
)

type UserRepo struct {
	BaseRepo
	log *slog.Logger
}

func NewUserRepo(db *pgxpool.Pool, log *slog.Logger) *UserRepo {
	repo := &UserRepo{}
	repo.db = db
	repo.table = UsersTName
	repo.log = log
	return repo
}

func (r *UserRepo) Create(ctx *gin.Context, user *models.User) (*models.User, *errors.CustomError) {
	id, err := r.BaseRepo.create(ctx, user)
	if err != nil {
		r.log.Error("failed to create user", slog.Any("err", err))
		return nil, err
	}
	user.ID = id
	return user, nil
}

func (r *UserRepo) Update(user *models.User) (string, error) {
	return "User updated successfully", nil
}

func (r *UserRepo) Delete(user *models.User) (string, error) {
	return "User deleted successfully", nil
}

func (r *UserRepo) Get(email string) (*models.User, error) {
	return nil, nil
}

func (r *UserRepo) GetByID(userID int) (*models.User, error) {
	return nil, nil
}
