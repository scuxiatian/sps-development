<template>
  <component
    :is="type === 'textarea' ? 'BaseTextarea' : 'BaseInput'"
    ref="baseInputRef"
    v-bind="getBindValue"
    v-model:value="model"
    @input="onChange">
    <template #default>
      <slot></slot>
    </template>
  </component>
</template>

<script>
import { Input } from 'ant-design-vue'
import { computed, ref } from 'vue'
import _ from 'lodash'

export default {
  name: 'AInputExpand',
  components: {
    BaseInput: Input,
    BaseTextarea: Input.TextArea
  },
  props: {
    value: String,
    type: String
  },
  inheritAttrs: false,
  setup (props, { attrs, emit }) {
    const baseInputRef = ref(null)
    const model = ref(props.value)

    const getBindValue = computed(() => {
      const omitAttrs = _.omit(attrs, ['onValueChange'])
      return { ...omitAttrs }
    })

    const onChange = (e) => {
      emit('valueChange', e.target.value)
    }

    return {
      baseInputRef,
      getBindValue,
      model,
      onChange
    }
  },
  watch: {
    value (val) {
      if (val !== this.model) {
        this.model = val
      }
    }
  }
}
</script>
