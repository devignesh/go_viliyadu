# Golang with jwt authendication application.

    1. Go
    2. gin-gonic
    3. mongo
    4. go-jwt

# How to run

    1. Clone this repo
        git clone {{git url}}

    2. Navigate to the project folder

    3. Build the application
        go build {{name}}

    4. Run
        ./ {{name}}

# Postman Collection

    1. Signup Curl: localhost:8888/users/signup

        curl --location --request POST 'localhost:8888/users/signup' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "First_name":"vicky",
                "Last_name":"smazz",
                "Password":"qwerty@123",
                "Email":"vigneshkumar.mca216@adhiyamaan.in",
                "Phone":"9047660920",
                "User_type":"ADMIN"
            }'

        Response:

            {
                "InsertedID": "61d9a21b3ada8a68884245ff"
            }

    2. Login Curl: localhost:8888/users/login

        curl --location --request POST 'localhost:8888/users/login' \
            --header 'Content-Type: application/json' \
            --data-raw '{
                "Email":"vigneshkumar.mca216@adhiyamaan.in",
                "Password":"qwerty@123"
            }'

        Response:

            {
                "ID": "61d9a21b3ada8a68884245ff",
                "first_name": "vicky",
                "last_name": "smazz",
                "Password": "$2a$14$lPKbBFLpkNj8Rvtprwd/5.CWQCoNPye3ooDR/u5g0AHZH/jjLnorO",
                "email": "vigneshkumar.mca216@adhiyamaan.in",
                "phone": "9047660920",
                "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InZpZ25lc2hrdW1hci5tY2EyMTZAYWRoaXlhbWFhbi5pbiIsIkZpcnN0X25hbWUiOiJ2aWNreSIsIkxhc3RfbmFtZSI6InNtYXp6IiwiVWlkIjoiNjFkOWEyMWIzYWRhOGE2ODg4NDI0NWZmIiwiVXNlcl90eXBlIjoiQURNSU4iLCJleHAiOjE2NDE4MzcwMDF9.zgXhJqwqHoNcG2-SM2oE8VZFhD6aelDMWoNoRXoimW8",
                "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X25hbWUiOiIiLCJMYXN0X25hbWUiOiIiLCJVaWQiOiIiLCJVc2VyX3R5cGUiOiIiLCJleHAiOjE2NDIzNTU0MDF9.8LQbikzo1MDOvsDKFoWib4Ru88jtxJoz6nKbkVB4uqc",
                "user_type": "ADMIN",
                "created_at": "2022-01-08T14:39:23Z",
                "updated_at": "2022-01-09T17:50:01Z",
                "user_id": "61d9a21b3ada8a68884245ff"
            }
