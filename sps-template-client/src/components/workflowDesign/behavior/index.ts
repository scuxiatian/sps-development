import G6 from '@antv/g6/lib'
import dragPanelItemAddNode from './dragPanelItemAddNode'
import hoverNodeActived from './hoverNodeActived'
import hoverAnchorActived from './hoverAnchorActived'
import clickSelected from './clickSelected'
import dragNode from './dragNode'
import dragEdge from './dragEdge'

export default function (g6: typeof G6) {
  dragPanelItemAddNode(g6)
  hoverNodeActived(g6)
  hoverAnchorActived(g6)
  clickSelected(g6)
  dragNode(g6)
  dragEdge(g6)
}
