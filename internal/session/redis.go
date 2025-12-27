package session
import (
	"context"
	"time"
	"github.com/redis/go-redis/v9"
)
var ctx = context.Background()

func CreateClient() (rdb *redis.Client){
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379", // ADDR
		Password: "password", // Password
		DB:       0,
	})
	return rdb
}
func SetSession(rdb *redis.Client, session_id string, expire time.Duration, user_id int, role string) error {
	err := rdb.HSet(ctx, session_id, map[string]interface{}{
		"user_id": user_id,
		"role":    role,
	}).Err()
	if err != nil {
		return err
	}
	rdb.Expire(ctx, session_id, time.Duration(expire))
	return nil
}
func GetSession(rdb *redis.Client, session_id string) (map[string]string, error){
	val, err := rdb.HGetAll(ctx, session_id).Result()
	if err != nil {
		return nil, err
	}
	return val, nil
}
func UpdateTTL(rdb *redis.Client, session_id string, expire time.Duration) error{
	return rdb.Expire(ctx, session_id, expire).Err()
}
func DeleteSession(rdb *redis.Client, session_id string) (int64, error){
	return rdb.Del(ctx, session_id).Result()
}