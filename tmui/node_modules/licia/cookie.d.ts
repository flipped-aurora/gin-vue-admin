declare namespace cookie {
    interface IOptions {
        path?: string;
        expires?: number;
        domain?: string;
        secure?: boolean;
    }
    interface ICookie {
        get(key: string, options?: cookie.IOptions): string;
        set(key: string, val: string, options?: cookie.IOptions): ICookie;
        remove(key: string, options?: cookie.IOptions): ICookie;
    }
}
declare const cookie: cookie.ICookie;

export = cookie;
