package v1

import (
	"github.com/qvvan/dragonfit/internal/app/repository"
)

type Manager struct {
	factory *repository.Factory
}

func NewManager(
	factory *repository.Factory,
) *Manager {
	return &Manager{
		factory: factory,
	}
}
