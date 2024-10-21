import 'pinia';
import { ComponentInternalInstance, ComponentPublicInstance } from "vue"
interface Data {
    [key: string]: any;
}

/**渐变方向 left:右->左，right:左->右。top:下->上，bottom:上->下。 */
export type linearType = 'left' | 'right' | 'bottom' | 'top' | ''
/**渐变色调，light,dark,accent亮系渐变和深色渐变。 */
export type linearDeepType = "accent" | "dark" | "light"

/**边线的访问 */
export type borderDirectionType = "all" | "bottom" | "bottomleft" | "bottomright" | "left" | "leftright" | "right" | "right" | "top" | "topbottom" | "topleft" | "topright" | "x" | "y"
/**边线的样式 */
export type borderStyleType = "solid" | "dashed" | "dotted"

type componentKey = 'app' | 'button' | 'card' | 'tag' | 'sheet'
/**定义对象字段的键值只能是规定的组件名称。 */
type componentKeys = Partial<Record<componentKey, Tmui.components.all>>


declare global {
    interface Uni {
        /**
         * tmui3.0 函数库
         * https://tmui.design
         */
        $tm: tmUtil,
    }
    namespace UniNamespace {
        interface CanvasToTempFilePathOptions {
            // @ts-ignore
            canvas?: CanvasContext
        }
        interface CanvasContext {
            width: number,
            height: number,
            /**dpr,屏幕的倍率 */
            dpr: number,
            /**nvue的gcanvas函数 */
            toTempFilePath(any0: any, any1: any, any2: any, any3: any, any4: any, any5: any, any6: any, any7: any, any8: any): void
        }
        interface NodesRef {
            /**节点，微信小程序上 */
            node(callback?: (result: any) => void): SelectorQuery,
            /**需要读取的字段属性 */
            fields(fields: NodeField, callback?: (result: NodeInfo) => void): SelectorQuery;
        }
        interface NodeField {
            /**包含的节点，部分小程序上会返回。 */
            node?: boolean
        }
    }

    namespace Tmui {
		interface fetchNetConfigType {
			/**
			 * 请求的数据，query,或者body，都可传递对象。自动转换数据格式。
			 */
			data?:any,
			/**
			 * 头部数据。
			 */
			header?:any;
			/**
			 * 请求方式
			 */
			method?:fetchConfigMethod,
			/** 
			* 定义成功的服务器返回的状态码成功的标志
			*/
			statusCode?:number,
			/** 
			 * 从返回的结果读取的数据字段，默认全部返回，
			 * 如果指定了就会读取指定字段
			 * 比如你的结果返回{code,data,msg},如果你指定：['data'],结果集中只返回data数据
			 * 如果是['data',code],结果集中只返回{data,code}
			 */
			pick?:string[],
			/**
			 * 是否监视参数的变化，如果参数发生了变化将会重新发起请求
			 */
			watchRefresh?:boolean,
			/**
			 * 如果不配置或者配置,默认为all
			 */
			toast?:'fail'|'success'|'all',
			/**
			 * 是否显示操作后的提示,默认为true
			 */
			showToast?:boolean,
			/**
			 * 读取提示信息的字段，如果读取不到会显示自动的文字。默认为msg
			 */
			toastKey?:string,
			/**
			 * 是否显示加载框,默认为true。
			 */
			showLoading?:boolean,
		}
		
        /**tmui配置表 */
        interface tmuiConfig {
            /** 自动跟随系统暗黑 */
            autoDark?: boolean,
            /** 开启全局分离功能，默认关闭 */
            shareDisable?: boolean,
            /**主题列表 */
            theme?: {},
            /**细化全局的主题配置表 */
            themeConfig?: {
                /**暗黑模式下的一些统一配置 */
                dark?: {
                    /**一般的卡片项目暗黑背景 */
                    cardColor?: string,
                    /**输入框，表单等暗黑背景 */
                    inputColor?: string,
                    /**禁用输入框，表单等暗黑背景 */
                    disableColor?: string,
                    /**暗黑下的页面背景 */
                    bodyColor?: string,
                    /**文本禁用色. */
                    textDisableColor?: string,
                },
                /**
                 * app整体字体大小的调整比例，
                 * 只对使用tm-text组件以及自身组件的字号才会有效果 
                 * 
                 * */
                globalFontSizeRatio?: number,
                /** 是否关闭弹层背景的模糊 */
                overflowBlur?: false,
                /**
                 * 针对不同的主题配置详细的配色方案。
                 * 注意这里影响的时主题计算功能的配置
                 */
                theme?: {
                    /**（如果该组件默认有值就会使用下列属性，如果组件默认为0不会使用。） */
                    [key: string]: {
                        /**待考虑编辑2023-2-18 00:23:35 */
                    }
                },
                /**各个组件的统一配置 */
                component?: componentKeys
            },
            /**router路由拦截代替外置文件 */
            router?: {
                /**页面访问前执行 */
                useTmRouterBefore(arg: beforeRouterOpts): void,
                /**页面访问后执行 */
                useTmRouterAfter(arg: beforeRouterOpts): void,
            },
            /**用户自定义全局数据 */
            custom?: {
                [key: string]: any
            }
        }
        interface beforeRouterOpts {
            path: string | null,//当前页面路径，不含前缀 /
            opts?: any,//页面参数
            openType?: string,//当前页面打开的类型
            context: ComponentPublicInstance | null,
        }
        /**actionMenu组件项目类型 */
        interface tmActionMenu {
            text?: string,
            disabled?: boolean,
            /** 各家小程序的openType。 */
            openType?: string
            [key: string]: any;
        }
        interface tmAlert {
            icon?: string,
            title?: string,
            content?: string
        }
        interface tabs {
            key?: string | number,
            title?: string,
            icon?: string,
            dot?: boolean,
            count?: string | number
            dotColor?: string,
            [key: string]: any
        }
        interface skuItem {
            title: string,
            id: string | number,
            num: number,
            children: skuItem[]
        }
        interface sku {
            data: skuItem[],
            product: {
                id: string,
                title: string,
                num: number,
                max_buy: number,
                /** 原价 */
                price: number,
                /** 优惠价 */
                salePrice: 54,
                tip: string,
                img: string
            }[]
        }
		
        interface tmFormSubmitResult {
            data: { [key: string]: any };
            isPass: boolean,
            result: {
                message: string,//校验后的提示文本
                validator: boolean,//是否校验通过
            }[]
        }
        interface tmFormRules {
            validator?: Function | boolean,//检验函数。可以是Promise异步回调。
            required?: boolean,//是否必填。
            message?: string,//检验不合格时的文本
            type?: string,//校验类型.
            [key: string]: any
        }
        /**组件的配置 */
        namespace components {
            type all = button & sheet
            interface button {
                round?: number,
                shadow?: number,
                color?: string
            }
            interface sheet {

            }
        }


    }

}
declare module 'pinia' {
    export interface PiniaCustomProperties {
        tmuiConfig: Tmui.tmuiConfig,
    }
    export interface PiniaCustomStateProperties {
        tmuiConfig: Tmui.tmuiConfig
    }
}


