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
  DeleteOutlined,
  CopyOutlined,
  SaveOutlined,
  LaptopOutlined,
  SearchOutlined
} from '@ant-design/icons-vue'

export const iconComponents = [
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
  DeleteOutlined,
  CopyOutlined,
  SaveOutlined,
  LaptopOutlined,
  SearchOutlined
]

const antdIcon = {
  install (Vue: any) {
    iconComponents.forEach(iconComponent => {
      Vue.component(iconComponent.name, iconComponent)
    })
  }
}

export default antdIcon
