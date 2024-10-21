import types = require('./types');

declare class Validator {
    constructor(options: types.PlainObj<any>);
    validate(object: any): string | boolean;
    static plugins: any;
    static addPlugin(name: string, plugin: types.AnyFn): void;
}

export = Validator;
