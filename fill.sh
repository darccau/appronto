#!/usr/bin/env bash

app='localhost:4000'

# Check healthcheck
curl $app/v1/healthcheck

# Insert User 
curl -X POST $app/v1/user -d '{"first_name": "darccau", "last_name": "silva", "password": "sample1", "email": "darccau@mail.com"}'
curl -X POST $app/v1/user -d '{"first_name": "jao", "last_name": "sauro", "password": "sample2", "email": "jao@mail.com"}'
curl -X POST $app/v1/user -d '{"first_name": "pedro", "last_name": "zulu", "password": "sample3", "email": "pedro@mail.com"}'
curl -X POST $app/v1/user -d '{"first_name": "jao", "last_name": "test", "password": "sample4", "email": "patter@mail.com"}'
curl -X POST $app/v1/user -d '{"first_name": "zezin", "last_name": "aaa", "password": "sample5", "email": "pedra@mail.com"}'
curl -X POST $app/v1/user -d '{"first_name": "prole", "last_name": "bbbb", "password": "sample6", "email": "pero@mail.com"}'
curl -X POST $app/v1/user -d '{"first_name": "mofilho", "last_name": "lllll", "password": "sample7", "email": "padre@mail.com"}'

# Update User
curl -X PATH $app/v1/user -d '{"first_name": "son", "last_name": "sonson"}'

curl -X GET $app/v1/user/2
#
# curl -X PATCH $app/v1/user/1 -d '{"first_name": "forgemaster", "last_name": "test"}'
#
#  curl -w '\nTime: %{time_total}s \n' $app/v1/user/1

