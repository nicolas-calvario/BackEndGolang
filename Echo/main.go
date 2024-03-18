package main

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

func main() {
	e := echo.New()
	e.GET("/", saludar)
	e.GET("/div", dividir)

	// e.POST("/personas/crear", crear)
	// e.GET("/personas/consulta", consulta)
	// e.PUT("/personas/actualiza", actulizar)
	// e.DELETE("/personas/eliminar", eliminar)
	//grupos de rutas
	personas := e.Group("/personas")
	personas.Use(MiddlewareLogPersona)
	personas.POST("", crear)
	personas.GET("/:id", consulta)
	personas.PUT("/:id", actulizar)
	personas.DELETE("/:id", eliminar)

	e.Logger.Fatal(e.Start(":1323"))
}

func saludar(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"saludo": "hola"})
}

func dividir(e echo.Context) error {
	d := e.QueryParam("d")
	f, _ := strconv.Atoi(d)
	if f <= 0 {
		return e.String(http.StatusBadRequest, "El valor no puede ser menor o igual a 0")

	}
	result := 3000 / f

	return e.String(http.StatusOK, strconv.Itoa(result))
}

func crear(e echo.Context) error {
	return e.JSON(http.StatusOK, map[string]string{"Status": "Creado"})
}

func consulta(e echo.Context) error {
	id := e.Param("id")
	return e.JSON(http.StatusOK, map[string]string{"Status": "Consultado " + id})
}
func actulizar(e echo.Context) error {
	id := e.Param("id")
	return e.JSON(http.StatusOK, map[string]string{"Status": "actualizado " + id})
}
func eliminar(e echo.Context) error {
	id := e.Param("id")
	return e.JSON(http.StatusOK, map[string]string{"Status": "eliminado " + id})
}

// Middleware
func MiddlewareLogPersona(f echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		log.Printf("Peticion realizada a /pesonas")
		return f(c)
	}
}
