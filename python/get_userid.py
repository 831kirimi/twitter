#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import requests
from base64 import b64decode, b64encode
import json

def main():
    bearer_token = "*******"
    username = "Twitter"
    print(username)

    url = "https://api.twitter.com/1.1/users/lookup.json"
    headers = {
        "Authorization": "Bearer {}".format(bearer_token)
    }
    params = {
        "screen_name": username,
    }

    r = requests.get(url, headers=headers, params=params)
    user_id = r.json()[0]["id"]
    print(user_id)


if __name__ == "__main__":
    main()
