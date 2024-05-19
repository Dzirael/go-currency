package main

import (
	_ "github.com/Dzirael/go-curenncy/docs"
	"github.com/Dzirael/go-curenncy/internal/api"
	"github.com/Dzirael/go-curenncy/internal/email"
)

func main() {
	email.Run()
	api.Run()
}
