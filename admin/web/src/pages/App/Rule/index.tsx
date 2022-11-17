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
import { CtlAppGatewayRulePage, CtlAppGatewayRuleRemove, SearchAppGatewayConfigMapDataWithMode } from '@/services/ant-design-pro/api';
import AddAppGatewayRuleDataForm from './components/AddAppGatewayRuleDataForm';
import { Button, Modal } from 'antd';
import EditAppGatewayRuleDataForm from './components/EditAppGatewayRuleDataForm';
import { ShowErrorMessage, ShowSuccessMessage } from '@/services/ant-design-pro/lib';

// 数据过滤处理
const processRespData = (dataList: API.AppGatewayRuleData[]) => {
  return dataList;
};

const RulePage: React.FC = () => {
  const [addModalVisible, setAddModalVisible] = useState<boolean>(false);
  const [editModalVisible, setEditModalVisible] = useState<boolean>(false);

  const actionRef = useRef<ActionType>();

  const [currentRow, setCurrentRow] = useState<API.AppGatewayRuleData>();

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

    const respData = await CtlAppGatewayRuleRemove(reqData);

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
  const columns: ProColumns<API.AppGatewayRuleData>[] = [
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
      title: '规则站点',
      dataIndex: 'site_data_name',
      ellipsis: true,
      hideInSearch: true,
    },
    {
      title: '规则序号',
      dataIndex: 'rule_order_no',
      sorter: true,
      hideInSearch: true,
      align: 'center',
    },
    {
      title: '规则名称',
      dataIndex: 'rule_name',
      sorter: false,
    },
    {
      title: '规则匹配对象',
      dataIndex: 'match_target',
      sorter: false,
      hideInSearch: true,
      valueType: 'select',
      valueEnum: {
        s_biz_merchant_id: { text: '业务租户ID', status: 's_biz_merchant_id' },
        s_http_header_useragent: { text: '客户端请求UA', status: 's_http_header_useragent' },    
        s_http_header_url: { text: '客户端请求URL', status: 's_http_header_url' },
        s_http_header_referer: { text: '客户端请求来源URL', status: 's_http_header_referer' },
        s_http_header_method: { text: '客户端请求方法', status: 's_http_header_method' },
        s_http_ip: { text: '客户端IP', status: 's_http_ip' },
      },
    },
    {
      title: '规则匹配操作',
      dataIndex: 'match_op',
      sorter: false,
      hideInSearch: true,
      valueType: 'select',
      valueEnum: {
        regex: { text: '正则匹配', status: 'regex' },
        contain: { text: '字符包含', status: 'contain' },
        in: { text: '包含于列表', status: 'in' },
        notregex: { text: '正则不匹配', status: 'notregex' },
        notcontain: { text: '字符不包含', status: 'notcontain' },
        notin: { text: '包不含于列表', status: 'notin' },
        eq: { text: '等于', status: 'eq' },
        neq: { text: '不等于', status: 'neq' },
        lt: { text: '小于', status: 'lt' },
        gt: { text: '大于', status: 'gt' },
        lte: { text: '小于等于', status: 'lte' },
        gte: { text: '大于等于', status: 'gte' },
      },
    },
    {
      title: '规则匹配内容',
      dataIndex: 'match_data',
      sorter: false,
      hideInSearch: true,
      ellipsis: true,
    },
    {
      title: '规则后端节点',
      dataIndex: 'action_data_name',
      sorter: false,
      hideInSearch: true,
      ellipsis: true,
    },
    {
      title: '规则状态',
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
      <ProTable<API.AppGatewayRuleData, API.AppGatewayRulePageRequest>
        bordered={true}
        columns={columns}
        actionRef={actionRef}
        rowKey="data_id"
        search={{
          filterType: 'light',
        }}
        request={CtlAppGatewayRulePage}
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
            添加规则
          </Button>,
        ]}
      />

      <AddAppGatewayRuleDataForm
        onModalVisible={onAddModalVisible}
        modalVisible={addModalVisible}
      />

      <EditAppGatewayRuleDataForm
        onModalVisible={onEditModalVisible}
        modalVisible={editModalVisible}
        dataId={currentRow?.data_id}
      />
    </PageContainer>
  );
};

export default RulePage;
