package main

import (
	"context"
	"os"
	"os/signal"
	"rlservice/internal/config"
	"rlservice/internal/delivery/grpc"
	"rlservice/internal/infra/mredis"
	"rlservice/internal/infra/postgresql"
	"rlservice/internal/usecase"
	"rlservice/pkg/logger"
	"syscall"
	"time"
)

func main() {
	cfg := config.MustUpload()
	log := logger.NewLogger(cfg.Env)

	psql, err := postgresql.NewPsqlStorage(log, &cfg.Psql)
	if err != nil {
		panic(err)
	}

	mred, err := mredis.NewRedisStorage(log, &cfg.Redis)
	if err != nil {
		panic(err)
	}

	su := usecase.NewServiceUsecase(log, psql)
	scu := usecase.NewScriptUsecase(log, mred)

	app := grpc.NewApplication(log, scu, su, cfg.Port)

	go app.MustRun()
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan
	_, ctxRelease := context.WithTimeout(context.Background(), 5*time.Second)
	defer ctxRelease()
	if err := app.Shutdown(); err != nil {
		panic(err)
	}
	log.Info("main", "application shutdown")
}
