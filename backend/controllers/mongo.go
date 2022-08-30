package mongocontroller

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/joho/godotenv"
	jwttokenhelper "github.com/shubhamkandiyal0101/go-basic-ecommerce-app/helpers"
	model "github.com/shubhamkandiyal0101/go-basic-ecommerce-app/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection
var categoryCollection *mongo.Collection
var productCollection *mongo.Collection
var userOrdersCollection *mongo.Collection
var userTokenCollection *mongo.Collection

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	connectionString := os.Getenv("MONGO_URI")

	clientOptions := options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(" MONGO CONNECTION ERROR: ", err)
	}

	userCollection = client.Database("my-shopping-app").Collection("users")
	productCollection = client.Database("my-shopping-app").Collection("products")
	categoryCollection = client.Database("my-shopping-app").Collection("categories")
	// shippingAddressCollection = client.Database("my-shopping-app").Collection("shipping_address")
	userTokenCollection = client.Database("my-shopping-app").Collection("user_token")
	userOrdersCollection = client.Database("my-shopping-app").Collection("order_info")

	fmt.Println(" MONGO DB INSTANCE IS READY FOR USE ")

}

// GENERATE AND VERIFY ENCRYPTED PASSWORD
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

// ENDS HERE ~ GENERATE AND VERIFY ENCRYPTED PASSWORD

// fetch basic user details
func FetchBasicDetails(email string) map[string]interface{} {
	var filterData map[string]interface{}
	userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&filterData)

	return filterData
}

// ends here ~ fetch basic user details

// JWT TOKEN
// ENDS HERE ~ JWT TOKEN

// USER SIGNUP
func UserSignupDb(user *model.User) (interface{}, int) {
	// fmt.Println(" -- user signup -- ", user)
	userEmail := user.Email
	userPassword := user.Password

	var filterData bson.M
	userCollection.FindOne(context.Background(), bson.M{"email": userEmail}).Decode(&filterData)

	if len(filterData) == 0 {
		hashPswd, _ := HashPassword(userPassword)

		user.IsActive = true
		user.Password = hashPswd

		userCollection.InsertOne(context.Background(), user)

		response := map[string]string{
			"Message": "User Created Successfully",
		}

		return response, 200

	} else {
		response := map[string]string{
			"Message": "User already Exists in DB. Please Login",
		}
		return response, 400
	}

}

// ENDS HERE ~ USER SIGNUP

// USER LOGIN
func UserLoginDb(user *model.User) (interface{}, int) {
	// fmt.Println(" -- user signup -- ", user)
	userEmail := user.Email
	userPassword := user.Password
	isAdmin := user.IsAdmin

	// var filterData bson.M
	var filterData map[string]interface{}
	userCollection.FindOne(context.Background(), bson.M{"email": userEmail, "is_admin": isAdmin}).Decode(&filterData)

	if len(filterData) == 0 {
		response := map[string]interface{}{
			"Message": "User is not exists with this email exists",
			"data": map[string]string{
				"email": userEmail,
			},
		}
		return response, 400
	}

	var userHashPswdStr string = ""
	userHashPswd := filterData["password"]
	userHashPswdStr = fmt.Sprint(userHashPswd)

	isPswdCorrect := CheckPasswordHash(userPassword, userHashPswdStr)
	if !isPswdCorrect {
		response := map[string]interface{}{
			"Message": "Please provide correct Password for this E-Mail",
			"data": map[string]string{
				"email": userEmail,
			},
		}
		return response, 400
	}

	newToken, _ := jwttokenhelper.GenerateJWT(userEmail)

	response := map[string]interface{}{
		"Message": "Login Successfully",
		"data": map[string]string{
			"token": newToken,
		},
	}

	return response, 200

}

// ENDS HERE ~ USER LOGIN

// Add Products Category in DB
func AddProductCatInDB(productCat *model.Categories) interface{} {
	catPermalink := productCat.CategoryPermalink

	var filterData map[string]interface{}
	categoryCollection.FindOne(context.Background(), bson.M{"cat_permalink": catPermalink}).Decode(&filterData)

	if len(filterData) > 0 {
		response := map[string]interface{}{
			"Message": "Category Already Exists",
			"data": map[string]string{
				"cat_permalink": catPermalink,
			},
		}
		return response
	}

	categoryCollection.InsertOne(context.Background(), productCat)

	response := map[string]string{
		"message": "Product Category Created Successfully",
	}

	return response
}

