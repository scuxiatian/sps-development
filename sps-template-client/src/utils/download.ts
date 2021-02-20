const path = process.env.VUE_APP_BASE_API

// 地址转换
export const convertPath = (src: string) => {
  if (src && src.slice(0, 4) !== 'http') {
    return path + '/' + src
  }
  return src
}

// 下载图片地址和图片名
export const downloadImage = (imgSrc: string, name: string) => {
  const image = new Image()
  image.setAttribute('crossOrigin', 'anonymous')
  image.onload = () => {
    const canvas = document.createElement('canvas')
    canvas.width = image.width
    canvas.height = image.height
    const context = canvas.getContext('2d') as CanvasRenderingContext2D
    context.drawImage(image, 0, 0, image.width, image.height)
    const url = canvas.toDataURL('image/png') // 得到图片的base64编码数据

    const a = document.createElement('a') // 生成一个a元素
    const event = new MouseEvent('click') // 创建一个单击事件
    a.download = name || 'photo' // 设置图片名称
    a.href = url // 将生成的URL设置为a.href属性
    a.dispatchEvent(event) // 触发a的单击事件
  }
  image.src = convertPath(imgSrc)
}
