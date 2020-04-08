package xhttp

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type RequestOption struct {
	ContentType string
	Method      string
	Headers     map[string]string
}

//生成最新请求Option
func makeOptions(options ...RequestOption) *RequestOption {
	opts := &RequestOption{
		ContentType: "application/json",
		Method:      "POST",
	}
	if options != nil {
		option := options[0]
		if option.ContentType != "" {
			opts.ContentType = option.ContentType
		}
		if option.Method != "" {
			opts.Method = option.Method
		}
	}
	return opts
}

//发送POST请求
//url:请求地址，data:POST请求提交的数据,contentType:请求体格式，如：application/json
//content:请求放回的内容
func HttpRequest(url string, data interface{}, options ...RequestOption) ([]byte, error) {
	var opts *RequestOption
	if options != nil {
		opts = makeOptions(options[0])
	} else {
		opts = makeOptions()
	}
	jsonStr, _ := json.Marshal(data)
	req, err := http.NewRequest(opts.Method, url, bytes.NewBuffer(jsonStr))
	req.Header.Add("content-type", opts.ContentType)
	if opts.Headers != nil {
		for k, v := range opts.Headers {
			if v != "content-type" {
				req.Header.Add(k, v)
			}
		}
	}

	if err != nil {
		return nil, err
	}
	defer req.Body.Close()

	client := &http.Client{Timeout: 5 * time.Second}
	resp, error := client.Do(req)
	if error != nil {
		return nil, err
	}
	defer resp.Body.Close()

	result, _ := ioutil.ReadAll(resp.Body)
	return result, nil
}

//代理请求（相当于请求转发）
func HttpProxy(url string, w http.ResponseWriter, r *http.Request) ([]byte, error) {
	client := &http.Client{}

	//提交请求
	request, err := http.NewRequest(r.Method, url, r.Body)
	request.Header = r.Header

	if err != nil {
		return nil, err
	}
	//处理返回结果
	response, _ := client.Do(request)
	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	return body, nil
}
