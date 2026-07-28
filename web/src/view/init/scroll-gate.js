const SCROLL_BOTTOM_TOLERANCE = 1

export const isScrollAtBottom = ({ scrollTop, clientHeight, scrollHeight }) =>
  scrollTop + clientHeight >= scrollHeight - SCROLL_BOTTOM_TOLERANCE

export const getScrollGateHint = (canConfirm) =>
  canConfirm ? '' : '请滚动阅读至底部后确认'
