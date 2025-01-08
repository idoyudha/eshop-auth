# Auth Service
Part of [eshop](https://github.com/idoyudha/eshop) Microservices Architecture.

## Overview
This service handles user authorization from all internal backend services that will redirect to [AWS Cognito](https://aws.amazon.com/cognito/). All of customer identity and access management (CIAM) are handled by [AWS Cognito](https://aws.amazon.com/cognito/).

## Architecture
```
eshop-auth
├── .github
│   └── workflows   # github workflows to automatically test, build, and push
├── cmd
│   └── app         # configuration and log initialization
│── config          # configuration
│── internal   
│   └── app         # one run function in the `app.go`
│   └── controller  # serve handler layer
│       └── http    
│           └── v1  # rest http
│   └── entity      # entities of business logic (models) can be used in any layer
│   └── usecase     # business logic
│       └── webapi  # API that business logic works with
└── pkg
    │── cognito     # initialize aws cognito
    │── httpserver  # initialize httpserver
    └── logger      # initialize logger
```

## Tech Stack
- Backend: Go
- CI/CD: Github Actions
- IAM: AWS Cognito
- Container: Docker

## API Documentation
tbd