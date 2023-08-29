# 原始Banner获取模块调用文档

* 版本：v0.0.1-20230109

## 功能

加载 host:port 格式文件列表，尽可能准确获取对应主机的对应端口Banner信息

## 运行方式

```shell
./tentacleye -h
A application to get host:port banner info

Usage:
  Tentacleye [flags]

Flags:
  -h, --help            help for Tentacleye
      --maxThread int   max thread (default 256) 
      --portFilter      use port filter, use false to load all probe to per port (default true)
      --rarity int      rarity of probe, 1-9, use -1 to load all probe (default 6)
      --target string   target file, per line one host:ip (default "target.txt")
      
# 默认加载可执行文件同目录的target.txt作为目标
# 不输入参数时，等价执行./tentacleye --target=target.txt -rarity=6 --portFilter=true --maxThread=256
# 可在cmd/root.go对命令行参数进行修改（使用cobra模块）
# 可在modules/banner.go对扫描逻辑进行修改（入口函数为InitTask）
# 结果会在程序运行目录下的./tmp/pb/生成json格式文件
```

## 格式样例

```json
{
	"address":"1.34.148.22:1433",
    "host":"1.34.148.22",
    "ip":"1.34.148.22",
    "port":"1433",
    "is_tls":false,
    "probe":"mssql",
 "resp_hex":"0401002500000100000015000601001b000102001c000103001d0000ff0b00083400000000",
    "err_str":""
}

// address：目标地址
// host：目标主机
// IP：目标主机IP
// port：目标端口
// is_tls：是否为TLS端口
// probe：使用探针
// err_str：错误信息
```

