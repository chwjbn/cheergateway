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


---获取后端节点记录信息
local function getRecordBackendNodeInfo(nodeKey)
    
	local data={}
	
	data.count=0
	data.time=0.0
	
	local dbHandle=ngx.shared.cheer_cache_db
	
	if not dbHandle then
        return data
    end
	
	local nodeInfoJson=dbHandle:get(nodeKey)
	
	if not nodeInfoJson then
		return data
	end
	
	local jsonLib = require('cjson')
    local nodeInfo=jsonLib.decode(nodeInfoJson)
	
    if not nodeInfo then
        return data
    end
	
	if type(nodeInfo.count)~='number' or type(nodeInfo.time)~='number' then
		return data
	end
	
	data=nodeInfo
	
	return data
	
end

----添加后端节点记录信息
local function addRecordBackendNodeInfo(nodeKey,nodeCount,nodeTime)
    
	if type(nodeCount)~='number' then
		return
	end
	
	if type(nodeTime)~='number' then
		return
	end
	
	local nodeInfo=getRecordBackendNodeInfo(nodeKey)
	
	wafLog('addRecordBackendNodeInfo.getRecordBackendNodeInfo nodeKey=[%s] nodeInfo.count=[%d],nodeInfo.time=[%f]', nodeKey,nodeInfo.count,nodeInfo.time)
	
	nodeInfo.count=nodeInfo.count+nodeCount
	
	nodeInfo.time=(nodeInfo.time+nodeTime)/2
	
	local dbHandle=ngx.shared.cheer_cache_db
	if not dbHandle then
        return
    end
	
	local jsonLib = require('cjson')
    local nodeInfoJson=jsonLib.encode(nodeInfo)

    wafLog('addRecordBackendNodeInfo.nodeInfoJson='..nodeInfoJson)
	
	----5分钟算一次
	local _,xErr,_=dbHandle:set(nodeKey,nodeInfoJson,300)
	if xErr~=nil then
		wafWarn('addRecordBackendNodeInfo.setRecordBackendNodeInfo nodeKey=[%s] nodeInfo.count=[%d],nodeInfo.time=[%f] error=[%s]', nodeKey,nodeInfo.count,nodeInfo.time,xErr)
	end
	
	wafLog('addRecordBackendNodeInfo.setRecordBackendNodeInfo nodeKey=[%s] nodeInfo.count=[%d],nodeInfo.time=[%f]', nodeKey,nodeInfo.count,nodeInfo.time)
	
	return
	
end

local function checkEnd()
	
	---获取当前的后端节点
	local backendNode=ngx.ctx.globalCurrentPeer
	local backendResponseTime=tonumber(ngx.var.upstream_response_time)
	
	if not backendNode then
		backendNode={}
		backendNode.server_ip='0.0.0.0'
		backendNode.server_port=80
	end
	
	local nodeKey=backendNode.server_ip..':'..backendNode.server_port
	
	local nodeAddCount=1
	
	if backendResponseTime>20 then
		nodeAddCount=20
	end
	
	if backendResponseTime>60 then
		nodeAddCount=100
	end
	
	if backendResponseTime>150 then
		nodeAddCount=200
	end
	
	addRecordBackendNodeInfo(nodeKey,nodeAddCount,backendResponseTime)
	
end

wafLog('====================================================================================================')
checkEnd()