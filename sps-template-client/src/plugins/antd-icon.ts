import {
  CodeSandboxCircleFilled,
  UserOutlined,
  LockOutlined,
  EyeOutlined,
  EyeInvisibleOutlined,
  MailOutlined,
  LoginOutlined,
  HomeOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  DownloadOutlined,
  DownOutlined,
  LogoutOutlined,
  SettingOutlined,
  TeamOutlined,
  MenuOutlined
} from '@ant-design/icons-vue'

const iconComponents = [
  CodeSandboxCircleFilled,
  UserOutlined,
  LockOutlined,
  EyeOutlined,
  EyeInvisibleOutlined,
  MailOutlined,
  LoginOutlined,
  HomeOutlined,
  MenuFoldOutlined,
  MenuUnfoldOutlined,
  DownloadOutlined,
  DownOutlined,
  LogoutOutlined,
  SettingOutlined,
  TeamOutlined,
  MenuOutlined
]

const antdIcon = {
  install (Vue: any) {
    iconComponents.forEach(iconComponent => {
      Vue.component(iconComponent.name, iconComponent)
    })
  }
}

export default antdIcon
