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

# Login Curl: localhost:8888/users/login

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

# Get Users: Admin login | localhost:8888/users

    curl --location --request GET 'localhost:8888/users' \
     --header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InZpZ25lc2hrdW1hci5tY2EyMTZAYWRoaXlhbWFhbi5pbiIsIkZpcnN0X25hbWUiOiJ2aWNreSIsIkxhc3RfbmFtZSI6InNtYXp6IiwiVWlkIjoiNjFkOWEyMWIzYWRhOGE2ODg4NDI0NWZmIiwiVXNlcl90eXBlIjoiQURNSU4iLCJleHAiOjE2NDE4MzcwMDF9.zgXhJqwqHoNcG2-SM2oE8VZFhD6aelDMWoNoRXoimW8'

    Response:

        {
            "total_count": 5,
            "user_items": [
                {
                    "_id": "61d09fe708fdfe83625749cb",
                    "created_at": "2022-01-02T00:09:35+05:30",
                    "email": "vinesh1865@gmail.com",
                    "first_name": "vigneshkumar",
                    "last_name": "Dharmalingam",
                    "password": "$2a$14$Be780/eU26G53s.AfCF7POQgglKuUTJyS6iFtjdYEZUZJ50x0rfn6",
                    "phone": "1234567890",
                    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X25hbWUiOiIiLCJMYXN0X25hbWUiOiIiLCJVaWQiOiIiLCJVc2VyX3R5cGUiOiIiLCJleHAiOjE2NDE2Njc5NjV9.08-y1LBa1up4a423aIg4tertZ32yfh8Jw4mX7BaDzVE",
                    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InZpbmVzaDE4NjVAZ21haWwuY29tIiwiRmlyc3RfbmFtZSI6InZpZ25lc2hrdW1hciIsIkxhc3RfbmFtZSI6IkRoYXJtYWxpbmdhbSIsIlVpZCI6IjYxZDA5ZmU3MDhmZGZlODM2MjU3NDljYiIsIlVzZXJfdHlwZSI6IkFETUlOIiwiZXhwIjoxNjQxMTQ5NTY1fQ.CcaPiIQ4SNnQH0zy8NaTGMpCoAYnIsKMjUkqDPqMaZg",
                    "updated_at": "2022-01-02T00:22:45+05:30",
                    "user_id": "61d09fe708fdfe83625749cb",
                    "user_type": "ADMIN"
                },
                {
                    "_id": "61d0a24708fdfe83625749cd",
                    "created_at": "2022-01-02T00:19:43+05:30",
                    "email": "vinesh1865+demo@gmail.com",
                    "first_name": "vicky",
                    "last_name": "Dharmalingam",
                    "password": "$2a$14$DGCi5G8JqJlr.JDuyHAa1erUSUDOLe4l.SgBboZnfCQAlrLrXcIXi",
                    "phone": "1234567891",
                    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X25hbWUiOiIiLCJMYXN0X25hbWUiOiIiLCJVaWQiOiIiLCJVc2VyX3R5cGUiOiIiLCJleHAiOjE2NDE2Njc4MTR9.ipYdSVsaTrGRJGlGAJ707gEddQXfsJvxEN1xPH5Z0bE",
                    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InZpbmVzaDE4NjUrZGVtb0BnbWFpbC5jb20iLCJGaXJzdF9uYW1lIjoidmlja3kiLCJMYXN0X25hbWUiOiJEaGFybWFsaW5nYW0iLCJVaWQiOiI2MWQwYTI0NzA4ZmRmZTgzNjI1NzQ5Y2QiLCJVc2VyX3R5cGUiOiJVU0VSIiwiZXhwIjoxNjQxMTQ5NDE0fQ.23BR0qw7Br8WjXfU_fX2i4SSOkzLM7eF_6vPYaPSRN8",
                    "updated_at": "2022-01-02T00:20:14+05:30",
                    "user_id": "61d0a24708fdfe83625749cd",
                    "user_type": "USER"
                },
                {
                    "_id": "61d0a2b365d7e3a383ed7e51",
                    "created_at": "2022-01-02T00:21:31+05:30",
                    "email": "vinesh1865+demo1@gmail.com",
                    "first_name": "vicky kumar",
                    "last_name": "dukey",
                    "password": "$2a$14$wHgzxw2Zf8T5G.fOcb62Tu3VbzWFEIkIx8.4N3F4BE0lvmvdivhl.",
                    "phone": "1234567892",
                    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X25hbWUiOiIiLCJMYXN0X25hbWUiOiIiLCJVaWQiOiIiLCJVc2VyX3R5cGUiOiIiLCJleHAiOjE2NDE2Njc4OTF9.QktwKOiexc_X0GvfCHOmzZHAoWAZHWHsCFXxC1zjZs0",
                    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InZpbmVzaDE4NjUrZGVtbzFAZ21haWwuY29tIiwiRmlyc3RfbmFtZSI6InZpY2t5IGt1bWFyIiwiTGFzdF9uYW1lIjoiZHVrZXkiLCJVaWQiOiI2MWQwYTJiMzY1ZDdlM2EzODNlZDdlNTEiLCJVc2VyX3R5cGUiOiJVU0VSIiwiZXhwIjoxNjQxMTQ5NDkxfQ.ziibuyKODQFxMqWuMFb1EfYPqYJ_8x4RYd7z-yUL-Bg",
                    "updated_at": "2022-01-02T00:21:31+05:30",
                    "user_id": "61d0a2b365d7e3a383ed7e51",
                    "user_type": "USER"
                },
                {
                    "_id": "61d46a3aafe036d989f6d88e",
                    "created_at": "2022-01-04T21:09:38+05:30",
                    "email": "vinesh1865+demo2@gmail.com",
                    "first_name": "vicky vicky",
                    "last_name": "koli",
                    "password": "$2a$14$dPWTf3kRWCJ85ozoivNXPO9TbzFMiU9YSGd2DPiXbrP6xxnqxndLG",
                    "phone": "1234567892",
                    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X25hbWUiOiIiLCJMYXN0X25hbWUiOiIiLCJVaWQiOiIiLCJVc2VyX3R5cGUiOiIiLCJleHAiOjE2NDE5MTU5NTl9.rhrdE5a1Zp5KKYxxO07fANPUxZeJoaBPEyGH5SKltM4",
                    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InZpbmVzaDE4NjUrZGVtbzJAZ21haWwuY29tIiwiRmlyc3RfbmFtZSI6InZpY2t5IHZpY2t5IiwiTGFzdF9uYW1lIjoia29saSIsIlVpZCI6IjYxZDQ2YTNhYWZlMDM2ZDk4OWY2ZDg4ZSIsIlVzZXJfdHlwZSI6IkFETUlOIiwiZXhwIjoxNjQxMzk3NTU5fQ.6k2TFIBUQFVz5TZmy15bVdWjLxN3kthuuvP2q9p9T1o",
                    "updated_at": "2022-01-04T21:15:59+05:30",
                    "user_id": "61d46a3aafe036d989f6d88e",
                    "user_type": "ADMIN"
                },
                {
                    "_id": "61d9a21b3ada8a68884245ff",
                    "created_at": "2022-01-08T20:09:23+05:30",
                    "email": "vigneshkumar.mca216@adhiyamaan.in",
                    "first_name": "vicky",
                    "last_name": "smazz",
                    "password": "$2a$14$lPKbBFLpkNj8Rvtprwd/5.CWQCoNPye3ooDR/u5g0AHZH/jjLnorO",
                    "phone": "9047660920",
                    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6IiIsIkZpcnN0X25hbWUiOiIiLCJMYXN0X25hbWUiOiIiLCJVaWQiOiIiLCJVc2VyX3R5cGUiOiIiLCJleHAiOjE2NDIzNTU0MDF9.8LQbikzo1MDOvsDKFoWib4Ru88jtxJoz6nKbkVB4uqc",
                    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InZpZ25lc2hrdW1hci5tY2EyMTZAYWRoaXlhbWFhbi5pbiIsIkZpcnN0X25hbWUiOiJ2aWNreSIsIkxhc3RfbmFtZSI6InNtYXp6IiwiVWlkIjoiNjFkOWEyMWIzYWRhOGE2ODg4NDI0NWZmIiwiVXNlcl90eXBlIjoiQURNSU4iLCJleHAiOjE2NDE4MzcwMDF9.zgXhJqwqHoNcG2-SM2oE8VZFhD6aelDMWoNoRXoimW8",
                    "updated_at": "2022-01-09T23:20:01+05:30",
                    "user_id": "61d9a21b3ada8a68884245ff",
                    "user_type": "ADMIN"
                }
            ]
        }

# Get User: Based on ID | localhost:8888/users/61d9a21b3ada8a68884245ff

        curl --location --request GET 'localhost:8888/users/61d9a21b3ada8a68884245ff' \
         --header 'token: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJFbWFpbCI6InZpZ25lc2hrdW1hci5tY2EyMTZAYWRoaXlhbWFhbi5pbiIsIkZpcnN0X25hbWUiOiJ2aWNreSIsIkxhc3RfbmFtZSI6InNtYXp6IiwiVWlkIjoiNjFkOWEyMWIzYWRhOGE2ODg4NDI0NWZmIiwiVXNlcl90eXBlIjoiQURNSU4iLCJleHAiOjE2NDE4MzcwMDF9.zgXhJqwqHoNcG2-SM2oE8VZFhD6aelDMWoNoRXoimW8'

        Response:

            {
                "Message": "requst accepted",
                "data": {
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
            }
