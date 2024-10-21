import Store = require('./Store');

declare class SessionStore extends Store {
    constructor(name: string, data?: any);
}

export = SessionStore;
