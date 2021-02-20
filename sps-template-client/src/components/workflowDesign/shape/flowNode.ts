import G6, { ModelConfig } from '@antv/g6'
import { deepMix } from '@antv/util'
import editorStyle from '../util/defaultStyle'

const startDefaultOptions = {
  icon: null,
  iconStyle: {
    width: 23,
    height: 23,
    left: 16,
    top: 16
  },
  style: {
    ...editorStyle.nodeStyle,
    fill: '#FEF7E8',
    stroke: '#FA8C16',
    cursor: 'default'
  },
  stateStyles: {
    selected: {
      fill: '#FCD49A'
    },
    hover: {
      cursor: editorStyle.cursor.hoverNode
    }
  }
}

const taskDefaultOptions = {
  icon: null,
  iconStyle: {
    width: 12,
    height: 12,
    left: 2,
    top: 2
  },
  style: {
    ...editorStyle.nodeStyle,
    fill: '#E7F7FE',
    stroke: '#1890FF',
    cursor: 'default'
  },
  stateStyles: {
    selected: {
      fill: '#95D6FB'
    },
    hover: {
      cursor: editorStyle.cursor.hoverNode
    }
  }
}

export default function (g6: typeof G6) {
  g6.registerNode('start-node', {
    shapeType: 'circle',
    labelPosition: 'bottom',
    options: deepMix({}, startDefaultOptions, { icon: require('../assets/icons/flow/start.svg') }),
    getShapeStyle (cfg: ModelConfig) {
      cfg.size = [55, 55]
      const width = cfg.size[0]
      const style = {
        x: 0,
        y: 0,
        r: width / 2,
        ...startDefaultOptions.style
      }
      if (cfg.color) {
        style.fill = cfg.color
      }
      return style
    },
    afterDraw (config, group) {
      const cfg = config as ModelConfig
      if (cfg.active) {
        const shape = group?.get('children')[0]
        shape.animate({
          repeat: true,
          onFrame (ratio: number) {
            const diff = ratio <= 0.5 ? ratio * 10 : (1 - ratio) * 10
            let radius = cfg.size as any
            if (isNaN(radius)) radius = radius[0]
            return {
              r: radius / 2 + diff
            }
          }
        }, 3000, 'easeCubic')
      }
    },
    getAnchorPoints () {
      return [
        [0.5, 0], // top
        [1, 0.5], // right
        [0.5, 1] // bottom
      ]
    }
  }, 'base-node')
  g6.registerNode('task-node', {
    shapeType: 'rect',
    options: {
      ...taskDefaultOptions
    },
    getShapeStyle (cfg: ModelConfig) {
      const _this = this as any
      cfg.size = [100, 55]
      const width = cfg.size[0]
      const height = cfg.size[1]
      const style = {
        x: 0 - width / 2,
        y: 0 - height / 2,
        width,
        height,
        ..._this.options.style
      }
      return style
    }
  }, 'base-node')
  g6.registerNode('user-task-node', {
    options: deepMix({}, taskDefaultOptions, {
      icon: require('../assets/icons/flow/icon_user.svg'),
      style: {
        fill: '#E7F7FE',
        stroke: '#1890FF'
      },
      stateStyles: {
        selected: {
          fill: '#95D6FB'
        }
      }
    })
  }, 'task-node')
}
