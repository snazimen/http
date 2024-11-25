package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// функция проверяющая заголовок, если админ возвращает запись в консоль
func checkRoleMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		userRole := c.Request().Header.Get("User-Role")
		//проверяем заголовок
		if userRole == "admin" {
			fmt.Println("admin tedected")

		}
		return next(c)
	}
}

func DateToUntil(c echo.Context) error {

	n := time.Date(2025, time.January, 01, 00, 00, 00, 00, time.UTC)
	a := time.Now()
	// расчитвыем дни до 2025 года 1 января
	// от 2025 отнимаем текущую дату и делм на 24, чтобы получить дни
	daysintil := int(n.Sub(a).Hours() / 24)
	//переменная выводящая ответ в виде интерфейса
	response := map[string]interface{}{
		"days_until_2025": daysintil,
	}
	//возвращаем статус код 200 и дни до 2025 года
	return c.JSON(http.StatusOK, response)
}
func main() {
	e := echo.New()

	// записывает инфу о запросах logger
	e.Use(middleware.Logger())
	//обрабатывает панические ошибки
	e.Use(middleware.Recover())
	//Проверяем с какой роли идет запрос
	e.Use(checkRoleMiddleware)

	// получение иноформации из функции date
	e.GET("/daysintil2025", DateToUntil)
	// старт сервера
	e.Logger.Fatal(e.Start(":8888"))
}
