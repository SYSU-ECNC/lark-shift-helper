package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	var (
		APPID     string
		APPSecret string
	)

	flag.StringVar(&APPID, "i", "", "APP ID (required)")
	flag.StringVar(&APPSecret, "s", "", "APP Secret (required)")
	flag.Parse()

	if APPID == "" || APPSecret == "" {
		fmt.Println("Error: APP ID and APP Secret are required")
		flag.Usage()
		os.Exit(1)
	}

	fmt.Println("APP ID:", APPID)
	fmt.Println("APP Secret:", APPSecret)
}
