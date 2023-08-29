# tlsx

##项目链接
* https://github.com/projectdiscovery/tlsx

## 用法

```shell
dnsx -h
Usage:
  ./tlsx [flags]

Flags:
INPUT:
   -u, -host string[]  target host to scan (-u INPUT1,INPUT2)
   -l, -list string    target list to scan (-l INPUT_FILE)
   -p, -port string[]  target port to connect (default 443)

SCAN-MODE:
   -sm, -scan-mode string     tls connection mode to use (ctls, ztls, auto) (default "auto")
   -ps, -pre-handshake        enable pre-handshake tls connection (early termination) using ztls
   -sa, -scan-all-ips         scan all ips for a host (default false)
   -iv, -ip-version string[]  ip version to use (4, 6) (default 4)

PROBES:
   -san                 display subject alternative names
   -cn                  display subject common names
   -so                  display subject organization name
   -tv, -tls-version    display used tls version
   -cipher              display used cipher
   -hash string         display certificate fingerprint hashes (md5,sha1,sha256)
   -jarm                display jarm fingerprint hash
   -ja3                 display ja3 fingerprint hash (using ztls)
   -wc, -wildcard-cert  display host with wildcard ssl certificate
   -tps, -probe-status  display tls probe status
   -ve, -version-enum   enumerate and display supported tls versions
   -ce, -cipher-enum    enumerate and display supported cipher
   -ch, -client-hello   include client hello in json output (ztls mode only)
   -sh, -server-hello   include server hello in json output (ztls mode only)

MISCONFIGURATIONS:
   -ex, -expired      display host with host expired certificate
   -ss, -self-signed  display host with self-signed certificate
   -mm, -mismatched   display host with mismatched certificate
   -re, -revoked      display host with revoked certificate

CONFIGURATIONS:
   -config string               path to the tlsx configuration file
   -r, -resolvers string[]      list of resolvers to use
   -cc, -cacert string          client certificate authority file
   -ci, -cipher-input string[]  ciphers to use with tls connection
   -sni string[]                tls sni hostname to use
   -rs, -random-sni             use random sni when empty
   -min-version string          minimum tls version to accept (ssl30,tls10,tls11,tls12,tls13)
   -max-version string          maximum tls version to accept (ssl30,tls10,tls11,tls12,tls13)
   -ac, -all-ciphers            send all ciphers as accepted inputs (default true)
   -cert, -certificate          include certificates in json output (PEM format)
   -tc, -tls-chain              include certificates chain in json output
   -vc, -verify-cert            enable verification of server certificate

OPTIMIZATIONS:
   -c, -concurrency int  number of concurrent threads to process (default 300)
   -timeout int          tls connection timeout in seconds (default 5)
   -retry int            number of retries to perform for failures (default 3)
   -delay string         duration to wait between each connection per thread (eg: 200ms, 1s)

OUTPUT:
   -o, -output string  file to write output to
   -j, -json           display json format output
   -ro, -resp-only     display tls response only
   -silent             display silent output
   -nc, -no-color      disable colors in cli output
   -v, -verbose        display verbose output
   -version            display project version
```

## 运行 tlsx

### tlsx 输入

* tlsx需要TLS连接并接受多种格式：
```bash
AS1449 # ASN input
173.0.84.0/24 # CIDR input
93.184.216.34 # IP input
example.com # DNS input
example.com:443 # DNS input with port
https://example.com:443 # URL input port
```

* 逗号分隔主机：
```console
$ tlsx -u 93.184.216.34,example.com,example.com:443,https://example.com:443 -silent
```

* 基于文件的主机输入：

```console
$ tlsx -list host_list.txt
```

**端口输入:**

tlsx默认连接到端口443，可以使用-port / -p标志自定义
* 逗号分隔的: 

```
$ tlsx -u hackerone.com -p 443,8443
```

* 基于文件: 

```
$ tlsx -u hackerone.com -p port_list.txt
```


### TLS 探测 （默认）

```console
$ echo 173.0.84.0/24 | tlsx 
  

  _____ _    _____  __
 |_   _| |  / __\ \/ /
   | | | |__\__ \>  < 
   |_| |____|___/_/\_\  v0.0.1

    projectdiscovery.io

[WRN] Use with caution. You are responsible for your actions.
[WRN] Developers assume no liability and are not responsible for any misuse or damage.

173.0.84.69:443
173.0.84.67:443
173.0.84.68:443
173.0.84.66:443
173.0.84.76:443
173.0.84.70:443
173.0.84.72:443
```

### SAN/CN 探针

TLS certificate contains DNS names under **subject alternative name** and **common name** field that can be extracted using `-san`, `-cn` flag.

