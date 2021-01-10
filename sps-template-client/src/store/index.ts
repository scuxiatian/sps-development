import { createStore } from 'vuex'
import VuexPersistence from 'vuex-persist'
import user from './modules/user'

const vuexLocal = new VuexPersistence({
  storage: window.localStorage,
  modules: ['user']
})

export default createStore({
  modules: {
    user
  },
  plugins: [vuexLocal.plugin]
})
