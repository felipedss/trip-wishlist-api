package dto

import "github.com/felipedss/trip-wishlist-api/domain/entity"

type CreateWishlistDTO struct {
	Destination string `json:"destination"`
}

type CreateWishlistResponseDTO struct {
	ID uint `json:"id"`
}

func (p *CreateWishlistDTO) ToEntity() entity.Wishlist {
	var wishlist entity.Wishlist
	wishlist.Destination = p.Destination
	return wishlist
}

func ToResponse(wishlist entity.Wishlist) CreateWishlistResponseDTO {
	return CreateWishlistResponseDTO{ID: wishlist.ID}
}
