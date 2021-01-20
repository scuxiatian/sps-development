import { createStore } from 'vuex'
import VuexPersistence from 'vuex-persist'
import user from './modules/user'
import router from './modules/router'

const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
  modules: ['user']
})

export default createStore({
  modules: {
    user,
    router
  },
  plugins: [vuexLocal.plugin]
})
