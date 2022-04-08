# 为什么需要WebSocket ？

WebSocket的出现是为了满足人们对基于Web的全双工协议的需求。

Web是互联网上最大的应用平台，随着越来越多的人网上冲浪，基于web的即时聊天、游戏等项目出现。为了实现这些功能，开发者使用各种黑科技来在HTTP1.x的**单向的粗粒度的请求响应模型**上模拟一个全双工的通信协议，包括但不限于轮训、XHR、Commet、iframe...

一个更好的方式是提供一个新的全双工协议，从而更正式的、更有效率的支持以上应用需求。

这个协议在[RFC6455](https://datatracker.ietf.org/doc/html/rfc6455)中进行的定义。

# WebSocket 的设计思路

WebSocket的想法是使用现有的HTTP协议承载[握手协议](https://zh.wikipedia.org/wiki/WebSocket#%E6%8F%A1%E6%89%8B%E5%8D%8F%E8%AE%AE)，好处是可以利用HTTP基础设施，如Cookie等。HTTP Tunnel、HTTP2也采用了类似的设计。

握手完成后，进入WebSocket的数据传输阶段，其本质是一个**基于帧**的通信协议，具体参考`协议实例`一文。

当需要关闭连接时，任意一方可以发起关闭连接的请求，并半关闭它。

# WebSocket 的野心

握手阶段的 `Sec-WebSocket-Protocol` HTTP请求字段提供了让WebSocket来承载多个子协议的可能性，甚至在RFC中提出了多路复用的愿景。

WebSocket本质是复用承载HTTP协议的运输层提供一个简单的运输层协议。



# HTTP2时代的WebSocket

WebSocket 想要成为一个多路复用的协议终究只是很美好，因为后续出现了HTTP2,HTTP3。

HTTP2主要为了解决TCP单个连接上的HTTP请求存在队头阻塞的问题，在tcp之上构建了一个支持多路复用的运输层，而这个运输层和HTTP协议深度耦合。

HTTP2的RFC并没有提及WebSocket，但[RFC8441](https://tools.ietf.org/html/rfc8441)提供了将HTTP2流升级为WebSocket 甚至其它协议的可能性。

使用HTTP2的流来承载WebSocket除了节省TCP连接的一点点用外，并没有任何优点，反而因为中间帧层带来了性能损失（增加了至少9字节的帧头）。因此目前业界还是使用HTTP1.1来升级为WebSocket。



# HTTP3时代的WebSocket

// more
