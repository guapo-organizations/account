package send_message

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/guapo-organizations/go-micro-secret/cache"
	"github.com/guapo-organizations/go-micro-secret/help"
	"github.com/lifei6671/gorand"
	email_lib "github.com/nilslice/email"
	"strings"
	"time"
)

const (
	FROM_EMAIL = "13647897503@163.com"
	//缓存前缀
	EMAIL_CODE_CACHE_PREFIX = "email:code:"
)

type SendMessageService struct {
}

//邮箱缓存key
func (this SendMessageService) getEmailCacheKey(email string) string {
	return EMAIL_CODE_CACHE_PREFIX + email
}

//邮箱发送验证码
func (this SendMessageService) SendEmailCode(email string) (bool, error) {

	if !help.VerifyEmailFormat(email) {
		return false, fmt.Errorf("邮箱格式不对")
	}

	redis_client := cache.GetRedisClient()

	_, err := redis_client.Get(this.getEmailCacheKey(email)).Result()
	if err == redis.Nil {
		//不存在key
		//发送验证码
		code := gorand.KRand(6, gorand.KC_RAND_KIND_ALL)

		msg := email_lib.Message{
			To:      email,
			From:    FROM_EMAIL,
			Subject: "来自最靓的仔团队",
			Body:    fmt.Sprintf("欢迎，欢迎，这是您的验证码：%s;请收好", code),
		}

		err := msg.Send()
		if err != nil {
			return false, err
		}

		_, err = redis_client.SetNX(this.getEmailCacheKey(email), code, 60*time.Second).Result()

		if err != nil {
			return false, err
		}

		return true, nil
	}

	//不知道什么错
	if err != nil {
		return false, err
	}

	//有值
	return false, fmt.Errorf("我丢，你已经发送过验证码了，稍后再发好吧？！")
}

//校验emailCode
func (this SendMessageService) CheckEmailCode(email string) (bool, error) {
	redis_client := cache.GetRedisClient()
	result, err := redis_client.Get(this.getEmailCacheKey(email)).Result()
	if err == redis.Nil {
		return false, fmt.Errorf("邮箱:%s没有发送邮件或者过期咯！", email)
	}
	//不知道什么错
	if err != nil {
		return false, err
	}

	//比较一下
	if !strings.EqualFold(result, email) {
		return false, nil
	}

	return true, nil
}
