import Store = require('./Store');

declare class LocalStore extends Store {
    constructor(name: string, data?: {});
}

export = LocalStore;
