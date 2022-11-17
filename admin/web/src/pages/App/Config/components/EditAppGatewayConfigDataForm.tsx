import React, { useEffect, useRef, useState } from 'react';
import type { ProFormInstance } from '@ant-design/pro-form';
import {
  ModalForm,
  ProFormSelect,
  ProFormText,
} from '@ant-design/pro-form';
import moment from 'moment';
import 'moment/locale/zh-cn';
import { CtlAppGatewayConfigSave, CtlAppGatewayConfigGet } from '@/services/ant-design-pro/api';
import { ShowErrorMessage, ShowSuccessMessage } from '@/services/ant-design-pro/lib';
import { Spin } from 'antd';

export type EditAppGatewayConfigDataFormProps = {
  onModalVisible: (flag: boolean) => void;
  modalVisible: boolean;
  dataId?: string;
};

const EditAppGatewayConfigDataForm: React.FC<EditAppGatewayConfigDataFormProps> = (props) => {
  const [xIsLoading, setIsLoading] = useState<boolean>(true);

  const xFormRef = useRef<ProFormInstance<API.AppGatewayConfigData>>();

  const onFormVisible = (flag: boolean) => {
    moment.locale('zh-cn');
    return props.onModalVisible(flag);
  };

  const loadDataInfo = async (dataId?: string) => {

    setIsLoading(true);

    if (dataId) {
      const respData = await CtlAppGatewayConfigGet({ data_id: dataId });
      if (respData && respData.errorCode === '0') {
        if (respData.data) {
          const xDataInfo = respData.data;
          xFormRef.current?.setFieldsValue(xDataInfo);
        }
      }
    }

    setIsLoading(false);
  };

  // 加载
  useEffect(() => {
    if (props.modalVisible) {
      loadDataInfo(props.dataId);
    }

    // eslint-disable-next-line react-hooks/exhaustive-deps
  }, [props.dataId, props.modalVisible]);

  const onSaveData = async (xData: API.AppGatewayConfigData) => {

    xData.data_id=props.dataId;

    const xRequest: API.AppGatewayConfigSaveRequest = {
      data: xData,
    };

    const xResp = await CtlAppGatewayConfigSave(xRequest);

    if (xResp.errorCode === '0') {
      ShowSuccessMessage(xResp.errorMessage);
      return true;
    }

    ShowErrorMessage(xResp.errorMessage);

    return false;
  };

  return (
    <ModalForm<API.AppGatewayConfigData>
      title={'修改站点'}
      width="80%"
      onVisibleChange={onFormVisible}
      visible={props.modalVisible}
      modalProps={{ destroyOnClose: true }}
      onFinish={async (values) => {
        return onSaveData(values as API.AppGatewayConfigData);
      }}
      formRef={xFormRef}
    >
      <Spin spinning={xIsLoading} delay={100}>
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


     <ProFormSelect
        name="status"
        label="数据状态"
        valueEnum={{
          enable: '启用',
          disable: '禁用',
        }}
        placeholder="请选择数据状态"
        rules={[{ required: true, message: '请选择数据状态.' }]}
      />

      </Spin>
    </ModalForm>
  );
};

export default EditAppGatewayConfigDataForm;
