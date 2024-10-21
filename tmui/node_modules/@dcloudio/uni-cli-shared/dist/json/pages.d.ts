export declare function isUniPageFile(file: string, inputDir?: string): boolean;
export declare function isUniPageSetupAndTs(file: string): boolean;
export declare function isUniPageSfcFile(file: string, inputDir?: string): boolean;
/**
 * 小程序平台慎用，因为该解析不支持 subpackages
 * @param inputDir
 * @param platform
 * @param normalize
 * @returns
 */
export declare const parsePagesJson: (inputDir: string, platform: UniApp.PLATFORM, normalize?: boolean) => UniApp.PagesJson;
/**
 * 该方法解析出来的是不支持 subpackages，会被合并入 pages
 */
export declare const parsePagesJsonOnce: (inputDir: string, platform: UniApp.PLATFORM, normalize?: boolean) => UniApp.PagesJson;
/**
 * 目前 App 和 H5 使用了该方法
 * @param jsonStr
 * @param platform
 * @param param2
 * @returns
 */
export declare function normalizePagesJson(jsonStr: string, platform: UniApp.PLATFORM, { subpackages, }?: {
    subpackages: boolean;
}): UniApp.PagesJson;
export declare function validatePages(pagesJson: Record<string, any>, jsonStr: string): void;
export declare function removePlatformStyle(pageStyle: Record<string, any>): Record<string, any>;
export declare function normalizePagesRoute(pagesJson: UniApp.PagesJson): UniApp.UniRoute[];
declare function parseSubpackagesRoot(inputDir: string, platform: UniApp.PLATFORM): string[];
export declare const parseSubpackagesRootOnce: typeof parseSubpackagesRoot;
export declare function filterPlatformPages(platform: UniApp.PLATFORM, pagesJson: UniApp.PagesJson): void;
export {};
