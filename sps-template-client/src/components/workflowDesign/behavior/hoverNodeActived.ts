import G6, { Graph, IG6GraphEvent } from '@antv/g6'

export default function (g6: typeof G6) {
  g6.registerBehavior('hoverNodeActived', {
    getEvents () {
      return {
        'node:mouseenter': 'onNodeEnter',
        'node:mouseleave': 'onNodeLeave',
        'anchor:mouseleave': 'onAnchorLeave'
      }
    },
    onNodeEnter (e: IG6GraphEvent) {
      if (!e.item) return
      const graph = this.graph as Graph
      const clazz = e.item.getModel().clazz
      if (clazz !== 'endEvent' && !graph.get('edgeDragging')) {
        graph.setItemState(e.item, 'show-anchor', true)
      }
    },
    onNodeLeave (e: IG6GraphEvent) {
      if (!e.item) return
      const graph = this.graph as Graph
      if (!(e.target.cfg.type === 'marker') && !graph.get('edgeDragging')) {
        graph.setItemState(e.item, 'show-anchor', false)
      }
    },
    onAnchorLeave (e: IG6GraphEvent) {
      if (!e.item) return
      const node = e.item.getContainer().getParent()
      const graph = this.graph as Graph
      if (node && !graph.get('edgeDragging')) {
        graph.setItemState(node.get('item'), 'show-anchor', false)
      }
    }
  })
}
