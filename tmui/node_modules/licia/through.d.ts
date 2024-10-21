import extend = require('./extend');

import stream = require('stream');
declare namespace through {
    interface ThroughConstructor extends stream.Transform {
        new (opts?: stream.DuplexOptions): stream.Transform;
        (opts?: stream.DuplexOptions): stream.Transform;
    }
    type TransformCallback = (err?: any, data?: any) => void;
    type TransformFunction = (
        this: stream.Transform,
        chunk: any,
        enc: string,
        callback: TransformCallback
    ) => void;
    type FlushCallback = (
        this: stream.Transform,
        flushCallback: () => void
    ) => void;
    function obj(
        transform?: TransformFunction,
        flush?: FlushCallback
    ): stream.Transform;
    function ctor(
        transform?: TransformFunction,
        flush?: FlushCallback
    ): ThroughConstructor;
    function ctor(
        opts?: stream.DuplexOptions,
        transform?: TransformFunction,
        flush?: FlushCallback
    ): ThroughConstructor;
}
declare function through(
    transform?: through.TransformFunction,
    flush?: through.FlushCallback
): stream.Transform;
declare function through(
    opts?: stream.DuplexOptions,
    transform?: through.TransformFunction,
    flush?: through.FlushCallback
): stream.Transform;

export = through;
