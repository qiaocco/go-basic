package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	//get()
	//getWithParams()
	post()
}

func get() {
	resp, err := http.Get("http://httpbin.org/get")
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}

func getWithParams() {
	apiUrl := "http://httpbin.org/get"
	data := url.Values{}
	data.Set("name", "qiaocc")
	data.Set("age", "29")
	u, err := url.ParseRequestURI(apiUrl)
	if err != nil {
		fmt.Printf("parse url url failed, err:%v\n", err)
	}
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("get failed, err:%v\n", err)
		return
	}
	fmt.Println(resp)
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("read from resp.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(body))
}

type AutoGenerated struct {
	Args struct {
	} `json:"args"`
	Data  string `json:"data"`
	Files struct {
	} `json:"files"`
	Form struct {
	} `json:"form"`
	Headers struct {
		AcceptEncoding string `json:"Accept-Encoding"`
		ContentLength  string `json:"Content-Length"`
		ContentType    string `json:"Content-Type"`
		Host           string `json:"Host"`
		UserAgent      string `json:"User-Agent"`
		XAmznTraceID   string `json:"X-Amzn-Trace-Id"`
	} `json:"headers"`
	JSON struct {
		Age  int    `json:"age"`
		Name string `json:"name"`
	} `json:"json"`
	Origin string `json:"origin"`
	URL    string `json:"url"`
}

func post() {
	apiUrl := "http://httpbin.org/post"
	contentType := "application/json"
	data := `{"name": "qiaocc", "age": 29}`
	resp, err := http.Post(apiUrl, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed, err: %v\n", err)
		return
	}

	var a AutoGenerated
	err = json.NewDecoder(resp.Body).Decode(&a)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(a.URL)

	//defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//if err != nil {
	//	fmt.Printf("read from resp.Body failed, err:%v\n", err)
	//	return
	//}
	//fmt.Println(string(body))

}