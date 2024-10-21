import type { Plugin } from 'vite';
import { type UniViteCopyPluginTarget, type UniVitePlugin } from '@dcloudio/uni-cli-shared';
import type { TemplateCompiler } from '@vue/compiler-sfc';
import type { VitePluginUniResolvedOptions } from '..';
export declare function initPluginUniOptions(UniVitePlugins: UniVitePlugin[]): {
    compiler?: TemplateCompiler;
    copyOptions: {
        assets: string[];
        targets: UniViteCopyPluginTarget[];
    };
    transformEvent: Record<string, string>;
    compilerOptions: Required<Required<UniVitePlugin>['uni']>['compilerOptions'];
    jsxOptions: Required<Required<UniVitePlugin>['uni']>['jsxOptions'];
    styleOptions: Required<Required<UniVitePlugin>['uni']>['styleOptions'];
};
export declare function initExtraPlugins(cliRoot: string, platform: UniApp.PLATFORM, options: VitePluginUniResolvedOptions): Plugin<any>[];
