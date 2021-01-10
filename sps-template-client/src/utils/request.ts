import axios from 'axios'
import { message } from 'ant-design-vue'

const request = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 999999
})

request.interceptors.response.use(
  response => {
    if (response.data.code === 0) {
      return response.data
    } else {
      message.error(response.data.msg)
      return response.data.msg ? response.data : response
    }
  },
  error => {
    message.error(error)
    return error
  }
)

export default request
