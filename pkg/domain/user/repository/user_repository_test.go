package repository

import (
	"database/sql"
	"fmt"
	"testing"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/user/model"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/Mangaba-Labs/ape-finance-api/pkg/domain/config"
	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/assert"
)

var db *sql.DB

func TestRepository(t *testing.T) {

	var err error
	pool, err := dockertest.NewPool("")

	if err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}

	resource, err := pool.Run("postgres", "9.6", []string{"POSTGRES_USER=postgres", "POSTGRES_PASSWORD=postgres", "POSTGRES_DB=tempoo"})

	if err != nil {
		t.Fatalf("Could not start resource: %s", err)
	}

	if err = pool.Retry(func() error {
		var err error
		db, err = sql.Open("postgres", fmt.Sprintf("postgres://postgres:postgres@localhost:%s/%s?sslmode=disable", resource.GetPort("5432/tcp"), "tempoo"))
		if err != nil {
			return err
		}
		return db.Ping()
	}); err != nil {
		t.Fatalf("Could not connect to docker: %s", err)
	}

	gormDB, err := gorm.Open(postgres.New(postgres.Config{
		Conn: db,
	}), &gorm.Config{})

	if err != nil {
		t.Fatalf("could not create gorm instance")
	}

	r := NewUserRepository(gormDB)

	migrationTool := config.Migrate{DB: gormDB}

	migrationTool.DB.AutoMigrate(&model.User{})

	r.Create(&model.User{
		Model:    gorm.Model{},
		Email:    "matheus.cumpian@hotmail.com",
		Name:     "Matheus Cumpian",
		Password: "20012000",
	})

	t.Run("Get all user records", func(t *testing.T) {

		users, err := r.FindAll()

		assert.Nil(t, err, "cannot find all users")

		fmt.Println(users)

		assert.Equal(t, len(users), 1)
	})

	t.Run("Find user by email", func(t *testing.T) {
		user, err := r.FindOneByEmail("matheus.cumpian@hotmail.com")

		assert.Nil(t, err, "cannot find user")

		assert.NotEmpty(t, user, "cannot find user")
	})

	t.Run("Find user by id", func(t *testing.T) {
		user, err := r.FindByID(1)

		assert.Nil(t, err, "cannot find user")

		assert.NotEmpty(t, user, "cannot find user")
	})

	t.Run("Delete user", func(t *testing.T) {
		err := r.Delete(1)

		assert.Nil(t, err, "cannot find user")

		user, err := r.FindOneByEmail("matheus.cumpian@hotmail.com")

		assert.Empty(t, user, "user not deleted")
	})

	t.Cleanup(func() {
		if err = pool.Purge(resource); err != nil {
			t.Fatalf("Could not purge resource: %s", err)
		}
	})

}
