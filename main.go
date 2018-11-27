package main

import (
	"GoAPP/routes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
)

// These two lines are important in order to allow access from the front-end side to the methods
//var allowedOrigins = handlers.AllowedOrigins([]string{"*"})

var allowedOrigins = handlers.AllowedOrigins([]string{"res.bhn.net"})
var allowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})

/*
The End Customer portal is requirement for all merchants who offer subscription-based products.
This Iteration 1 implementation is strictly to allow customers to update their billing information.
See ##3914887 for Iteration 2 features
* A branded URL for every merchant
* Secure login for every customer
* Change/Update payment method -- typical use cases include
	1.)update the expiration date of the credit card
	2.)change the credit card number
	3.)change payment method
* Update customer contact info
	1.)change email address
	2.)change name
	3.)change contact phone number
	4.)change mailing address
* Change my Login and/or Password
* View list of orders and their current status
* Able to cancel a subscription (a recurring order)
*/
func main() {
	//var cancelfunc context.CancelFunc

	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("caught sig: %+v\t\n", sig)
		fmt.Println("Stopping Http Server")
		time.Sleep(10 * time.Nanosecond)
		os.Exit(0)
	}()

	router := routes.GetRoutes()

	//http.Handle("/assets/", http.StripPrefix("/assets/", http.FileServer(http.Dir("assets"))))
	router.Handle("/assets/", http.StripPrefix("/assets", http.FileServer(http.Dir("./assets"))))
	// launch server with CORS validations
	http.ListenAndServe("173.168.101.178:9090", handlers.CORS(
		allowedOrigins, allowedMethods)(router))

}