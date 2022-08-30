package homecontroller

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"

	"github.com/gorilla/mux"
	mongocontroller "github.com/shubhamkandiyal0101/go-basic-ecommerce-app/controllers"
)

func smtpAuth() smtp.Auth {
	hostname := "smtp.gmail.com"
	auth := smtp.PlainAuth("", "naco.shubhamkandiyal@gmail.com", "ahxpbvebkuzsxgoh", hostname)

	return auth
}

// Get All Products
func Products(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- Get All Products -- ")
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	response := mongocontroller.GetAllProductsDb()

	json.NewEncoder(w).Encode(response)

	return

}

// ends here ~ Get All Products

// GET SINGLE PRODUCT
func GetSingleProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- GET SINGLE PRODUCT -- ")
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r)
	product_id := params["product_id"]

	isProdIdExists := mongocontroller.IsProdIdExists(product_id)

	if !isProdIdExists {
		response := map[string]interface{}{
			"message": "Please provide Valid Product ID",
			"data": map[string]interface{}{
				"product_id": product_id,
			},
		}
		json.NewEncoder(w).Encode(response)
		return
	}

	response := mongocontroller.GetSingleProductDetailsDb(product_id)
	json.NewEncoder(w).Encode(response)
	return

}

// ENDS HERE ~ GET SINGLE PRODUCT

// Send Test Email
func SendTestEmail(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- Send Test Email -- ")
	w.Header().Set("Content-Type", "application/json")
	defer r.Body.Close()

	type emailTemplateVars struct {
		Name     string
		Link     string
		LinkText string
	}

	temlateData := &emailTemplateVars{"Deepak", "https://instagram.com/democraticnaari", "Democratic Naari"}

	parseTemplate := template.Must(template.ParseFiles("emailtemplates/sample.html"))

	var tpl bytes.Buffer

	parseTemplate.Execute(&tpl, temlateData)

	// fmt.Println(tpl.String())

	auth := smtpAuth()

	to := []string{"shubhamkandiyal0101@gmail.com"}
	// msg := []byte("dummy message")
	// msg := []byte("To: recipient@example.net\r\n" +
	// 	"Subject: discount Gophers!\r\n" +
	// 	"\r\n" +
	// 	"This is the email body.\r\n")

	subject := "Subject: Test email from Go!\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"
	// body := "<html><body><h1>Hello World!</h1></body></html>"
	body := tpl.String()
	from := mail.Address{"Go Shopping App", "naco.shubhamkandiyal@gmail.com"}
	fromString := "From: " + from.String() + "\r\n"

	msg := []byte(fromString + subject + mime + body)

	err := smtp.SendMail("smtp.gmail.com:587", auth, from.Address, to, msg)
	if err != nil {
		log.Fatal(err)
	}

	response := map[string]string{
		"Message": "Email Sent Successfully",
	}
	json.NewEncoder(w).Encode(response)
}

// ends here ~ Send Test Email

// get all categories
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
	fmt.Println(" -- -- GET ALL CATEGORIES -- -- ")
	defer r.Body.Close()

	w.Header().Set("Content-Type", "application/json")

	response, responseCode := mongocontroller.GetAllCategories()

	w.WriteHeader(responseCode)

	json.NewEncoder(w).Encode(response)
}

// ends here ~ Get All Categories
