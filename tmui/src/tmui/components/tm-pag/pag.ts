
const e = "android" === wx.getSystemInfoSync().platform,
	t = `${wx.env.USER_DATA_PATH}/pag/`,
	i = wx.getFileSystemManager();
let r;
try {
	r = performance.now.bind(performance)
} catch (e) {
	r = Date.now
}
class s {
	constructor() {
		this.startTime = r(), this.markers = {}
	}
	reset() {
		this.startTime = r(), this.markers = {}
	}
	mark(e) {
		e ? Object.keys(this.markers).find((t => t === e)) ? console.log(
			`Clock.mark(): The specified marker name '${e}' already exists!`) : this.markers[e] = r() : console
				.log("Clock.mark(): An empty marker name was specified!")
	}
	measure(e, t) {
		let i, s;
		if (e) {
			if (!Object.keys(this.markers).find((t => t === e))) return console.log(
				`Clock.measure(): The specified makerFrom '${e}' does not exist!`), 0;
			i = this.markers[e]
		} else i = this.startTime;
		if (t) {
			if (!Object.keys(this.markers).find((e => e === t))) return console.log(
				`Clock.measure(): The specified makerTo '${t}' does not exist!`), 0;
			s = this.markers[t]
		} else s = r();
		return s - i
	}
}

function a(e) {
	let t = Object.getOwnPropertyNames(e.prototype).filter((t => "constructor" !== t && "function" == typeof e
		.prototype[t]));
	const i = (e, t) => {
		const i = e[t];
		e[t] = function (...e) {
			if (!this.destroyed) return i.call(this, ...e);
			console.error(`Don't call ${t} of the PAGView that is destroyed.`)
		}
	};
	t.forEach((t => i(e.prototype, t)))
}
const n = {
	alpha: !0,
	depth: !1,
	stencil: !1,
	antialias: !1
};
var o = (e => (e[e.Normal = 0] = "Normal", e[e.Multiply = 1] = "Multiply", e[e.Screen = 2] = "Screen", e[e.Overlay =
	3] = "Overlay", e[e.Darken = 4] = "Darken", e[e.Lighten = 5] = "Lighten", e[e.ColorDodge = 6] =
	"ColorDodge", e[e.ColorBurn = 7] = "ColorBurn", e[e.HardLight = 8] = "HardLight", e[e.SoftLight = 9] =
	"SoftLight", e[e.Difference = 10] = "Difference", e[e.Exclusion = 11] = "Exclusion", e[e.Hue = 12] = "Hue",
	e[e.Saturation = 13] = "Saturation", e[e.Color = 14] = "Color", e[e.Luminosity = 15] = "Luminosity", e[e
		.DestinationIn = 21] = "DestinationIn", e[e.DestinationOut = 22] = "DestinationOut", e[e
			.DestinationATop = 23] = "DestinationATop", e[e.SourceIn = 24] = "SourceIn", e[e.SourceOut = 25] =
	"SourceOut", e[e.Xor = 26] = "Xor", e))(o || {}),
	h = (e => (e[e.None = 0] = "None", e[e.Linear = 1] = "Linear", e[e.Bezier = 2] = "Bezier", e[e.Hold = 3] = "Hold",
		e))(h || {});
const l = navigator && /(ios|ipad|iphone)/.test(navigator.userAgent.toLowerCase()),
	d = e => {
		const t = [];
		for (let i = 0; i < e.length; i++) t.push(e.charCodeAt(i));
		return t
	},
	c = e => [e >> 24, e >> 16 & 255, e >> 8 & 255, 255 & e],
	u = (e, ...t) => {
		let i = 8,
			r = t.length;
		const s = r;
		for (; r;) r -= 1, i += t[r].byteLength;
		const a = new Uint8Array(i);
		for (a[0] = i >> 24 & 255, a[1] = i >> 16 & 255, a[2] = i >> 8 & 255, a[3] = 255 & i, a.set(e, 4), r = 0, i =
			8; r < s; ++r) a.set(t[r], i), i += t[r].byteLength;
		return a
	};
class m {
	constructor(e) {
		this.param = e
	}
	ftyp() {
		return u(d("ftyp"), new Uint8Array(d("isom")), new Uint8Array([0, 0, 0, 1]), new Uint8Array(d("isom")),
			new Uint8Array(d("iso2")), new Uint8Array(d("avc1")), new Uint8Array(d("mp41")))
	}
	moov() {
		const e = this.param.tracks.map((e => this.trak(e))).reverse();
		return u(d("moov"), this.mvhd(), ...e, this.mvex())
	}
	moof() {
		return u(d("moof"), this.mfhd(), this.traf())
	}
	mdat() {
		const e = new Uint8Array(this.param.track.len);
		let t = 0;
		return this.param.videoSequence.headers.forEach((i => {
			e.set(new Uint8Array(i.data.data()), t), t += i.length
		})), this.param.videoSequence.frames.forEach(((i, r) => {
			e.set(new Uint8Array(i.fileBytes.data.data()), t), t += i.fileBytes.length
		})), u(d("mdat"), e)
	}
	mvhd() {
		return u(d("mvhd"), new Uint8Array([0, 0, 0, 0, ...c(2082873600), ...c(2082873600), ...c(this.param
			.timescale), ...c(this.param.duration), 0, 1, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 64, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 2
		]))
	}
	trak(e) {
		return u(d("trak"), this.tkhd(e), this.edts(e), this.mdia(e))
	}
	tkhd(e) {
		return u(d("tkhd"), new Uint8Array([0, 0, 0, 1, ...c(2082873600), ...c(2082873600), ...c(e.id), 0, 0, 0, 0,
			...c(e.duration), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 64, 0, 0, 0, e.width >>
			8 & 255, 255 & e.width, 0, 0, e.height >> 8 & 255, 255 & e.height, 0, 0
		]))
	}
	edts(e) {
		return u(d("edts"), this.elst(e))
	}
	elst(e) {
		return u(d("elst"), new Uint8Array([0, 0, 0, 0, ...c(1), ...c(e.duration), ...c(e.implicitOffset * Math
			.floor(e.duration / e.samples.length)), 0, 1, 0, 0]))
	}
	mdia(e) {
		return u(d("mdia"), this.mdhd(), this.hdlr(), this.minf(e))
	}
	mdhd() {
		return u(d("mdhd"), new Uint8Array([0, 0, 0, 0, ...c(2082873600), ...c(2082873600), ...c(this.param
			.timescale), ...c(0), 85, 196, 0, 0]))
	}
	hdlr() {
		return u(d("hdlr"), new Uint8Array([0, 0, 0, 0, 0, 0, 0, 0, 118, 105, 100, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			0, 0, 0, 86, 105, 100, 101, 111, 72, 97, 110, 100, 108, 101, 114, 0
		]))
	}
	minf(e) {
		return u(d("minf"), this.vmhd(), this.dinf(), this.stbl(e))
	}
	vmhd() {
		return u(d("vmhd"), new Uint8Array([0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0]))
	}
	dinf() {
		return u(d("dinf"), this.dref())
	}
	dref() {
		return u(d("dref"), new Uint8Array([0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 12, 117, 114, 108, 32, 0, 0, 0, 1]))
	}
	stbl(e) {
		return u(d("stbl"), this.stsd(e), this.stts(e), this.ctts(e), this.stss(e), this.stsc(), this.stsz(), this
			.stco())
	}
	stsd(e) {
		return u(d("stsd"), new Uint8Array([0, 0, 0, 0, 0, 0, 0, 1]), this.avc1(e))
	}
	avc1(e) {
		const t = [0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, e.width >> 8 & 255, 255 &
			e.width, e.height >> 8 & 255, 255 & e.height, 0, 72, 0, 0, 0, 72, 0, 0, 0, 0, 0, 0, 0, 1, 18, 98,
			105, 110, 101, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0,
			24, 255, 255
		];
		return u(d("avc1"), new Uint8Array(t), this.avcc(e))
	}
	avcc(e) {
		let t = [],
			i = [];
		e.sps.forEach((e => {
			const i = e.length - 4;
			t.push(i >>> 8 & 255), t.push(255 & i), t = t.concat(Array.prototype.slice.call(
				new Uint8Array(e.data.data(), 4)))
		})), e.pps.forEach((e => {
			const t = e.length - 4;
			i.push(t >>> 8 & 255), i.push(255 & t), i = i.concat(Array.prototype.slice.call(
				new Uint8Array(e.data.data(), 4)))
		}));
		const r = [1, t[3], t[4], t[5], 255, 224 | e.sps.length].concat(t).concat([e.pps.length]).concat(i);
		return u(d("avcC"), new Uint8Array(r))
	}
	stts(e) {
		return u(d("stts"), new Uint8Array([0, 0, 0, 0, ...c(1), ...c(e.samples.length), ...c(Math.floor(e
			.duration / e.samples.length))]))
	}
	ctts(e) {
		const t = e.pts.length,
			i = Math.floor(e.duration / t),
			r = [0, 0, 0, 0, ...c(t)];
		for (let s = 0; s < t; s++) {
			r.push(...c(1));
			const t = s * i,
				a = (e.pts[s] + e.implicitOffset) * i;
			r.push(...c(a - t))
		}
		return u(d("ctts"), new Uint8Array(r))
	}
	stss(e) {
		const t = e.samples.filter((e => e.flags.isKeyFrame)).map((e => e.index + 1)),
			i = [0, 0, 0, 0, ...c(t.length)];
		return t.forEach((e => {
			i.push(...c(e))
		})), u(d("stss"), new Uint8Array(i))
	}
	stsc() {
		return u(d("stsc"), new Uint8Array([0, 0, 0, 0, 0, 0, 0, 0]))
	}
	stsz() {
		return u(d("stsz"), new Uint8Array([0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0]))
	}
	stco() {
		return u(d("stco"), new Uint8Array([0, 0, 0, 0, 0, 0, 0, 0]))
	}
	mvex() {
		const e = this.param.tracks.map((e => this.trex(e))).reverse();
		return u(d("mvex"), ...e)
	}
	trex(e) {
		return u(d("trex"), new Uint8Array([0, 0, 0, 0, e.id >> 24, e.id >> 16 & 255, e.id >> 8 & 255, 255 & e.id,
			0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 1
		]))
	}
	mfhd() {
		return u(d("mfhd"), new Uint8Array([0, 0, 0, 0, this.param.sequenceNumber >> 24, this.param
			.sequenceNumber >> 16 & 255, this.param.sequenceNumber >> 8 & 255, 255 & this.param
				.sequenceNumber
		]))
	}
	traf() {
		const e = this.sdtp();
		return this.param.offset = e.length + 72, u(d("traf"), this.tfhd(), this.tfdt(), this.trun(), e)
	}
	tfhd() {
		return u(d("tfhd"), new Uint8Array([0, 0, 0, 0, ...c(this.param.track.id)]))
	}
	tfdt() {
		return u(d("tfdt"), new Uint8Array([0, 0, 0, 0, ...c(this.param.baseMediaDecodeTime)]))
	}
	trun() {
		const e = (this.param.track.samples || []).length,
			t = 12 + 16 * e;
		this.param.offset += 8 + t;
		const i = [0, 0, 15, 1, ...c(e), ...c(this.param.offset)];
		return this.param.track.samples.forEach((e => {
			const {
				duration: t,
				size: r,
				flags: s,
				cts: a
			} = e;
			i.push(...c(t)), i.push(...c(r)), i.push(s.isLeading << 2 | s.dependsOn), i.push(s
				.isDependedOn << 6 | s.hasRedundancy << 4 | 0 | s.isNonSync), i.push(61440 & s
					.degradPrio), i.push(15 & s.degradPrio), i.push(...c(a))
		})), u(d("trun"), new Uint8Array(i))
	}
	sdtp() {
		const e = new Uint8Array(4 + this.param.track.samples.length);
		return this.param.track.samples.forEach(((t, i) => {
			e[i + 4] = t.flags.dependsOn << 4 | t.flags.isDependedOn << 2 | t.flags.hasRedundancy
		})), u(d("sdtp"), e)
	}
}
var f = Object.defineProperty,
	p = Object.getOwnPropertySymbols,
	g = Object.prototype.hasOwnProperty,
	y = Object.prototype.propertyIsEnumerable,
	w = (e, t, i) => t in e ? f(e, t, {
		enumerable: !0,
		configurable: !0,
		writable: !0,
		value: i
	}) : e[t] = i,
	v = (e, t) => {
		for (var i in t || (t = {})) g.call(t, i) && w(e, i, t[i]);
		if (p)
			for (var i of p(t)) y.call(t, i) && w(e, i, t[i]);
		return e
	};
