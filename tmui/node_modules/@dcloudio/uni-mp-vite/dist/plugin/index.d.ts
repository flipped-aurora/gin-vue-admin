import type { AliasOptions } from 'vite';
import { type AppJson, type CopyOptions, type MiniProgramCompilerOptions, type UniVitePlugin, type findMiniProgramTemplateFiles } from '@dcloudio/uni-cli-shared';
import type { CompilerOptions } from '@dcloudio/uni-mp-compiler';
export interface UniMiniProgramPluginOptions {
    cdn?: number;
    vite: {
        alias: AliasOptions;
        copyOptions: CopyOptions;
        inject: {
            [name: string]: [string, string];
        };
    };
    global: string;
    json?: {
        windowOptionsMap?: Record<string, string>;
        tabBarOptionsMap?: Record<string, string>;
        tabBarItemOptionsMap?: Record<string, string>;
        formatAppJson?: (appJson: Record<string, any>, manifestJson: Record<string, any>, pagesJson: Record<string, any>) => void;
    };
    app: {
        /**
         * 是否支持darkmode
         */
        darkmode?: boolean;
        /**
         * 是否支持subpackages
         */
        subpackages?: boolean;
        /**
         * 是否支持发行插件
         */
        plugins?: boolean;
        /**
         * 是否支持全局组件
         */
        usingComponents: boolean;
        normalize?: (appJson: AppJson) => AppJson;
    };
    project?: {
        filename: string;
        config: string[];
        source: Record<string, any>;
        normalize?: (projectJson: Record<string, unknown>) => Record<string, unknown>;
    };
    template: {
        extname: string;
        directive: string;
        event?: MiniProgramCompilerOptions['event'];
        class: MiniProgramCompilerOptions['class'];
        slot: MiniProgramCompilerOptions['slot'];
        lazyElement?: MiniProgramCompilerOptions['lazyElement'];
        component?: MiniProgramCompilerOptions['component'];
        customElements?: string[];
        filter?: {
            lang: string;
            extname: string;
            generate: Parameters<typeof findMiniProgramTemplateFiles>[0];
        };
        compilerOptions?: CompilerOptions;
    };
    style: {
        extname: string;
    };
}
export declare function uniMiniProgramPlugin(options: UniMiniProgramPluginOptions): UniVitePlugin;
