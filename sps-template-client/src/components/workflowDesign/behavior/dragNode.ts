import G6, { Graph, ICanvas, IG6GraphEvent, IGroup, IShape, IShapeBase, Item } from '@antv/g6'
import { Point } from '@antv/g-base/lib'
import editorStyle from '../util/defaultStyle'

export default function (g6: typeof G6) {
  g6.registerBehavior('dragNode', {
    getDefaultCfg () {
      return {
        updateEdge: true,
        delegate: true,
        delegateStyle: {},
        align: true
      }
    },
    getEvents () {
      return {
        'node:dragstart': 'onDragStart',
        'node:drag': 'onDrag',
        'node:dragend': 'onDragEnd'
      }
    },
    onDragStart (e: IG6GraphEvent) {
      if (!this.shouldBegin?.call(this, e)) return
      this.target = e.item
      this.origin = {
        x: e.x,
        y: e.y
      }
    },
    onDrag (e: IG6GraphEvent) {
      const _this = this as any
      if (!_this.origin) return
      if (!_this.get('shouldUpdate').call(this, e)) return
      const origin = _this.origin
      const groupId = _this.target.get('groupId')
      const model = _this.target.get('model')
      if (!_this.point) {
        _this.point = {
          x: model.x,
          y: model.y
        }
      }
      if (groupId) {} else {
        const x = e.x - origin.x + _this.point.x
        const y = e.y - origin.y + _this.point.y
        _this.origin = { x: e.x, y: e.y }
        _this.point = { x, y }
        if (_this.delegate) {
          _this._updateDelegate(_this.target, x, y)
        }
      }
    },
    onDragEnd (e: IG6GraphEvent) {
      const _this = this as any
      if (!_this.shouldEnd?.call(this, e)) return
      if (!_this.origin) return
      const delegateShape = e.item?.get('delegateShape') as IShape | undefined
      const groupId = _this.target.get('groupId')
      if (groupId) {} else {
        if (delegateShape) {
          const bbox = delegateShape.getBBox()
          const x = bbox.x + bbox.width / 2
          const y = bbox.y + bbox.height / 2
          delegateShape.remove()
          _this.target.set('delegateShape', null)
          _this._updateItem(_this.target, { x, y })
        }
      }
      _this.point = null
      _this.origin = null
      _this.graph.emit('afternodedragend', _this.target)
    },
    _updateDelegate (item: IShapeBase & ICanvas, x: number, y: number) {
      const graph = this.graph as Graph
      let shape = item.get('delegateShape') as IShape
      const bbox = item.get('keyShape').getBBox()
      if (!shape) {
        const parent = graph.get('group') as IGroup
        const attrs = editorStyle.nodeDelegationStyle
        // model上的x, y是相对于图形中心的，delegateShape是g实例，x,y是绝对坐标
        shape = parent.addShape('rect', {
          attrs: {
            width: bbox.width,
            height: bbox.height,
            x: x - bbox.width / 2,
            y: y - bbox.height / 2,
            nodeId: item.get('id'),
            ...attrs
          }
        })
        shape.set('capture', false)
        item.set('delegateShape', shape)
      }
      shape.attr({ x: x - bbox.width / 2, y: y - bbox.height / 2 })
      graph.paint()
      graph.emit('afternodedrag', shape)
    },
    _updateItem (item: Item, point: Point) {
      const graph = this.graph as Graph
      item.updatePosition(point)
      graph.paint()
    }
  })
}
