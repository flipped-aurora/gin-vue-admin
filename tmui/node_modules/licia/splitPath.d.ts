declare function splitPath(
    path: string
): {
    dir: string;
    name: string;
    ext: string;
};

export = splitPath;
