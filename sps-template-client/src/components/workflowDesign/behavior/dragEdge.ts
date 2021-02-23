import G6, { Graph, IG6GraphEvent, IGroup, IShape, Item, Node } from '@antv/g6'
import Anchor from '../item/anchor'
import { CommandGraph } from '../plugins/command'
import { NodeGroup } from '../shape/node'
import editorStyle from '../util/defaultStyle'

interface Origin {
  x: number;
  y: number;
  sourceNode: Node;
  sourceAnchor: number;
  targetNode?: Node | null;
  targetAnchor?: Anchor | null;
}

export default function (g6: typeof G6) {
  g6.registerBehavior('dragEdge', {
    getDefaultCfg () {
      return {
        updateEdge: true,
        delegate: true,
        delegateStyle: {},
        dragEdge: false
      }
    },
    getEvents () {
      return {
        'anchor:dragstart': 'onDragStart',
        'anchor:drag': 'onDrag',
        'anchor:dragend': 'onDragEnd',
        'anchor:dragenter': 'onDragEnter',
        'anchor:dragleave': 'onDragLeave'
      } as any
    },
    onDragStart (e: IG6GraphEvent) {
      const _this = this as any
      const graph = this.graph as Graph
      const node = e.target.getParent().getParent().get('item') as Node
      const anchorIndex = e.item?.get('index')
      const point = node.getAnchorPoints()[anchorIndex]
      _this.target = e.item
      const groupId = node.get('groupId')
      if (groupId) {} else {
        _this.origin = {
          x: point.x,
          y: point.y,
          sourceNode: node,
          sourceAnchor: anchorIndex
        } as Origin
        _this.dragEdgeBeforeShowAnchor(e)
      }
      graph.set('edgeDragging', true)
    },
    onDrag (e: IG6GraphEvent) {
      if (!this.origin) return
      const _this = this as any
      _this._updateEdge(this.target, e)
    },
    onDragEnd (e: IG6GraphEvent) {
      const _this = this as any
      const origin = this.origin as Origin
      const graph = this.graph as Graph
      if (!origin) return
      const delegateShape = e.item?.get('edgeDelegate')
      if (delegateShape) {
        delegateShape.remove()
        _this.target.set('edgeDelegate', null)
      }
      _this._updateEdge(_this.target, e, true)
      graph.setItemState(origin.sourceNode, 'show-anchor', false)
      _this.target = null
      _this.origin = null
      graph.set('edgeDragging', false)
    },
    onDragEnter (e: IG6GraphEvent) {
      if (!this.origin) return
      const _this = this as any
      const origin = this.origin as Origin
      if (!_this.sameNode(e)) {
        const item = e.item as any
        item.setHotspotActived(true)
        origin.targetNode = e.target.getParent().getParent().get('item')
        origin.targetAnchor = item.get('index')
      }
    },
    onDragLeave (e: IG6GraphEvent) {
      if (!this.origin) return
      const _this = this as any
      const origin = this.origin as Origin
      if (!_this.sameNode(e)) {
        const item = e.item as any
        item.setHotspotActived(false)
        origin.targetNode = null
        origin.targetAnchor = null
      }
    },
    dragEdgeBeforeShowAnchor () {
      // const origin = this.origin as Origin
      const graph = this.graph as Graph
      // const sourceGroupId = origin.sourceNode.getModel().groupId
      graph.getNodes().forEach(node => {
        if (node.getModel().clazz === 'start') return
        const group = node.getContainer() as NodeGroup
        group.showAnchor()
        group.anchorShape.forEach(a => a.get('item').showHotpot())
      })
    },
    sameNode (e: IG6GraphEvent) {
      const origin = this.origin as Origin
      return e.target.cfg.type === 'marker' &&
        e.target.getParent() &&
        e.target.getParent().getParent().get('item').get('id') === origin.sourceNode.get('id')
    },
    _updateEdge (item: Item, e: IG6GraphEvent, force: boolean) {
      const _this = this as any
      const x = e.x
      const y = e.y
      if (_this.delegate && !force) {
        _this._updateEdgeDelegate(item, x, y)
        return
      }
      const node = e.target.getParent().getParent().get('item')
      const groupId = node.get('groupId')
      if (groupId) {} else {
        _this._addEdge(e)
      }
      _this._clearAllAnchor()
      _this.graph.paint()
    },
    _updateEdgeDelegate (item: Item, x: number, y: number) {
      const _this = this as any
      const graph = _this.graph as Graph
      const origin = _this.origin as Origin
      let edgeShape = item.get('edgeDelegate') as IShape
      if (!edgeShape) {
        const parent = graph.get('group') as IGroup
        edgeShape = parent.addShape('line', {
          attrs: {
            x1: origin.x,
            y1: origin.y,
            x2: x,
            y2: y,
            ...editorStyle.edgeDelegationStyle
          }
        })
        edgeShape.set('capture', false)
        item.set('edgeDelegate', edgeShape)
      }
      edgeShape.attr({ x2: x, y2: y })
      graph.paint()
    },
    _addEdge () {
      const origin = this.origin as Origin
      const graph = this.graph as CommandGraph
      if (origin.targetNode) {
        const timestamp = new Date().getTime()
        const addModel = {
          id: 'flow' + timestamp,
          clazz: 'flow',
          source: origin.sourceNode.get('id'),
          target: origin.targetNode.get('id'),
          sourceAnchor: origin.sourceAnchor,
          targetAnchor: origin.targetAnchor
        }
        graph.executeCommand('add', {
          type: 'edge',
          addModel
        })
        // graph.add('edge', addModel)
      }
    },
    _clearAllAnchor () {
      const graph = this.graph as Graph
      graph.getNodes().forEach(node => {
        const group = node.getContainer() as NodeGroup
        group.clearAnchor()
      })
    }
  })
}
