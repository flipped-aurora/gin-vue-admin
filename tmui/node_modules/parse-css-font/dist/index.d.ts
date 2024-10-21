export interface ISystemFont {
    system: string;
}
export interface IFont {
    style?: string;
    variant?: string;
    weight?: string;
    stretch?: string;
    size?: string;
    lineHeight?: string | number;
    family?: string[];
}
export default function parseCSSFont(value: string): ISystemFont | IFont;
