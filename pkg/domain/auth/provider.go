package auth

import (
	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/auth/handler"
	"github.com/google/wire"
)

var Set = wire.NewSet(
	handler.NewAuthHandler,
)
