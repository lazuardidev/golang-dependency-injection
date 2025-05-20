package main

import (
	"belajar-golang-dependency-injection/app"
	"belajar-golang-dependency-injection/controller"
	"belajar-golang-dependency-injection/helper"
	"belajar-golang-dependency-injection/middleware"
	"belajar-golang-dependency-injection/repository"
	"belajar-golang-dependency-injection/service"
	"net/http"

	"github.com/go-playground/validator"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := app.NewDB()
	validate := validator.New()
	categoryRepository := repository.NewCategoryRepository()
	categoryService := service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	router := app.NewRouter(categoryController)

	server := http.Server{
		Addr:    "localhost:3002",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}
