local confShowHttpHeader=true

local function setHttpHeader(data_key,data_val)
    if confShowHttpHeader then
		ngx.header[data_key]=data_val
	end
end

--写日志
local function wafLog(fmt,...)
    local arg = { ... }
    local data=string.format(fmt,unpack(arg))
    ngx.log(ngx.INFO,data)
end

--警告日志
local function wafWarn(fmt,...)
    local arg = { ... }
    local data=string.format(fmt,unpack(arg))
    ngx.log(ngx.WARN,data)
end

local function checkBackend()

	local backendNode=ngx.ctx.globalCurrentPeer
	
	local balancerHandle = require('ngx.balancer')
	
	wafLog('checkBackend backendNode.server_ip=%s,backendNode.server_port=%d',backendNode.server_ip,backendNode.server_port)
	
	local dataGatewayNode=backendNode.server_ip..':'..backendNode.server_port
	
	setHttpHeader("X-Gateway-Node",dataGatewayNode)
	
	local ok, err = balancerHandle.set_current_peer(backendNode.server_ip,backendNode.server_port)
	if not ok then
		wafWarn('checkBackend set_current_peer err=%s',err)
		wafWarn('checkBackend error backendNode.server_ip=%s,backendNode.server_port=%d',backendNode.server_ip,backendNode.server_port)
		ngx.exit(ngx.HTTP_NOT_ACCEPTABLE)
	end
	
end

wafLog('====================================================================================================')
checkBackend()