// Add Products Category in DB

// check is category exists or not
func IsCatIdExists(cat_id string) bool {
	var filterData map[string]interface{}

	objId, err := primitive.ObjectIDFromHex(cat_id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objId}

	err = categoryCollection.FindOne(context.Background(), filter).Decode(&filterData)

	if err != nil {
		return false
	} else {
		return true
	}

}

// ends here ~ check is category exists or not

// Add New Product DB
func AddNewProductDb(product *model.Product) interface{} {

	productCollection.InsertOne(context.Background(), product)

	response := map[string]string{
		"message": "Product Added Successfully",
	}

	return response
}

// ends here ~ Add New Product DB

// delete category
func DeleteCategoryDb(cat_id string) interface{} {
	is_cat_exist := IsCatIdExists(cat_id)

	if is_cat_exist == false {
		response := map[string]interface{}{
			"message": "Please provide Valid Category ID",
			"data": map[string]string{
				"cat_id": cat_id,
			},
		}
		return response
	}

	catObjectId, err := primitive.ObjectIDFromHex(cat_id)

	// fmt.Println(" catObjectId >> ", catObjectId, " >> err >> ", err, " >> cat_id >> ", cat_id)

	if err != nil {
		response := map[string]interface{}{
			"message": "Please provide Valid Category ID.",
			"data": map[string]string{
				"cat_id": cat_id,
			},
		}
		return response
	}

	var filterData map[string]interface{}
	err = productCollection.FindOne(context.Background(), bson.M{"cat_id": cat_id}).Decode(&filterData)

	if len(filterData) > 0 {
		response := map[string]interface{}{
			"message": "We have many products which are connected with this Category ID. So, its not possible to delete this Category now",
			"data": map[string]string{
				"cat_id": cat_id,
			},
		}
		return response
	}

	categoryCollection.DeleteOne(context.Background(), bson.M{"_id": catObjectId})

	response := map[string]interface{}{
		"message": "Category Deleted Successfully",
		"data": map[string]string{
			"cat_id": cat_id,
		},
	}
	return response

}

// ends here ~ delete category

// Get all Products
func GetAllProductsDb() interface{} {
	var products []model.Product
	cursor, err := productCollection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {

		var product model.Product
		err := cursor.Decode(&product)

		if err != nil {
			log.Fatal(err)
		}

		products = append(products, product)

	}

	return products
}

// ends here ~ Get All Products

// check product id exists or not
func IsProdIdExists(product_id string) bool {
	var filterData map[string]interface{}

	objId, err := primitive.ObjectIDFromHex(product_id)
	if err != nil {
		log.Fatal(err)
	}
	filter := bson.M{"_id": objId}

	err = productCollection.FindOne(context.Background(), filter).Decode(&filterData)

	if err != nil {
		return false
	} else {
		return true
	}
}

// ends here ~ check product id exists or not

// update product details
func UpdateProductDetailsDb(product_id string, productData map[string]interface{}) interface{} {

	objId, _ := primitive.ObjectIDFromHex(product_id)

	productCollection.FindOneAndUpdate(context.Background(), bson.M{"_id": objId}, bson.M{"$set": productData})

	response := map[string]string{
		"message": "Product Updated Successfully",
	}

	return response

}

// ends here ~ update product details

// get single product details
func GetSingleProductDetailsDb(product_id string) map[string]interface{} {
	var product *model.Product

	objId, _ := primitive.ObjectIDFromHex(product_id)
	err := productCollection.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&product)

	if err != nil {
		response := map[string]interface{}{
			"message": "This Product ID is not exists in Our Records. Please provide Valid Product ID",
			"data": map[string]interface{}{
				"product_id": product_id,
			},
		}

		return response
	}

	response := map[string]interface{}{
		"message": "Product Data Fetch Succesfully",
		"data": map[string]interface{}{
			"product_details": product,
		},
	}

	return response

}

