declare namespace isIp {
    function v4(str: string): boolean;
    function v6(str: string): boolean;
}
declare function isIp(str: string): boolean;

export = isIp;
