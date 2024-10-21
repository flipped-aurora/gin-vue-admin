declare function lazyImport<T>(
    importFn: (moduleId: string) => T,
    dirname?: string
): (moduleId: string) => T;

export = lazyImport;
