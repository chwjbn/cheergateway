export default [
  {
    path: '/user',
    routes: [
      {
        path: '/user',
        routes: [
          {
            name: 'login',
            path: '/user/login',
            layout: false,
            component: './User/Login',
          },
          {
            name: 'settings',
            path: '/user/settings',
            component: './User/Setting',
          },
        ],
      },
      {
        component: './404',
      },
    ],
  },
  {
    name: 'config',
    icon: 'database',
    path: '/app/config',
    component: '@/pages/App/Config',
  },
  {
    name: 'backend',
    icon: 'cluster',
    path: '/app/backend',
    component: '@/pages/App/Backend',
  },
  {
    name: 'site',
    icon: 'deployment-unit',
    path: '/app/site',
    component: '@/pages/App/Site',
  },
  {
    name: 'rule',
    icon: 'control',
    path: '/app/rule',
    component: '@/pages/App/Rule',
  },
  {
    path: '/',
    redirect: '/app/site',
  },
  {
    component: './404',
  },
];
