declare function lazyRequire<T>(
    requireFn: (moduleId: string) => T
): (moduleId: string) => T;

export = lazyRequire;
