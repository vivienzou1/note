-- require 'api.ip_dict'


ngx.say("hello world<br>");

IP  = ngx.var.remote_addr 

ngx.say("<hr>IP:",IP);

local headers=ngx.req.get_headers()  
local ip=headers["X-REAL-IP"] or headers["X_FORWARDED_FOR"] or ngx.var.remote_addr or "0.0.0.0" 

ngx.say("<hr>ip:",ip);


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