export type pagesCustomType = 'default' | 'custom'
export interface pagesType {
    //页面地址
    path: string,
    //导航栏模式
    custom: pagesCustomType,
    navigationBarBackgroundColor: string,
    navigationBarTextStyle: string,
    subPackages?: Array<any>
}
export interface tabBarItemType {
    pagePath: string,
    iconPath: string,
    selectedIconPath: string,
    text: string
}
export interface tabBarType {
    color: string,
    selectedColor: string,
    borderStyle: string,
    backgroundColor: string,
    list?: Array<tabBarItemType>
}
export interface beforeRouterOpts {
    path: string | null,//当前页面路径，不含前缀 /
    opts?: any,//页面参数
    openType?: string,//当前页面打开的类型
    context: ComponentPublicInstance | null,
}
type fetchConfigResponseType = "arraybuffer" | "text";
type fetchConfigDataType = "json" | "text";
type fetchConfigMethod = "GET" | "POST" | "PUT" | "DELETE" | "CONNECT" | "HEAD" | "OPTIONS" | "TRACE";

interface fetchConfigSuccessType {
    data: object | string | ArrayBuffer,
    statusCode: number,
    header: object,
    cookies: Array<string>
}
interface fetchConfig {
    url?: string,
    data?: object | string | ArrayBuffer,
    header?: object,
    method?: fetchConfigMethod,
    timeout?: number,
    dataType?: fetchConfigDataType,
    responseType?: fetchConfigResponseType,
    sslVerify?: boolean,
    withCredentials?: boolean,
    firstIpv4?: boolean,
    success?: Function,
    fail?: Function,
    complete?: Function
}
type openUrlType = "navigate" | "redirect" | "reLaunch" | "switchTab" | "navigateBack"
declare interface Touch {
    readonly clientX: number;
    readonly clientY: number;
    readonly force: number;
    readonly identifier: number;
    readonly pageX: number;
    readonly pageY: number;
    readonly radiusX: number;
    readonly radiusY: number;
    readonly rotationAngle: number;
    readonly screenX: number;
    readonly screenY: number;
    readonly target: EventTarget;
    readonly x: number;
    readonly y: number;
}

