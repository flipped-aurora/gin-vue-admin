declare namespace intersectRange {
    interface IRange {
        start: number;
        end: number;
    }
}
declare function intersectRange(
    a: intersectRange.IRange,
    b: intersectRange.IRange
): intersectRange.IRange | void;

export = intersectRange;
