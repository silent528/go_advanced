package main

import (
	"context"
	"errors"
	"github.com/go-redis/redis/v8"
	"strconv"
	"time"
)

const (
	periodScript = `local limit = tonumber(ARGV[1])
local window = tonumber(ARGV[2])
local current = redis.call("INCRBY", KEYS[1], 1)
if current == 1 then
    redis.call("expire", KEYS[1], window)
    return 1
elseif current < limit then
    return 1
elseif current == limit then
    return 2
else
    return 0
end`
	zoneDiff = 3600 * 8
)

const (
	Unknown = iota

	Allowed

	HitQuota

	OverQuota

	internalOverQuota = 0
	internalAllowed   = 1
	internalHitQuota  = 2
)

type PeriodLimit struct {
	period     int
	quota      int
	limitStore *redis.Client
	align      bool
}

func NewPeriodLimit(period, quota int, limitStore *redis.Client) *PeriodLimit {
	limiter := &PeriodLimit{
		period:     period,
		quota:      quota,
		limitStore: limitStore,
	}
	return limiter
}

func (h *PeriodLimit) Take(key string) (int, error) {
	c := h.limitStore.Eval(context.Background(), periodScript, []string{key}, []string{
		strconv.Itoa(h.quota),
		strconv.Itoa(h.calcExpireSeconds()),
	})
	resp, err := c.Result()
	if err != nil {
		return Unknown, err
	}

	code, ok := resp.(int64)
	if !ok {
		return Unknown, errors.New("REDIS RESPONSE IS NOT INT")
	}

	switch code {
	case internalOverQuota:
		return OverQuota, nil
	case internalAllowed:
		return Allowed, nil
	case internalHitQuota:
		return HitQuota, nil
	default:
		return Unknown, errors.New("REDIS RETURN STATE ERROR")
	}
}

func (h *PeriodLimit) calcExpireSeconds() int {
	if h.align {
		unix := time.Now().Unix() + zoneDiff
		return h.period - int(unix%int64(h.period))
	}
	return h.period
}
