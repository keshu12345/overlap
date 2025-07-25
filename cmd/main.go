package main

import "github.com/keshu12345/overlap/internal/router"

func main() {
	r := router.InitRouter()
	r.Run(":8080")
}
