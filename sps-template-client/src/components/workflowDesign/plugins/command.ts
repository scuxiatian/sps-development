import { Graph } from '@antv/g6'
import { clone, isString, mix } from '@antv/util'
import BasePlugin from './basePlugin'

export interface CommandGraph extends Graph {
  executeCommand: (name: string, cfg: any) => void;
  commandEnable: (name: string) => boolean;
  getCommands: () => Array<any>;
  getCurrentCommand: () => any;
}

interface CommandConfig {
  queue?: boolean;
  enable?: (graph: Graph) => void;
  execute?: (graph: Graph) => void;
  back?: (graph: Graph) => void;
  method?: (graph: Graph) => void;
  shortcutCodes?: Array<any>;
}

class Command extends BasePlugin {
  list: any[]
  queue: any[]
  cmds: any

  constructor () {
    super({})
    this._cfgs = this.getDefaultCfg()
    this.list = []
    this.queue = []
    this.cmds = {}
  }

  getDefaultCfg () {
    return { _command: { zoomDelta: 0.1, queue: [], current: 0, clipboard: [] } }
  }

  initPlugin (graph: CommandGraph) {
    this.initCommands()
    graph.getCommands = () => {
      return this.get('_command').queue
    }
    graph.getCurrentCommand = () => {
      const c = this.get('_command')
      return c.queue[c.current - 1]
    }
    graph.executeCommand = (name, cfg) => {
      this.execute(name, graph, cfg)
    }
    graph.commandEnable = (name) => {
      return this.enable(name, graph)
    }
  }

  initCommands () {
    const cmdPlugin = this
    cmdPlugin.registerCommand('add', {
      enable () {
        const _this = this as any
        return _this.type && _this.addModel
      },
      execute (graph) {
        const _this = this as any
        const item = graph.add(_this.type, _this.addModel)
        if (_this.executeTimes === 1) {
          _this.addId = item.get('id')
        }
      },
      back (graph) {
        const _this = this as any
        graph.remove(_this.addId)
      }
    })
    cmdPlugin.registerCommand('update', {
      enable () {
        const _this = this as any
        return _this.itemId && _this.updateModel
      },
      execute (graph) {
        const _this = this as any
        const item = graph.findById(_this.itemId)
        if (item) {
          if (_this.executeTimes === 1) {
            _this.originModel = mix({}, item.getModel())
          }
          graph.update(item, _this.updateModel)
        }
      },
      back (graph) {
        const _this = this as any
        const item = graph.findById(_this.itemId)
        graph.update(item, _this.originModel)
      }
    })
    cmdPlugin.registerCommand('redo', {
      queue: false,
      enable (graph) {
        const mode = graph.getCurrentMode()
        const manager = cmdPlugin.get('_command')
        return mode === 'edit' && manager.current < manager.queue.length
      },
      execute (graph) {
        const manager = cmdPlugin.get('_command')
        const cmd = manager.queue[manager.current]
        cmd && cmd.execute(graph)
        manager.current++
      }
    })
    cmdPlugin.registerCommand('undo', {
      queue: false,
      enable (graph) {
        const mode = graph.getCurrentMode()
        return mode === 'edit' && cmdPlugin.get('_command').current > 0
      },
      execute (graph) {
        const manager = cmdPlugin.get('_command')
        const cmd = manager.queue[manager.current - 1]
        if (cmd) {
          cmd.executeTimes++
          cmd.back(graph)
        }
        manager.current--
      },
      shortcutCodes: [['metaKey', 'z'], ['ctrlKey', 'z']]
    })
    cmdPlugin.registerCommand('copy', {
      queue: false,
      enable (graph) {
        const mode = graph.getCurrentMode()
        const items = graph.get('selectedItems')
        return mode === 'edit' && items && items.length > 0
      },
      method (graph) {
        const manager = cmdPlugin.get('_command')
        manager.clipboard = []
        const items = graph.get('selectedItems')
        if (items && items.length > 0) {
          const item = graph.findById(items[0])
          if (item) {
            manager.clipboard.push({ type: item.get('type'), model: item.getModel() })
          }
        }
      }
    })
    cmdPlugin.registerCommand('paste', {
      enable (graph) {
        const mode = graph.getCurrentMode()
        return mode === 'edit' && cmdPlugin.get('_command').clipboard.length > 0
      },
      method (graph) {
        const _this = this as any
        const manager = cmdPlugin.get('_command')
        _this.pasteData = clone(manager.clipboard[0])
        const addModel = _this.pasteData.model
        addModel.x && (addModel.x += 10)
        addModel.y && (addModel.y += 10)
        const { clazz = 'userTask' } = addModel
        const timeStamp = new Date().getTime()
        const id = clazz + timeStamp
        addModel.id = id
        const item = graph.add(_this.pasteData.type, addModel)
        item.toFront()
      }
    })
    cmdPlugin.registerCommand('delete', {
      enable (graph) {
        const mode = graph.getCurrentMode()
        const selectedItems = graph.get('selectedItems')
        return mode === 'edit' && selectedItems && selectedItems.length > 0
      },
      method (graph) {
        const selectedItems = graph.get('selectedItems') as string[]
        graph.emit('beforedelete', { item: selectedItems })
        if (selectedItems && selectedItems.length > 0) {
          selectedItems.forEach(i => {
            const node = graph.findById(i)
            if (node) {
              graph.remove(i)
            }
          })
        }
        graph.set('selectedItems', null)
        graph.emit('afterdelete', { items: selectedItems })
      },
      shortcutCodes: ['Delete', 'Backspace']
    })
    cmdPlugin.registerCommand('zoomIn', {
      enable (graph) {
        const zoom = graph.getZoom()
        const maxZoom = graph.get('maxZoom')
        const minZoom = graph.get('minZoom')
        return zoom <= maxZoom && zoom >= minZoom
      },
      execute (graph) {
        const _this = this as any
        const manager = cmdPlugin.get('_command')
        const maxZoom = graph.get('maxZoom')
        const zoom = graph.getZoom()
        _this.originZoom = zoom
        let currentZoom = zoom + manager.zoomDelta
        if (currentZoom > maxZoom) {
          currentZoom = maxZoom
        }
        graph.zoomTo(currentZoom)
      },
      back (graph) {
        const _this = this as any
        graph.zoomTo(_this.originZoom)
      },
      shortcutCodes: [['metaKey', '='], ['ctrlKey', '=']]
    })
    cmdPlugin.registerCommand('zoomOut', {
      enable (graph) {
        const zoom = graph.getZoom()
        const maxZoom = graph.get('maxZoom')
        const minZoom = graph.get('minZoom')
        return zoom <= maxZoom && zoom >= minZoom
      },
      execute (graph) {
        const _this = this as any
        const manager = cmdPlugin.get('_command')
        const minZoom = graph.get('minZoom')
        const zoom = graph.getZoom()
        _this.originZoom = zoom
        let currentZoom = zoom - manager.zoomDelta
        if (currentZoom < minZoom) {
          currentZoom = minZoom
        }
        graph.zoomTo(currentZoom)
      },
      back (graph) {
        const _this = this as any
        graph.zoomTo(_this.originZoom)
      },
      shortcutCodes: [['metaKey', '-'], ['ctrlKey', '-']]
    })
    cmdPlugin.registerCommand('resetZoom', {
      execute (graph) {
        const _this = this as any
        const zoom = graph.getZoom()
        _this.originZoom = zoom
        graph.zoomTo(1)
      },
      back (graph) {
        const _this = this as any
        graph.zoomTo(_this.originZoom)
      }
    })
    cmdPlugin.registerCommand('autoFit', {
      execute (graph) {
        const _this = this as any
        const zoom = graph.getZoom()
        _this.originZoom = zoom
        graph.fitView(5)
      },
      back (graph) {
        const _this = this as any
        graph.zoomTo(_this.originZoom)
      }
    })
    cmdPlugin.registerCommand('toFront', {
      queue: false,
      enable (graph) {
        const items = graph.get('selectedItems')
        return items && items.length > 0
      },
      execute (graph) {
        const items = graph.get('selectedItems')
        if (items && items.length > 0) {
          const item = graph.findById(items[0])
          item.toFront()
          graph.paint()
        }
      }
    })
    cmdPlugin.registerCommand('toBack', {
      queue: false,
      enable (graph) {
        const items = graph.get('selectedItems')
        return items && items.length > 0
      },
      execute (graph) {
        const items = graph.get('selectedItems')
        if (items && items.length > 0) {
          const item = graph.findById(items[0])
          item.toBack()
          graph.paint()
        }
      }
    })
  }

