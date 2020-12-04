import editorStyle from "../util/defaultStyle";

export default function(G6){
  G6.registerBehavior('dragNode', {
    getDefaultCfg() {
      return {
        updateEdge: true,
        delegate: true,
        delegateStyle: {},
        align: true,
      };
    },
    getEvents() {
      return {
        'node:dragstart': 'onDragStart',
        'node:drag': 'onDrag',
        'node:dragend': 'onDragEnd'
      };
    },
    onDragStart(e) {
      if (!this.shouldBegin.call(this, e)) {
        return;
      }
      this.target = e.item;
      this.origin = {
        x: e.x,
        y: e.y
      };
    },
    onDrag(e) {
      if (!this.origin) {
        return;
      }
      if (!this.get('shouldUpdate').call(this, e)) {
        return;
      }
      const origin = this.origin;
      const groupId = this.target.get('groupId');
      const model = this.target.get('model');
      if (!this.point) {
        this.point = {
          x: model.x,
          y: model.y
        };
      }
      if (groupId) {
        const subProcessNode = this.graph.findById(groupId);
        const subProcessBBox = subProcessNode.getBBox();
        const x = e.x - origin.x + this.point.x + subProcessBBox.x + subProcessBBox.width / 2;
        const y = e.y - origin.y + this.point.y + subProcessBBox.y + subProcessBBox.height / 2;
        this.origin = { x: e.x, y: e.y };
        this.point = {
          x: x - subProcessBBox.x - subProcessBBox.width / 2,
          y: y - subProcessBBox.y - subProcessBBox.height / 2
        };
        if (this.delegate) {
          this._updateDelegate(this.target, x, y);
        }
      } else {
        const x = e.x - origin.x + this.point.x;
        const y = e.y - origin.y + this.point.y;
        this.origin = { x: e.x, y: e.y };
        this.point = { x, y };
        if (this.delegate) {
          this._updateDelegate(this.target, x, y);
        }
      }
    },
    onDragEnd(e) {
      if (!this.shouldEnd.call(this, e)) {
        return;
      }
      if (!this.origin) {
        return;
      }
      const delegateShape = e.item.get('delegateShape');
      const groupId = this.target.get('groupId');
      if (groupId) {
        if (delegateShape) {
          const subProcessNode = this.graph.findById(groupId);
          const subProcessBBox = subProcessNode.getBBox();
          const bbox = delegateShape.getBBox();
          const x = bbox.x + bbox.width / 2 - subProcessBBox.x - subProcessBBox.width / 2;
          const y = bbox.y + bbox.height / 2 - subProcessBBox.y - subProcessBBox.height / 2;
          delegateShape.remove();
          this.target.set('delegateShape', null);
          const group = subProcessNode.getContainer();
          const id = this.target.get('id');
          const resultModel = group.updateNodeModel(subProcessNode, id, { x, y });
          this._updateItem(subProcessNode, resultModel);
        }
      } else {
        if (delegateShape) {
          const bbox = delegateShape.getBBox();
          const x = bbox.x + bbox.width / 2;
          const y = bbox.y + bbox.height / 2;
          delegateShape.remove();
          this.target.set('delegateShape', null);
          this._updateItem(this.target, { x, y });
        }
      }
      this.point = null;
      this.origin = null;
      this.graph.emit('afternodedragend', this.target);
    },
    _updateItem(item, point) {
      if(this.graph.executeCommand) {
        this.graph.executeCommand('update', {
          itemId: item.get('id'),
          updateModel: point
        });
      }else {
        if (this.get('updateEdge')) {
          this.graph.updateItem(item, point);
        } else {
          item.updatePosition(point);
          this.graph.paint();
        }
      }
    },
    _updateDelegate(item, x, y) {
      const self = this;
      let shape = item.get('delegateShape');
      const bbox = item.get('keyShape').getBBox();
      if (!shape) {
        const parent = self.graph.get('group');
        const attrs = editorStyle.nodeDelegationStyle;
        // model上的x, y是相对于图形中心的，delegateShape是g实例，x,y是绝对坐标
        shape = parent.addShape('rect', {
          attrs: {
            width: bbox.width,
            height: bbox.height,
            x: x - bbox.width / 2,
            y: y - bbox.height / 2,
            nodeId:item.get('id'),
            ...attrs
          }
        });
        shape.set('capture', false);
        item.set('delegateShape', shape);
      }
      shape.attr({ x: x - bbox.width / 2, y: y - bbox.height / 2 });
      this.graph.paint();
      this.graph.emit('afternodedrag',shape);
    },
  });
}
