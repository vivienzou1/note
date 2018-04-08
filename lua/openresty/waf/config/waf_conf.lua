--全局规则配置

--总开关，是否启用waf,on or off(这里必须设置为on，建议在rule_conf下的base.json文件配置)
soyoung_waf_switch = "on"

--规则存放目录
RulePath = "/usr/local/nginx/conf/waf/rule_conf/"

--- base.json 文件绝对路径 [需要自行根据自己服务器情况设置]
base_rule_json = "/usr/local/nginx/conf/waf/rule_conf/base.json"

-- 状态码
sy_error_code_100	= "100"
sy_error_code_200	= "200"
sy_error_code_301	= "301"
sy_error_code_403	= "403"
sy_error_code_404	= "404"
sy_error_code_502	= "502"
sy_error_code_503	= "503"

-- 拦截提示语
-- waf_show_error_html = "soyoung access denied"
waf_show_error_html=[[
<html><head>
<html xmlns="http://www.w3.org/1999/xhtml"><head>
<meta http-equiv="Content-Type" content="text/html; charset=utf-8">
<title>新氧美容整形</title>
<style>p {line-height:20px;}
ul{ list-style-type:none;}
li{ list-style-type:none;}
</style>
</head>
<body style=" padding:0; margin:0; font:14px/1.5 Microsoft Yahei, 宋体,sans-serif; color:#555;">
 <div style="width:200px;
            height: 200px;
            background:green;
            position: absolute;
            left:35%;
            top:30%;
            margin-left:-100px;
            margin-top:-100px;">
  <div style="width:600px; float:left;">
    <div style=" height:40px; line-height:40px; color:#fff; font-size:16px; overflow:hidden; background:#000000; padding-left:20px;"><b>soyoung.com</b></div>
    <div style="border:1px dashed #cdcece; border-top:none; font-size:14px; background:#fff; color:#555; line-height:15px; height:150px; padding:20px 20px 0 20px; overflow-y:auto;background:#f3f7f9;">
      <p style=" margin-top:0px; margin-bottom:0px; margin-left:0px; margin-right:0px; -qt-block-indent:0; text-indent:0px;"><span style=" font-weight:600; color:#fc4f03;">access denied</span></p>
    </div>
  </div>
</div>
</body></html>
]]