const b = e => {
	const t = l ? T(e) : e,
		i = S(t);
	if (!i || 0 === i.len) throw new Error("mp4Track is empty");
	const r = {
		offset: 0,
		tracks: [i],
		track: i,
		duration: i.duration,
		timescale: i.timescale,
		sequenceNumber: 1,
		baseMediaDecodeTime: 0,
		nalusBytesLen: i.len,
		videoSequence: t
	},
		s = new m(r);
	return (e => {
		let t = 0;
		for (const i of e) t += i.byteLength;
		const i = new Uint8Array(t);
		let r = 0;
		for (const t of e) i.set(t, r), r += t.byteLength;
		return i
	})([s.ftyp(), s.moov(), s.moof(), s.mdat()])
},
	S = e => {
		if (e.headers.length < 2) throw new Error("Bad header data in video sequence!");
		if (0 === e.frames.length) throw new Error("There is no frame data in the video sequence!");
		const t = {
			id: 1,
			type: "video",
			timescale: 6e3,
			duration: Math.floor(6e3 * e.frames.length / e.frameRate),
			width: e.getVideoWidth(),
			height: e.getVideoHeight(),
			sps: [e.headers[0]],
			pps: [e.headers[1]],
			implicitOffset: E(e.frames),
			len: 0,
			pts: [],
			samples: []
		},
			i = e.headers.reduce(((e, t) => e + t.length), 0),
			r = t.duration / e.frames.length;
		return e.frames.forEach(((e, s) => {
			var a;
			let n = null != (a = e.fileBytes.length) ? a : 0;
			0 === s && (n += i), t.len += n, t.pts.push(e.frame), t.samples.push({
				index: s,
				size: n,
				duration: r,
				cts: (e.frame + t.implicitOffset - s) * r,
				flags: {
					isKeyFrame: e.isKeyframe,
					isNonSync: e.isKeyframe ? 0 : 1,
					dependsOn: e.isKeyframe ? 2 : 1,
					isLeading: 0,
					isDependedOn: 0,
					hasRedundancy: 0,
					degradPrio: 0
				}
			})
		})), t
	},
	E = e => Math.max(...e.map(((e, t) => t - e.frame))),
	T = e => {
		const t = e.frames.length;
		for (let i = 0; i < e.frames.length; i++) {
			const r = v({}, e.frames[i]);
			if (r.isKeyframe && i > 0) break;
			r.frame += t, e.frames.push(r)
		}
		return e
	},
	x = (e, t, i) => {
		const r = e.createProgram();
		if (!r) throw new Error("Failed to create program.");
		const s = R(e, t, e.VERTEX_SHADER);
		if (!s) throw new Error("Failed to create vertex shader.");
		e.attachShader(r, s);
		const a = R(e, i, e.FRAGMENT_SHADER);
		if (!a) throw new Error("Failed to create fragment shader.");
		e.attachShader(r, a), e.linkProgram(r);
		const n = e.getProgramInfoLog(r);
		n && console.log(n);
		const o = e.getShaderInfoLog(s);
		o && console.log(o);
		const h = e.getShaderInfoLog(a);
		return h && console.log(h), r
	},
	R = (e, t, i) => {
		const r = e.createShader(i);
		if (!r) throw new Error("Failed to create shader.");
		return e.shaderSource(r, t), e.compileShader(r), r
	},
	P = e => e.replace(/^\s+|\s+$/g, ""),
	A = e => {
		const t = e.createTexture();
		if (!t) throw new Error("Failed to create texture.");
		return e.bindTexture(e.TEXTURE_2D, t), e.texParameteri(e.TEXTURE_2D, e.TEXTURE_WRAP_S, e.CLAMP_TO_EDGE), e
			.texParameteri(e.TEXTURE_2D, e.TEXTURE_WRAP_T, e.CLAMP_TO_EDGE), e.texParameteri(e.TEXTURE_2D, e
				.TEXTURE_MIN_FILTER, e.NEAREST), e.texParameteri(e.TEXTURE_2D, e.TEXTURE_MAG_FILTER, e.NEAREST), t
	};
let D = {};
const k = (e, t, i, r = !1) => {
	var s;
	t in D || (D[t] = []), null == (s = D[t]) || s.push({
		node: e,
		handler: i,
		capture: r
	}), e.addEventListener(t, i, r)
};
var F = Object.defineProperty,
	V = Object.getOwnPropertyDescriptor;
const C = navigator && /MicroMessenger/i.test(navigator.userAgent),
	L = async e => {
		C && window.WeixinJSBridge && await new Promise((e => {
			window.WeixinJSBridge.invoke("getNetworkType", {}, (() => {
				e()
			}), (() => {
				e()
			}))
		}));
		try {
			await e.play()
		} catch (e) {
			throw new Error(e.message)
		}
	};
let B = class {
	constructor(e) {
		this.destroyed = !1, this.frameRate = 0, this._duration = e.frameCount / e.frameRate, this.frameRate = e
			.frameRate
	}
	static create(e) {
		const t = new B(e),
			i = t.load(e);
		return {
			videoReader: t,
			debugData: i
		}
	}
	getVideoElement() {
		return this.videoElement
	}
	progress() {
		return Math.round(this.videoElement.currentTime / this._duration * 100) / 100
	}
	duration() {
		return this._duration
	}
	currentTime() {
		return this.videoElement.currentTime || 0
	}
	start() {
		return L(this.videoElement)
	}
	pause() {
		var e;
		null == (e = this.videoElement) || e.pause()
	}
	seek(e) {
		return new Promise((t => {
			const i = () => {
				var e, r, s, a, n, o, h;
				e = this.videoElement, s = i, (r = "seeked") in D && (s ? (null == (a = D[r]) ||
					a.filter((({
						node: t,
						handler: i
					}) => t === e && i === s)).forEach((({
						node: e,
						handler: t,
						capture: i
					}) => e.removeEventListener(r, t, i))), D[r] = null == (n = D[r]) ?
						void 0 : n.filter((({
							node: t,
							handler: i
						}) => !(t === e && i === s)))) : (null == (o = D[r]) || o.filter((({
							node: t
						}) => t === e)).forEach((({
							node: e,
							handler: t,
							capture: i
						}) => e.removeEventListener(r, t, i))), D[r] = null == (h = D[r]) ?
							void 0 : h.filter((({
								node: t
							}) => t !== e)))), t()
			};
			k(this.videoElement, "seeked", i), this.videoElement.currentTime = e
		}))
	}
	addListener(e, t) {
		k(this.videoElement, e, t)
	}
	removeAllListeners() {
		var e;
		e = this.videoElement, Object.keys(D).forEach((t => {
			var i, r;
			const s = t;
			null == (i = D[s]) || i.filter((({
				node: t
			}) => t === e)).forEach((({
				node: e,
				handler: t,
				capture: i
			}) => e.removeEventListener(s, t, i))), D[s] = null == (r = D[s]) ? void 0 : r
				.filter((({
					node: t
				}) => t !== e))
		}))
	}
	getFrameData(e) { }
	clearCallback() { }
	destroy() {
		this.removeAllListeners(), this.videoElement = void 0, this.destroyed = !0
	}
	load(e) {
		this.videoElement = document.createElement("video"), this.videoElement.style.display = "none", this
			.videoElement.muted = !0, this.videoElement.playsInline = !0;
		const t = new s,
			i = b(e);
		return t.mark("coverMP4"), this.videoElement.src = URL.createObjectURL(new Blob([i], {
			type: "video/mp4"
		})), this.videoElement.load(), {
			coverMP4: t.measure("", "coverMP4")
		}
	}
};
B = ((e, t, i, r) => {
	for (var s, a = r > 1 ? void 0 : r ? V(t, i) : t, n = e.length - 1; n >= 0; n--)(s = e[n]) && (a = (r ? s(t,
		i, a) : s(a)) || a);
	return r && a && F(t, i, a), a
})([a], B);
var _ = Object.defineProperty,
	U = Object.getOwnPropertyDescriptor;
const M = (e, t) => {
	const i = new ArrayBuffer(t.data.byteLength);
	return new Uint8Array(i).set(new Uint8Array(t.data)), {
		id: e,
		data: i,
		width: t.width,
		height: t.height
	}
};
let I = class extends B {
	constructor() {
		super(...arguments), this.started = !1, this.frameDataBuffers = [], this.getFrameDataLooping = !1, this
			.getFrameDataCallback = null, this.getFrameDataLoopTimer = null, this.currentFrame = 0, this
				.seeking = !1
	}
	static create(e) {
		const t = new I(e),
			i = t.load(e);
		return {
			videoReader: t,
			debugData: i
		}
	}
	getVideoElement() {
		throw new Error("WeChat mini program does not support video element as decoder!")
	}
	progress() {
		return 0
	}
	duration() {
		return 0
	}
	currentTime() {
		return 0
	}
	async start() {
		this.started || (this.started = !0, await this.loadedPromise, this.startGetFrameDataLoop())
	}
	pause() { }
	async seek(e) {
		var t;
		const i = Math.floor(e * this.frameRate);
		if (this.frameDataBuffers.length > 0) {
			const e = this.frameDataBuffers.findIndex((e => e.id === i));
			if (-1 !== e) return void (this.frameDataBuffers = this.frameDataBuffers.slice(e));
			this.frameDataBuffers = []
		}
		i !== this.currentFrame && (this.seeking = !0, await (null == (t = this.videoDecoder) ? void 0 : t.seek(
			e)), this.seeking = !1, this.currentFrame = i)
	}
	getFrameData(e) {
		if (this.frameDataBuffers.length <= 2 && !this.getFrameDataLooping && this.startGetFrameDataLoop(),
			0 === this.frameDataBuffers.length) return void (this.getFrameDataCallback = e);
		const t = this.frameDataBuffers.shift();
		t ? e(t) : this.getFrameDataCallback = e
	}
	addListener(e, t) {
		var i;
		const r = e;
		null == (i = this.videoDecoder) || i.on(r, t)
	}
	removeAllListeners() { }
	destroy() {
		var e;
		this.clearFrameDataLoop(), this.clearCallback(), null == (e = this.videoDecoder) || e.remove(), this
			.videoDecoder = void 0, this.mp4Path && (e => {
				try {
					i.accessSync(e), i.unlinkSync(e)
				} catch (e) {
					console.error(e)
				}
			})(this.mp4Path), this.destroyed = !0
	}
	clearCallback() {
		this.getFrameDataCallback = null
	}
	load(e) {
		const r = new s;
		this.mp4Path = `${t}${Date.now()}.mp4`, (e => {
			try {
				i.accessSync(e)
			} catch (t) {
				try {
					i.mkdirSync(e)
				} catch (e) {
					console.error(t)
				}
			}
		})(t), r.mark("createDir");
		const a = b(e);
		return r.mark("coverMP4"), ((e, t) => {
			try {
				i.writeFileSync(e, t, "utf8")
			} catch (e) {
				throw new Error(e)
			}
		})(this.mp4Path, a.buffer), r.mark("writeFile"), this.videoDecoder = wx.createVideoDecoder(), r
			.mark("createDecoder"), this.loadedPromise = this.videoDecoder.start({
				source: this.mp4Path,
				mode: 1
			}), this.videoDecoder.on("ended", (() => {
				var e;
				null == (e = this.videoDecoder) || e.seek(0).then((() => {
					this.currentFrame = 0
				}))
			})), {
			createDir: r.measure("", "createDir"),
			coverMP4: r.measure("createDir", "coverMP4"),
			writeFile: r.measure("coverMP4", "writeFile"),
			createDecoder: r.measure("writeFile", "createDecoder")
		}
	}
	startGetFrameDataLoop() {
		this.getFrameDataLooping = !0, this.getFrameDataLoopTimer = setInterval((() => {
			this.getFrameDataLoop()
		}), 2)
	}
	getFrameDataLoop() {
		if (this.seeking) return;
		if (!this.videoDecoder) throw this.clearFrameDataLoop(), new Error("VideoDecoder is not ready!");
		if (this.frameDataBuffers.length >= 6) return this.getFrameDataLooping = !1, void this
			.clearFrameDataLoop();
		const e = this.videoDecoder.getFrameData();
		null !== e && (this.getFrameDataCallback ? (this.getFrameDataCallback(M(this.currentFrame, e)), this
			.getFrameDataCallback = null) : this.frameDataBuffers.push(M(this.currentFrame, e)), this
				.currentFrame += 1)
	}
	clearFrameDataLoop() {
		this.getFrameDataLoopTimer && (clearInterval(this.getFrameDataLoopTimer), this.getFrameDataLoopTimer =
			null), this.getFrameDataLooping = !1
	}
};
I = ((e, t, i, r) => {
	for (var s, a = r > 1 ? void 0 : r ? U(t, i) : t, n = e.length - 1; n >= 0; n--)(s = e[n]) && (a = (r ? s(t,
		i, a) : s(a)) || a);
	return r && a && _(t, i, a), a
})([a], I);
var O = (e => (e.Canvas2D = "2d", e.WebGL = "WebGL", e))(O || {}),
	q = (e => (e.onAnimationStart = "onAnimationStart", e.onAnimationEnd = "onAnimationEnd", e.onAnimationCancel =
		"onAnimationCancel", e.onAnimationRepeat = "onAnimationRepeat", e.onAnimationUpdate = "onAnimationUpdate", e
			.onAnimationPlay = "onAnimationPlay", e.onAnimationPause = "onAnimationPause", e))(q || {}),
	G = (e => (e.None = "None", e.Stretch = "Stretch", e.LetterBox = "LetterBox", e.Zoom = "Zoom", e))(G || {}),
	N = Object.freeze({
		__proto__: null,
		RenderingMode: O,
		EventName: q,
		ScaleMode: G
	});
