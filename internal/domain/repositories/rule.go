package repositories

import (
	"fmt"
	rlservicepb "rlservice/gen/v1"
	"rlservice/internal/domain/models"
	"rlservice/pkg/logger"
	"strconv"
	"strings"
	"time"
)

const (
	SECOND = 's'
	MINUTE = 'm'
	HOUR   = 'h'
	DAY    = 'd'
)

const (
	TIME_COST = 1
)

// 5m/10 -> limit = window = 300s
func ParseTimeRule(log *logger.Logger, req *rlservicepb.CheckRequest) (*models.TimeRule, error) {
	const op = "repositories.ParseTimeRule"
	details := strings.Split(req.Rule, "/")
	if len(details) != 2 {
		log.Error(op, "uncorrect rule format")
		return nil, ErrUncorrectRuleFormat
	}
	if len(details[0]) < 2 {
		log.Error(op, "uncorrect time format")
		return nil, ErrUncorrectRuleFormat
	}
	number, err := strconv.Atoi(details[0][0:len(details[0])-1])
	if err != nil {
		log.Error(op, fmt.Sprintf("uncorrect format number for time type: %v", err))
		return nil, ErrUncorrectRuleFormat
	}
	var win time.Duration
	switch details[0][len(details[0])-1] {
	case SECOND:
		win = time.Duration(number) * time.Second
	case MINUTE:
		win = time.Duration(number) * time.Minute
	case HOUR:
		win = time.Duration(number) * time.Hour
	case DAY:
		win = time.Duration(number) * time.Hour * 24
	default:
		log.Error(op, "undefined time type")
		return nil, ErrUndefinedTimeType
	}
	lim, err := strconv.Atoi(details[1])
	if err != nil {
		log.Error(op, fmt.Sprintf("cannot parse limit: %v", err))
		return nil, ErrUncorrectRuleFormat
	}
	return &models.TimeRule{
		Key:          req.GetKey(),
		ServiceToken: req.GetServiceToken(),
		Limit:        int64(lim),
		Window:       win,
		Cost:         TIME_COST,
	}, nil
}
