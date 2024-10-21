declare namespace $safeEls {
    type El = Element | Element[] | NodeListOf<Element> | string;
}
declare function $safeEls(val: $safeEls.El): Element[];

export = $safeEls;
