package content_type

var (
	// TextPlainCharsetUtf8 文本类型
	TextPlainCharsetUtf8      = "text/plain; charset=utf-8"      // 普通文本
	TextHtmlCharsetUtf8       = "text/html; charset=utf-8"       // HTML 文档
	TextCssCharsetUtf8        = "text/css; charset=utf-8"        // CSS 样式表
	TextCsvCharsetUtf8        = "text/csv; charset=utf-8"        // 逗号分隔值
	TextJavascriptCharsetUtf8 = "text/javascript; charset=utf-8" // JavaScript 文件

	// ApplicationJsonCharsetUtf8 应用类型
	ApplicationJsonCharsetUtf8               = "application/json; charset=utf-8"                  // JSON 格式
	ApplicationXmlCharsetUtf8                = "application/xml; charset=utf-8"                   // XML 格式
	ApplicationZip                           = "application/zip"                                  // ZIP 压缩文件
	ApplicationPdf                           = "application/pdf"                                  // PDF 文档
	ApplicationOctetStream                   = "application/octet-stream"                         // 二进制流
	ApplicationXWwwFormUrlencodedCharsetUtf8 = "application/x-www-form-urlencoded; charset=utf-8" // 表单数据
	ApplicationJavascriptCharsetUtf8         = "application/javascript; charset=utf-8"            // JavaScript 文件

	// ImageJpeg 图像类型
	ImageJpeg   = "image/jpeg"    // JPEG 图像
	ImagePng    = "image/png"     // PNG 图像
	ImageGif    = "image/gif"     // GIF 动画
	ImageBmp    = "image/bmp"     // BMP 图像
	ImageSvgXml = "image/svg+xml" // SVG 矢量图

	// AudioMpeg 音频类型
	AudioMpeg = "audio/mpeg" // MPEG 音频
	AudioWav  = "audio/wav"  // WAV 音频
	AudioOgg  = "audio/ogg"  // OGG 音频

	// VideoMp4 视频类型
	VideoMp4  = "video/mp4"  // MP4 视频
	VideoMpeg = "video/mpeg" // MPEG 视频
	VideoOgg  = "video/ogg"  // OGG 视频
	VideoWebm = "video/webm" // WebM 视频
)
