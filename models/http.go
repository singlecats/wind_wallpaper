package models

import (
    "crypto/tls"
    "fmt"
    "github.com/astaxie/beego/httplib"
    "net"
    "net/http"
    "reflect"
    "strings"
    "time"
)

type Search struct {
    Categories string `query:"categories"`
    Purity     string `query:"purity"`
    Color      int    `query:"color"`
    Page       int    `query:"page"`
    Q          string `query:"q"`
    Ratios     string `query:"ratios"`
    Sorting    string `query:"sorting"`
    Order      string `query:"order"`
    Seed       string `query:"seed"`
}

var (
    httpQueryConnectTimeOut   = 5
    httpQueryReadWriteTimeOut = 5
)

var tp http.RoundTripper = &http.Transport{
    DialContext: (&net.Dialer{
        Timeout:   30 * time.Second,
        KeepAlive: 30 * time.Second,
        DualStack: true,
    }).DialContext,
    MaxIdleConns:          100,
    IdleConnTimeout:       90 * time.Second,
    ExpectContinueTimeout: 1 * time.Second,
}

var host = "wallhaven.cc"
var baseUrl = "https://" + host
var searchUrl = baseUrl + "/search"
var FullImgUrl = "https://w.wallhaven.cc/full/%s/wallhaven-%s.%s"

func (s *Search) SearchQuery() string {
    params := make(map[string]interface{})
    rType := reflect.TypeOf(s)
    rVal := reflect.ValueOf(s)
    rType = rType.Elem()
    rVal = rVal.Elem()
    for i := 0; i < rVal.NumField(); i++ {
        t := rType.Field(i)
        f := rVal.Field(i)
        params[t.Tag.Get("query")] = f.Interface()
    }
    url := "?"
    if len(params) > 0 {
        for k, v := range params {
            url = url + k + "=" + fmt.Sprintf("%v", v) + "&"
        }
    }
    return searchUrl + strings.TrimRight(url, "&")
}

func Get(url string) *http.Response {
    req := httplib.Get(url)
    req = setHeader(req)
    //req.SetTransport(tp)
    req.SetTimeout(time.Duration(httpQueryConnectTimeOut)*time.Second, time.Duration(httpQueryReadWriteTimeOut)*time.Second)
    req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
    response, err := req.Response()
    if err != nil {
        panic(err)
    }
    return response
}

func setHeader(req *httplib.BeegoHTTPRequest) *httplib.BeegoHTTPRequest {
    req.Header("Host", host)
    req.Header("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
    req.Header("Content-Type", "text/html; charset=UTF-8")
    //req.Header("Accept-Encoding","gzip,deflate,br")
    req.Header("Accept-Language", "zh-CN,zh;q=0.9")
    req.Header("Accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.9")
    req.Header("Referer", "https://www/baidu.com")
    return req
}