const z =
	"\n      attribute vec2 a_position;\n      attribute vec2 a_texCoord;\n      \n      uniform vec2 u_resolution;\n      uniform vec2 u_scale;\n      \n      varying vec2 v_texCoord;\n    \n      \n      void main() {\n         // convert the rectangle from pixels to 0.0 to 1.0\n         vec2 zeroToOne = a_position / u_resolution;\n      \n         // convert from 0->1 to 0->2\n         vec2 zeroToTwo = zeroToOne * 2.0;\n      \n         // convert from 0->2 to -1->+1 (clipspace)\n         vec2 clipSpace = zeroToTwo - 1.0;\n      \n         gl_Position = vec4(clipSpace * vec2(1, -1), 0, 1);\n      \n         // pass the texCoord to the fragment shader\n         // The GPU will interpolate this value between points.\n         v_texCoord = a_texCoord / u_scale;\n      }\n        ";
class W {
	constructor() {
		this.listenersMap = {}
	}
	on(e, t) {
		void 0 === this.listenersMap[e] && (this.listenersMap[e] = []), this.listenersMap[e].push(t)
	}
	off(e, t) {
		const i = this.listenersMap[e];
		if (void 0 === i) return;
		if (void 0 === t) return void delete this.listenersMap[e];
		const r = i.findIndex((e => e === t));
		i.splice(r, 1)
	}
	emit(e, ...t) {
		const i = this.listenersMap[e];
		if (void 0 === i || i.length < 1) return !1;
		for (const e of i) e(...t);
		return !0
	}
}
var X = Object.defineProperty,
	H = Object.getOwnPropertyDescriptor,
	j = Object.getOwnPropertySymbols,
	K = Object.prototype.hasOwnProperty,
	Y = Object.prototype.propertyIsEnumerable,
	$ = (e, t, i) => t in e ? X(e, t, {
		enumerable: !0,
		configurable: !0,
		writable: !0,
		value: i
	}) : e[t] = i,
	Z = (e, t) => {
		for (var i in t || (t = {})) K.call(t, i) && $(e, i, t[i]);
		if (j)
			for (var i of j(t)) Y.call(t, i) && $(e, i, t[i]);
		return e
	};
let J = class {
	constructor(e, t, i) {
		this.canvasSize = {
			width: 0,
			height: 0
		}, this.playing = !1, this.viewportSize = {
			x: 0,
			y: 0,
			width: 0,
			height: 0,
			scaleX: 1,
			scaleY: 1
		}, this.destroyed = !1, this.renderTimer = null, this.repeatCount = 0, this.viewScaleMode = G
			.LetterBox, this.debugData = {
				FPS: 0,
				decodePAGFile: 0,
				createDir: 0,
				coverMP4: 0,
				writeFile: 0,
				createDecoder: 0,
				getFrame: 0,
				draw: 0
			};
		const r = e.getVideoSequence();
		if (!r) throw new Error("PAGFile has no BMP video sequence!");
		delete r.composition, this.videoSequence = r, this.canvas = t, this.videoParam = ((e, t) => ({
			width: e.mainComposition.width,
			height: e.mainComposition.height,
			hasAlpha: t.alphaStartX > 0 || t.alphaStartY > 0,
			alphaStartX: t.alphaStartX,
			alphaStartY: t.alphaStartY,
			sequenceWidth: t.width,
			sequenceHeight: t.height,
			MP4Width: (t.width + t.alphaStartX) % 2 == 0 ? t.width + t.alphaStartX : t.width + t
				.alphaStartX + 1,
			MP4Height: (t.height + t.alphaStartY) % 2 == 0 ? t.height + t.alphaStartY : t.height + t
				.alphaStartY + 1
		}))(e, r), this.eventManager = new W, this.renderingMode = i.renderingMode || O.WebGL, this
			.updateSize(i.useScale), this.setScaleMode()
	}
	isPlaying() {
		return this.playing
	}
	isDestroyed() {
		return this.destroyed
	}
	duration() {
		return this.videoSequence.frameCount / this.videoSequence.frameRate
	}
	frameRate() {
		return this.videoSequence.frameRate
	}
	setRepeatCount(e = 1) {
		this.repeatCount = e < 0 ? -1 : e - 1
	}
	addListener(e, t) {
		return this.eventManager.on(e, t)
	}
	removeListener(e, t) {
		return this.eventManager.off(e, t)
	}
	scaleMode() {
		return this.viewScaleMode
	}
	setScaleMode(e = G.LetterBox) {
		switch (this.viewScaleMode = e, e) {
			case G.None:
				this.viewportSize = {
					x: 0,
					y: this.renderingMode === O.WebGL ? this.canvas.height - this.videoParam.height : 0,
					width: this.videoParam.width,
					height: this.videoParam.height,
					scaleX: 1,
					scaleY: 1
				};
				break;
			case G.Stretch:
				this.viewportSize = {
					x: 0,
					y: 0,
					width: this.canvas.width,
					height: this.canvas.height,
					scaleX: this.canvas.width / this.videoParam.sequenceWidth,
					scaleY: this.canvas.height / this.videoParam.sequenceHeight
				};
				break;
			case G.LetterBox: {
				const e = this.canvas.width / this.videoParam.sequenceWidth,
					t = this.canvas.height / this.videoParam.sequenceHeight,
					i = Math.min(e, t);
				this.viewportSize = {
					x: (this.canvas.width - this.videoParam.sequenceWidth * i) / 2,
					y: (this.canvas.height - this.videoParam.sequenceHeight * i) / 2,
					width: this.videoParam.sequenceWidth * i,
					height: this.videoParam.sequenceHeight * i,
					scaleX: i,
					scaleY: i
				}
			}
				break;
			case G.Zoom: {
				const e = this.canvas.width / this.videoParam.sequenceWidth,
					t = this.canvas.height / this.videoParam.sequenceHeight,
					i = Math.max(e, t);
				this.viewportSize = {
					x: (this.canvas.width - this.videoParam.sequenceWidth * i) / 2,
					y: (this.canvas.height - this.videoParam.sequenceHeight * i) / 2,
					width: this.videoParam.sequenceWidth * i,
					height: this.videoParam.sequenceHeight * i,
					scaleX: i,
					scaleY: i
				}
			}
		}
	}
	updateSize(e = !0) {
		if (!this.canvas) throw new Error("Canvas element is not found!");
		let t;
		const i = getComputedStyle(this.canvas),
			r = {
				width: Number(i.width.replace("px", "")),
				height: Number(i.height.replace("px", ""))
			};
		if (r.width > 0 && r.height > 0) t = r;
		else {
			const e = {
				width: Number(this.canvas.style.width.replace("px", "")),
				height: Number(this.canvas.style.height.replace("px", ""))
			};
			t = e.width > 0 && e.height > 0 ? e : {
				width: this.canvas.width,
				height: this.canvas.height
			}
		}
		if (!e) return this.canvas.width = this.canvas.width || t.width, void (this.canvas.height = this.canvas
			.height || t.height);
		this.canvas.style.width = `${t.width}px`, this.canvas.style.height = `${t.height}px`, this.canvas
			.width = t.width * window.devicePixelRatio, this.canvas.height = t.height * window.devicePixelRatio
	}
	getDebugData() {
		return this.debugData
	}
	setDebugData(e) {
		this.debugData = Z(Z({}, this.debugData), e)
	}
	loadContext() { }
	clearRender() { }
};
J = ((e, t, i, r) => {
	for (var s, a = r > 1 ? void 0 : r ? H(t, i) : t, n = e.length - 1; n >= 0; n--)(s = e[n]) && (a = (r ? s(t,
		i, a) : s(a)) || a);
	return r && a && X(t, i, a), a
})([a], J);
var Q = Object.defineProperty,
	ee = Object.getOwnPropertyDescriptor;
navigator && /MicroMessenger/i.test(navigator.userAgent);
let te = class extends J {
	constructor(e, t, i) {
		super(e, t, i), this.fpsBuffer = [], this.currentFrame = -1, this.needSeek = !1, this.videoReader = this
			.createVideoReader(this.videoSequence)
	}
	async play() {
		this.playing || (this.playing = !0, await this.videoReader.start(), await this.flushLoop(), 0 === this
			.getProgress() && this.eventManager.emit(q.onAnimationStart), this.eventManager.emit(q
				.onAnimationPlay))
	}
	pause() {
		this.playing && (this.videoReader.pause(), this.clearTimer(), this.playing = !1, this.eventManager.emit(
			q.onAnimationPause))
	}
	stop() {
		this.videoReader.pause(), this.videoReader.seek(0), this.clearRender(), this.playing = !1, this
			.eventManager.emit(q.onAnimationCancel)
	}
	destroy() {
		this.clearTimer(), this.clearRender(), this.canvas = null, this.videoReader.destroy(), this
			.destroyed = !0
	}
	getProgress() {
		return this.currentFrame / this.videoSequence.frameCount
	}
	setProgress(e) {
		if (e < 0 || e > 1) throw new Error("progress must be between 0.0 and 1.0!");
		const t = Math.round(e * this.videoSequence.frameCount);
		this.currentFrame !== t && (this.needSeek = !0, this.currentFrame = t)
	}
	flush() {
		return this.flushInternal(!0)
	}
	draw() { }
	createVideoReader(e) {
		const {
			videoReader: t,
			debugData: i
		} = B.create(e);
		return this.setDebugData(i), l || t.addListener("ended", (() => {
			this.repeat()
		})), t
	}
	async repeat() {
		return 0 === this.repeatCount ? (this.setProgress(1), await this.flushInternal(!0), this.videoReader
			.pause(), this.clearTimer(), this.playing = !1, this.eventManager.emit("onAnimationEnd"), !1) :
			(this.repeatCount -= 1, l ? await this.videoReader.seek(0) : this.videoReader.start(), this
				.eventManager.emit("onAnimationRepeat"), !0)
	}
	flushLoop() {
		return this.renderTimer = window.requestAnimationFrame((() => {
			this.flushLoop()
		})), l && this.duration() - this.videoReader.currentTime() <= 1 / this.frameRate() && this.repeat(),
			this.flushInternal(!1)
	}
	clearTimer() {
		this.renderTimer && (window.cancelAnimationFrame(this.renderTimer), this.renderTimer = null)
	}
	updateFPS() {
		let e;
		try {
			e = performance.now()
		} catch (t) {
			e = Date.now()
		}
		this.fpsBuffer = this.fpsBuffer.filter((t => e - t <= 1e3)), this.fpsBuffer.push(e), this.setDebugData({
			FPS: this.fpsBuffer.length
		})
	}
	async flushInternal(e) {
		const t = new s;
		this.needSeek ? (e ? await this.videoReader.seek(this.currentFrame / this.frameRate()) : this
			.videoReader.seek(this.currentFrame / this.frameRate()), this.needSeek = !1) : this
				.currentFrame = Math.floor(this.videoReader.currentTime() * this.frameRate()), this.draw(), t.mark(
					"draw"), this.setDebugData({
						draw: t.measure("", "draw")
					}), this.updateFPS(), this.eventManager.emit(q.onAnimationUpdate)
	}
};
te = ((e, t, i, r) => {
	for (var s, a = r > 1 ? void 0 : r ? ee(t, i) : t, n = e.length - 1; n >= 0; n--)(s = e[n]) && (a = (r ? s(
		t, i, a) : s(a)) || a);
	return r && a && Q(t, i, a), a
})([a], te);
var ie = Object.defineProperty,
	re = Object.getOwnPropertyDescriptor,
	se = Object.getOwnPropertySymbols,
	ae = Object.prototype.hasOwnProperty,
	ne = Object.prototype.propertyIsEnumerable,
	oe = (e, t, i) => t in e ? ie(e, t, {
		enumerable: !0,
		configurable: !0,
		writable: !0,
		value: i
	}) : e[t] = i;
