package main 

import (
	"fmt"
	"noweibo"
	"flag"
	"bufio"
	"os"
	"time"
)

var (
	access_token = flag.String("access_token", "", "用户的访问令牌")
)

func CatchCommand(ch chan string) {
	reader := bufio.NewReader(os.Stdin)
	data, _, _ := reader.ReadLine()
	command := string(data)
	ch <- command
}

func GetWeibo(ch chan string) {
	var command string
	command = <-ch
	if (command == "start") {
		var statuses noweibo.Statuses
		flag.Parse()
		weibo := noweibo.Weibo{}
		weibo.CallNet("statuses/home_timeline", "GET", *access_token, &statuses)

		for _, status := range statuses.Statuses {
			fmt.Println("----------")
			fmt.Println(status.User.Screen_Name)
			fmt.Println(status.Text)
			fmt.Println("\n")
		}
	}
}

func main() {
	ch := make(chan string)
	go CatchCommand(ch)
	go GetWeibo(ch)
	time.Sleep(1e15)
}