#  nmap -script external baidu.com

Starting Nmap 7.01 ( https://nmap.org ) at 2023-01-12 19:32 PST
Pre-scan script results:
| targets-asn:
|_  targets-asn.asn is a mandatory parameter
Nmap scan report for baidu.com (39.156.66.10)
Host is up (0.00054s latency).
Other addresses for baidu.com (not scanned): 110.242.68.66
Not shown: 998 filtered ports
PORT    STATE SERVICE
80/tcp  open  http
|_http-cross-domain-policy: ERROR: Script execution failed (use -d to debug)
| http-xssed:
|
|     UNFIXED XSS vuln.
|
|     	http://youxi.m.baidu.com/softlist.php?cateid=75&amp;phoneid=&amp;url=%22%3E%3Ciframe%20src=http://www.xssed.<br>com%3E
|
|     	http://utility.baidu.com/traf/click.php?id=215&amp;url=http://log0.wordpress.com
|
|     	http://passport.baidu.com/?reg&amp;tpl=sp&amp;return_method=%22%3E%3Ciframe%20src=%22http://xssed.com%22%3E
|
|     	http://zhangmen.baidu.com/search.jsp?f=ms&amp;tn=baidump3&amp;ct=134217728&amp;lf=&amp;rn=&amp;word=%3Cscript%3Ealert%28<br>%27XSS+by+Domino%27%29%3C%2Fscript%3E
|
|     	http://www2.baidu.com/agent/agent_user.php
|
|     	http://www.baidu.com/s?wd=&quot;&gt;&lt;script&gt;alert(document.cookie)&lt;/script&gt;
|
|     	http://www1.baidu.com/s?wd=&quot;&gt;&lt;script&gt;alert(document.cookie)&lt;/script&gt;
|
|     	http://post.baidu.com/f?kw=&quot;&gt;&lt;script&gt;alert(document.cookie)&lt;/script&gt;
|
|
|     FIXED XSS vuln.
|
|     	http://ir.baidu.com/phoenix.zhtml?c=188488&amp;p=irol-infoReqSuccess&amp;t=InfoRequestOnlySave&amp;submit=Submit
|
|_    	http://ir.baidu.com/phoenix.zhtml?c=188488&amp;p=irol-alertsLong&amp;t=AlertPost
443/tcp open  https
| http-cross-domain-policy:
|   VULNERABLE:
|   Cross-domain and Client Access policies.
|     State: LIKELY VULNERABLE
|       A cross-domain policy file specifies the permissions that a web client such as Java, Adobe Flash, Adobe Reader,
|       etc. use to access data across different domains. A client acces policy file is similar to cross-domain policy
|       but is used for M$ Silverlight applications. Overly permissive configurations enables Cross-site Request
|       Forgery attacks, and may allow third parties to access sensitive data meant for the user.
|     Check results:
|       /crossdomain.xml:
|         <?xml version="1.0"?>
|         <cross-domain-policy>
|         	<allow-access-from domain="*.baidu.com" />
|             <allow-access-from domain="*.bdstatic.com" />
|         	<allow-http-request-headers-from domain="*.baidu.com" headers="*"/>
|             <allow-http-request-headers-from domain="*.bdstatic.com" headers="*"/>
|         </cross-domain-policy>
|
|     Extra information:
|       Trusted domains:baidu.com, bdstatic.com
|
|     References:
|       https://www.adobe.com/devnet/articles/crossdomain_policy_file_spec.html
|       https://www.owasp.org/index.php/Test_RIA_cross_domain_policy_%28OTG-CONFIG-008%29
|       http://acunetix.com/vulnerabilities/web/insecure-clientaccesspolicy-xml-file
|       http://sethsec.blogspot.com/2014/03/exploiting-misconfigured-crossdomainxml.html
|       https://www.adobe.com/devnet-docs/acrobatetk/tools/AppSec/CrossDomain_PolicyFile_Specification.pdf
|_      http://gursevkalra.blogspot.com/2013/08/bypassing-same-origin-policy-with-flash.html
| http-xssed:
|
|     UNFIXED XSS vuln.
|
|     	http://youxi.m.baidu.com/softlist.php?cateid=75&amp;phoneid=&amp;url=%22%3E%3Ciframe%20src=http://www.xssed.<br>com%3E
|
|     	http://utility.baidu.com/traf/click.php?id=215&amp;url=http://log0.wordpress.com
|
|     	http://passport.baidu.com/?reg&amp;tpl=sp&amp;return_method=%22%3E%3Ciframe%20src=%22http://xssed.com%22%3E
|
|     	http://zhangmen.baidu.com/search.jsp?f=ms&amp;tn=baidump3&amp;ct=134217728&amp;lf=&amp;rn=&amp;word=%3Cscript%3Ealert%28<br>%27XSS+by+Domino%27%29%3C%2Fscript%3E
|
|     	http://www2.baidu.com/agent/agent_user.php
|
|     	http://www.baidu.com/s?wd=&quot;&gt;&lt;script&gt;alert(document.cookie)&lt;/script&gt;
|
|     	http://www1.baidu.com/s?wd=&quot;&gt;&lt;script&gt;alert(document.cookie)&lt;/script&gt;
|
|     	http://post.baidu.com/f?kw=&quot;&gt;&lt;script&gt;alert(document.cookie)&lt;/script&gt;
|
|
|     FIXED XSS vuln.
|
|     	http://ir.baidu.com/phoenix.zhtml?c=188488&amp;p=irol-infoReqSuccess&amp;t=InfoRequestOnlySave&amp;submit=Submit
|
|_    	http://ir.baidu.com/phoenix.zhtml?c=188488&amp;p=irol-alertsLong&amp;t=AlertPost
| ssl-google-cert-catalog:
|_  No DB entry

Host script results:
| asn-query:
| BGP: 39.156.0.0/17 | Country: CN
|   Origin AS: 9808 - CHINAMOBILE-CN China Mobile Communications Group Co., Ltd., CN
|_    Peer AS: 58453
| dns-blacklist:
|   SPAM
|     bl.nszones.com - DYNAMIC
|     list.quorum.to - FAIL
|     dnsbl.inps.de - FAIL
|     l2.apews.org - FAIL
|     sbl.spamhaus.org - FAIL
|     spam.dnsbl.sorbs.net - FAIL
|     all.spamrats.com - FAIL
|     bl.spamcop.net - FAIL
|   ATTACK
|     all.bl.blocklist.de - FAIL
|   PROXY
|     socks.dnsbl.sorbs.net - FAIL
|     tor.dan.me.uk - FAIL
|     http.dnsbl.sorbs.net - FAIL
|     dnsbl.tornevall.org - FAIL
|_    misc.dnsbl.sorbs.net - FAIL
| hostmap-ip2hosts:
|   hosts:
|     <html xmlns="www.w3.org/1999/xhtml">
| <head>
| <script>document.title='\xBE\xA3\xC3\xC5\xC0\xB7\xB6\xBA\xCA\xB3\xC6\xB7\xD3\xD0\xCF\xDE\xB9\xAB\xCB\xBE';</script>
| <title>&#20122;&#27954;&#22823;&#23610;&#24230;&#97;&#118;&#26080;&#30721;&#19987;&#21306;&#44;&#20122;&#27954;&#20154;&#25104;&#20154;&#55;&#55;&#55;&#55;&#55;&#32593;&#31449;&#44;&#31179;&#38686;&#22312;&#32447;&#35266;&#30475;&#29255;&#26080;&#30721;&#20813;&#36153;&#29233;&#29255;&#44;&#26085;&#26085;&#25720;&#22812;&#22812;&#28155;&#22812;&#22812;&#28155;&#22269;&#20135;&#50;&#48;&#50;&#49;</title>
| <meta name="keywords" content="&#20122;&#27954;&#22823;&#23610;&#24230;&#97;&#118;&#26080;&#30721;&#19987;&#21306;&#44;&#20122;&#27954;&#20154;&#25104;&#20154;&#55;&#55;&#55;&#55;&#55;&#32593;&#31449;&#44;&#31179;&#38686;&#22312;&#32447;&#35266;&#30475;&#29255;&#26080;&#30721;&#20813;&#36153;&#29233;&#29255;&#44;&#26085;&#26085;&#25720;&#22812;&#22812;&#28155;&#22812;&#22812;&#28155;&#22269;&#20135;&#50;&#48;&#50;&#49;" />
| <meta name="description" content="&#22312;&#32447;&#35266;&#30475;&#257;&#29255;&#20813;&#36153;&#20813;&#25773;&#25918;&#22120;
|     &#27431;&#32654;&#29087;&#22919;&#20081;&#23376;&#20262;&#120;&#120;&#35270;&#39057;
|     &#22269;&#20135;&#20122;&#27954;&#39321;&#34121;&#32447;&#25773;&#25918;&#945;&#118;&#51;&#56;
|     &#20122;&#27954;&#25104;&#65;&#86;&#20154;&#19981;&#21345;&#26080;&#30721;&#24433;&#29255;
|     &#27431;&#32654;&#29298;&#20132;&#65;&#27431;&#32654;&#29298;&#20132;&#65;&#8548;&#21478;&#31867;
|     &#65;&#86;&#21943;&#27700;&#39640;&#28526;&#21943;&#27700;&#22312;&#32447;&#35266;&#30475;&#67;&#79;&#77;
|     &#26085;&#26412;&#29305;&#40644;&#29305;&#33394;&#65;&#65;&#65;&#22823;&#29255;&#20813;&#36153;
|     &#104;&#32905;&#26080;&#20462;&#21160;&#28459;&#22312;&#32447;&#35266;&#30475;&#24212;&#29992;" />
| <meta http-equiv="Content-Type" content="text/html; charset=gb2312" />
| </head>
| <script language="javascript" type="text/javascript" src="/common.js"></script>
| <script language="javascript" type="text/javascript" src="/tj.js"></script>
| </body>
|_</html>
|_hostmap-robtex:
| ip-geolocation-geoplugin:
| 39.156.66.10 (baidu.com)
|   coordinates (lat,lon): 34.7732,113.722
|_  state: , China
|_ip-geolocation-maxmind: ERROR: Script execution failed (use -d to debug)
| whois-domain:
|
| Domain name record found at whois.verisign-grs.com
|    Domain Name: BAIDU.COM
|    Registry Domain ID: 11181110_DOMAIN_COM-VRSN
|    Registrar WHOIS Server: whois.markmonitor.com
|    Registrar URL: http://www.markmonitor.com
|    Updated Date: 2022-09-01T03:54:43Z
|    Creation Date: 1999-10-11T11:05:17Z
|    Registry Expiry Date: 2026-10-11T11:05:17Z
|    Registrar: MarkMonitor Inc.
|    Registrar IANA ID: 292
|    Registrar Abuse Contact Email: abusecomplaints@markmonitor.com
|    Registrar Abuse Contact Phone: +1.2086851750
|    Domain Status: clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited
|    Domain Status: clientTransferProhibited https://icann.org/epp#clientTransferProhibited
|    Domain Status: clientUpdateProhibited https://icann.org/epp#clientUpdateProhibited
|    Domain Status: serverDeleteProhibited https://icann.org/epp#serverDeleteProhibited
|    Domain Status: serverTransferProhibited https://icann.org/epp#serverTransferProhibited
|    Domain Status: serverUpdateProhibited https://icann.org/epp#serverUpdateProhibited
|    Name Server: NS1.BAIDU.COM
|    Name Server: NS2.BAIDU.COM
|    Name Server: NS3.BAIDU.COM
|    Name Server: NS4.BAIDU.COM
|    Name Server: NS7.BAIDU.COM
|    DNSSEC: unsigned
|    URL of the ICANN Whois Inaccuracy Complaint Form: https://www.icann.org/wicf/
| >>> Last update of whois database: 2023-01-13T03:36:18Z <<<
|
| For more information on Whois status codes, please visit https://icann.org/epp
|
| NOTICE: The expiration date displayed in this record is the date the
| registrar's sponsorship of the domain name registration in the registry is
| currently set to expire. This date does not necessarily reflect the expiration
| date of the domain name registrant's agreement with the sponsoring
| registrar.  Users may consult the sponsoring registrar's Whois database to
| view the registrar's reported date of expiration for this registration.
|
| TERMS OF USE: You are not authorized to access or query our Whois
| database through the use of electronic processes that are high-volume and
| automated except as reasonably necessary to register domain names or
| modify existing registrations; the Data in VeriSign Global Registry
| Services' ("VeriSign") Whois database is provided by VeriSign for
| information purposes only, and to assist persons in obtaining information
| about or related to a domain name registration record. VeriSign does not
| guarantee its accuracy. By submitting a Whois query, you agree to abide
| by the following terms of use: You agree that you may use this Data only
| for lawful purposes and that under no circumstances will you use this Data
| to: (1) allow, enable, or otherwise support the transmission of mass
| unsolicited, commercial advertising or solicitations via e-mail, telephone,
| or facsimile; or (2) enable high volume, automated, electronic processes
| that apply to VeriSign (or its computer systems). The compilation,
| repackaging, dissemination or other use of this Data is expressly
| prohibited without the prior written consent of VeriSign. You agree not to
| use electronic processes that are automated and high-volume to access or
| query the Whois database except as reasonably necessary to register
| domain names or modify existing registrations. VeriSign reserves the right
| to restrict your access to the Whois database in its sole discretion to ensure
| operational stability.  VeriSign may restrict or terminate your access to the
| Whois database for failure to abide by these terms of use. VeriSign
| reserves the right to modify these terms at any time.
|
| The Registry database contains ONLY .COM, .NET, .EDU domains and
|_Registrars.
| whois-ip: Record found at whois.apnic.net
| inetnum: 39.128.0.0 - 39.191.255.255
| netname: CMNET
| descr: China Mobile Communications Corporation
| country: CN
| orgname: China Mobile
| organisation: ORG-CM1-AP
|_email: hostmaster@chinamobile.com

Nmap done: 1 IP address (1 host up) scanned in 412.38 seconds



# nmap -script external http://www.zorelworld.com/

Starting Nmap 7.01 ( https://nmap.org ) at 2023-01-12 19:47 PST
Pre-scan script results:
| targets-asn:
|_  targets-asn.asn is a mandatory parameter
Unable to split netmask from target expression: "http://www.zorelworld.com/"
WARNING: No targets were specified, so 0 hosts scanned.
Nmap done: 0 IP addresses (0 hosts up) scanned in 0.14 seconds
root@ubuntu:/zrtx/log/cyberspace/2023-01-12# nmap -script external zorelworld.com

Starting Nmap 7.01 ( https://nmap.org ) at 2023-01-12 19:47 PST
Pre-scan script results:
| targets-asn:
|_  targets-asn.asn is a mandatory parameter
Nmap scan report for zorelworld.com (182.92.154.36)
Host is up (0.00035s latency).
Not shown: 998 filtered ports
PORT   STATE SERVICE
22/tcp open  ssh
80/tcp open  http
|_http-cross-domain-policy: ERROR: Script execution failed (use -d to debug)
|_http-xssed: No previously reported XSS vuln.

Host script results:
| asn-query:
| BGP: 182.92.128.0/17 | Country: CN
|   Origin AS: 37963 - ALIBABA-CN-NET Hangzhou Alibaba Advertising Co.,Ltd., CN
|_    Peer AS: 4808 23724
| dns-blacklist:
|   PROXY
|     http.dnsbl.sorbs.net - FAIL
|     misc.dnsbl.sorbs.net - FAIL
|     socks.dnsbl.sorbs.net - FAIL
|     dnsbl.tornevall.org - FAIL
|     tor.dan.me.uk - FAIL
|   SPAM
|     list.quorum.to - FAIL
|     dnsbl.inps.de - FAIL
|     spam.dnsbl.sorbs.net - FAIL
|     bl.spamcop.net - FAIL
|     all.spamrats.com - FAIL
|     bl.nszones.com - FAIL
|     sbl.spamhaus.org - FAIL
|     l2.apews.org - FAIL
|   ATTACK
|_    all.bl.blocklist.de - FAIL
| hostmap-ip2hosts:
|   hosts:
|     <html xmlns="www.w3.org/1999/xhtml">
| <head>
| <script>document.title='\xBE\xA3\xC3\xC5\xC0\xB7\xB6\xBA\xCA\xB3\xC6\xB7\xD3\xD0\xCF\xDE\xB9\xAB\xCB\xBE';</script>
| <title>&#20122;&#27954;&#22823;&#23610;&#24230;&#97;&#118;&#26080;&#30721;&#19987;&#21306;&#44;&#20122;&#27954;&#20154;&#25104;&#20154;&#55;&#55;&#55;&#55;&#55;&#32593;&#31449;&#44;&#31179;&#38686;&#22312;&#32447;&#35266;&#30475;&#29255;&#26080;&#30721;&#20813;&#36153;&#29233;&#29255;&#44;&#26085;&#26085;&#25720;&#22812;&#22812;&#28155;&#22812;&#22812;&#28155;&#22269;&#20135;&#50;&#48;&#50;&#49;</title>
| <meta name="keywords" content="&#20122;&#27954;&#22823;&#23610;&#24230;&#97;&#118;&#26080;&#30721;&#19987;&#21306;&#44;&#20122;&#27954;&#20154;&#25104;&#20154;&#55;&#55;&#55;&#55;&#55;&#32593;&#31449;&#44;&#31179;&#38686;&#22312;&#32447;&#35266;&#30475;&#29255;&#26080;&#30721;&#20813;&#36153;&#29233;&#29255;&#44;&#26085;&#26085;&#25720;&#22812;&#22812;&#28155;&#22812;&#22812;&#28155;&#22269;&#20135;&#50;&#48;&#50;&#49;" />
| <meta name="description" content="&#22312;&#32447;&#35266;&#30475;&#257;&#29255;&#20813;&#36153;&#20813;&#25773;&#25918;&#22120;
|     &#27431;&#32654;&#29087;&#22919;&#20081;&#23376;&#20262;&#120;&#120;&#35270;&#39057;
|     &#22269;&#20135;&#20122;&#27954;&#39321;&#34121;&#32447;&#25773;&#25918;&#945;&#118;&#51;&#56;
|     &#20122;&#27954;&#25104;&#65;&#86;&#20154;&#19981;&#21345;&#26080;&#30721;&#24433;&#29255;
|     &#27431;&#32654;&#29298;&#20132;&#65;&#27431;&#32654;&#29298;&#20132;&#65;&#8548;&#21478;&#31867;
|     &#65;&#86;&#21943;&#27700;&#39640;&#28526;&#21943;&#27700;&#22312;&#32447;&#35266;&#30475;&#67;&#79;&#77;
|     &#26085;&#26412;&#29305;&#40644;&#29305;&#33394;&#65;&#65;&#65;&#22823;&#29255;&#20813;&#36153;
|     &#104;&#32905;&#26080;&#20462;&#21160;&#28459;&#22312;&#32447;&#35266;&#30475;&#24212;&#29992;" />
| <meta http-equiv="Content-Type" content="text/html; charset=gb2312" />
| </head>
| <script language="javascript" type="text/javascript" src="/common.js"></script>
| <script language="javascript" type="text/javascript" src="/tj.js"></script>
| </body>
|_</html>
|_hostmap-robtex:
| ip-geolocation-geoplugin:
| 182.92.154.36 (zorelworld.com)
|   coordinates (lat,lon): 34.7732,113.722
|_  state: , China
|_ip-geolocation-maxmind: ERROR: Script execution failed (use -d to debug)
| whois-domain:
|
| Domain name record found at whois.verisign-grs.com
|    Domain Name: ZORELWORLD.COM
|    Registry Domain ID: 1857671320_DOMAIN_COM-VRSN
|    Registrar WHOIS Server: whois.ename.com
|    Registrar URL: http://www.ename.net
|    Updated Date: 2021-10-27T07:27:32Z
|    Creation Date: 2014-05-07T09:03:50Z
|    Registry Expiry Date: 2025-05-07T09:03:50Z
|    Registrar: eName Technology Co., Ltd.
|    Registrar IANA ID: 1331
|    Registrar Abuse Contact Email: abuse@ename.com
|    Registrar Abuse Contact Phone: 86.4000044400
|    Domain Status: clientDeleteProhibited https://icann.org/epp#clientDeleteProhibited
|    Domain Status: clientTransferProhibited https://icann.org/epp#clientTransferProhibited
|    Name Server: F1G1NS1.DNSPOD.NET
|    Name Server: F1G1NS2.DNSPOD.NET
|    DNSSEC: unsigned
|    URL of the ICANN Whois Inaccuracy Complaint Form: https://www.icann.org/wicf/
| >>> Last update of whois database: 2023-01-13T03:54:36Z <<<
|
| For more information on Whois status codes, please visit https://icann.org/epp
|
| NOTICE: The expiration date displayed in this record is the date the
| registrar's sponsorship of the domain name registration in the registry is
| currently set to expire. This date does not necessarily reflect the expiration
| date of the domain name registrant's agreement with the sponsoring
| registrar.  Users may consult the sponsoring registrar's Whois database to
| view the registrar's reported date of expiration for this registration.
|
| TERMS OF USE: You are not authorized to access or query our Whois
| database through the use of electronic processes that are high-volume and
| automated except as reasonably necessary to register domain names or
| modify existing registrations; the Data in VeriSign Global Registry
| Services' ("VeriSign") Whois database is provided by VeriSign for
| information purposes only, and to assist persons in obtaining information
| about or related to a domain name registration record. VeriSign does not
| guarantee its accuracy. By submitting a Whois query, you agree to abide
| by the following terms of use: You agree that you may use this Data only
| for lawful purposes and that under no circumstances will you use this Data
| to: (1) allow, enable, or otherwise support the transmission of mass
| unsolicited, commercial advertising or solicitations via e-mail, telephone,
| or facsimile; or (2) enable high volume, automated, electronic processes
| that apply to VeriSign (or its computer systems). The compilation,
| repackaging, dissemination or other use of this Data is expressly
| prohibited without the prior written consent of VeriSign. You agree not to
| use electronic processes that are automated and high-volume to access or
| query the Whois database except as reasonably necessary to register
| domain names or modify existing registrations. VeriSign reserves the right
| to restrict your access to the Whois database in its sole discretion to ensure
| operational stability.  VeriSign may restrict or terminate your access to the
| Whois database for failure to abide by these terms of use. VeriSign
| reserves the right to modify these terms at any time.
|
| The Registry database contains ONLY .COM, .NET, .EDU domains and
|_Registrars.
| whois-ip: Record found at whois.apnic.net
| inetnum: 182.92.0.0 - 182.92.255.255
| netname: ALISOFT
| descr: Aliyun Computing Co., LTD
|_country: CN

Nmap done: 1 IP address (1 host up) scanned in 582.86 seconds
