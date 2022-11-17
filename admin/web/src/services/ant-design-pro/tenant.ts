import { request } from 'umi';
import { ClearLocalCache, GetLocalCache, SetLocalCache } from './lib';

export function getEnvData(): API.XEnvData | undefined {
  const xHost = window.location.host;

  const xData: API.XEnvData = {
    appEnv: 'prod',
    appHost: xHost,
    appWsHost: xHost,
  };

  if (xHost.indexOf('127.0.0.1') > -1 || xHost.indexOf('localhost') > -1) {
    xData.appEnv = 'dev';
    xData.appHost = xHost;
    xData.appWsHost = '127.0.0.1:20202';
  }

  if (xData.appEnv === 'dev') {
    window.console.log(xData);
  }

  return xData;
}

export function setLoginTenantInfo(tenantInfo?: API.TenantLoginRespData) {
  SetLocalCache('x_tenant_info', JSON.stringify(tenantInfo));
}

export function getLoginTenantInfo(): undefined | API.TenantLoginRespData {
  let xTenantInfo: undefined | API.TenantLoginRespData = undefined;

  const tenantInfoJson = GetLocalCache('x_tenant_info');
  if (!tenantInfoJson) {
    return xTenantInfo;
  }

  xTenantInfo = JSON.parse(tenantInfoJson);

  if (xTenantInfo) {
    xTenantInfo.env_data = getEnvData();
  }

  return xTenantInfo;
}

export function clearLoginTenantInfo() {
   ClearLocalCache();
}

export function getLoginTokenId(){

  const xTenantUserInfo = getLoginTenantInfo();

  if(xTenantUserInfo){
      return xTenantUserInfo.token_id;
  }

  return "";
}


export async function CtlCheckCodeImage() {
  return request<API.CheckCodeImageResponse>('/xapi/cms/check-code-image', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: null,
  });
}

export async function GetCheckCodeImage(): Promise<string> {
  
  let xData="";

  const xResp=await CtlCheckCodeImage();

  if(xResp&&xResp.errorCode&&xResp.errorCode=="0"){

    if(xResp.data?.code_image_data&&xResp.data?.code_image_id){
      xData=xResp.data?.code_image_data;
      SetLocalCache("x_check_code_image_id",xResp.data?.code_image_id);
    }
    
  }

  return xData;
}

export function GetCheckCodeImageId(): string{
  return GetLocalCache('x_check_code_image_id');
}

export async function CtlUserLogin(reqData: API.TenantLoginRequest) {
  return request<API.TenantLoginResponse>('/xapi/cms/tenant-login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}

export async function CtlTenantCurrent() {
  
  return request<API.TenantInfoResponse>('/xapi/cms/tenant-current', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: {},
  });
  
}

export async function CtlTenantUpdateInfo(reqData: API.TenantInfoRequest) {
  return request<API.BaseResponse>('/xapi/cms/tenant-update-info', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}


export async function CtlTenantUpdatePassword(reqData: API.TenantInfoRequest) {
  return request<API.BaseResponse>('/xapi/cms/tenant-update-password', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: reqData,
  });
}