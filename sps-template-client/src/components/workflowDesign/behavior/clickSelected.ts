import G6, { Graph, IG6GraphEvent } from '@antv/g6'

export default function (g6: typeof G6) {
  g6.registerBehavior('clickSelected', {
    getDefaultCfg () {
      return {
        multiple: false
      }
    },
    getEvents () {
      return {
        'node:click': 'onClick',
        'edge:click': 'onClick',
        'canvas:click': 'onCanvasClick',
        'node:mouseover': 'onNodeMouseOver',
        'edge:mouseover': 'onEdgeMouseOver',
        'edge:mouseleave': 'onEdgeMouseLeave'
      }
    },
    onClick (e: IG6GraphEvent) {
      if (!e.item) return
      const _this = this as any
      _this._clearSelected()
      const graph = this.graph as Graph
      graph.setItemState(e.item, 'selected', true)
      let selectedItems = graph.get('selectedItems')
      if (!selectedItems) selectedItems = []
      selectedItems = [e.item.get('id')]
      graph.set('selectedItems', selectedItems)
      graph.emit('afteritemselected', selectedItems)
    },
    onCanvasClick () {
      const _this = this as any
      const graph = this.graph as Graph
      _this._clearSelected()
      graph.emit('afteritemselected', [])
    },
    onNodeMouseOver (e: IG6GraphEvent) {
      if (!e.item) return
      const graph = this.graph as Graph
      if (graph.getCurrentMode() === 'edit') {
        graph.setItemState(e.item, 'hover', true)
      } else {
        graph.setItemState(e.item, 'hover', false)
      }
    },
    onEdgeMouseOver (e: IG6GraphEvent) {
      if (!e.item) return
      const graph = this.graph as Graph
      if (graph.getCurrentMode() === 'edit' && !e.item.hasState('selected')) {
        graph.setItemState(e.item, 'hover', true)
      }
    },
    onEdgeMouseLeave (e: IG6GraphEvent) {
      if (!e.item) return
      const graph = this.graph as Graph
      if (graph.getCurrentMode() === 'edit' && !e.item.hasState('selected')) {
        graph.setItemState(e.item, 'hover', false)
      }
    },
    _clearSelected () {
      const graph = this.graph as Graph
      let selected = graph.findAllByState('node', 'selected')
      selected.forEach(node => {
        graph.setItemState(node, 'selected', false)
      })
      selected = graph.findAllByState('edge', 'selected')
      selected.forEach(edge => {
        graph.setItemState(edge, 'selected', false)
      })
      graph.set('selectedItems', [])
    }
  })
}
