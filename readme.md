# 基于GO的代理池

基于Golang实现的代理IP池项目，主要功能为解析代理服务商所提供的代理验证入库，定时验证入库的代理保证代理的可用性，提供API访问实现快速搭建代理。同时也支持无限扩展代理服务商，同时支持IPV6换IPV4并入库。

- 文档：todo

## 运行项目

下载代码

**clone代码**

```
git clone https://github.com/guanjiangtao/go-project.git
```

**安装依赖并编译**

```
go build
```

**更新配置**

```go
// cd config/
var (
	Version   = 1.0                     // 版本号
	Name      = "proxy pool"            // 名称
	CronCheck = "@every 1m"             // 校验定时（目前设置为1m更新一次，也可以改成1h，1d或者1s）
	CronPull  = "@every 1m"             // 拉取定时
	Proxy     = map[string]interface{}{ // 填写解析方法 （此处填写需要处理的代理方法，可以无限扩展，目前为只提供了一个QinGuo的解析方法，如果还需要其他的，可以前往该文件实现）
		"QinGuo": common.QinGuo,
	}
	DBKey         = "proxy_pool"           // 代理池在Redis中存放的Key，可自定义修改
	DBHost        = "" // Redis的主机地址
	DBPassword    = "" // Redis的密码
	DB            = 0
	ProxyAuth     = "" // 代理权限，当然也可以设置白名单
	HttpsValidUrl = "https://www.qq.com"
	HttpValidUrl  = "http://httpbin.org"
)
```

**启动项目**

```
./main -type server // 服务端主要提供暴露给使用方调用的接口
./main -type schedule // 开启定时任务，主要负责拉取代理并进行校验
```

## 使用代理

启动web服务后, 默认配置下会开启 http://127.0.0.1:8080的api接口服务:

| api     | method | Description      | params                                      |
| ------- | ------ | ---------------- | ------------------------------------------- |
| /get    | GET    | 随机获取一个代理 | 可选参数: `?type=https` 过滤支持https的代理 |
| /getall | GET    | 获取所有代理     | 无                                          |
| /delete | GET    | 删除一个代理     | 必填参数: `?id=host` 传入为待删除的ip的host |
| /clear  | GET    | 清除所有代理     | 无                                          |

## 扩展代理

项目默认提供的青果云的解析能力，如果觉得无法满足要求可以通过所下方代理进行进一步扩展。

**编写实现解析方法**

```go
func Proxy1() []*ProxyGetter {
	var proxyResult = make([]*ProxyGetter, 0)
	var result map[string]interface{}
	resp, err := http.Get("") // 填写待爬取的URL
	if err != nil {
		log.Fatalln("request error！")
		return nil
	}
	if body, err := ioutil.ReadAll(resp.Body); err == nil {
		err = json.Unmarshal(body, &result)
	}
 
	if result["Code"].(float64) == 0 {
		  // 处理请求结果
	}
	return proxyResult
}
```

**配置文件添加对应的方法**

```go
Proxy = map[string]interface{}{ // 填写解析方法 （此处填写需要处理的代理方法，可以无限扩展，目前为只提供了一个QinGuo的解析方法，如果还需要其他的，可以前往该文件实现）
	"QinGuo": common.QinGuo,
  "Proxy1": common.Proxy1, // 添加解析方法
}
```

## 问题反馈

任何问题欢迎在[Issues](https://github.com/guanjiangtao/go-project/issues) 中反馈，你的反馈会让此项目变得更加完美。

## 贡献代码

本项目接受贡献代码，如果有比较好的方法，可以通过提MR的方式进行提交，作者会不定期进行Code Review。

### 