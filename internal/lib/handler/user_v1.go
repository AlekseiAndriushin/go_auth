package handler

import (
	"context"
	"crypto/rand"
	"fmt"
	"math/big"
	"strings"

	"github.com/AlekseiAndriushin/go_auth/internal/lib/logger"
	desc "github.com/AlekseiAndriushin/go_auth/pkg/user_v1"
	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserRPCServerV1 struct {
    desc.UnimplementedUserV1Server
}

func NewUserRPCServerV1() *UserRPCServerV1 {
    return &UserRPCServerV1{}
}

func (s *UserRPCServerV1) Create(ctx context.Context, req *desc.CreateRequest) (*desc.CreateResponse, error) {
	resStr := fmt.Sprintf("Received Create:\n\tName: %v,\n\tEmail: %v,\n\tPassword: %v,\n\tPassword confirm: %v,\n\tRole: %v\n",
		req.GetName(),
		req.GetEmail(),
		req.GetPassword(),
		req.GetPasswordConfirm(),
		req.GetRole(),
	)
  logger.LogInfo(resStr)

	if dline, ok := ctx.Deadline(); ok {
		logger.LogInfo(fmt.Sprintf("Deadline: %v", dline))
	}

	randInt64, err := rand.Int(rand.Reader, new(big.Int).SetInt64(1<<62))
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	id := randInt64.Int64()

	respStr := fmt.Sprintf("Response Create:\n\tId: %v\n", id)
	logger.LogInfo(respStr)

	return &desc.CreateResponse{
		Id: id,
	}, nil
}

func (s *UserRPCServerV1) Get(ctx context.Context, req *desc.GetRequest) (*desc.GetResponse, error) {
	logger.LogInfo(fmt.Sprintf("Received Get:\n\tId: %v", req.GetId()))

	if dline, ok := ctx.Deadline(); ok {
		logger.LogInfo(fmt.Sprintf("Deadline: %v", dline))
	}

	role := "ADMIN" // Вместо gofakeit используем жестко заданные значения
	resp := desc.GetResponse{
		Id:        req.GetId(),
		Name:      gofakeit.Name(),
		Email:     gofakeit.Email(),
		Role:      desc.Role(desc.Role_value[role]),
		CreatedAt: timestamppb.New(gofakeit.Date()),
		UpdatedAt: timestamppb.New(gofakeit.Date()),
	}

	respStr := fmt.Sprintf("Response Get:\n\tId: %v,\n\tName: %v,\n\tEmail: %v,\n\tRole: %v,\n\tCreatedAt: %v,\n\tUpdatedAt: %v\n",
		resp.Id,
		resp.Name,
		resp.Email,
		resp.Role,
		resp.CreatedAt,
		resp.UpdatedAt,
	)

	logger.LogInfo(respStr)

	return &resp, nil
}

func (s *UserRPCServerV1) Update(ctx context.Context, req *desc.UpdateRequest) (*emptypb.Empty, error) {
	buf := strings.Builder{}
	buf.WriteString("Received Update:\n")
	idStr := fmt.Sprintf("\tId: %v\n", req.GetId())
	buf.WriteString(idStr)

	if req.Name != nil {
		buf.WriteString(fmt.Sprintf("\tName: %v\n", req.GetName().GetValue()))
	}

	if req.Email != nil {
		buf.WriteString(fmt.Sprintf("\tEmail: %v\n", req.GetEmail().GetValue()))
	}

	if req.Role != desc.Role_UNKNOWN {
		buf.WriteString(fmt.Sprintf("\tRole: %v\n", req.GetRole()))
	}

	logger.LogInfo(buf.String())

	return &emptypb.Empty{}, nil
}

func (s *UserRPCServerV1) Delete(ctx context.Context, req *desc.DeleteRequest) (*emptypb.Empty, error) {
	logger.LogInfo(fmt.Sprintf("Received Delete:\n\tId: %v", req.GetId()))

	return &emptypb.Empty{}, nil
}