import G6, { Graph, ICombo, IEdge, IG6GraphEvent, INode } from '@antv/g6'

export default function (g6: typeof G6) {
  g6.registerBehavior('clickSelected', {
    getDefaultCfg () {
      return {
        multiple: false
      }
    },
    getEvents () {
      return {
        'node:mouseover': 'onNodeMouseOver'
      }
    },
    onNodeMouseOver (e: IG6GraphEvent) {
      const graph = this.graph as Graph
      const item = e.item as INode | IEdge | ICombo
      if (graph.getCurrentMode() === 'edit') {
        graph.setItemState(item, 'hover', true)
      } else {
        graph.setItemState(item, 'hover', false)
      }
    }
  })
}
