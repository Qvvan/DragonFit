package v1

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/qvvan/dragonfit/internal/app/models"
	errorDb "github.com/qvvan/dragonfit/pkg/client/postgresql/utils"
)

type RegisterRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

func (m *Manager) Register(c *gin.Context) {
	token, err := m.RegisterService(c)

	if err != nil {
		var pubErr *PublicError
		switch {
		case errors.As(err, &pubErr):
			c.JSON(pubErr.status, gin.H{"error": pubErr.Error()})
			return
		default:
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
	}

	c.JSON(http.StatusOK, token)
}

func (m *Manager) RegisterService(c *gin.Context) (*models.User, error) {
	var req RegisterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		return nil, NewPublicErr(err, http.StatusBadRequest)
	}
	newUser := &models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	createdUser, err := m.factory.UserRepo.Create(c, newUser)
	if err != nil {
		if err.Code == errorDb.PGErrDuplicateCode {
			return nil, NewPublicErr(err.Message, http.StatusConflict)
		}
		return nil, NewPublicErr(err.Message, http.StatusInternalServerError)
	}

	return createdUser, nil
}
