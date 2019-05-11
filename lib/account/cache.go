package account

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/guapo-organizations/go-micro-secret/cache"
	"time"
)

const (
	EMAIL_TOKEN_PREFIX = "email:token:"
	PHONE_TOKEN_PREFIX = "phone:token"
)

//email在缓存中的key
func GetEmailToken(email string) string {
	return EMAIL_TOKEN_PREFIX + email
}

//phone在缓存中的key
func GetPhoneToken(phone string) string {
	return PHONE_TOKEN_PREFIX + phone
}

//email存入缓存
func StoreEmail(email string) (bool, error) {
	redis_client := cache.GetRedisClient()

	_, err := redis_client.SetNX(GetEmailToken(email), email, 5*60*time.Second).Result()

	if err != nil {
		return false, err
	}

	return true, nil
}

//在缓存中获取email
func GetEmailFormCache(token string) (email string, err error) {
	redis_client := cache.GetRedisClient()
	email, err = redis_client.Get(token).Result()

	if err == redis.Nil {
		email = ""
		err = fmt.Errorf("token:%s无法找到email", token)
	}

	return email, err
}
