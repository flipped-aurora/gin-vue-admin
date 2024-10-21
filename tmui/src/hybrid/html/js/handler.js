
// 等待初始化完毕
document.addEventListener('UniAppJSBridgeReady', () => {
	
  document.body.onclick = () =>
    uni.postMessage({
      data: {
        action: 'onClick'
      }
    })
  uni.postMessage({
    data: {
      action: 'onJSBridgeReady'
    }
  })
})
let options
let medias = []

/**
 * @description 获取标签的所有属性
 * @param {Element} ele
 */
function getAttrs (ele) {
  const attrs = Object.create(null)
  for (let i = ele.attributes.length; i--;) {
    attrs[ele.attributes[i].name] = ele.attributes[i].value
  }
  return attrs
}

/**
 * @description 图片加载出错
 */
function onImgError () {
  if (options[1]) {
    this.src = options[1]
    this.onerror = null
  }
  // 取消监听点击
  this.onclick = null
  this.ontouchstart = null
  uni.postMessage({
    data: {
      action: 'onError',
      source: 'img',
      attrs: getAttrs(this)
    }
  })
}

/**
 * @description 检查是否所有图片加载完毕
 */
function checkReady () {
  window.unloadimgs -= 1
  if (window.unloadimgs === 0) {
    // 所有图片加载完毕
    uni.postMessage({
      data: {
        action: 'onReady'
      }
    })
  }
}

/**
 * @description 创建 dom 结构
 * @param {object[]} nodes 节点数组
 * @param {Element} parent 父节点
 * @param {string} namespace 命名空间
 */
function createDom (nodes, parent, namespace) {
  for (let i = 0; i < nodes.length; i++) {
    const node = nodes[i]
    let ele
    if (!node.type || node.type === 'node') {
      let name = node.name
      // svg 需要设置 namespace
      if (name === 'svg') {
        namespace = 'http://www.w3.org/2000/svg'
      }
      if (name === 'html' || name === 'body') {
        name = 'div'
      }
      // 创建标签
      if (!namespace) {
        ele = document.createElement(name)
      } else {
        ele = document.createElementNS(namespace, name)
      }
      // 设置属性
      for (const item in node.attrs) {
        ele.setAttribute(item, node.attrs[item])
      }
      // 递归创建子节点
      if (node.children) {
        createDom(node.children, ele, namespace)
      }

      // 处理图片
      if (name === 'img') {
        window.unloadimgs += 1
        ele.onload = checkReady
        ele.onerror = checkReady
        if (!ele.src && ele.getAttribute('data-src')) {
          ele.src = ele.getAttribute('data-src')
        }
        if (!node.attrs.ignore) {
          // 监听图片点击事件
          ele.onclick = function (e) {
            e.stopPropagation()
            uni.postMessage({
              data: {
                action: 'onImgTap',
                attrs: getAttrs(this)
              }
            })
          }
        }
        if (options[2]) {
          const image = new Image()
          image.src = ele.src
          ele.src = options[2]
          image.onload = function () {
            ele.src = this.src
          }
          image.onerror = function () {
            ele.onerror()
          }
        }
        ele.onerror = onImgError
      } else if (name === 'a') {
        // 处理链接
        ele.addEventListener('click', function (e) {
          e.stopPropagation()
          e.preventDefault() // 阻止默认跳转
          const href = this.getAttribute('href')
          let offset
          if (href && href[0] === '#') {
            offset = (document.getElementById(href.substr(1)) || {}).offsetTop
          }
          uni.postMessage({
            data: {
              action: 'onLinkTap',
              attrs: getAttrs(this),
              offset
            }
          })
        }, true)
      } else if (name === 'video' || name === 'audio') {
        // 处理音视频
        medias.push(ele)
        if (!node.attrs.autoplay && !node.attrs.controls) {
          ele.setAttribute('controls', 'true')
        }
        ele.onplay = function () {
          uni.postMessage({
            data: {
              action: 'onPlay'
            }
          })
          if (options[3]) {
            for (let i = 0; i < medias.length; i++) {
              if (medias[i] !== this) {
                medias[i].pause()
              }
            }
          }
        }
        ele.onerror = function () {
          uni.postMessage({
            data: {
              action: 'onError',
              source: name,
              attrs: getAttrs(this)
            }
          })
        }
      } else if (name === 'table' && options[4] && !ele.style.cssText.includes('inline')) {
        // 处理表格
        const div = document.createElement('div')
        div.style.overflow = 'auto'
        div.appendChild(ele)
        ele = div
      } else if (name === 'svg') {
        namespace = undefined
      }
    } else {
      ele = document.createTextNode(node.text.replace(/&amp;/g, '&'))
    }
    parent.appendChild(ele)
  }
}

// 设置 html 内容
window.setContent = function (nodes, opts, append) {
  const ele = document.getElementById('content')
  document.body.style.overflow = "scroll"
  // 容器样式
  if (opts[0]) {
    document.body.style.cssText = opts[0]
  }


  // 长按复制
  if (!opts[5]) {
    ele.style.userSelect = 'none'
  }

  if (!append) {
    ele.innerHTML = '' // 不追加则先清空
    medias = []
  }

  options = opts
  window.unloadimgs = 0
  const fragment = document.createDocumentFragment()
  createDom(nodes, fragment)
  ele.appendChild(fragment)

  // 触发事件
  let height = ele.scrollHeight
  uni.postMessage({
    data: {
      action: 'onLoad',
      height
    }
  })
  if (!window.unloadimgs) {
    uni.postMessage({
      data: {
        action: 'onReady',
        height
      }
    })
  }

  clearInterval(window.timer)
  window.timer = setInterval(() => {
    if (ele.scrollHeight !== height) {
      height = ele.scrollHeight
      uni.postMessage({
        data: {
          action: 'onHeightChange',
          height: height
        }
      })
    }
  }, 350)
}

// 回收计时器
window.onunload = function () {
  clearInterval(window.timer)
}