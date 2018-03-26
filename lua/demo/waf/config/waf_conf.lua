--全局规则配置

--总开关，是否启用waf,on or off(这里必须设置为on，建议在rule_conf下的base.json文件配置)
soyoung_waf_switch = "on"

--规则存放目录
RulePath = "/usr/local/nginx/conf/waf/rule_conf/"

--- base.json 文件绝对路径 [需要自行根据自己服务器情况设置]
base_rule_json = "/usr/local/nginx/conf/waf/rule_conf/base.json"

--是否开启攻击信息记录log
attacklog = "on"

--开启攻击信息记录log，需要配置logdir
logdir = "/usr/local/nginx/logs/hack/"

--是否拦截url访问
UrlDeny = "on"

--是否拦截后重定向
Redirect = "on"

--是否拦截cookie攻击
CookieMatch = "on"

--是否拦截post攻击
postMatch = "on"

--是否开启URL白名单
whiteModule = "on" 

--填写不允许上传文件后缀类型
black_fileExt = {"php","jsp"}

--ip白名单，多个ip用逗号分隔
ipWhitelist = {"127.0.0.1",""}

--ip黑名单，多个ip用逗号分隔
ipBlocklist = {"1.0.0.1",""}

--是否开启拦截cc攻击(IP)
Xy_CCDeny = "on"
----设置cc攻击频率，单位为秒.--默认1分钟同一个IP只能请求100次
Xy_CCrate = "100/60"

--是否开启拦截cc攻击(IP+URL)
Xy_CCDeny_URL = "on"
----设置cc攻击频率，单位为秒.--默认1分钟同一个IP只能请求同一个URL 15次
Xy_CCrate_URL = "15/10"

-- 状态码
sy_error_code_100	= "100"
sy_error_code_200	= "200"
sy_error_code_301	= "301"
sy_error_code_403	= "403"
sy_error_code_404	= "404"
sy_error_code_502	= "502"
sy_error_code_503	= "503"

-- 拦截提示语
waf_show_error_html = "access denied"
