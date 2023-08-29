# whois信息获取模块调用文档
- 版本：v0.0.1-20230109

参考链接：

https://github.com/likexian/whois
https://github.com/likexian/whois-parser



`go install https://github.com/likexian/whois/cmd/whois@latest`

## 功能

加载`domain` `ip` `asn`等信息，输出对应的whois信息

## 运行方式
```bash
root@xxx:~# ./whois
Usage:
	./whois [-j] [-h server] domain

domain:
  domain or IPv4 or IPv6 or ASN for query

options:
  -h string specify the whois server
  -j        output format as json
  -v        show the whois version
```

## 使用及输出
### 域名形式输入
#### 输入
```bash
root@xxx:~#./whois -j baidu.com
```
#### 输出
```json
{
  "domain": {
    "id": "11181110_DOMAIN_COM-VRSN",
    "domain": "baidu.com",
    "punycode": "baidu.com",
    "name": "baidu",
    "extension": "com",
    "whois_server": "whois.markmonitor.com",
    "status": [
      "clientdeleteprohibited",
      "clienttransferprohibited",
      "clientupdateprohibited",
      "serverdeleteprohibited",
      "servertransferprohibited",
      "serverupdateprohibited"
    ],
    "name_servers": [
      "ns1.baidu.com",
      "ns2.baidu.com",
      "ns3.baidu.com",
      "ns4.baidu.com",
      "ns7.baidu.com"
    ],
    "created_date": "1999-10-11T11:05:17Z",
    "created_date_in_time": "1999-10-11T11:05:17Z",
    "updated_date": "2022-09-01T03:54:43Z",
    "updated_date_in_time": "2022-09-01T03:54:43Z",
    "expiration_date": "2026-10-11T11:05:17Z",
    "expiration_date_in_time": "2026-10-11T11:05:17Z"
  },
  "registrar": {
    "id": "292",
    "name": "MarkMonitor Inc.",
    "phone": "+1.2086851750",
    "email": "abusecomplaints@markmonitor.com",
    "referral_url": "http://www.markmonitor.com"
  },
  "registrant": {
    "organization": "Beijing Baidu Netcom Science Technology Co., Ltd.",
    "province": "Beijing",
    "country": "CN",
    "email": "select request email form at https://domains.markmonitor.com/whois/baidu.com"
  },
  "administrative": {
    "organization": "Beijing Baidu Netcom Science Technology Co., Ltd.",
    "province": "Beijing",
    "country": "CN",
    "email": "select request email form at https://domains.markmonitor.com/whois/baidu.com"
  },
  "technical": {
    "organization": "Beijing Baidu Netcom Science Technology Co., Ltd.",
    "province": "Beijing",
    "country": "CN",
    "email": "select request email form at https://domains.markmonitor.com/whois/baidu.com"
  }
}
```
### ip形式输入
#### 输入
```bash
root@xxx:~#./whois 1.1.1.1
```

