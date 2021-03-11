package user

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/handler"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/repository"
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/services"
	"github.com/google/wire"
)

// Set wire for user
var Set = wire.NewSet(
	repository.NewUserRepository,
	repository.Set,
	services.NewUserService,
	handler.NewUserHandler,
)
