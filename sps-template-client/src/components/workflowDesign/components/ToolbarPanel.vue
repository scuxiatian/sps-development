<template>
  <div class="toolbar">
    <template
      v-for="action of actionList"
      :key="action.command">
      <a-tooltip
        :title="action.title">
        <span class="command" :data-command="action.command">
          <span>
            <component :is="action.icon"></component>
          </span>
        </span>
      </a-tooltip>
      <span v-if="action.separator" class="separator" />
    </template>
  </div>
</template>

<script>
export default {
  name: 'ToolbarPanel',
  setup () {
    const actionList = [
      { icon: 'UndoOutlined', command: 'undo', title: '撤销' },
      { icon: 'RedoOutlined', command: 'redo', title: '重做', separator: true },
      { icon: 'CopyOutlined', command: 'copy', title: '复制' },
      { icon: 'SnippetsOutlined', command: 'paste', title: '粘贴' },
      { icon: 'DeleteOutlined', command: 'delete', title: '删除', separator: true },
      { icon: 'ZoomInOutlined', command: 'zoomIn', title: '放大' },
      { icon: 'ZoomOutOutlined', command: 'zoomOut', title: '缩小' },
      { icon: 'ExpandOutlined', command: 'resetZoom', title: '实际大小' },
      { icon: 'ScanOutlined', command: 'autoFit', title: '适应屏幕', separator: true },
      { icon: 'VerticalAlignTopOutlined', command: 'toFront', title: '移到上一层' },
      { icon: 'VerticalAlignBottomOutlined', command: 'toBack', title: '移到下一层' }
    ]

    return {
      actionList
    }
  }
}
</script>

<style lang="less" scoped>
.toolbar {
  text-align: left;
  width: 100%;
  padding: 8px 0;
  background-color: #fff;
  border: 1px solid #E9E9E9;
  box-shadow: 0 8px 12px 0 rgba(0, 52, 107, 0.04);
  .command {
    display: inline-block;
    margin: 0 6px;
    line-height: 27px;
    border: 1px solid rgba(2, 2, 2, 0);
    border-radius: 2px;

    span {
      margin: 0 6px;
    }

    &:nth-of-type(1) {
      margin-left: 24px;
    }

    &:hover {
      border: 1px solid #E9E9E9;
      cursor: pointer;
    }
  }
  .separator {
    margin: 4px;
    border-left: 1px solid #E9E9E9;
  }
}
</style>
