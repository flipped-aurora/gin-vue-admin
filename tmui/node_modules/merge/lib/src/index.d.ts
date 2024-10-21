export default main;
export declare function main(clone: boolean, ...items: any[]): any;
export declare namespace main {
    var clone: typeof import(".").clone;
    var isPlainObject: typeof import(".").isPlainObject;
    var recursive: typeof import(".").recursive;
}
export declare function main(...items: any[]): any;
export declare namespace main {
    var clone: typeof import(".").clone;
    var isPlainObject: typeof import(".").isPlainObject;
    var recursive: typeof import(".").recursive;
}
export declare function merge(clone: boolean, ...items: any[]): any;
export declare function merge(...items: any[]): any;
export declare function recursive(clone: boolean, ...items: any[]): any;
export declare function recursive(...items: any[]): any;
export declare function clone<T>(input: T): T;
export declare function isPlainObject(input: any): input is Object;
