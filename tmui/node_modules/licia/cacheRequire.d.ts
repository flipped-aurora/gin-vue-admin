declare function cacheRequire(options?: {
    dir?: string;
    requirePath?: boolean;
    code?: boolean;
    compileCache?: boolean;
}): void;

export = cacheRequire;
