package message

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/guapo-organizations/go-micro-secret/cache"
	"strings"
)

//邮箱验证码校验
func CheckEmailCode(email string) (bool, error) {
	redis_client := cache.GetRedisClient()
	result, err := redis_client.Get(getEmailCacheKey(email)).Result()
	if err == redis.Nil {
		return false, fmt.Errorf("邮箱:%s没有发送邮件或者过期咯！", email)
	}

	//可能连接出错
	if err != nil {
		return false, err
	}

	//比较一下
	if !strings.EqualFold(result, email) {
		return false, nil
	}

	return true, nil
}
