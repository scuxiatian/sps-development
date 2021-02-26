<template>
  <div :data-clazz="model.clazz">
    <div class="panel-title">审批节点</div>
    <div class="panel-body">
      <DefaultDetail
        :model="model"
        :onChange="onChange"
        :readOnly="readOnly"
      />
      <div class="panel-row">
        <div>选择审批人类别</div>
        <a-select
          style="width: 90%; font-size: 12px"
          placeholder="请选择角色（与用户互斥）"
          :value="model.assignType"
          :disabled="readOnly"
          @change="(value) => {
            onChange('assignType', value)
            onChange('assignValue', [])
          }">
          <a-select-option key="user" value="user">用户</a-select-option>
          <a-select-option key="authority" value="authority">角色</a-select-option>
          <a-select-option key="self" value="self">发起人本人</a-select-option>
        </a-select>
      </div>
      <div v-if="model.assignType === 'user'" class="panel-row">
        <div>选择用户</div>
        <a-select
          style="width: 90%; font-size: 12px"
          placeholder="请选择用户"
          :value="model.assignValue"
          allowClear
          mode="multiple"
          show-search
          :disabled="readOnly"
          @change="(value) => onChange('assignValue', value)">
          <a-select-option
            v-for="user in users"
            :key="user.id"
            :value="user.id">
            {{ user.name }}
          </a-select-option>
        </a-select>
      </div>
      <div v-else-if="model.assignType === 'authority'" class="panel-row">
        <div>选择角色</div>
        <a-select
          style="width: 90%; font-size: 12px"
          placeholder="请选择角色"
          :value="model.assignValue"
          allowClear
          mode="multiple"
          show-search
          :disabled="readOnly"
          @change="(value) => onChange('assignValue', value)">
          <a-select-option
            v-for="authoritie in authorities"
            :key="authoritie.id"
            :value="authoritie.id">
            {{ authoritie.name }}
          </a-select-option>
        </a-select>
      </div>
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
  name: 'UserTaskDetail',
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
    },
    users: {
      type: Array,
      default: () => ([])
    },
    authorities: {
      type: Array,
      default: () => ([])
    },
    groups: {
      type: Array,
      default: () => ([])
    }
  }
}
</script>
