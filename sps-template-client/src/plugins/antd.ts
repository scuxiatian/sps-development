import {
  Card,
  Form,
  Input,
  Row,
  Col,
  Button,
  Image
} from 'ant-design-vue'
// import 'ant-design-vue/dist/antd.css'

const components = [
  Card,
  Form,
  Input,
  Row,
  Col,
  Button,
  Image
]

export default (app: any) => {
  components.forEach(component => {
    app.use(component)
  })
}
