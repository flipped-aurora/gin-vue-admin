export declare function initI18nOptions(platform: UniApp.PLATFORM, inputDir: string, warning?: boolean, withMessages?: boolean): {
    locale: string;
    locales: Record<string, Record<string, string>>;
    delimiters: [string, string];
} | undefined;
export declare const initI18nOptionsOnce: typeof initI18nOptions;
export declare function isUniAppLocaleFile(filepath: string): boolean;
export declare function getLocaleFiles(cwd: string): string[];
export declare function initLocales(dir: string, withMessages?: boolean): Record<string, Record<string, string>>;
export declare function resolveI18nLocale(platform: UniApp.PLATFORM, locales: string[], locale?: string): string;