let he = class extends te {
	constructor(e, t, i) {
		var r;
		super(e, t, i), this.scale = {
			x: 1,
			y: 1
		}, this.positionLocation = 0, this.texcoordLocation = 0, this.alphaStartLocation = null, this
			.scaleLocation = null, this.resolutionLocation = null, this.positionBuffer = null, this
				.texcoordBuffer = null, this.originalVideoTexture = null, this.renderingTexture = null, this
					.renderingFbo = null;
		const s = null == (r = this.canvas) ? void 0 : r.getContext("webgl", ((e, t) => {
			for (var i in t || (t = {})) ae.call(t, i) && oe(e, i, t[i]);
			if (se)
				for (var i of se(t)) ne.call(t, i) && oe(e, i, t[i]);
			return e
		})({}, n));
		if (!s) throw new Error("Can't get WebGL context!");
		this.gl = s, this.videoParam.hasAlpha ? this.program = x(this.gl, P(z), P(
			"\n      precision mediump float;\n      // our texture\n      uniform sampler2D u_image;\n      \n      // the texCoords passed in from the vertex shader.\n      varying vec2 v_texCoord;\n      uniform vec2 v_alphaStart;\n      \n      void main() {\n         vec4 color = texture2D(u_image, v_texCoord);\n         vec4 alpha = texture2D(u_image, vec2(v_texCoord.x + v_alphaStart.x, v_texCoord.y + v_alphaStart.y));\n         gl_FragColor = vec4(color.rgb * alpha.r, alpha.r);\n      }     \n         "
		)) : this.program = x(this.gl, P(z), P(
			"\n      precision mediump float;\n      // our texture\n      uniform sampler2D u_image;\n      \n      // the texCoords passed in from the vertex shader.\n      varying vec2 v_texCoord;\n      \n      void main() {\n         gl_FragColor = texture2D(u_image, v_texCoord);\n      }\n         "
		)), this.loadContext()
	}
	loadContext() {
		if (!this.program) throw new Error("program is not initialized");
		if (this.positionLocation = this.gl.getAttribLocation(this.program, "a_position"), -1 === this
			.positionLocation) throw new Error("unable to get attribute location for a_position");
		if (this.scaleLocation = this.gl.getUniformLocation(this.program, "u_scale"), -1 === this.scaleLocation)
			throw new Error("unable to get attribute location for u_scale");
		if (this.texcoordLocation = this.gl.getAttribLocation(this.program, "a_texCoord"), -1 === this
			.texcoordLocation) throw new Error("unable to get attribute location for a_texCoord");
		if (this.videoParam.hasAlpha && (this.alphaStartLocation = this.gl.getUniformLocation(this.program,
			"v_alphaStart"), !this.alphaStartLocation)) throw new Error(
				"unable to get attribute location for v_alphaStart");
		if (this.resolutionLocation = this.gl.getUniformLocation(this.program, "u_resolution"), -1 === this
			.positionLocation) throw new Error("unable to get attribute location for u_resolution");
		if (this.positionBuffer = this.gl.createBuffer(), this.gl.bindBuffer(this.gl.ARRAY_BUFFER, this
			.positionBuffer), this.setRectangle(this.gl, 0, 0, this.videoParam.MP4Width, this.videoParam
				.MP4Height), this.texcoordBuffer = this.gl.createBuffer(), this.gl.bindBuffer(this.gl
					.ARRAY_BUFFER, this.texcoordBuffer), this.gl.bufferData(this.gl.ARRAY_BUFFER, new Float32Array([
						0, 0, 1, 0, 0, 1, 0, 1, 1, 0, 1, 1
					]), this.gl.STATIC_DRAW), this.renderingTexture = A(this.gl), this.gl.texImage2D(this.gl.TEXTURE_2D,
						0, this.gl.RGBA, this.videoParam.sequenceWidth, this.videoParam.sequenceHeight, 0, this.gl.RGBA,
						this.gl.UNSIGNED_BYTE, null), this.renderingFbo = this.gl.createFramebuffer(), !this
							.renderingFbo) throw new Error("unable to create framebuffer");
		this.gl.bindFramebuffer(this.gl.FRAMEBUFFER, this.renderingFbo), this.gl.framebufferTexture2D(this.gl
			.FRAMEBUFFER, this.gl.COLOR_ATTACHMENT0, this.gl.TEXTURE_2D, this.renderingTexture, 0), this
				.originalVideoTexture = A(this.gl)
	}
	draw() {
		this.gl.bindTexture(this.gl.TEXTURE_2D, this.originalVideoTexture), this.texImage2D(), this.gl
			.clearColor(0, 0, 0, 0), this.gl.clear(this.gl.COLOR_BUFFER_BIT), this.gl.useProgram(this.program),
			this.gl.enableVertexAttribArray(this.positionLocation), this.gl.bindBuffer(this.gl.ARRAY_BUFFER,
				this.positionBuffer);
		const e = this.gl.FLOAT;
		this.gl.vertexAttribPointer(this.positionLocation, 2, e, false, 0, 0), this.gl.enableVertexAttribArray(
			this.texcoordLocation), this.gl.bindBuffer(this.gl.ARRAY_BUFFER, this.texcoordBuffer), this.gl
				.vertexAttribPointer(this.texcoordLocation, 2, e, false, 0, 0), this.videoParam.hasAlpha && this.gl
					.uniform2f(this.alphaStartLocation, this.videoParam.alphaStartX / this.videoParam.MP4Width / this
						.scale.x, this.videoParam.alphaStartY / this.videoParam.MP4Height / this.scale.y), this.gl
							.bindTexture(this.gl.TEXTURE_2D, this.originalVideoTexture), this.gl.bindFramebuffer(this.gl
								.FRAMEBUFFER, this.renderingFbo), this.gl.uniform2f(this.resolutionLocation, this.videoParam
									.sequenceWidth, this.videoParam.sequenceHeight), this.gl.uniform2f(this.scaleLocation, this
										.scale.x, this.scale.y), this.gl.viewport(0, 0, this.videoParam.sequenceWidth, this.videoParam
											.sequenceHeight);
		const t = this.gl.TRIANGLES;
		this.gl.drawArrays(t, 0, 6), this.gl.bindFramebuffer(this.gl.FRAMEBUFFER, null), this.gl.uniform2f(this
			.resolutionLocation, this.videoParam.sequenceWidth, this.videoParam.sequenceHeight), this.gl
				.viewport(this.viewportSize.x, this.viewportSize.y, this.viewportSize.width, this.viewportSize
					.height), this.gl.drawArrays(t, 0, 6)
	}
	clearRender() {
		this.gl.clearColor(0, 0, 0, 0), this.gl.clear(this.gl.COLOR_BUFFER_BIT)
	}
	detectWebGLContext() {
		return (() => {
			const e = document.createElement("canvas");
			return !(!e.getContext("webgl") && !e.getContext("experimental-webgl"))
		})()
	}
	texImage2D() {
		this.gl.texImage2D(this.gl.TEXTURE_2D, 0, this.gl.RGBA, this.gl.RGBA, this.gl.UNSIGNED_BYTE, this
			.videoReader.getVideoElement())
	}
	setRectangle(e, t, i, r, s) {
		const a = t,
			n = t + r,
			o = i,
			h = i + s;
		e.bufferData(e.ARRAY_BUFFER, new Float32Array([a, o, n, o, a, h, a, h, n, o, n, h]), e.STATIC_DRAW)
	}
};
he = ((e, t, i, r) => {
	for (var s, a = r > 1 ? void 0 : r ? re(t, i) : t, n = e.length - 1; n >= 0; n--)(s = e[n]) && (a = (r ? s(
		t, i, a) : s(a)) || a);
	return r && a && ie(t, i, a), a
})([a], he);
var le = (e => (e[e.Unknown = 0] = "Unknown", e[e.Vector = 1] = "Vector", e[e.Bitmap = 2] = "Bitmap", e[e.Video = 3] =
	"Video", e))(le || {}),
	de = (e => (e[e.End = 0] = "End", e[e.FontTables = 1] = "FontTables", e[e.VectorCompositionBlock = 2] =
		"VectorCompositionBlock", e[e.CompositionAttributes = 3] = "CompositionAttributes", e[e.ImageTables = 4] =
		"ImageTables", e[e.LayerBlock = 5] = "LayerBlock", e[e.LayerAttributes = 6] = "LayerAttributes", e[e
			.SolidColor = 7] = "SolidColor", e[e.TextSource = 8] = "TextSource", e[e.TextPathOption = 9] =
		"TextPathOption", e[e.TextMoreOption = 10] = "TextMoreOption", e[e.ImageReference = 11] = "ImageReference",
		e[e.CompositionReference = 12] = "CompositionReference", e[e.Transform2D = 13] = "Transform2D", e[e
			.MaskBlock = 14] = "MaskBlock", e[e.ShapeGroup = 15] = "ShapeGroup", e[e.Rectangle = 16] = "Rectangle",
		e[e.Ellipse = 17] = "Ellipse", e[e.PolyStar = 18] = "PolyStar", e[e.ShapePath = 19] = "ShapePath", e[e
			.Fill = 20] = "Fill", e[e.Stroke = 21] = "Stroke", e[e.GradientFill = 22] = "GradientFill", e[e
				.GradientStroke = 23] = "GradientStroke", e[e.MergePaths = 24] = "MergePaths", e[e.TrimPaths = 25] =
		"TrimPaths", e[e.Repeater = 26] = "Repeater", e[e.RoundCorners = 27] = "RoundCorners", e[e.Performance =
		28] = "Performance", e[e.DropShadowStyle = 29] = "DropShadowStyle", e[e.InnerShadowStyle = 30] =
		"InnerShadowStyle", e[e.OuterGlowStyle = 31] = "OuterGlowStyle", e[e.InnerGlowStyle = 32] =
		"InnerGlowStyle", e[e.BevelAndEmbossStyle = 33] = "BevelAndEmbossStyle", e[e.SatinStyle = 34] =
		"SatinStyle", e[e.ColorOverlayStyle = 35] = "ColorOverlayStyle", e[e.GradientOverlayStyle = 36] =
		"GradientOverlayStyle", e[e.StrokeStyle = 37] = "StrokeStyle", e[e.TintEffect = 38] = "TintEffect", e[e
			.FillEffect = 39] = "FillEffect", e[e.StrokeEffect = 40] = "StrokeEffect", e[e.TritoneEffect = 41] =
		"TritoneEffect", e[e.DropShadowEffect = 42] = "DropShadowEffect", e[e.RadialWipeEffect = 43] =
		"RadialWipeEffect", e[e.DisplacementMapEffect = 44] = "DisplacementMapEffect", e[e.BitmapCompositionBlock =
		45] = "BitmapCompositionBlock", e[e.BitmapSequence = 46] = "BitmapSequence", e[e.ImageBytes = 47] =
		"ImageBytes", e[e.ImageBytes2 = 48] = "ImageBytes2", e[e.ImageBytes3 = 49] = "ImageBytes3", e[e
			.VideoCompositionBlock = 50] = "VideoCompositionBlock", e[e.VideoSequence = 51] = "VideoSequence", e[e
				.LayerAttributesV2 = 52] = "LayerAttributesV2", e[e.Count = 53] = "Count", e))(de || {});
const ce = e => {
	const t = e.readUint16();
	let i = (63 & t) >>> 0;
	const r = t >> 6;
	return 63 === i && (i = e.readUint32()), e.context.tagLevel < r && (e.context.tagLevel = r), {
		code: r,
		length: i
	}
};

