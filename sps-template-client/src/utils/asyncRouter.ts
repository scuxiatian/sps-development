import { Router } from '@/store/modules/router'
const _import = require('./_import') // 获取组件的方法

export const asyncRouterHandle = (asyncRouter: Array<Router>) => {
  asyncRouter.map((item: Router) => {
    if (item.component) {
      item.component = _import(item.component)
    } else {
      delete item.component
    }
    if (item.children) {
      asyncRouterHandle(item.children)
    }
  })
}
