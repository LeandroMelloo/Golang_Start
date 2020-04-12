package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	res, _ := http.Get("http://google.com")
	body, _ := ioutil.ReadAll(res.Body)
	res.Body.Close() // Close() está em outro pacote tem que estar com a letra maiuscula.
	fmt.Printf("%s", body)
}
