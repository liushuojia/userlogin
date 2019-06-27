# userlogin

初次使用go研发, 如果有啥不足, 请大家见谅, 谢谢

主要用到go的资源
框架beego  github.com/astaxie/beego

数据库 github.com/gohouse/gorose  之前用gorose写的一个操作方法, 故未使用beego自带的

redis github.com/garyburd/redigo/redis



数据库文件 TUserDB.sql

用户登录后

产生一个UID  格式如 xxxx-xxxx-asdfsxxxxxxxadf   可以api获取, 或登录后获取, 其中第一位为自动id 第二位为linux时间戳 第三位为加密Md字符串

如果用户未登录 使用接口前需获取一个授权 token  xxxx-asdfsxxxxxxxadf  第一位为linux时间戳 第二位为加密Md字符串


发送短信/邮件 只是写进redis, 由独立的模块处理


nginx 做跨域转发配置

server {
        listen 80;
        server_name  user.xxx.com;

        root   html;
        index  index.html index.htm;

        ## send request back to apache ##
        location / {

          add_header Access-Control-Allow-Origin *;
          add_header Access-Control-Allow-Methods 'GET, POST, OPTIONS, PUT, DELETE';
          add_header Access-Control-Allow-Headers 'Token';
          if ($request_method = 'OPTIONS') {
            #跨域在发包前会先发一个options去测试是否连接可用
            return 204;
          }

                proxy_pass  http://127.0.0.1:8080;  #go 服务器可以指定到其他的机器上，这里设定本机服务器

                proxy_redirect     off;
                proxy_set_header   Host             $host;
                proxy_set_header   X-Real-IP        $remote_addr;
                proxy_set_header   X-Forwarded-For  $proxy_add_x_forwarded_for;
        }
}
