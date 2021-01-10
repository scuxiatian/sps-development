import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import loadAntd from './plugins/antd'
import loadAntdGlobal from './plugins/antd-global'
import AntdIcon from './plugins/antd-icon'

const app = createApp(App).use(store).use(router).use(AntdIcon)
loadAntd(app)
loadAntdGlobal(app)
app.mount('#app')
