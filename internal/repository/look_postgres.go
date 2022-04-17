package repository

import (
	"fmt"

	"github.com/gogaeva/shmot-shprot/internal/domain"
	"github.com/jmoiron/sqlx"
)

type LookPostgres struct {
	db *sqlx.DB
}

func NewLookPostgres(db *sqlx.DB) *LookPostgres {
	return &LookPostgres{db: db}
}

func (r *LookPostgres) CreateLook(look domain.Look) (uint, error) {
	var id uint

	query := fmt.Sprintf("INSERT INTO %s (photo_id, owner_id, description, season, temperature_range, purpose, priority) RETURNING id", looksTable)
	row := r.db.QueryRow(query, look.PhotoPath, look.OwnerId, look.Description, look.Season, look.TemperatureRange, look.Purpose, look.Priority)
	if err := row.Scan(&id); err != nil {
		return 0, err
	}
	return id, nil
}
