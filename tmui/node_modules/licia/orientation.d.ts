import Emitter = require('./Emitter');

declare namespace orientation {
    interface IOrientation extends Emitter {
        get(): string;
    }
}
declare const orientation: orientation.IOrientation;

export = orientation;
