<template>
  <div :data-clazz="model.clazz">
    <div class="panel-title">连接线</div>
    <div class="panel-body">
      <DefaultDetail
        :model="model"
        :onChange="onChange"
        :readOnly="readOnly"
        :iconConfig="false"
      />
      <div class="panel-row">
        <div>条件参数</div>
        <a-input-expand
          style="width: 90%; font-size: 12px"
          type="textarea"
          :disabled="readOnly"
          :value="model.conditionExpression"
          @valueChange="(value) => onChange('conditionExpression', value)"
        />
      </div>
      <div class="panel-row">
        <div>序号</div>
        <a-input-expand
          style="width: 90%; font-size: 12px"
          type="textarea"
          :disabled="readOnly"
          :value="model.seq"
          @valueChange="(value) => onChange('seq', value)"
        />
      </div>
      <!-- <div class="panel-row">
        <a-checkbox
          :checked="!!model.reverse"
          :disabled="readOnly"
          @change="(value) => onChange('reverse', value.target.checked)">
          反向
        </a-checkbox>
      </div> -->
    </div>
  </div>
</template>

<script>
import DefaultDetail from './DefaultDetail'

export default {
  name: 'FlowDetail',
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
