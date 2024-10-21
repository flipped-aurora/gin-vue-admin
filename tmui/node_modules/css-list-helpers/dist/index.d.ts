/**
 * Splits a CSS declaration value (shorthand) using provided separators
 * as the delimiters.
 */
export declare function split(
    /**
     * A CSS declaration value (shorthand).
     */
    value: string, 
    /**
     * Any number of separator characters used for splitting.
     */
    separators: string[], {last}?: {
    last?: boolean;
}): string[];
/**
 * Splits a CSS declaration value (shorthand) using whitespace characters
 * as the delimiters.
 */
export declare function splitBySpaces(
    /**
     * A CSS declaration value (shorthand).
     */
    value: string): string[];
/**
 * Splits a CSS declaration value (shorthand) using commas as the delimiters.
 */
export declare function splitByCommas(
    /**
     * A CSS declaration value (shorthand).
     */
    value: string): string[];
