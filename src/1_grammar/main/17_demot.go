package main

import (
	"fmt"
	"github.com/go-resty/resty/v2"
	"net/http"
)

func main() {
	//resp, err := http.Get("http://10.32.210.101:30014")
	//if err == nil && resp.StatusCode==http.StatusOK {
	//	fmt.Println("success")
	//} else {
	//	fmt.Println(resp.StatusCode)
	//}

	client := resty.New()
	res, err := client.R().Get("http://10.32.210.101:31032")
	if err == nil && res.StatusCode() == http.StatusOK {
		fmt.Println("success")
	} else {
		fmt.Println(res.StatusCode())
	}
}
