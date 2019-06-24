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


