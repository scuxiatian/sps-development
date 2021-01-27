<template>
  <div class="button-box">
    <a-button type="primary" @click="addMenu(0)"><PlusOutlined />新增根菜单</a-button>
  </div>
  <!-- 菜单列表 -->
  <a-table
    size="small"
    bordered
    :columns="columns"
    :dataSource="tableData"
    rowKey="ID"
    :pagination="false"
    :scroll="{ x: 1650 }"
  >
    <template #hidden="{ text }">
      {{ text ? '隐藏' : '显示' }}
    </template>
    <template #action="{ record }">
      <a-button type="primary" @click="addMenu(record.ID)"><PlusCircleOutlined />添加子菜单</a-button>
      <a-button class="table-button" @click="editMenu(record.ID)"><EditOutlined />编辑</a-button>
      <a-popconfirm title="此操作将永久删除所有角色下该菜单, 是否继续?" @confirm="deleteMenu(record.ID)">
        <a-button class="table-button" type="danger"><DeleteOutlined />删除</a-button>
      </a-popconfirm>
    </template>
  </a-table>
  <!-- 新增 / 编辑菜单对话框 -->
  <a-modal
    v-model:visible="dialogVisible"
    :title="dialogTitle"
    width="60%"
    :afterClose="initForm"
    @ok="okDialog">
    <a-form
      ref="menuForm"
      :model="formModel"
      :rules="rules">
      <a-row>
        <a-col :span="8">
          <a-form-item label="菜单Name" name="name">
            <a-input
              v-model:value="formModel.name"
              placeholder="唯一英文字符串"></a-input>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item label="菜单Path" name="path">
            <a-input
              v-model:value="formModel.path"
              placeholder="建议只在后方拼接参数"></a-input>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item label="是否隐藏" name="hidden">
            <a-radio-group
              v-model:value="formModel.hidden"
              placeholder="是否在列表隐藏">
              <a-radio :value="false">否</a-radio>
              <a-radio :value="true">是</a-radio>
            </a-radio-group>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item label="父节点Id" name="parentId">
            <a-tree-select
              v-model:value="formModel.parentId"
              :tree-data="menuOption">
            </a-tree-select>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item label="文件路径" name="component">
            <a-input
              v-model:value="formModel.component"></a-input>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item name="meta.title">
            <template #label><span class="required">*</span> 展示名称</template>
            <a-input
              v-model:value="formModel.meta.title"></a-input>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item label="图标" name="meta.icon">
            <a-select class="icon-select" v-model:value="formModel.meta.icon">
              <a-select-option
                v-for="icon of icons"
                :key="icon"
                :value="icon">
                <span><component :is="icon"></component> {{ icon }}</span>
              </a-select-option>
            </a-select>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item label="排序标记" name="sort">
            <a-input-number v-model:value="formModel.sort"></a-input-number>
          </a-form-item>
        </a-col>
        <a-col :span="8">
          <a-form-item label="keepAlive" name="meta.keepAlive">
            <a-radio-group
              v-model:value="formModel.meta.keepAlive"
              placeholder="是否keepAlive缓存页面">
              <a-radio :value="false">否</a-radio>
              <a-radio :value="true">是</a-radio>
            </a-radio-group>
          </a-form-item>
        </a-col>
      </a-row>
    </a-form>
    <a-button type="primary" @click="addFormParameter"><PlusCircleOutlined />新增菜单参数</a-button>
    <a-table
      size="small"
      :columns="paramsColumns"
      :dataSource="formModel.parameters"
      :rowKey="generateRowKey"
      :pagination="false">
      <template #type="{ record }">
        <a-select v-model:value="record.type">
          <a-select-option key="query" value="query">query</a-select-option>
          <a-select-option key="params" value="params">params</a-select-option>
        </a-select>
      </template>
      <template #key="{ record }">
        <a-input v-model:value="record.key"></a-input>
      </template>
      <template #value="{ record }">
        <a-input v-model:value="record.value"></a-input>
      </template>
      <template #action="{ index }">
        <a-button type="danger" @click="deleteFormParameter(formModel.parameters, index)"><DeleteOutlined/>删除</a-button>
      </template>
    </a-table>
  </a-modal>
</template>

<script>
import { reactive, ref, toRefs, getCurrentInstance } from 'vue'
import { getMenuList, getBaseMenuById, updateBaseMenu, addBaseMenu, deleteBaseMenu } from '@/api/menu'
import { v4 as uuidv4 } from 'uuid'
import useInfoList from '@/mixins/infoList'
import icons from './icons'

