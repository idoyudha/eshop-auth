# Auth Service
Part of [eshop](https://github.com/idoyudha/eshop) Microservices Architecture.

## Overview
This service handles user authorization from all internal backend services that will redirect to [AWS Cognito](https://aws.amazon.com/cognito/). All of customer identity and access management (CIAM) are handled by [AWS Cognito](https://aws.amazon.com/cognito/).

## Architecture
eshop-auth
├── .github
|   |── workflows
├── cmd
│   ├── app
|── config
|── internal   
|   |── app
|   |── controller
|       |── http
|           |── v1
|   |── entity
|   |── usecase
|       |── webapi
└── pkg
    |── cognito
    |── httpserver
    |── logger


## Tech Stack
- Backend: Go
- CI/CD: Github Actions
- IAM: AWS Cognito
- Container: Docker

## API Documentation
tbd