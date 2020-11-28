/* eslint-disable */
import editorStyle from '../util/defaultStyle';
import { getShapeName } from '../util/clazz';
import ControlPoint from '../item/controlPoint';
const Node = require('@antv/g6/lib/item/node');
const Edge = require('@antv/g6/lib/item/edge');

export default function(G6) {
  G6.registerNode('sub-process-node', {
    shapeType: 'rect',
    options: {
      icon: null,
      iconStyle: {
        width: 12,
        height: 12,
        left: 2,
        top: 2,
      },
      style: {
        ...editorStyle.nodeStyle,
        fill: '#FFFFFF',
        stroke: '#1890FF',
        cursor: 'default',
      },
      stateStyles: {
        selected: {
          fill: '#E7F7FE',
        },
        addNode: {
          fill: '#E7F7FE',
        },
        hover: {
          cursor: editorStyle.cursor.hoverNode,
        },
      },
    },
    drawControlPoints(group) {
      const bbox = group.get('children')[0].getBBox();
      this.getControlPoints().forEach((p, i) => {
        const anchorContainer = group.addGroup();
        let cursor = 'default';
        if (p[0] === 0 && p[1] === 0) {
          cursor = 'nwse-resize';
        } else if (p[0] === 0 && p[1] === 0.5) {
          cursor = 'ew-resize';
        } else if (p[0] === 1 && p[1] === 0) {
          cursor = 'nesw-resize';
        } else if (p[0] === 1 && p[1] === 0.5) {
          cursor = 'ew-resize';
        } else if (p[0] === 1 && p[1] === 1) {
          cursor = 'nwse-resize';
        } else if (p[0] === 0.5 && p[1] === 1) {
          cursor = 'ns-resize';
        } else if (p[0] === 0 && p[1] === 1) {
          cursor = 'nesw-resize';
        } else if (p[0] === 0.5 && p[1] === 0) {
          cursor = 'ns-resize';
        }
        const anchor = new ControlPoint({
          group: anchorContainer,
          index: i,
          model:{
            style:{
              x: bbox.minX + bbox.width * p[0],
              y: bbox.minY + bbox.height * p[1],
              cursor,
            }
          }
        });
        // fuck 坑了好几天，全局注册的东西，在Item内部调用 Shape.getFactory('anchor') 查询不到。
        // 导致 item 初始化不成功。 只能暂时手工在外部把item 初始化了。 有点坑。
        // const shapeFactory = G6.Shape.getFactory('anchor');
        // anchor.set('shapeFactory', shapeFactory);
        // anchor.draw();
        anchor.toFront();
        group.controlPointShapes.push(anchorContainer);
        group.getAllAnchors = () => {
          return group.controlPointShapes;
        };
        group.getAnchor = (j) => {
          return group.controlPointShapes.filter(a => a.get('index') === j);
        };
      });
    },
    drawNodes(cfg, group) {
      const nodes = [];
      if (cfg.content && cfg.content.nodes) {
        cfg.content.nodes.forEach(nodeCfg => {
          const nodeContainer = group.addGroup();
          const node = new Node({
            group: nodeContainer,
            capture: false,
            id: nodeCfg.id,
            groupId: cfg.id,
            model: {
              ...nodeCfg,
              shape: getShapeName(nodeCfg.clazz),
            },
          });
          node.toFront();
          nodeCfg.shape = getShapeName(nodeCfg.clazz);
          nodes.push(node);
        });
      }
      return nodes;
    },
    drawEdges(cfg, group) {
      const edges = [];
      if (cfg.content && cfg.content.edges) {
        cfg.content.edges.forEach(edgeCfg => {
          let source = edgeCfg.source;
          let target = edgeCfg.target;
          if (source && G6.Util.isString(source)) {
            source = group.findById(source);
            source = source.get('item');
          }
          if (target && G6.Util.isString(target)) {
            target = group.findById(target);
            target = target.get('item');
          }
          const edgeContainer = group.addGroup();
          const edge = new Edge({
            group: edgeContainer,
            capture: false,
            source,
            target,
            id: edgeCfg.id,
            model: {
              ...edgeCfg,
              shape: 'flow-polyline-round',
            },
          });
          edge.toFront();
          const model = edge.get('model');
          edgeCfg.startPoint = model.startPoint;
          edgeCfg.endPoint = model.endPoint;
          edgeCfg.shape = 'flow-polyline-round';
          edges.push(edge);
        });
      }
      return edges;
    },
    _addNode(node, cfg) {
      const nodeModel = node.get('model');
      const nodes = nodeModel.content ? nodeModel.content.nodes : [];
      const edges = nodeModel.content ? nodeModel.content.edges : [];
      nodes.push(cfg);
      return { content: { nodes, edges } };
    },
    _addEdge(node, cfg) {
      const nodeModel = node.get('model');
      const nodes = nodeModel.content ? nodeModel.content.nodes : [];
      const edges = nodeModel.content ? nodeModel.content.edges : [];
      edges.push(cfg);
      return { content: { nodes, edges } };
    },
    _updateNode(node, nodeId, cfg) {
      const nodeModel = node.get('model');
      const nodes = nodeModel.content ? nodeModel.content.nodes : [];
      const edges = nodeModel.content ? nodeModel.content.edges : [];
      const tempNode = nodes.find(a => a.id === nodeId);
      Object.assign(tempNode, cfg);
      return { content: { nodes, edges } };
    },
    _getItem(node, itemId) {
      const containerGroup = node.getContainer();
      return containerGroup.nodes.find(item => item.get('id') === itemId);
    },
    _removeItem(node, itemId) {
      const nodeModel = node.get('model');
      if (nodeModel && nodeModel.content) {
        let index = nodeModel.content.nodes.findIndex(item => item.id === itemId);
        if (index !== -1) {
          nodeModel.content.nodes.splice(index, 1);
          return { content: nodeModel.content };
        } else {
          index = nodeModel.content.edges.findIndex(item => item.id === itemId);
          if (index !== -1) {
            nodeModel.content.edges.splice(index, 1);
            return { content: nodeModel.content };
          }
        }
      }
      return {};
    },
    drawShape(cfg, group) {
      const shapeType = this.shapeType;
      const style = this.getShapeStyle(cfg);
      const shape = group.addShape(shapeType, {
        attrs: {
          ...style,
        },
      });

      this.drawIcon(cfg,group);
      this.initAnchor(group);

      const subGroup = group.addGroup({ id: 'sub_' + cfg.id });
      group.subGroup = subGroup;
      // 渲染内部Nodes
      const nodes = this.drawNodes(cfg, subGroup);
      // 渲染内部edges
      const edges = this.drawEdges(cfg, subGroup);

      group.nodes = nodes;
      group.edges = edges;
      group.controlPointShapes = [];
      group.showControlPoints = (_group) => {
        this.drawControlPoints(_group);
      };
      group.clearControlPoints = (_group) => {
        if (_group.controlPointShapes) {
          _group.controlPointShapes.forEach(a => a.remove());
        }
        _group.controlPointShapes = [];
      };
      group.addNodeModel = (_node, _cfg) => {
        return this._addNode(_node, _cfg);
      };
      group.addEdgeModel = (_node, _cfg) => {
        return this._addEdge(_node, _cfg);
      };
      group.updateNodeModel = (_node, _itemId, _cfg) => {
        return this._updateNode(_node, _itemId, _cfg);
      };
      group.getItem = (_node, _itemId) => {
        return this._getItem(_node, _itemId);
      };
      group.removeItem = (_node, _itemId) => {
        return this._removeItem(_node, _itemId);
      };

      return shape;
    },
    setCustomState(name, value, item){
      const group = item.getContainer();
      if (name === 'selected') {
        const rect = group.getChildByIndex(0);
        if (value) {
          group.clearAnchor(group);
          group.showControlPoints(group);
          rect.attr('fill', this.options.stateStyles.selected.fill);
        } else {
          group.clearControlPoints(group);
          rect.attr('fill', this.options.style.fill);
        }
      }
    },
    getControlPoints() {
      return [
        [0, 0],
        [0.5, 0],
        [1, 0],
        [1, 0.5],
        [1, 1],
        [0.5, 1],
        [0, 1],
        [0, 0.5],
      ];
    },
    getShapeStyle(cfg) {
      if (!cfg.size || !Array.isArray(cfg.size)) {
        cfg.size = [80, 44];
      }
      const width = cfg.size[0];
      const height = cfg.size[1];
      const style = {
        x: 0 - width / 2,
        y: 0 - height / 2,
        width,
        height,
        ...this.options.style,
      };
      return style;
    },
    afterUpdate(cfg, node) {
      // TODO 因为没有改变 shape 所以不会更新
      const bbox = node.getBBox();
      const group = node.getContainer();
      const subGroup = group.subGroup;
      if (subGroup) {
        subGroup.clear();
        // 重新渲染内部Nodes
        const nodes = this.drawNodes(cfg, subGroup);
        // 重新渲染内部edges
        const edges = this.drawEdges(cfg, subGroup);
        group.nodes = nodes;
        group.edges = edges;
      }
    },
  }, 'base-node');
}
