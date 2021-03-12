<template>
  <div>
    <suspense>
      <template #default>
        <div>
          <FlowContainer />
        </div>
      </template>
      <template #fallback>
        <div>加载中...</div>
      </template>
    </suspense>
  </div>
</template>

<script>
import FlowContainer from './flowContainer'

export default {
  name: 'workflowUse',
  components: {
    FlowContainer
  }
}
</script>
