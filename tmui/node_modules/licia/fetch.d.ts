import has = require('./has');
import types = require('./types');

declare namespace fetch {
    interface IResult {
        ok: boolean;
        status: number;
        statusText: string;
        url: string;
        clone(): IResult;
        text(): Promise<string>;
        json(): Promise<any>;
        xml(): Promise<Document | null>;
        blob(): Promise<Blob>;
        headers: {
            keys(): string[];
            entries(): Array<string[]>;
            get(name: string): string;
            has(name: string): boolean;
        };
    }
}
declare function fetch(
    url: string,
    options?: {
        method?: string;
        timeout?: number;
        headers?: types.PlainObj<string>;
        body?: any;
    }
): Promise<fetch.IResult>;

export = fetch;
