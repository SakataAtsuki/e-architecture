package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/SakataAtsuki/e-architecture/pkg/entity"
	"github.com/SakataAtsuki/e-architecture/pkg/repository"
	"github.com/SakataAtsuki/e-architecture/pkg/repository/postgres"
	"github.com/SakataAtsuki/e-architecture/pkg/usecase"
	"github.com/SakataAtsuki/e-architecture/pkg/util/errcode"
)

func main() {
	uri := fmt.Sprintf("postgres://%s/%s?sslmode=disable&user=%s&password=%s&port=%s&timezone=Asia/Tokyo",
		os.Getenv("DB_HOST"), os.Getenv("DB_NAME"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_PORT"))
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Println(errcode.New(err))
		os.Exit(1)
	}
	if err := db.Ping(); err != nil {
		log.Println(errcode.New(err))
		os.Exit(1)
	}
	log.Println("successfully connected to database")

	ctx := context.Background()
	cfg := &usecase.Config{
		DB: &repository.Database{User: postgres.NewUser(db)},
	}
	uc := usecase.New(cfg)

	req := &usecase.CreateUserRequest{
		User: &entity.User{
			ID:   "test-id",
			Name: "test-name",
		},
	}

	resp, err := uc.CreateUser(ctx, req)
	if err != nil {
		log.Println(errcode.New(err))
		os.Exit(1)
	}
	log.Println(resp.User.ID, resp.User.Name)
}
