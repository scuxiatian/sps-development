import { MenuParams } from '@/api/model/menu'
const _import = require('./_import') // 获取组件的方法

export const asyncRouterHandle = (asyncRouter: Array<MenuParams>) => {
  asyncRouter.map((item: MenuParams) => {
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
