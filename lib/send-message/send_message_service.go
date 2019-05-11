package send_message

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/guapo-organizations/account-service/lib/register"
	"github.com/guapo-organizations/go-micro-secret/cache"
	"github.com/guapo-organizations/go-micro-secret/help"
	"github.com/lifei6671/gorand"
	email_lib "github.com/nilslice/email"
	"strings"
	"sync"
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
func (this SendMessageService) SendEmailCode(email string) (result bool, err error) {

	if !help.VerifyEmailFormat(email) {
		result = false
		err = fmt.Errorf("邮箱格式不对")
		return
	}

	//一边异步存储，一边异步发送验证码
	var wg sync.WaitGroup
	wg.Add(2)

	//异步的将email存起来
	go func(email string) {
		defer wg.Done()
		register_service := register.RegisterService{}
		_, store_err := register_service.StoreEmail(email)
		if store_err != nil {
			result = false
			err = store_err
		}
	}(email)

	//调用三方发送验证码
	go func(email string) {
		defer wg.Done()
		redis_client := cache.GetRedisClient()
		//如果err没有报错，则是获取到了值
		_, send_err := redis_client.Get(this.getEmailCacheKey(email)).Result()

		if send_err == redis.Nil {
			//邮箱没有发送过
			code := gorand.KRand(6, gorand.KC_RAND_KIND_ALL)
			msg := email_lib.Message{
				To:      email,
				From:    FROM_EMAIL,
				Subject: "来自最靓的仔团队",
				Body:    fmt.Sprintf("欢迎，欢迎，这是您的验证码：%s;请收好", code),
			}
			send_err = msg.Send()
			if send_err != nil {
				result = false
				err = send_err
				return
			}
			_, err = redis_client.SetNX(this.getEmailCacheKey(email), code, 60*time.Second).Result()
			if send_err != nil {
				result = false
				err = send_err
				return
			}
			//发送成功
			result = true
			err = nil
			return
		}

		//邮箱已经发送
		result = false
		err = fmt.Errorf("我丢，你已经发送过验证码了，稍后再发好吧？！")
		if send_err != nil {
			//也可能是连接出错了
			err = send_err
		}

	}(email)

	wg.Wait()

	return
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
