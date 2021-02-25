function downloadFile (data: Uint8Array | string, type: string, filename: string) {
  const blob = new Blob([data], { type })
  const link = document.createElement('a')
  if (link.download !== undefined) {
    const url = URL.createObjectURL(blob)
    link.setAttribute('href', url)
    link.setAttribute('download', filename)
    link.style.visibility = 'hidden'
    document.body.appendChild(link)
    link.click()
    document.body.removeChild(link)
  }
}

interface ProcessModel {
  id: string;
  name: string;
  category: string;
  clazz: string;
}

export function exportJSON (json: any, canvas: ProcessModel, createFile = true) {
  const id = canvas.id || 'flow'
  const name = canvas.name || 'flow'

  const jsonData = {
    id,
    name,
    ...json
  }
  const jsonStr = JSON.stringify(jsonData, null, 2)

  if (createFile) {
    downloadFile(jsonStr, 'application/json;charset=utf-8;', `${name}.json`)
  }
}

export function exportImg (canvasPanel: HTMLElement, filename: string, createFile = true) {
  filename = filename || 'flow'
  const canvas = canvasPanel.querySelector('canvas') as HTMLCanvasElement
  const context = canvas.getContext('2d')

  const imgData = context?.getImageData(0, 0, canvas.width, canvas.height).data
  if (!imgData) return
  let left = canvas.width
  let right = 0
  let top = canvas.height
  let bottom = 0
  for (let i = 0; i < canvas.width; i++) {
    for (let j = 0; j < canvas.height; j++) {
      const pos = (i + canvas.width * j) * 4
      if (imgData[pos] > 0 || imgData[pos + 1] > 0 || imgData[pos + 2] > 0 || imgData[pos + 3] > 0) {
        bottom = Math.max(j, bottom) // 找到有色彩的最下端
        right = Math.max(i, right) // 找到有色彩的最右端
        top = Math.min(j, top) // 找到有色彩的最上端
        left = Math.min(i, left) // 找到有色彩的最左端
      }
    }
  }

  const c = document.createElement('canvas')
  // 四周空白余量
  const blankWidth = 60
  c.width = right - left + blankWidth * 2
  c.height = bottom - top + blankWidth * 2
  const ctx = c.getContext('2d') as CanvasRenderingContext2D
  // 设置白底
  ctx.fillStyle = '#fff'
  ctx.fillRect(0, 0, c.width, c.height)
  ctx.drawImage(canvas, left - blankWidth, top - blankWidth, c.width, c.height, 0, 0, c.width, c.height)
  const data = c.toDataURL('image/jpeg')
  if (createFile) {
    const parts = data.split(';base64,')
    const contentType = parts[0].split(':')[1]
    const raw = window.atob(parts[1])
    const uInt8Array = new Uint8Array(raw.length)
    for (let i = 0; i < raw.length; ++i) {
      uInt8Array[i] = raw.charCodeAt(i)
    }
    downloadFile(uInt8Array, contentType, `${filename}.jpg`)
  }
  return data
}
