package fakes

import (
	"log"
	"ms-auth/internal/user/repository"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/go-redis/redis"
)

var (
	client *redis.Client
)

type UserRedisRepositoryFake struct {
	repo *repository.UserRedisRepository
}

func NewUserRedisRepositoryFake() *UserRedisRepositoryFake {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("error '%s' database connection", err)
	}

	client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	repo := repository.NewUserRedisRepository(client)

	return &UserRedisRepositoryFake{
		repo: repo,
	}
}

func (f *UserRedisRepositoryFake) Get(key string) (string, error) {

	result, err := f.repo.Get(key)

	if err != nil {
		return "", err
	}

	return result, nil

}

func (f *UserRedisRepositoryFake) Set(key string, value interface{}, expiration time.Duration) error {

	err := f.repo.Set(key, value, expiration)

	if err != nil {
		return err
	}

	return nil

}

func (f *UserRedisRepositoryFake) Del(key string) error {
	err := f.repo.Del(key)

	if err != nil {
		return err
	}

	return nil

}

func (f *UserRedisRepositoryFake) ZRemRangeByScore(key, min, max string) (int64, error) {

	result, err := f.repo.ZRemRangeByScore(key, min, max)

	if err != nil {
		return 0, err
	}

	return result, nil

}

func (f *UserRedisRepositoryFake) ZRange(key string, start, stop int64) ([]string, error) {

	result, err := f.repo.ZRange(key, start, stop)

	if err != nil {
		return nil, err
	}

	return result, nil

}

func (f *UserRedisRepositoryFake) Expire(key string, expiration time.Duration) bool {

	result := f.repo.Expire(key, expiration)

	return result

}

func (f *UserRedisRepositoryFake) ZAddNX(key string, score float64, member interface{}) int64 {

	result := f.repo.ZAddNX(key, score, member)

	return result

}
