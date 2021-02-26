// 驼峰转换下划线
export const toSQLLine = (str: string) => {
  if (!str) return ''
  if (str === 'ID') return 'ID'
  return str.replace(/([A-Z])/g, '_$1').toLowerCase()
}
