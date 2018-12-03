package main

import (
	"fmt"
	"goapp/routes"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gorilla/handlers"
)

// These two lines are important in order to allow access from the front-end side to the methods
//var allowedOrigins = handlers.AllowedOrigins([]string{"*"})

var allowedOrigins = handlers.AllowedOrigins([]string{"*"})
var allowedMethods = handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT", "OPTIONS", "HEAD"})
var allowedHeaders = handlers.AllowedHeaders([]string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "X-XSRF-Token", "X-HTTP-Method-Override", "X-Requested-With", "Mobile-Cookie"})

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

	gracefulStop := make(chan os.Signal)
	signal.Notify(gracefulStop, syscall.SIGTERM)
	signal.Notify(gracefulStop, syscall.SIGINT)
	go func() {
		sig := <-gracefulStop
		fmt.Printf("\n\tcaught sig: %+v\t\n", sig)
		fmt.Println("\tStopping Http Server")
		time.Sleep(10 * time.Nanosecond)
		os.Exit(0)
	}()

	router := routes.GetRoutes()

	// This will serve files under http://IP:PORT/assets/<filename>
	router.PathPrefix("/assets/").Handler(http.FileServer(http.Dir(".")))

	Server := &http.Server{
		Handler: handlers.CORS(allowedOrigins, allowedMethods)(router),
		Addr:    "127.0.0.1:8080",
	}

	fmt.Println(Server.ListenAndServe())
}
