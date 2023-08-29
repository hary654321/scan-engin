# Nuclei工具说明

版本：v0.0.1-20230203



## 工具相关链接

[projectdiscovery/nuclei: Fast and customizable vulnerability scanner based on simple YAML based DSL. (github.com)](https://github.com/projectdiscovery/nuclei)

[projectdiscovery/nuclei-templates: Community curated list of templates for the nuclei engine to find security vulnerabilities. (github.com)](https://github.com/projectdiscovery/nuclei-templates)

[Introduction - Nuclei - Community Powered Vulnerability Scanner (projectdiscovery.io)](https://nuclei.projectdiscovery.io/templating-guide/)

[Index - Nuclei - Community Powered Vulnerability Scanner (projectdiscovery.io)](https://nuclei.projectdiscovery.io/nuclei/get-started/#running-nuclei)



## 工具介绍

nuclei工具通过加载nuclei-templates中的模板yaml文件，对目标执行特定探测。

模板yaml文件内容可参考官方Github：[projectdiscovery/nuclei-templates: Community curated list of templates for the nuclei engine to find security vulnerabilities. (github.com)](https://github.com/projectdiscovery/nuclei-templates)

自定义编写规范可参考官方文档：[Introduction - Nuclei - Community Powered Vulnerability Scanner (projectdiscovery.io)](https://nuclei.projectdiscovery.io/templating-guide/)



调用方式

```bash
nuclei -l nuclei_target.txt -t custom -json -silent -o test_nuclei_target.json -irr -sresp

# -l：指定待探测目标文件（IP/域名:端口），每行一个
# -t：指定待扫描模板文件或路径
# -json：输出为json格式
# -irr：保存请求和响应至json
# -sresp：保存请求和响应至当前目录下output目录

其他配置参见https://nuclei.projectdiscovery.io/nuclei/get-started/#running-nuclei
```



参考文档：

[nuclei 应用服务识别调研 (yuque.com)](https://www.yuque.com/jinjiedexiaoqiao/pcggu1/tsfta941dblegscc?singleDoc#) 《nuclei 应用服务识别调研》 密码：sfq4

示例模板和结果格式见文件夹内容