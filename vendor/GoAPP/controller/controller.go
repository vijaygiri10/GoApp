package controller

import (
	"fmt"
	"net/http"
)

//GetOrderswithStatus List all Orders With thier Status
func GetOrderswithStatus(w http.ResponseWriter, res *http.Request) {
	fmt.Fprintf(w, "GetOrderswithStatus")
}

//UpdatePaymentMethod
//	Cases Include
//		1.)Update the expiration date of the credit card
//		2.)Change the credit card number
//		3.)Change payment method
func UpdatePaymentMethod(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "UpdatePaymentMethod")
}

//UpdateCustomerContactInfo
//	Case Include
//		1.)change email address
//		2.)change name
//		3.)change contact phone number
//		4.)change mailing address
func UpdateCustomerContactInfo(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "UpdateCustomerContactInfo")
}