```console
$ echo 173.0.84.0/24 | tlsx -san -cn -silent

173.0.84.104:443 [uptycspay.paypal.com]
173.0.84.104:443 [api-3t.paypal.com]
173.0.84.104:443 [api-m.paypal.com]
173.0.84.104:443 [payflowpro.paypal.com]
173.0.84.104:443 [pointofsale-s.paypal.com]
173.0.84.104:443 [svcs.paypal.com]
173.0.84.104:443 [uptycsven.paypal.com]
173.0.84.104:443 [api-aa.paypal.com]
173.0.84.104:443 [pilot-payflowpro.paypal.com]
173.0.84.104:443 [pointofsale.paypal.com]

```

* 可选-resp-only标志可用于在 CLI 输出中仅列出 dns 名称
```console
$ echo 173.0.84.0/24 | tlsx -san -cn -silent -resp-only

api-aa-3t.paypal.com
pilot-payflowpro.paypal.com
pointofsale-s.paypal.com
uptycshon.paypal.com
a.paypal.com
adjvendor.paypal.com
zootapi.paypal.com
api-aa.paypal.com
payflowpro.paypal.com
pointofsale.paypal.com
uptycspay.paypal.com

```


### TLS / 密码探测

```console
$ subfinder -d hackerone.com | tlsx -tls-version -cipher

mta-sts.hackerone.com:443 [TLS1.3] [TLS_AES_128_GCM_SHA256]
hackerone.com:443 [TLS1.3] [TLS_AES_128_GCM_SHA256]
api.hackerone.com:443 [TLS1.3] [TLS_AES_128_GCM_SHA256]
mta-sts.managed.hackerone.com:443 [TLS1.3] [TLS_AES_128_GCM_SHA256]
mta-sts.forwarding.hackerone.com:443 [TLS1.3] [TLS_AES_128_GCM_SHA256]
www.hackerone.com:443 [TLS1.3] [TLS_AES_128_GCM_SHA256]
support.hackerone.com:443 [TLS1.2] [TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256]
```

# TLS配置错误

可以向 tlsx 提供主机列表以检测过期/自签名/不匹配/吊销的证书。
```console
$ tlsx -l hosts.txt -expired -self-signed -mismatched -revoked
  

  _____ _    _____  __
 |_   _| |  / __\ \/ /
   | | | |__\__ \>  < 
   |_| |____|___/_/\_\  v0.0.1

    projectdiscovery.io

[WRN] Use with caution. You are responsible for your actions.
[WRN] Developers assume no liability and are not responsible for any misuse or damage.

wrong.host.badssl.com:443 [mismatched]
self-signed.badssl.com:443 [self-signed]
expired.badssl.com:443 [expired]
revoked.badssl.com:443 [revoked]
```

### JARM TLS 指纹
```console
$ echo hackerone.com | tlsx -jarm -silent

hackerone.com:443 [29d3dd00029d29d00042d43d00041d5de67cc9954cc85372523050f20b5007]
```

### JA3 TLS 指纹
```console
$ echo hackerone.com | tlsx -ja3 -silent

hackerone.com:443 [20c9baf81bfe96ff89722899e75d0190]
```

### JSON输出

```console
echo example.com | tlsx -json -silent | jq .
```

```json
{
  "timestamp": "2022-08-22T21:22:59.799053+05:30",
  "host": "example.com",
  "ip": "93.184.216.34",
  "port": "443",
  "probe_status": true,
  "tls_version": "tls13",
  "cipher": "TLS_AES_256_GCM_SHA384",
  "not_before": "2022-03-14T00:00:00Z",
  "not_after": "2023-03-14T23:59:59Z",
  "subject_dn": "CN=www.example.org, O=Internet Corporation for Assigned Names and Numbers, L=Los Angeles, ST=California, C=US",
  "subject_cn": "www.example.org",
  "subject_org": [
    "Internet Corporation for Assigned Names and Numbers"
  ],
  "subject_an": [
    "www.example.org",
    "example.net",
    "example.edu",
    "example.com",
    "example.org",
    "www.example.com",
    "www.example.edu",
    "www.example.net"
  ],
  "issuer_dn": "CN=DigiCert TLS RSA SHA256 2020 CA1, O=DigiCert Inc, C=US",
  "issuer_cn": "DigiCert TLS RSA SHA256 2020 CA1",
  "issuer_org": [
    "DigiCert Inc"
  ],
  "fingerprint_hash": {
    "md5": "c5208a47259d540a6e3404dddb85af91",
    "sha1": "df81dfa6b61eafdffffe1a250240db5d2e6cee25",
    "sha256": "7f2fe8d6b18e9a47839256cd97938daa70e8515750298ddba2f3f4b8440113fc"
  },
  "tls_connection": "ctls",
  "sni": "example.com"
}
```

## 配置

### 扫描方式

tlsx 提供了多种模式来建立 TLS 连接——

- `auto`（**失败时自动回退到其他模式**）-默认
- `ctls` (**crypto/tls**)
- `ztls` (**zcrypto/tls**)
- `openssl` (**openssl**)


