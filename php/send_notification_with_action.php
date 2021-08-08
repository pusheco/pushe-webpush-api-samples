<?php
// Obtain token -> https://docs.pushe.co/docs/web-api/authentication
$TOKEN = "YOUR_TOKEN";

$data = array(
    "app_ids" => ["YOUR_APP_ID",],
    "data" => array(
        "title" => "Title",
        "content" => "Content",
        "action" => array(
            "action_type" => "U",
            "url" => "https://pushe.co"
        ),
        "buttons" => array(
            array(
                "btn_action" => array(
                    "action_type" => "U",
                    "url" => "https://pushe.co"
                ),
                "btn_content" => "YOUR_CONTENT",
                "btn_order" => 0,
            ),
            array(
                "btn_action" => array(
                    "action_type" => "U",
                    "url" => "https://pushe.co"
                ),
                "btn_content" => "YOUR_CONTENT",
                "btn_order" => 1,
            )
        ),
    ),
    // additional keywords -> https://docs.pushe.co/docs/web-api/notification-actions
);

// initialize curl
$ch = curl_init("https://api.pushe.co/v2/messaging/notifications/web/");

// set header parameters
curl_setopt($ch, CURLOPT_HTTPHEADER, array(
    "Content-Type: application/json",
    "Accept: application/json",
    "Authorization: Token " . $TOKEN,
));
curl_setopt($ch, CURLOPT_POST, 1);

// set data
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));

$response = curl_exec($ch);

echo "response \n";
print_r($response);


?>
