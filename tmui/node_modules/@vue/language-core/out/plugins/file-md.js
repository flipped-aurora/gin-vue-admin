"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
const source_map_1 = require("@volar/source-map");
const parseSfc_1 = require("../utils/parseSfc");
const codeblockReg = /```[\s\S]+?```/g;
const inlineCodeblockReg = /`[^\n`]+?`/g;
const scriptSetupReg = /\\\<[\s\S]+?\>\n?/g;
const sfcBlockReg = /\<(script|style)\b[\s\S]*?\>([\s\S]*?)\<\/\1\>/g;
const angleBracketReg = /\<\S*\:\S*\>/g;
const linkReg = /\[[\s\S]*?\]\([\s\S]*?\)/g;
const codeSnippetImportReg = /^\s*<<<\s*.+/gm;
const plugin = () => {
    return {
        version: 1,
        parseSFC(fileName, content) {
            if (fileName.endsWith('.md')) {
                content = content
                    // code block
                    .replace(codeblockReg, match => '```' + ' '.repeat(match.length - 6) + '```')
                    // inline code block
                    .replace(inlineCodeblockReg, match => `\`${' '.repeat(match.length - 2)}\``)
                    // # \<script setup>
                    .replace(scriptSetupReg, match => ' '.repeat(match.length))
                    // <<< https://vitepress.dev/guide/markdown#import-code-snippets
                    .replace(codeSnippetImportReg, match => ' '.repeat(match.length));
                const codes = [];
                for (const match of content.matchAll(sfcBlockReg)) {
                    if (match.index !== undefined) {
                        const matchText = match[0];
                        codes.push([matchText, undefined, match.index]);
                        codes.push('\n\n');
                        content = content.substring(0, match.index) + ' '.repeat(matchText.length) + content.substring(match.index + matchText.length);
                    }
                }
                content = content
                    // angle bracket: <http://foo.com>
                    .replace(angleBracketReg, match => ' '.repeat(match.length))
                    // [foo](http://foo.com)
                    .replace(linkReg, match => ' '.repeat(match.length));
                codes.push('<template>\n');
                codes.push([content, undefined, 0]);
                codes.push('\n</template>');
                const file2VueSourceMap = new source_map_1.SourceMap((0, source_map_1.buildMappings)(codes));
                const sfc = (0, parseSfc_1.parse)((0, source_map_1.toString)(codes));
                if (sfc.descriptor.template) {
                    transformRange(sfc.descriptor.template);
                }
                if (sfc.descriptor.script) {
                    transformRange(sfc.descriptor.script);
                }
                if (sfc.descriptor.scriptSetup) {
                    transformRange(sfc.descriptor.scriptSetup);
                }
                for (const style of sfc.descriptor.styles) {
                    transformRange(style);
                }
                for (const customBlock of sfc.descriptor.customBlocks) {
                    transformRange(customBlock);
                }
                return sfc;
                function transformRange(block) {
                    block.loc.start.offset = file2VueSourceMap.toSourceOffset(block.loc.start.offset)?.[0] ?? -1;
                    block.loc.end.offset = file2VueSourceMap.toSourceOffset(block.loc.end.offset)?.[0] ?? -1;
                }
            }
            ;
        }
    };
};
exports.default = plugin;
//# sourceMappingURL=file-md.js.map