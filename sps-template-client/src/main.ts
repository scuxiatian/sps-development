import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import './router/permission'
import loadAntd from './plugins/antd'
import loadAntdGlobal from './plugins/antd-global'
import AntdIcon from './plugins/antd-icon'

const app = createApp(App).use(store).use(router).use(AntdIcon)
loadAntd(app)
loadAntdGlobal(app)
// const win: any = window
// if (process.env.NODE_ENV === 'development') {
//   if ('__VUE_DEVTOOLS_GLOBAL_HOOK__' in win) {
//     // 这里__VUE_DEVTOOLS_GLOBAL_HOOK__.Vue赋值一个createApp实例
//     win.__VUE_DEVTOOLS_GLOBAL_HOOK__.Vue = app
//   }
//   console.log(win.__VUE_DEVTOOLS_GLOBAL_HOOK__)
// }
app.mount('#app')
