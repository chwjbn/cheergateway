import React, { useState, useRef } from 'react';
import {
  DeleteOutlined,
  EditOutlined,
  ExclamationCircleOutlined,
  PlusOutlined,
  SendOutlined,
} from '@ant-design/icons';
import { PageContainer } from '@ant-design/pro-layout';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import ProTable from '@ant-design/pro-table';
import {
  CtlAppGatewayConfigPage,
  CtlAppGatewayConfigPub,
  CtlAppGatewayConfigRemove,
} from '@/services/ant-design-pro/api';
import AddSiteDataForm from './components/AddAppGatewayConfigDataForm';
import { Button, Modal, Tag } from 'antd';
import EditSiteDataForm from './components/EditAppGatewayConfigDataForm';
import { ShowErrorMessage, ShowSuccessMessage } from '@/services/ant-design-pro/lib';

// 数据过滤处理
const processRespData = (dataList: API.AppGatewayConfigData[]) => {
  return dataList;
};

const ConfigPage: React.FC = () => {
  const [addModalVisible, setAddModalVisible] = useState<boolean>(false);
  const [editModalVisible, setEditModalVisible] = useState<boolean>(false);

  const actionRef = useRef<ActionType>();

  const [currentRow, setCurrentRow] = useState<API.AppGatewayConfigData>();

  const onAddModalVisible = (flag: boolean) => {
    if (!flag) {
      actionRef.current?.reload();
    }
    setAddModalVisible(flag);
  };

  const onEditModalVisible = (flag: boolean) => {
    if (!flag) {
      actionRef.current?.reload();
    }
    setEditModalVisible(flag);
  };

  const doRemoveDataAction = async (dataId?: string) => {
    const reqData: API.AppDataInfoRequest = {};
    reqData.data_id = dataId;

    const respData = await CtlAppGatewayConfigRemove(reqData);

    if (respData.errorCode === '0') {
      ShowSuccessMessage(respData.errorMessage);
      actionRef.current?.reload();
      return;
    }

    ShowErrorMessage(respData.errorMessage);
  };

  const doRemoveData = async (dataId?: string) => {
    Modal.confirm({
      title: '系统操作确认',
      icon: <ExclamationCircleOutlined />,
      content: '此数据删除操作不可恢复,你确定要删除此数据吗?',
      okText: '确认',
      cancelText: '取消',
      onOk: async () => {
        await doRemoveDataAction(dataId);
      },
    });
  };

  const doPubDataAction = async (dataId?: string) => {
    const reqData: API.AppDataInfoRequest = {};
    reqData.data_id = dataId;

    const respData = await CtlAppGatewayConfigPub(reqData);

    if (respData.errorCode === '0') {
      ShowSuccessMessage(respData.errorMessage);
      actionRef.current?.reload();
      return;
    }

    ShowErrorMessage(respData.errorMessage);
  };

  const doPubData = async (dataId?: string) => {
    Modal.confirm({
      title: '系统操作确认',
      icon: <ExclamationCircleOutlined />,
      content: '此配置数据发布操作直接影响网关结果,你确定要发布此配置吗?',
      okText: '确认',
      cancelText: '取消',
      onOk: async () => {
        await doPubDataAction(dataId);
      },
    });
  };

  // 定义列
  const columns: ProColumns<API.AppGatewayConfigData>[] = [
    {
      title: '系统ID',
      dataIndex: 'data_id',
      hideInSearch: true,
      hideInTable: true,
    },

    {
      title: '配置名称',
      dataIndex: 'config_name',
      align:'center',
    },
    {
      title: '环境类别',
      dataIndex: 'env_type',
      valueType: 'select',
      valueEnum: {
        prod: { text: '生产', rule_type: 'prod' },
        test: { text: '测试', rule_type: 'test' },
        dev: { text: '开发', rule_type: 'dev' },
      },
    },
    {
      title: '服务器地址',
      dataIndex: 'server_addr',
      hideInSearch: true,
    },
    {
      title: '最后更新版本',
      dataIndex: 'last_ver',
      valueType: 'dateTime',
      hideInSearch: true,
    },
    {
      title: '最后发布版本',
      dataIndex: 'last_pub_ver',
      valueType: 'dateTime',
      sorter: false,
      hideInSearch: true,
      render: (_, record) => (
        <Tag color={record.last_pub_ver == record.last_ver ? 'green' : 'red'}>
          {record.last_pub_ver}
        </Tag>
      ),
    },
    {
      title: '配置状态',
      dataIndex: 'status',
      valueType: 'select',
      sorter: true,
      valueEnum: {
        enable: { text: '启用', status: 'enable' },
        disable: { text: '禁用', status: 'disable' },
      },
    },
    {
      title: '操作',
      key: 'option',
      valueType: 'option',
      render: (_, record) => [
        <a
          key="act-edit"
          onClick={() => {
            setCurrentRow(record);
            onEditModalVisible(true);
          }}
        >
          <EditOutlined />
          &nbsp;修改
        </a>,
        <a
          key="act-pub"
          onClick={() => {
            doPubData(record.data_id);
          }}
        >
          <SendOutlined />
          &nbsp;发布
        </a>,
        <a
          key="act-del"
          onClick={() => {
            doRemoveData(record.data_id);
          }}
        >
          <DeleteOutlined />
          &nbsp;删除
        </a>,
      ],
      width: '230px',
    },
  ];

  return (
    <PageContainer>
      <ProTable<API.AppGatewayConfigData, API.AppGatewayConfigPageRequest>
        bordered={true}
        columns={columns}
        actionRef={actionRef}
        rowKey="data_id"
        search={{
          filterType: 'light',
        }}
        request={CtlAppGatewayConfigPage}
        postData={processRespData}
        toolBarRender={() => [
          <Button
            type="primary"
            key="primary"
            onClick={() => {
              onAddModalVisible(true);
            }}
            icon={<PlusOutlined />}
          >
            添加配置
          </Button>,
        ]}
      />

      <AddSiteDataForm onModalVisible={onAddModalVisible} modalVisible={addModalVisible} />

      <EditSiteDataForm
        onModalVisible={onEditModalVisible}
        modalVisible={editModalVisible}
        dataId={currentRow?.data_id}
      />
    </PageContainer>
  );
};

export default ConfigPage;
