package main

import (
	"api"
	"net/http"
)

func main() {
	http.HandleFunc("/upload", api.UploadImage)
	http.ListenAndServe(":8000", nil)
}
