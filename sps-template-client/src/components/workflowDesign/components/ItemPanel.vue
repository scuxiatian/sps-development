<template>
  <div class="item-panel" :style="{ height: height + 'px' }">
    <a-collapse>
      <a-collapse-panel
        v-for="panel of panels"
        :header="panel.header"
        :key="panel.key"
        :forceRender="true">
        <template v-for="item of panel.children" :key="item.title">
          <img :data-item="item.config" :src="require(`../assets/flow/${item.src}`)" />
          <div>{{ item.title }}</div>
        </template>
      </a-collapse-panel>
    </a-collapse>
  </div>
</template>

<script>
export default {
  name: 'ItemPanel',
  props: {
    height: {
      type: Number,
      default: 800
    }
  },
  setup () {
    const panels = [
      {
        header: '开始事件',
        key: 'start',
        children: [
          {
            title: '开始节点',
            src: 'start.svg',
            config: "{clazz:'start',size:'50*50',label:''}"
          }
        ]
      },
      {
        header: '活动',
        key: 'task',
        children: [
          {
            title: '审批节点',
            src: 'user-task.svg',
            config: "{clazz:'userTask',size:'80*44',label:'审批节点'}"
          }
        ]
      },
      {
        header: '网关',
        key: 'gateway',
        children: [
          {
            title: '排他网关',
            src: 'exclusive-gateway.svg',
            config: "{clazz:'exclusiveGateway',size:'60*60',label:''}"
          },
          {
            title: '并行网关',
            src: 'parallel-gateway.svg',
            config: "{clazz:'parallelGateway',size:'60*60',label:''}"
          },
          {
            title: '包容网关',
            src: 'inclusive-gateway.svg',
            config: "{clazz:'inclusiveGateway',size:'60*60',label:''}"
          }
        ]
      },
      {
        header: '结束事件',
        key: 'end',
        children: [
          {
            title: '结束节点',
            src: 'end.svg',
            config: "{clazz:'end',size:'50*50',label:''}"
          }
        ]
      }
    ]
    return {
      panels
    }
  }
}
</script>

<style lang="less" scoped>
.item-panel {
  float: left;
  width: 10%;
  background: #f0f2f5;
  overflow-y: auto;
  border-left: 1px solid #E9E9E9;
  border-bottom: 1px solid #E9E9E9;
  img {
    width: 48px;
    height: 48px;
    padding: 4px;
    border: 1px solid rgba(0,0,0,0);
    border-radius: 2px;
    &:hover{
        border: 1px solid #ccc;
        cursor: move;
    }
  }
  .ant-collapse {
    border: 0;
    .ant-collapse-item {
      > div[role=tab] > div {
        padding-left: 10px;
        border: 1px solid #E9E9E9;
        border-left:0;
      }
      &:first-child{
        > div[role=tab] > div {
          border-top: 0;
        }
      }
      &:last-child{
        > div[role=tab] > div {
          border-bottom: 1px solid #E9E9E9;
        }
      }
      /deep/ .ant-collapse-content-box {
        border-top: 0;
        background: #f0f2f5;
        text-align: center;
      }
    }
  }
}
</style>
