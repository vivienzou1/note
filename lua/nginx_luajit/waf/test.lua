local ngx_var = ngx.var
local ngx_ctx = ngx.ctx
local ngx_unescape_uri = ngx.unescape_uri

local optl = require("optl")
local get_argsByName = optl.get_argsByName

--[[
posts = ngx.req.get_post_args()
posts_data = optl.get_table(posts)
ngx.say(posts_data);
ngx.exit(ngx.HTTP_FORBIDDEN)
--]]

--http请求参数  http://47.104.140.87/lua?act=del_ip&user_name=admin&passwd=123456&ip=111.198.24.208
--参数值获取
local _act = get_argsByName("act")
local _userName = get_argsByName("user_name")
local _passwd = get_argsByName("passwd")
local _ip = get_argsByName("ip")

--请求方法
local _method = ngx_var.request_method
local _request_uri = ngx_unescape_uri(ngx_var.request_uri)
local _uri = ngx_unescape_uri(ngx_var.uri)

ngx.say("request:\t",_method,"<hr>request uri:\t",_request_uri,"<hr>method:\t",_uri,"<hr>act:\tact",_act,"<hr>userName:\t",_userName,"<hr>passwd:\t",_passwd,"<hr>ip:\t",_ip);

if _method == "GET" then
	ngx.say("<h3>GET REQUEST</h3>")
elseif _method == "POST" then
	ngx.say("<h3>POST REQUEST</h3>")
else
	ngx.say("<h3>ERROR REQUEST</h3>")
end

ngx.exit(ngx.HTTP_FORBIDDEN)


ngx.say("hello world<br>");

IP  = ngx.var.remote_addr 

ngx.say("<hr>IP:",IP);

local headers=ngx.req.get_headers()  
local ip=headers["X-REAL-IP"] or headers["X_FORWARDED_FOR"] or ngx.var.remote_addr or "0.0.0.0" 


ngx.say("<hr>ip:",ip,"<hr>");

--删除ip
ip = "111.198.24.208"
local ip_vv = ngx.shared.ip_dict:get(ip)
ngx.shared.ip_dict:delete(ip)
ngx.header.content_type = "text/html"
ngx.say(ip,"<hr>",ip_vv,"<hr><br>");

-- 记录错误日志 
-- local ngx_log = ngx.log
-- local ngx_ERR = ngx.ERR
-- ngx_log(ngx_ERR, "failed to startup handler worker...","test...")

--删除ip 111.198.24.208
--ip_dict:delete(ip)

-- test
--[[
if soyoung(true) then
	ngx.header.content_type = "text/html"
	ngx.status = ngx.HTTP_FORBIDDEN
	
	request_guid = ngx_var.request_id
	ngx.say("are you ok<hr>",request_guid)
	ngx.exit(ngx.status)
end
--]]

ngx.exit(ngx.HTTP_FORBIDDEN)