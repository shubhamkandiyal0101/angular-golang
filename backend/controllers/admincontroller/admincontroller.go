package admincontroller

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	mongocontroller "github.com/shubhamkandiyal0101/go-basic-ecommerce-app/controllers"
	jwttokenhelper "github.com/shubhamkandiyal0101/go-basic-ecommerce-app/helpers"
	model "github.com/shubhamkandiyal0101/go-basic-ecommerce-app/models"
)

// add product categpry
func AddProductCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- Add Product Categories -- ")

	w.Header().Set("Content-Type", "application/json")

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	email, err := jwttokenhelper.VerifyToken(reqToken)

	if err != nil {
		log.Fatal(err)
	}

	userBasicDetails := mongocontroller.FetchBasicDetails(email)

	if userBasicDetails["is_admin"] == false {
		response := map[string]string{
			"message": "User don't have enough Permission",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	var productCategory *model.Categories

	json.NewDecoder(r.Body).Decode(&productCategory)

	response := mongocontroller.AddProductCatInDB(productCategory)

	json.NewEncoder(w).Encode(response)

	defer r.Body.Close()

}

// ends here ~ add product categpry

// add new product
func AddNewProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- -- Add New Product -- -- ")

	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	email, err := jwttokenhelper.VerifyToken(reqToken)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(email)

	params := mux.Vars(r)
	cat_id := params["cat_id"]

	isCatIdExists := mongocontroller.IsCatIdExists(cat_id)

	if !isCatIdExists {
		response := map[string]interface{}{
			"message": "Please provide valid Category ID",
			"data": map[string]interface{}{
				"category_id": cat_id,
			},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	var product *model.Product

	json.NewDecoder(r.Body).Decode(&product)
	product.CategoryId = cat_id

	response := mongocontroller.AddNewProductDb(product)
	json.NewEncoder(w).Encode(response)
	return

}

// ends here ~ add new product

// delete product category
func DeleteProductCategory(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- Delete Product Categories -- ")

	w.Header().Set("Content-Type", "application/json")

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	email, err := jwttokenhelper.VerifyToken(reqToken)

	if err != nil {
		log.Fatal(err)
	}

	userBasicDetails := mongocontroller.FetchBasicDetails(email)

	if userBasicDetails["is_admin"] == false {
		response := map[string]string{
			"message": "User don't have enough Permission",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	params := mux.Vars(r)
	cat_id := params["cat_id"]

	response := mongocontroller.DeleteCategoryDb(cat_id)

	json.NewEncoder(w).Encode(response)

	defer r.Body.Close()

}

// ends here ~ delete product category

// Update Product Details
func UpdateProductDetails(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- Delete Product Categories -- ")

	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	email, err := jwttokenhelper.VerifyToken(reqToken)

	if err != nil {
		log.Fatal(err)
	}

	userBasicDetails := mongocontroller.FetchBasicDetails(email)

	if userBasicDetails["is_admin"] == false {
		response := map[string]string{
			"message": "User don't have enough Permission",
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	var reqData map[string]interface{}

	json.NewDecoder(r.Body).Decode(&reqData)

	if len(reqData) == 0 {
		response := map[string]string{
			"message": "Plesae Provide Data to Update Details",
		}

		json.NewEncoder(w).Encode(response)
		return
	}

	params := mux.Vars(r)
	product_id := params["product_id"]

	isProdIdExists := mongocontroller.IsProdIdExists(product_id)

	if !isProdIdExists {
		response := map[string]interface{}{
			"message": "Please provide valid Product ID",
			"data": map[string]interface{}{
				"product_id": product_id,
			},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := mongocontroller.UpdateProductDetailsDb(product_id, reqData)
	json.NewEncoder(w).Encode(response)
	return
}

// ends here ~ Update Product Details
