exports =
    typeof window === 'object' &&
    typeof document === 'object' &&
    document.nodeType === 9;

module.exports = exports;
