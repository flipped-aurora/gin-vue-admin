export default function(G6) {
  G6.registerBehavior('dragPoint', {
    getDefaultCfg() {
      return {
        updateEdge: true,
        delegate: true,
        delegateStyle: {},
        dragEdge: false,
      };
    },
    getEvents() {
      return {
        'controlPoint:dragstart': 'onDragStart',
        'controlPoint:drag': 'onDrag',
        'controlPoint:dragend': 'onDragEnd',
      };
    },
    onDragStart(e) {
      const node = e.target.getParent().getParent().get('item');
      const anchorIndex = e.item.get('index');
      this.target = e.item;
      this.origin = {
        x: e.x,
        y: e.y,
        sourceNode: node,
        sourceAnchor: anchorIndex,
      };
      this.graph.set('edgeDragging', true);
    },
    onDrag(e) {
      if (!this.origin) {
        return;
      }

      const node = e.target.getParent().getParent().get('item');
      const anchorIndex = e.item.get('index');

      const model = node.getModel();
      const currentWidth = model.size[0];
      const currentHeight = model.size[1];
      const addWidth = e.x - this.origin.x;
      const addHeight = e.y - this.origin.y;
      let width = currentWidth;
      let height = currentHeight;
      // 0,0 // 两个都是负的变大
      // 1,0 // 一正 一负 变大
      // 1,1 // 两个都是负变大
      // 0,1 // 一负 一正 变大
      const pointIndex = this.origin.sourceAnchor;
      if (pointIndex === 0) {
        width = currentWidth - addWidth;
        height = currentHeight - addHeight;
      } else if (pointIndex === 1) {
        height = currentHeight - addHeight;
      } else if (pointIndex === 2) {
        width = currentWidth + addWidth;
        height = currentHeight - addHeight;
      } else if (pointIndex === 3) {
        width = currentWidth + addWidth;
      } else if (pointIndex === 4) {
        width = currentWidth + addWidth;
        height = currentHeight + addHeight;
      } else if (pointIndex === 5) {
        height = currentHeight + addHeight;
      } else if (pointIndex === 6) {
        width = currentWidth - addWidth;
        height = currentHeight + addHeight;
      } else if (pointIndex === 7) {
        width = currentWidth - addWidth;
      }
      const group = node.getContainer();
      // 移动过程中的控制点全部隐藏。
      group.controlPointShapes.forEach(a => a.hide());

      this.graph.updateItem(node, { size: [width, height], });
      this.origin = {
        x: e.x,
        y: e.y,
        sourceNode: node,
        sourceAnchor: anchorIndex,
      };

    },
    onDragEnd(e) {
      if (!this.origin) {
        return;
      }
      const node = e.target.getParent().getParent().get('item');
      this.target = null;
      this.origin = null;
      const group = node.getContainer();
      group.clearControlPoints(group);
      group.showControlPoints(group);
      this.graph.set('edgeDragging', false);
    },
  });
}
