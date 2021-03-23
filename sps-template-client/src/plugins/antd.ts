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
  Collapse,
  DatePicker,
  Steps,
  Switch
} from 'ant-design-vue'

import { registerAntDesignExpand } from '@/components/antDesignExpand'

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
  Collapse,
  DatePicker,
  Steps,
  Switch
]

export default (app: any) => {
  registerAntDesignExpand(app)
  components.forEach(component => {
    app.use(component)
  })
}
