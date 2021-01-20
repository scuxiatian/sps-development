import { asyncMenu } from '@/api/menu'
import { asyncRouterHandle } from '@/utils/asyncRouter'

interface Meta {
  title: string;
  keepAlive?: boolean;
  defaultMenu?: boolean;
  icon?: string;
  hidden?: boolean;
}

export interface Router {
  path?: string;
  name: string;
  component?: any;
  meta: Meta;
  children?: Array<Router>;
  hidden?: boolean;
  menuId?: string;
}

interface RouterShow {
  label: string;
  value: string;
}

const routerList: Array<RouterShow> = []
const formatRouter = (routes: Array<Router>) => {
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
  asyncRouters: Array<Router>;
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
  setAsyncRouter (state: State, asyncRouters: Array<Router>) {
    state.asyncRouters = asyncRouters
  }
}

const actions = {
  async SetAsyncRouter ({ commit }: any) {
    const baseRouter: Array<Router> = [{
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
