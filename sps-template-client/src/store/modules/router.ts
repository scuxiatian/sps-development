import { RouteRecordRaw } from 'vue-router'

interface Route {
  path: string;
  name: string;
  component: string;
  title: string;
  
}

const routerList: Array<RouteRecordRaw> = []
// const formatRouter = (routes)