//go:build wireinject
// +build wireinject

package main

import (
	"golang-rest-api/app"
	"golang-rest-api/controller"
	"golang-rest-api/middleware"
	"golang-rest-api/repository"
	"golang-rest-api/service"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/wire"
	"github.com/julienschmidt/httprouter"
)

// var categorySet = wire.NewSet(
//
//	repository.NewCategoryRepository,
//	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
//	service.NewCategoryService,
//	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
//	controller.NewCategoryController,
//	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
//
// )
func ProvideValidator() *validator.Validate {
	return validator.New()
}

var categorySet = wire.NewSet(
	repository.NewCategoryRepository,
	wire.Bind(new(repository.CategoryRepository), new(*repository.CategoryRepositoryImpl)),
	service.NewCategoryService,
	wire.Bind(new(service.CategoryService), new(*service.CategoryServiceImpl)),
	controller.NewCategoryController,
	wire.Bind(new(controller.CategoryController), new(*controller.CategoryControllerImpl)),
)

//other set

func InitializeServer() *http.Server {
	wire.Build(
		app.NewDB,
		ProvideValidator,
		categorySet,
		app.NewRouter,
		wire.Bind(new(http.Handler), new(*httprouter.Router)),
		middleware.NewAuthMiddleware,
		NewServer,
	)
	return nil
}
