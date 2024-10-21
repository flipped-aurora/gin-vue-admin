var Class = require('./Class');
var trim = require('./trim');
var repeat = require('./repeat');
var defaults = require('./defaults');
var camelCase = require('./camelCase');

exports = {
    parse: function(css) {
        return new Parser(css).parse();
    },
    stringify: function(stylesheet, options) {
        return new Compiler(stylesheet, options).compile();
    }
};
var regComments = /(\/\*[\s\S]*?\*\/)/gi;
var regOpen = /^{\s*/;
var regClose = /^}/;
var regWhitespace = /^\s*/;
var regProperty = /^(\*?[-#/*\\\w]+(\[[0-9a-z_-]+\])?)\s*/;
var regValue = /^((?:'(?:\\'|.)*?'|"(?:\\"|.)*?"|\([^)]*?\)|[^};])+)/;
var regSelector = /^([^{]+)/;
var regSemicolon = /^[;\s]*/;
var regColon = /^:\s*/;
var regMedia = /^@media *([^{]+)/;
var regKeyframes = /^@([-\w]+)?keyframes\s*/;
var regFontFace = /^@font-face\s*/;
var regSupports = /^@supports *([^{]+)/;
var regIdentifier = /^([-\w]+)\s*/;
var regKeyframeSelector = /^((\d+\.\d+|\.\d+|\d+)%?|[a-z]+)\s*/;
var regComma = /^,\s*/;
var Parser = Class({
    initialize: function Parser(css) {
        this.input = stripCmt(css);
        this.open = this._createMatcher(regOpen);
        this.close = this._createMatcher(regClose);
        this.whitespace = this._createMatcher(regWhitespace);
        this.atImport = this._createAtRule('import');
        this.atCharset = this._createAtRule('charset');
        this.atNamespace = this._createAtRule('namespace');
    },
    parse: function() {
        return this.stylesheet();
    },
    stylesheet: function() {
        return {
            type: 'stylesheet',
            rules: this.rules()
        };
    },
    rules: function() {
        var rule;
        var rules = [];
        this.whitespace();
        while (
            this.input.length &&
            this.input[0] !== '}' &&
            (rule = this.atRule() || this.rule())
        ) {
            rules.push(rule);
            this.whitespace();
        }
        return rules;
    },
    atRule: function() {
        if (this.input[0] !== '@') return;
        return (
            this.atKeyframes() ||
            this.atMedia() ||
            this.atSupports() ||
            this.atImport() ||
            this.atCharset() ||
            this.atNamespace() ||
            this.atFontFace()
        );
    },
    atKeyframes: function() {
        var matched = this.match(regKeyframes);
        if (!matched) return;
        var vendor = matched[1] || '';
        matched = this.match(regIdentifier);
        if (!matched) throw Error('@keyframes missing name');
        var name = matched[1];
        if (!this.open()) throw Error("@keyframes missing '{'");
        var keyframes = [];
        var keyframe;
        while ((keyframe = this.keyframe())) {
            keyframes.push(keyframe);
        }
        if (!this.close()) throw Error("@keyframes missing '}'");
        return {
            type: 'keyframes',
            name: name,
            vendor: vendor,
            keyframes: keyframes
        };
    },
    keyframe: function() {
        var selector = [];
        var matched;
        while ((matched = this.match(regKeyframeSelector))) {
            selector.push(matched[1]);
            this.match(regComma);
        }
        if (!selector.length) return;
        this.whitespace();
        return {
            type: 'keyframe',
            selector: selector.join(', '),
            declarations: this.declarations()
        };
    },
    atSupports: function() {
        var matched = this.match(regSupports);
        if (!matched) return;
        var supports = trim(matched[1]);
        if (!this.open()) throw Error("@supports missing '{'");
        var rules = this.rules();
        if (!this.close()) throw Error("@supports missing '}'");
        return {
            type: 'supports',
            supports: supports,
            rules: rules
        };
    },
    atFontFace: function() {
        var matched = this.match(regFontFace);
        if (!matched) return;
        if (!this.open()) throw Error("@font-face missing '{'");
        var declaration;
        var declarations = [];
        while ((declaration = this.declaration())) {
            declarations.push(declaration);
        }
        if (!this.close()) throw Error("@font-face missing '}'");
        return {
            type: 'font-face',
            declarations: declarations
        };
    },
    atMedia: function() {
        var matched = this.match(regMedia);
        if (!matched) return;
        var media = trim(matched[1]);
        if (!this.open()) throw Error("@media missing '{'");
        this.whitespace();
        var rules = this.rules();
        if (!this.close()) throw Error("@media missing '}'");
        return {
            type: 'media',
            media: media,
            rules: rules
        };
    },
    rule: function() {
        var selector = this.selector();
        if (!selector) throw Error('missing selector');
        return {
            type: 'rule',
            selector: selector,
            declarations: this.declarations()
        };
    },
    declarations: function() {
        var declarations = [];
        if (!this.open()) throw Error("missing '{'");
        this.whitespace();
        var declaration;
        while ((declaration = this.declaration())) {
            declarations.push(declaration);
        }
        if (!this.close()) throw Error("missing '}'");
        this.whitespace();
        return declarations;
    },
    declaration: function() {
        var property = this.match(regProperty);
        if (!property) return;
        property = trim(property[0]);
        if (!this.match(regColon)) throw Error("property missing ':'");
        var value = this.match(regValue);
        this.match(regSemicolon);
        this.whitespace();
        return {
            type: 'declaration',
            property: property,
            value: value ? trim(value[0]) : ''
        };
    },
    selector: function() {
        var matched = this.match(regSelector);
        if (!matched) return;
        return trim(matched[0]);
    },
    match: function(reg) {
        var matched = reg.exec(this.input);
        if (!matched) return;
        this.input = this.input.slice(matched[0].length);
        return matched;
    },
    _createMatcher: function(reg) {
        var _this = this;
        return function() {
            return _this.match(reg);
        };
    },
    _createAtRule: function(name) {
        var reg = new RegExp('^@' + name + '\\s*([^;]+);');
        return function() {
            var matched = this.match(reg);
            if (!matched) return;
            var ret = {
                type: name
            };
            ret[name] = trim(matched[1]);
            return ret;
        };
    }
});
var Compiler = Class({
    initialize: function Compiler(input) {
        var options =
            arguments.length > 1 && arguments[1] !== undefined
                ? arguments[1]
                : {};
        defaults(options, {
            indent: '  '
        });
        this.input = input;
        this.indentLevel = 0;
        this.indentation = options.indent;
    },
    compile: function() {
        return this.stylesheet(this.input);
    },
    stylesheet: function(node) {
        return this.mapVisit(node.rules, '\n\n');
    },
    media: function(node) {
        return (
            '@media ' +
            node.media +
            ' {\n' +
            this.indent(1) +
            this.mapVisit(node.rules, '\n\n') +
            this.indent(-1) +
            '\n}'
        );
    },
    keyframes: function(node) {
        return (
            '@'.concat(node.vendor, 'keyframes ') +
            node.name +
            ' {\n' +
            this.indent(1) +
            this.mapVisit(node.keyframes, '\n') +
            this.indent(-1) +
            '\n}'
        );
    },
    supports: function(node) {
        return (
            '@supports ' +
            node.supports +
            ' {\n' +
            this.indent(1) +
            this.mapVisit(node.rules, '\n\n') +
            this.indent(-1) +
            '\n}'
        );
    },
    keyframe: function(node) {
        return this.rule(node);
    },
    mapVisit: function(nodes, delimiter) {
        var str = '';
        for (var i = 0, len = nodes.length; i < len; i++) {
            var node = nodes[i];
            str += this[camelCase(node.type)](node);
            if (delimiter && i < len - 1) str += delimiter;
        }
        return str;
    },
    fontFace: function(node) {
        return (
            '@font-face {\n' +
            this.indent(1) +
            this.mapVisit(node.declarations, '\n') +
            this.indent(-1) +
            '\n}'
        );
    },
    rule: function(node) {
        return (
            this.indent() +
            node.selector +
            ' {\n' +
            this.indent(1) +
            this.mapVisit(node.declarations, '\n') +
            this.indent(-1) +
            '\n' +
            this.indent() +
            '}'
        );
    },
    declaration: function(node) {
        return this.indent() + node.property + ': ' + node.value + ';';
    },
    import: function(node) {
        return '@import '.concat(node.import, ';');
    },
    charset: function(node) {
        return '@charset '.concat(node.charset, ';');
    },
    namespace: function(node) {
        return '@namespace '.concat(node.namespace, ';');
    },
    indent: function(level) {
        if (level) {
            this.indentLevel += level;
            return '';
        }
        return repeat(this.indentation, this.indentLevel);
    }
});
var stripCmt = function(str) {
    return str.replace(regComments, '');
};

module.exports = exports;
