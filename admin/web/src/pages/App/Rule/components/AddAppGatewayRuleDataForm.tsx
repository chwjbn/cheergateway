import React, { useRef, useState } from 'react';
import type { ProFormInstance } from '@ant-design/pro-form';
import { ProFormTextArea } from '@ant-design/pro-form';
import { ProFormDigit } from '@ant-design/pro-form';
import { ModalForm, ProFormSelect, ProFormText } from '@ant-design/pro-form';
import moment from 'moment';
import 'moment/locale/zh-cn';
import {
  CtlAppGatewayRuleAdd,
  SearchAppGatewayBackendMapData,
  SearchAppGatewayConfigMapData,
  SearchAppGatewaySiteMapData,
} from '@/services/ant-design-pro/api';
import { ShowErrorMessage, ShowSuccessMessage } from '@/services/ant-design-pro/lib';

export type AddAppGatewayRuleDataFormProps = {
  onModalVisible: (flag: boolean) => void;
  modalVisible: boolean;
};

const AddAppGatewayRuleDataForm: React.FC<AddAppGatewayRuleDataFormProps> = (props) => {
  const xFormRef = useRef<ProFormInstance<API.AppGatewayRuleData>>();

  const [xBackendMapData, setBackendMapData] = useState<any>();
  const [xSiteMapData, setSiteMapData] = useState<any>();

  const onFormVisible = (flag: boolean) => {
    moment.locale('zh-cn');
    return props.onModalVisible(flag);
  };

  const loadBackendMapData = async () => {
    const xData = xFormRef.current?.getFieldsValue();
    const mapData = await SearchAppGatewayBackendMapData(xData?.config_data_id);
    setBackendMapData(mapData);
  };

  const loadSiteMapData = async () => {
    const xData = xFormRef.current?.getFieldsValue();
    const mapData = await SearchAppGatewaySiteMapData(xData?.config_data_id);
    setSiteMapData(mapData);
  };

  const loadMapData = async () => {
    await loadSiteMapData();
    await loadBackendMapData();
  };

  const onSaveData = async (xData: API.AppGatewayRuleData) => {
    const xRequest: API.AppGatewayRuleAddRequest = {
      data: xData,
    };

    const xResp = await CtlAppGatewayRuleAdd(xRequest);

    if (xResp.errorCode === '0') {
      ShowSuccessMessage(xResp.errorMessage);
      return true;
    }

    ShowErrorMessage(xResp.errorMessage);

    return false;
  };

  return (
    <ModalForm<API.AppGatewayRuleData>
      title={'添加规则'}
      width="80%"
      onVisibleChange={onFormVisible}
      visible={props.modalVisible}
      modalProps={{ destroyOnClose: true }}
      onFinish={async (values) => {
        return onSaveData(values as API.AppGatewayRuleData);
      }}
      formRef={xFormRef}
    >
      <ProFormSelect
        name="config_data_id"
        label="配置服务"
        request={SearchAppGatewayConfigMapData}
        placeholder="请选择配置服务"
        rules={[{ required: true, message: '请选择配置服务.' }]}
        fieldProps={{  showSearch: true,onChange: loadMapData }}
      />

      <ProFormText
        name="rule_name"
        label="规则名称"
        placeholder="请输入规则名称"
        rules={[{ required: true, message: '请输入规则名称.' }]}
      />

      <ProFormSelect
        name="site_data_id"
        label="规则站点"
        placeholder="请选择规则站点"
        rules={[{ required: true, message: '请选择规则站点.' }]}
        fieldProps={{  showSearch: true,options: xSiteMapData }}
      />

      <ProFormDigit
        label="规则序号"
        name="rule_order_no"
        min={1}
        max={1000}
        fieldProps={{ precision: 0 }}
        rules={[{ required: true, message: '请输入规则序号.' }]}
      />

      <ProFormSelect
        name="match_target"
        label="规则匹配对象"
        valueEnum={{
          s_biz_merchant_id: { text: '业务租户ID', status: 's_biz_merchant_id' },
          s_http_header_useragent: { text: '客户端请求UA', status: 's_http_header_useragent' },    
          s_http_header_url: { text: '客户端请求URL', status: 's_http_header_url' },
          s_http_header_referer: { text: '客户端请求来源URL', status: 's_http_header_referer' },
          s_http_header_method: { text: '客户端请求方法', status: 's_http_header_method' },
          s_http_ip: { text: '客户端IP', status: 's_http_ip' },
        }}
        placeholder="请选择规则匹配对象"
        rules={[{ required: true, message: '请选择规则匹配对象.' }]}
      />

      <ProFormSelect
        name="match_op"
        label="规则匹配操作"
        valueEnum={{
          regex: '正则匹配',
          contain: '字符包含',
          in: '包含于列表',
          notregex: '正则不匹配',
          notcontain: '字符不包含',
          notin: '不包含于列表',
          eq: '等于',
          neq: '不等于',
          lt: '小于',
          gt: '大于',
          lte: '小于等于',
          gte: '大于等于',
        }}
        placeholder="请选择规则匹配操作"
        rules={[{ required: true, message: '请选择规则匹配操作.' }]}
      />

      <ProFormTextArea
        name="match_data"
        label="规则匹配内容"
        placeholder="请输入规则匹配内容"
        rules={[{ required: true, message: '请输入规则匹配内容.' }]}
      />

      <ProFormSelect
        name="action_data"
        label="规则后端节点"
        placeholder="请选择规则后端节点"
        rules={[{ required: true, message: '请选择规则后端节点.' }]}
        fieldProps={{  showSearch: true,options: xBackendMapData }}
      />
    </ModalForm>
  );
};

export default AddAppGatewayRuleDataForm;