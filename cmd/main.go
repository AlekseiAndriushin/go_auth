package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	"github.com/AlekseiAndriushin/go_auth/internal/config"
	"github.com/AlekseiAndriushin/go_auth/internal/lib/handler"
	"github.com/AlekseiAndriushin/go_auth/internal/lib/logger"
	"github.com/AlekseiAndriushin/go_auth/pkg/user_v1"
	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.MustConfig()
	iLog := log.New(os.Stdout, color.CyanString("[INFO] "), log.LstdFlags)

	iLog.Println("Starting auth service...")

	url := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.GRPCPort)

	lis, err := net.Listen("tcp", url)
	if err != nil {
		errStr := fmt.Sprintf("failed to listen: %v", err)
		logger.LogError(errStr) // Используем новый логгер
		os.Exit(1)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	rpcSrvV1 := handler.NewUserRPCServerV1() 
	user_v1.RegisterUserV1Server(s, rpcSrvV1)

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.Serve(lis); err != nil {
			errStr := fmt.Sprintf("failed to serve: %v", err)
			logger.LogError(errStr)
			os.Exit(1)
		}
	}()

	iLog.Println(color.GreenString("Service started successfully "), color.BlueString(url))

	<-done
	s.GracefulStop()
	iLog.Println(color.YellowString("Service stopped"))
}