// ends here ~ get single product details

// check is email id exists or not
func IsEmailIdExists(email string) bool {

	var userDetails map[string]interface{}

	err := userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&userDetails)
	if err != nil {
		return false
	} else {
		return true
	}

}

// ends here ~ check is email id exists or not

// send email for forgot password
func ForgotPasswordDb(email string) (bool, interface{}, string) {
	purpose := "reset_password"

	var userTokenDetails map[string]interface{}

	filter := bson.M{"email": email, "purpose": purpose}
	err := userTokenCollection.FindOne(context.Background(), filter).Decode(&userTokenDetails)

	if err != nil {
		var userToken model.UserToken
		uniqueToken := uuid.NewString()
		userToken.Token = uniqueToken
		userToken.Email = email
		userToken.Count = 1
		userToken.CreatedOn = time.Now()
		userToken.UpdateOn = time.Now()
		userToken.Purpose = purpose

		fmt.Println(" Error >> ", err)

		userTokenCollection.InsertOne(context.Background(), &userToken)

		response := map[string]interface{}{
			"message": "Forgot Password E-Mail Sent Successfully. Please check your email for Reset Token",
			"data": map[string]interface{}{
				"email": email,
			},
		}

		return true, response, uniqueToken

	} else {
		response := map[string]interface{}{
			"message": "Please check your email. Token is already sent to your email to Reset Password",
			"data": map[string]interface{}{
				"email": email,
			},
		}
		return false, response, ""
	}

	// data := model.UserToken{nil, email, 1, time.Now(), time.Now(), purpose}

	// userTokenCollection

	return false, map[string]interface{}{}, ""

}

// ends here ~ send email for forgot password

// check token is valid or not
func CheckTokenIsValid(token string) (bool, interface{}, interface{}) {

	var userTokenDetails map[string]interface{}

	userTokenCollection.FindOne(context.Background(), bson.M{"token": token}).Decode(&userTokenDetails)

	fmt.Println(" >> ", userTokenDetails)
	fmt.Println(" >> ", len(userTokenDetails))
	// fmt.Println(err, " >> ",len(userTokenDetails))

	if len(userTokenDetails) >= 1 {

		response := map[string]interface{}{
			"data": map[string]interface{}{
				"token": token,
			},
			"message": "Token is Valid",
		}

		return true, response, userTokenDetails["email"]

	} else {
		response := map[string]interface{}{
			"data": map[string]interface{}{
				"token": token,
			},
			"message": "Token is Invalid",
		}

		return false, response, ""
	}

}

// ends here ~ check token is valid or not

// set new password
func SetNewPasswordDb(password string, email string) {
	hashPswd, _ := HashPassword(password)

	userData := map[string]string{"password": hashPswd}

	userCollection.FindOneAndUpdate(context.Background(), bson.M{"email": email}, bson.M{"$set": userData})

}

// ends here ~ set new password

func PlaceOrderDb(productId string, productQty interface{}, email string) interface{} {
	objId, _ := primitive.ObjectIDFromHex(productId)

	var response map[string]interface{}

	var productData map[string]interface{}
	err := productCollection.FindOne(context.Background(), bson.M{"_id": objId}).Decode(&productData)

	if err != nil {
		response = map[string]interface{}{
			"data": map[string]interface{}{
				"product_id": productId,
			},
			"message": "Product is not Available",
		}
	}

	var userDetails map[string]interface{}
	userCollection.FindOne(context.Background(), bson.M{"email": email}).Decode(&userDetails)

	var orderDetails model.UserOrder

	// UserId
	orderDetails.ProductId = productId
	orderDetails.Address = ""
	// orderDetails.ProductQty = productQty

	return response
}

// get all products
func GetAllCategories() (interface{}, int) {
	var categoies []model.Categories
	cursor, err := categoryCollection.Find(context.Background(), bson.M{})

	if err != nil {
		log.Fatal(err)
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {

		var category model.Categories
		err := cursor.Decode(&category)

		if err != nil {
			log.Fatal(err)
		}

		categoies = append(categoies, category)

	}

	return categoies, 200
}

// ends here ~ get all products
