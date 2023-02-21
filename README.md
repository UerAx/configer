# Detail
golang 多配置文件读取 支持ini,toml,yaml,json格式

# Usage

``` golang
// init

go get github.com/uerax/goconf

goconf.LoadConfig(path)

VarString("section", "key")
VarStringOrDefault("default", "section", "key")

```
