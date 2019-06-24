package conf

import (
	"fmt"
	"encoding/json"
	"errors"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"

)


var RedisConn redis.Conn
func ConnectRedis() ( err error ) {
	var connFlag bool = true
	if RedisConn==nil {
		connFlag = true
	}else{
		ping, _ := redis.String(RedisConn.Do("PING"));
		if( (ping)=="PONG" ){
			connFlag = false
		}
	}
	if connFlag {
		redis_host := beego.AppConfig.String("redis::host")
		redis_passwd := beego.AppConfig.String("redis::passwd")
		redis_db,_ := beego.AppConfig.Int("redis::db")

		if RedisConn, err = redis.Dial("tcp", redis_host); err != nil {
			return
		}

		if _, err = RedisConn.Do("AUTH", redis_passwd); err != nil {
			RedisConn.Close()
			return
		}

		if _, err = RedisConn.Do("SELECT", redis_db); err != nil {
			RedisConn.Close()
			return
		}
	}
	return
}

func CloseRedis() {
	if( RedisConn!=nil ){
		RedisConn.Close()
		RedisConn = nil
	}
}

func RedisGet( key string ) ( value string, err error ){

	if err = ConnectRedis(); err != nil {
		return
	}
	defer CloseRedis()

	num, err := redis.Int( RedisConn.Do("EXISTS", key ) )

	if err != nil {
		return
	}
	if( num==0 ) {
		err = errors.New("验证码已经失效,请重新获取")
		return
	}

	value, err = redis.String(RedisConn.Do("GET", key))
	return
}


func GetSmsCode(mobile string) ( smscode string,err error ){

	checkKey := beego.AppConfig.String("sms::checkKey") + "_" + mobile

	smscode, err = RedisGet( checkKey )
	return
}


//  smsKey = "SmsQueue"             //短信发送队列
//    smsCheckKey = "SmsCheck"        //短信检验数据
func SendSms(mobile string, smscode string) ( err error) {
	
	if err = ConnectRedis(); err != nil {
		return
	}
	defer CloseRedis()

	queueKey := beego.AppConfig.String("sms::queueKey")
	checkKey := beego.AppConfig.String("sms::checkKey") + "_" + mobile

	num, err := redis.Int( RedisConn.Do("EXISTS", checkKey ) )
	if err != nil {
		return
	}
	if( num>0 ) {
		err = errors.New("短信已发送过了, 请60秒后重新获取")
		return
	}

	_, err = RedisConn.Do("set", checkKey, smscode , "EX", "60")
	if err != nil {
		return
	}

	tmpArray := make( map[string]string )
	tmpArray["mobile"] = mobile
	tmpArray["content"] = "您的验证码是：" + smscode + "。请不要把验证码泄露给其他人。"
	j, _ := json.Marshal(tmpArray)
	_, err = RedisConn.Do("rPush", queueKey, string(j))
	if err != nil {
		return
	}
	return
	fmt.Println(err)
	return
}

	/*

	_, err = c.Do("rPush", "mykey", smscode , "EX", "5")
	if err != nil {
		fmt.Println("redis set failed:", err)
	}

	username, err := redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}

	time.Sleep(8 * time.Second)

	username, err = redis.String(c.Do("GET", "mykey"))
	if err != nil {
		fmt.Println("redis get failed:", err)
	} else {
		fmt.Printf("Get mykey: %v \n", username)
	}
	*/