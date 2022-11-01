package dto

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID             primitive.ObjectID `json:"_id" bson:"_id"`
	FistName       *string            `json:"first_name" validate:"required,min=2,max=30"`
	LastName       *string            `json:"last_name"  validate:"required,min=2,max=30"`
	Password       *string            `json:"password"   validate:"required,min=6"`
	Email          *string            `json:"email"      validate:"email,required"`
	Phone          *string            `json:"phone"      validate:"required"`
	Token          *string            `json:"token"`
	RefreshToken   *string            `josn:"refresh_token"`
	CreatedAt      time.Time          `json:"created_at"`
	UpdatedAt      time.Time          `json:"updated_at"`
	UserID         string             `json:"user_id"`
	UserCart       []ProductUser      `json:"usercart" bson:"usercart"`
	AddressDetails []Address          `json:"address" bson:"address"`
	OrderStatus    []Order            `json:"orders" bson:"orders"`
}

type Product struct {
	ProductID   primitive.ObjectID `bson:"_id"`
	ProductName *string            `json:"product_name"`
	Price       *uint64            `json:"price"`
	Ratings     *uint8             `json:"ratings"`
	Image       *string            `json:"image"`
}

type ProductUser struct {
	ProductID
	ProductName
	Price
	Ratings
	Image
}
type Address struct {
	AddressID
	House
	Street
	City
	Pincode
}
type Order struct {
	OrderID
	OrderCart
	OrderedAt
	Price
	Discount
	PaymentMethod
}

type Payment struct {
	Digital bool `json:"digital" bson:"digital"`
	COD     bool `json:"cod"     bson:"cod"`
}
