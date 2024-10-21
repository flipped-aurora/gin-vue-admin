import freeze = require('./freeze');

declare function freezeDeep<T>(obj: T): T;

export = freezeDeep;
