# 同源策略

[同源策略](https://developer.mozilla.org/zh-CN/docs/Web/Security/Same-origin_policy)是Web的一个重要安全策略。

传统网页模式下，客户端（主要是浏览器）和服务端均可基于具体需求部署该安全策略来保证安全，但默认情况下，主要是浏览器采取保守的安全策略。

但是WebSocket有所不同，使用WebSocket的普遍场景对于跨域的请求比较频繁，浏览器没有采取任何限制，而是将权利交给服务端，服务端根据请求头信息自行判断。