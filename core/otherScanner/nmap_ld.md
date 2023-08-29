$ nmap --script=vuln 127.0.0.1

Starting Nmap 6.40 ( http://nmap.org ) at 2023-01-13 11:31 CST
Pre-scan script results:
| broadcast-avahi-dos:
|   Discovered hosts:
|     172.16.130.166
|     172.16.130.103
|     172.16.130.240
|     172.16.130.104
|   After NULL UDP avahi packet DoS (CVE-2011-1002).
|_  Hosts are all up (not vulnerable).
Nmap scan report for localhost (127.0.0.1)
Host is up (0.000017s latency).
Not shown: 984 closed ports
PORT     STATE SERVICE
21/tcp   open  ftp
22/tcp   open  ssh
25/tcp   open  smtp
| smtp-vuln-cve2010-4344:
|_  The SMTP server is not Exim: NOT VULNERABLE
80/tcp   open  http
|_http-fileupload-exploiter:
|_http-frontpage-login: false
|_http-stored-xss: Couldn't find any stored XSS vulnerabilities.
| http-vuln-cve2011-3192:
|   VULNERABLE:
|   Apache byterange filter DoS
|     State: VULNERABLE
|     IDs:  OSVDB:74721  CVE:CVE-2011-3192
|     Description:
|       The Apache web server is vulnerable to a denial of service attack when numerous
|       overlapping byte ranges are requested.
|     Disclosure date: 2011-08-19
|     References:
|       http://nessus.org/plugins/index.php?view=single&id=55976
|       http://osvdb.org/74721
|       http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2011-3192
|_      http://seclists.org/fulldisclosure/2011/Aug/175
111/tcp  open  rpcbind
443/tcp  open  https
|_http-fileupload-exploiter:
|_http-frontpage-login: false
|_http-stored-xss: Couldn't find any stored XSS vulnerabilities.
5000/tcp open  upnp
5555/tcp open  freeciv
5678/tcp open  rrac
8000/tcp open  http-alt
| http-enum:
|   /docs/README: Interesting, a readme.
|   /docs/: Potentially interesting folder
|_  /web/: Potentially interesting folder
|_http-frontpage-login: false
|_http-huawei-hg5xx-vuln: false
|_http-majordomo2-dir-traversal: ERROR: Script execution failed (use -d to debug)
|_http-vuln-cve2010-0738: false
8001/tcp open  vcom-tunnel
8080/tcp open  http-proxy
| http-enum:
|   /admin/download/backup.sql: Possible database backup
|   /_layouts/download.aspx: MS Sharepoint
|   /downloadFile.php: NETGEAR WNDAP350 2.0.1 to 2.0.9 potential file download and SSH root password disclosure
|   /download/: Potentially interesting folder
|_  /downloads/: Potentially interesting folder
|_http-frontpage-login: false
|_http-huawei-hg5xx-vuln: false
|_http-majordomo2-dir-traversal: ERROR: Script execution failed (use -d to debug)
|_http-phpmyadmin-dir-traversal: ERROR: Script execution failed (use -d to debug)
|_http-vuln-cve2010-0738: false
8088/tcp open  radan-http
|_http-frontpage-login: false
|_http-huawei-hg5xx-vuln: false
|_http-majordomo2-dir-traversal: ERROR: Script execution failed (use -d to debug)
| http-slowloris-check:
|   VULNERABLE:
|   Slowloris DOS attack
|     State: VULNERABLE
|     Description:
|       Slowloris tries to keep many connections to the target web server open and hold them open as long as possible.
|       It accomplishes this by opening connections to the target web server and sending a partial request. By doing
|       so, it starves the http server's resources causing Denial Of Service.
|
|     Disclosure date: 2009-09-17
|     References:
|_      http://ha.ckers.org/slowloris/
|_http-vuln-cve2010-0738: false
| http-vuln-cve2011-3192:
|   VULNERABLE:
|   Apache byterange filter DoS
|     State: VULNERABLE
|     IDs:  OSVDB:74721  CVE:CVE-2011-3192
|     Description:
|       The Apache web server is vulnerable to a denial of service attack when numerous
|       overlapping byte ranges are requested.
|     Disclosure date: 2011-08-19
|     References:
|       http://nessus.org/plugins/index.php?view=single&id=55976
|       http://osvdb.org/74721
|       http://cve.mitre.org/cgi-bin/cvename.cgi?name=CVE-2011-3192
|_      http://seclists.org/fulldisclosure/2011/Aug/175
8443/tcp open  https-alt
8888/tcp open  sun-answerbook
9200/tcp open  wap-wsp

Nmap done: 1 IP address (1 host up) scanned in 734.35 seconds
