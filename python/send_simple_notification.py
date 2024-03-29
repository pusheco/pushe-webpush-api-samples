#!/usr/bin/env python
# vim: ts=4 sw=4 et

import requests

# Obtain token -> https://docs.pushe.co/docs/web-api/authentication
TOKEN = 'YOUR_TOKEN'

# set header
headers = {
    'Authorization': 'Token ' + TOKEN,
    'Content-Type': 'application/json'
}

# Webpush doc -> http://docs.pushe.co/docs/web-api/simple-notification/

data = {
    'app_ids': ['YOUR_APP_ID', ],
    'data': {
        'title': 'Title',
        'content': 'Content',
    }
}

# send request
response = requests.post(
    'https://api.pushe.co/v2/messaging/notifications/web/',
    json=data,
    headers=headers,
)

# get status_code and response
print('status code => ', response.status_code)
print('response => ', response.json())
print('==========')

if response.status_code == 201:
    print('Success!')

    data = response.json()

    # hashed_id just generated for Non-Free plan
    if data['hashed_id']:
        report_url = 'https://pushe.co/report?id=%s' % data['hashed_id']
    else:
        report_url = 'no report url for your plan'

    notif_id = data['wrapper_id']
    print('report_url: %s' % report_url)
    print('notification id: %s' % notif_id)
else:
    print('failed')
