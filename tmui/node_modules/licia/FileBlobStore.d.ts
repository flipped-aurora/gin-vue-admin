import types = require('./types');
import each = require('./each');
import Emitter = require('./Emitter');
import keys = require('./keys');

declare class FileBlobStore extends Emitter {
    constructor(path: string, data?: types.PlainObj<Buffer>);
    set(key: string, buf: Buffer): void;
    set(values: types.PlainObj<Buffer>): void;
    get(key: string): Buffer | void;
    get(keys: string[]): types.PlainObj<Buffer>;
    remove(key: string): void;
    remove(keys: string[]): void;
    clear(): void;
    each(fn: (val: Buffer, key: string) => void): void;
    save(): void;
}

export = FileBlobStore;
