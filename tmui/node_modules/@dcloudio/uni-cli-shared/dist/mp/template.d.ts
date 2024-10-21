import type { EmittedAsset } from 'rollup';
import type { ElementNode } from '@vue/compiler-core';
import type { MiniProgramComponentsType } from '../json/mp/types';
type LazyElementFn = (node: ElementNode, context: {
    isMiniProgramComponent(name: string): MiniProgramComponentsType | undefined;
}) => {
    [name: string]: {
        name: 'on' | 'bind';
        arg: string[];
    }[] | true;
} | boolean;
export interface MiniProgramCompilerOptions {
    /**
     * 需要延迟渲染的组件，通常是某个组件的某个事件会立刻触发，需要延迟到首次 render 之后，比如微信 editor 的 ready 事件，快手 switch 的 change
     */
    lazyElement?: {
        [name: string]: {
            name: 'on' | 'bind';
            arg: string[];
        }[] | true;
    } | LazyElementFn;
    event?: {
        key?: boolean;
        format?(name: string, opts: {
            isCatch?: boolean;
            isCapture?: boolean;
            isComponent?: boolean;
        }): string;
    };
    class: {
        /**
         * 是否支持绑定 array 类型
         */
        array: boolean;
    };
    slot: {
        /**
         * 是否支持 $slots.default 访问
         */
        $slots?: boolean;
        /**
         * 是否支持后备内容
         */
        fallbackContent?: boolean;
        /**
         * 是否支持动态插槽名
         */
        dynamicSlotNames?: boolean;
    };
    filter?: {
        lang: string;
    };
    component?: {
        /**
         * 平台自定义组件目录，如 wxcomponents
         */
        dir?: string;
        /**
         * 自定义组件自定义 hidden 属性用于实现 v-show
         */
        vShow?: string;
        /**
         * 父组件 setData 后，子组件的 properties 是否可以同步获取，目前仅 mp-weixin，mp-qq，mp-alipay 支持
         */
        getPropertySync?: boolean;
        /**
         * 格式化组件名称，比如 wx-btn => weixin-btn (微信不允许以 wx 命名自定义组件)
         */
        normalizeName?: (name: string) => string;
        /**
         * 合并虚拟化节点属性（class、style）
         */
        mergeVirtualHostAttributes?: boolean;
    };
    directive: string;
    emitFile?: (emittedFile: EmittedAsset) => string;
}
export interface MiniProgramFilterOptions {
    id: string;
    type: string;
    name: string;
    src?: string;
    code: string;
}
type GenFilterFn = (filter: MiniProgramFilterOptions, filename: string) => string | void;
export declare function findMiniProgramTemplateFiles(genFilter?: GenFilterFn): Record<string, string>;
export declare function clearMiniProgramTemplateFiles(): void;
export declare function addMiniProgramTemplateFile(filename: string, code: string): void;
export declare function clearMiniProgramTemplateFilter(filename: string): void;
export declare function addMiniProgramTemplateFilter(filename: string, filter: MiniProgramFilterOptions): void;
export {};
