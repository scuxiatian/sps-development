import { asyncMenu } from '@/api/menu'
import { asyncRouterHandle } from '@/utils/asyncRouter'
import { MenuParams } from '@/api/model/menu'

interface RouterShow {
  label: string;
  value: string;
}

const routerList: Array<RouterShow> = []
const formatRouter = (routes: Array<MenuParams>) => {
  routes && routes.map(item => {
    if ((!item.children || item.children.every(ch => ch.hidden)) && item.name !== '404') {
      routerList.push({
        label: item.meta.title,
        value: item.name
      })
    }
    item.meta.hidden = item.hidden
    if (item.children && item.children.length > 0) {
      formatRouter(item.children)
    }
  })
}

interface State {
  asyncRouters: Array<MenuParams>;
  routerList: Array<RouterShow>;
}

const state: State = {
  asyncRouters: [],
  routerList
}

const mutations = {
  setRouterList (state: State, routerList: Array<RouterShow>) {
    state.routerList = routerList
  },
  setAsyncRouter (state: State, asyncRouters: Array<MenuParams>) {
    state.asyncRouters = asyncRouters
  }
}

const actions = {
  async SetAsyncRouter ({ commit }: any) {
    const baseRouter: Array<MenuParams> = [{
      path: '/layout',
      name: 'layout',
      component: 'layout/index.vue',
      meta: {
        title: '底层layout'
      },
      children: []
    }]
    const asyncRouterRes = await asyncMenu()
    const asyncRouter = asyncRouterRes.data.menus
    formatRouter(asyncRouter)
    baseRouter[0].children = asyncRouter
    asyncRouterHandle(baseRouter)
    commit('setAsyncRouter', baseRouter)
    commit('setRouterList', routerList)
    return true
  }
}

const getters = {
  asyncRouters (state: State) {
    return state.asyncRouters
  },
  routerList (state: State) {
    return state.routerList
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
