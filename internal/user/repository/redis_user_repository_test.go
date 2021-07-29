package repository_test

import (
	"fmt"
	"log"
	"ms-auth/internal/user/repository"
	"os"
	"testing"
	"time"

	"github.com/alicebob/miniredis"
	"github.com/elliotchance/redismock"
	"github.com/go-redis/redis"
	"github.com/stretchr/testify/assert"
)

var (
	client *redis.Client
	key    = "key"
	val    = "value"
)

func TestMain(m *testing.M) {
	mr, err := miniredis.Run()
	if err != nil {
		log.Fatalf("database failed %s", err)
	}

	client = redis.NewClient(&redis.Options{
		Addr: mr.Addr(),
	})

	code := m.Run()
	os.Exit(code)
}

func Test_Get(t *testing.T) {

	mock := redismock.NewNiceMock(client)
	mock.On("Get", key).Return(redis.NewStringResult(val, nil))

	r := repository.NewUserRedisRepository(mock)
	res, err := r.Get(key)
	assert.NoError(t, err)
	assert.Equal(t, val, res)

}

func TestSet(t *testing.T) {
	exp := time.Duration(0)

	mock := redismock.NewNiceMock(client)
	mock.On("Set", key, val, exp).Return(redis.NewStatusResult("", nil))

	r := repository.NewUserRedisRepository(mock)
	err := r.Set(key, val, exp)
	assert.NoError(t, err)
}

func Test_Del(t *testing.T) {

	mock := redismock.NewNiceMock(client)
	mock.On("Del", []string{key}).
		Return(redis.NewIntResult(1, nil))

	r := repository.NewUserRedisRepository(mock)
	err := r.Del(key)
	assert.NoError(t, err)

}

func Test_ZRemRangeByScore(t *testing.T) {

	max := fmt.Sprint(time.Now().UnixNano())

	mock := redismock.NewNiceMock(client)
	mock.On("ZRemRangeByScore", key, "0", max).
		Return(redis.NewIntResult(1, nil))

	r := repository.NewUserRedisRepository(mock)

	_, err := r.ZRemRangeByScore(key, "0", max)
	assert.NoError(t, err)

}

func Test_ZRange(t *testing.T) {

	start := int64(0)
	stop := int64(-1)

	mock := redismock.NewNiceMock(client)
	mock.On("ZRange", key, start, stop).
		Return(redis.NewStringSliceCmd(1, nil))

	r := repository.NewUserRedisRepository(mock)

	_, err := r.ZRange(key, start, stop)
	assert.NoError(t, err)

}

func Test_Expire(t *testing.T) {

	expiration := time.Duration(1000)

	mock := redismock.NewNiceMock(client)
	mock.On("Expire", key, expiration).
		Return(redis.NewBoolCmd("set", key, false))

	r := repository.NewUserRedisRepository(mock)

	res := r.Expire(key, expiration)

	assert.Equal(t, res, false)

}

func Test_ZAddNX(t *testing.T) {

	currentTime := time.Now().UnixNano()

	mock := redismock.NewNiceMock(client)
	mock.On("ZAddNX", key, []redis.Z{
		{
			Score: float64(currentTime), Member: float64(currentTime),
		},
	}).
		Return(redis.NewIntResult(1, nil))

	r := repository.NewUserRedisRepository(mock)
	res := r.ZAddNX(key, float64(currentTime), float64(currentTime))

	assert.Equal(t, res, int64(1))

}
