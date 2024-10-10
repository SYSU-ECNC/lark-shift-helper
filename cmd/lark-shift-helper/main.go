package main

import (
	"fmt"

	"github.com/SYSU-ECNC/lark-shift-helper/internal/config"
)

func main() {
	var cfg config.Config
	cfg.Parse()

	fmt.Println("APP ID:", cfg.APPID)
	fmt.Println("APP Secret:", cfg.APPSecret)
}
