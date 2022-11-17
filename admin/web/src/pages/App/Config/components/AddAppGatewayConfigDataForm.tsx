import React from 'react';
import {
  ModalForm,
  ProFormSelect,
  ProFormText,
} from '@ant-design/pro-form';
import moment from 'moment';
import 'moment/locale/zh-cn';
import { CtlAppGatewayConfigAdd } from '@/services/ant-design-pro/api';
import { ShowErrorMessage, ShowSuccessMessage } from '@/services/ant-design-pro/lib';

export type AddAppGatewayConfigDataFormProps = {
  onModalVisible: (flag: boolean) => void;
  modalVisible: boolean;
};

const AddAppGatewayConfigDataForm: React.FC<AddAppGatewayConfigDataFormProps> = (props) => {
  const onFormVisible = (flag: boolean) => {
    moment.locale('zh-cn');
    return props.onModalVisible(flag);
  };

  const onSaveData = async (xData: API.AppGatewayConfigData) => {
    const xRequest: API.AppGatewayConfigAddRequest = {
      data: xData,
    };

    const xResp = await CtlAppGatewayConfigAdd(xRequest);

    if (xResp.errorCode === '0') {
      ShowSuccessMessage(xResp.errorMessage);
      return true;
    }

    ShowErrorMessage(xResp.errorMessage);

    return false;
  };

  return (
    <ModalForm<API.AppGatewayConfigData>
      title={'添加配置'}
      width="80%"
      onVisibleChange={onFormVisible}
      visible={props.modalVisible}
      modalProps={{ destroyOnClose: true }}
      onFinish={async (values) => {
        return onSaveData(values as API.AppGatewayConfigData);
      }}
    >
      <ProFormSelect
        name="env_type"
        label="环境类别"
        valueEnum={{
          prod: '生产',
          test: '测试',
          dev: '开发',
        }}
        placeholder="请选择环境类别"
        rules={[{ required: true, message: '请选择环境类别.' }]}
      />

      <ProFormText
        name="config_name"
        label="配置名称"
        placeholder="请输入配置名称"
        rules={[{ required: true, message: '请输入配置名称.' }]}
      />

      <ProFormText
        name="server_addr"
        label="服务器地址"
        placeholder="请输入Redis服务器地址(地址:端口),如:127.0.0.1:6379"
        rules={[{ required: true, message: '请输入Redis服务器地址(地址:端口),如:127.0.0.1:6379.' }]}
      />

      <ProFormText
        name="user_name"
        label="服务器用户名"
        placeholder="请输入Redis服务器用户名"
        rules={[{ required: false, message: '请输入Redis服务器用户名.' }]}
      />

      <ProFormText
        name="password"
        label="服务器密码"
        placeholder="请输入Redis服务器密码"
        rules={[{ required: false, message: '请输入Redis服务器密码.' }]}
      />
    </ModalForm>
  );
};

export default AddAppGatewayConfigDataForm;
