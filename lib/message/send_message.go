package message

import (
	"fmt"
	"github.com/go-mail/mail"
	"github.com/go-redis/redis"
	"github.com/guapo-organizations/account-service/lib/account"
	"github.com/guapo-organizations/go-micro-secret/cache"
	"github.com/guapo-organizations/go-micro-secret/help"
	"github.com/lifei6671/gorand"
	"log"
	"sync"
	"time"
)

const (
	FROM_EMAIL = "song8339702@163.com"
	PASSWD_EMAIL = "s123456"
	//缓存前缀
	EMAIL_CODE_CACHE_PREFIX = "account-service:email:code:"
)

//一分钟内不能重发验证码

//用于判断是否一分钟内的缓存
func getEmailCacheKey(email string) string {
	return EMAIL_CODE_CACHE_PREFIX + email
}

//发送邮箱验证码,并且存入缓存，注册的时候通过token从缓存中获取email
func SendEmailCode(email string) (result bool, err error) {

	if !help.VerifyEmailFormat(email) {
		result = false
		err = fmt.Errorf("邮箱格式不对")
		return
	}

	//一边异步存储，一边异步发送验证码
	var wg sync.WaitGroup
	wg.Add(2)

	//异步的将email存起来，注册的时候通过token去获取
	go func(email string) {
		defer wg.Done()
		_, store_err := account.StoreEmail(email)
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
		_, send_err := redis_client.Get(getEmailCacheKey(email)).Result()

		if send_err == redis.Nil {
			//邮箱没有发送过
			code := gorand.KRand(6, gorand.KC_RAND_KIND_ALL)
			d := mail.NewDialer("smtp.163.com", 465, FROM_EMAIL, PASSWD_EMAIL)
			m := mail.NewMessage()
			m.SetHeader("From", FROM_EMAIL)
			m.SetHeader("To", email)
			m.SetHeader("Subject", "来自最靓的仔团队")
			m.SetBody("text/html", fmt.Sprintf("欢迎，欢迎，这是您的验证码：%s;请收好", code))
			send_err = d.DialAndSend(m)
			log.Println("邮箱错误信息是什么:", send_err)
			if send_err != nil {
				result = false
				err = send_err
				return
			}
			_, err = redis_client.SetNX(getEmailCacheKey(email), code, 60*time.Second).Result()
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
