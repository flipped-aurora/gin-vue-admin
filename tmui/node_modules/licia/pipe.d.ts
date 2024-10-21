import stream = require('stream');
declare function pipe(...streams: stream.Stream[]): void;

export = pipe;
