package routes

import (
	controller "GoAPP/controller"
	logger "GoAPP/logger"
	"GoAPP/template"

	//"fmt"
	"net/http"

	mux "github.com/gorilla/mux"
)

type route struct {
	FuncName    string
	MethodType  string
	URLPattern  string
	HandlerFunc http.HandlerFunc
}

//Routes defines the list of routes of our API
var Routes = []route{
	route{
		//Display Index Page
		FuncName:    "Index",
		MethodType:  "GET",
		URLPattern:  "/",
		HandlerFunc: template.Index,
	},
	route{
		//Display About Page
		FuncName:    "About",
		MethodType:  "GET",
		URLPattern:  "/about",
		HandlerFunc: template.About,
	},
	route{
		//Display Contact Page
		FuncName:    "Contact",
		MethodType:  "GET",
		URLPattern:  "/contact",
		HandlerFunc: template.Contact,
	},
	route{
		//Display Signup Page
		FuncName:    "Signup",
		MethodType:  "POST",
		URLPattern:  "/signup",
		HandlerFunc: template.Signup,
	},
	route{
		//Display Signup Page
		FuncName:    "Signup",
		MethodType:  "GET",
		URLPattern:  "/signup",
		HandlerFunc: template.Signup,
	},
	route{
		//Display Process Page
		FuncName:    "Process",
		MethodType:  "POST",
		URLPattern:  "/process",
		HandlerFunc: template.Process,
	},

	route{
		//View list of orders and their current status
		FuncName:    "GetOrderswithStatus",
		MethodType:  "GET",
		URLPattern:  "/GetOrderswithStatus",
		HandlerFunc: controller.GetOrderswithStatus,
	},
	route{
		//Change/Update payment method -- typical use cases include
		FuncName:    "UpdatePaymentMethod",
		MethodType:  "PUT",
		URLPattern:  "/UpdatePaymentMethod",
		HandlerFunc: controller.UpdatePaymentMethod,
	},
	route{
		//Update customer contact info
		FuncName:    "UpdateCustomerContactInfo",
		MethodType:  "PUT",
		URLPattern:  "/UpdateCustomerContactInfo",
		HandlerFunc: controller.UpdateCustomerContactInfo,
	},
}

//GetRoutes configures a new router to the API
func GetRoutes() *mux.Router {
	MuxRouter := mux.NewRouter().StrictSlash(true)
	for _, route := range Routes {
		var handler http.Handler
		handler = route.HandlerFunc

		MuxRouter.Methods(route.MethodType).Path(route.URLPattern).Name(route.FuncName).Handler(logger.Logger(handler, route.FuncName))
	}
	return MuxRouter
}
