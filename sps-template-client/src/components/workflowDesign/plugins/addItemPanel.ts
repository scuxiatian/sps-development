import { createDom } from '@antv/dom-util'
import { Graph } from '@antv/g6'
import { each } from '@antv/util'
import BasePlugin from './basePlugin'

class AddItemPanel extends BasePlugin {
  initPlugin (graph: Graph) {
    const parentNode = this.get('container') as HTMLElement
    const ghost = createDom('<img src="data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7"' + ' style="opacity:0"/>')
    const children = parentNode.querySelectorAll('.ant-collapse > .ant-collapse-item > .ant-collapse-content > .ant-collapse-content-box > img[data-item]')
    each(children, (child: Element) => {
      const addModel = (new Function('return ' + child.getAttribute('data-item')))()
      child.addEventListener('dragstart', (e: any) => {
        e.dataTransfer.setDragImage(ghost, 0, 0)
        graph.set('addNodeDragging', true)
        graph.set('addModel', addModel)
      })
      child.addEventListener('dragend', e => {
        graph.emit('canvas:mouseup', e)
        graph.set('addNodeDragging', false)
        graph.set('addModel', null)
      })
    })
  }
}

export default AddItemPanel
