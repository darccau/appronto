#!/usr/bin/env bash

app='localhost:4000'
body=' { "first_name": "darccau", "last_name": "forgemaster", "email": "darccau@mail.com", "password": "forgemaster" }'


# User querys
## Insert User 
curl -X POST $app/v1/user -d '{"first_name": "darccau", "last_name": "forgemaster", "password": "password", "email": "darccau@mail.com"}'

## Activate user
curl -X PUT $app/v1/user/activated '{"token": "2ZKJ7B4PF4IA7OC4OFTWVMNQHU"}'

## Authenticate user
curl -X POST $app/v1/token/authentication -d '{"password": "password", "email": "darccau@mail.com"}'

## Check user permissions with authentication
curl -X POST $app/v1/appointments -H "Authorization: Bearer AHG4MAECEONVK5LIWZHWWJ7DSM" 

# Appointments querys

## Insert
curl $app/v1/appointments -H "Authorization: Bearer AHG4MAECEONVK5LIWZHWWJ7DSM"
