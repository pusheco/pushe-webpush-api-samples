<?php

$TOKEN = "YOUR_TOKEN";

// Webpush -> https://pushe.co/docs/webpush-api/

$data = array(
    "app_ids" => ["YOUR_APP_ID",],

    "data" => array(
        "title" => "Title",
        "content" => "Content",
    ),
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
curl_setopt($ch, CURLOPT_RETURNTRANSFER, true);

// set data
curl_setopt($ch, CURLOPT_POSTFIELDS, json_encode($data));

$response = curl_exec($ch);

// report url only generated on Non-Free plan
$hashed_id = json_decode($response)->hashed_id;

print_r($hashed_id);

curl_close($ch);
?>
