<template>
  <div :data-clazz="model.clazz">
    <div class="panel-title">开始节点</div>
    <div class="panel-body">
      <DefaultDetail
        :model="model"
        :onChange="onChange"
        :readOnly="readOnly"
      />
      <div class="panel-row">
        <div>视图文件路径</div>
        <a-input-expand
          style="width: 90%; font-size: 12px"
          :disabled="readOnly"
          placeholder="请输入视图文件路径"
          :value="model.view"
          @valueChange="(value) => onChange('view', value)"
        />
      </div>
      <div class="panel-row">
        <div>步骤</div>
        <a-input-expand
          style="width: 90%; font-size: 12px"
          :disabled="readOnly"
          placeholder="请输入步骤"
          :value="model.step"
          @valueChange="(value) => onChange('step', value)"
        />
      </div>
    </div>
  </div>
</template>

<script>
import DefaultDetail from './DefaultDetail'

export default {
  name: 'StartEventDetail',
  components: {
    DefaultDetail
  },
  props: {
    model: {
      type: Object,
      default: () => ({})
    },
    onChange: {
      type: Function,
      default: () => {}
    },
    readOnly: {
      type: Boolean,
      default: false
    }
  }
}
</script>
