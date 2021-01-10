import { login, UserLogin } from '@/api/user'

interface UserInfo {
  uuid: string;
  nickName: string;
  headerImg: string;
  authority: string;
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
  setUserInfo (state: UserState, userInfo: UserInfo) {
    state.userInfo = userInfo
  },
  setToken (state: UserState, token: string) {
    state.token = token
  }
}

const actions = {
  async login ({ commit }: any, loginInfo: UserLogin) {
    console.log('123')
    const res: any = await login(loginInfo)
    if (res.code === 0) {
      commit('setUserInfo', res.data.user)
      commit('setToken', res.data.token)
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
