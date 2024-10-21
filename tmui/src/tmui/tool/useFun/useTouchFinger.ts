import { ref, reactive } from 'vue';
interface positionType {
    x: number | null,
    y: number | null
}

/**
 * 触摸事件
 * @example
 ```
   import { useTouchFinger } from '@/tmui/tool/useFun/useTouchFinger'
    const {touchstart,touchmove,touchend,touchcancel,addEventListener,direction,deltaXY,preTapPosition,angle,scale } = useTouchFinger();
    <image 
    @touchstart="touchstart"
    @touchmove="touchmove"
    @touchend="touchend"
    @touchcancel="touchcancel"
    src="https://pic.rmb.bdstatic.com/bjh/beautify/aee57799c6885386bb748e07fe43f78f.jpeg" 
    :style="{transform:`scale(${scale}) rotate(${angle}deg)`,width:'100px',height:'100px'}"></image>
    addEventListener('rotate',(evt)=>{
        console.log(direction.value)
    })
 ```
 */
export const useTouchFinger = () => {
    const x1 = ref<number | null>(null)
    const x2 = ref<number | null>(null)
    const y1 = ref<number | null>(null)
    const y2 = ref<number | null>(null)
    //当前点击的位置
    const preTapPosition = reactive<positionType>({ x: null, y: null });
    const preV = reactive<positionType>({ x: null, y: null });
    //计算的距离是相对上一次移动的距离。
    const deltaXY = reactive<positionType>({ x: null, y: null });
    //计算的位置是相对上次触摸开始的位置，位置以当前元素在相对父级左上角开始的位置。
    const positionXY =  reactive<positionType>({ x: 0, y: 0 });

    const pinchStartLen = ref(0);
    const scale = ref(1);
    const angle = ref(0);
    const direction = ref<"" | "Left" | "Right" | "Up" | "Down">("");
    const testScale = ref(0)


    let delta = 0;
    let last = 0;
    let now = 0;
    let isDoubleTap = false;
    let tapTimeout: any = NaN;
    let singleTapTimeout: any = NaN;
    let longTapTimeout: any = NaN;
    let swipeTimeout: any = NaN;
    const zoomFactor = 0.55; // 缩放速率
    const zoomFactorAb = 0.03; // 旋转速率
    // 点按
    let tap = (evt: TouchEvent) => { }
    //双指捏合
    let pinch = (evt: TouchEvent) => { }
    //双指撮合旋转
    let rotate = (evt: TouchEvent) => { }
    let touchStart = (evt: TouchEvent) => { }
    let touchMove = (evt: TouchEvent) => { }
    let touchEnd = (evt: TouchEvent) => { }
    let touchCancel = (evt: TouchEvent) => { }
    //手指按住移动
    let pressMove = (evt: TouchEvent) => { }
    //双击
    let doubleTap = (evt: TouchEvent) => { }
    //长按
    let longTap = (evt: TouchEvent) => { }
    //手指滑动
    let swipe = (evt: TouchEvent) => { }
    //单击
    let singleTap = (evt: TouchEvent) => { }
    //一个手指以上触发
    let multipointStart = (evt: TouchEvent) => { }
    //一个手指以上触发结束
    let multipointEnd = (evt: TouchEvent) => { }

    function start(evt: TouchEvent) {
        if (!evt.touches) return;
        touchStart(evt)
        now = Date.now();
        x1.value = evt.touches[0].pageX;
        y1.value = evt.touches[0].pageY;
        delta = now - (last || now);
        if(x1.value===null) x1.value = 0
        if(y1.value===null) y1.value = 0
        if (preTapPosition.x !== null && preTapPosition.y !== null) {
            isDoubleTap = (delta > 0 && delta <= 250 && Math.abs(preTapPosition.x - x1.value) < 30 && Math.abs(preTapPosition.y - y1.value) < 30);
        }
        preTapPosition.x = x1.value;
        preTapPosition.y = y1.value;
        last = now;
        let len = evt.touches.length;

        if (len > 1) {
            cancelLongTap();
            cancelSingleTap();
            
            let otx = evt.touches[1].pageX;
            let oty = evt.touches[1].pageY;
            let v = { x: otx - x1.value, y: oty - y1.value };
            preV.x = v.x;
            preV.y = v.y;
           
            pinchStartLen.value = getLen(preV);
            multipointStart(evt)
        }
        longTapTimeout = setTimeout(function () {
            longTap(evt);
        }, 750);


    }
    function move(evt: TouchEvent) {
        if (!evt.touches) return;
        let len = evt.touches.length,
            currentX = evt.touches[0].pageX,
            currentY = evt.touches[0].pageY;
        isDoubleTap = false;
        if (len > 1) {
            let otx = evt.touches[1].pageX;
            let oty = evt.touches[1].pageY;
            let v = { x: otx - currentX, y: oty - currentY };

            if (preV.x !== null) {
                if (pinchStartLen.value > 0) {
                    let temsc = (getLen(v) / pinchStartLen.value);
                    // 计算缩放比例
                    const deltaScale = (temsc - 1) * zoomFactor + scale.value;
                    scale.value = Math.max(deltaScale,0.1)
                   
                    pinch(evt);
                }
                let testjd = getRotateAngle(v, preV);
                angle.value = Math.floor((testjd - 1) * zoomFactorAb) + angle.value;
                pinchStartLen.value = getLen(v);
                rotate(evt);
            }
            preV.x = v.x;
            preV.y = v.y;
        } else {
            if (x2.value !== null && y2.value !== null) {
                deltaXY.x = currentX - x2.value;
                deltaXY.y = currentY - y2.value;

            } else {
                deltaXY.x = 0;
                deltaXY.y = 0;
            }
            pressMove(evt);
        }

        touchMove(evt);

        cancelLongTap();
        x2.value = currentX;
        y2.value = currentY;
        if (len > 1) {
            evt?.preventDefault();
        }

    }
    function end(evt: TouchEvent) {
        if (!evt.changedTouches) return;
        cancelLongTap();
        if (evt.touches.length < 2) {
            multipointEnd(evt);
        }
        touchEnd(evt);
        if(x1.value===null) x1.value = 0
        if(x2.value===null) x2.value = 0
        if(y1.value===null) y1.value = 0
        if(y2.value===null) y2.value = 0
        //swipe
        if ((x2.value && Math.abs(x1.value - x2.value) > 30) ||
            (y2.value && Math.abs(y1.value - y2.value) > 30)) {
            direction.value = swipeDirection(x1.value, x2.value, y1.value, y2.value);
            swipeTimeout = setTimeout(function () {
                swipe(evt);
            }, 0)
        } else {
            tapTimeout = setTimeout(function () {
                tap(evt);
                if (isDoubleTap) {
                    doubleTap(evt);
                    scale.value = 1;
                    angle.value = 0
                    clearTimeout(singleTapTimeout);
                    isDoubleTap = false;
                }
            }, 0)
            if (!isDoubleTap) {
                singleTapTimeout = setTimeout(function () {
                    singleTap(evt);
                }, 250);
            }
        }

        preV.x = 0;
        preV.y = 0;
        // scale.value = 1;
        pinchStartLen.value = 0;
        x1.value = x2.value = y1.value = y2.value = null;
    }

    function bind(
        type: "tap" | "pinch" | "rotate" | "touchStart" | "touchMove" | "touchEnd" | "touchCancel"
            | "pressMove" | "doubleTap" | "longTap" | "swipe" | "singleTap" | "multipointStart" | "multipointEnd",
        call: (evt: TouchEvent) => void
    ) {
        if (type == 'tap') tap = call;
        if (type == 'pinch') pinch = call;
        if (type == 'rotate') rotate = call;
        if (type == 'touchStart') touchStart = call;
        if (type == 'touchMove') touchMove = call;
        if (type == 'touchEnd') touchEnd = call;
        if (type == 'touchCancel') touchCancel = call;
        if (type == 'pressMove') pressMove = call;
        if (type == 'doubleTap') doubleTap = call;
        if (type == 'longTap') longTap = call;
        if (type == 'swipe') swipe = call;
        if (type == 'singleTap') singleTap = call;
        if (type == 'multipointStart') multipointStart = call;
        if (type == 'multipointEnd') multipointEnd = call;
    }
    function swipeDirection(x1: number, x2: number, y1: number, y2: number) {
        return Math.abs(x1 - x2) >= Math.abs(y1 - y2) ? (x1 - x2 > 0 ? 'Left' : 'Right') : (y1 - y2 > 0 ? 'Up' : 'Down')
    }

    function getLen(v: positionType) {
        if(v.x===null||v.y===null) return 0

        return Math.hypot(v.x,v.y)


        // return Math.sqrt(v.x * v.x + v.y * v.y);
    }

    function dot(v1: positionType, v2: positionType) {
        if(v1.x===null||v1.y===null||v2.x===null||v2.y===null) return 0
        return v1.x * v2.x + v1.y * v2.y;
    }

    function getAngle(v1: positionType, v2: positionType) {
        
        let mr = getLen(v1) * getLen(v2);
        if (mr === 0) return 0;
        let r = dot(v1, v2) / mr;
        if (r > 1) r = 1;
        let jd =  Math.acos(r);
        jd = jd * (180 / Math.PI);
        return (jd + 360) % 360;;
    }

    function cross(v1: positionType, v2: positionType) {
        if(v1.x===null||v1.y===null||v2.x===null||v2.y===null) return 0
        return v1.x * v2.y - v2.x * v1.y;
    }

    function getRotateAngle(v1: positionType, v2: positionType) {
        let angle = getAngle(v1, v2);
        if (cross(v1, v2) > 0) {
            angle *= -1;
        }

        return angle * 180 / Math.PI;
    }


    function cancel(evt: TouchEvent) {
        clearTimeout(singleTapTimeout);
        clearTimeout(tapTimeout);
        clearTimeout(longTapTimeout);
        clearTimeout(swipeTimeout);
        touchCancel(evt);

    }

    function cancelLongTap() {
        clearTimeout(longTapTimeout);
    }

    function cancelSingleTap() {
        clearTimeout(singleTapTimeout);
    }


    return {
        
        touchstart: start, 
        touchmove: move, 
        touchend: end, 
        touchcancel: cancel, 
        addEventListener: bind,
        deltaXY, 
        pinchStartLen, 
        scale, 
        angle, 
        direction, 
        preTapPosition,
        testScale
    }
}