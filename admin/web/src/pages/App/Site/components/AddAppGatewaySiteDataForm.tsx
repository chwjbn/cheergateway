import React, { useRef, useState } from 'react';
import type { ProFormInstance } from '@ant-design/pro-form';
import { ProFormDigit } from '@ant-design/pro-form';
import { ModalForm, ProFormSelect, ProFormText } from '@ant-design/pro-form';
import moment from 'moment';
import 'moment/locale/zh-cn';
import {
  CtlAppGatewaySiteAdd,
  SearchAppGatewayBackendMapData,
  SearchAppGatewayConfigMapData,
} from '@/services/ant-design-pro/api';
import { ShowErrorMessage, ShowSuccessMessage } from '@/services/ant-design-pro/lib';

export type AddAppGatewaySiteDataFormProps = {
  onModalVisible: (flag: boolean) => void;
  modalVisible: boolean;
};

const AddAppGatewaySiteDataForm: React.FC<AddAppGatewaySiteDataFormProps> = (props) => {
  const xFormRef = useRef<ProFormInstance<API.AppGatewaySiteData>>();

  const [xBackendMapData, setBackendMapData] = useState<any>();

  const onFormVisible = (flag: boolean) => {
    moment.locale('zh-cn');
    return props.onModalVisible(flag);
  };

  const loadBackendMapData = async () => {
    const xData = xFormRef.current?.getFieldsValue();
    const mapData = await SearchAppGatewayBackendMapData(xData?.config_data_id);
    setBackendMapData(mapData);
  };

  const onSaveData = async (xData: API.AppGatewaySiteData) => {
    const xRequest: API.AppGatewaySiteAddRequest = {
      data: xData,
    };

    const xResp = await CtlAppGatewaySiteAdd(xRequest);

    if (xResp.errorCode === '0') {
      ShowSuccessMessage(xResp.errorMessage);
      return true;
    }

    ShowErrorMessage(xResp.errorMessage);

    return false;
  };

  return (
    <ModalForm<API.AppGatewaySiteData>
      title={'添加站点'}
      width="80%"
      onVisibleChange={onFormVisible}
      visible={props.modalVisible}
      modalProps={{ destroyOnClose: true }}
      onFinish={async (values) => {
        return onSaveData(values as API.AppGatewaySiteData);
      }}
      formRef={xFormRef}
    >
      <ProFormSelect
        name="config_data_id"
        label="配置服务"
        request={SearchAppGatewayConfigMapData}
        placeholder="请选择配置服务"
        rules={[{ required: true, message: '请选择配置服务.' }]}
        fieldProps={{  showSearch: true,onChange: loadBackendMapData }}
      />

      <ProFormDigit
        label="站点序号"
        name="site_order_no"
        min={1}
        max={1000}
        fieldProps={{ precision: 0 }}
      />

      <ProFormText
        name="site_name"
        label="站点名称"
        placeholder="请输入站点名称"
        rules={[{ required: true, message: '请输入站点名称.' }]}
      />

      <ProFormSelect
        name="rule_type"
        label="域名匹配方式"
        valueEnum={{
          regex: '正则匹配',
          contain: '字符包含',
          in: '包含于列表',
          notregex: '正则不匹配',
          notcontain: '字符不包含',
          notin: '不包含于列表',
          eq: '等于',
          neq: '不等于',
        }}
        placeholder="请选择域名匹配方式"
        rules={[{ required: true, message: '请选择域名匹配方式.' }]}
      />

      <ProFormText
        name="rule_data"
        label="域名匹配规则"
        placeholder="请输入域名匹配规则"
        rules={[{ required: true, message: '请输入域名匹配规则.' }]}
      />

      <ProFormSelect
        name="default_backend_data_id"
        label="默认后端节点"
        placeholder="请选择默认后端节点"
        rules={[{ required: true, message: '请选择默认后端节点.' }]}
        fieldProps={{ showSearch: true,options: xBackendMapData }}
      />
    </ModalForm>
  );
};

export default AddAppGatewaySiteDataForm;
