import { jsonInBlacklist } from '@/api/jwt'
import { login } from '@/api/user'
import { LoginParams } from '@/api/model/user'
import router from '@/router'
import { ResponseData } from '@/utils/request'
// import router from '@/router'

interface UserInfo {
  uuid?: string;
  nickName?: string;
  headerImg?: string;
  authority?: string;
}

interface UserState {
  userInfo: UserInfo;
  token: string;
}

const state: UserState = {
  userInfo: {
    uuid: '',
    nickName: '',
    headerImg: '',
    authority: ''
  },
  token: ''
}

const mutations = {
  SET_USER_INFO (state: UserState, userInfo: UserInfo) {
    state.userInfo = userInfo
  },
  SET_TOKEN (state: UserState, token: string) {
    state.token = token
  },
  LOGOUT (state: UserState) {
    state.userInfo = {}
    state.token = ''
    router.push({ name: 'login', replace: true })
    sessionStorage.clear()
    window.location.reload()
  }
}

const actions = {
  async Login ({ commit }: any, loginInfo: LoginParams) {
    const res: ResponseData = await login(loginInfo)
    if (res.code === 0) {
      commit('SET_USER_INFO', res.data.user)
      commit('SET_TOKEN', res.data.token)
    }
  },
  async Logout ({ commit }: any) {
    const res: ResponseData = await jsonInBlacklist()
    if (res.code === 0) {
      commit('LOGOUT')
    }
  }
}

const getters = {
  userInfo (state: UserState) {
    return state.userInfo
  },
  token (state: UserState) {
    return state.token
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions,
  getters
}