#### 输出
```
% [whois.apnic.net]
% Whois data copyright terms    http://www.apnic.net/db/dbcopyright.html

% Information related to '1.1.1.0 - 1.1.1.255'

% Abuse contact for '1.1.1.0 - 1.1.1.255' is 'helpdesk@apnic.net'

inetnum:        1.1.1.0 - 1.1.1.255
netname:        APNIC-LABS
descr:          APNIC and Cloudflare DNS Resolver project
descr:          Routed globally by AS13335/Cloudflare
descr:          Research prefix for APNIC Labs
country:        AU
org:            ORG-ARAD1-AP
admin-c:        AR302-AP
tech-c:         AR302-AP
abuse-c:        AA1412-AP
status:         ASSIGNED PORTABLE
remarks:        ---------------
remarks:        All Cloudflare abuse reporting can be done via
remarks:        resolver-abuse@cloudflare.com
remarks:        ---------------
mnt-by:         APNIC-HM
mnt-routes:     MAINT-AU-APNIC-GM85-AP
mnt-irt:        IRT-APNICRANDNET-AU
last-modified:  2020-07-15T13:10:57Z
source:         APNIC

irt:            IRT-APNICRANDNET-AU
address:        PO Box 3646
address:        South Brisbane, QLD 4101
address:        Australia
e-mail:         helpdesk@apnic.net
abuse-mailbox:  helpdesk@apnic.net
admin-c:        AR302-AP
tech-c:         AR302-AP
auth:           # Filtered
remarks:        helpdesk@apnic.net was validated on 2021-02-09
mnt-by:         MAINT-AU-APNIC-GM85-AP
last-modified:  2021-03-09T01:10:21Z
source:         APNIC

organisation:   ORG-ARAD1-AP
org-name:       APNIC Research and Development
country:        AU
address:        6 Cordelia St
phone:          +61-7-38583100
fax-no:         +61-7-38583199
e-mail:         helpdesk@apnic.net
mnt-ref:        APNIC-HM
mnt-by:         APNIC-HM
last-modified:  2017-10-11T01:28:39Z
source:         APNIC

role:           ABUSE APNICRANDNETAU
address:        PO Box 3646
address:        South Brisbane, QLD 4101
address:        Australia
country:        ZZ
phone:          +000000000
e-mail:         helpdesk@apnic.net
admin-c:        AR302-AP
tech-c:         AR302-AP
nic-hdl:        AA1412-AP
remarks:        Generated from irt object IRT-APNICRANDNET-AU
abuse-mailbox:  helpdesk@apnic.net
mnt-by:         APNIC-ABUSE
last-modified:  2021-03-09T01:10:22Z
source:         APNIC

role:           APNIC RESEARCH
address:        PO Box 3646
address:        South Brisbane, QLD 4101
address:        Australia
country:        AU
phone:          +61-7-3858-3188
fax-no:         +61-7-3858-3199
e-mail:         research@apnic.net
nic-hdl:        AR302-AP
tech-c:         AH256-AP
admin-c:        AH256-AP
mnt-by:         MAINT-APNIC-AP
last-modified:  2018-04-04T04:26:04Z
source:         APNIC

% Information related to '1.1.1.0/24AS13335'

route:          1.1.1.0/24
origin:         AS13335
descr:          APNIC Research and Development
                6 Cordelia St
mnt-by:         MAINT-AU-APNIC-GM85-AP
last-modified:  2018-03-16T16:58:06Z
source:         APNIC

% This query was served by the APNIC Whois Service version 1.88.16 (WHOIS-AU1)

;; Query time: 2292 msec
;; WHEN: Tue Jan 10 11:47:21 CST 2023
```

