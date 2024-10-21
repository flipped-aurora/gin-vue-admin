declare class Delegator {
    constructor(host: object, target: object | string);
    method(name: string, target?: string): Delegator;
    getter(name: string, target?: string): Delegator;
    setter(name: string, target?: string): Delegator;
    access(name: string, target?: string): Delegator;
}

export = Delegator;
