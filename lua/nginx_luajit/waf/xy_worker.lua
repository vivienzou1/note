local _worker_count = ngx.worker.count()
local _worker_id = ngx.worker.id()

local ngx_shared = ngx.shared
local require = require
local ipairs = ipairs
local ngx_log = ngx.log
local ngx_ERR = ngx.ERR
local ngx_thread = ngx.thread
local timer_at = ngx.timer.at
local timer_every = ngx.timer.every
local config_dict = ngx_shared.config_dict
local cjson_safe = require "cjson.safe"

local handler
local handler_all
local handler_zero

local config_base


-- dict 清空过期内存
local function flush_expired_dict()
	local dict_list = {"token_dict","count_dict","config_dict","host_dict","ip_dict","limit_ip_dict"}
	for _,v in ipairs(dict_list) do
		ngx_shared[v]:flush_expired()
	end
end

-- 拉取config_dict配置数据(redis)

-- 推送config_dict、host_dict、count_dict到redis

-- 推送count_dict统计、计数等(redis)

-- 保存config_dict、host_dict到本机文件

handler_zero = function ()
	-- do something	
	
	local config = cjson_safe.decode(config_dict:get("config")) or {}
	config_base = config.base or {}
	local timeAt = config_base.autoSync.timeAt or 5
	-- 如果 auto Sync 开启 就定时从redis 拉取配置并推送一些计数
	if config_base.autoSync.state == "Master" then
		push_Master()
	elseif config_base.autoSync.state == "Slave" then
		if pull_redisConfig() then
			local _debug = "no"
			if config_base.debug_Mod then
				_debug = "yes"
			end
			save_configFile(_debug)
		end
		--推送count_dict到redis
		push_count_dict()
	else
		-- nothing todo
	end

	--清空过期内存
	ngx_thread.spawn(flush_expired_dict)

	--
	local ok, err = timer_at(timeAt, handler_zero)
	if not ok then
		ngx_log(ngx_ERR, "failed to startup handler_zero worker...", err)
	end
end

handler_all = function ()
	local optl = require("optl")
	local dict_config_version = config_dict:get("config_version")
	local optl_config_version = optl.config_version
	if dict_config_version ~= optl_config_version then
		local config = cjson_safe.decode(config_dict:get("config"))
		-- 简单判断config,最好是内容规则的判断
		if config ~= nil then
			optl.config = config
			optl.config_version = dict_config_version
		end
	end
end

handler = function()
	if _worker_id == 0 then
		handler_zero()
	end
	timer_every(1,handler_all)
end

local ok, err = timer_at(0, handler)
if not ok then
	ngx_log(ngx_ERR, "failed to startup handler worker...", err)
end