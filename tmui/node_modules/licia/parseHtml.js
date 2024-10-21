var last = require('./last');
var arrToMap = require('./arrToMap');
var startWith = require('./startWith');
var lowerCase = require('./lowerCase');

exports = function(html, handler) {
    var stack = [];
    var text;
    var lastHtml = html;
    while (html) {
        text = true;
        if (!last(stack) || !SPECIAL[last(stack)]) {
            if (startWith(html, '<!--')) {
                var endIdx = html.indexOf('-->');
                if (endIdx >= 0) {
                    if (handler.comment) {
                        handler.comment(html.substring(4, endIdx));
                    }
                    html = html.substring(endIdx + 3);
                    text = false;
                }
            } else if (startWith(html, '<!')) {
                var match = html.match(regDoctype);
                if (match) {
                    if (handler.text)
                        handler.text(html.substring(0, match[0].length));
                    html = html.substring(match[0].length);
                    text = false;
                }
            } else if (startWith(html, '</')) {
                var _match = html.match(regEndTag);
                if (_match) {
                    html = html.substring(_match[0].length);
                    _match[0].replace(regEndTag, parseEndTag);
                    text = false;
                }
            } else if (startWith(html, '<')) {
                var _match2 = html.match(regStartTag);
                if (_match2) {
                    html = html.substring(_match2[0].length);
                    _match2[0].replace(regStartTag, parseStartTag);
                    text = false;
                }
            }
            if (text) {
                var _endIdx = html.indexOf('<');
                var _text = _endIdx < 0 ? html : html.substring(0, _endIdx);
                html = _endIdx < 0 ? '' : html.substring(_endIdx);
                if (handler.text) handler.text(_text);
            }
        } else {
            var execRes = new RegExp('</'.concat(last(stack), '[^>]*>')).exec(
                html
            );
            if (execRes) {
                var _text2 = html.substring(0, execRes.index);
                html = html.substring(execRes.index + execRes[0].length);
                if (_text2 && handler.text) handler.text(_text2);
            }
            parseEndTag('', last(stack));
        }
        if (lastHtml === html) {
            throw Error('Parse Error: ' + html);
        }
        lastHtml = html;
    }
    parseEndTag();
    function parseStartTag(tag, tagName, rest, unary) {
        tagName = lowerCase(tagName);
        unary = !!unary;
        if (!unary) stack.push(tagName);
        if (handler.start) {
            var attrs = {};
            rest.replace(regAttr, function(all, $1, $2, $3, $4) {
                attrs[$1] = $2 || $3 || $4 || '';
            });
            handler.start(tagName, attrs, unary);
        }
    }
    function parseEndTag(tag, tagName) {
        tagName = lowerCase(tagName);
        var pos;
        if (!tagName) {
            pos = 0;
        } else {
            for (pos = stack.length - 1; pos >= 0; pos--) {
                if (stack[pos] === tagName) break;
            }
        }
        if (pos >= 0) {
            for (var i = stack.length - 1; i >= pos; i--) {
                if (handler.end) handler.end(stack[i]);
            }
            stack.length = pos;
        }
    }
};
var regDoctype = /^<!\s*doctype((?:\s+[\w:]+(?:\s*=\s*(?:(?:"[^"]*")|(?:'[^']*')|[^>\s]+))?)*)\s*(\/?)>/i;
var regEndTag = /^<\/([-A-Za-z0-9_]+)[^>]*>/;
var regStartTag = /^<([-A-Za-z0-9_]+)((?:\s+[-A-Za-z0-9_:@.]+(?:\s*=\s*(?:(?:"[^"]*")|(?:'[^']*')|[^>\s]+))?)*)\s*(\/?)>/i;
var regAttr = /([-A-Za-z0-9_:@.]+)(?:\s*=\s*(?:(?:"((?:\\.|[^"])*)")|(?:'((?:\\.|[^'])*)')|([^>\s]+)))?/g;

var SPECIAL = arrToMap('script,style'.split(','));

module.exports = exports;
