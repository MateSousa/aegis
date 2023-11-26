package health

import (
	"context"

	"gorm.io/gorm"
)

type healthDBRepository struct {
	DB *gorm.DB
}

func NewHealthRepository(db *gorm.DB) *healthDBRepository {
	return &healthDBRepository{DB: db}
}

func (h healthDBRepository) GetHealthDB(ctx context.Context) error {
	tx := h.DB.WithContext(ctx).Exec("SELECT 1")

	return tx.Error
}
