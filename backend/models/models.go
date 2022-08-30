package model

import "time"

type Categories struct {
	CategoryId        string `json:"_id,omitempty" bson:"_id,omitempty"`
	CategoryName      string `json:"category_name,omitempty" bson:"category_name,omitempty"`
	CategoryPermalink string `json:"cat_permalink,omitempty" bson:"cat_permalink,omitempty"`
}

type Product struct {
	ProductId          string  `json:"_id,omitempty" bson:"_id,omitempty"`
	CategoryId         string  `json:"cat_id,omitempty" bson:"cat_id,omitempty"`
	ProductName        string  `json:"product_name,omitempty" bson:"product_name,omitempty"`
	ProductPrice       float64 `json:"product_price,omitempty" bson:"product_price,omitempty"`
	ProductUnit        int     `json:"product_unit,omitempty" bson:"product_unit,omitempty"`
	ProductBrand       string  `json:"product_brand,omitempty" bson:"product_brand,omitempty"`
	ProductDescription string  `json:"product_desc,omitempty" bson:"product_desc,omitempty"`
}

type User struct {
	UserId     string `json:"_id,omitempty" bson:"_id,omitempty"`
	FullName   string `json:"full_name,omitempty" bson:"full_name,omitempty"`
	Email      string `json:"email,omitempty"`
	Password   string `json:"password,-"`
	IsActive   bool   `json:"is_active" bson:"is_active"`
	IsAdmin    bool   `json:"is_admin" bson:"is_admin"`
	ProfilePic string `json:"profile_pic,omitempty" bson:"profile_pic,omitempty"`
}

type ShippingAddress struct {
	AddressId string `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId    string `json:"user_id,omitempty"`
	Address   string `json:"address,omitempty"`
	City      string `json:"city,omitempty"`
	State     string `json:"state,omitempty"`
	Pincode   string `json:"pincode,omitempty"`
	Country   string `json:"country,omitempty" default:"India"`
}

type UserToken struct {
	TokenId   string    `json:"_id,omitempty" bson:"_id,omitempty"`
	Token     string    `json:"token,omitempty" bson:"token,omitempty"`
	Email     string    `json:"email,omitempty" bson:"email,omitempty"`
	Count     int       `json:"count,omitempty" bson:"count,omitempty"`
	CreatedOn time.Time `json:"created_on,omitempty" bson:"created_on,omitempty"`
	UpdateOn  time.Time `json:"updated_on,omitempty" bson:"updated_on,omitempty"`
	Purpose   string    `json:"purpose" bson:"purpose"`
}

type UserOrder struct {
	OrderId    string    `json:"_id,omitempty" bson:"_id,omitempty"`
	UserId     string    `json:"user_id,omitempty"`
	ProductId  string    `json:"product_id,omitempty"`
	Address    string    `json:"address,omitempty"`
	ProductQty int       `json:"product_qty,omitempty"`
	CreatedOn  time.Time `json:"created_on,omitempty"`
}
