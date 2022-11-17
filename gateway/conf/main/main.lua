local confRedisHost='127.0.0.1'
local confRedisPort='6379'
local confRedisDb=0
local confRedisPassword='Ik#idc#redis#888#999'

local confShowHttpHeader=true

--iphash/mincount/mintime
local confBalancePolicy='mincount'

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

--分解字符串
local function stringSplit(strData,strDelim)
	
	local tData={}
	
	if type(strDelim) ~= "string" or string.len(strDelim) <= 0 then
		return tData
	end
	
	local startIndex = 1
	
	while true do
	
		 local pos = string.find (strData,strDelim, startIndex, true)
		 
		 if not pos then
			break
		 end
		 
		 table.insert (tData,string.sub(strData, startIndex, pos - 1))
		 startIndex = pos + string.len (strDelim)
	
	end
	
	table.insert (tData,string.sub(strData,startIndex))
	
	return tData
end

--字符串包含,不能用string.find
local function stringContain(strData,subStrData)

	local bRet=false
	
	local nLen=string.len(strData)
	local nSubLen=string.len(subStrData)
	
	if nSubLen<1 then
		bRet=true
		return bRet
	end
	
	if nLen<nSubLen then
		return bRet
	end
	
	local startIndex=1
	
	while(startIndex<=nLen)
	do
		if string.byte(strData,startIndex)==string.byte(subStrData,1) then
			
			local endIndex=startIndex+nSubLen-1
			if endIndex>nLen then
				endIndex=nLen
			end
			
			local matchSubStr=string.sub(strData,startIndex,endIndex)
			
			if matchSubStr==subStrData then
				bRet=true
				return bRet
			end
			
		end
		
		startIndex=startIndex+1
	end
	
	return bRet

end

--字符串正则匹配
local function stringRegexMatched(strData,strRule)

	local bRet=false
	
	if ngx.re.match(strData,strRule,'jio') then
		bRet=true
	end
	
	return bRet

end

--字符串在字符串列表
local function stringInList(strData,strListData)
	
	local bRet=false
	
	local strList=stringSplit(strListData,'|')
	
	for k,v in ipairs(strList) do
		
		if strData==v then
			bRet=true
			return bRet
		end
		
	end
	
	return bRet
end

--获取一个redis客户端
local function getRedisClient()
    local redisLib = require('resty.redis')
    local redisHandle=redisLib.new()
    local ok,err=redisHandle.connect(redisHandle, confRedisHost, confRedisPort)

    if not ok then
        wafWarn('getRedisClient redis error: %s', err)
        redisHandle=nil
    end

    if not confRedisPassword then
        return redisHandle
    end
	
	if not redisHandle then
		return redisHandle
	end

    local authRet,authErr=redisHandle:auth(confRedisPassword)

    if not authRet then
        wafWarn('getRedisClient redis auth error: %s', authErr)
        redisHandle=nil
    end

    return redisHandle
end

--关闭redis连接池
local function closeRedis(redisHandle)
    if not redisHandle then
        return
    end

    local pool_max_idle_time = 10000
    local pool_size = 100
    local ok, err = redisHandle:set_keepalive(pool_max_idle_time, pool_size)

    if not ok then
        wafWarn('closeRedis set keepalive error: %s',err)
    end

end

--从Redis获取
local function getDataFromRedis(dataKey)
    local data=nil

    local redisHandle=getRedisClient()
    if not redisHandle then
        return data
    end

    --选择指定的redis库
    redisHandle:select(confRedisDb)

    local res=redisHandle:get(dataKey)

    closeRedis(redisHandle)

    wafLog('getDataFromRedis dataKey=%s',dataKey)

    if not res then
        wafLog('getDataFromRedis redis get error: %s', dataKey)
    else
        data=res
    end

    return data
end


--从全局缓存中获取
local function getDataFromCacheDb(dataKey)
    local data={}
    local dbHandle=ngx.shared.cheer_cache_db

    if not dbHandle then
        return data
    end

    --缓存中获取
    local dataVal=dbHandle:get(dataKey)

    --是否需要更新到缓存
    local bNeedUpdate=false

    --缓存中没有,从Redis获取
    if not dataVal then
        dataVal=getDataFromRedis(dataKey)
        bNeedUpdate=true
    end

    --都没有直接返回空table
    if not dataVal then
        return data
    end

    --不是字符串
    if type(dataVal)~='string' then
        return data
    end

    --更新到缓存
    if bNeedUpdate then
        dbHandle:set(dataKey,dataVal,30);
    end

    --反解析为table
    local jsonLib = require('cjson')
    data=jsonLib.decode(dataVal)

    if not data then
        data={}
    end

    return data
