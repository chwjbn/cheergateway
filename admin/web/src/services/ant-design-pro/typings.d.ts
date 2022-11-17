declare namespace API {
  type XEnvData = {
    appHost: string;
    appWsHost: string;
    appEnv: string;
  };

  type SelectItemNode = {
    label: string;
    value: any;
  };

  type DataMapNode = {
    data_id: string;
    data_name: string;
  };

  type BaseResponse = {
    success?: boolean;
    errorCode?: string;
    errorMessage?: string;
  };

  type BasePageDataResponse = {
    total?: number;
    success?: boolean;
  };

  type PageParams = {
    current?: number;
    pageSize?: number;
  };

  type CheckCodeImageRespData = {
    code_image_id?: string;
    code_image_data?: string;
  };

  type CheckCodeImageResponse = BaseResponse & {
    data?: CheckCodeImageRespData;
  };

  type TenantInfoData={
    tenant_id?: string;
    user_name?: string;
    password?: string;
    email?: string;
    user_img_url?: string;
    create_time?: string;
    create_ip?: string;
    status?: string;
  };

  type TenantInfoRequest={
    data?: TenantInfoData;
  };

  type TenantInfoResponse = BaseResponse & {
    data?: TenantInfoData;
  };

  type TenantLoginRequest = {
    user_name?: string;
    password?: string;
    check_code_img_id?: string;
    check_code_img_data?: string;
  };

  type TenantLoginRespData = {
    token_id?: string;
    user_name?: string;
    user_id?: string;
    user_img_url?: string;
    role?: string;
    env_data?: XEnvData;
  };

  type TenantLoginResponse = BaseResponse & {
    data?: TenantLoginRespData;
  };

  type AppDataInfoSearchRequest = {
    keyword?: string;
  };

  type AppData = {
    data_id?: string;
    create_time?: string;
    update_time?: string;
    status?: string;
  };

  type AppDataInfoRequest = {
    data_id?: string;
  };

  // gateway config
  type AppGatewayConfigData = AppData & {
    env_type?: string;
    config_name?: string;
    server_addr?: string;
    user_name?: string;
    password?: string;
    last_pub_ver?: string;
    last_ver?: string;
  };

  type AppGatewayConfigPageRequest = PageParams & {
    env_type?: string;
    config_name?: string;
    status?: string;
  };

  type AppGatewayConfigPageResponse = BasePageDataResponse & {
    data?: AppGatewayConfigData[];
  };

  type AppGatewayConfigAddRequest = {
    data?: AppGatewayConfigData;
  };

  type AppGatewayConfigGetResponse = BaseResponse & {
    data?: AppGatewayConfigData;
  };

  type AppGatewayConfigSaveRequest = {
    data?: AppGatewayConfigData;
  };

  type AppGatewayConfigMapDataRequest = {
    env_type?: string;
    mode?: string;
  };

  type AppGatewayConfigMapDataResponse = BaseResponse & {
    data?: DataMapNode[];
  };

  // gateway backend
  type AppGatewayBackendData = AppData & {
    config_data_id?: string;
    config_data_name?: string;
    backend_name?: string;
    node_addr?: string;
  };

  type AppGatewayBackendPageRequest = PageParams & {
    config_data_id?: string;
    backend_name?: string;
    node_addr?: string;
    status?: string;
  };

  type AppGatewayBackendPageResponse = BasePageDataResponse & {
    data?: AppGatewayBackendData[];
  };

  type AppGatewayBackendAddRequest = {
    data?: AppGatewayBackendData;
  };

  type AppGatewayBackendGetResponse = BaseResponse & {
    data?: AppGatewayBackendData;
  };

  type AppGatewayBackendSaveRequest = {
    data?: AppGatewayBackendData;
  };

  type AppGatewayBackendMapDataRequest = {
    config_data_id?: string;
  };

  type AppGatewayBackendMapDataResponse = BaseResponse & {
    data?: DataMapNode[];
  };

  // gateway site
  type AppGatewaySiteData = AppData & {
    config_data_id?: string;
    config_data_name?: string;
    site_order_no?: number;
    site_name?: string;
    rule_type?: string;
    rule_data?: string;
    default_backend_data_id?: string;
  };

  type AppGatewaySitePageRequest = PageParams & {
    config_data_id?: string;
    site_name?: string;
    status?: string;
  };

  type AppGatewaySitePageResponse = BasePageDataResponse & {
    data?: AppGatewaySiteData[];
  };

  type AppGatewaySiteAddRequest = {
    data?: AppGatewaySiteData;
  };

  type AppGatewaySiteGetResponse = BaseResponse & {
    data?: AppGatewaySiteData;
  };

  type AppGatewaySiteSaveRequest = {
    data?: AppGatewaySiteData;
  };

  type AppGatewaySiteMapDataRequest = {
    config_data_id?: string;
  };

  type AppGatewaySiteMapDataResponse = BaseResponse & {
    data?: DataMapNode[];
  };

  //gateway rule
  type AppGatewayRuleData = AppData & {
    config_data_id?: string;
    config_data_name?: string;
    site_data_id?: string;
    site_data_name?: string;
    rule_order_no?: number;
    rule_name?: string;
    match_target?: string;
    match_op?: string;
    match_data?: string;
    action_type?: string;
    action_data?: string;
    action_data_name?: string;
  };

  type AppGatewayRulePageRequest = PageParams & {
    config_data_id?: string;
    rule_name?: string;
    status?: string;
  };

  type AppGatewayRulePageResponse = BasePageDataResponse & {
    data?: AppGatewayRuleData[];
  };

  type AppGatewayRuleAddRequest = {
    data?: AppGatewayRuleData;
  };

  type AppGatewayRuleGetResponse = BaseResponse & {
    data?: AppGatewayRuleData;
  };

  type AppGatewayRuleSaveRequest = {
    data?: AppGatewayRuleData;
  };
}
