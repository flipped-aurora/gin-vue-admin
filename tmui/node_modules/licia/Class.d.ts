import extend = require('./extend');
import inherits = require('./inherits');
import types = require('./types');

declare namespace Class {
    class Base {
        toString(): string;
    }
    class IConstructor extends Base {
        constructor(...args: any[]);
        static extend(methods: any, statics: any): IConstructor;
        static inherits(Class: types.AnyFn): void;
        static methods(methods: any): IConstructor;
        static statics(statics: any): IConstructor;
        [method: string]: any;
    }
}
declare function Class(methods: any, statics?: any): Class.IConstructor;

export = Class;
