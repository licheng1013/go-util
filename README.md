# go-util
## 介绍
- 2022/8/8
- 一个简单的go-util工具库,简易封装一些api来使用。
- 目前命名规则并不完善，欢迎加入贡献！

## 版本变化
- 小版本迭代
- v0.0.1 > v0.0.2 废弃和优化
- 中版本迭代
- v0.0.0 > v0.1.0 增加新工具类
- 大版本迭代
- v0.0.0 > v1.0.0 删除废弃的api

## 未来如何升级
- 升级大版本前的最后一个版本然后修改所有使用过期方法的地方在升级到大版本

## 安装教程
- Github安装
- 无

- Gitee安装
- 以下是配置格式
```text
require (
	go-util v0.0.5
)

replace go-util => gitee.com/licheng1013/go-util v0.0.5
```

## 工具类
- 废弃的方法不会在表格内显示

### file_util
- 对文件进行的处理

### gin_util
- 对gin框架的封装
- 在一个请求上下文中注入gin的context即可使用：例如在拦截器时候注入

| 方法名             | 作用                |
|-----------------|-------------------|
| GinUrlPath      | 获取段路径 /xx/hh      |
| GinParamsUrlPah | 获取全路径 /xx/nn?xx=a |
| GinGetContext   | 获取上下文             |
| GinSetContext   | 设置上下文             |



