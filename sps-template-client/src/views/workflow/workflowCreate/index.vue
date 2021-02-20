<template>
  <div>
    <WorkflowDesign
      :height="550" />
  </div>
</template>

<script>
import WorkflowDesign from '@/components/workflowDesign'

export default {
  name: 'workflowCreate',
  components: {
    WorkflowDesign
  }
}
</script>
