import { Graph } from '@antv/g6'
import BasePlugin from './basePlugin'

class CanvasPanel extends BasePlugin {
  initPlugin (graph: Graph) {
    const parentNode = this.get('container') as HTMLElement
    parentNode.addEventListener('dragover', e => {
      graph.emit('canvas:mousemove', e)
    })
    parentNode.addEventListener('dragleave', e => {
      graph.emit('canvas:mouseleave', e)
    })
  }
}

export default CanvasPanel