export default {
  name: 'MenuManage',
  setup () {
    const { ctx } = getCurrentInstance()
    const state = reactive({
      columns: [
        { title: 'ID', dataIndex: 'ID', width: 100 },
        { title: '路由Name', dataIndex: 'name', width: 160 },
        { title: '路由Path', dataIndex: 'path', width: 160 },
        { title: '是否隐藏', dataIndex: 'hidden', width: 100, slots: { customRender: 'hidden' } },
        { title: '父节点', dataIndex: 'parentId', width: 90 },
        { title: '排序', dataIndex: 'sort', width: 70 },
        { title: '文件路径', dataIndex: 'component', width: 360 },
        { title: '展示名', dataIndex: 'meta.title', width: 120 },
        { title: '图标', dataIndex: 'meta.icon', width: 140 },
        { title: '操作', key: 'action', fixed: 'right', slots: { customRender: 'action' } }
      ],
      paramsColumns: [
        { title: '参数类型', dataIndex: 'type', width: 150, slots: { customRender: 'type' } },
        { title: '参数key', dataIndex: 'key', width: 180, slots: { customRender: 'key' } },
        { title: '参数值', dataIndex: 'value', slots: { customRender: 'value' } },
        { title: '操作', dataIndex: 'action', slots: { customRender: 'action' } }
      ],
      dialogVisible: false,
      dialogTitle: '',
      formModel: {
        meta: {}
      },
      rules: {
        name: [{ required: true, message: '请输入路由Name', trigger: 'blur' }],
        path: [{ required: true, message: '请输入菜单Path', trigger: 'blur' }],
        component: [
          { required: true, message: '请输入文件路径', trigger: 'blur' }
        ],
        'meta.title': [
          {
            validator () {
              if (!state.formModel.meta.title) {
                // eslint-disable-next-line prefer-promise-reject-errors
                return Promise.reject('请输入菜单展示名称')
              }
              return Promise.resolve()
            },
            message: '请输入菜单展示名称',
            trigger: 'blur'
          }
        ]
      },
      isEdit: false,
      checkFlag: false,
      menuOption: [],
      icons
    })

    const {
      tableData,
      getTableData,
      handleSizeChange
    } = useInfoList(getMenuList)

    handleSizeChange(999)

    const menuForm = ref(null)

    // 生成唯一key值
    const generateRowKey = () => {
      return uuidv4()
    }

    // 初始化表单（关闭对话框时调用）
    const initForm = () => {
      menuForm.value.resetFields()
      state.formModel = {
        ID: 0,
        meta: {}
      }
    }

    // 递归生成菜单选项
    const setMenuOptions = (menuData, optionsData, disabled) => {
      menuData && menuData.forEach(item => {
        if (item.children && item.children.length) {
          const option = {
            label: item.meta.title,
            value: String(item.ID),
            // 禁止将菜单的父节点修改为当前菜单自身或者当前菜单的子菜单
            disabled: disabled || item.ID === state.formModel.ID,
            children: []
          }
          setMenuOptions(
            item.children,
            option.children,
            disabled || item.ID === state.formModel.ID
          )
          optionsData.push(option)
        } else {
          const option = {
            label: item.meta.title,
            value: String(item.ID),
            disabled: disabled || item.ID === state.formModel.ID
          }
          optionsData.push(option)
        }
      })
    }

    // 初始化菜单选项
    const setOptions = () => {
      state.menuOption = [
        {
          value: '0',
          label: '根目录'
        }
      ]
      setMenuOptions(tableData.value, state.menuOption, false)
    }

    // 添加菜单方法，id为 0则为添加根菜单
    const addMenu = (id) => {
      state.dialogTitle = '新增菜单'
      state.formModel.parentId = String(id)
      state.isEdit = false
      setOptions()
      state.dialogVisible = true
    }

    // 修改菜单方法
    const editMenu = async (id) => {
      state.dialogTitle = '编辑菜单'
      const res = await getBaseMenuById({ id })
      state.formModel = res.data.menu
      state.isEdit = true
      setOptions()
      state.dialogVisible = true
    }

    // 删除菜单
    const deleteMenu = async (id) => {
      console.log(id)
      const res = await deleteBaseMenu(id)
      if (res.code === 0) {
        ctx.$success('删除成功')
      }
      getTableData()
    }

    // 新增菜单参数行
    const addFormParameter = () => {
      if (!state.formModel.parameters) {
        state.formModel.parameters = []
      }
      state.formModel.parameters.push({
        type: 'query',
        keyParam: '',
        value: ''
      })
    }

    // 删除菜单参数行
    const deleteFormParameter = (parameters, index) => {
      parameters.splice(index, 1)
    }

    // 提交表单
    const okDialog = async () => {
      try {
        await menuForm.value.validate()
        let res
        if (state.isEdit) {
          res = await updateBaseMenu(state.formModel)
        } else {
          res = await addBaseMenu(state.formModel)
        }
        if (res.code === 0) {
          ctx.$success(`${state.isEdit ? '编辑成功' : '添加成功'}!`)
          getTableData()
        }
        initForm()
        state.dialogVisible = false
      } catch (err) {}
    }

    return {
      ...toRefs(state),
      tableData,
      menuForm,
      generateRowKey,
      initForm,
      addMenu,
      editMenu,
      deleteMenu,
      addFormParameter,
      deleteFormParameter,
      okDialog
    }
  }
}
</script>

<style lang="less" scoped>
.icon-select {
  min-width: 200px;
}
</style>
