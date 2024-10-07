package httpx

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

func NewHttpProxy(proxyUrl string, handel ...func(req *http.Request)) (*httputil.ReverseProxy, error) {
	targetURL, err := url.Parse(proxyUrl)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	// 创建代理处理器
	proxy := &httputil.ReverseProxy{
		Director: func(req *http.Request) {
			//todo 这里可以负载均衡，添加元数据，鉴权用户是否登录，验证权限等等
			if handel != nil {
				f := handel[0]
				f(req)
			}

			//query := req.URL.Query()
			//query.Add("_version", "v1")
			//query.Add("_type", "windows")
			//req.URL.RawQuery = query.Encode()
			req.URL.Scheme = targetURL.Scheme
			req.URL.Host = targetURL.Host

			req.Header.Set("X-Forwarded-Host", req.Header.Get("Host"))
			req.Header.Set("X-Origin-Host", targetURL.Host)
		},
	}
	return proxy, nil
}
