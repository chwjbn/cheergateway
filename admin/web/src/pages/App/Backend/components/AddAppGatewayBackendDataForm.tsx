import React from 'react';
import {
  ModalForm,
  ProFormSelect,
  ProFormText,
} from '@ant-design/pro-form';
import moment from 'moment';
import 'moment/locale/zh-cn';
import { CtlAppGatewayBackendAdd, SearchAppGatewayConfigMapData } from '@/services/ant-design-pro/api';
import { ShowErrorMessage, ShowSuccessMessage } from '@/services/ant-design-pro/lib';

export type AddAppGatewayBackendDataFormProps = {
  onModalVisible: (flag: boolean) => void;
  modalVisible: boolean;
};

const AddAppGatewayBackendDataForm: React.FC<AddAppGatewayBackendDataFormProps> = (props) => {
  const onFormVisible = (flag: boolean) => {
    moment.locale('zh-cn');
    return props.onModalVisible(flag);
  };

  const onSaveData = async (xData: API.AppGatewayBackendData) => {
    const xRequest: API.AppGatewayBackendAddRequest = {
      data: xData,
    };

    const xResp = await CtlAppGatewayBackendAdd(xRequest);

    if (xResp.errorCode === '0') {
      ShowSuccessMessage(xResp.errorMessage);
      return true;
    }

    ShowErrorMessage(xResp.errorMessage);

    return false;
  };

  return (
    <ModalForm<API.AppGatewayBackendData>
      title={'添加节点'}
      width="80%"
      onVisibleChange={onFormVisible}
      visible={props.modalVisible}
      modalProps={{ destroyOnClose: true }}
      onFinish={async (values) => {
        return onSaveData(values as API.AppGatewayBackendData);
      }}
    >
      <ProFormSelect
        name="config_data_id"
        label="配置服务"
        request={SearchAppGatewayConfigMapData}
        placeholder="请选择配置服务"
        rules={[{ required: true, message: '请选择配置服务.' }]}
      />

      <ProFormText
        name="backend_name"
        label="节点名称"
        placeholder="请输入节点名称"
        rules={[{ required: true, message: '请输入节点名称.' }]}
      />

      <ProFormText
        name="node_addr"
        label="服务器地址"
        placeholder="请输入节点地址(地址:端口,多个地址用|隔开),如:172.16.1.10:80|172.16.1.11:80"
        rules={[{ required: true, message: '请输入节点地址(地址:端口,多个地址用|隔开),如:172.16.1.10:80|172.16.1.11:80.' }]}
      />
    </ModalForm>
  );
};

export default AddAppGatewayBackendDataForm;