  registerCommand (name: string, cfg: CommandConfig) {
    if (this.cmds[name]) {
      mix(this.cmds[name], cfg)
      return
    }
    const base: any = {}
    const cmd = mix(base, {
      name: name,
      shortcutCodes: [],
      queue: true,
      executeTimes: 1,
      init () {},
      enable () {
        return true
      },
      execute (graph: Graph) {
        const _this = this as any
        _this.snapShot = graph.save()
        _this.selectedItems = graph.get('selectedItems')
        _this.method && (isString(_this.method)) ? graph.get(_this.method)() : _this.method(graph)
      },
      back (graph: Graph) {
        const _this = this as any
        graph.read(_this.snapShot)
        graph.set('selectedItems', _this.selectedItems)
      }
    }, cfg)
    this.cmds[name] = cmd
    this.list.push(cmd)
  }

  execute (name: string, graph: CommandGraph, cfg: any) {
    const cmd = mix({}, this.cmds[name], cfg)
    const manager = this.get('_command')
    if (cmd.enable(graph)) {
      cmd.init()
      if (cmd.queue) {
        manager.queue.splice(manager.current, manager.queue.length - manager.current, cmd)
        manager.current++
      }
      graph.emit('beforecommandexecute', { command: cmd })
      cmd.execute(graph)
      graph.emit('aftercommandexecute', { command: cmd })
      return cmd
    }
  }

  enable (name: string, graph: CommandGraph) {
    return this.cmds[name].enable(graph)
  }

  destroyPlugin (this: any) {
    this._events = null
    this._cfgs = null
    this.list = []
    this.queue = []
    this.destroyed = true
  }
}

export default Command
