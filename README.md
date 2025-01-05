# A Restful API for sending email and sms with Golang

This is a simple RESTful API for sending email and sms with Golang. It is built with the Gin web framework and GORM for database operations. The API provides endpoints for sending email and sms, as well as user authentication with JWT.

## Features

<ul>
    <li>RESTful API with Gin.</li>
    <li>Sending mails and sms with Twilio and SMTP </li>
    <li>Middleware for securing endpoints.</li>
    <li>Configuration via environment variables.</li>
    <li>Modular and clean project structure.</li>
</ul>

## Prerequisites

<ul>
  <li>Go 1.19 or newer</li>
  <li><a href="https://www.postman.com/" target="_blank">Postman</a> (optional, for API testing)</li>
</ul>

## Installation

* Clone the repository:
   ``` bash
    git clone https://github.com/tnqbao/gau_validation_service.git
    cd gau_validation_service
   ```
* Setup your module:
  ``` bash
   go mod edit -module=your-link-github-repo 
  ```
* Install dependencies:
  ``` bash
    go mod tidy 
  ``` 

* Set up environment variables:
    * Create a `.env` file in the project root and configure the following variables:
  ```dotenv
    JWT_SECRET=your_jwt_secret
  
    SMTP_USERNAME=your_smtp_username
    SMTP_PASSWORD=your_smtp_password
  
    REDIS_ADDR=your-host:6379
    REDIS_PASSWORD=your-redis-password
    REDIS_DB=int-number
  
    TWILIO_ACCOUNT_SID=your_twilio_account_sid
    TWILIO_AUTH_TOKEN=your_twilio_auth_token
    TWILIO_PHONE_NUMBER=your_twilio_phone_number
  
    GLOBAL_DOMAIN=your_global_domain
    LINK_DOMAIN=your_link_domain
    ```
* Run database migrations:
    ``` bash
     go run main.go migrate
   ```

* Start the server:
    ``` bash 
    go run main.go
    ```

  <li>Access the API at: <a href="http://localhost:8080" target="_blank">http://localhost:8080</a></li>

## Project Structure
   ``` 
   Directory structure:
└── tnqbao-gau_validation_service/
    ├── Dockerfile
    ├── go.mod
    ├── go.sum
    ├── main.go
    ├── api/
    │   ├── healthcheck/
    │   │   └── healthcheck.go
    │   ├── mail/
    │   │   ├── checkOTP.go
    │   │   ├── mailSender.go
    │   │   └── sendOTP.go
    │   └── sms/
    │       ├── checkOTP.go
    │       └── sendOTP.go
    ├── config/
    │   ├── database.go
    │   └── redis.go
    ├── middlewares/
    │   ├── auth.go
    │   └── cors.go
    ├── providers/
    │   ├── helpers.go
    │   └── types.go
    ├── routes/
    │   └── routes.go
    └── .github/
        └── workflows/
            ├── deploy-production.yml
            └── deploy-staging.yml

   ```



<h2>Future Improvements</h2>
<ul>
  <li>Add unit and integration tests.</li>
  <li>Implement role-based access control (RBAC).</li>
  <li>Add API versioning.</li>
  <li>Improve error handling and logging.</li>
</ul>

<h2>License</h2>
<p>This project is licensed under the MIT License. See the <a href="LICENSE" target="_blank">LICENSE</a> file for details.</p>
