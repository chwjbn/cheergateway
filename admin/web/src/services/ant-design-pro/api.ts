import { request } from 'umi';

export async function CtlAppGatewayConfigPage(reqData: API.AppGatewayConfigPageRequest) {
  const xDataRet = request<API.AppGatewayConfigPageResponse>('/xapi/cms/app-gateway-config-page', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });

  return xDataRet;
}

export async function CtlAppGatewayConfigAdd(reqData: API.AppGatewayConfigAddRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-config-add', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewayConfigGet(reqData: API.AppDataInfoRequest) {
  return request<API.AppGatewayConfigGetResponse>('/xapi/cms/app-gateway-config-get', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewayConfigSave(reqData: API.AppGatewayConfigSaveRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-config-save', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}


export async function CtlAppGatewayConfigRemove(reqData: API.AppDataInfoRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-config-remove', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewayConfigPub(reqData: API.AppDataInfoRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-config-pub', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewayConfigMapData(reqData: API.AppGatewayConfigMapDataRequest) {
  return request<API.AppGatewayConfigMapDataResponse>('/xapi/cms/app-gateway-config-map-data', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function SearchAppGatewayConfigMapDataWithMode(xSearchKeyword: string,mode?: string) {
  const dataList: any[] = [];

  const reqData: API.AppGatewayConfigMapDataRequest = {};
  reqData.env_type = xSearchKeyword;
  reqData.mode=mode;

  const xRetData = await CtlAppGatewayConfigMapData(reqData);

  if (!xRetData.data) {
    return dataList;
  }

  xRetData.data.forEach((xItem) => {
    dataList.push({ label: xItem.data_name, value: xItem.data_id });
  });

  return dataList;
}

export async function SearchAppGatewayConfigMapData() {
  return SearchAppGatewayConfigMapDataWithMode("",'full');
}

export async function CtlAppGatewayBackendPage(reqData: API.AppGatewayBackendPageRequest) {
  const xDataRet = request<API.AppGatewayBackendPageResponse>(
    '/xapi/cms/app-gateway-backend-page',
    {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      data: reqData,
    },
  );

  return xDataRet;
}

export async function CtlAppGatewayBackendAdd(reqData: API.AppGatewayBackendAddRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-backend-add', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewayBackendGet(reqData: API.AppDataInfoRequest) {
  return request<API.AppGatewayBackendGetResponse>('/xapi/cms/app-gateway-backend-get', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewayBackendSave(reqData: API.AppGatewayBackendSaveRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-backend-save', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewayBackendRemove(reqData: API.AppDataInfoRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-backend-remove', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewayBackendMapData(reqData: API.AppGatewayBackendMapDataRequest) {
  return request<API.AppGatewayBackendMapDataResponse>('/xapi/cms/app-gateway-backend-map-data', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function SearchAppGatewayBackendMapData(configDataId?: string) {
  const dataList: any[] = [];

  const reqData: API.AppGatewayBackendMapDataRequest = {};
  reqData.config_data_id = configDataId;

  const xRetData = await CtlAppGatewayBackendMapData(reqData);

  if (!xRetData.data) {
    return dataList;
  }

  xRetData.data.forEach((xItem) => {
    dataList.push({ label: xItem.data_name, value: xItem.data_id });
  });

  return dataList;
}

export async function CtlAppGatewaySitePage(reqData: API.AppGatewaySitePageRequest) {
  const xDataRet = request<API.AppGatewaySitePageResponse>('/xapi/cms/app-gateway-site-page', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });

  return xDataRet;
}

export async function CtlAppGatewaySiteAdd(reqData: API.AppGatewaySiteAddRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-site-add', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewaySiteGet(reqData: API.AppDataInfoRequest) {
  return request<API.AppGatewaySiteGetResponse>('/xapi/cms/app-gateway-site-get', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewaySiteSave(reqData: API.AppGatewaySiteSaveRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-site-save', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewaySiteRemove(reqData: API.AppDataInfoRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-site-remove', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewaySiteMapData(reqData: API.AppGatewaySiteMapDataRequest) {
  return request<API.AppGatewaySiteMapDataResponse>('/xapi/cms/app-gateway-site-map-data', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function SearchAppGatewaySiteMapData(configDataId?: string) {
  const dataList: any[] = [];

  const reqData: API.AppGatewaySiteMapDataRequest = {};
  reqData.config_data_id = configDataId;

  const xRetData = await CtlAppGatewaySiteMapData(reqData);

  if (!xRetData.data) {
    return dataList;
  }

  xRetData.data.forEach((xItem) => {
    dataList.push({ label: xItem.data_name, value: xItem.data_id });
  });

  return dataList;
}

export async function CtlAppGatewayRulePage(reqData: API.AppGatewayRulePageRequest) {
  const xDataRet = request<API.AppGatewayRulePageResponse>('/xapi/cms/app-gateway-rule-page', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });

  return xDataRet;
}

export async function CtlAppGatewayRuleAdd(reqData: API.AppGatewayRuleAddRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-rule-add', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewayRuleGet(reqData: API.AppDataInfoRequest) {
  return request<API.AppGatewayRuleGetResponse>('/xapi/cms/app-gateway-rule-get', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewayRuleSave(reqData: API.AppGatewayRuleSaveRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-rule-save', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlAppGatewayRuleRemove(reqData: API.AppDataInfoRequest) {
  return request<API.BaseResponse>('/xapi/cms/app-gateway-rule-remove', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}