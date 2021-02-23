<template>
  <div class="root">
    <ToolbarPanel ref="toolbarRef" />
    <div style="display: flex">
      <ItemPanel ref="addItemPanelRef" :height="height" />
      <div
        ref="containerRef"
        class="canvas-panel"
        :style="{
          height: height + 'px',
          width: isView ? '100%' : '70%' }">
      </div>
      <DetailPanel
        ref="detailPanelRef"
        :model="selectedModel"
      />
    </div>
  </div>
</template>

<script>
import ToolbarPanel from './ToolbarPanel'
import ItemPanel from './ItemPanel'
import DetailPanel from './DetailPanel'
import G6 from '@antv/g6/lib'
import { onMounted, ref, unref } from 'vue'
import registerShape from '../shape'
import registerBehavior from '../behavior'
import AddItemPanel from '../plugins/addItemPanel'
import CanvasPanel from '../plugins/canvasPanel'
import Toolbar from '../plugins/toolbar'
import Command from '../plugins/command'

registerShape(G6)
registerBehavior(G6)

export default {
  name: 'WorkflowDesign',
  components: {
    ToolbarPanel,
    ItemPanel,
    DetailPanel
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
    const toolbarRef = ref(null)
    const detailPanelRef = ref(null)
    const data = {
      // 点集
      nodes: [],
      // 边集
      edges: []
    }
    onMounted(() => {
      const container = unref(containerRef)
      const addItemPanel = unref(addItemPanelRef)
      const toolbar = unref(toolbarRef)
      let plugins = []
      if (!props.isView) {
        const cmdPlugin = new Command()
        const addItemPanelPlugin = new AddItemPanel({ container: addItemPanel.$el })
        const canvasPanelPlugin = new CanvasPanel({ container })
        const toolbarPlugin = new Toolbar({ container: toolbar.$el })
        plugins = [cmdPlugin, addItemPanelPlugin, canvasPanelPlugin, toolbarPlugin]
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
            'hoverAnchorActived', 'clickSelected', 'dragNode', 'dragEdge', 'itemAlign']
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
      addItemPanelRef,
      toolbarRef,
      detailPanelRef
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
