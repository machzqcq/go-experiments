package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8000", http.FileServer(http.Dir("src/root_mdot_10_")))
}
