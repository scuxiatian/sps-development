import G6 from '@antv/g6/lib'
import registerEdge from './edge'
import registerFlowNode from './flowNode'
import registerNode from './node'
import registerAnchor from './anchor'

export default function (g6: typeof G6) {
  registerAnchor()
  registerNode(g6)
  registerFlowNode(g6)
  registerEdge(g6)
}