declare interface TouchEvent {
    readonly changedTouches: TouchList
}




type tmUtil = {
    //pagejson下的pages配置。
    pages: Array<{ path: string, custom: 'custom' | 'default' }>,
    //pagejson下的配置。
    tabBar: tabBarType,
    globalNavStyle: "custom" | "default",
    /**
     * 判断是否是颜色值
     * @param color 颜色值
     */
    isColor(color: string): boolean,
    /**
     * 国际化
     * @param key 关键字
     * @return 语言值
     */
    language(key: string): string,
    fetch: {
        /**
         * GET请求
         * @param url 请求地址
         * @param data 请求的数据
         * @param opts 请求的配置
         * @help https://tmui.design/doc/JSTool/fetch.html
         */
        get(url: string, data?: object, opts?: fetchConfig): Promise<UniApp.GeneralCallbackResult | UniApp.RequestSuccessCallbackResult>,
        /**
         * POST请求
         * @param url 请求地址
         * @param data 请求的数据
         * @param opts 请求的配置
         * @help https://tmui.design/doc/JSTool/fetch.html
         */
        post(url: string, data?: object, opts?: fetchConfig): Promise<UniApp.GeneralCallbackResult | UniApp.RequestSuccessCallbackResult>,
        /**
         * 自定义请求
         * @param cog 请求的配置
         * @param beforeFun 请求前执行的函数
         * @param afterFun 请求后执行的函数
         * @param complete 请求完成的函数
         * @help https://tmui.design/doc/JSTool/fetch.html
         */
        request(cog: fetchConfig, beforeFun?: Function, afterFun?: Function, complete?: Function): Promise<UniApp.GeneralCallbackResult | UniApp.RequestSuccessCallbackResult>,
    },
    /**
     * tmui3.0函数工具
     */
    u: {

        /**
         * 检测是否是数字
         * @param arg 待检测的字符
         * @param defaultNum 0,如果不符合值时设置默认值
         * @returns number类型数值
         */
        isNumber(arg: string | number | undefined | null, defaultNum: number): number,

        /**
         * 检测是否是字符串
         * @param arg 待检测的字符
         * @param defaultNum 默认"",如果不符合值是设置默认值
         * @returns 字符串
         */
        isString(arg: string | number | undefined | null, defaultStr: string): string,

        /**
         * 把一个数字进行分页返回数字数组
         * @param total 总数
         * @param pageSize 分页大小
         * @returns 数字数组
         */
        paginate(total: number, pageSize: number): number[],

        /**
           * 取对象数据值（可深层次取值）
           * @example getValue(data,"a.b.c")
           * @param data 对象数据
           * @param keys 键值
           * @returns 返回值
           * @description 注意不会去改变原来的数据
           */
        getValue(data: Data, keys: string): any,
        /**
         * 设置对象键值（可深层次设置值）
         * @example setValue(data,"a.b.c","haha")
         * @param data 对象数据
         * @param keys 键值
         * @returns 修改后的对象数据。
         * @description 改变原来的数据
         */
        setValue(data: Data, keys: string, value: any): void,
        /**
         * 计算并返回一个对象中最大的层级数
         * @param data 待检测对象数据
         * @returns 最大层级数
         */
        getMaxDepth(data: Data): number,

        /**
        * 预览图片。
        * @param {Object} url 必填 当前预览的图片链接。
        * @param {Object} list 可以是url数组，也可以是对象，数据比如：["http:url"] or [{url:"https:url",...}]
        * @param {Object} rangKey 如果list是对象数组，需要提供url字段。
        */
        preview(url: string, list?: Array<string>, rangKey?: string): void,

        /**
        * 数据分组
        * @param {Array} oArr - 原数组列表
        * @param {Number} length - 单个数组长度
        * @return {Array}  arr - 分组后的新数组
        */
        splitData<T>(arr: Array<T>, size: number): Array<T[]>,

        /**
        * 剩余时间格式化
        * @param {Number} t - 剩余多少秒
        * @return {Object}  format - 格式后的天时分秒对象
        */
        timeMuch(t: number): string,
        /**
         * 获取时间距离当前时间
         * @param timestamp 当前时间
         * @return 返回比如：不久前，1年前这样的格式文本
         */
        getDateToNewData(timestamp?: number | string | Date): string,
        /**
        * 打电话
        * @param {String<Number>} phoneNumber - 数字字符串
        * @return Promise<boolean>
        */
        callPhone(phoneNumber: string): Promise<boolean | any>,
        /**
         * 调起客户端相机扫码。
         * @param {Boolean} onlyFromCamera true 是否只允许相机扫码识别
         * @param {Array<string>} scanType ['barCode', 'qrCode', 'datamatrix','datamatrix']
         * @returns Promise 成功返回相关数据结构
         */
        scanCode(onlyFromCamera: boolean, scanType: Array<any>): Promise<UniApp.ScanCodeSuccessRes | string>,
        /**
         * 设置剪切板内容。
         * @param {String} data 
         * @returns Promise true/false
         */
        setClipboardData(data: string): Promise<boolean | string>,
        /**
         * 获取剪切板内容
         * @returns Promise 剪切板内容
         */
        getClipboardData(): Promise<string | boolean>,
        /**
         * 设置cookie数据
         * @param {String} key 键值
         * @param {String} data 值
         * @returns Boolean
         */
        setCookie(key: string, data: any): boolean,
        /**
         * 删除一个本地cookie
         * @param {String} key 键值
         * @returns Boolean
         */
        delCookie(key: string): boolean,
        /**
         * 获取一个cookie数据
         * 如果存入的是对象，返回的也是对象。如果是string返回的也是字符串。
         * @param {String} key 键
         * @returns json/string
         */
        getCookie(key: string): Object | string,
        /**
         * 向地址连接追加参数。
         * @param {string} uri 网址
         * @param {string} key 字段
         * @param {string} value 字段值
         * @returns 网址
         */
        httpUrlAddKey(uri: string, key: string, value: string): string,
        /**
         * 取url参数
         * @param {string} uri 网址
         * @param {string} key 字段
         * @returns string|undefined
         */
        getQueryString(url: string, key: string): string,
        /**
         * 唯一标识
         * @param rdix 1 随机因子
         * @param length  12 取的长度
         * @param isAddStr false 是否限制随机结果中的长度,不允许输出长度
         * @returns String
         */
        getUid(rdix?: number, length?: number, isAddStr?: boolean): number | string,
        /**
         * 防抖
         * 防抖原理：在一定时间内，只有最后一次操作，再过wait毫秒后才执行函数
         * @param {Function} func 要执行的回调函数
         * @param {Number} wait 延迟的时间
         * @param {Boolean} immediate 是否要立即执行
        */
        debounce(func: Function, wait?: number, immediate?: boolean): void,
        /**
         * 节流
         * 节流原理：在一定时间内，只能触发一次
         * @param {Function} func 要执行的回调函数 
         * @param {Number} wait 延时的时间
         * @param {Boolean} immediate 是否立即执行
         * @return void
         */
        throttle(func: Function, wait?: number, immediate?: boolean, timer?: number, flags?: boolean): void,
        /**
         * 深度克隆
         * @param {T} data 待大克隆复制的数据
         * @return {T} any
         */
        deepClone<T>(data: T): T,
        /**
         * 等同：queryDom
         */
        quereyDom(t: ComponentInternalInstance, node: string): Promise<UniApp.NodeInfo | UniApp.NodeInfo[]>,
        /**
         * 查询文档节点信息
         * @param t Vue上下文对象
         * @param node 提供带#的id比如：'#id',在nvue中应该是元素上写明ref='id'
         * @returns vue页面返回查询的节点信息，nvue返回weex的节点信息。
         */
        queryDom(t: ComponentInternalInstance, node: string): Promise<UniApp.NodeInfo | UniApp.NodeInfo[]>,
        /**
         * 深度合并对象
         * @param FirstOBJ 需要合并的对象
         * @param SecondOBJ 被合并的对象
         * @returns 返回合并后的对象 
         */
        deepObjectMerge(FirstOBJ: Record<string, any>, SecondOBJ: Record<string, any>): Record<string, any>,
        /**
         * 是否是手机号码
         * @param phone 号码
         * @returns Boolean
         */
        isPhone(phone: string | number): boolean,
        /**
         * 是否含有中文
         * @param s 字符串
         * @returns Boolean
         */
        isChina(s: string): boolean,
        /**
         * 是否为空
         * @description 判断是否是null,对象是否为空，数组是否为空。是否为 undefaind，是否为 “”空字符串。
         * @param s 任意
         */
        isEmpty(s: any): boolean,
        /**
         * 是否邮箱
         * @param s 字符串
         * @returns Boolean
         */
        isEmail(s: string): boolean,
        /**
         * 是否身份证号
         * @param val 字符号或者数字
         * @returns Boolean
         * @author https://cloud.tencent.com/developer/article/1114323
         */
        isIdCard(val: string | number): boolean,
        /**
         * 是否车牌
         * @description 蓝牌5位，绿牌6位。
         * @param s 字符串
         * @returns Boolean
         */
        isIdCar(s: string): boolean,
        /**
         * 纯数字密码验证
         * @param s 字符串或者数字
         * @param len 最小长度，默认6
         * @param maxLen 最大长度，默认20
         * @returns Boolean
         */
        isPasswordOfNumber(s: number | string, le: number, maxLen: number): boolean,
        /**
         * 密码验证
         * @param s 字符串或者数字
         * @param len 最小长度，默认6
         * @param maxLen 最大长度，默认20
         * @param model 0数字和英文，1数字，英文必须包含，不允许有特殊字符，2数字和字母必须包含，可以有特殊字符。
         * @returns Boolean
         */
        isPasswordOfOther(s: string | number, len: number, maxLen: number, model: number): boolean,
        /**
         * 是否是一个有效的日期
         * @param s 字符串，数字，日期对象
         * @returns Boolean
         */
        isDate(s: string | number | Date): boolean,
        /**
         * 显示信息
         * @param word 标题
         * @param mask 不允许穿透
         * @param icon 图标
         */
        toast(word: string, mask?: boolean, icon?: any): void,
        /**
         * 获取屏幕窗口安全高度和宽度
         * 注意是针对种屏幕的统一计算，统一高度，不再让uni获取有效高度而烦恼。
         * 请一定要在onMounted或者onLoad中调用，否则不准确在h5端。
         * @return {height,width,top,isCustomHeader,sysinfo}
         */
        getWindow(): { width: number, height: number, top: number, bottom: number, statusBarHeight: number, isCustomHeader: Boolean, sysinfo: UniApp.GetSystemInfoResult },

        /**
         * 打开页面路径
         * @param url string 打开的页面路径
         * @param type openUrlType "navigate" | "redirect" | "reLaunch" | "switchTab" | "navigateBack"
         */
        routerTo(url: string, type?: openUrlType): void,
        /**
         * 将rpx转换为px
         * @param v 待转换的数字
         * @param screenWidth 屏幕的宽度，如果不提供默认自动获取
         * @return number
         */
        torpx(v: number, screenWidth?: number): number
        /**
         * 将rpx转换为px
         * @param v 待转换的数字
         * @return number
         */
        topx(v: number, screenWidth?: number): number,
        /**
         * 在下一次前执行回调函数
         * @param callback 回调函数
         * @returns 一个id值，取消时cancelAnimationFrame(id)来取消
         */
        requestAnimationFrame(callback: Function): number,
        /**
         * 取消回调执行
         * @param id requestAnimationFrame产生的id
         */
        cancelAnimationFrame(id: number): void


    },
    tmicon: Array<{
        font: string,
        prefix: string,
        fontJson: Array<{
            icon_id: string,
            name: string,
            font_class: string,
            unicode: string,
            unicode_decimal: number
        }>
    }>,
    /**tmui3.0的全局配置，以代替router,theme等的外围文件夹。解耦相关目录 */
    config: Tmui.tmuiConfig
}


