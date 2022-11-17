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
import { CtlAppGatewaySitePage, CtlAppGatewaySiteRemove, SearchAppGatewayConfigMapDataWithMode } from '@/services/ant-design-pro/api';
import AddAppGatewaySiteDataForm from './components/AddAppGatewaySiteDataForm';
import { Button, Modal } from 'antd';
import EditAppGatewaySiteDataForm from './components/EditAppGatewaySiteDataForm';
import { ShowErrorMessage, ShowSuccessMessage } from '@/services/ant-design-pro/lib';

// 数据过滤处理
const processRespData = (dataList: API.AppGatewaySiteData[]) => {
  return dataList;
};

const SitePage: React.FC = () => {
  const [addModalVisible, setAddModalVisible] = useState<boolean>(false);
  const [editModalVisible, setEditModalVisible] = useState<boolean>(false);

  const actionRef = useRef<ActionType>();

  const [currentRow, setCurrentRow] = useState<API.AppGatewaySiteData>();

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

    const respData = await CtlAppGatewaySiteRemove(reqData);

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
  const columns: ProColumns<API.AppGatewaySiteData>[] = [
    {
      title: '站点ID',
      dataIndex: 'data_id',
      hideInSearch: true,
      hideInTable: true,
    },
    {
      title: '配置服务',
      dataIndex: 'config_data_id',
      sorter: true,
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
      title: '站点序号',
      dataIndex: 'site_order_no',
      sorter: true,
      hideInSearch: true,
      align: 'center',
    },
    {
      title: '站点名称',
      dataIndex: 'site_name',
      sorter: false,
    },
    {
      title: '域名匹配方式',
      dataIndex: 'rule_type',
      sorter: false,
      hideInSearch: true,
      valueEnum: {
        regex: { text: '正则匹配', status: 'regex' },
        contain: { text: '字符包含', status: 'contain' },
        in: { text: '包含于列表', status: 'in' },
        notregex: { text: '正则不匹配', status: 'notregex' },
        notcontain: { text: '字符不包含', status: 'notcontain' },
        notin: { text: '包不含于列表', status: 'notin' },
        eq: { text: '等于', status: 'eq' },
        neq: { text: '不等于', status: 'neq' },
      },
    },
    {
      title: '域名匹配规则',
      dataIndex: 'rule_data',
      sorter: false,
      hideInSearch: true,
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
      title: '站点状态',
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
      <ProTable<API.AppGatewaySiteData, API.AppGatewaySitePageRequest>
        bordered={true}
        columns={columns}
        actionRef={actionRef}
        rowKey="data_id"
        search={{
          filterType: 'light',
        }}
        request={CtlAppGatewaySitePage}
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
            添加站点
          </Button>,
        ]}
      />

      <AddAppGatewaySiteDataForm
        onModalVisible={onAddModalVisible}
        modalVisible={addModalVisible}
      />

      <EditAppGatewaySiteDataForm
        onModalVisible={onEditModalVisible}
        modalVisible={editModalVisible}
        dataId={currentRow?.data_id}
      />
    </PageContainer>
  );
};

export default SitePage;
