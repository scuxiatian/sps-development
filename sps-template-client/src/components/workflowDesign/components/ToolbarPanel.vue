<template>
  <div class="toolbar">
    <a-tooltip
      v-for="action of actionList"
      :key="action.command"
      :title="action.title">
      <span class="command">
        <span>
          <component :is="action.icon"></component>
        </span>
      </span>
    </a-tooltip>
  </div>
</template>

<script>
export default {
  name: 'ToolbarPanel',
  setup () {
    const actionList = [
      { icon: 'UndoOutlined', command: 'undo', title: '撤销' },
      { icon: 'RedoOutlined', command: 'redo', title: '重复' }
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
}
</style>