function ue(e, t, i) {
	let r = ce(e);
	for (; r.code !== de.End;) {
		const s = e.readBytes(r.length);
		i(s, r.code, t), e.context.tagLevel < s.context.tagLevel && (e.context.tagLevel = s.context.tagLevel), r = ce(e)
	}
}
const me = {
	red: 0,
	green: 0,
	blue: 0
},
	fe = {
		red: 255,
		green: 255,
		blue: 255
	},
	pe = () => {
		console.error("PAG Verify Failed!")
	},
	ge = e => !!e || (console.error("PAG Verify Failed!"), !1),
	ye = class {
		constructor() {
			this.id = 0, this.width = 0, this.height = 0, this.duration = 0, this.frameRate = 30, this
				.backgroundColor = fe, this.cacheID = 0, this.cacheID = ye.cacheIDCount, ye.cacheIDCount += 1
		}
		type() {
			return le.Unknown
		}
		getStaticTimeRanges() {
			return []
		}
		verify() {
			return ge(this.width > 0 && this.height > 0 && this.duration > 0 && this.frameRate > 0)
		}
	};
let we = ye;
we.cacheIDCount = 1;
class ve extends we {
	constructor() {
		super(...arguments), this.hasAlpha = !1, this.sequences = [], this.staticTimeRanges = [], this
			.staticTimeRangeUpdated = !1
	}
	type() {
		return le.Video
	}
	getStaticTimeRanges() {
		return this.staticTimeRangeUpdated || (this.staticTimeRangeUpdated = !0, this.updateStaticTimeRanges()),
			this.staticTimeRanges
	}
	updateStaticTimeRanges() {
		if (!(this.duration <= 1))
			if (this.sequences.length > 0) {
				let e = this.sequences[0];
				for (let t = 1; t < this.sequences.length; t++) {
					const i = this.sequences[t];
					i.frameRate > e.frameRate && (e = i)
				}
				const t = this.frameRate / e.frameRate;
				for (const i of e.staticTimeRanges) i.start = Math.round(i.start * t), i.end = Math.round(i.end *
					t), this.staticTimeRanges.push(i)
			} else {
				const e = {
					start: 0,
					end: this.duration - 1
				};
				this.staticTimeRanges.push(e)
			}
	}
	hasImageContent() {
		return !0
	}
	verify() {
		if (!super.verify() || this.sequences.length <= 0) return pe(), !1;
		for (const e of this.sequences)
			if (!e || !e.verify()) return pe(), !1;
		return !0
	}
}
class be {
	constructor(e, t) {
		this.numerator = 1, this.denominator = 1, this.numerator = e, this.denominator = t
	}
	value() {
		return this.numerator / this.denominator
	}
}
const Se = new be(1, 1);
class Ee {
	constructor(e, t) {
		this.x = e, this.y = t
	}
}
const Te = new Ee(0, 0);
var xe = (e => (e[e.None = 0] = "None", e[e.Alpha = 1] = "Alpha", e[e.AlphaInverted = 2] = "AlphaInverted", e[e.Luma =
	3] = "Luma", e[e.LumaInverted = 4] = "LumaInverted", e))(xe || {}),
	Re = (e => (e[e.Unknown = 0] = "Unknown", e[e[void 0] = 1] = "undefined", e[e.Solid = 2] = "Solid", e[e.Text = 3] =
		"Text", e[e.Shape = 4] = "Shape", e[e.Image = 5] = "Image", e[e.PreCompose = 6] = "PreCompose", e))(Re || {});
class Pe {
	constructor() {
		this.id = 0, this.parent = void 0, this.containingComposition = void 0, this.stretch = Se, this.startTime =
			0, this.duration = 0, this.autoOrientation = !1, this.transform = void 0, this.isActive = !0, this
				.blendMode = o.Normal, this.trackMatteType = 0, this.trackMatteLayer = void 0, this.timeRemap = void 0,
			this.masks = void 0, this.effects = void 0, this.layerStyles = void 0, this.layerCache = void 0, this
				.maxScale = void 0
	}
	type() {
		return 0
	}
	excludeVaryingRanges(e) {
		var t;
		if (null == (t = this.transform) || t.excludeVaryingRanges(e), void 0 !== this.timeRemap && this.timeRemap
			.excludeVaryingRanges(e), void 0 !== this.masks)
			for (const t of this.masks) t.excludeVaryingRanges(e);
		if (void 0 !== this.effects && this.effects.length > 0)
			for (const t of this.effects) t.excludeVaryingRanges(e);
		if (void 0 !== this.layerStyles && this.layerStyles.length > 0)
			for (const t of this.layerStyles) t.excludeVaryingRanges(e)
	}
	gotoFrame(e) {
		var t;
		if (null == (t = this.transform) || t.gotoFrame(e), void 0 !== this.timeRemap && this.timeRemap.gotoFrame(
			e), void 0 !== this.masks && this.masks.length > 0)
			for (const t of this.masks) t.gotoFrame(e);
		if (void 0 !== this.effects && this.effects.length > 0)
			for (const t of this.effects) t.gotoFrame(e);
		if (void 0 !== this.layerStyles && this.layerStyles.length > 0)
			for (const t of this.layerStyles) t.gotoFrame(e)
	}
	verify() {
		if (!this.containingComposition || this.duration <= 0 || !this.transform) return pe(), !1;
		if (!this.transform.verify()) return pe(), !1;
		if (this.masks && this.masks.length > 0)
			for (const e of this.masks)
				if (!e || !e.verify()) return pe(), !1;
		if (this.layerStyles && this.layerStyles.length > 0)
			for (const e of this.layerStyles)
				if (!e || !e.verify()) return pe(), !1;
		if (this.effects && this.effects.length > 0)
			for (const e of this.effects)
				if (!e || !e.verify()) return pe(), !1;
		return !0
	}
	getMaxScaleFactor() {
		if (void 0 !== this.maxScale) return this.maxScale;
		this.maxScale = new Ee(1, 1);
		const e = this.transform.scale;
		if (e.animatable()) {
			const {
				keyframes: t
			} = e;
			let i = Math.abs(t[0].startValue.x),
				r = Math.abs(t[0].startValue.y);
			if (void 0 !== t && t.length > 0)
				for (const e of t) {
					const t = Math.abs(e.endValue.x),
						s = Math.abs(e.endValue.y);
					i < t && (i = t), r < s && (r = s)
				}
			this.maxScale.x = i, this.maxScale.y = r
		} else this.maxScale.x = Math.abs(e.value.x), this.maxScale.y = Math.abs(e.value.y);
		if (void 0 !== this.parent) {
			const e = this.parent.getMaxScaleFactor();
			this.maxScale.x *= e.x, this.maxScale.y *= e.y
		}
		return this.maxScale
	}
}
const Ae = e => ({
	red: e.readUint8(),
	green: e.readUint8(),
	blue: e.readUint8()
}),
	De = e => e.readEncodedUint64(),
	ke = e => {
		const t = e.readFloat32(),
			i = e.readFloat32();
		return new Ee(t, i)
	},
	Fe = (e, t) => {
		t.width = e.readEncodeInt32(), t.height = e.readEncodeInt32(), t.duration = De(e), t.frameRate = e
			.readFloat32(), t.backgroundColor = Ae(e)
	};
