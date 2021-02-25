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
        <a-checkbox
          :checked="!!model.success"
          :disabled="readOnly"
          @change="(value) => onChange('success', value.target.checked)">
          是否成功
        </a-checkbox>
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
  name: 'EndEventDetail',
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
