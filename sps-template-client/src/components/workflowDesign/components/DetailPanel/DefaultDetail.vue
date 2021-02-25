<template>
  <div>
    <div class="panel-row">
      <div>标题</div>
      <a-input-expand
        style="width: 90%; font-size: 12px"
        :disabled="readOnly"
        :value="model.label"
        placeholder="请输入标题"
        @valueChange="(value) => onChange('label', value)"
      />
    </div>
    <div v-if="iconConfig" class="panel-row">
      <a-checkbox
        :checked="!!model.hideIcon"
        :disabled="readOnly"
        @change="(value) => onChange('hideIcon', value.target.checked)">
        隐藏图标
      </a-checkbox>
    </div>
  </div>
</template>

<script>
export default {
  name: 'DefaultDetail',
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
    },
    iconConfig: {
      type: Boolean,
      default: true
    }
  }
}
</script>
