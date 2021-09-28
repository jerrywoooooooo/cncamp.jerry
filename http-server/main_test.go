package main_test

import (
	"github.com/stretchr/testify/assert"
	"io/ioutil"
	"net/http"
	"testing"
)

const host = "http://127.0.0.1"
const indexUrl = host + "/index"
const healthUrl = host + "/healthz"
const customUrl = host + "/custom"

// 1. 接收客户端request，并将request中带的header写入response header
func TestHeader(t *testing.T) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", indexUrl, nil)
	//增加header选项
	const testHeader = "MyHeaderTest"
	request.Header.Add("MyHeader", testHeader)
	//处理返回结果
	response, _ := client.Do(request)
	defer response.Body.Close()
	t.Log(response.Header.Get("MyHeader"))
	assert.Equal(t, response.Header.Get("MyHeader"), testHeader)
}

// 2. 读取当前系统的环境变量中的VERSION配置，并写入response header
func TestVersion(t *testing.T) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", indexUrl, nil)
	//处理返回结果
	response, _ := client.Do(request)
	defer response.Body.Close()
	t.Log(response.Header.Get("VERSION"))
	assert.True(t, response.Header.Get("VERSION") != "")
}

// 4. 当访问localhost/healthz时, 应返回200
func TestHealth(t *testing.T) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", healthUrl, nil)
	//处理返回结果
	response, _ := client.Do(request)
	defer response.Body.Close()
	result, _ := ioutil.ReadAll(response.Body)
	t.Log(string(result))
	assert.True(t, string(result) == "200")
}

// 访问不存在的url
func Test404(t *testing.T) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", host+"/path-not-exist", nil)
	//处理返回结果
	response, _ := client.Do(request)
	defer response.Body.Close()
	assert.True(t, response.StatusCode == 404)
}

// 访问custom
func TestCustom(t *testing.T) {
	client := &http.Client{}
	request, _ := http.NewRequest("GET", customUrl, nil)
	//处理返回结果
	response, _ := client.Do(request)
	defer response.Body.Close()
	result, _ := ioutil.ReadAll(response.Body)
	t.Log(string(result))
	assert.True(t, response.StatusCode == 200 && string(result) != "")
}
