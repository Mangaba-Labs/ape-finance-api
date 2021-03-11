// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/api/router"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/auth"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user"
	"github.com/google/wire"
)

func initializeServer() (*router.Server, error) {

	wire.Build(user.Set, auth.Set, router.NewServer)

	return &router.Server{}, nil
}
