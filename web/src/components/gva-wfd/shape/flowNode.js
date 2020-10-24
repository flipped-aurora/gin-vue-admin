import editorStyle from "../util/defaultStyle";

const taskDefaultOptions = {
  icon: null,
  iconStyle: {
    width: 12,
    height: 12,
    left: 2,
    top: 2,
  },
  style:{
    ...editorStyle.nodeStyle,
    fill: '#E7F7FE',
    stroke:'#1890FF',
    cursor: 'default',
  },
  stateStyles: {
    selected: {
      fill: '#95D6FB',
    },
    hover: {
      cursor: editorStyle.cursor.hoverNode,
    }
  }
};

const gatewayDefaultOptions = {
  icon: null,
  iconStyle: {
    width: 20,
    height: 20,
    left: -10,
    top: -10,
  },
  style:{
    ...editorStyle.nodeStyle,
    fill: '#E8FEFA',
    stroke:'#13C2C2',
    cursor: 'default',
  },
  stateStyles: {
    selected: {
      fill: '#8CE8DE',
    },
    hover: {
      cursor: editorStyle.cursor.hoverNode,
    }
  }
};

const startDefaultOptions = {
  icon: null,
  iconStyle: {
    width: 23,
    height: 23,
    left: 16,
    top: 16,
  },
  style:{
    ...editorStyle.nodeStyle,
    fill: '#FEF7E8',
    stroke:'#FA8C16',
    cursor: 'default',
  },
  stateStyles: {
    selected: {
      fill: '#FCD49A',
    },
    hover: {
      cursor: editorStyle.cursor.hoverNode,
    }
  }
};

const endDefaultOptions = {
  icon: null,
  iconStyle: {
    width: 23,
    height: 23,
    left: 16,
    top: 16,
  },
  style:{
    ...editorStyle.nodeStyle,
    fill: '#EFF7E8',
    stroke:'#F5222D',
    cursor: 'default',
  },
  stateStyles: {
    selected: {
      fill: '#CFD49A',
    },
    hover: {
      cursor: editorStyle.cursor.hoverNode,
    }
  }
};

const catchDefaultOptions = {
  icon: null,
  iconStyle: {
    width: 20,
    height: 20,
    left: -10,
    top: -8,
  },
  style:{
    ...editorStyle.nodeStyle,
    fill: '#FEF7E8',
    stroke:'#FA8C16',
    cursor: 'default',
  },
  stateStyles: {
    selected: {
      fill: '#FCD49A',
    },
    hover: {
      cursor: editorStyle.cursor.hoverNode,
    }
  }
};

