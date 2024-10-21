export default class WxCanvas {
  constructor(ctx, canvasId, isNew, canvasNode) {
    this.ctx = ctx;
    this.canvasId = canvasId;
    this.chart = null;
    this.isNew = isNew
    if (isNew) {
      this.canvasNode = canvasNode;
    }
    else {
      this._initStyle(ctx);
    }

    // this._initCanvas(zrender, ctx);

    this._initEvent();
  }

  getContext(contextType) {
    if (contextType === '2d') {
      return this.ctx;
    }
  }

  // canvasToTempFilePath(opt) {
  //   if (!opt.canvasId) {
  //     opt.canvasId = this.canvasId;
  //   }
  //   return wx.canvasToTempFilePath(opt, this);
  // }

  setChart(chart) {
    this.chart = chart;
  }

  attachEvent() {
    // noop
  }

  detachEvent() {
    // noop
  }

  _initCanvas(zrender, ctx) {
    zrender.util.getContext = function () {
      return ctx;
    };

    zrender.util.$override('measureText', function (text, font) {
      ctx.font = font || '12px sans-serif';
      return ctx.measureText(text);
    });
  }

  _initStyle(ctx) {
    ctx.createRadialGradient = () => {
      return ctx.createCircularGradient(arguments);
    };
  }

  _initEvent() {
    this.event = {};
    const eventNames = [{
      wxName: 'touchStart',
      ecName: 'mousedown'
    }, {
      wxName: 'touchMove',
      ecName: 'mousemove'
    }, {
      wxName: 'touchEnd',
      ecName: 'mouseup'
    }, {
      wxName: 'touchEnd',
      ecName: 'click'
    }];

    eventNames.forEach(name => {
      this.event[name.wxName] = e => {
        const touch = e.touches[0];
        this.chart.getZr().handler.dispatch(name.ecName, {
          zrX: name.wxName === 'tap' ? touch.clientX : touch.x,
          zrY: name.wxName === 'tap' ? touch.clientY : touch.y
        });
      };
    });
  }
  addEventListener(el, name, handler, opt) {
    // el?.addEventListener(name, handler, opt);
  }
  set width(w) {
    if (this.canvasNode) this.canvasNode.width = w
  }
  set height(h) {
    if (this.canvasNode) this.canvasNode.height = h
  }

  get width() {
    if (this.canvasNode)
      return this.canvasNode.width
    return 0
  }
  get height() {
    if (this.canvasNode)
      return this.canvasNode.height
    return 0
  }
}