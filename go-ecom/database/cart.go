package database

import "errors"

var (
	ErrCantFindProduct    = errors.New("can't find product")
	ErrCantDecodeProducts = errors.New("can't find product")
	ErrUserIdIsNotValid   = errors.New("user is not valid")
	ErrCantUpdateUser     = errors.New("can't add product to cart")
	ErrCantRemoveItem     = errors.New("can't remove item from cart")
	ErrCantGetItem        = errors.New("can't get item from cart")
	ErrCantBuyCartItem    = errors.New("can't update the purchase")
)

func AddProductToCart() {}

func RemoveCartItem() {}

func BuyItemFromCart() {}

func InstantBuyer() {}
