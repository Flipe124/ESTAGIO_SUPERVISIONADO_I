// Package main (principal) é o pacote principal que start a aplicação.
package main

import (
	"log"
	"net/http"
	"server/internal/pkg/ipa"
)

func main() {

	log.Fatal(http.ListenAndServe(":"+ipa.API.GetPortString(), nil))

}
