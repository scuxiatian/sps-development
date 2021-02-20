import editorStyle from '../util/defaultStyle'
import G6 from '@antv/g6/lib'

// const uniqBy = (arr: Array<any>, key: string) => {
//   const result: Array<any> = []
//   arr.forEach(i => {
//     if (!result.find(r => r[key] === i[key])) {
//       result.push(i)
//     }
//   })
//   return result
// }

export default function (g6: typeof G6) {
  g6.registerEdge('flow-polyline-round', {
    options: {
      style: {
        ...editorStyle.edgeStyle
      },
      stateStyles: {
        selected: {
          lineWidth: editorStyle.edgeSelectedStyle.lineWidth
        },
        hover: {
          stroke: editorStyle.edgeActivedStyle.stroke
        }
      }
    },
    setState (name, value, item) {
      const group = item?.getContainer()
      const path = group?.getChildByIndex(0)
      const options: any = this.options
      if (name === 'selected') {
        if (value) {
          path?.attr('lineWidth', options.stateStyles.selected.lineWidth)
          path?.attr('stroke', options.style.stroke)
        } else {
          path?.attr('lineWidth', options.style.lineWidth)
        }
      } else if (name === 'hover') {
        if (value) {
          path?.attr('stroke', options.stateStyles.hover.stroke)
        } else {
          path?.attr('stoke', options.style.stroke)
        }
      }
    }
  }, 'polyline')
}
