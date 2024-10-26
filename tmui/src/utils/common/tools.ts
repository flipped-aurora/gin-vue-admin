

//获取链接中的邀请地址
export const getQueryVariable = (variable:string)=> {
  const query = window.location.search.substring(1);
  const vars = query.split("&");
  for (var i=0;i<vars.length;i++) {
    const pair = vars[i].split("=");
    if(pair[0] == variable){
      return pair[1];
    }
  }
  return('0x00000x0000');
}