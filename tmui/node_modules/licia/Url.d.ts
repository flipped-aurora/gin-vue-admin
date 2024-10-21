import query = require('./query');
import types = require('./types');
import toStr = require('./toStr');

declare namespace Url {
    interface IUrl {
        protocol: string;
        auth: string;
        hostname: string;
        hash: string;
        query: any;
        port: string;
        pathname: string;
        slashes: boolean;
    }
}
declare class Url {
    protocol: string;
    auth: string;
    hostname: string;
    hash: string;
    query: any;
    port: string;
    pathname: string;
    slashes: boolean;
    constructor(url?: string);
    setQuery(name: string, val: string | number): Url;
    setQuery(query: types.PlainObj<string | number>): Url;
    rmQuery(name: string | string[]): Url;
    toString(): string;
    static parse(url: string): Url.IUrl;
    static stringify(object: Url.IUrl): string;
}

export = Url;
