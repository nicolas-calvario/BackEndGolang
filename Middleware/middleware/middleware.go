package middleware

import (
	"fmt"
	"time"
)

type MyFunction func(string)

func MiddlewareLog(f MyFunction) MyFunction {
	return func(name string) {
		fmt.Println("Inicio: ", time.Now().Format("15:04:04"))
		f(name)
		fmt.Println("Fin ", time.Now().Format("15:04:04"))
	}
}
