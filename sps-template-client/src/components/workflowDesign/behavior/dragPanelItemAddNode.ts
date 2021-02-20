import G6, { Graph, IG6GraphEvent, IGroup, IShape } from '@antv/g6/lib'
import type { Point } from '@antv/g-base/lib'
import editorStyle from '../util/defaultStyle'
import { getShapeName } from '../util/clazz'

export default function (g6: typeof G6) {
  g6.registerBehavior('dragPanelItemAddNode', {
    getEvents () {
      return {
        'canvas:mousemove': 'onMouseMove',
        'canvas:mouseup': 'onMouseUp',
        'canvas:mouseleave': 'onMouseLeave'
      }
    },
    onMouseMove (e: IG6GraphEvent) {
      const graph = this.graph as Graph
      if (graph.get('addNodeDragging')) {
        let delegateShape = graph.get('addDelegateShape') as IShape | undefined
        const addModel = graph.get('addModel')
        const width = parseInt(addModel.size.split('*')[0])
        const height = parseInt(addModel.size.split('*')[1])
        const point = graph.getPointByClient(e.x, e.y)
        const x = point.x
        const y = point.y
        if (!delegateShape) {
          const parent = graph.get('group') as IGroup
          delegateShape = parent.addShape('rect', {
            attrs: {
              width,
              height,
              x: x - width / 2,
              y: y - height / 2,
              ...editorStyle.nodeDelegationStyle
            }
          })
          delegateShape.set('capture', false)
          graph.set('addDelegateShape', delegateShape)
        }
        delegateShape.attr({ x: x - width / 2, y: y - height / 2 })
        graph.paint()
        graph.emit('afternodedrag', delegateShape)
      }
    },
    onMouseUp (e: IG6GraphEvent) {
      const _this = this as any
      const graph = this.graph as Graph
      if (graph.get('addNodeDragging')) {
        const p = graph.getPointByClient(e.clientX, e.clientY)
        if (p.x > 0 && p.y > 0) {
          _this._addNode(p)
        }
      }
    },
    onMouseLeave () {
      const _this = this as any
      const graph = this.graph as Graph
      if (graph.get('addDelegateShape')) {
        _this._clearDelegate()
        graph.emit('afternodedragend')
      }
    },
    _clearDelegate () {
      const graph = this.graph as Graph
      const delegateShape = graph.get('addDelegateShape') as IShape
      if (delegateShape) {
        delegateShape.remove()
        graph.set('addDelegateShape', null)
        graph.paint()
      }
      graph.emit('afternodedragend')
    },
    _addNode (p: Point) {
      const _this = this as any
      const graph = this.graph as Graph
      if (graph.get('addNodeDragging')) {
        const addModel = graph.get('addModel')
        const { clazz = 'userTask' } = addModel
        addModel.type = getShapeName(clazz)
        addModel.shape = getShapeName(clazz)
        const timestamp = new Date().getTime()
        const id = clazz + timestamp
        const x = p.x
        const y = p.y

        graph.add('node', {
          ...addModel,
          x,
          y,
          id
        })
      }
      _this._clearDelegate()
    }
  })
}
