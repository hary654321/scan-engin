# go-portScan 主机存活、开放端口探测、原始banner获取
- 版本：v0.0.1-20230109

## 功能
1. 加载用户输入的ipv4地址信息/cidr网络段，设置判断主机存活等参数，
2. 加载用户输入的ipv4地址信息/cidr网络段，设置需要扫描的端口和并发数量等参数，获得目标地址范围的端口开放情况。
3. 加载用户输入的ipv4地址信息/cidr网络段，设置需要扫描的端口和并发数量等参数，获得目标地址范围的端口原始banners信息

## 运行方式
```bash
root@xxxx go-port-scan # ./go-portScan 
NAME:
   PortScan - A new cli application

USAGE:
   PortScan [global options] command [command options] [arguments...]

DESCRIPTION:
   High-performance port scanner

COMMANDS:
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --ip value                   target ip, eg: "1.1.1.1/30,1.1.1.1-1.1.1.2,1.1.1.1-2"
   --iL value                   target ip file, eg: "ips.txt"
   --port value, -p value       eg: "top1000,5612,65120,-" (default: "top1000")
   --Pn                         no ping probe (default: false)
   --rateP value, --rp value    concurrent num when ping probe each ip (default: 300)
   --sT                         TCP-mode(support IPv4 and IPv6) (default: false)
   --timeout value, --to value  TCP-mode SYN-mode timeout. unit is ms. (default: 800)
   --sS                         Use SYN-mode(Only IPv4) (default: true)
   --sQ                         Use tcp-mode(tcp result collection) (default: false)
   --dev value                  specified pcap dev name
   --rate value, -r value       number of packets sent per second. If set -1, TCP-mode is 1000, SYN-mode is 2000(SYN-mode is restricted by the network adapter, 2000=1M) (default: -1)
   --devices, --ld              list devices name (default: false)
   --sV                         port service identify (default: false)
   --httpx                      http server identify (default: false)
   --netLive                    Detect live C-class networks, eg: -ip 192.168.0.0/16,172.16.0.0/12,10.0.0.0/8 (default: false)
   --resume value               go-portScan starts in recovery mode
   --ipLive                     Detect ip is alive (default: false)
   --es                         Result send to es (default: false)
   --json                       go-portScan use json output (default: false)
   --es_url value               elasticsearch server url
   --es_db value                elasticsearch server db name
   --es_user value              elasticsearch server username
   --es_passwd value            elasticsearch server password
   --help, -h                   show help (default: false)
```

### 主机存活方式
#### 输入
```bash
root@xxx go-port-scan # ./go-portScan -ip 110.42.1.2/16 -ipLive true 
```
#### 输出
```bash
2023/01/10 18:22:28 [+] 110.42.1.118:22 is open
2023/01/10 18:22:28 [+] 110.42.1.207:22 is open
2023/01/10 18:22:28 [+] 110.42.2.213:22 is open
2023/01/10 18:22:28 [+] 110.42.2.82:49 is open
2023/01/10 18:22:28 [+] 110.42.3.134:80 is open
2023/01/10 18:22:28 [+] 110.42.1.141:80 is open
2023/01/10 18:22:28 [+] 110.42.1.203:80 is open
2023/01/10 18:22:29 [+] 110.42.2.138:80 is open
2023/01/10 18:22:29 [+] 110.42.1.219:80 is open
2023/01/10 18:22:29 [+] 110.42.2.213:80 is open
```

### 开放端口探测
#### 输入
```bash
root@xxx go-port-scan # ./go-portScan -ip 110.42.1.2/16 -port top1000 
```
#### 输出
```bash
2023/01/10 18:15:01 [+] 110.42.2.82:42 is open
2023/01/10 18:15:01 [+] 110.42.3.189:70 is open
2023/01/10 18:15:01 [+] 110.42.3.153:43 is open
2023/01/10 18:15:01 [+] 110.42.2.78:43 is open
2023/01/10 18:15:01 [+] 110.42.2.213:49 is open
2023/01/10 18:15:01 [+] 110.42.2.78:49 is open
2023/01/10 18:15:01 [+] 110.42.3.153:49 is open
2023/01/10 18:15:01 [+] 110.42.3.189:79 is open
2023/01/10 18:15:01 [+] 110.42.2.82:43 is open
2023/01/10 18:15:01 [+] 110.42.2.82:49 is open
2023/01/10 18:15:01 [+] 110.42.3.189:80 is open
```

### 原始banner获取
#### 输入
```bash
root@xxx go-port-scan # ./go-portScan -ip 110.42.1.2/16 -port top1000 -sQ true

# 如果需要将结果推送到es数据库，则需要配置es数据库相关信息
```

