import Select = require('./Select');
import $offset = require('./$offset');
import last = require('./last');
import types = require('./types');

declare namespace $ {
    class $ extends Select {
        find(selector: string): $;
        each(fn: types.AnyFn): $;
        offset(): $offset.IOffset;
        hide(): $;
        show(): $;
        first(): $;
        last(): $;
        get(index: number): Element;
        eq(index: number): $;
        on(event: string, selector: string, handler: types.AnyFn): $;
        on(event: string, handler: types.AnyFn): $;
        off(event: string, selector: string, handler: types.AnyFn): $;
        off(event: string, handler: types.AnyFn): $;
        html(): string;
        html(value: string): $;
        text(): string;
        text(value: string): $;
        val(): string;
        val(value: string): $;
        css(name: string): string;
        css(name: string, value: string): $;
        css(properties: types.PlainObj<string | number>): $;
        attr(name: string): string;
        attr(name: string, value: string): $;
        attr(attributes: types.PlainObj<string>): $;
        data(name: string): string;
        data(name: string, value: string): $;
        data(attributes: types.PlainObj<string>): $;
        rmAttr(name: string): $;
        remove(): $;
        addClass(name: string | string[]): $;
        rmClass(name: string): $;
        toggleClass(name: string): $;
        hasClass(name: string): boolean;
        parent(): $;
        append(content: string | Element): $;
        prepend(content: string | Element): $;
        before(content: string | Element): $;
        after(content: string | Element): $;
    }
}
declare function $(selector: string | Element | Document): $.$;

export = $;
