import React, { useEffect, useRef, useState } from 'react';
import type { ProFormInstance } from '@ant-design/pro-form';
import {
  ModalForm,
  ProFormSelect,
  ProFormText,
} from '@ant-design/pro-form';
import moment from 'moment';
import 'moment/locale/zh-cn';
import { SearchAppGatewayConfigMapData, CtlAppGatewayBackendSave, CtlAppGatewayBackendGet } from '@/services/ant-design-pro/api';
import { ShowErrorMessage, ShowSuccessMessage } from '@/services/ant-design-pro/lib';
import { Spin } from 'antd';

export type EditAppGatewayBackendDataFormProps = {
  onModalVisible: (flag: boolean) => void;
  modalVisible: boolean;
  dataId?: string;
};

const EditAppGatewayBackendDataForm: React.FC<EditAppGatewayBackendDataFormProps> = (props) => {
  const [xIsLoading, setIsLoading] = useState<boolean>(true);

  const xFormRef = useRef<ProFormInstance<API.AppGatewayBackendData>>();

  const onFormVisible = (flag: boolean) => {
    moment.locale('zh-cn');
    return props.onModalVisible(flag);
  };

  const loadDataInfo = async (dataId?: string) => {

    setIsLoading(true);

    if (dataId) {
      const respData = await CtlAppGatewayBackendGet({ data_id: dataId });
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

  const onSaveData = async (xData: API.AppGatewayBackendData) => {

    xData.data_id=props.dataId;

    const xRequest: API.AppGatewayBackendSaveRequest = {
      data: xData,
    };

    const xResp = await CtlAppGatewayBackendSave(xRequest);

    if (xResp.errorCode === '0') {
      ShowSuccessMessage(xResp.errorMessage);
      return true;
    }

    ShowErrorMessage(xResp.errorMessage);

    return false;
  };

  return (
    <ModalForm<API.AppGatewayBackendData>
      title={'修改节点'}
      width="80%"
      onVisibleChange={onFormVisible}
      visible={props.modalVisible}
      modalProps={{ destroyOnClose: true }}
      onFinish={async (values) => {
        return onSaveData(values as API.AppGatewayBackendData);
      }}
      formRef={xFormRef}
    >
      <Spin spinning={xIsLoading} delay={100}>
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

export default EditAppGatewayBackendDataForm;
