module.exports = {
  devServer: {
    port: 8889,
    proxy: {
      [process.env.VUE_APP_BASE_API]: {
        target: 'http://127.0.0.1:8888',
        changeOrigin: true,
        pathRewrite: {
          ['^' + process.env.VUE_APP_BASE_API]: ''
        }
      }
    }
  },
  css: {
    loaderOptions: {
      less: {
        globalVars: {
          primary: '#1890ff',
          'white-bg': '#fff'
        }
      }
    }
  }
}
