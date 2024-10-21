import types = require('./types');

declare namespace ajax {
    function get(
        url: string,
        data: string | {},
        success: types.AnyFn,
        dataType?: string
    ): XMLHttpRequest;
    function get(
        url: string,
        success: types.AnyFn,
        dataType?: string
    ): XMLHttpRequest;
    function post(
        url: string,
        data: string | {},
        success: types.AnyFn,
        dataType?: string
    ): XMLHttpRequest;
    function post(
        url: string,
        success: types.AnyFn,
        dataType?: string
    ): XMLHttpRequest;
}
declare function ajax(options: {
    type?: string;
    url: string;
    data?: string | {};
    dataType?: string;
    contentType?: string;
    success?: types.AnyFn;
    error?: types.AnyFn;
    complete?: types.AnyFn;
    timeout?: number;
}): XMLHttpRequest;

export = ajax;
