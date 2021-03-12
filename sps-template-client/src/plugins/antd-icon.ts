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
  SearchOutlined,
  BookOutlined,
  UploadOutlined,
  ApartmentOutlined,
  UndoOutlined,
  RedoOutlined,
  SnippetsOutlined,
  ZoomInOutlined,
  ZoomOutOutlined,
  ExpandOutlined,
  ScanOutlined,
  VerticalAlignTopOutlined,
  VerticalAlignBottomOutlined,
  InfoCircleOutlined,
  FileExcelOutlined,
  CloudUploadOutlined,
  ContactsOutlined,
  ToolOutlined,
  CodeOutlined,
  FormOutlined,
  FileOutlined,
  HistoryOutlined,
  CloudOutlined,
  UnorderedListOutlined,
  AppstoreAddOutlined,
  BellOutlined,
  CaretRightOutlined,
  RollbackOutlined,
  CheckOutlined,
  CloseOutlined,
  AuditOutlined
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
  SearchOutlined,
  BookOutlined,
  UploadOutlined,
  ApartmentOutlined,
  UndoOutlined,
  RedoOutlined,
  SnippetsOutlined,
  ZoomInOutlined,
  ZoomOutOutlined,
  ExpandOutlined,
  ScanOutlined,
  VerticalAlignTopOutlined,
  VerticalAlignBottomOutlined,
  InfoCircleOutlined,
  FileExcelOutlined,
  CloudUploadOutlined,
  ContactsOutlined,
  ToolOutlined,
  CodeOutlined,
  FormOutlined,
  FileOutlined,
  HistoryOutlined,
  CloudOutlined,
  UnorderedListOutlined,
  AppstoreAddOutlined,
  BellOutlined,
  CaretRightOutlined,
  RollbackOutlined,
  CheckOutlined,
  CloseOutlined,
  AuditOutlined
]

const antdIcon = {
  install (Vue: any) {
    iconComponents.forEach(iconComponent => {
      Vue.component(iconComponent.name, iconComponent)
    })
  }
}

export default antdIcon
