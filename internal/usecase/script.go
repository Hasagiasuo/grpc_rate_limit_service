package usecase

import (
	"context"
	"fmt"
	rlservicepb "rlservice/gen/v1"
	"rlservice/internal/domain/repositories"
	"rlservice/internal/infra/mredis"
	"rlservice/pkg/logger"
	"time"
)

// todo: redis methods
type ScriptRepository interface {
	Incr(ctx context.Context, key string, cost int64) (int64, error)
	ExpireOnce(ctx context.Context, key string, ttl time.Duration) error
	GetTTL(ctx context.Context, key string) (time.Duration, error)
}

type ScriptUsecase struct {
	log *logger.Logger
	rs  *mredis.RedisStorage
}

func NewScriptUsecase(log *logger.Logger, rs *mredis.RedisStorage) *ScriptUsecase {
	return &ScriptUsecase{
		log: log,
		rs:  rs,
	}
}

func (su *ScriptUsecase) ServeTimeRule(ctx context.Context, req *rlservicepb.CheckRequest) (*rlservicepb.CheckResponse, int, error) {
	timeRule, err := repositories.ParseTimeRule(su.log, req)
	if err != nil {
		return nil, -1, err
	}
	windowStart := time.Now().Unix() / int64(timeRule.Window.Seconds())
	rKey := fmt.Sprintf("rl:%s:%s:%d", timeRule.ServiceToken, timeRule.Key, windowStart)
	count, err := su.rs.Incr(ctx, rKey, timeRule.Cost)
	if err != nil {
		return nil, -1, err
	}
	if count == timeRule.Cost {
		if err := su.rs.ExpireOnce(ctx, rKey, timeRule.Window); err != nil {
			return nil, -1, err
		}
	}
	if count > timeRule.Limit {
		ttl, err := su.rs.GetTTL(ctx, rKey)
		if err != nil {
			return nil, -1, err
		}
		return &rlservicepb.CheckResponse{
			Allowed:    false,
			RetryAfter: int64(ttl.Seconds()),
		}, repositories.TIME_COST, nil
	}
	return &rlservicepb.CheckResponse{
		Allowed:    true,
		RetryAfter: 0,
	}, repositories.TIME_COST, nil
}
