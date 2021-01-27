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
  MenuOutlined,
  PlusCircleOutlined,
  PlusOutlined,
  EditOutlined,
  DeleteOutlined
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
  MenuOutlined,
  PlusCircleOutlined,
  PlusOutlined,
  EditOutlined,
  DeleteOutlined
]

const antdIcon = {
  install (Vue: any) {
    iconComponents.forEach(iconComponent => {
      Vue.component(iconComponent.name, iconComponent)
    })
  }
}

export default antdIcon
