function sendErrorTip(errorInfo) {
  console.groupCollapsed(`捕获到错误: ${errorInfo.type}`);
  console.log('错误类型:', errorInfo.type);
  console.log('错误信息:', errorInfo.message);
  console.log('调用栈:', errorInfo.stack);
  if (errorInfo.component) {
    console.log('组件名:', errorInfo.component.name)
    console.log('组件地址:', errorInfo.component.__file)
  }
  if (errorInfo.vueInfo) console.log('Vue 信息:', errorInfo.vueInfo);
  if (errorInfo.source) console.log('来源文件:', errorInfo.source);
  if (errorInfo.lineno) console.log('行号:', errorInfo.lineno);
  if (errorInfo.colno) console.log('列号:', errorInfo.colno);
  console.groupEnd();
}

function initVueErrorHandler(app) {
  app.config.errorHandler = (err, vm, info) => {
    let errorType = 'Vue Error';
  
    sendErrorTip({
      type: errorType,
      message: err.message,
      stack: err.stack,
      component: vm.$options || 'Unknown Vue Component',
      vueInfo: info
    });
  };
}

function initJsErrorHandler() {
  window.onerror = (message, source, lineno, colno, error) => {
    let errorType = 'JS Error';
  
    sendErrorTip({
      type: errorType,
      message: message,
      stack: error ? error.stack : 'No stack available',
      source: source,
      lineno: lineno,
      colno: colno
    });
  
    return false;
  };
}

export function initErrorHandler(app) {
  initVueErrorHandler(app)
  initJsErrorHandler()
}
