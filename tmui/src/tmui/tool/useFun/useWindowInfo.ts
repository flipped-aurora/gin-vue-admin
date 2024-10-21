import { onMounted, reactive } from "vue"
/**
 * 获取当前窗口尺寸信息
 */
export const useWindowInfo = () => {

    let winSize = reactive({
        top:0,
        topSafe: 0,
        width: uni.upx2px(750),
        height: uni.upx2px(750),
        statusBar: 24,
        navigatorBar: 44,
        bottomSafe: 0,
        nvue:false,
        // 是否存在系统自带的状态栏，针对h5优化。
        h5IsSystemBar:true
    })
    // #ifdef APP-NVUE
    winSize.nvue = true;
    // #endif

    function init(){
        let sysinfo = uni.getSystemInfoSync();
        winSize.topSafe = (sysinfo?.statusBarHeight??24) + (sysinfo?.navigationBarHeight??44);
        winSize.width = sysinfo?.windowWidth??uni.upx2px(750);
        winSize.height =  sysinfo?.windowHeight??uni.upx2px(750);
        // #ifdef H5
        winSize.height =  window.innerHeight;
        let ele = document.querySelector(".uni-page-head") as HTMLElement|null;
        if(ele){
            winSize.top = 44
            winSize.h5IsSystemBar = true
            ele.style.zIndex='400'
        }else{
            winSize.h5IsSystemBar = false
        }
        // #endif
        winSize.statusBar =  sysinfo?.statusBarHeight??24;
        winSize.navigatorBar =  sysinfo?.navigationBarHeight??44;
        winSize.bottomSafe =  sysinfo.safeAreaInsets?.bottom??0

    }
    init();
    onMounted(()=>{
        setTimeout(() => {
            init();
        }, 0);
    })

    return winSize;
}