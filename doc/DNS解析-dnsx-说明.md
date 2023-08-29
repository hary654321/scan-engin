# dnsx

##项目链接
* https://github.com/projectdiscovery/dnsx

## 用法

```shell
dnsx -h

INPUT:
   -l, -list string      list of sub(domains)/hosts to resolve (file or stdin)
   -d, -domain string    list of domain to bruteforce (file or comma separated or stdin)
   -w, -wordlist string  list of words to bruteforce (file or comma separated or stdin)

QUERY:
   -a      query A record (default)
   -aaaa   query AAAA record
   -cname  query CNAME record
   -ns     query NS record
   -txt    query TXT record
   -ptr    query PTR record
   -mx     query MX record
   -soa    query SOA record
   -axfr   query AXFR
   -caa    query CAA record

FILTER:
   -re, -resp          display dns response
   -ro, -resp-only     display dns response only
   -rc, -rcode string  filter result by dns status code (eg. -rcode noerror,servfail,refused)

PROBE:
   -cdn  display cdn name

RATE-LIMIT:
   -t, -threads int      number of concurrent threads to use (default 100)
   -rl, -rate-limit int  number of dns request/second to make (disabled as default) (default -1)

OUTPUT:
   -o, -output string  file to write output
   -json               write output in JSONL(ines) format

DEBUG:
   -hc, -health-check  run diagnostic check up
   -silent             display only results in the output
   -v, -verbose        display verbose output
   -raw, -debug        display raw dns response
   -stats              display stats of the running scan
   -version            display version of dnsx

OPTIMIZATION:
   -retry int                number of dns attempts to make (must be at least 1) (default 2)
   -hf, -hostsfile           use system host file
   -trace                    perform dns tracing
   -trace-max-recursion int  Max recursion for dns trace (default 32767)
   -resume                   resume existing scan
   -stream                   stream mode (wordlist, wildcard, stats and stop/resume will be disabled)

CONFIGURATIONS:
   -r, -resolver string          list of resolvers to use (file or comma separated)
   -wt, -wildcard-threshold int  wildcard filter threshold (default 5)
   -wd, -wildcard-domain string  domain name for wildcard filtering (other flags will be ignored)
```

## 样例
### DNS解析：

* PTR使用查询从给定的网络范围中提取子域：
```json
echo 173.0.84.0/24 | dnsx -silent -resp-only -ptr
        
api-aa-3t.paypal.com
business.paypal.com
www.cybercash.com
waf-origin-2-api-11.paypal.com
ssp.paypal.com
reports.paypal.com
business-cog.paypal.com
sspserv.paypal.com
cld-edge-origin-api.paypal.com
m.paypal.fr
m.paypal.co.uk
merchant.paypal.com
he.paypal.com
```

* 在给定的（子）域列表上使用DNS 状态代码进行探测：
```shell
subfinder -silent -d hackerone.com | dnsx -silent -rcode noerror,servfail,refused

ns.hackerone.com [NOERROR]
a.ns.hackerone.com [NOERROR]
b.ns.hackerone.com [NOERROR]
support.hackerone.com [NOERROR]
resources.hackerone.com [NOERROR]
mta-sts.hackerone.com [NOERROR]
www.hackerone.com [NOERROR]
mta-sts.forwarding.hackerone.com [NOERROR]
```

* 从各种来源获得的被动子域列表中过滤活动主机名：:

```console
subfinder -silent -d hackerone.com | dnsx -silent

a.ns.hackerone.com
www.hackerone.com
api.hackerone.com
docs.hackerone.com
mta-sts.managed.hackerone.com
mta-sts.hackerone.com
resources.hackerone.com
b.ns.hackerone.com
mta-sts.forwarding.hackerone.com
events.hackerone.com
support.hackerone.com
```
* 为给定的子域列表打印A记录：:

```console
subfinder -silent -d hackerone.com | dnsx -silent -a -resp

www.hackerone.com [104.16.100.52]
www.hackerone.com [104.16.99.52]
hackerone.com [104.16.99.52]
hackerone.com [104.16.100.52]
api.hackerone.com [104.16.99.52]
api.hackerone.com [104.16.100.52]
mta-sts.forwarding.hackerone.com [185.199.108.153]
mta-sts.forwarding.hackerone.com [185.199.109.153]
mta-sts.forwarding.hackerone.com [185.199.110.153]
mta-sts.forwarding.hackerone.com [185.199.111.153]
a.ns.hackerone.com [162.159.0.31]
resources.hackerone.com [52.60.160.16]

```

* 提取A记录:

```console
subfinder -silent -d hackerone.com | dnsx -silent -a -resp-only

104.16.99.52
104.16.100.52
162.159.1.31
104.16.99.52
104.16.100.52
185.199.110.153
185.199.111.153
185.199.108.153
185.199.109.153
104.16.99.52
104.16.100.52
104.16.51.111
104.16.53.111
185.199.108.153
185.199.111.153
185.199.110.153
185.199.111.153
```

* 提取CNAME:

```console
subfinder -silent -d hackerone.com | dnsx -silent -cname -resp

support.hackerone.com [hackerone.zendesk.com]
resources.hackerone.com [read.uberflip.com]
mta-sts.hackerone.com [hacker0x01.github.io]
mta-sts.forwarding.hackerone.com [hacker0x01.github.io]
events.hackerone.com [whitelabel.bigmarker.com]
```

* 在给定的（子）域列表上使用DNS 状态代码进行探测:
```console
subfinder -silent -d hackerone.com | dnsx -silent -rcode noerror,servfail,refused

ns.hackerone.com [NOERROR]
a.ns.hackerone.com [NOERROR]
b.ns.hackerone.com [NOERROR]
support.hackerone.com [NOERROR]
resources.hackerone.com [NOERROR]
mta-sts.hackerone.com [NOERROR]
www.hackerone.com [NOERROR]
mta-sts.forwarding.hackerone.com [NOERROR]
docs.hackerone.com [NOERROR]
```

### DNS暴力破解

* 使用“d”和“w”标志的给定域或域列表的暴力破解子域:

```console
dnsx -silent -d facebook.com -w dns_worldlist.txt

blog.facebook.com
booking.facebook.com
api.facebook.com
analytics.facebook.com
beta.facebook.com
apollo.facebook.com
ads.facebook.com
box.facebook.com
alpha.facebook.com
apps.facebook.com
connect.facebook.com
c.facebook.com
careers.facebook.com
code.facebook.com
```

* 使用单个或多个关键字输入的暴力破解目标子域：

```console
dnsx -silent -d domains.txt -w jira,grafana,jenkins

grafana.1688.com
grafana.8x8.vc
grafana.airmap.com
grafana.aerius.nl
jenkins.1688.com
jenkins.airbnb.app
jenkins.airmap.com
jenkins.ahn.nl
jenkins.achmea.nl
jira.amocrm.com
jira.amexgbt.com
jira.amitree.com
jira.arrival.com
jira.atlassian.net
jira.atlassian.com
```


### 通配符筛选


```console
dnsx -l subdomain_list.txt -wd airbnb.com -o output.txt
```
