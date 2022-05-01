package dataprovider

import "github.com/felipedss/trip-wishlist-api/domain/entity"

type WishlistRepository interface {
	Insert(wishlist entity.Wishlist) error
}
