import React, { useEffect, useRef, useState } from 'react';
import type { ProFormInstance } from '@ant-design/pro-form';
import ProForm, { ProFormText } from '@ant-design/pro-form';
import 'moment/locale/zh-cn';
import { ShowErrorMessage, ShowSuccessMessage } from '@/services/ant-design-pro/lib';
import { Spin } from 'antd';
import { CtlTenantCurrent, CtlTenantUpdatePassword } from '@/services/ant-design-pro/tenant';

const SecurityView: React.FC = () => {
  const [xIsLoading, setIsLoading] = useState<boolean>(true);

  const xFormRef = useRef<ProFormInstance<API.TenantInfoData>>();

  const loadDataInfo = async () => {
    setIsLoading(true);

    const respData = await CtlTenantCurrent();
    if (respData && respData.errorCode === '0') {
      if (respData.data) {
        const xDataInfo = respData.data;
        xFormRef.current?.setFieldsValue(xDataInfo);
      }
    }

    setIsLoading(false);
  };

  // 加载
  useEffect(() => {
    loadDataInfo();
  }, []);

  const onSaveData = async (xData: API.TenantInfoData) => {
    
    const xRequest: API.TenantInfoRequest = {
      data: xData,
    };

    const xResp = await CtlTenantUpdatePassword(xRequest);

    if (xResp.errorCode === '0') {
      ShowSuccessMessage(xResp.errorMessage);
      return true;
    }

    ShowErrorMessage(xResp.errorMessage);

    return false;
  };

  return (
    <ProForm<API.TenantInfoData>
      onFinish={async (values) => {
        return onSaveData(values as API.TenantInfoData);
      }}
      submitter={{
        resetButtonProps: {
          style: {
            display: 'none',
          },
        },
      }}
      formRef={xFormRef}
    >
      <Spin spinning={xIsLoading} delay={100}>

      <ProFormText
          name="tenant_id"
          label="用户ID"
          placeholder="请输入用户ID"
          disabled
          rules={[{ required: true, message: '请输入用户ID.' }]}
        />

      <ProFormText
          name="user_name"
          label="用户名"
          placeholder="请输入用户名"
          disabled
          rules={[{ required: true, message: '请输入用户名.' }]}
        />
        
        <ProFormText.Password
          name="password"
          label="登录密码"
          placeholder="请输入登录密码"
          rules={[{ required: true, message: '请输入登录密码.' }]}
        />

       
      </Spin>
    </ProForm>
  );
};

export default SecurityView;
