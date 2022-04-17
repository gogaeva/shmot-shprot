package repository

import (
	"fmt"

	"github.com/gogaeva/shmot-shprot/internal/domain"
	"github.com/jmoiron/sqlx"
)

type ClothRepo struct {
	db *sqlx.DB
}

func NewClothRepo(db *sqlx.DB) *ClothRepo {
	return &ClothRepo{db: db}
}

func (r *ClothRepo) AddCloth(cloth domain.Cloth) (uint, error) {
	//var id int64
	//? need RETURNING id or not?
	query := fmt.Sprintf("INSERT INTO %s (photo_id, owner_id, class, brand, color) VALUES ($1, $2, $3, $4, $5) RETURNING id", clothesTable)
	// row := r.db.QueryRow(query, cloth.PhotoId, cloth.OwnerId, cloth.Class, cloth.Brand, cloth.Color)
	// if err := row.Scan(&id); err != nil {
	// 	return 0, err
	// }
	stmt, err := r.db.Preparex(query)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	// res, err := stmt.Exec(cloth.PhotoId, cloth.OwnerId, cloth.Class, cloth.Brand, cloth.Color)
	// if err != nil {
	// 	return 0, err
	// }

	// id, err := res.LastInsertId()
	// if err != nil {
	// 	return 0, err
	// }

	var id uint
	err = stmt.Get(&id, cloth.PhotoId, cloth.OwnerId, cloth.Class, cloth.Brand, cloth.Color)

	return id, err
}

func (r *ClothRepo) DeleteCloth(id uint) error {
	stmt, err := r.db.Prepare(fmt.Sprintf("DELETE FROM %s WHERE id = $1", clothesTable))
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)

	return err
}

func (r *ClothRepo) GetCloth(id uint) (domain.Cloth, error) {
	var cloth domain.Cloth
	query := fmt.Sprintf("SELECT * FROM %s WHERE id=$1", clothesTable)
	err := r.db.Get(&cloth, query, id)

	return cloth, err
}

func (r *ClothRepo) GetAllClothes(userId uint) ([]domain.Cloth, error) {
	var clothes []domain.Cloth
	query := fmt.Sprintf("SELECT * FROM %s WHERE owner_id=$1", clothesTable)
	err := r.db.Select(clothes, query, userId)
	if err != nil {
		return nil, err
	}

	return clothes, nil
}

func (r *ClothRepo) UpdateCloth(id uint, new *domain.Cloth) error {
	stmt, err := r.db.Prepare("UPDATE %s SET photo_id=$1, owner_id=$2, class=$3, brand=$4, color=$5 WHERE id=$6")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(new.PhotoId, new.OwnerId, new.Class, new.Brand, new.Color, new.Id)

	return err
}
