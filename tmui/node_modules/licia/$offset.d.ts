import $safeEls = require('./$safeEls');

declare namespace $offset {
    interface IOffset {
        left: number;
        top: number;
        width: number;
        height: number;
    }
}
declare function $offset(element: $safeEls.El): $offset.IOffset;

export = $offset;
