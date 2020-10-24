import { each } from '@antv/util';
import { vec2 } from '@antv/matrix-util';
export default function(G6){
  const {mix} = G6.Util;
  G6.registerBehavior('itemAlign', {
    getDefaultCfg() {
      return {
        alignLineStyle: { stroke: '#FA8C16', lineWidth: 1 },
        tolerance: 5,
        _alignLines: [],
      };
    },
    getEvents() {
      return {
        'afternodedrag': 'onDrag',
        'afternodedragend': 'onDragEnd',
      };
    },
    onDrag(shape) {
      this._clearAlignLine();
      this._itemAlign(shape);
    },
    onDragEnd() {
      this._clearAlignLine();
    },
    _itemAlign(item){
      const bbox = item.getBBox();
      const ct = { x: bbox.x + bbox.width / 2, y: bbox.y };
      const cc = { x: bbox.x + bbox.width / 2, y: bbox.y + bbox.height / 2 };
      const cb = { x: bbox.x + bbox.width / 2, y: bbox.y + bbox.height };
      const lc = { x: bbox.x, y: bbox.y + bbox.height / 2 };
      const rc = { x: bbox.x + bbox.width, y: bbox.y + bbox.height / 2 };
      const nodes = this.graph.getNodes();
      each(nodes, (node) => {
        const horizontalLines = [];
        const verticalLines = [];
        let p = null;
        const bbox1 = node.getBBox();
        each(this.getHorizontalLines(bbox1), (line) => {
          horizontalLines.push(this.getDistance(line,ct));
          horizontalLines.push(this.getDistance(line,cc));
          horizontalLines.push(this.getDistance(line,cb));
        });
        each(this.getVerticalLines(bbox1), (line) => {
          verticalLines.push(this.getDistance(line,lc));
          verticalLines.push(this.getDistance(line,cc));
          verticalLines.push(this.getDistance(line,rc));
        });
        horizontalLines.sort((a,b)=> a.dis-b.dis);
        verticalLines.sort((a,b)=> a.dis-b.dis);
        if(horizontalLines.length > 0 && horizontalLines[0].dis < this.tolerance){
          item.attr({y: horizontalLines[0].line[1] - horizontalLines[0].point.y + bbox.y });
          p = { horizontals: [horizontalLines[0]] };
          for (let i = 1; i < 3; i++) horizontalLines[0].dis === horizontalLines[i].dis && p.horizontals.push(horizontalLines[i]);
        }
        if(verticalLines.length > 0 && verticalLines[0].dis < this.tolerance){
          item.attr({ x: verticalLines[0].line[0] - verticalLines[0].point.x + bbox.x });
          p ? p.verticals = [verticalLines[0]] : p = { verticals: [verticalLines[0]] };
          for (let i = 1; i < 3; i++) verticalLines[0].dis === verticalLines[i].dis && p.verticals.push(verticalLines[i]);
        }
        if(p){
          p.bbox = bbox;
          this._addAlignLine(p);
        }
      })
    },
    _addAlignLine(p){
      const group = this.graph.get('group');
      const bbox = p.bbox;
      const lineStyle = this.alignLineStyle;
      const lineArr = this._alignLines;
      if(p.horizontals){
        each(p.horizontals, function(lineObj) {
          const line = lineObj.line;
          const point = lineObj.point;
          const lineHalf = (line[0] + line[2]) / 2;
          let x1,x2;
          if(point.x < lineHalf){
            x1 = point.x - bbox.width / 2;
            x2 = Math.max(line[0], line[2]);
          }else{
            x1 = point.x + bbox.width / 2;
            x2 = Math.min(line[0], line[2]);
          }
          const shape = group.addShape('line', { attrs: mix({ x1, y1: line[1], x2, y2: line[1] }, lineStyle), capture: false });
          lineArr.push(shape);
        })
      }
      if(p.verticals){
        each(p.verticals, function(lineObj) {
          const line = lineObj.line;
          const point = lineObj.point;
          const lineHalf = (line[1] + line[3]) / 2;
          let y1,y2;
          if(point.y < lineHalf){
            y1 = point.y - bbox.height / 2;
            y2 = Math.max(line[1], line[3]);
          }else{
            y1 = point.y + bbox.height / 2;
            y2 = Math.min(line[1], line[3]);
          }
          const shape = group.addShape('line', { attrs: mix({ x1: line[0], y1, x2: line[0], y2 }, lineStyle), capture: false });
          lineArr.push(shape);
        })
      }
    },
    getHorizontalLines(bbox){
      return [
        [bbox.minX, bbox.minY, bbox.maxX, bbox.minY],       // tltr
        [bbox.minX, bbox.centerY, bbox.maxX, bbox.centerY], // lcrc
        [bbox.minX, bbox.maxY, bbox.maxX, bbox.maxY],       // blbr
      ]
    },
    getVerticalLines(bbox){
      return [
        [bbox.minX, bbox.minY, bbox.minX, bbox.maxY],       // tlbl
        [bbox.centerX, bbox.minY, bbox.centerX, bbox.maxY], // tcbc
        [bbox.maxX, bbox.minY, bbox.maxX, bbox.maxY],       // trbr
      ]
    },
    getDistance(line, point) {
      return { line, point, dis: this.pointLineDistance(line[0], line[1], line[2], line[3], point.x, point.y) };
    },
    pointLineDistance: function(lineX1, lineY1, lineX2, lineY2, pointX, pointY) {
      const lineLength = [lineX2 - lineX1, lineY2 - lineY1];
      if (vec2.exactEquals(lineLength, [0, 0])) return NaN;
      let s = [-lineLength[1], lineLength[0]];
      vec2.normalize(s, s);
      return Math.abs(vec2.dot([pointX - lineX1, pointY - lineY1], s));
    },
    _clearAlignLine(){
      each(this._alignLines, (line) => {
        line.remove();
      });
      this._alignLines = [];
      this.graph.paint();
    },

  });
}
