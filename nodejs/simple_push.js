var request = require('request');

// More info about device_id:
//    (web): https://pushe.co/docs/webpush/#unique_id

var TOKEN = "YOUR_TOKEN";

// Webpush -> https://pushe.co/docs/webpush-api/

request.post(
    {
        uri: "https://api.pushe.co/v2/messaging/notifications/web/",
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            "Accept": "application/json",
            "Authorization": "Token " + TOKEN,
        },
        body: JSON.stringify({
            "app_ids": ["YOUR_APP_ID"],
            "data": {
                "title": "Title",
                "content": "Content",
            }
            // additional keywords -> https://pushe.co/docs/webpush-api/#api_send_advance_notification
        }),
    },
    function (error, response, body) {
        console.log("status_code: " + response.statusCode);
        console.log("response: " + body);

        if (response.statusCode == 201) {
            console.log("Success!");

            var data = JSON.parse(body);
            var report_url;

            // report url only generated on Non-Free plan
            if (data.hashed_id) {
                report_url = "https://pushe.co/report?id=" + data.hashed_id;
            } else {
                report_url = "no report url for your plan";
            }

            console.log("report_url: " + report_url);

            console.log("notification id: " + data.wrapper_id);
        } else {
            console.log("failed!");
        }
    }
);
