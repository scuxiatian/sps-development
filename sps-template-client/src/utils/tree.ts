export const getFullPath = (tree: any, key: string, query: string) => {
  const result: Array<any> = []
  try {
    function getNodePath (node: any) {
      result.push(node)

      if (node[key] === query) {
        throw new Error('find')
      }

      if (node.children && node.children.length > 0) {
        node.children.forEach((child: any) => {
          getNodePath(child)
        })
      } else {
        result.pop()
      }
    }
    getNodePath(tree)
  } catch (e) {
    return result
  }
  return []
}