end

--逻辑节点计算
local function logicNodeCalc(contextData,logicMatchTarget,logicMatchOp,logicMatchData)
	local bRet=false
	
	local xCalcTargetData=contextData[logicMatchTarget]
	
	--上下文中不存在此逻辑计算对象
	if not xCalcTargetData then
		return bRet
	end
	
	--数据类型
	local xDataTypePrefix=string.sub(logicMatchTarget,1,2)
	
	if logicMatchOp=='eq' then
		if xCalcTargetData==logicMatchData then
			bRet=true
			return bRet
		end
	end
	
	if logicMatchOp=='lt' then
		if xCalcTargetData<logicMatchData then
			bRet=true
			return bRet
		end
	end
	
	if logicMatchOp=='gt' then
		if xCalcTargetData>logicMatchData then
			bRet=true
			return bRet
		end
	end
	
	if logicMatchOp=='lte' then
		if xCalcTargetData<=logicMatchData then
			bRet=true
			return bRet
		end
	end
	
	if logicMatchOp=='gte' then
		if xCalcTargetData>=logicMatchData then
			bRet=true
			return bRet
		end
	end
	
	if logicMatchOp=='neq' then
		if xCalcTargetData~=logicMatchData then
			bRet=true
			return bRet
		end
	end
	
	--针对字符串生效的逻辑判断
	if xDataTypePrefix=='s_' then
	
		if logicMatchOp=='regex' then
			if stringRegexMatched(xCalcTargetData,logicMatchData) then
				bRet=true
				return bRet
			end
		end

		if logicMatchOp=='notregex' then
			if not stringRegexMatched(xCalcTargetData,logicMatchData) then
				bRet=true
				return bRet
			end
		end
		
		if logicMatchOp=='contain' then
			if stringContain(xCalcTargetData,logicMatchData) then
				bRet=true
				return bRet
			end
		end
		
		if logicMatchOp=='notcontain' then
			if not stringContain(xCalcTargetData,logicMatchData) then
				bRet=true
				return bRet
			end
		end
		
		if logicMatchOp=='in' then
			if stringInList(xCalcTargetData,logicMatchData) then
				bRet=true
				return bRet
			end
		end
		
		if logicMatchOp=='notin' then
			if not stringInList(xCalcTargetData,logicMatchData) then
				bRet=true
				return bRet
			end
		end
	
	end
	
	return bRet
end


--获取站点规则列表
local function getSiteDataTable()
    local data=getDataFromCacheDb('cheer-gateway-site')
    if not data then
        data={}
    end
    return data
end


--获取站点后端节点路由规则列表
local function getRuleDataTable(site_data_id)
    local dataKey=string.format('cheer-gateway-rule-%s',site_data_id)
    local data=getDataFromCacheDb(dataKey)
    if not data then
        data={}
    end
    return data
end


--获取后端服务器节点信息
local function getBackendDataTable(backend_data_id)
    local dataKey=string.format('cheer-gateway-backend-%s',backend_data_id)
    local data=getDataFromCacheDb(dataKey)
    if not data then
        data={}
    end
    return data
end


--获取客户端IP
local function getClientIp()
    
	local header = ngx.req.get_headers()
    
	local data=header['X-Real-IP']
	
	if not data then
	
		local xClientIpStr=ngx.var.proxy_add_x_forwarded_for
		
		if xClientIpStr then
			local xClientIpTable=stringSplit(xClientIpStr,',')
			if table.getn(xClientIpTable)>1 then
				data=xClientIpTable[1]
			end
		end

	end

	if not data then
        data=header['X-Forwarded-For']
    end
	
    if not data then
        data=ngx.var.remote_addr
    end

    if not data then
        data='0.0.0.0'
    end

    return data
end


--获取业务租户ID
local function getMerchantId()
    
	local httpHeader=ngx.req.get_headers()
	
	--从http头里面拿
	local data=httpHeader['X-Tenant']
	

    --从请求参数拿ik_merchant_id
	if not data then
		data=ngx.var.arg_ik_merchant_id
	end

    --没有从cookie拿ik_merchant_id
    if not data then
        data=ngx.var.cookie_ik_merchant_id
    end
	
	--从请求参数拿merchantId
	if not data then
        data=ngx.var.arg_merchantId
    end
	
	--从请求参数拿merchant_id
	if not data then
        data=ngx.var.arg_merchant_id
    end
	
	--从请求参数拿sso_merchant_id
	if not data then
        data=ngx.var.arg_sso_merchant_id
    end
	
	--从cookie拿sso_merchant_id
	if not data then
        data=ngx.var.cookie_sso_merchant_id
    end

    if not data then
        data='none'
    end

    return data
