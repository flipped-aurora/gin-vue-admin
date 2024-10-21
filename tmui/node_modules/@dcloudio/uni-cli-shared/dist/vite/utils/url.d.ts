export interface VueQuery {
    vue?: boolean;
    src?: boolean;
    type?: 'script' | 'template' | 'style' | 'custom' | 'page';
    index?: number;
    lang?: string;
    raw?: boolean;
    setup?: boolean;
    'lang.ts'?: string;
    'lang.js'?: string;
}
export declare function parseVueRequest(id: string): {
    filename: string;
    query: VueQuery;
};
export declare const isImportRequest: (url: string) => boolean;
/**
 * Prefix for resolved fs paths, since windows paths may not be valid as URLs.
 */
export declare const FS_PREFIX = "/@fs/";
/**
 * Prefix for resolved Ids that are not valid browser import specifiers
 */
export declare const VALID_ID_PREFIX = "/@id/";
export declare const CLIENT_PUBLIC_PATH = "/@vite/client";
export declare const ENV_PUBLIC_PATH = "/@vite/env";
export declare const isInternalRequest: (url: string) => boolean;
export declare const queryRE: RegExp;
export declare const hashRE: RegExp;
export declare const cleanUrl: (url: string) => string;
export declare function isJsFile(id: string): boolean;
