package main 

import (
	"fmt"
	"noweibo"
	"flag"
)

var (
	access_token = flag.String("access_token", "", "用户的访问令牌")
)

func main() {
	flag.Parse()
	weibo := noweibo.Weibo{}
	iRet := weibo.CallNet("statuses/user_timeline", "GET", *access_token)
	fmt.Println(iRet)
}