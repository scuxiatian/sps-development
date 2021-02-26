import { IGroup, Item, ModelConfig, Shape } from '@antv/g6'
import { shapeBase } from '@antv/g6-core/lib/element/shapeBase'
import editorStyle from '../util/defaultStyle'

export default function () {
  Shape.registerFactory('anchor', {
    defaultShapeType: 'marker',
    getShape (type: string) {
      const shapeObj = Object.assign({}, shapeBase, {
        type: 'marker',
        itemType: type,
        drawShape (cfg: ModelConfig, group: IGroup) {
          const _this = this as any
          const style = _this.getShapeStyle(cfg)
          const shape = group.addShape('marker', {
            attrs: style,
            name: 'anchor-shape',
            draggable: true
          })
          return shape
        },
        setState (name: string, value: any, item: Item) {
          const _this = this as any
          if (name === 'active-anchor') {
            if (value) {
              _this.update({ style: { ...editorStyle.anchorPointHoverStyle } }, item)
            } else {
              _this.update({ style: { ...editorStyle.anchorPointStyle } }, item)
            }
          }
        }
      })
      return shapeObj
    }
  })
}
