import { App } from 'vue'
import { Input } from './input'

const componets = [
  Input
]

export const registerAntDesignExpand = (app: App) => {
  componets.forEach((component) => {
    app.component(component.name, component)
  })
}
