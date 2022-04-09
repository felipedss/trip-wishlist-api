package service

import (
	"github.com/felipedss/trip-wishlist-api/domain/dataprovider"
	"github.com/felipedss/trip-wishlist-api/domain/entity"
)

type WishlistService interface {
	Insert(wishlist entity.Wishlist) error
}

type WishlistServiceSt struct {
	repository dataprovider.WishlistRepository
}

func NewWishlistService(wishlistRepository dataprovider.WishlistRepository) *WishlistServiceSt {
	return &WishlistServiceSt{
		repository: wishlistRepository,
	}
}

func (serviceSt *WishlistServiceSt) Insert(wishlist entity.Wishlist) error {
	return serviceSt.repository.Insert(wishlist)
}
