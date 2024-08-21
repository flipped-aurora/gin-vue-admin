// import axios from "axios";
import service from "@/utils/request";

export const getUploadToken = async function () {

     let res= await service({
        url: '/bizAppHub/getUploadToken',
        method: 'get',
        params: null
    })

    // let res = await axios.get('/bizAppHub/getUploadToken');
    // console.log(res.data.data.token);
    console.log("res.data.token:======",res.data.token)
    return res.data.token;
};