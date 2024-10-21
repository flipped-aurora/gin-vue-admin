import type { UTSTargetLanguage } from './uts';
export type DefineOptions = {
    name?: string;
    app?: boolean | {
        js?: boolean;
        kotlin?: boolean;
        swift?: boolean;
    };
    [key: string]: any;
};
export type Define = string | string[] | Record<string, string | DefineOptions> | false;
export type Defines = {
    [name: string]: Define;
};
export interface Exports {
    [name: string]: Define | Defines | false;
}
export declare function getUniExtApiProviders(): {
    plugin: string;
    service: string;
    name?: string | undefined;
    servicePlugin?: string | undefined;
}[];
export declare function getUniExtApiPlugins(): {
    plugin: string;
}[];
export declare function formatExtApiProviderName(service: string, name: string): string;
export declare function getUniExtApiProviderRegisters(): {
    name: string;
    plugin: string;
    service: string;
    class: string;
}[];
export declare function parseUniExtApis(vite: boolean | undefined, platform: typeof process.env.UNI_UTS_PLATFORM, language?: UTSTargetLanguage): Injects;
export declare function parseUniExtApi(pluginDir: string, pluginId: string, vite: boolean | undefined, platform: typeof process.env.UNI_UTS_PLATFORM, language?: UTSTargetLanguage): Injects | undefined;
export type Injects = {
    [name: string]: string | [string, string] | [string, string, DefineOptions['app']] | false;
};
/**
 *  uni:'getBatteryInfo'
 * import getBatteryInfo from '..'
 *
 * uni:['getBatteryInfo']
 * import { getBatteryInfo } from '..'
 *
 * uni:['openLocation','chooseLocation']
 * import { openLocation, chooseLocation } from '..'
 *
 * uni:{
 *  onUserCaptureScreen: "onCaptureScreen"
 *  offUserCaptureScreen: "offCaptureScreen"
 * }
 *
 * uni.getBatteryInfo = getBatteryInfo
 * @param source
 * @param globalObject
 * @param define
 * @returns
 */
export declare function parseInjects(vite: boolean | undefined, platform: typeof process.env.UNI_UTS_PLATFORM, language: UTSTargetLanguage, source: string, uniModuleRootDir: string, exports?: Exports): Injects;
/**
 * @private
 */
export declare const camelize: (str: string) => string;
/**
 * @private
 */
export declare const capitalize: (str: string) => string;
/**
 * 解析 UTS 类型的模块依赖列表
 * @param deps
 * @param inputDir
 * @returns
 */
export declare function parseUTSModuleDeps(deps: string[], inputDir: string): string[];
export declare function genEncryptEasyComModuleIndex(platform: typeof process.env.UNI_UTS_PLATFORM, components: Record<string, '.vue' | '.uvue'>): string;
export declare function parseUniModulesWithComponents(inputDir: string): Record<string, Record<string, ".vue" | ".uvue">>;
/**
 * 解析 easyCom 组件列表
 * @param pluginId
 * @param inputDir
 * @returns
 */
export declare function parseEasyComComponents(pluginId: string, inputDir: string, detectBinary?: boolean): Record<string, ".vue" | ".uvue">;
export declare function findEncryptUniModules(inputDir: string, cacheDir?: string): Record<string, EncryptPackageJson | undefined>;
export declare function findUploadEncryptUniModulesFiles(uniModules: Record<string, EncryptPackageJson | undefined>, platform: typeof process.env.UNI_UTS_PLATFORM, inputDir: string): Record<string, string[]>;
export declare function packUploadEncryptUniModules(uniModules: Record<string, EncryptPackageJson | undefined>, platform: typeof process.env.UNI_UTS_PLATFORM, inputDir: string, cacheDir: string): {
    zipFile: string;
    modules: string[];
};
interface EncryptPackageJson {
    id: string;
    version: string;
    uni_modules: {
        dependencies: string[];
        artifacts: {
            env: {
                compilerVersion: string;
            } & Record<string, any>;
            apis: string[];
            components: string[];
            scopedSlots: string[];
            declaration: string;
        };
    };
}
export declare function initCheckEnv(): Record<string, string>;
export declare function resolveEncryptUniModule(id: string, platform: typeof process.env.UNI_UTS_PLATFORM, isX?: boolean): string | undefined;
export declare function checkEncryptUniModules(inputDir: string, params: {
    mode: 'development' | 'production';
    packType: 'debug' | 'release';
    compilerVersion: string;
    appid: string;
    appname: string;
    platform: typeof process.env.UNI_UTS_PLATFORM;
    'uni-app-x': boolean;
}): Promise<{} | undefined>;
export declare function parseUniModulesArtifacts(): {
    name: string;
    package: string;
    scopedSlots: string[];
    declaration: string;
}[];
export {};
