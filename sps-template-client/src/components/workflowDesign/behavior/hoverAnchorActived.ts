import G6, { Graph, ICombo, IEdge, IG6GraphEvent, INode } from '@antv/g6'

export default function (g6: typeof G6) {
  g6.registerBehavior('hoverAnchorActived', {
    getEvents () {
      return {
        'anchor:mouseenter': 'onAnchorEnter',
        'anchor:mousemove': 'onAnchorEnter',
        'anchor:mouseleave': 'onAnchorLeave'
      } as any
    },
    onAnchorEnter (e: IG6GraphEvent) {
      const graph = this.graph as Graph
      const item = e.item as INode | IEdge | ICombo
      if (!graph.get('edgeDragging')) {
        graph.setItemState(item, 'active-anchor', true)
      }
    },
    onAnchorLeave (e: IG6GraphEvent) {
      if (!e.item) return
      const graph = this.graph as Graph
      if (!graph.get('edgeDragging')) {
        const node = e.item.getContainer().getParent()
        if (node) {
          graph.setItemState(e.item, 'active-anchor', false)
        }
      }
    }
  })
}
