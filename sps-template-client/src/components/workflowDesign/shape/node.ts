import G6, { IGroup, IShape, ModelConfig, ShapeOptions } from '@antv/g6'
import Anchor from '../item/anchor'

export interface NodeGroup extends IGroup {
  icon: IShape;
  anchorShape: Array<IGroup>;
  showAnchor: () => void;
  clearAnchor: () => void;
}

let g: typeof G6

const nodeDefinition: ShapeOptions = {
  drawAnchor (group: NodeGroup) {
    const bbox = group.get('children')[0].getBBox()
    if (!this.getAnchorPoints) return
    const anchorPoints = this.getAnchorPoints() as number[][]
    anchorPoints.forEach((p, i) => {
      const anchorContainer = group.addGroup()
      const anchor = new Anchor({
        group: anchorContainer,
        index: i,
        model: {
          style: {
            x: bbox.minX + bbox.width * p[0],
            y: bbox.minY + bbox.height * p[1]
          }
        }
      })
      const shapeFactory = g.Shape.getFactory('anchor')
      anchor.set('shapeFactory', shapeFactory)
      anchor.draw()
      group.anchorShape.push(anchorContainer)
    })
  },
  initAnchor (group: NodeGroup) {
    group.anchorShape = []
    group.showAnchor = () => {
      this.drawAnchor(group)
    }
    group.clearAnchor = () => {
      group.anchorShape && group.anchorShape.forEach(a => a.remove())
      group.anchorShape = []
    }
  },
  drawShape (cfg, groupParams) {
    const shapeType = this.shapeType as string
    const group = groupParams as IGroup
    const style = this.getShapeStyle(cfg)
    const shape = group.addShape(shapeType, {
      attrs: {
        ...style
      }
    })
    this.drawIcon(cfg, group)
    this.initAnchor(group)
    return shape
  },
  drawIcon (cfg: ModelConfig, group: NodeGroup) {
    const style = this.getShapeStyle(cfg)
    const options: any = this.options
    if (options.icon) {
      let attrs = {
        x: style.x + options.iconStyle.left,
        y: style.y + options.iconStyle.top,
        width: options.iconStyle.width,
        height: options.iconStyle.height
      }
      if (this.shapeType === 'circle') {
        attrs = {
          x: style.x - style.r + options.iconStyle.left,
          y: style.y - style.r + options.iconStyle.top,
          width: options.iconStyle.width,
          height: options.iconStyle.height
        }
      } else if (this.shapeType === 'path') {
        attrs = {
          x: options.iconStyle.left,
          y: options.iconStyle.top,
          width: options.iconStyle.width,
          height: options.iconStyle.height
        }
      }
      group.icon = group.addShape('image', {
        attrs: {
          img: options.icon,
          ...attrs
        },
        draggable: true
      })
      if (cfg.hideIcon) {
        group.icon.hide()
      }
    }
  },
  setState (name, value, item) {
    const group = item?.getContainer() as NodeGroup
    if (name === 'show-anchor') {
      if (value) {
        group.showAnchor()
      } else {
        group.clearAnchor()
      }
    } else if (name === 'selected') {
      const rect = group.getChildByIndex(0)
      const options = this.options as any
      if (value) {
        rect.attr('fill', options.stateStyles.selected.fill)
      } else {
        rect.attr('fill', options.style.fill)
      }
    } else if (name === 'hover') {
      const rect = group.getChildByIndex(0)
      const text = group.getChildByIndex(1)
      const options = this.options as ModelConfig
      if (value) {
        rect.attr('cursor', options.stateStyles?.hover.cursor)
        text && text.attr('cursor', options.stateStyles?.hover.cursor)
      } else {
        rect.attr('cursor', options.style?.cursor)
        text && text.attr('cursor', options.style?.cursor)
      }
    }
  },
  getAnchorPoints () {
    return [
      [0.5, 0], // top
      [1, 0.5], // right
      [0.5, 1], // bottom
      [0, 0.5] // left
    ]
  }
}

export default function (g6: typeof G6) {
  g = g6
  g6.registerNode('base-node', nodeDefinition, 'single-node')
}
