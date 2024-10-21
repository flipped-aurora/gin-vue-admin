'use strict'
var attrs: string[] = ['top', 'left', 'right', 'bottom']
var inited: boolean
var elementComputedStyle = {}
var support: string

function getSupport() {
    if(!('CSS' in window) || typeof CSS.supports != 'function') {
        support = ''
    } else if(CSS.supports('top: env(safe-area-inset-top)')) {
        support = 'env'
    } else if(CSS.supports('top: constant(safe-area-inset-top)')) {
        support = 'constant'
    } else {
        support = ''
    }
    return support
}

function init() {
    support = typeof support === 'string' ? support : getSupport()
    if(!support) {
        attrs.forEach((attr: string) => {
            elementComputedStyle[attr] = 0
        })
        return
    }

    function setStyle(el: HTMLElement, style) {
        var elStyle: CSSStyleDeclaration = el.style
        Object.keys(style).forEach(key => {
            var val: string = style[key]
            elStyle[key] = val
        })
    }

    var cbs: Function[] = []
    function parentReady(callback?: Function) {
        if(callback) {
            cbs.push(callback)
        } else {
            cbs.forEach(cb => {
                cb()
            })
        }
    }

    var passiveEvents: any = false
    try {
        var opts = Object.defineProperty({}, 'passive', {
            get: function() {
                passiveEvents = { passive: true }
            }
        })
        window.addEventListener('test', null, opts)
    } catch(e) {

    }

    function addChild(parent: HTMLElement, attr: string) {
        var a1: HTMLElement = document.createElement('div')
        var a2: HTMLElement = document.createElement('div')
        var a1Children: HTMLElement = document.createElement('div')
        var a2Children: HTMLElement = document.createElement('div')
        var W: number = 100
        var MAX: number = 10000
        var aStyle = {
            position: 'absolute',
            width: W + 'px',
            height: '200px',
            boxSizing: 'border-box',
            overflow: 'hidden',
            paddingBottom: `${support}(safe-area-inset-${attr})`
        }
        setStyle(a1, aStyle)
        setStyle(a2, aStyle)
        setStyle(a1Children, {
            transition: '0s',
            animation: 'none',
            width: '400px',
            height: '400px'
        })
        setStyle(a2Children, {
            transition: '0s',
            animation: 'none',
            width: '250%',
            height: '250%'
        })
        a1.appendChild(a1Children)
        a2.appendChild(a2Children)
        parent.appendChild(a1)
        parent.appendChild(a2)

        parentReady(() => {
            a1.scrollTop = a2.scrollTop = MAX
            var a1LastScrollTop: number = a1.scrollTop
            var a2LastScrollTop: number = a2.scrollTop
            function onScroll() {
                if(this.scrollTop === (this === a1 ? a1LastScrollTop : a2LastScrollTop)) {
                    return
                }
                a1.scrollTop = a2.scrollTop = MAX
                a1LastScrollTop = a1.scrollTop
                a2LastScrollTop = a2.scrollTop
                attrChange(attr)
            }
            a1.addEventListener('scroll', onScroll, passiveEvents)
            a2.addEventListener('scroll', onScroll, passiveEvents)
        })

        var computedStyle: CSSStyleDeclaration = getComputedStyle(a1)
        Object.defineProperty(elementComputedStyle, attr, {
            configurable: true,
            get() {
                return parseFloat(computedStyle.paddingBottom)
            }
        })
    }

    var parentDiv: HTMLElement = document.createElement('div')
    setStyle(parentDiv, {
        position: 'absolute',
        left: '0',
        top: '0',
        width: '0',
        height: '0',
        zIndex: '-1',
        overflow: 'hidden',
        visibility: 'hidden',
    })
    attrs.forEach(key => {
        addChild(parentDiv, key)
    })
    document.body.appendChild(parentDiv)
    parentReady()
    inited = true
}

function getAttr(attr: string): number {
    if(!inited) {
        init()
    }
    return elementComputedStyle[attr]
}

var changeAttrs: string[] = []
function attrChange(attr: string) {
    if(!changeAttrs.length) {
        setTimeout(() => {
            var style = {}
            changeAttrs.forEach(attr => {
                style[attr] = elementComputedStyle[attr]
            })
            changeAttrs.length = 0
            callbacks.forEach(callback => {
                callback(style)
            })
        }, 0)
    }
    changeAttrs.push(attr)
}

var callbacks: Function[] = []
function onChange(callback: Function) {
    if(!getSupport()) {
        return
    }
    if(!inited) {
        init()
    }
    if(typeof callback === 'function') {
        callbacks.push(callback)
    }
}

function offChange(callback: Function) {
    var index = callbacks.indexOf(callback)
    if(index >= 0){
        callbacks.splice(index, 1)
    }
}

var safeAreaInsets = {
    get support(): boolean {
        return (typeof support === 'string' ? support : getSupport()).length != 0
    },
    get top(): number {
        return getAttr('top')
    },
    get left(): number {
        return getAttr('left')
    },
    get right(): number {
        return getAttr('right')
    },
    get bottom(): number {
        return getAttr('bottom')
    },
    onChange,
    offChange
}

export = safeAreaInsets
