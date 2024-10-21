/// <reference types="node" />
/// <reference types="node" />
import type * as UTSCompiler from '@dcloudio/uni-uts-v1';
import type { EasycomMatcher } from './easycom';
/**
 * 解析 app 平台的 uts 插件，任意平台（android|ios）存在即可
 * @param id
 * @param importer
 * @returns
 */
export declare function resolveUTSAppModule(platform: typeof process.env.UNI_UTS_PLATFORM, id: string, importer: string, includeUTSSDK?: boolean): string | undefined;
export declare function resolveUTSModule(id: string, importer: string, includeUTSSDK?: boolean): string | undefined;
export declare function resolveUTSCompiler(): typeof UTSCompiler;
export declare function isUTSComponent(name: string): boolean;
export declare function getUTSComponentAutoImports(): Record<string, [[string]]>;
export declare function parseUTSComponent(name: string, type: 'kotlin' | 'swift'): {
    className: string;
    namespace: string;
    source: string;
} | undefined;
export declare function initUTSComponents(inputDir: string, platform: UniApp.PLATFORM): EasycomMatcher[];
export declare function parseKotlinPackageWithPluginId(id: string, is_uni_modules: boolean): string;
export declare function parseSwiftPackageWithPluginId(id: string, is_uni_modules: boolean): string;
export type UTSTargetLanguage = typeof process.env.UNI_UTS_TARGET_LANGUAGE;
export declare const parseUniExtApiNamespacesOnce: (platform: typeof process.env.UNI_UTS_PLATFORM, language: UTSTargetLanguage) => Record<string, [string, string]>;
export declare const parseUniExtApiNamespacesJsOnce: (platform: typeof process.env.UNI_UTS_PLATFORM, language: UTSTargetLanguage) => Record<string, [string, string]>;
export declare function resolveUniTypeScript(): any;
