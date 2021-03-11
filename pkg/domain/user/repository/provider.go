package repository

import (
	"github.com/Mangaba-Labs/ape-finance-api/database"
	"github.com/google/wire"
)

// Set wire
var Set = wire.NewSet(
	database.NewDatabase,
	wire.Struct(new(Repository), "DB"),
)
