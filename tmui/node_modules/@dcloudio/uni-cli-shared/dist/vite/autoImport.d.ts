import type { Options } from 'unplugin-auto-import/types';
export type AutoImportOptions = Options;
export declare function initAutoImportOptions(platform: typeof process.env.UNI_UTS_PLATFORM, { imports, ...userOptions }: AutoImportOptions): AutoImportOptions;
