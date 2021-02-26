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
        :height="height"
        :users="users"
        :authorities="authorities"
        :onChange="(key, val) => onItemCfgChange(key, val)"
      />
    </div>
  </div>
</template>

<script>
import ToolbarPanel from './ToolbarPanel'
import ItemPanel from './ItemPanel'
import DetailPanel from './DetailPanel'
import G6 from '@antv/g6/lib'
import { onMounted, reactive, ref, toRefs, unref } from 'vue'
import registerShape from '../shape'
import registerBehavior from '../behavior'
import AddItemPanel from '../plugins/addItemPanel'
import CanvasPanel from '../plugins/canvasPanel'
import Toolbar from '../plugins/toolbar'
import Command from '../plugins/command'
import { exportImg, exportJSON } from '../util/bpmn'

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
    },
    users: {
      type: Array,
      default: () => ([])
    },
    authorities: {
      type: Array,
      default: () => ([])
    },
    propProcessModel: {
      type: Object,
      default: () => ({
        id: '',
        name: '',
        category: '',
        clazz: 'process',
        dataObjs: [],
        signalDefs: [],
        messageDefs: []
      })
    }
  },
  setup (props) {
    let graph
    const containerRef = ref(null)
    const addItemPanelRef = ref(null)
    const toolbarRef = ref(null)
    const detailPanelRef = ref(null)

    const state = reactive({
      selectedModel: {},
      processModel: {}
    })

    const data = {
      // 点集
      nodes: [],
      // 边集
      edges: []
    }

    const init = () => {
      state.processModel = props.propProcessModel
      state.selectedModel = state.processModel
    }
    init()

    const initEvents = () => {
      // 节点选中状态变化事件
      graph.on('afteritemselected', (items) => {
        if (items && items.length > 0) {
          const item = graph.findById(items[0])
          state.selectedModel = { ...item.getModel() }
        } else {
          state.selectedModel = state.processModel
        }
      })
      // 撤销 / 重做指令执行后更新选中模型
      graph.on('aftercommandexecute', ({ command }) => {
        if (command.name !== 'undo' && command.name !== 'redo') return
        if (state.selectedModel.clazz === 'process') {
          const selected = graph.findAllByState('node', 'selected')[0]
          selected && (state.selectedModel = { ...selected.getModel() })
        }
      })
    }

    // 配置项内容变化事件
    const onItemCfgChange = (key, value) => {
      const items = graph.get('selectedItems')
      if (items && items.length > 0) {
        const item = graph.findById(items[0])
        graph.executeCommand('update', {
          itemId: items[0],
          updateModel: { [key]: value }
        })
        state.selectedModel = { ...item.getModel() }
      } else {
        const canvasModel = { ...state.processModel, [key]: value }
        state.selectedModel = canvasModel
        state.processModel = canvasModel
      }
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
      initEvents()
    })

    const saveImg = (createFile = true) => exportImg(unref(containerRef), state.processModel.name, createFile)

    const saveJSON = (createFile = true) => exportJSON(graph.save(), state.processModel, createFile)

    const saveFlow = () => graph.save()

    return {
      ...toRefs(state),
      graph,
      containerRef,
      addItemPanelRef,
      toolbarRef,
      detailPanelRef,
      onItemCfgChange,
      saveImg,
      saveJSON,
      saveFlow
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
