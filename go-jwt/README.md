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

    1. Signup Curl:

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
