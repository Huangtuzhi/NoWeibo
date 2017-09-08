package noweibo

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"encoding/json"
)

type Params map[string]interface{}

type Weibo struct {
	httpClient http.Client 
}

func (weibo *Weibo) CallNet(apiMethod string, method string, accessToken string, response interface{}) error{
    url := fmt.Sprintf("%s/%s/%s%s?access_token=%s", ApiDomain, ApiVersion, apiMethod, ApiNamePostfix, accessToken)
    if method == "GET" {
		fmt.Println("GET METHOD")
		resp, err := weibo.httpClient.Get(url)
		if err != nil {
			fmt.Printf("inner error %s.\n", err)
			return err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("inner error %s.\n", err)
			return err
		}

		json.Unmarshal(body, &response)
		return nil
	} else if method == "POST" {
		fmt.Println("POST METHOD")
		resp, err := weibo.httpClient.Post(url, "application/x-www-form-urlencoded",
			strings.NewReader("name=titus"))
		if err != nil {
			return err
		}

		defer resp.Body.Close()
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Printf("inner error %s.\n", err)
			return err
		}
		json.Unmarshal(body, &response)
		return nil
	}

	return nil
}
