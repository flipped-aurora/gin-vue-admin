declare class Enum {
    size: number;
    constructor(map: string[] | { [member: string]: any });
    [key: string]: any;
}

export = Enum;
