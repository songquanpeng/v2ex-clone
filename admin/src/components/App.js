import React from 'react';
import {
  Layout,
  Menu,
  Avatar,
  Dropdown,
  Space,
  Row,
  Col,
  Button,
  Upload,
  Popconfirm,
  Divider,
  message,
} from 'antd';

import {
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  UserOutlined,
  EditOutlined,
  DashboardOutlined,
  FileTextOutlined,
  CommentOutlined,
  CloudUploadOutlined,
  SettingOutlined,
  LogoutOutlined,
} from '@ant-design/icons';

import './App.css';

const { Header, Sider, Content } = Layout;
const { SubMenu } = Menu;

class App extends React.Component {
  state = {
    collapsed: false,
  };

  toggle = () => {
    this.setState({
      collapsed: !this.state.collapsed,
    });
  };

  render() {
    const menu = (
      <Menu>
        <Menu.Item key="1" icon={<UserOutlined />}>
          My Account
        </Menu.Item>
        <Menu.Item key="2" icon={<SettingOutlined />}>
          My Setting
        </Menu.Item>
        <Menu.Divider />
        <Menu.Item key="3" icon={<LogoutOutlined />}>
          Logout
        </Menu.Item>
      </Menu>
    );

    return (
      <Layout style={{ height: '100%' }}>
        <Sider trigger={null} collapsible collapsed={this.state.collapsed}>
          <div className="logo">
            <h1>System Admin</h1>
          </div>
          <Menu theme="dark" mode="inline" defaultSelectedKeys={['0']}>
            <Menu.Item key="0" icon={<DashboardOutlined />}>
              Dashboard
            </Menu.Item>
            <SubMenu key="sub0" icon={<EditOutlined />} title="Editor">
              <Menu.Item key="1">Rich Text Editor</Menu.Item>
              <Menu.Item key="2">Code Editor</Menu.Item>
            </SubMenu>
            <Menu.Item key="3" icon={<FileTextOutlined />}>
              Posts
            </Menu.Item>
            <Menu.Item key="4" icon={<CommentOutlined />}>
              Comments
            </Menu.Item>
            <Menu.Item key="5" icon={<CloudUploadOutlined />}>
              Files
            </Menu.Item>
            <Menu.Item key="6" icon={<UserOutlined />}>
              Users
            </Menu.Item>
            <Menu.Item key="7" icon={<SettingOutlined />}>
              Settings
            </Menu.Item>
          </Menu>
        </Sider>
        <Layout className="site-layout">
          <Header className="site-layout-background" style={{ padding: 0 }}>
            {React.createElement(
              this.state.collapsed ? MenuUnfoldOutlined : MenuFoldOutlined,
              {
                className: 'trigger',
                onClick: this.toggle,
              }
            )}
            <Space style={{ float: 'right', marginRight: '16px' }}>
              <Dropdown overlay={menu} placement="bottomCenter">
                <Button type={'text'} icon={<UserOutlined />} size={'large'}>
                  Admin
                </Button>
              </Dropdown>
            </Space>
          </Header>
          <Content
            className="site-layout-background"
            style={{
              margin: '24px 16px',
              padding: 24,
              minHeight: 280,
            }}
          >
            Content
          </Content>
        </Layout>
      </Layout>
    );
  }
}
export default App;
