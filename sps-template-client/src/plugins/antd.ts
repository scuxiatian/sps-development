import {
  Card,
  Form,
  Input,
  Row,
  Col,
  Button,
  Image,
  Layout,
  Menu,
  Breadcrumb,
  Avatar,
  Dropdown,
  Tabs
} from 'ant-design-vue'
// import 'ant-design-vue/dist/antd.css'

const components = [
  Card,
  Form,
  Input,
  Row,
  Col,
  Button,
  Image,
  Layout,
  Menu,
  Breadcrumb,
  Avatar,
  Dropdown,
  Tabs
]

export default (app: any) => {
  components.forEach(component => {
    app.use(component)
  })
}