class Ve {
	constructor() {
		this.tagLevel = 0, this.compositions = [], this.errorMessages = []
	}
	throwException(e) {
		this.errorMessages.push(e)
	}
	releaseCompositions() {
		const e = this.compositions.slice();
		return this.compositions = [], e
	}
}
const Ce = "PAG file decode error!";
class Le {
	constructor(e, t) {
		this._position = 0, this.bitPosition = 0, this.dataView = new DataView(e), this.littleEndian = !!t, this
			.context = new Ve
	}
	get length() {
		return this.dataView.byteLength
	}
	get bytesAvailable() {
		return this.dataView.byteLength - this._position
	}
	data() {
		return this.dataView.buffer
	}
	get position() {
		return this._position
	}
	alignWithBytes() {
		this.bitPosition = 8 * this._position
	}
	readBoolean() {
		const e = this.dataView.getInt8(this._position);
		return this._position += 1, this.positonChanged(), Boolean(e)
	}
	readChar() {
		if (this._position >= this.length) throw new Error(Ce);
		const e = this.dataView.getInt8(this._position);
		return this._position += 1, this.positonChanged(), String.fromCharCode(e)
	}
	readUint8() {
		if (this._position >= this.length) throw new Error(Ce);
		const e = this.dataView.getUint8(this._position);
		return this._position += 1, this.positonChanged(), e
	}
	readInt8() {
		if (this._position >= this.length) throw new Error(Ce);
		const e = this.dataView.getInt8(this._position);
		return this._position += 1, this.positonChanged(), e
	}
	readInt16() {
		if (this._position >= this.length - 1) throw new Error(Ce);
		const e = this.dataView.getInt16(this._position, this.littleEndian);
		return this._position += 2, this.positonChanged(), e
	}
	readUint16() {
		if (this._position >= this.length - 1) throw new Error(Ce);
		const e = this.dataView.getUint16(this._position, this.littleEndian);
		return this._position += 2, this.positonChanged(), e
	}
	readInt24() {
		if (this._position >= this.length - 2) throw new Error(Ce);
		const e = this.dataView.getInt16(this._position, this.littleEndian),
			t = this.dataView.getInt8(this._position + 2);
		return this._position += 3, this.positonChanged(), this.littleEndian ? e + 65536 * t : 65536 * e + t
	}
	readUint24() {
		if (this._position >= this.length - 2) throw new Error(Ce);
		const e = this.dataView.getUint16(this._position, this.littleEndian),
			t = this.dataView.getUint8(this._position + 2);
		return this._position += 3, this.positonChanged(), this.littleEndian ? e + 65536 * t : 65536 * e + t
	}
	readInt32() {
		if (this._position >= this.length - 3) throw new Error(Ce);
		const e = this.dataView.getInt32(this._position, this.littleEndian);
		return this._position += 4, this.positonChanged(), e
	}
	readUint32() {
		if (this._position >= this.length - 3) throw new Error(Ce);
		const e = this.dataView.getUint32(this._position, this.littleEndian);
		return this._position += 4, this.positonChanged(), e
	}
	readInt64() {
		if (this._position >= this.length - 7) throw new Error(Ce);
		const e = this.dataView.getInt32(this._position, this.littleEndian),
			t = this.dataView.getInt32(this._position + 4, this.littleEndian);
		return this._position += 8, this.positonChanged(), this.littleEndian ? e + 2 ** 32 * t : 2 ** 32 * e + t
	}
	readUint64() {
		if (this._position >= this.length - 7) throw new Error(Ce);
		const e = this.dataView.getUint32(this._position, this.littleEndian),
			t = this.dataView.getUint32(this._position + 4, this.littleEndian);
		return this._position += 8, this.positonChanged(), this.littleEndian ? e + 2 ** 32 * t : 2 ** 32 * e + t
	}
	readFloat32() {
		if (this._position >= this.length - 3) throw new Error(Ce);
		const e = this.dataView.getFloat32(this._position, this.littleEndian);
		return this._position += 4, this.positonChanged(), e
	}
	readDouble() {
		if (this._position >= this.length - 7) throw new Error(Ce);
		const e = this.dataView.getFloat64(this._position, this.littleEndian);
		return this._position += 8, this.positonChanged(), e
	}
	readUTF8String() {
		if (this._position >= this.length) throw new Error(Ce);
		let e = "",
			t = 0;
		for (let i = this._position; i < this.length && 0 !== this.dataView.getUint8(i); i++) e +=
			`%${this.dataView.getUint8(i).toString(16)}`, t += 1;
		return this._position += t, this.positonChanged(), decodeURIComponent(e)
	}
	readEncodedUint32() {
		let e = 0,
			t = 0;
		for (let i = 0; i < 32; i += 7) {
			if (this._position >= this.length) throw Error("readEncodedUint32 End of file was encountered.");
			if (t = this.dataView.getUint8(this._position), this._position += 1, e |= (127 & t) << i, 0 == (128 &
				t)) break
		}
		return this.positonChanged(), e
	}
	readEncodeInt32() {
		const e = this.readEncodedUint32(),
			t = e >> 1;
		return (1 & e) > 0 ? -t : t
	}
	readEncodedUint64() {
		let e = 0,
			t = 0;
		for (let i = 0; i < 64; i += 7) {
			if (this._position >= this.length) throw Error("readEncodedUint64 End of file was encountered.");
			if (t = this.dataView.getUint8(this._position), this._position += 1, e |= (127 & t) << i, 0 == (128 &
				t)) break
		}
		return this.positonChanged(), e
	}
	readEncodeInt64() {
		const e = this.readEncodedUint64(),
			t = e << 0;
		return (1 & e) > 0 ? -t : t
	}
	readBytes(e) {
		const t = e || this.length - this._position;
		if (this._position > this.length - t) throw new Error(Ce);
		const i = this.dataView.buffer.slice(this._position, this._position + t);
		return this._position += t, this.positonChanged(), new Le(i, this.littleEndian)
	}
	readUBits(e) {
		const t = [0, 1, 3, 7, 15, 31, 63, 127, 255];
		let i = 0;
		if (this.bitPosition > 8 * this.length - e) throw new Error(Ce);
		let r = 0;
		for (; r < e;) {
			const s = Math.floor(.125 * this.bitPosition),
				a = this.bitPosition % 8;
			let n = this.dataView.getUint8(s) >> a;
			const o = Math.min(8 - a, e - r);
			n &= t[o], i |= n << r, r += o, this.bitPosition += o
		}
		return this.bitPositionChanged(), i
	}
	readBits(e) {
		let t = this.readUBits(e);
		t <<= 32 - e;
		return t << 0 >> 32 - e
	}
	readNumBits() {
		return this.readUBits(5) + 1
	}
	readInt32List(e) {
		const t = this.readNumBits(),
			i = new Array(e);
		for (let r = 0; r < e; r++) i[r] = this.readBits(t);
		return i
	}
	readUint32List(e) {
		const t = this.readNumBits(),
			i = new Array(e);
		for (let r = 0; r < e; r++) i[r] = this.readUBits(t);
		return i
	}
	readBitBoolean() {
		return 0 !== this.readUBits(1)
	}
	readFloatList(e, t) {
		const i = this.readNumBits(),
			r = new Array(e);
		for (let s = 0; s < e; s++) r[s] = this.readBits(i) * t;
		return r
	}
	bitPositionChanged() {
		this._position = Math.ceil(.125 * this.bitPosition)
	}
	positonChanged() {
		this.bitPosition = 8 * this._position
	}
}
class Be {
	constructor(e, t) {
		this.length = 0, this.data = e, this.length = t
	}
}
class _e {
	constructor() {
		this.isKeyframe = !1, this.frame = 0, this.fileBytes = new Be(new Le(new ArrayBuffer(0)), 0)
	}
}
class Ue extends class {
	constructor() {
		this.composition = void 0, this.id = 0, this.width = 0, this.height = 0, this.frameRate = 0, this
			.frameCount = 0, this.isKeyFrameFlags = []
	}
	verify() {
		return ge(void 0 !== this.composition && this.width > 0 && this.height > 0 && this.frameRate > 0)
	}
} {
	constructor() {
		super(...arguments), this.alphaStartX = 0, this.alphaStartY = 0, this.frames = [], this.headers = [], this
			.staticTimeRanges = []
	}
	verify() {
		if (!super.verify() || this.frames.length <= 0) return pe(), !1;
		for (const e of this.frames)
			if (!e || !e.fileBytes) return pe(), !1;
		for (const e of this.headers)
			if (!e) return pe(), !1;
		return !0
	}
	getVideoWidth() {
		let e = this.alphaStartX + this.width;
		return e % 2 == 1 && (e += 1), e
	}
	getVideoHeight() {
		let e = this.alphaStartY + this.height;
		return e % 2 == 1 && (e += 1), e
	}
}
const Me = e => {
	const t = e.readEncodedUint32(),
		i = e.readBytes(t);
	if (0 === t) throw new Error("Read start code with length 0!");
	const r = new ArrayBuffer(t + 4);
	return new DataView(r).setUint32(0, t), ((e, t, i, r, s) => {
		if (t >= e.byteLength || r >= i.byteLength || i.byteLength - r > e.byteLength - t || s > i
			.byteLength) return;
		const a = new Uint8Array(e),
			n = new Uint8Array(i, r, s);
		a.set(n, t)
	})(r, 4, i.data(), 0, t), new Be(new Le(r), t + 4)
},
	Ie = (e, t, i) => {
		const {
			composition: r
		} = i;
		switch (t) {
			case de.CompositionAttributes:
				Fe(e, r);
				break;
			case de.VideoSequence: {
				const t = ((e, t) => {
					const i = new Ue;
					i.width = e.readEncodeInt32(), i.height = e.readEncodeInt32(), i.frameRate = e
						.readFloat32(), t && (i.alphaStartX = e.readEncodeInt32(), i.alphaStartY = e
							.readEncodeInt32());
					const r = Me(e),
						s = Me(e);
					i.headers.push(r, s), i.frameCount = e.readEncodedUint32();
					for (let t = 0; t < i.frameCount; t++) {
						const t = new _e;
						t.isKeyframe = e.readBitBoolean(), i.frames.push(t)
					}
					for (let t = 0; t < i.frameCount; t++) {
						const r = i.frames[t];
						r.frame = De(e), r.fileBytes = Me(e)
					}
					if (e.bytesAvailable > 0) {
						const t = e.readEncodedUint32();
						for (let r = 0; r < t; r++) {
							const t = {
								start: 0,
								end: 0
							};
							t.start = De(e), t.end = De(e), i.staticTimeRanges.push(t)
						}
					}
					return i
				})(e, i.hasAlpha);
				t.composition = r, r.sequences.push(t);
				break
			}
		}
	};
var Oe = (e => (e[e.Unknown = 0] = "Unknown", e[e.Tint = 1] = "Tint", e[e.Fill = 2] = "Fill", e[e.Stroke = 3] =
	"Stroke", e[e.Tritone = 4] = "Tritone", e[e.DropShadow = 5] = "DropShadow", e[e.RadialWipe = 6] =
	"RadialWipe", e[e.DisplacementMap = 7] = "DisplacementMap", e))(Oe || {});

function qe(e, t, i) {
	if (i < t) return;
	for (let r = e.length - 1; r >= 0; r--) {
		const s = e[r];
		if (!(s.end < t || s.start > i)) {
			if (s.start < t && s.end > i) {
				const a = {
					start: i + 1,
					end: s.end
				};
				s.end = t - 1, a.end > a.start && e.splice(r + 1, 0, a), s.end <= s.start && e.splice(r, 1);
				break
			}
			s.start >= t && s.end <= i ? e.splice(r, 1) : s.start < t ? (s.end = t - 1, s.end <= s.start && e.splice(r,
				1)) : (s.start = i + 1, s.end <= s.start && e.splice(r, 1))
		}
	}
}

function Ge(e, t) {
	for (let i = e.length - 1; i >= 0; i--) {
		const r = e[i];
		if (r.start === t || r.end <= t) break;
		if (r.start < t && r.end > t) {
			const s = {
				start: t,
				end: r.end
			};
			r.end = t - 1, s.end > s.start && e.splice(i + 1, 0, s), r.end <= r.start && e.splice(i, 1);
			break
		}
	}
}
class Ne extends we {
	constructor() {
		super(...arguments), this.layers = [], this.staticTimeRanges = [], this.staticTimeRangeUpdated = !1
	}
	type() {
		return le.Vector
	}
	getStaticTimeRanges() {
		return this.staticTimeRangeUpdated || (this.staticTimeRangeUpdated = !0, this.updateStaticTimeRanges()),
			this.staticTimeRanges
	}
	verify() {
		if (!super.verify()) return pe(), !1;
		for (const e of this.layers)
			if (!e || !e.verify()) return pe(), !1;
		return !0
	}
	updateStaticTimeRanges() {
		if (this.duration > 1) {
			const e = {
				start: 0,
				end: this.duration - 1
			};
			this.staticTimeRanges = [e];
			for (const e of this.layers) {
				if (this.staticTimeRanges.length <= 0) break;
				e.excludeVaryingRanges(this.staticTimeRanges), Ge(this.staticTimeRanges, e.startTime), Ge(this
					.staticTimeRanges, e.startTime + e.duration)
			}
		}
	}
}
class ze {
	constructor(e) {
		this.value = e
	}
	animatable() {
		return !1
	}
	excludeVaryingRanges(e) { }
	gotoFrame(e) { }
}
class We {
	static createDefaultTransform2D() {
		return new We
	}
	constructor() {
		this.anchorPoint = new ze(Te), this.position = new ze(Te), this.xPosition = new ze(0), this.yPosition =
			new ze(0), this.scale = new ze(new Ee(1, 1)), this.rotation = new ze(0), this.opacity = new ze(255)
	}
	excludeVaryingRanges(e) {
		this.anchorPoint.excludeVaryingRanges(e), void 0 !== this.position ? this.position.excludeVaryingRanges(e) :
			(this.xPosition.excludeVaryingRanges(e), this.yPosition.excludeVaryingRanges(e)), this.scale
				.excludeVaryingRanges(e), this.rotation.excludeVaryingRanges(e), this.opacity.excludeVaryingRanges(e)
	}
	gotoFrame(e) {
		this.anchorPoint.gotoFrame(e), void 0 !== this.position ? this.position.gotoFrame(e) : (this.xPosition
			.gotoFrame(e), this.yPosition.gotoFrame(e)), this.scale.gotoFrame(e), this.rotation.gotoFrame(e),
			this.opacity.gotoFrame(e)
	}
	verify() {
		return void 0 !== this.anchorPoint && (void 0 !== this.position || void 0 !== this.xPosition && void 0 !==
			this.yPosition) && void 0 !== this.scale && void 0 !== this.rotation && void 0 !== this.opacity
	}
}
class Xe extends Pe {
	constructor() {
		super(...arguments), this.composition = void 0, this.compositionStartTime = 0, this.staticTimeRanges =
			void 0, this.staticTimeRangeUpdated = !1
	}
	static wrap(e) {
		const t = new Xe;
		t.duration = e.duration;
		const i = new We;
		return t.transform = i, t.composition = e, t
	}
	type() {
		return Re.PreCompose
	}
	excludeVaryingRanges(e) {
		super.excludeVaryingRanges(e), e && 0 !== e.length && this.updateStaticTimeRanges()
	}
	gotoFrame(e) {
		super.gotoFrame(e)
	}
	verify() {
		return !!super.verify() && !!this.composition
	}
	updateStaticTimeRanges() {
		var e;
		if (this.staticTimeRangeUpdated) return;
		this.staticTimeRangeUpdated = !0;
		const t = null == (e = this.composition) ? void 0 : e.getStaticTimeRanges();
		if (t) {
			for (let e = t.length - 1; e >= 0; e--) {
				const i = t[e];
				i.start += this.compositionStartTime, i.end += this.compositionStartTime, i.end <= this.startTime ?
					t.pop() : i.start < this.startTime ? i.start = 0 : i.start >= this.startTime + this.duration -
						1 ? t.pop() : i.end > this.startTime + this.duration - 1 && (i.end = this.startTime + this
							.duration - 1)
			}
			this.staticTimeRanges = t
		}
	}
}
class He extends Pe {
	constructor() {
		super(...arguments), this.contents = []
	}
	type() {
		return Re.Shape
	}
	excludeVaryingRanges(e) {
		super.excludeVaryingRanges(e);
		for (const t of this.contents) t.excludeVaryingRanges(e)
	}
	gotoFrame(e) {
		super.gotoFrame(e);
		for (const t of this.contents) t.gotoFrame(e)
	}
	verify() {
		if (!super.verify()) return !1;
		for (const e of this.contents)
			if (void 0 === e || !e.verify()) return !1;
		return !0
	}
}
class je extends Pe {
	constructor() {
		super(...arguments), this.solidColor = me, this.width = 0, this.height = 0
	}
	type() {
		return Re.Solid
	}
	excludeVaryingRanges(e) {
		super.excludeVaryingRanges(e)
	}
	gotoFrame(e) {
		super.gotoFrame(e)
	}
	verify() {
		return super.verify() ? ge(this.width > 0 && this.height > 0) : (pe(), !1)
	}
}
class Ke extends Pe {
	type() {
		return Re.undefined
	}
}
class Ye {
	constructor() {
		this.startTime = 0, this.endTime = 0, this.interpolationType = h.Hold, this.bezierOut = [], this
			.bezierIn = [], this.spatialOut = Te, this.spatialIn = Te
	}
	initialize() { }
	getValue(e) {
		return this.startValue
	}
	containsTime(e) {
		return e >= this.startTime && e < this.endTime
	}
}
class $e extends ze {
	constructor(e) {
		if (!e || 0 === e.length) throw new Error("keyframes is required");
		if (void 0 === e[0].startValue) throw new Error("startValue is required");
		super(e[0].startValue), this.keyframes = e, this.lastKeyframeIndex = 0;
		for (const t of e) t.initialize()
	}
	animatable() {
		return !0
	}
	excludeVaryingRanges(e) {
		for (const t of this.keyframes) switch (t.interpolationType) {
			case h.Bezier:
			case h.Linear:
				qe(e, t.startTime, t.endTime - 1);
				break;
			default:
				Ge(e, t.startTime), Ge(e, t.endTime)
		}
	}
	gotoFrame(e) {
		let t = this.keyframes[this.lastKeyframeIndex];
		if (t.containsTime(e)) this.value = t.getValue(e);
		else {
			if (e < t.startTime)
				for (; this.lastKeyframeIndex > 0 && (this.lastKeyframeIndex -= 1, !this.keyframes[this
					.lastKeyframeIndex].containsTime(e)););
			else
				for (; this.lastKeyframeIndex < this.keyframes.length - 1 && (this.lastKeyframeIndex += 1, !this
					.keyframes[this.lastKeyframeIndex].containsTime(e)););
			t = this.keyframes[this.lastKeyframeIndex], void 0 !== t.startValue && e <= t.startTime ? this.value = t
				.startValue : void 0 !== t.endValue && e >= t.endTime ? this.value = t.endValue : this.value = t
					.getValue(e)
		}
	}
}
var Ze = (e => (e[e.Value = 0] = "Value", e[e.FixedValue = 1] = "FixedValue", e[e.SimpleProperty = 2] =
	"SimpleProperty", e[e.DiscreteProperty = 3] = "DiscreteProperty", e[e.MultiDimensionProperty = 4] =
	"MultiDimensionProperty", e[e.SpatialProperty = 5] = "SpatialProperty", e[e.BitFlag = 6] = "BitFlag", e[e
		.Custom = 7] = "Custom", e))(Ze || {});
