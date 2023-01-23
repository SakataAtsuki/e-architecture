package grpc

import (
	"time"

	"github.com/SakataAtsuki/e-architecture/pkg/proto/api"
	"github.com/SakataAtsuki/e-architecture/pkg/usecase"
)

// var _ api.EArchitectureServer = (*Service)(nil)

type Service struct {
	api.UnimplementedEArchitectureServer
	uc  usecase.Usecase
	now func() time.Time
}

func New(uc usecase.Usecase) *Service {
	return &Service{uc: uc, now: time.Now}
}
