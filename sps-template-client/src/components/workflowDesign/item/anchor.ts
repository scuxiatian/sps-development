import { IShape } from '@antv/g6'
import Item from '@antv/g6-core/lib/item/item'
import { deepMix } from '@antv/util'
import editorStyle from '../util/defaultStyle'

export default class Anchor extends Item {
  isAnchor: boolean
  hotpot: IShape | undefined
  constructor (cfg: any) {
    super(deepMix(cfg, {
      type: 'anchor',
      isActived: false,
      model: {
        type: 'anchor',
        style: {
          ...editorStyle.anchorPointStyle,
          cursor: editorStyle.cursor.hoverEffectiveAnchor
        }
      }
    }))
    this.enableCapture(true)
    this.isAnchor = true
    this.toFront()
  }

  showHotpot () {
    this.hotpot = this.getContainer().addShape('marker', {
      attrs: {
        ...this.get('model').style,
        ...editorStyle.anchorHotsoptStyle
      }
    })
    this.hotpot.toFront()
    this.getKeyShape().toFront()
  }

  setHotspotActived (act: boolean) {
    if (!this.hotpot) return
    act
      ? this.hotpot.attr(editorStyle.anchorHotsoptActivedStyle)
      : this.hotpot.attr(editorStyle.anchorHotsoptStyle)
  }

  // setActived () {
  //   this.update({ style: { ...editorStyle.anchorPointHoverStyle } })
  // }

  // clearActived () {
  //   this.update({ style: { ...editorStyle.anchorPointStyle } })
  // }
}