end


--生成CheerId
local function genCheerId()
	local data=getClientIp()
	local timeStr=os.date("%Y%m%d%H%M%S")
	local randNum=math.random()
	data=data..timeStr..randNum
	data=ngx.md5(data)
	data='cheerid_'..data
	return data
end


--获取环境数据
local function getContextData()
    local data={}
	
	local cookieNeedSet={}

    --获取客户端IP
    data.s_http_ip=getClientIp()
    if not data.s_http_ip then
        data.s_http_ip='0.0.0.0'
    end

    --获取域名
    data.s_http_header_host=ngx.var.host
    if not data.s_http_header_host then
        data.s_http_header_host=''
    end
	
	--获取http方法
	data.s_http_header_method=ngx.var.request_method
	if not data.s_http_header_method then
		data.s_http_header_method=''
	end
	
	--获取useragent
	data.s_http_header_useragent=ngx.var.http_user_agent
	if not data.s_http_header_useragent then
		data.s_http_header_useragent=''
	end
	
	--获取请求url
	data.s_http_header_url=ngx.var.request_uri
	if not data.s_http_header_url then
		data.s_http_header_url=''
	end
	
	--获取来源url
	data.s_http_header_referer=ngx.var.http_referer
	if not data.s_http_header_referer then
		data.s_http_header_referer=''
	end

    --获取业务租户ID
    data.s_biz_merchant_id=getMerchantId()
    if not data.s_biz_merchant_id then
        data.s_biz_merchant_id='none'
    end
	
	--设置访问ID,链路追踪用
	data.s_cookie_cheerid=ngx.var.cookie_cheer_cheerid
	if not data.s_cookie_cheerid then
		data.s_cookie_cheerid=genCheerId()
		local cookieCheerId= 'cheer_cheerid='..data.s_cookie_cheerid..'; Path=/; Expires=' .. ngx.cookie_time(ngx.time() + 2592000)..'; HttpOnly'
		table.insert(cookieNeedSet,cookieCheerId)
	end
	
	--设置租户代码
	local cookieMerchantId = 'cheer_merchant_id='..data.s_biz_merchant_id..'; Path=/; Expires=' .. ngx.cookie_time(ngx.time() + 2592000)..'; HttpOnly'
	table.insert(cookieNeedSet,cookieMerchantId)
	
	setHttpHeader('Set-Cookie',cookieNeedSet)
	

    return data
end

--获取对应的域名匹配的规则站点
local function getHostMatchedSite(host)

    local data={}

    data.data_id=''
    data.rule_type='none'
    data.rule_data=''
    data.default_backend_data_id=''

    local siteData=getSiteDataTable()

    for k,v in ipairs(siteData) do
	
        local xSiteDataId=v.data_id
        local xRuleType=v.rule_type
        local xRuleData=v.rule_data
        local xDefaultBackendDataId=v.default_backend_data_id

        if xRuleType=='eq' then
		
            if xRuleData==host then
				data.data_id=xSiteDataId
				data.rule_type=xRuleType
				data.rule_data=xRuleData
				data.default_backend_data_id=xDefaultBackendDataId
                return data
            end
			
        end
		
		if xRuleType=='neq' then
		
            if xRuleData~=host then
				data.data_id=xSiteDataId
				data.rule_type=xRuleType
				data.rule_data=xRuleData
				data.default_backend_data_id=xDefaultBackendDataId
                return data
            end
			
        end

        if xRuleType=='regex' then
		
            if stringRegexMatched(host,xRuleData) then
                data.data_id=xSiteDataId
				data.rule_type=xRuleType
				data.rule_data=xRuleData
				data.default_backend_data_id=xDefaultBackendDataId
                return data
            end
			
        end
		
		if xRuleType=='notregex' then
		
            if not stringRegexMatched(host,xRuleData) then
                data.data_id=xSiteDataId
				data.rule_type=xRuleType
				data.rule_data=xRuleData
				data.default_backend_data_id=xDefaultBackendDataId
                return data
            end
			
        end
		
		
		if xRuleType=='contain' then
		
            if stringContain(host,xRuleData) then
                data.data_id=xSiteDataId
				data.rule_type=xRuleType
				data.rule_data=xRuleData
				data.default_backend_data_id=xDefaultBackendDataId
                return data
            end
			
        end
		
		if xRuleType=='notcontain' then
		
            if not stringContain(host,xRuleData) then
                data.data_id=xSiteDataId
				data.rule_type=xRuleType
				data.rule_data=xRuleData
				data.default_backend_data_id=xDefaultBackendDataId
                return data
            end
			
        end
		
		if xRuleType=='in' then
		
            if stringInList(host,xRuleData) then
                data.data_id=xSiteDataId
				data.rule_type=xRuleType
				data.rule_data=xRuleData
				data.default_backend_data_id=xDefaultBackendDataId
                return data
            end
			
        end
		
		if xRuleType=='notin' then
		
            if not stringInList(host,xRuleData) then
                data.data_id=xSiteDataId
				data.rule_type=xRuleType
				data.rule_data=xRuleData
				data.default_backend_data_id=xDefaultBackendDataId
                return data
            end
			
        end
		

    end

    return data

