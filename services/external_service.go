package services

import (
	"net/http"
	"io"
)
// Fungsi untuk mengambil produk eksternal dari API
func GetExternalProducts() ([]byte, error) {
	resp, err := http.Get("https://dummyjson.com/products")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	return body, err
}