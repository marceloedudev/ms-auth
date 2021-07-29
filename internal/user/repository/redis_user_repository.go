package repository

import (
	"ms-auth/internal/user/domain"
	"time"

	"github.com/go-redis/redis"
	"github.com/pkg/errors"
)

type UserRedisRepository struct {
	cacheDB redis.Cmdable
}

func NewUserRedisRepository(cacheDB redis.Cmdable) *UserRedisRepository {
	return &UserRedisRepository{
		cacheDB: cacheDB,
	}
}

func (r UserRedisRepository) Get(key string) (string, error) {
	result, err := r.cacheDB.Get(key).Result()

	if err != nil {
		return "", errors.Wrap(err, "db.Get")
	}

	return result, nil

}

func (r UserRedisRepository) Set(key string, value interface{}, expiration time.Duration) error {
	if err := r.cacheDB.Set(key, value, expiration).Err(); err != nil {
		return errors.Wrap(err, "db.Set")
	}

	return nil

}

func (r UserRedisRepository) Del(key string) error {
	result, err := r.cacheDB.Del(key).Result()

	if err != nil {
		return errors.Wrap(err, "db.Del")
	}

	if result == 0 {
		return domain.ErrUnknownError
	}

	return nil

}

func (r UserRedisRepository) ZRemRangeByScore(key, min, max string) (int64, error) {
	result, err := r.cacheDB.ZRemRangeByScore(key, min, max).Result()

	if err != nil {
		return 0, errors.Wrap(err, "db.ZRemRangeByScore")
	}

	return result, nil

}

func (r UserRedisRepository) ZRange(key string, start, stop int64) ([]string, error) {
	result, err := r.cacheDB.ZRange(key, start, stop).Result()

	if err != nil {
		return nil, errors.Wrap(err, "db.ZRange")
	}

	return result, nil

}

func (r UserRedisRepository) Expire(key string, expiration time.Duration) bool {

	result := r.cacheDB.Expire(key, expiration)

	return result.Val()

}

func (r UserRedisRepository) ZAddNX(key string, score float64, member interface{}) int64 {

	result := r.cacheDB.ZAddNX(key, redis.Z{Score: score, Member: member})

	return result.Val()

}
