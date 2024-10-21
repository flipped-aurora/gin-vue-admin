import Store = require('./Store');

declare class FileStore extends Store {
    constructor(path: string, data?: any);
}

export = FileStore;
