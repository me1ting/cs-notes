# http提交数据的方式
## get

`get`请求可以携带少量的查询参数(query params)，该方式受到浏览器的最大URL长度限制，但对于用户而言很直观。

这种方式不需要设置`content-type`头字段。

```
:authority: github.com
:method: GET
:path: /shrimpbighead/showBossActiveTime

```

## post
### application/x-www-form-urlencoded

这是浏览器表单的默认提交方式，其内容与get请求的查询参数部分相同，只是由body来承载数据：

```
POST http://www.example.com HTTP/1.1
Content-Type: application/x-www-form-urlencoded;charset=utf-8

title=ahweg&sub%5B%5D=1&sub%5B%5D=2&sub%5B%5D=3
```

### multipart/form-data

使用表单提交时，需要显式设置其`[`enctype`](https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form#enctype)`属性值为`multipart/form-data`。

```
POST http://www.example.com HTTP/1.1
Content-Type:multipart/form-data; boundary=----WebKitFormBoundaryrGKCBY7qhFd3TrwA

------WebKitFormBoundaryrGKCBY7qhFd3TrwA
Content-Disposition: form-data; name="ahweg"

title
------WebKitFormBoundaryrGKCBY7qhFd3TrwA
Content-Disposition: form-data; name="file"; filename="ahweg.png"
Content-Type: image/png

PNG ... content of ahweg.png ...
------WebKitFormBoundaryrGKCBY7qhFd3TrwA--
```

使用一种特殊的结构来分割不同的字段，这也是浏览器上传文件的原生方式，对服务端而言，通常需要使用特定的处理插件。

### 自定义content-type

如果使用Ajax，可以灵活设置content-type。

常用的`content-type`包括：

- application/json，json格式
- text/plain，纯文本
- application/octet-stream，二进制数据

服务端需要使用`插件`来解析数据。

## 参考资料

[content-type 汇总](https://stackoverflow.com/questions/23714383/what-are-all-the-possible-values-for-http-content-type-header)