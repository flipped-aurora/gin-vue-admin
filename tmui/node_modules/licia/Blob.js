var root = require('./root');
var each = require('./each');
exports =
    root.Blob ||
    function Blob(parts, options) {
        options = options || {};
        var blobBuilder = new BlobBuilder();
        each(parts, function(part) {
            blobBuilder.append(part);
        });
        return options.type
            ? blobBuilder.getBlob(options.type)
            : blobBuilder.getBlob();
    };
var BlobBuilder =
    root.BlobBuilder ||
    root.WebKitBlobBuilder ||
    root.MSBlobBuilder ||
    root.MozBlobBuilder;

module.exports = exports;
