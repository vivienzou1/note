

-- redis 操作
local redis = require "resty.redis"
local red = redis:new()
red:set_timeout(2000) -- 2 sec
local ok, err = red:connect("172.16.16.4",19000)
if not ok then
    ngx.say("redis failed to connect: ", err)
    return
end
local count, err = red:get_reused_times()

-- 授权，redis有密码
local redis_pwd = ""
if 0 == count then
    if redis_pwd ~= "" then
        local ok, err = red:auth(redis_mod.Password)
        if not ok then
            ngx.say("failed to auth1: ", err)
            return
        end
    end
elseif err then
    ngx.say("failed to auth2: ", err)
    return
end


red:set("soyoung_test_redis_key", "test redis value")
local xy_val = red:get("soyoung_test_redis_key")

ngx.header.content_type = "text/html"
ngx.say(type(xy_val),"<hr>",xy_val)
ngx.exit(ngx.HTTP_FORBIDDEN)


-- test



