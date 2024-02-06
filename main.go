package main

import (
	"fmt"
	"github.com/latindarkrai/luhn/car"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	car1 := car.NewCar("Mercedez", 4)
	fmt.Println(car1)
	var car2 car.Car = car.Car{}
	car2.Name = "Audi :("
	fmt.Println(car2.String())

	twoArguments(true, "/")
	http.HandleFunc("/", luhnHandler)

	log.Fatal(http.ListenAndServe(":8080", nil))
	// fmt.Println("Hello World")
	// luhn := isLuhn("4242424242424243")
	//
	// if luhn {
	// 	fmt.Println("valid credit card number")
	// } else {
	// 	fmt.Println("not a valid credit card number")
	// }
}

func luhnHandler(writer http.ResponseWriter, request *http.Request) {
	// var luhn bool
	creditCardNumber := strings.Trim(request.URL.Path, "/")
	luhn := isLuhn(creditCardNumber)
	message := "The credit card number %s is %s!"
	var status string

	if luhn {
		status = "valid"
	} else {
		status = "invalid"
	}

	fmt.Fprintf(writer, message, creditCardNumber, status)
	// fmt.Fprintf(w, "luhn\n")
	// fmt.Fprintf(w, "Hello, %q\n", html.EscapeString(r.URL.Path))
}

func zeroArguments() string {
	myString := "hello world"
	return myString
}

func twoArguments(boolean bool, stringy string) {
	fmt.Println(boolean, stringy)
	return
}

// func declares the functions intent
// function name
// argument inside the parantheses
// return types at the end before the curly brace
// inside the curly brace is the code for the function
func isLuhn(cardNumber string) bool {
	var sum int
	parity := len(cardNumber) % 2

	// for (initialization; condition; step)
	// for i=0; i<10; i++
	for i := len(cardNumber) - 1; i >= 0; i-- {
		digit, _ := strconv.Atoi(string(cardNumber[i]))

		if i%2 == parity {
			digit = digit * 2
		}

		if digit > 9 {
			digit -= 9
		}

		sum += digit
	}

	return sum%10 == 0
}
