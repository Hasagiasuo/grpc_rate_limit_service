package grpc

import (
	"fmt"
	"net"

	pb "rlservice/gen/v1"
	"rlservice/internal/grpc/rlgrpc"
	"rlservice/pkg/logger"

	"google.golang.org/grpc"
)

type Application struct {
	log      *logger.Logger
	server   *grpc.Server
	listener net.Listener
	port     int
}

func NewApplication(log *logger.Logger, scriptProvider rlgrpc.ScriptProvider, serviceProvider rlgrpc.ServiceProvider, port int) *Application {
	grpcServer := grpc.NewServer()
	service := rlgrpc.NewService(scriptProvider, serviceProvider, log)
	pb.RegisterRateLimitServiceServer(grpcServer, service)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		panic(fmt.Sprintf("cannot run tcp listener: %v", err))
	}
	return &Application{
		log:      log,
		server:   grpcServer,
		port:     port,
		listener: listener,
	}
}

func (a *Application) MustRun() {
	const op = "grpc.MustRun"
	a.log.Info(op, fmt.Sprintf("tcp listener started at :%d", a.port))
	a.server.Serve(a.listener)
}

func (a *Application) Shutdown() error {
	const op = "grpc.Shutdown"
	a.log.Info(op, "listener closed")
	return a.listener.Close()
}
