import component from './zh-CN/component';
import globalHeader from './zh-CN/globalHeader';
import menu from './zh-CN/menu';
import pwa from './zh-CN/pwa';
import settingDrawer from './zh-CN/settingDrawer';
import settings from './zh-CN/settings';
import pages from './zh-CN/pages';

export default {
  'navBar.lang': '语言',
  'layout.user.link.help': '帮助',
  'layout.user.link.privacy': '隐私',
  'layout.user.link.terms': '条款',
  'app.copyright.produced': 'CheerGateway出品',
  ...pages,
  ...globalHeader,
  ...menu,
  ...settingDrawer,
  ...settings,
  ...pwa,
  ...component,

  'app.server.msg.common.server.error': '远程服务升级维护中，请稍候再试!',
  'app.server.msg.common.op.succ': '操作成功!',
  'app.server.msg.common.request.succ': '请求成功!',
  'app.server.msg.common.request.invalid': '你提交的请求数据不正确!',
  'app.server.msg.common.data.notexists': '你请求的数据不存在!',

  'app.server.msg.common.token.invalid': '账号尚未登录或者登录超时,请先登录!',

  'app.server.msg.tenant.login.succ': '账号登录成功!',

  'app.server.msg.tenant.username.invalid':'你输入的账号用户名无效!',
  'app.server.msg.tenant.password.invalid':'你输入的账号密码无效!',
  'app.server.msg.tenant.email.invalid':'你输入的账号电子邮箱无效!',
  'app.server.msg.tenant.checkimgcode.invalid':'你输入的图形验证码无效!',

  'app.server.msg.tenant.checkimgcode.error': '你输入的图形验证码不正确!',
  'app.server.msg.tenant.account.error':'你输入的账号用户名或者密码不正确!',
  'app.server.msg.tenant.account.status.noactive':'此账号未激活或者被禁用,请联系系统管理员!',


};
