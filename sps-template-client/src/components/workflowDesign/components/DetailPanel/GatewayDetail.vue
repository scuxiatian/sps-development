<template>
  <div :data-clazz="model.clazz">
    <div class="panel-title">{{ titleMap[model.clazz] }}</div>
    <div class="panel-body">
      <DefaultDetail
        :model="model"
        :onChange="onChange"
        :readOnly="readOnly"
      />
    </div>
  </div>
</template>

<script>
import DefaultDetail from './DefaultDetail'

export default {
  name: 'GatewayDetail',
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
  },
  setup () {
    const titleMap = {
      exclusiveGateway: '排他网关',
      parallelGateway: '并行网关',
      inclusiveGateway: '包容网关'
    }

    return {
      titleMap
    }
  }
}
</script>
