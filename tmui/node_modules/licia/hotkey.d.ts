import types = require('./types');

declare namespace hotkey {
    interface IOptions {
        element?: HTMLElement;
    }
}
declare const hotkey: {
    on(key: string, options: hotkey.IOptions, listener: types.AnyFn): void;
    on(key: string, listener: types.AnyFn): void;
    off(key: string, options: hotkey.IOptions, listener: types.AnyFn): void;
    off(key: string, listener: types.AnyFn): void;
};

export = hotkey;
