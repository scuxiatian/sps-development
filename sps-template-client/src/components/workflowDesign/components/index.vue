<template>
  <div class="root">
    <ToolbarPanel />
    <img src="data:image/gif;base64,R0lGODlhAQABAIAAAAAAAP///yH5BAEAAAAALAAAAAABAAEAAAIBRAA7" style="opacity:1"/>
    <div style="display: flex">
      <ItemPanel ref="addItemPanelRef" :height="height" />
      <div
        ref="containerRef"
        class="canvas-panel"
        :style="{
          height: height + 'px',
          width: isView ? '100%' : '70%' }"></div>
    </div>
  </div>
</template>

<script>
import ToolbarPanel from './ToolbarPanel'
import ItemPanel from './ItemPanel'
import G6 from '@antv/g6/lib'
import { onMounted, ref, unref } from 'vue'
import registerShape from '../shape'
import registerBehavior from '../behavior'
import AddItemPanel from '../plugins/addItemPanel'
import CanvasPanel from '../plugins/canvasPanel'

registerShape(G6)
registerBehavior(G6)

export default {
  name: 'WorkflowDesign',
  components: {
    ToolbarPanel,
    ItemPanel
  },
  props: {
    isView: {
      type: Boolean,
      default: false
    },
    mode: {
      type: String,
      default: 'edit'
    },
    height: {
      type: Number,
      default: 800
    }
  },
  setup (props) {
    let graph
    const containerRef = ref(null)
    const addItemPanelRef = ref(null)
    const data = {
      // 点集
      nodes: [],
      // 边集
      edges: []
    }
    onMounted(() => {
      const container = unref(containerRef)
      const addItemPanel = unref(addItemPanelRef)
      let plugins = []
      if (!props.isView) {
        const addItemPanelPlugin = new AddItemPanel({ container: addItemPanel.$el })
        const canvasPanelPlugin = new CanvasPanel({ container })
        plugins = [addItemPanelPlugin, canvasPanelPlugin]
      }
      graph = new G6.Graph({
        container,
        plugins,
        width: container.offsetWidth,
        height: props.height,
        modes: {
          default: ['drag-canvas', 'click-select'],
          view: [],
          edit: ['drag-canvas', 'click-select', 'dragPanelItemAddNode', 'hoverNodeActived',
            'hoverAnchorActived', 'clickSelected', 'dragNode', 'dragEdge']
        },
        defaultEdge: {
          type: 'flow-polyline-round'
        }
      })
      if (props.isView) {
        graph.setMode('view')
      } else {
        graph.setMode(props.mode)
      }
      graph.data(data)
      graph.render()
    })
    return {
      containerRef,
      addItemPanelRef
    }
  }
}
</script>

<style lang="less" scoped>
.root {
  width: 100%;
  height: 100%;
  background-color: #fff;
  display: block;
}
</style>
