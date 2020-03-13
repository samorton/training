package hello

import (
	"log"

	"rsc.io/quote"
)

func Hello() string {

	log.Print(quote.Glass())
	return quote.Hello()

}