end

--获取当前请求环境下命中的节点路由规则
local function getSiteMatchedRule(contextData,site_data_id)

    local data={}

    data.data_id=''
    data.site_data_id=''
	
	data.action_type=''
	data.action_data=''
	
    --获取站点对应的规则列表
    local xRuleData=getRuleDataTable(site_data_id)

    for k,v in ipairs(xRuleData) do

		local xRuleDataId=v.data_id
        local xRuleSiteDataId=v.site_data_id
		
		local xRuleMatchTarget=v.match_target
		local xRuleMatchOp=v.match_op
        local xRuleMatchData=v.match_data
		
        local xRuleActionType=v.action_type
		local xRuleActionData=v.action_data

       
		--当前规则逻辑判断
        if logicNodeCalc(contextData,xRuleMatchTarget,xRuleMatchOp,xRuleMatchData) then
            
			data.data_id=xRuleDataId
			data.site_data_id=xRuleSiteDataId

			data.action_type=xRuleActionType
			data.action_data=xRuleActionData
			
			wafLog('logicNodeCalc matched xRuleDataId=%s',xRuleDataId)
			
            return data
        end
		
    end

    return data
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

---通过IPHash策略获取后端节点
local function getBackendNodeInfoByIpHash(backendNodeList,contextData)

	local data={}
	
	data.server_ip='0.0.0.0'
	data.server_port=9000

	local nodeCount=table.getn(backendNodeList)
	
	local clientKey=contextData.s_cookie_cheerid	
	if not clientKey then
		clientKey=contextData.s_http_ip
	end
	
	local clientHash=ngx.crc32_long(clientKey)
	local clientIndex=(clientHash%nodeCount)+1
	
	data=backendNodeList[clientIndex]
	
	return data
	
end

---通过最小访问次数获取后端节点
local function getBackendNodeInfoByMinCount(backendNodeList)
	
	local data={}
	
	data.server_ip='0.0.0.0'
	data.server_port=9000
	
	local xMinCount=99999999
	
	for k,v in ipairs(backendNodeList) do
	
		local xNodeKey=v.server_ip..':'..v.server_port
		local xNodeInfo=getRecordBackendNodeInfo(xNodeKey)
		
		if xMinCount>xNodeInfo.count then
			data=v
			xMinCount=xNodeInfo.count
		end
		
	end
	
	return data
	
end


---通过最小响应时间获取后端节点
local function getBackendNodeInfoByMinTime(backendNodeList)
	
	local data={}
	
	data.server_ip='0.0.0.0'
	data.server_port=9000
	
	local xMinTime=99999999
	
	for k,v in ipairs(backendNodeList) do
	
		local xNodeKey=v.server_ip..':'..v.server_port
		local xNodeInfo=getRecordBackendNodeInfo(xNodeKey)
		
		if xMinTime>xNodeInfo.time then
			data=v
			xMinTime=xNodeInfo.time
		end
		
	end
	
	return data
	
end


