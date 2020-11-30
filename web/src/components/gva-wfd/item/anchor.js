import editorStyle from '../util/defaultStyle';
import Item from '@antv/g6/lib/item/item';
import { deepMix } from '@antv/util';

export default class Anchor extends Item {
  constructor(cfg) {
    super(deepMix(cfg,{
      type: 'anchor',
      isActived: false,
      model: {
        type: 'anchor',
        style: {
          ...editorStyle.anchorPointStyle,
          cursor: editorStyle.cursor.hoverEffectiveAnchor,
        }
      },
    }));
    this.enableCapture(true);
    this.isAnchor = true;
    this.toFront();
  }

  showHotpot(){
    this.hotpot = this.getContainer().addShape('marker', {
      attrs: {
        ...this.get('model').style,
        ...editorStyle.anchorHotsoptStyle
      },
      name: 'hotpot-shape',
      draggable: true,
    });
    this.hotpot.toFront();
    this.getKeyShape().toFront();
  }
  setActived(){
    this.update({style: {...editorStyle.anchorPointHoverStyle}});
  }
  clearActived(){
    this.update({style: {...editorStyle.anchorPointStyle}});
  }
  setHotspotActived(act){
    this.hotpot &&
    (act ?
      this.hotpot.attr(editorStyle.anchorHotsoptActivedStyle)
      : this.hotpot.attr(editorStyle.anchorHotsoptStyle))
  }
}
