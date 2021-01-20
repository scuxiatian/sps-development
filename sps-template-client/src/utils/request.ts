import axios from 'axios'
import { message } from 'ant-design-vue'
import store from '@/store'

export interface ResponseData {
  code?: number;
  data?: any;
  msg?: string;
}

const request = axios.create({
  baseURL: process.env.VUE_APP_BASE_API,
  timeout: 999999
})

request.interceptors.request.use(
  config => {
    const token = store.getters['user/token']
    const user = store.getters['user/userInfo']
    config.data = JSON.stringify(config.data)
    config.headers = {
      'Content-Type': 'application/json',
      'x-token': token,
      'x-user-id': user.ID
    }
    return config
  },
  error => {
    message.error(error)
    return error
  }
)

request.interceptors.response.use(
  response => {
    if (response.headers['new-token']) {
      store.commit('user/SET_TOKEN', response.headers['new-token'])
    }
    if (response.data.code === 0) {
      return response.data
    } else {
      message.error(response.data.msg)
      if (response.data.data && response.data.data.reload) {
        store.commit('user/LOGOUT')
      }
      return response.data.msg ? response.data : response
    }
  },
  error => {
    message.error(error)
    return error
  }
)

export default request
