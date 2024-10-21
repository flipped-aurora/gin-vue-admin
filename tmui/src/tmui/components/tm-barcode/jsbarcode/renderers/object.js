class ObjectRenderer {
	constructor(object, encodings, options) {
		this.object = object;
		this.encodings = encodings;
		this.options = options;
	}

	render() {
		this.object.encodings = this.encodings;
	}
}


export default ObjectRenderer;
