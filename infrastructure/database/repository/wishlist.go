package repository

import (
	"github.com/felipedss/trip-wishlist-api/domain/entity"
	"github.com/felipedss/trip-wishlist-api/infrastructure/db"
)

type WishlistRepositorySt struct {
}

func NewWishlistRepository() *WishlistRepositorySt {
	return &WishlistRepositorySt{}
}

func (a *WishlistRepositorySt) Insert(wishlist entity.Wishlist) error {
	db := db.GetDB()

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, errInsert := tx.Exec("INSERT INTO WISHLIST (DESTINATION) VALUES(?)", wishlist.Destination)
	if errInsert != nil {
		if err := tx.Rollback(); err != nil {
			return err
		}
		return errInsert
	}

	if err := tx.Commit(); err != nil {
		return err
	}
	return nil

}
