import { deepMix } from '@antv/util'

class BasePlugin {
  protected _cfgs: any;

  constructor (cfgs: any) {
    this._cfgs = deepMix(this.getDefaultCfg(), cfgs)
  }

  getDefaultCfg (): any {
    return { container: null }
  }

  get (key: string) {
    return this._cfgs[key]
  }

  set (key: string, val: any) {
    this._cfgs[key] = val
  }

  destroy () {
    this.get('canvas').destroy()
    const container = this.get('container')
    container.parentNode.removeChild(container)
  }
}

export default BasePlugin