export default function(G6) {
  G6.registerNode('task-node', {
    shapeType: 'rect',
    options:{
      ...taskDefaultOptions
    },
    getShapeStyle(cfg) {
      cfg.size = [100, 55];
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
    }
  }, 'base-node');
  G6.registerNode('gateway-node', {
    shapeType: 'path',
    labelPosition: 'bottom',
    options:{
      ...gatewayDefaultOptions
    },
    getShapeStyle(cfg) {
      cfg.size = [55, 55];
      const width = cfg.size[0];
      const height = cfg.size[1];
      const gap = 4;
      const style = {
        path: [
          ['M', 0 - gap, 0 - height / 2 + gap],
          ['Q', 0, 0 - height / 2, gap, 0 - height / 2 + gap],
          ['L', width / 2 - gap, 0 - gap],
          ['Q', width / 2, 0, width / 2 - gap, gap],
          ['L', gap, height / 2 - gap],
          ['Q', 0, height / 2, 0 - gap, height / 2 - gap],
          ['L', -width / 2 + gap, gap],
          ['Q', -width / 2, 0, -width / 2 + gap, 0 - gap],
          ['Z']
        ],
        ...this.options.style,
      };
      return style;
    },
  }, 'base-node');
  G6.registerNode('exclusive-gateway-node', {
    options:  G6.Util.deepMix({},gatewayDefaultOptions,{icon: require('../assets/icons/flow/exclusive-gateway.svg')}),
  }, 'gateway-node');
  G6.registerNode('parallel-gateway-node', {
    options:  G6.Util.deepMix({},gatewayDefaultOptions,{icon: require('../assets/icons/flow/parallel-gateway.svg')}),
  }, 'gateway-node');
  G6.registerNode('inclusive-gateway-node', {
    options:  G6.Util.deepMix({},gatewayDefaultOptions,{icon: require('../assets/icons/flow/inclusive-gateway.svg')}),
  }, 'gateway-node');
  G6.registerNode('start-node', {
    shapeType: 'circle',
    labelPosition: 'bottom',
    options:  G6.Util.deepMix({},startDefaultOptions,{icon: require('../assets/icons/flow/start.svg')}),
    getShapeStyle(cfg) {
      cfg.size = [55, 55];
      const width = cfg.size[0];
      const style = {
        x: 0,
        y: 0,
        r: width / 2,
        ...this.options.style,
      };
      if(cfg.hasOwnProperty('color')){
        style.fill = cfg.color
      }
      return style;
    },
    afterDraw(cfg, group) {
      if(cfg.active) {
        const shape = group.get('children')[0];
        shape.animate({
          repeat: true,
          onFrame(ratio) {
            const diff = ratio <=0.5 ? ratio * 10 : (1 - ratio) * 10;
            let radius = cfg.size;
            if (isNaN(radius)) radius = radius[0];
            return {
              r: radius / 2 + diff
            }
          }
        }, 3000, 'easeCubic');
      }
    },
    getAnchorPoints() {
      return [
        [0.5, 0], // top
        [1, 0.5], // right
        [0.5, 1], // bottom
      ]
    }
  }, 'base-node');
  G6.registerNode('end-node', {
    shapeType: 'circle',
    labelPosition: 'bottom',
    options:  G6.Util.deepMix({},endDefaultOptions,{icon: require('../assets/icons/flow/end.svg')}),
    getShapeStyle(cfg) {
      cfg.size = [55, 55];
      const width = cfg.size[0];
      const style = {
        x: 0,
        y: 0,
        r: width / 2,
        ...this.options.style,
      };
      if(cfg.hasOwnProperty('color')){
        style.fill = cfg.color
      }
      return style;
    },
    afterDraw(cfg, group) {
      if(cfg.active) {

      }
    },
    getAnchorPoints() {
      return [
        [0.5, 0], // top
        [0.5, 1], // bottom
        [0, 0.5], // left
      ]
    }
  }, 'base-node');
  G6.registerNode('catch-node', {
    shapeType: 'path',
    labelPosition: 'bottom',
    options: {
      ...catchDefaultOptions
    },
    getShapeStyle(cfg) {
      cfg.size = [70, 55];
      const width = cfg.size[0];
      const height = cfg.size[1];
      const style = {
        path: [
          ['M', 0 , -height/3],
          ['L', width/2, -height/3],
          ['L', 0, height/3*2],
          ['L', -width/2, -height/3],
          ['Z'] // close
        ],
        ...this.options.style,
      };
      return style;
    },
    getAnchorPoints() {
      return [
        [0.5, 0], // top
        [0.8, 0.38], // right
        [0.5, 1], // bottom
        [0.2, 0.38], // left
      ]
    }
  }, 'base-node');
  G6.registerNode('user-task-node', {
    options:  G6.Util.deepMix({},taskDefaultOptions,{
      icon: require('../assets/icons/flow/icon_user.svg'),
      style: {
        fill: '#E7F7FE',
        stroke: '#1890FF',
      },
      stateStyles: {
        selected: {
          fill: '#95D6FB',
        },
      }
    }),
  }, 'task-node');
  G6.registerNode('script-task-node', {
    options:  G6.Util.deepMix({},taskDefaultOptions,{
      icon: require('../assets/icons/flow/icon_script.svg'),
      style: {
        fill: '#FFF7E6',
        stroke: '#FFA940',
      },
      stateStyles: {
        selected: {
          fill: '#FFE7BA',
        },
      }
    }),
  }, 'task-node');
  G6.registerNode('java-task-node', {
    options:  G6.Util.deepMix({},taskDefaultOptions,{
      icon: require('../assets/icons/flow/icon_java.svg'),
      style: {
        fill: '#FFF1F0',
        stroke: '#FF4D4F',
      },
      stateStyles: {
        selected: {
          fill: '#FFCCC7',
        },
      }
    }),
  }, 'task-node');
  G6.registerNode('mail-task-node', {
    options:  G6.Util.deepMix({},taskDefaultOptions,{
      icon: require('../assets/icons/flow/icon_mail.svg'),
      style: {
        fill: '#F6FFED',
        stroke: '#73D13D',
      },
      stateStyles: {
        selected: {
          fill: '#D9F7BE',
        },
      }
    }),
  }, 'task-node');
  G6.registerNode('receive-task-node', {
    options:  G6.Util.deepMix({},taskDefaultOptions,{
      icon: require('../assets/icons/flow/icon_receive.svg'),
      style: {
        fill: '#FFF0F6',
        stroke: '#FF85C0',
      },
      stateStyles: {
        selected: {
          fill: '#FFD6E7',
        },
      }
    }),
  }, 'task-node');
  G6.registerNode('timer-start-node', {
    options:  G6.Util.deepMix({},startDefaultOptions,{icon: require('../assets/icons/flow/icon_timer.svg')}),
    afterDraw(cfg, group) { this.runAnimate(cfg,group) },
  }, 'start-node');
  G6.registerNode('message-start-node', {
    options:  G6.Util.deepMix({},startDefaultOptions,{icon: require('../assets/icons/flow/icon_message.svg')}),
    afterDraw(cfg, group) { this.runAnimate(cfg,group) },
  }, 'start-node');
  G6.registerNode('signal-start-node', {
    options:  G6.Util.deepMix({},startDefaultOptions,{icon: require('../assets/icons/flow/icon_signal.svg')}),
    afterDraw(cfg, group) { this.runAnimate(cfg,group) },
  }, 'start-node');
  G6.registerNode('timer-catch-node', {
    options:  G6.Util.deepMix({},catchDefaultOptions,{icon: require('../assets/icons/flow/icon_timer.svg')}),
  }, 'catch-node');
  G6.registerNode('signal-catch-node', {
    options:  G6.Util.deepMix({},catchDefaultOptions,{icon: require('../assets/icons/flow/icon_signal.svg')}),
  }, 'catch-node');
  G6.registerNode('message-catch-node', {
    options:  G6.Util.deepMix({},catchDefaultOptions,{icon: require('../assets/icons/flow/icon_message.svg')}),
  }, 'catch-node');
}
