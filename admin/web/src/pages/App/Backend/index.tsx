import React, { useState, useRef } from 'react';
import {
  DeleteOutlined,
  EditOutlined,
  ExclamationCircleOutlined,
  PlusOutlined,
} from '@ant-design/icons';
import { PageContainer } from '@ant-design/pro-layout';
import type { ProColumns, ActionType } from '@ant-design/pro-table';
import ProTable from '@ant-design/pro-table';
import {
  CtlAppGatewayBackendPage,
  CtlAppGatewayBackendRemove,
  SearchAppGatewayConfigMapDataWithMode,
} from '@/services/ant-design-pro/api';
import AddAppGatewayBackendDataForm from './components/AddAppGatewayBackendDataForm';
import { Button, Modal } from 'antd';
import EditAppGatewayBackendDataForm from './components/EditAppGatewayBackendDataForm';
import { ShowErrorMessage, ShowSuccessMessage } from '@/services/ant-design-pro/lib';

// 数据过滤处理
const processRespData = (dataList: API.AppGatewayBackendData[]) => {
  return dataList;
};

const BackendPage: React.FC = () => {
  const [addModalVisible, setAddModalVisible] = useState<boolean>(false);
  const [editModalVisible, setEditModalVisible] = useState<boolean>(false);

  const actionRef = useRef<ActionType>();

  const [currentRow, setCurrentRow] = useState<API.AppGatewayBackendData>();

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

    const respData = await CtlAppGatewayBackendRemove(reqData);

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

  // 定义列
  const columns: ProColumns<API.AppGatewayBackendData>[] = [
    {
      title: '节点ID',
      dataIndex: 'data_id',
      hideInSearch: true,
      hideInTable: true,
    },
    {
      title: '配置服务',
      dataIndex: 'config_data_id',
      hideInSearch: false,
      hideInTable: true,
      request:async () => {
        return SearchAppGatewayConfigMapDataWithMode('','simple');
      },
      valueType: 'select',
      fieldProps:{showSearch: true},
    },
    {
      title: '配置服务',
      dataIndex: 'config_data_name',
      align:'center',
      hideInSearch: true,
    },
    {
      title: '节点名称',
      dataIndex: 'backend_name',
      sorter: false,
      ellipsis: true,
    },
    {
      title: '节点地址',
      dataIndex: 'node_addr',
      sorter: false,
      ellipsis: true,
    },
    {
      title: '最后更新',
      dataIndex: 'last_sync_time',
      valueType: 'dateTime',
      sorter: true,
      hideInSearch: true,
    },
    {
      title: '节点状态',
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
          key="act-del"
          onClick={() => {
            doRemoveData(record.data_id);
          }}
        >
          <DeleteOutlined />
          &nbsp;删除
        </a>,
      ],
      width: '180px',
    },
  ];

  return (
    <PageContainer>
      <ProTable<API.AppGatewayBackendData, API.AppGatewayBackendPageRequest>
        bordered={true}
        columns={columns}
        actionRef={actionRef}
        rowKey="data_id"
        search={{
          filterType: 'light',
        }}
        request={CtlAppGatewayBackendPage}
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
            添加节点
          </Button>,
        ]}
      />

      <AddAppGatewayBackendDataForm
        onModalVisible={onAddModalVisible}
        modalVisible={addModalVisible}
      />

      <EditAppGatewayBackendDataForm
        onModalVisible={onEditModalVisible}
        modalVisible={editModalVisible}
        dataId={currentRow?.data_id}
      />
    </PageContainer>
  );
};

export default BackendPage;
