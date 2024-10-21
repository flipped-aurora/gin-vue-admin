export interface childrenData {
    id: number | string,
    text: string,
    disabled?: boolean,
    children?: Array<childrenData>
}