#### 输出
```bash
2023/01/11 10:56:26 [+] {"ip":"110.42.1.141","port":135,"result":{"tls":"{\"domainName\":\"110.42.1.141\",\"ip\":\"\",\"issuer\":\"\",\"commonName\":\"\",\"sans\":null,\"notBefore\":\"\",\"notAfter\":\"\",\"error\":\"context deadline exceeded\",\"certs\":null,\"port_cont\":\"\"}"},"time":1673405784}
2023/01/11 10:56:26 [+] {"ip":"110.42.1.207","port":749,"result":{"tcp":"485454502f312e312033303220466f756e640d0a5365727665723a206e67696e780d0a446174653a205765642c203131204a616e20323032332030323a35353a313920474d540d0a436f6e74656e742d547970653a20746578742f68746d6c3b20636861727365743d7574662d380d0a5472616e736665722d456e636f64696e673a206368756e6b65640d0a436f6e6e656374696f6e3a20636c6f73650d0a43616368652d636f6e74726f6c3a206e6f2d63616368652c6d7573742d726576616c69646174650d0a4c6f636174696f6e3a202f696e6465782f757365722f6c6f67696e0d0a0d0a300d0a0d0a","tls":"{\"domainName\":\"110.42.1.207\",\"ip\":\"\",\"issuer\":\"\",\"commonName\":\"\",\"sans\":null,\"notBefore\":\"\",\"notAfter\":\"\",\"error\":\"tls: first record does not look like a TLS handshake\",\"certs\":null,\"port_cont\":\"\"}"},"time":1673405786}
2023/01/11 10:56:27 [+] {"ip":"110.42.2.225","port":3306,"result":{"tcp":"46000000ff6a04486f7374202733362e3131302e3131362e333827206973206e6f7420616c6c6f77656420746f20636f6e6e65637420746f2074686973204d7953514c20736572766572","tls":"{\"domainName\":\"110.42.2.225\",\"ip\":\"\",\"issuer\":\"\",\"commonName\":\"\",\"sans\":null,\"notBefore\":\"\",\"notAfter\":\"\",\"error\":\"tls: first record does not look like a TLS handshake\",\"certs\":null,\"port_cont\":\"\"}"},"time":1673405786}
2023/01/11 10:56:27 [+] {"ip":"110.42.1.197","port":80,"result":{"tcp":"485454502f312e312033303720466f7262696464656e2052656469726563740d0a4c6f636174696f6e3a20687474703a2f2f3138332e3133362e3133322e32340d0a436f6e74656e742d4c656e6774683a203130350d0a436f6e74656e742d547970653a20746578742f68746d6c0d0a436f6e6e656374696f6e3a20436c6f73650d0a0d0a3c68746d6c3e3c686561643e3c7469746c653e446f6d61696e204e616d6520466f7262696464656e3c2f7469746c653e3c2f686561643e3c626f64793e3c68313e446f6d61696e204e616d6520466f7262696464656e3c2f68313e3c2f626f64793e3c2f68746d6c3e","tls":"{\"domainName\":\"110.42.1.197\",\"ip\":\"\",\"issuer\":\"\",\"commonName\":\"\",\"sans\":null,\"notBefore\":\"\",\"notAfter\":\"\",\"error\":\"EOF\",\"certs\":null,\"port_cont\":\"\"}"},"time":1673405786}
2023/01/11 10:56:27 [+] {"ip":"110.42.1.207","port":1009,"result":{"tcp":"485454502f312e312033303220466f756e640d0a5365727665723a206e67696e780d0a446174653a205765642c203131204a616e20323032332030323a35353a313920474d540d0a436f6e74656e742d547970653a20746578742f68746d6c3b20636861727365743d7574662d380d0a5472616e736665722d456e636f64696e673a206368756e6b65640d0a436f6e6e656374696f6e3a20636c6f73650d0a43616368652d636f6e74726f6c3a206e6f2d63616368652c6d7573742d726576616c69646174650d0a4c6f636174696f6e3a202f696e6465782f757365722f6c6f67696e0d0a0d0a300d0a0d0a","tls":"{\"domainName\":\"110.42.1.207\",\"ip\":\"\",\"issuer\":\"\",\"commonName\":\"\",\"sans\":null,\"notBefore\":\"\",\"notAfter\":\"\",\"error\":\"tls: first record does not look like a TLS handshake\",\"certs\":null,\"port_cont\":\"\"}"},"time":1673405786}
2023/01/11 10:56:27 [+] {"ip":"110.42.2.128","port":443,"result":{"tcp":"485454502f312e312033303720466f7262696464656e2052656469726563740d0a4c6f636174696f6e3a20687474703a2f2f3138332e3133362e3133322e32340d0a436f6e74656e742d4c656e6774683a203130350d0a436f6e74656e742d547970653a20746578742f68746d6c0d0a436f6e6e656374696f6e3a20436c6f73650d0a0d0a3c68746d6c3e3c686561643e3c7469746c653e446f6d61696e204e616d6520466f7262696464656e3c2f7469746c653e3c2f686561643e3c626f64793e3c68313e446f6d61696e204e616d6520466f7262696464656e3c2f68313e3c2f626f64793e3c2f68746d6c3e","tls":"{\"domainName\":\"110.42.2.128\",\"ip\":\"\",\"issuer\":\"\",\"commonName\":\"\",\"sans\":null,\"notBefore\":\"\",\"notAfter\":\"\",\"error\":\"EOF\",\"certs\":null,\"port_cont\":\"\"}"},"time":1673405786}
2023/01/11 10:56:27 [+] {"ip":"110.42.1.197","port":21,"result":{"tcp":"3232302046696c655a696c6c61205365727665722076657273696f6e20302e392e34362062657461207772697474656e2062792054696d204b6f737365202854696d2e4b6f73736540676d782e64652920506c6561736520766973697420687474703a2f2f736f75726365666f7267652e0d0a3530302053796e746178206572726f722c20636f6d6d616e6420756e7265636f676e697a65642e0d0a3530302053796e746178206572726f722c20636f6d6d616e6420756e7265636f676e697a65642e0d0a3530302053796e746178206572726f722c20636f6d6d616e6420756e7265636f676e697a65642e0d0a3530302053796e746178206572726f722c20636f6d6d616e6420756e7265636f676e697a65642e0d0a3530302053796e746178206572726f722c20636f6d6d616e6420756e7265636f676e697a65642e0d0a","tls":"{\"domainName\":\"110.42.1.197\",\"ip\":\"\",\"issuer\":\"\",\"commonName\":\"\",\"sans\":null,\"notBefore\":\"\",\"notAfter\":\"\",\"error\":\"tls: first record does not look like a TLS handshake\",\"certs\":null,\"port_cont\":\"\"}"},"time":1673405787}
```
