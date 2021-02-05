import {
  ConfigProvider,
  Card,
  Form,
  Input,
  InputNumber,
  Radio,
  Checkbox,
  Select,
  TreeSelect,
  Row,
  Col,
  Button,
  Image,
  Layout,
  Menu,
  Breadcrumb,
  Avatar,
  Dropdown,
  Tabs,
  Table,
  Modal,
  Popconfirm,
  Drawer,
  Tree,
  Tag
} from 'ant-design-vue'
// import 'ant-design-vue/dist/antd.css'

const components = [
  ConfigProvider,
  Card,
  Form,
  Input,
  InputNumber,
  Radio,
  Checkbox,
  Select,
  TreeSelect,
  Row,
  Col,
  Button,
  Image,
  Layout,
  Menu,
  Breadcrumb,
  Avatar,
  Dropdown,
  Tabs,
  Table,
  Modal,
  Popconfirm,
  Drawer,
  Tree,
  Tag
]

export default (app: any) => {
  components.forEach(component => {
    app.use(component)
  })
}
