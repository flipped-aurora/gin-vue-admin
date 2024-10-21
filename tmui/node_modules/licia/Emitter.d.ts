import once = require('./once');
import types = require('./types');

declare class Emitter {
    on(event: string, listener: types.AnyFn): Emitter;
    off(event: string, listener: types.AnyFn): Emitter;
    once(event: string, listener: types.AnyFn): Emitter;
    emit(event: string, ...args: any[]): Emitter;
    removeAllListeners(event?: string): Emitter;
    static mixin(obj: any): any;
}

export = Emitter;
