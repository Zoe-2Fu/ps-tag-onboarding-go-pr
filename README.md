# ps-tag-onboarding-go
## Overview
The project is a Go implementation of a Java application, serving as an onboarding exercise. It leverages the `echo` framework, a fast and minimalist Go web framework. Focus on building a RESTful API for managing user information, with features such as user creation, retrieval, and error handling.

## Onboarding Exercise Details
This repository is part of an onboarding exercise. For the original Java project and onboarding details, please check out: https://wexinc.atlassian.net/wiki/spaces/TGT/pages/153576505378/Developer+Onboarding+Exercise+-+Advanced

## Setup
1. Run the docker container under `/ps-tag-onboarding-go-pr/` folder
    ``` shell
        docker compose up --build -d
    ```

2. Run the below commands to test the api endpoints
- Save new user
    ``` shell
        curl -X POST http://localhost:8080/user \
        -H "Content-Type: application/json" \
        -d '{"firstname": "Sam", "lastname": "Smith", "email": "good@example.com", "age": 20}' 
    ```
(There will display the new user data with objectID in terminal, we can copy & paste it to the following `find` command)

- Find user by id
    ``` shell
        curl localhost:8080/user/657a853c45eb642c2ab54ca3
    ```
3. For testing the unit test, run the below command to run all the tests
    ```
        go test ./...
    ```