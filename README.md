# Detail
golang 多配置文件读取 支持ini,toml,yaml,json格式

# Usage

``` golang

go get github.com/uerax/goconf

goconf.LoadConfig(path)

// 直接使用
goconf.VarString("section", "key")
goconf.VarStringOrDefault("default", "section", "key")

// 转换成自己需要的结构体
cfg := goconf.NewConfigFile()
mar, _ := json.Marshal(cfg)
json.Unmarshal(mar, &YOUR_STRUCT)
```
