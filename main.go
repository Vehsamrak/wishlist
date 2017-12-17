package main

import "github.com/vehsamrak/wishlist/src"

func main() {
    wishService := src.WishService{}
    wishService.Start()
}
