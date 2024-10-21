import type { Plugin } from 'vite';
import type { TemplateCompiler } from '@vue/compiler-sfc';
import type { Options as VueOptions } from '@vitejs/plugin-vue';
import { type UniVitePlugin } from '@dcloudio/uni-cli-shared';
import type { ViteLegacyOptions, VitePluginUniResolvedOptions } from '..';
/**
 * 每次创建新的 plugin-vue 实例。因为该插件内部会 cache  descriptor，而相同的vue文件在编译到vue页面和nvue页面时，不能共享缓存（条件编译，css scoped等均不同）
 * @returns
 */
export declare function createPluginVueInstance(options: VueOptions): Plugin<any>;
export declare function initPluginVueOptions(options: VitePluginUniResolvedOptions, UniVitePlugins: UniVitePlugin[], uniPluginOptions: Required<Omit<Required<UniVitePlugin>['uni'], 'compiler'>> & {
    compiler?: TemplateCompiler;
}): VueOptions;
export declare function initPluginVueJsxOptions(options: VitePluginUniResolvedOptions, { isCustomElement, }: Required<Required<UniVitePlugin>['uni']>['compilerOptions'], jsxOptions: Required<Required<UniVitePlugin>['uni']>['jsxOptions']): import("@vue/babel-plugin-jsx").VueJSXPluginOptions & {
    babelPlugins?: any[] | undefined;
};
export declare function initPluginViteLegacyOptions(options: VitePluginUniResolvedOptions): ViteLegacyOptions;