const Je = (e, t, i) => {
	const r = i,
		s = [];
	if (!r.configs || 0 === r.configs.length) return t;
	for (const t of r.configs) {
		const i = rt(e, t);
		s.push(i)
	}
	e.alignWithBytes();
	let a = 0;
	for (const i of r.configs) {
		const r = s[a],
			n = i.key;
		i.readAttribute(e, r, t, n), a += 1
	}
	return t
};
class Qe {
	constructor(e, t, i) {
		this.attributeType = t, this.defaultValue = i, this.key = e
	}
	readAttribute(e, t, i, r) { }
	readValue(e) { }
	readValueList(e, t, i) { }
	dimensionality() {
		return 1
	}
	newKeyframe(e) {
		return new Ye
	}
}
const et = (e, t, i, r, s) => {
	6 === s.attributeType ? i[r] = t.exist : 1 === s.attributeType ? i[r] = s.readValue(e) : 0 === s.attributeType ?
		i[r] = it(e, s, t) : i[r] = tt(e, s, t)
},
	tt = (e, t, i) => {
		let r;
		if (i.exist)
			if (i.animatable) {
				const s = st(e, t, i);
				if (!s || 0 === s.length) throw new Error("Wrong number of keyframes!");
				at(e, s, t), nt(e, s, t), i.hasSpatial, r = new $e(s)
			} else r = new ze(it(e, t, i));
		else r = new ze(t.defaultValue);
		return r
	},
	it = (e, t, i) => i.exist ? t.readValue(e) : t.defaultValue,
	rt = (e, t) => {
		const i = {
			exist: !1,
			animatable: !1,
			hasSpatial: !1
		},
			{
				attributeType: r
			} = t;
		return 1 === r ? (i.exist = !0, i) : (i.exist = e.readBitBoolean(), i.exist && 0 !== r && 6 !== r && 7 !== r ? (
			i.animatable = e.readBitBoolean(), i.animatable && 5 === r ? (i.hasSpatial = e.readBitBoolean(),
				i) : i) : i)
	},
	st = (e, t, i) => {
		const r = [],
			s = e.readEncodedUint32();
		for (let a = 0; a < s; a++) {
			let s;
			if (3 === t.attributeType) s = new Ye;
			else {
				const r = e.readUBits(2);
				r === h.Hold ? s = new Ye : (s = t.newKeyframe(i), s.interpolationType = r)
			}
			r.push(s)
		}
		return r
	},
	at = (e, t, i) => {
		const r = t.length;
		t[0].startTime = De(e);
		for (let i = 0; i < r; i++) {
			const s = De(e);
			t[i].endTime = s, i < r - 1 && (t[i + 1].startTime = s)
		}
		const s = [];
		i.readValueList(e, s, r + 1);
		let a = 0;
		t[0].startValue = s[a], a += 1;
		for (let e = 0; e < r; e++) {
			const i = s[a];
			a += 1, t[e].endValue = i, e < r - 1 && (t[e + 1].startValue = i)
		}
	},
	nt = (e, t, i) => {
		const r = 4 === i.attributeType ? i.dimensionality() : 1,
			s = e.readNumBits();
		for (const i of t) {
			if (i.interpolationType !== h.Bezier) continue;
			let t, a;
			for (let n = 0; n < r; n++) t = .005 * e.readBits(s), a = .005 * e.readBits(s), i.bezierOut.push({
				x: t,
				y: a
			}), t = .005 * e.readBits(s), a = .005 * e.readBits(s), i.bezierIn.push({
				x: t,
				y: a
			})
		}
	};

function ot(e, t, i) {
	return e + (t - e) * i
}
class ht {
	getInterpolation(e) {
		return e
	}
}
class lt extends Ye {
	initialize() {
		super.initialize(), this.interpolationType === h.Bezier || (this.xInterpolator = new ht, this
			.yInterpolator = new ht)
	}
	getValue(e) {
		var t, i, r, s;
		const a = (e - this.startTime) / (this.endTime - this.startTime),
			n = null != (i = null == (t = this.xInterpolator) ? void 0 : t.getInterpolation(a)) ? i : a,
			o = null != (s = null == (r = this.yInterpolator) ? void 0 : r.getInterpolation(a)) ? s : a;
		return {
			x: ot(this.startValue.x, this.endValue.x, n),
			y: ot(this.startValue.y, this.endValue.y, o)
		}
	}
}
class dt extends Ye {
	initialize() {
		this.interpolationType === h.Bezier || (this.interpolator = new ht)
	}
	getProgress(e) {
		var t, i;
		const r = (e - this.startTime) / (this.endTime - this.startTime);
		return null != (i = null == (t = this.interpolator) ? void 0 : t.getInterpolation(r)) ? i : r
	}
	getValue(e) {
		const t = this.getProgress(e);
		return ot(this.startValue, this.endValue, t)
	}
}
class ct extends Qe {
	constructor(e, t, i) {
		super(e, t, i)
	}
	readAttribute(e, t, i, r) {
		et(e, t, i, r, this)
	}
	readValue(e) {
		return e.readFloat32()
	}
	readValueList(e, t, i) {
		for (let r = 0; r < i; r++) t.push(this.readValue(e))
	}
	dimensionality() {
		return 1
	}
	newKeyframe(e) {
		return new dt
	}
}
class ut extends Qe {
	constructor(e, t, i) {
		super(e, t, i)
	}
	readAttribute(e, t, i, r) {
		et(e, t, i, r, this)
	}
	readValue(e) {
		return e.readBoolean()
	}
	readValueList(e, t, i) {
		for (let r = 0; r < i; r++) t.push(e.readBitBoolean())
	}
	dimensionality() {
		return 1
	}
	newKeyframe(e) {
		return new Ye
	}
}
class mt extends Qe {
	constructor(e, t, i) {
		super(e, t, i)
	}
	readAttribute(e, t, i, r) {
		et(e, t, i, r, this)
	}
	readValue(e) {
		return e.readUint8()
	}
	readValueList(e, t, i) {
		const r = e.readUint32List(i);
		for (let e = 0; e < i; e++) t.push(r[e])
	}
	dimensionality() {
		return 1
	}
	newKeyframe(e) {
		return new dt
	}
}
class ft extends Qe {
	constructor(e, t, i) {
		super(e, t, i)
	}
	readAttribute(e, t, i, r) {
		et(e, t, i, r, this)
	}
	readValue(e) {
		return De(e)
	}
	readValueList(e, t, i) {
		for (let r = 0; r < i; r++) t[r] = this.readValue(e)
	}
	dimensionality() {
		return 1
	}
	newKeyframe(e) {
		return new dt
	}
}
class pt extends Qe {
	constructor(e, t, i) {
		super(e, t, i)
	}
	readAttribute(e, t, i, r) {
		et(e, t, i, r, this)
	}
	readValue(e) {
		return ke(e)
	}
	readValueList(e, t, i) {
		if (this.attributeType === Ze.SpatialProperty) {
			const r = e.readFloatList(2 * i, .05);
			for (let e = 0; e < i; e++) t[e] || (t[e] = new Ee(0, 0)), t[e].x = r[e]
		} else
			for (let r = 0; r < i; r++) t[r] = ke(e)
	}
	dimensionality() {
		return 2
	}
	newKeyframe(e) {
		return this.attributeType === Ze.MultiDimensionProperty ? new lt : new dt
	}
}
class gt extends Qe {
	constructor(e, t, i) {
		super(e, t, i)
	}
	readAttribute(e, t, i, r) {
		et(e, t, i, r, this)
	}
	readValue(e) {
		return (e => {
			const t = e.readEncodeInt32(),
				i = e.readEncodedUint32();
			return new be(t, i)
		})(e)
	}
	readValueList(e, t, i) {
		for (let r = 0; r < i; r++) t[r] = this.readValue(e)
	}
	dimensionality() {
		return 1
	}
	newKeyframe(e) {
		return new dt
	}
}
class yt extends Qe {
	constructor(e, t, i) {
		super(e, t, i)
	}
	readAttribute(e, t, i, r) {
		et(e, t, i, r, this)
	}
	readValue(e) {
		return (e => {
			const t = e.readEncodedUint32();
			if (0 === t) throw new Error("Layer ID is 0");
			const i = new Pe;
			return i.id = t, i
		})(e)
	}
	readValueList(e, t, i) {
		for (let r = 0; r < i; r++) t[r] = this.readValue(e)
	}
	dimensionality() {
		return 1
	}
	newKeyframe(e) {
		return new dt
	}
}
const wt = {
	tagCode: de.LayerAttributes,
	configs: [new ut("isActive", Ze.BitFlag, !0), new ut("autoOrientation", Ze.BitFlag, !1), new yt("parent", Ze
		.Value, void 0), new gt("stretch", Ze.Value, Se), new ft("startTime", Ze.Value, 0), new mt(
			"blendMode", Ze.Value, o.Normal), new mt("trackMatteType", Ze.Value, xe.None), new ct("timeRemap",
				Ze.SimpleProperty, 0), new ft("duration", Ze.FixedValue, 0)]
},
	vt = {
		tagCode: de.LayerAttributesV2,
		configs: [new ut("isActive", Ze.BitFlag, !0), new ut("autoOrientation", Ze.BitFlag, !1), new yt("parent", Ze
			.Value, void 0), new gt("stretch", Ze.Value, Se), new ft("startTime", Ze.Value, 0), new mt(
				"blendMode", Ze.Value, o.Normal), new mt("trackMatteType", Ze.Value, xe.None), new ct("timeRemap",
					Ze.SimpleProperty, 0), new ft("duration", Ze.FixedValue, 0), new class extends Qe {
						constructor(e, t, i) {
							super(e, t, i)
						}
						readAttribute(e, t, i, r) {
							et(e, t, i, r, this)
						}
						readValue(e) {
							return e.readUTF8String()
						}
						readValueList(e, t, i) {
							for (let r = 0; r < i; r++) t[r] = this.readValue(e)
						}
						dimensionality() {
							return 1
						}
						newKeyframe(e) {
							return new dt
						}
					}("name", Ze.Value, "")]
	},
	bt = {
		tagCode: de.Transform2D,
		configs: [new pt("anchorPoint", Ze.SpatialProperty, Te), new pt("position", Ze.SpatialProperty, Te), new ct(
			"xPosition", Ze.SimpleProperty, 0), new ct("yPosition", Ze.SimpleProperty, 0), new pt("scale", Ze
				.MultiDimensionProperty, new Ee(1, 1)), new ct("rotation", Ze.SimpleProperty, 0), new mt("opacity",
					Ze.SimpleProperty, 255)]
	};
