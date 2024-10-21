import type { Plugin } from 'vite';
import type { VitePluginUniResolvedOptions } from '../..';
declare const BUILT_IN_MODULES: {
    'vue-router': string;
    vuex: string;
    'vue-i18n': string;
    '@dcloudio/uni-app': string;
    '@dcloudio/uni-cloud': string;
    '@dcloudio/uni-i18n': string;
    '@dcloudio/uni-shared': string;
    '@dcloudio/uni-stacktracey': string;
    '@vue/shared': string;
    pinia: string;
};
export type BuiltInModulesKey = keyof typeof BUILT_IN_MODULES;
export declare function uniResolveIdPlugin(options: VitePluginUniResolvedOptions): Plugin;
export {};
