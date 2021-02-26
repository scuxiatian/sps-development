import { modifyCSS } from '@antv/dom-util'
import { Graph } from '@antv/g6'
import { each, wrapBehavior } from '@antv/util'
import BasePlugin from './basePlugin'
import { CommandGraph } from './command'

class Toolbar extends BasePlugin {
  private _events: any
  initPlugin (graph: Graph) {
    const _this = this as any
    this.set('graph', graph)
    const events = this.getEvents()
    const bindEvents: any = {}
    each(events, (v, k) => {
      const event = wrapBehavior(_this, v) as any
      bindEvents[k] = event
      graph.on(k, event)
    })
    this._events = bindEvents

    this.initEvents()
    this.updateToolbar()
  }

  getEvents () {
    return { afteritemselected: 'updateToolbar', aftercommandexecute: 'updateToolbar' }
  }

  initEvents () {
    const graph = this.get('graph') as CommandGraph
    const parentNode = this.get('container') as HTMLElement
    const children = parentNode.querySelectorAll('div > span[data-command]')
    each(children, (child: HTMLElement) => {
      const cmdName = child.getAttribute('data-command') as string
      child.addEventListener('click', () => {
        graph.commandEnable(cmdName) && graph.executeCommand(cmdName, {})
      })
    })
  }

  updateToolbar () {
    const graph = this.get('graph') as CommandGraph
    const parentNode = this.get('container') as HTMLElement
    const children = parentNode.querySelectorAll('div > span[data-command]')
    each(children, (child: HTMLElement) => {
      const cmdName = child.getAttribute('data-command') as string
      const icon = child.children[0] as HTMLElement
      if (graph.commandEnable(cmdName)) {
        modifyCSS(child, {
          cursor: 'pointer'
        })
        modifyCSS(icon, {
          color: '#666'
        })
        icon.setAttribute('color', '#666')
      } else {
        modifyCSS(child, {
          cursor: 'default'
        })
        modifyCSS(icon, {
          color: '#bfbfbf'
        })
        icon.setAttribute('color', '#bfbfbf')
      }
    })
  }
}

export default Toolbar