--获取后端节点对应的服务器信息
local function getBackendNodeInfo(contextData,backend_data_id)
	
	local data={}
	
	data.server_ip='0.0.0.0'
	data.server_port=9000
	
	local xBackendData=getBackendDataTable(backend_data_id)
	
	if not xBackendData then
		return data
	end
	
	if not xBackendData.node_addr then
		return data
	end
	
	local xBackendNodeDataList=stringSplit(xBackendData.node_addr,'|')
	if not xBackendNodeDataList then
		return data
	end
		
	local backendNodeList={}
	
	for k,v in ipairs(xBackendNodeDataList) do
		
		local dataItem={}
		local xNodeInfoTable=stringSplit(v,':')
		
		if table.getn(xNodeInfoTable)==2 then
			dataItem.server_ip=xNodeInfoTable[1]
			dataItem.server_port=tonumber(xNodeInfoTable[2])
			table.insert(backendNodeList,dataItem)
		end
		
		
	end
	
	if not backendNodeList then
		return data
	end
	
	
	local xNodeInfo={}
	
	---IPHash负载
	if confBalancePolicy=='iphash' then
		xNodeInfo=getBackendNodeInfoByIpHash(backendNodeList,contextData)
	end
	
	---最小连接次数
	if confBalancePolicy=='mincount' then
		xNodeInfo=getBackendNodeInfoByMinCount(backendNodeList)
	end
	
	----最小响应时间
	if confBalancePolicy=='mintime' then
		xNodeInfo=getBackendNodeInfoByMinTime(backendNodeList)
	end
	
	if not xNodeInfo then
		return data
	end
	
	if not xNodeInfo.server_ip then
		return data
	end
	
	if not xNodeInfo.server_port then
		return data
	end
	
	data=xNodeInfo
	
	return data
end

--获取当前站点后端节点
local function getCurrentContextBackend(contextData)

    local xBackendDataId=''

    local xMatchedSiteData=getHostMatchedSite(contextData.s_http_header_host)

    wafLog('getHostMatchedSite xMatchedSiteData.data_id=[%s],xMatchedSiteData.default_backend_data_id=[%s]', xMatchedSiteData.data_id,xMatchedSiteData.default_backend_data_id)

    --没有找到匹配站点
    if xMatchedSiteData.data_id=='' then
        return xBackendDataId
    end
	
	setHttpHeader("X-Gateway-SiteId",xMatchedSiteData.data_id)
	
	--初始赋值给默认节点ID
	xBackendDataId=xMatchedSiteData.default_backend_data_id
 
	--获取当前请求环境下命中的节点路由规则
	local xMatchedRuleData=getSiteMatchedRule(contextData,xMatchedSiteData.data_id)
	
	--没有找到匹配规则
	if xMatchedRuleData.data_id=='' then
		return xBackendDataId
	end
	
	wafLog('getSiteMatchedRule xMatchedRuleData.data_id=[%s],xMatchedRuleData.action_type=[%s],xMatchedSiteData.action_data=[%s]', xMatchedRuleData.data_id,xMatchedRuleData.action_type,xMatchedRuleData.action_data)
	
	--如果动作是后端节点
	if xMatchedRuleData.action_type=='backend' then
		xBackendDataId=xMatchedRuleData.action_data
		return xBackendDataId
	end
	
    return xBackendDataId
end




--拒绝访问
local function rejectAccess(errorCode)

	local pageData='<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd"><html xmlns="http://www.w3.org/1999/xhtml"><head><meta http-equiv="Content-Type" content="text/html; charset=utf-8" /><title>CheerGateway</title></head><body><center>系统升级或者访问用户超出系统限制,服务暂时不可用,错误码:CHEER_CODE</center></body></html>'
	pageData=string.gsub(pageData,'CHEER_CODE',errorCode)

	ngx.status=503
	ngx.send_headers()
	ngx.say(pageData)
	ngx.exit(0)
end


local function checkAccess()
    
	setHttpHeader("X-Gateway-Name",'CheerGateway/5.0')

	--获取环境上下文数据
    local xContextData=getContextData()
	
	--获取当前环境上下文下的后端节点
    local xBackendDataId=getCurrentContextBackend(xContextData)

    --没有找到后端节点
    if not xBackendDataId then
        rejectAccess(200404)
    end
	
	setHttpHeader("X-Gateway-NodeId",xBackendDataId)
	
	--没有找到后端节点
	if xBackendDataId=='' then
        rejectAccess(200404)
    end
	
	local backendNode=getBackendNodeInfo(xContextData,xBackendDataId)
	
	if not backendNode then
		rejectAccess(200405)
	end
	
	if backendNode.server_ip=='0.0.0.0' then
		rejectAccess(200406)
	end
	
	wafLog('checkBackend backendNode.server_ip=%s,backendNode.server_port=%d',backendNode.server_ip,backendNode.server_port)
	
	ngx.ctx.globalCurrentPeer=backendNode
	

end

wafLog('====================================================================================================')
checkAccess()