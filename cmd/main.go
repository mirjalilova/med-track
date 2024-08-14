package main

import (
	"github.com/mirjalilova/med-track/config"
	"github.com/mirjalilova/med-track/pkg/app"
)

func main() {
	cnf := config.Load()

	app.Run(&cnf)
}