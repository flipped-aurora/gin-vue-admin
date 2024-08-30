import service from '@/utils/request'


export const getUserInfo = async function (username) {

    let res= await service({
        url: '/bizAppHub/getUserInfo',
        method: 'get',
        params:{user:username}
    })

    // let res = await axios.get('/bizAppHub/getUploadToken');
    // console.log(res.data.data.token);
    console.log("getUserInfo:======",res.data)
    return res;
};