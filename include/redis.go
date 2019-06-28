package conf

import (
	"fmt"
	//"encoding/json"
	"errors"

	"github.com/astaxie/beego"
	"github.com/garyburd/redigo/redis"

)

type RedisConn struct {
	Conn redis.Conn
}

func (c *RedisConn) Connect() ( err error ) {
	var connFlag bool = true
	if c.Conn==nil {
		connFlag = true
	}else{
		ping, _ := redis.String(c.Conn.Do("PING"));
		if( ping=="PONG" ){
			connFlag = false
		}
	}
	if connFlag {
		redis_host := beego.AppConfig.String("redis::host")
		redis_passwd := beego.AppConfig.String("redis::passwd")
		redis_db,_ := beego.AppConfig.Int("redis::db")
		c.Conn, err = redis.Dial("tcp", redis_host);

		if  err != nil {
			return
		}
		if _, err = c.Conn.Do("AUTH", redis_passwd); err != nil {
			c.Close()
			return
		}
		if redis_db>0 {
			if _, err = c.Conn.Do("SELECT", redis_db); err != nil {
				c.Close()
				return
			}			
		}
	}
	return
}

func (c *RedisConn) Close() {
	if( c.Conn!=nil ){
		c.Conn.Close()
		c.Conn = nil
	}
	return
	fmt.Println(c.Conn)
	return
}

func (c RedisConn) Set( rdKey string, rdVal string, rdExTime int64 ) (err error) {
	if err = c.Connect(); err != nil {
		return
	}

	c.Conn.Do("set", rdKey, rdVal)
	if rdExTime>0 {
		c.Conn.Do("Expire", rdKey, rdExTime)
	}
	return
}

func (c RedisConn) EXISTS( key string ) ( is_key_exit bool, err error ){
	if err = c.Connect(); err != nil {
		return
	}

	is_key_exit, err = redis.Bool( c.Conn.Do("EXISTS", key) )

	if err != nil {
		return
	}
	return
}


func (c RedisConn) Get( key string ) ( value string, err error ){

	if err = c.Connect(); err != nil {
		return
	}

	if_exit, err := c.EXISTS(key)
	if err != nil {
		return
	}
	if if_exit==false {
		err = errors.New("数据已经过期")
		return
	}

	value, err = redis.String(c.Conn.Do("GET", key))
	return
}

func (c RedisConn) Del( key string ) ( err error ){

	if err = c.Connect(); err != nil {
		return
	}

	c.Conn.Do("DEL", key)
	return
}




//队列管理插入队列
func (c RedisConn) RPush( key string, val string ) ( err error ){
	if err = c.Connect(); err != nil {
		return
	}

	_, err = c.Conn.Do("rPush", key, val )
	if err != nil {
		return
	}

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