const axios = require("axios");

const apiUrl = 'https://api.pushe.co/v2/messaging/notifications/';

// You can find your api key token in your console.pushe.co and in your account
// Your token would be like ea8f6698565927e0*******56be614a2839217dc
const token = 'Your Token Goes Here ...';

// Documentation is available https://pushe.co/docs/webpush-api/
const data = {
    app_ids: [],
    platform: 2,
    data: {
        title: "تیتر پیام",
        content: "متن پیام"
    }
};

// Use axios library to send post request
// Axios documentation : https://github.com/axios/axios
const response = axios({
    method: 'post',
    url: apiUrl,
    headers: {
        authorization: 'Token ' + token,
        'content-type': 'application/json'
    },
    data
});

// To get the response you can use promise then, catch
response.then((resp) => {
    console.log('Response is: ', resp.data);
}).catch((error) => {
    console.log('Error is: ', error.response.data, error.response.status);
});
