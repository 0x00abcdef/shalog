package image

import (
	"context"
	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/icowan/blog/src/repository/types"
	"time"
)

type loggingServer struct {
	logger log.Logger
	Service
}

func NewLoggingServer(logger log.Logger, s Service) Service {
	return &loggingServer{
		logger:  level.Info(logger),
		Service: s,
	}
}

func (s *loggingServer) List(ctx context.Context, pageSize, offset int) (images []types.Image, count int64, err error) {
	defer func(begin time.Time) {
		_ = s.logger.Log(
			"method", "List",
			"pageSize", pageSize,
			"offset", offset,
			"total", count,
			"took", time.Since(begin),
			"err", err,
		)
	}(time.Now())
	return s.Service.List(ctx, pageSize, offset)
}
