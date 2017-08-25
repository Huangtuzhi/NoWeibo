# NoWeibo
Known instant post of Weibo based on Golang

## Why this repo
最初的想法是想用 Go 实现关注监听某些人的博客，一旦有通知则拉微博数据。这里的异步监听功能使用 Go 的协程和通道功能是可以很优雅实现的。

## Useage
之前曾经使用爬虫去抓微博的数据，但它的鉴权机制有点复杂，不好实现。后来发现可以使用官方提供给开发者的 API，微博提供给开发者 [access_token](http://open.weibo.com/tools/console) 来标识自己的身份。

进入 [demo](./src/noweibo/examples/)，执行

`go run weibo_demo.go --access_token='your-token'`