package main

import (
	"log"

	"github.com/azybk/simple-forum/internal/configs"
	"github.com/azybk/simple-forum/internal/handler/memberships"
	membershipRepo "github.com/azybk/simple-forum/internal/repository/memberships"
	"github.com/azybk/simple-forum/pkg/internalsql"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	var cfg *configs.Config

	err := configs.Init(
		configs.WithConfigFolder([]string{"./internal/configs"}),
		configs.WithConfigFile("config"),
		configs.WithConfigType("yaml"),
	)

	if err != nil {
		log.Fatal("Gagal inisiasi config", err)
	}

	cfg = configs.GetConfig()
	log.Println("config", cfg)

	db, err := internalsql.Connect(cfg.Database.DataSourceName)
	if err != nil {
		log.Fatal("Gagal inisiasi database", err)
	}

	_ = membershipRepo.NewRepository(db)

	memberships := memberships.NewHandler(r)
	memberships.RegisterRoute()

	r.Run(cfg.Service.Port)
}