### asn形式输入
#### 输入
```bash
root@xxx:~#./whois 60614
```
#### 输出
```bash
% [whois.apnic.net]
% Whois data copyright terms    http://www.apnic.net/db/dbcopyright.html

% Information related to '1.1.1.0 - 1.1.1.255'

% Abuse contact for '1.1.1.0 - 1.1.1.255' is 'helpdesk@apnic.net'

inetnum:        1.1.1.0 - 1.1.1.255
netname:        APNIC-LABS
descr:          APNIC and Cloudflare DNS Resolver project
descr:          Routed globally by AS13335/Cloudflare
descr:          Research prefix for APNIC Labs
country:        AU
org:            ORG-ARAD1-AP
admin-c:        AR302-AP
tech-c:         AR302-AP
abuse-c:        AA1412-AP
status:         ASSIGNED PORTABLE
remarks:        ---------------
remarks:        All Cloudflare abuse reporting can be done via
remarks:        resolver-abuse@cloudflare.com
remarks:        ---------------
mnt-by:         APNIC-HM
mnt-routes:     MAINT-AU-APNIC-GM85-AP
mnt-irt:        IRT-APNICRANDNET-AU
last-modified:  2020-07-15T13:10:57Z
source:         APNIC

irt:            IRT-APNICRANDNET-AU
address:        PO Box 3646
address:        South Brisbane, QLD 4101
address:        Australia
e-mail:         helpdesk@apnic.net
abuse-mailbox:  helpdesk@apnic.net
admin-c:        AR302-AP
tech-c:         AR302-AP
auth:           # Filtered
remarks:        helpdesk@apnic.net was validated on 2021-02-09
mnt-by:         MAINT-AU-APNIC-GM85-AP
last-modified:  2021-03-09T01:10:21Z
source:         APNIC

organisation:   ORG-ARAD1-AP
org-name:       APNIC Research and Development
country:        AU
address:        6 Cordelia St
phone:          +61-7-38583100
fax-no:         +61-7-38583199
e-mail:         helpdesk@apnic.net
mnt-ref:        APNIC-HM
mnt-by:         APNIC-HM
last-modified:  2017-10-11T01:28:39Z
source:         APNIC

role:           ABUSE APNICRANDNETAU
address:        PO Box 3646
address:        South Brisbane, QLD 4101
address:        Australia
country:        ZZ
phone:          +000000000
e-mail:         helpdesk@apnic.net
admin-c:        AR302-AP
tech-c:         AR302-AP
nic-hdl:        AA1412-AP
remarks:        Generated from irt object IRT-APNICRANDNET-AU
abuse-mailbox:  helpdesk@apnic.net
mnt-by:         APNIC-ABUSE
last-modified:  2021-03-09T01:10:22Z
source:         APNIC

role:           APNIC RESEARCH
address:        PO Box 3646
address:        South Brisbane, QLD 4101
address:        Australia
country:        AU
phone:          +61-7-3858-3188
fax-no:         +61-7-3858-3199
e-mail:         research@apnic.net
nic-hdl:        AR302-AP
tech-c:         AH256-AP
admin-c:        AH256-AP
mnt-by:         MAINT-APNIC-AP
last-modified:  2018-04-04T04:26:04Z
source:         APNIC

% Information related to '1.1.1.0/24AS13335'

route:          1.1.1.0/24
origin:         AS13335
descr:          APNIC Research and Development
                6 Cordelia St
mnt-by:         MAINT-AU-APNIC-GM85-AP
last-modified:  2018-03-16T16:58:06Z
source:         APNIC

% This query was served by the APNIC Whois Service version 1.88.16 (WHOIS-AU1)

;; Query time: 2292 msec
;; WHEN: Tue Jan 10 11:47:21 CST 2023

root@iZ2ze6btrk4y6cvivep069Z:~#
root@iZ2ze6btrk4y6cvivep069Z:~# ./whois 60614
% This is the RIPE Database query service.
% The objects are in RPSL format.
%
% The RIPE Database is subject to Terms and Conditions.
% See http://www.ripe.net/db/support/db-terms-conditions.pdf

% Note: this output has been filtered.
%       To receive output for a database update, use the "-B" flag.

% Information related to 'AS59392 - AS61261'

as-block:       AS59392 - AS61261
descr:          RIPE NCC ASN block
remarks:        These AS Numbers are assigned to network operators in the RIPE NCC service region.
mnt-by:         RIPE-NCC-HM-MNT
created:        2020-06-22T15:23:11Z
last-modified:  2020-06-22T15:23:11Z
source:         RIPE

% Information related to 'AS60614'

% Abuse contact for 'AS60614' is 'abuse@steveyi.net'

aut-num:        AS60614
as-name:        STEVEYI-NETWORK
org:            ORG-TY18-RIPE
descr:          SteveYi Network Service
remarks:        ---------------
remarks:        Website: https://network.steveyi.net/
remarks:        Looking Glass: https://lg.steveyi.net/
remarks:        PeeringDB: https://www.peeringdb.com/asn/60614
remarks:        ---------------
remarks:        BGP Communities
remarks:        60614:10001 HKG IT2
remarks:        60614:10002 HKG Apernet
remarks:        60614:10003 HKG HK1
remarks:        60614:10004 Japan Tokyo
remarks:        60614:10005 Taiwan Taipei
remarks:        ---------------
remarks:        60614:100 Learned from Customer / Downstream
remarks:        60614:200 Learned from Upstream
remarks:        60614:300 Learned from Peer
remarks:        60614:400 Learned from IX
remarks:        ---------------
sponsoring-org: ORG-AIL60-RIPE
admin-c:        YT1698-RIPE
tech-c:         YT1698-RIPE
status:         ASSIGNED
mnt-by:         RIPE-NCC-END-MNT
mnt-by:         STEVEYI-MNT
created:        2020-12-18T16:09:01Z
last-modified:  2023-01-06T13:57:20Z
source:         RIPE

organisation:   ORG-TY18-RIPE
descr:          SteveYi Network Service
org-name:       Tsung-Yi Yu
country:        TW
org-type:       OTHER
address:        No. 12, Ln. 60, Sec. 2, Zhongshan Rd. Changhua City, Changhua County 500 , Taiwan (R.O.C.)
abuse-c:        YT1698-RIPE
mnt-ref:        STEVEYI-MNT
mnt-by:         STEVEYI-MNT
created:        2020-01-26T17:35:48Z
last-modified:  2022-12-01T17:28:54Z
source:         RIPE # Filtered

role:           Tsung-Yi Yu
address:        No. 6, Aly. 3, Ln. 108, Fushan St., Changhua City, Changhua County 50036 , Taiwan (R.O.C.)
abuse-mailbox:  abuse@steveyi.net
nic-hdl:        YT1698-RIPE
mnt-by:         STEVEYI-MNT
created:        2020-01-26T19:09:13Z
last-modified:  2022-02-24T20:17:37Z
source:         RIPE # Filtered

% This query was served by the RIPE Database Query Service version 1.105 (DEXTER)

;; Query time: 2161 msec
;; WHEN: Tue Jan 10 11:50:08 CST 2023

```