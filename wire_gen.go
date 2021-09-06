// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package main

import (
	"github.com/Mangaba-Labs/ape-finance-api/database"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/api/router"
	handler3 "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/auth/handler"
	handler5 "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/handler"
	repositoryCategory "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/repository"
	servicesCategory "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/category/service"
	handler4 "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/handler"
	repositoryStock "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/repository"
	servicesStock "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/stock/services"
	handler6 "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/handler"
	repositoryTransaction "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/repository"
	serviceTransaction "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/transaction/service"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/handler"
	repositoryUser "github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/repository"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/services"
)

// Injectors from wire.go:

func initializeServer() (*router.Server, error) {
	db, err := database.NewDatabase()
	if err != nil {
		return nil, err
	}
	userRepository := repositoryUser.Repository{
		DB: db,
	}
	stockRepository := repositoryStock.Repository{
		DB: db,
	}

	categoryRepository := repositoryCategory.NewCategoryRepository(db)
	transactionRepository := repositoryTransaction.NewTransactionRepository(db)

	userService := services.NewUserService(userRepository)
	handlerHandler := handler.NewUserHandler(userService)
	authHandler := handler3.NewAuthHandler(userService)
	categoryService := servicesCategory.NewCategoryService(categoryRepository)
	categoryHandler := handler5.NewCategoryHandler(categoryService)
	stockService := servicesStock.NewUserService(stockRepository)
	stockHandler := handler4.NewStockHandler(stockService)
	transactionService := serviceTransaction.NewTransactionService(transactionRepository)
	transactionHandler := handler6.NewTransactionHandler(transactionService)
	server := router.NewServer(handlerHandler, authHandler, stockHandler, categoryHandler, transactionHandler)
	return server, nil
}
