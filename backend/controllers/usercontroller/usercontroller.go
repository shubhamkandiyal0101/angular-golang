package usercontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
	"os"
	"strings"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	mongocontroller "github.com/shubhamkandiyal0101/go-basic-ecommerce-app/controllers"
	jwttokenhelper "github.com/shubhamkandiyal0101/go-basic-ecommerce-app/helpers"
	model "github.com/shubhamkandiyal0101/go-basic-ecommerce-app/models"
)

func smtpAuth() smtp.Auth {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	EMAIL := os.Getenv("EMAIL")
	EMAIL_PASSWORD := os.Getenv("EMAIL_PASSWORD")

	hostname := "smtp.gmail.com"
	// auth := smtp.PlainAuth("", "naco.shubhamkandiyal@gmail.com", "ahxpbvebkuzsxgoh", hostname)
	auth := smtp.PlainAuth("", EMAIL, EMAIL_PASSWORD, hostname)

	return auth
}

func UserSignup(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- -- USER SIGNUP -- -- ")
	w.Header().Set("Content-Type", "application/json")

	var userDetails *model.User

	json.NewDecoder(r.Body).Decode(&userDetails)

	response, responseCode := mongocontroller.UserSignupDb(userDetails)
	w.WriteHeader(responseCode)

	json.NewEncoder(w).Encode(response)
	defer r.Body.Close()

}

func UserLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- -- USER LOGIN -- -- ")
	w.Header().Set("Content-Type", "application/json")

	var userDetails *model.User
	json.NewDecoder(r.Body).Decode(&userDetails)

	response, responseCode := mongocontroller.UserLoginDb(userDetails)
	w.WriteHeader(responseCode)

	json.NewEncoder(w).Encode(response)
	defer r.Body.Close()

}

func VerifyUserToken(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- -- VERIFY USER TOKEN -- -- ")
	w.Header().Set("Content-Type", "application/json")

	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	email, err := jwttokenhelper.VerifyToken(reqToken)

	if err != nil {
		log.Fatal(err)
	}

	response := map[string]string{
		"email": email,
	}

	json.NewEncoder(w).Encode(response)
	defer r.Body.Close()
}

func ForgotPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- -- Forgot Password -- -- ")
	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()

	var reqData map[string]interface{}

	json.NewDecoder(r.Body).Decode(&reqData)

	var email string
	emailVal := reqData["email"]
	email = fmt.Sprintf("%v", emailVal)

	isEmailExists := mongocontroller.IsEmailIdExists(email)

	if !isEmailExists {
		response := map[string]interface{}{
			"message": "Please Provide Valid E-Mail ID",
		}
		json.NewEncoder(w).Encode(response)

		return
	}

	is_new_data_inserted, response, token := mongocontroller.ForgotPasswordDb(email)

	if is_new_data_inserted {

		// set email template variable data
		type ForgotPswdEmailVars struct {
			Token string
		}
		templateData := &ForgotPswdEmailVars{token}
		// ends here ~ set email template variable data

		// parse template and store into buffer
		parseTemplate := template.Must(template.ParseFiles("emailtemplates/forgot_password.html"))
		var tpl bytes.Buffer
		parseTemplate.Execute(&tpl, templateData)

		templateString := tpl.String()
		// ends here ~ parse template and store into buffer

		auth := smtpAuth()

		to := []string{email}
		subject := "Subject: Forgot Password - Go Shopping App\n"
		mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
		body := templateString

		from := mail.Address{"Go Shopping App", "naco.shubhamkandiyal@gmail.com"}
		fromString := "From: " + from.String() + "\r\n"

		msg := []byte(fromString + subject + mime + body)

		err := smtp.SendMail("smtp.gmail.com:587", auth, from.Address, to, msg)
		if err != nil {
			log.Fatal(err)
		}

	}

	json.NewEncoder(w).Encode(response)

	return

}

func ResetPassword(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- RESET PASSWORD -- ")

	w.Header().Set("Content-Type", "application/json")

	defer r.Body.Close()

	var reqData map[string]interface{}

	json.NewDecoder(r.Body).Decode(&reqData)

	token := mux.Vars(r)["token"]

	is_token_valid, response, email := mongocontroller.CheckTokenIsValid(token)

	// fmt.Println(reqData["password"]," -- response -- ",email)
	// fmt.Println(fmt.Sprint(reqData["password"])," -- response -- ",email)

	if is_token_valid {

		passwordStr := fmt.Sprint(reqData["password"])
		emailStr := fmt.Sprint(email)

		mongocontroller.SetNewPasswordDb(passwordStr, emailStr)

		// fmt.Println(passwordStr," -- response -- ",emailStr)

		response := map[string]interface{}{
			"data":    map[string]interface{}{},
			"message": "Password Changed Successfully",
		}

		json.NewEncoder(w).Encode(response)
	} else {
		json.NewEncoder(w).Encode(response)
	}

}

func PlaceOrder(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- -- -- PLACE ORDER -- -- -- ")

	w.Header().Set("Content-Type", "application/json")

	// check user is logged in or not
	reqToken := r.Header.Get("Authorization")
	splitToken := strings.Split(reqToken, "Bearer ")
	reqToken = splitToken[1]

	email, err := jwttokenhelper.VerifyToken(reqToken)

	if err != nil {
		log.Fatal(err)

		response := map[string]interface{}{
			"data":    "",
			"message": "Unauthorised User",
		}

		json.NewEncoder(w).Encode(response)
	}
	// ends here ~ check user is logged in or not

	defer r.Body.Close()

	var reqData map[string]interface{}

	json.NewDecoder(r.Body).Decode(&reqData)

	productId := mux.Vars(r)["product_id"]
	var prodQty = reqData["product_qty"]

	fmt.Println(productId, " -- -- ", prodQty, "------", email)

	mongocontroller.PlaceOrderDb(productId, prodQty, email)

}
