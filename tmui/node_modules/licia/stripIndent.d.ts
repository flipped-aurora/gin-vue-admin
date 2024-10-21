declare function stripIndent(str: string): string;
declare function stripIndent(
    literals: TemplateStringsArray,
    ...placeholders: any[]
): string;

export = stripIndent;
