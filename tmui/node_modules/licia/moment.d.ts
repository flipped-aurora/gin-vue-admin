import toDate = require('./toDate');
import isLeapYear = require('./isLeapYear');
import toStr = require('./toStr');

declare namespace moment {
    class M {
        constructor(value: string | Date);
        format(mask: string): string;
        isValid(): boolean;
        isLeapYear(): boolean;
        isSame(that: M): boolean;
        valueOf(): number;
        isBefore(that: M): boolean;
        isAfter(that: M): boolean;
        year(): number;
        year(number): M;
        month(): number;
        month(number): M;
        date(): number;
        date(number): M;
        hour(): number;
        hour(number): M;
        minute(): number;
        minute(number): M;
        second(): number;
        second(number): M;
        millisecond(): number;
        millisecond(number): M;
        unix(): number;
        clone(): M;
        toDate(): Date;
        toArray(): number[];
        toJSON(): string;
        toISOString(): string;
        toObject(): any;
        toString(): string;
        set(unit: string, num: number): M;
        startOf(unit: string): M;
        endOf(unit: string): M;
        daysInMonth(): number;
        add(num: number, unit: string): M;
        subtract(num: number, unit: string): M;
        diff(input: M | string | Date, unit: string, asFloat: boolean): number;
    }
}
declare function moment(value: string | Date): moment.M;

export = moment;
