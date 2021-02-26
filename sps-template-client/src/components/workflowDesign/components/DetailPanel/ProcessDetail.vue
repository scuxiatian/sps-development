<template>
  <div :data-clazz="model.clazz">
    <div class="panel-title">流程</div>
    <div class="panel-body">
      <div class="panel-row">
        <div>流程标识(id)</div>
        <a-input-expand
          style="width: 90%; font-size: 12px"
          :disabled="readOnly"
          placeholder="必填（唯一标识）"
          :value="model.id"
          @valueChange="(value) => onChange('id', value)"
        />
      </div>
      <DefaultDetail
        :model="model"
        :onChange="onChange"
        :readOnly="readOnly"
        :iconConfig="false"
      />
      <div class="panel-row">
        <div>流程名称</div>
        <a-input-expand
          style="width: 90%; font-size: 12px"
          :disabled="readOnly"
          placeholder="请输入流程名称"
          :value="model.name"
          @valueChange="(value) => onChange('name', value)"
        />
      </div>
      <div class="panel-row">
        <div>默认视图文件路径</div>
        <a-input-expand
          style="width: 90%; font-size: 12px"
          :disabled="readOnly"
          placeholder="请输入视图文件路径"
          :value="model.view"
          @valueChange="(value) => onChange('view', value)"
        />
      </div>
      <div class="panel-row">
        <div>详情说明</div>
        <a-input-expand
          style="width: 90%; font-size: 12px"
          :disabled="readOnly"
          type="textarea"
          placeholder="请输入详情说明"
          :value="model.description"
          @valueChange="(value) => onChange('description', value)"
        />
      </div>
    </div>
  </div>
</template>

<script>
import DefaultDetail from './DefaultDetail'

export default {
  name: 'ProcessDetail',
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
