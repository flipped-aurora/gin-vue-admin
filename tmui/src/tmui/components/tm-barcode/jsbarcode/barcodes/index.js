import { CODE39 } from './CODE39/';
import { CODE128, CODE128A, CODE128B, CODE128C } from './CODE128/';
import { EAN13, EAN8, EAN5, EAN2, UPC, UPCE } from './EAN_UPC/';
import { ITF, ITF14 } from './ITF/';
import { MSI, MSI10, MSI11, MSI1010, MSI1110 } from './MSI/';
import { pharmacode } from './pharmacode/';
import { codabar } from './codabar';
import { GenericBarcode } from './GenericBarcode/';

export default {
	CODE39,
	CODE128, CODE128A, CODE128B, CODE128C,
	EAN13, EAN8, EAN5, EAN2, UPC, UPCE,
	ITF14,
	ITF,
	MSI, MSI10, MSI11, MSI1010, MSI1110,
	pharmacode,
	codabar,
	GenericBarcode
};
