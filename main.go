package main

import (
	"golang-restfulapi/app"
	"golang-restfulapi/controller"
	"golang-restfulapi/helper"
	"golang-restfulapi/middleware"
	"golang-restfulapi/repository"
	"golang-restfulapi/service"
	"net/http"

	_ "github.com/go-sql-driver/mysql"

	"github.com/go-playground/validator/v10"
	
)

func main() {
	validate := validator.New()
	db :=app.NewDB()
	categoryRepository := repository.NewCategoryRepository()
	categoryService	:= service.NewCategoryService(categoryRepository, db, validate)
	categoryController := controller.NewCategoryController(categoryService)
	
	router := app.NewRouter(categoryController)
	server := http.Server{
		Addr: "localhost:3000",
		Handler: middleware.NewAuthMiddleware(router),
	}

	err := server.ListenAndServe()
	helper.PanicIfError(err)
}