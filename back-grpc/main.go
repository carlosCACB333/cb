package main

import (
	"context"
	"net"
	"net/http"
	"os"

	"github.com/carlosCACB333/cb-grpc/pb"
	"github.com/carlosCACB333/cb-grpc/servers"
	"github.com/carlosCACB333/cb-grpc/utils"
	"github.com/carlosCACB333/cb-grpc/worker"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	cfg, err := utils.LoadConfig()
	if err != nil {
		panic(err)
	}

	if cfg.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	dialector := postgres.Open(cfg.DBUrl)
	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// model.Migrate(db)
	redisOpts := asynq.RedisClientOpt{
		Addr: cfg.RedisAddr,
	}

	taskDistributor := worker.NewRedisTaskDistributor(redisOpts)

	go runGatewayServer(cfg, db, taskDistributor)
	go runTaskProcessor(cfg, db, &redisOpts)
	runGrpcServer(cfg, db, taskDistributor)

}

func runGrpcServer(cfg *utils.Config, db *gorm.DB, distributor worker.TaskDistributor) {

	authServer := servers.NewAuthServer(cfg, db, distributor)
	postServer := servers.NewPostServer(cfg, db)

	logger := grpc.UnaryInterceptor(utils.GrpcLogger)
	grpcServer := grpc.NewServer(logger)

	pb.RegisterAuthServiceServer(grpcServer, authServer)
	pb.RegisterPostServiceServer(grpcServer, postServer)
	reflection.Register(grpcServer)

	listener, err := net.Listen("tcp", ":"+cfg.GrpcPort)
	if err != nil {
		log.Fatal().Msg("cannot start listener: ")
	}

	log.Info().Msg("Running gRPC server on " + listener.Addr().String())

	defer listener.Close()

	if err := grpcServer.Serve(listener); err != nil {
		log.Fatal().Msg("cannot start server: ")
	}

}

func runGatewayServer(cfg *utils.Config, db *gorm.DB, distributor worker.TaskDistributor) {

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	authServer := servers.NewAuthServer(cfg, db, distributor)
	grpcMux := runtime.NewServeMux(
		runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.JSONPb{
			MarshalOptions:   protojson.MarshalOptions{UseProtoNames: true},
			UnmarshalOptions: protojson.UnmarshalOptions{DiscardUnknown: true},
		}),
	)

	err := pb.RegisterAuthServiceHandlerServer(ctx, grpcMux, authServer)

	if err != nil {
		log.Fatal().Msg("cannot register gateway: ")
	}

	mux := http.NewServeMux()
	mux.Handle("/", grpcMux)

	listener, err := net.Listen("tcp", ":"+cfg.HttpPort)
	if err != nil {
		log.Fatal().Msg("cannot start listener: ")
	}

	log.Info().Msg("Running gRPC gateway on " + listener.Addr().String())

	defer listener.Close()

	if err := http.Serve(listener, utils.HttpLogger(mux)); err != nil {
		log.Fatal().Msg("cannot start gateway: ")
	}

}

func runTaskProcessor(cfg *utils.Config, db *gorm.DB, opts *asynq.RedisClientOpt) {
	processor := worker.NewRedisTaskProcessor(cfg, *opts, db)
	log.Info().Msg("Running task processor")
	if err := processor.Start(); err != nil {
		log.Fatal().Msg("cannot start task processor: ")
	}
}
