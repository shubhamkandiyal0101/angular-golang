package router

import (
	"github.com/gorilla/mux"
	"github.com/shubhamkandiyal0101/go-basic-ecommerce-app/controllers/admincontroller"
	"github.com/shubhamkandiyal0101/go-basic-ecommerce-app/controllers/homecontroller"
	"github.com/shubhamkandiyal0101/go-basic-ecommerce-app/controllers/usercontroller"
)

func Router() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/user-signup", usercontroller.UserSignup).Methods("POST")
	r.HandleFunc("/user-login", usercontroller.UserLogin).Methods("POST")
	r.HandleFunc("/verify-token", usercontroller.VerifyUserToken).Methods("POST")
	r.HandleFunc("/add-product-category", admincontroller.AddProductCategory).Methods("POST")
	r.HandleFunc("/add-product/{cat_id}", admincontroller.AddNewProduct).Methods("POST")
	r.HandleFunc("/delete-product-category/{cat_id}", admincontroller.DeleteProductCategory).Methods("DELETE")
	r.HandleFunc("/products", homecontroller.Products).Methods("GET")
	r.HandleFunc("/update-product-details/{product_id}", admincontroller.UpdateProductDetails).Methods("PUT")
	r.HandleFunc("/product/{product_id}", homecontroller.GetSingleProduct).Methods("GET")
	r.HandleFunc("/send-test-email", homecontroller.SendTestEmail).Methods("POST")
	r.HandleFunc("/forgot-password", usercontroller.ForgotPassword).Methods("POST")
	r.HandleFunc("/reset-password/{token}", usercontroller.ResetPassword).Methods("POST")
	r.HandleFunc("/place-order/{product_id}", usercontroller.PlaceOrder).Methods("POST")
	r.HandleFunc("/get-all-categories", homecontroller.GetAllCategories).Methods("GET")

	return r
}

// TODO:
// 1. Buy Product
// 2. Update User Details
// 3. Update User Image