de.MaskBlock;
const St = (e, t, i) => {
	switch (t) {
		case de.LayerAttributes:
			Je(e, i, wt), i.duration <= 0 && (i.duration = 1);
			break;
		case de.LayerAttributesV2:
			Je(e, i, vt), i.duration <= 0 && (i.duration = 1);
			break;
		case de.Transform2D:
			i.transform = new We, Je(e, i.transform, bt), i.transform.position.animatable() || i.transform.position
				.value !== Te || !i.transform.xPosition.animatable() && 0 === i.transform.xPosition.value && !i
					.transform.yPosition.animatable() && 0 === i.transform.yPosition.value ? (i.transform.xPosition =
						new ze(0), i.transform.yPosition = new ze(0)) : i.transform.position = new ze(Te);
			break;
		case de.SolidColor:
			i.type() === Re.Solid && function (e, t) {
				t.solidColor = Ae(e), t.width = e.readEncodeInt32(), t.height = e.readEncodeInt32()
			}(e, i);
			break;
		case de.CompositionReference:
			i.type() === Re.PreCompose && function (e, t) {
				const i = e.readEncodedUint32();
				i > 0 && (t.composition = new we, t.composition.id = i), t.compositionStartTime = De(e)
			}(e, i)
	}
},
	Et = (e, t, i) => {
		switch (t) {
			case de.CompositionAttributes:
				Fe(e, i);
				break;
			case de.LayerBlock:
				i.layers.push((e => {
					let t;
					switch (e.readUint8()) {
						case Re.undefined:
							t = new Ke;
							break;
						case Re.Solid:
							t = new je;
							break;
						case Re.Shape:
							t = new He;
							break;
						case Re.PreCompose:
							t = new Xe;
							break;
						default:
							t = new Pe
					}
					return t.id = e.readEncodedUint32(), ue(e, t, St), t
				})(e))
		}
	},
	Tt = e => {
		if (e && 0 === e.length) return;
		const t = new Map;
		for (const i of e) i && (xt(i), t.set(i.id, i));
		let i = 0;
		for (const r of e)
			if (r) {
				if (void 0 !== r.parent) {
					const e = r.parent.id,
						i = t.get(e);
					void 0 !== i && (r.parent = i)
				}
				if (i > 0 && Rt(r.trackMatteType) && (r.trackMatteLayer = e[i - 1]), void 0 !== r.effects && r.effects
					.length > 0)
					for (const e of r.effects) e && (e.type(), Oe.DisplacementMap);
				i += 1
			}
	},
	xt = e => {
		var t;
		if (!e || !e.masks || 0 === e.masks.length) return;
		const i = new Map;
		for (const t of e.masks) t && i.set(t.id, t);
		if (null == (t = e.effects) || t.forEach((e => {
			if (e) {
				if (void 0 !== e.maskReferences && e.maskReferences.length > 0) {
					const t = new Array;
					e.maskReferences.forEach((e => {
						const r = e.id,
							s = i.get(r);
						void 0 !== s && t.push(s)
					})), e.maskReferences = t
				}
				switch (e.type()) {
					case Oe.Fill:
						if (void 0 !== e.fillMask) {
							const t = e.fillMask.id,
								r = i.get(t);
							void 0 !== r && (e.fillMask = r)
						}
						break;
					case Oe.Stroke: {
						const t = e;
						if (void 0 !== t.path) {
							const e = t.path.id,
								r = i.get(e);
							void 0 !== r && (t.path = r)
						}
						break
					}
				}
			}
		})), e.type() === Re.Text) {
			const {
				pathOption: t
			} = e;
			if (null == t ? void 0 : t.path) {
				const e = t.path.id,
					r = i.get(e);
				void 0 !== r && (t.path = r)
			}
		}
	},
	Rt = e => {
		switch (e) {
			case xe.Alpha:
			case xe.AlphaInverted:
				return !0;
			default:
				return !1
		}
	};

function Pt(e, t, i) {
	switch (t) {
		case de.VectorCompositionBlock:
			i.compositions.push((e => {
				const t = new Ne;
				return t.id = e.readEncodedUint32(), ue(e, t, Et), Tt(t.layers), t
			})(e));
			break;
		case de.VideoCompositionBlock:
			i.compositions.push((e => {
				const t = new ve;
				return t.id = e.readEncodedUint32(), t.hasAlpha = e.readBoolean(), ue(e, {
					composition: t,
					hasAlpha: t.hasAlpha
				}, Ie), t
			})(e))
	}
}
const At = e => {
	const t = Dt(e),
		{
			context: i
		} = t;
	ue(t, i, Pt),
		function (e) {
			if (!e || 0 === e.length) return;
			const t = new Map;
			e.forEach((e => {
				e && t.set(e.id, e)
			})), e.forEach((e => {
				if (e && e.type() === le.Vector) {
					const i = e;
					i.layers && i.layers.length > 0 && i.layers.forEach((e => {
						e.containingComposition = i;
						const r = e;
						if (r.type() === Re.PreCompose && r.composition) {
							const e = t.get(r.composition.id);
							e && (r.composition = e)
						}
					}))
				}
			}))
		}(i.compositions);
	const r = i.releaseCompositions();
	return (e => {
		let t = e.length > 0;
		for (const i of e)
			if (!i || !i.verify()) {
				t = !1;
				break
			} if (!t) throw new Error("Verify composition failed!")
	})(r), {
		compositions: r,
		tagLevel: i.tagLevel
	}
},
	Dt = e => {
		if (e.length < 11) throw new Error("PAG file is invalid!");
		const t = e.readInt8(),
			i = e.readInt8(),
			r = e.readInt8();
		if (80 !== t || 65 !== i || 71 !== r) throw new Error("invalid PAG header!");
		return e.readInt8(), e.readUint32(), e.readInt8(), e.readBytes()
	};
class kt {
	constructor(e, t) {
		this.tagLevel = 1, this.compositions = [], this.numLayers = 0, this.scaledTimeRange = {
			start: 0,
			end: 0
		}, this.mainComposition = e[e.length - 1], this.scaledTimeRange.start = 0, this.scaledTimeRange.end =
			this.mainComposition.duration, this.compositions = e, this.duration = this.mainComposition.duration,
			this.implDuration = 1e3 * this.mainComposition.duration / this.mainComposition.frameRate;
		for (const t of e)
			if (t.type() === le.Vector)
				for (const e of t.layers) e.type() !== Re.PreCompose && (this.numLayers += 1);
			else this.numLayers += 1;
		this.tagLevel = t
	}
	static fromArrayBuffer(e) {
		if (!e || 0 === e.byteLength) throw new Error("Can't read empty array buffer!");
		const t = new Le(e, !0),
			{
				compositions: i,
				tagLevel: r
			} = At(t);
		return new kt(i, r)
	}
	getVideoSequence() {
		const e = this.mainComposition.type();
		return e === le.Video ? Ft(this.mainComposition) : e === le.Vector ? Vt(this.mainComposition) : void 0
	}
}
const Ft = e => {
	if (!e.sequences || 0 === e.sequences.length) throw new Error("PAGFile has no BMP video sequence!");
	return e.sequences[e.sequences.length - 1]
},
	Vt = e => {
		const t = Ct(e);
		if (t.length > 1) throw new Error("PAGFile has more than one BMP video sequence!");
		if (t.length < 1) throw new Error("PAGFile has no BMP video sequence!");
		const i = t[0];
		return Ft(i)
	},
	Ct = e => {
		const t = [];
		return e.layers.forEach((e => {
			if (e.type() !== Re.PreCompose) return;
			const {
				composition: i
			} = e;
			(null == i ? void 0 : i.type()) !== le.Video ? (null == i ? void 0 : i.type()) === le.Vector &&
				t.push(...Ct(i)) : t.push(i)
		})), t
	};
class Lt extends he {
	constructor(e, t) {
		super(e, t, {
			renderingMode: O.WebGL
		}), this.frameData = void 0, this.needGetFrame = !1, this.flushBaseTime = Math.floor(1e3 / this
			.videoSequence.frameRate), this.videoReader.start()
	}
	static init(e, t) {
		const i = new s,
			r = kt.fromArrayBuffer(e);
		i.mark("decode");
		const a = new Lt(r, t);
		return a.setDebugData({
			decodePAGFile: i.measure("", "decode")
		}), a
	}
	async play() {
		this.playing || (this.playing = !0, await this.videoReader.start(), this.flushLoop(), 0 === this
			.getProgress() && this.eventManager.emit(q.onAnimationStart), this.eventManager.emit(q
				.onAnimationPlay))
	}
	pause() {
		this.playing && (this.clearTimer(), this.playing = !1, this.eventManager.emit(q.onAnimationPause))
	}
	stop() {
		this.clearTimer(), this.seekToStart(), this.clearRender(), this.playing = !1, this.eventManager.emit(q
			.onAnimationCancel)
	}
	destroy() {
		this.clearTimer(), this.videoReader.destroy(), this.clearRender(), this.canvas = null, this.destroyed = !0
	}
	updateSize() { }
	setProgress(e) {
		if (e < 0 || e > 1) throw new Error("progress must be between 0.0 and 1.0!");
		const t = Math.round(e * this.videoSequence.frameCount);
		if (this.currentFrame !== t) return t !== this.currentFrame + 1 && (this.needSeek = !0), this
			.needGetFrame = !0, void (this.currentFrame = t);
		this.needSeek = !1, this.needGetFrame = !1
	}
	async flushInternal() {
		const t = () => {
			const e = Date.now();
			this.draw(), this.setDebugData({
				draw: Date.now() - e
			}), this.updateFPS(), this.eventManager.emit(q.onAnimationUpdate)
		};
		if (this.needSeek && (this.needSeek = !1, await this.videoReader.seek(this.currentFrame / this.frameRate() *
			1e3)), this.needGetFrame) {
			this.needGetFrame = !1;
			const i = Date.now();
			this.videoReader.getFrameData((r => {
				this.frameData = r, this.currentFrame = r.id, this.setDebugData({
					getFrame: Date.now() - i
				}), e && (this.scale = {
					x: 16 * Math.ceil(this.frameData.width / 16) / this.frameData.width,
					y: 16 * Math.ceil(this.frameData.height / 16) / this.frameData.height
				}), t(), this.currentFrame === this.videoSequence.frameCount - 1 && this.repeat()
			}))
		} else t()
	}
	flushLoop() {
		return this.renderTimer = setInterval((async () => {
			this.needGetFrame = !0, this.flushInternal()
		}), this.flushBaseTime), Promise.resolve()
	}
	detectWebGLContext() {
		return !0
	}
	createVideoReader(e) {
		const {
			videoReader: t,
			debugData: i
		} = I.create(e);
		return this.setDebugData(i), t
	}
	texImage2D() {
		this.gl.texImage2D(this.gl.TEXTURE_2D, 0, this.gl.RGBA, this.frameData.width, this.frameData.height, 0, this
			.gl.RGBA, this.gl.UNSIGNED_BYTE, new Uint8Array(this.frameData.data))
	}
	repeat() {
		return 0 === this.repeatCount ? (this.clearTimer(), this.seekToStart(), this.playing = !1, this.eventManager
			.emit("onAnimationEnd"), Promise.resolve(!1)) : (this.repeatCount -= 1, this.seekToStart(), this
				.eventManager.emit("onAnimationRepeat"), Promise.resolve(!0))
	}
	clearTimer() {
		this.renderTimer && (clearTimeout(this.renderTimer), this.renderTimer = null), this.videoReader
			.clearCallback()
	}
	async seekToStart() {
		this.videoReader.seek(0), this.currentFrame = -1
	}
}
export const PAGView = Lt;
export const clearCache = () => (e => {
	try {
		const t = i.readdirSync(e);
		return t.length > 0 && t.forEach((t => {
			i.unlinkSync(`${e}${t}`)
		})), !0
	} catch (e) {
		return console.error(e), !1
	}
})(t); export const types = N;
