# go-util
## 介绍
- 2022/8/8
- 一个简单的go-util工具库,正在努力优化中
- 欢迎参与贡献
- 采用对象式调用工具类
- 目前已开发的工具类


```goland
var FileUtil = util.NewFileUtil()
var IpUtil = util.NewIpUtil()
var GinUitl = util.NewGinUtil()
var RandomUtil = util.NewRandomUtil()
var CryptoUtil = util.NewCryptoUtil()
var TimeUtil = util.NewTimeUtil()
var JsonUtil = util.NewJsonUtil()
var TaskUtil = util.NewTaskUtil()
var SortUtil = util.NewSortUtil()
var UrlUtil = util.NewUrlUtil()
var HttpUtil = util.NewHttpUtil()
```

## 安装教程
- Github安装
- 以下是配置格式
```text
go get github.com/licheng1013/go-util
```
### 测试
-  go test .\test\
### 第三方库
- [https://github.com/mitchellh/mapstructure](https://github.com/mitchellh/mapstructure)

## 工具类
- 各工具类的作用
### Json
- json转换对象
- 对象转json
- map转换结构体

### File(文件)
- 读取文件
- 写入文件
- 列出所有文件和目录(不包括子目录)
- 列出所有文件和目录(包括子目录)
- 列出所有文件(包括子目录)

### Token(令牌)
- 生成token
- 解析token
- 自定义token

### Gin(web框架)
- 保存gin上下文
- 获取带参数路径
- 获取不带参数路径

### Random(随机)
- 获取指定长度的随机数字字符串
- 获取固定范围的随机数字
- 获取指定长度的随机字符串(大小写26字母和1-9)

### Time(时间)
- 获取今日时间 -> 不带时分秒 
- 获取今日时间 -> 时分秒
- 获取今年开始时间
- 获取今年结束时间
- 获取当月开始时间
- 获取当月结束时间

### Crypto(密钥)
- md5加密
- RSA密钥生成
- RSA加密
- RSA解密


### Url(编码)
- 编码
- 解码

### Http(请求)
- get请求
- 文件下载

### Sort(排序)
- 字符串字典排序
- 数字从低到高排序
