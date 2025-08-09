package main

import (
	"fmt"

	"github.com/ehaxi/url-shortener/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)
}
