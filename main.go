package main

import "github.com/vehsamrak/wishlist/src"

func main() {
    server := src.WishService{}
    server.Start()
}
