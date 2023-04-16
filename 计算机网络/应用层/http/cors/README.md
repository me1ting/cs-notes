# 背景
在JavaScirpt被引入到浏览器后，[SOP](https://en.wikipedia.org/wiki/Same-origin_policy)就成为了web安全的基础。

SOP规定，在同一域(协议、域名、端口)中的资源对于外部是隔离的，避免通过JavaScript访问不属于其域的资源。

>除了JavaScript，哪些资源访问受到CORS影响还有待进一步了解

CORS是在SOP之下，提供一种安全的跨域资源访问功能。

总结，SOP是浏览器对于`JavaScript`的默认限制，CORS则是提供例外。

# 基本模型
当域A上的脚本想要访问域B上的资源时，浏览器会使用`http options`发起`预检`请求，服务器会检查并显示声明`允许`，浏览器根据服务器的允许，发出请求。

## 预检请求
http options请求，并带有请求来源以及可选的请求头：
```
Origin: http://www.example.com
```

## 服务器响应
服务器会返回多个头信息：
```
Access-Control-Allow-Origin: *
Access-Control-Allow-Methods: GET, HEAD, POST, OPTIONS
Access-Control-Allow-Headers": Content-Type
...
```
包括：

- 允许的请求来源
- 允许的访问方法
- 允许携带的头信息
- ...

# 例外
某些浏览器在某些场景可能会执行更严格的CORS策略，比如chrome禁止内容脚本进行CORS访问。