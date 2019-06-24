package main

import (
    "fmt"
    "time"

    "github.com/garyburd/redigo/redis"
)

func main() {

    // 从配置文件获取redis配置并连接
    c, err := redis.Dial("tcp", "127.0.0.1:6379")
    if err != nil {
        fmt.Println("Connect to redis error", err)
        return
    }
    defer c.Close()

    //
    c.Do("SELECT", 0)

    _, err = c.Do("SET", "mykey", "superWang", "EX", "5")
    if err != nil {
        fmt.Println("redis set failed:", err)
    }

    username, err := redis.String(c.Do("GET", "mykey"))
    if err != nil {
        fmt.Println("redis get failed:", err)
    } else {
        fmt.Printf("Get mykey: %v \n", username)
    }

    time.Sleep( 8 * time.Second )

    username, err = redis.String(c.Do("GET", "mykey"))
    if err != nil {
        fmt.Println("redis get failed:", err)
    } else {
        fmt.Printf("Get mykey: %v \n", username)
    }
}
