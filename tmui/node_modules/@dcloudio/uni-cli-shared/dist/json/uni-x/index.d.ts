export { parseUniXFlexDirection, parseUniXSplashScreen } from './manifest';
export declare function checkPagesJson(jsonStr: string, inputDir: string): boolean;
export declare function normalizeUniAppXAppPagesJson(jsonStr: string): UniApp.PagesJson;
/**
 * TODO 应该闭包，通过globalThis赋值？
 * @param pagesJson
 * @param manifestJson
 * @returns
 */
export declare function normalizeUniAppXAppConfig(pagesJson: UniApp.PagesJson, manifestJson: Record<string, any>): string;
