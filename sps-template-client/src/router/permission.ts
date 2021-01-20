import router from '@/router'
import store from '@/store'

let asyncRouterFlag = 0

const whiteList = ['login']

router.beforeEach(async (to: any, from, next) => {
  const token = store.getters['user/token']
  if (whiteList.indexOf(to.name) > -1) {
    if (token) {
      next({ path: '/layout/dashboard' })
    } else {
      next()
    }
  } else {
    if (token) {
      if (!asyncRouterFlag) {
        asyncRouterFlag++
        await store.dispatch('router/SetAsyncRouter')
        const asyncRouters = store.getters['router/asyncRouters']
        router.addRoute(asyncRouters[0])
        next({ ...to, replace: true })
      } else {
        next()
      }
    }
    // 不在白名单中并且未登陆的时候
    if (!token) {
      next({
        name: 'login',
        query: {
          redirect: document.location.hash
        }
      })
    }
  }
})
