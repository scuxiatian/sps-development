import { message } from 'ant-design-vue'

message.config({
  duration: 2,
  maxCount: 3
})

export default (app: any) => {
  app.config.globalProperties.$success = message.success
  app.config.globalProperties.$error = message.error
  app.config.globalProperties.$warning = message.warning
  app.config.globalProperties.$info = message.info
}
