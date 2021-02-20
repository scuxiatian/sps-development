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
  Tag,
  Upload,
  Tooltip,
  Collapse
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
  Tag,
  Upload,
  Tooltip,
  Collapse
]

export default (app: any) => {
  components.forEach(component => {
    app.use(component)
  })
}
