<template>
  <div class="login-container">
    <!-- 登录框 -->
    <a-card class="login-card">
      <!-- 登录框标题 -->
      <template class="title-wrap" #title>
        <div class="title-wrap text-primary">
          <CodeSandboxCircleFilled class="title-logo" />
          <span class="title-text">死胖兽开发平台</span>
        </div>
      </template>
      <!-- 登录框表单 -->
      <a-form ref="loginForm" :model="loginModel" :rules="rules" :wrapperCol="{ span: 24 }">
        <a-form-item name="username" has-feedback>
          <a-input
            v-model:value="loginModel.username"
            placeholder="请输入用户名"
          >
            <template #prefix><UserOutlined /></template>
          </a-input>
        </a-form-item>
        <a-form-item name="password" has-feedback>
          <a-input
            v-model:value="loginModel.password"
            placeholder="请输入密码"
            :type="showPassword ? 'text' : 'password'"
          >
            <template #prefix><LockOutlined /></template>
            <template #suffix>
              <EyeInvisibleOutlined v-if="showPassword" class="cursor-pointer" @click="showPassword = false" />
              <EyeOutlined v-else class="cursor-pointer" @click="showPassword = true" />
            </template>
          </a-input>
        </a-form-item>
        <a-form-item name="captcha" has-feedback>
          <a-input
            class="captcha-input"
            v-model:value="loginModel.captcha"
            placeholder="请输入验证码"
            @keyup.enter="submitForm">
            <template #prefix><MailOutlined /></template>
          </a-input>
          <div class="captcha-img">
            <a-image
              class="cursor-pointer"
              :src="captchaSrc"
              width="100%"
              height="100%"
              :preview="false"
              @click="getCaptcha"
            />
          </div>
        </a-form-item>
        <a-form-item>
          <a-button type="primary" shape="round" block @click="submitForm">
            <template #icon><LoginOutlined /></template>登录
          </a-button>
        </a-form-item>
      </a-form>
    </a-card>
  </div>
</template>

<script>
import { reactive, ref, toRefs, getCurrentInstance } from 'vue'
import { captcha } from '@/api/user'
import { useStore, mapActions } from 'vuex'

export default {
  setup () {
    const { ctx } = getCurrentInstance()
    const store = useStore()
    const actions = mapActions('user', ['login'])
    const login = actions.login.bind({ $store: store })

    const state = reactive({
      loginModel: {
        username: 'admin',
        password: '123',
        captcha: '',
        captchaId: ''
      },
      showPassword: false,
      captchaSrc: '',
      rules: {
        username: [{ required: true, message: '请输入用户名', trigger: 'change' }],
        password: [{ required: true, message: '请输入密码', trigger: 'change' }],
        captcha: [
          { required: true, message: '请输入验证码', trigger: 'blur' },
          { len: 6, message: '请输入6位验证码', trigger: 'blur' }
        ]
      }
    })

    // 更新验证码
    const getCaptcha = async () => {
      const res = await captcha()
      state.loginModel.captchaId = res.data.captchaId
      state.captchaSrc = res.data.picPath
    }
    const loginForm = ref(null)
    // 提交登录表单
    const submitForm = async () => {
      try {
        await loginForm.value.validate()
        await login(state.loginModel)
      } catch (err) {
        if (err.errorFields && err.errorFields.length > 0) {
          ctx.$error(err.errorFields[0].errors[0])
        }
      }
    }

    // 启动时先更新一次验证码
    getCaptcha()

    return {
      ...toRefs(state),
      loginForm,
      getCaptcha,
      submitForm
    }
  }
}
</script>

<style lang="less" scoped>
.login-container {
  width: 100%;
  height: 100%;
  background: url('../../assets/login_background.jpg') no-repeat;
  background-size: cover;
  display: flex;
  justify-content: center;
  align-items: center;
  .login-card {
    width: 320px;
    min-height: 300px;
    background-color: rgba(144, 147, 153, 0.2);
    .title-wrap {
      text-align: center;
      .title-logo {
        display: inline-block;
        vertical-align: middle;
        font-size: 30px;
        margin-right: 5px;
      }
      .title-text {
        display: inline-block;
        vertical-align: middle;
        font-size: 20px;
      }
    }
    .ant-form-item {
      margin-bottom: 10px;
      .captcha-input {
        width: 60%;
      }
      .captcha-img {
        width: 30%;
        float: right;
      }
    }
  }
}
</style>>
