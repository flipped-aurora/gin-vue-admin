import { A_START_CHAR, B_START_CHAR, C_START_CHAR, A_CHARS, B_CHARS, C_CHARS } from './constants';

// Match Set functions
const matchSetALength = (string) => string.match(new RegExp(`^${A_CHARS}*`))[0].length;
const matchSetBLength = (string) => string.match(new RegExp(`^${B_CHARS}*`))[0].length;
const matchSetC = (string) => string.match(new RegExp(`^${C_CHARS}*`))[0];

// CODE128A or CODE128B
function autoSelectFromAB(string, isA) {
	const ranges = isA ? A_CHARS : B_CHARS;
	const untilC = string.match(new RegExp(`^(${ranges}+?)(([0-9]{2}){2,})([^0-9]|$)`));

	if (untilC) {
		return (
			untilC[1] +
			String.fromCharCode(204) +
			autoSelectFromC(string.substring(untilC[1].length))
		);
	}

	const chars = string.match(new RegExp(`^${ranges}+`))[0];

	if (chars.length === string.length) {
		return string;
	}

	return (
		chars +
		String.fromCharCode(isA ? 205 : 206) +
		autoSelectFromAB(string.substring(chars.length), !isA)
	);
}

// CODE128C
function autoSelectFromC(string) {
	const cMatch = matchSetC(string);
	const length = cMatch.length;

	if (length === string.length) {
		return string;
	}

	string = string.substring(length);

	// Select A/B depending on the longest match
	const isA = matchSetALength(string) >= matchSetBLength(string);
	return cMatch + String.fromCharCode(isA ? 206 : 205) + autoSelectFromAB(string, isA);
}

// Detect Code Set (A, B or C) and format the string
export default (string) => {
	let newString;
	const cLength = matchSetC(string).length;

	// Select 128C if the string start with enough digits
	if (cLength >= 2) {
		newString = C_START_CHAR + autoSelectFromC(string);
	} else {
		// Select A/B depending on the longest match
		const isA = matchSetALength(string) > matchSetBLength(string);
		newString = (isA ? A_START_CHAR : B_START_CHAR) + autoSelectFromAB(string, isA);
	}

	return newString.replace(
		/[\xCD\xCE]([^])[\xCD\xCE]/, // Any sequence between 205 and 206 characters
		(match, char) => String.fromCharCode(203) + char
	);
};
