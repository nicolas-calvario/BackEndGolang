package main

import (
	"go-middleware/funciones"
	"go-middleware/middleware"
)

func main() {
	name := "Nicolas"
	execute(name, middleware.MiddlewareLog(funciones.Saludar))
	execute(name, middleware.MiddlewareLog(funciones.Despedirse))
}

func execute(name string, f middleware.MyFunction) {
	f(name)
}
