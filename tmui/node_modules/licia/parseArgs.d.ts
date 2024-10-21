declare function parseArgs(
    args: string[],
    options: {
        names: any;
        shorthands: any;
    }
): any;

export = parseArgs;
