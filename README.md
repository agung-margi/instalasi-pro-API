# instalasi-pro-API

# API Documentation

This document provides a comprehensive overview of the API endpoints for user management, including login, registration, updating user details, and retrieving user information.

## Base URL

## Endpoints

### 1. User Login

Login endpoint for user authentication.

Example URL: http://localhost:8080/api/users/login

**Request**

```json
{
  "email": "customer1@mail.com",
  "password": "password"
}
```

**Response**

```json
{
  "meta": {
    "message": "Login success",
    "code": 200,
    "status": "success"
  },
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

### 2. User Registration

Registration endpoint for user registration.

Example URL: http://localhost:8080/api/users/register

**Request**

```json
{
  "email": "customer2@gmail.com",
  "password": "password"
}
```

**Response**

```json
{
  "meta": {
    "message": "success",
    "code": 200,
    "status": "success"
  },
  "data": {
    "email": "customer2@gmail.com",
    "role": "customer"
  }
}
```

### 3. Update User Details

Update user details endpoint for updating user details.

Example URL: http://localhost:8080/api/users/:2

**Request**

```json
{
  "email": "customer2@mail.com",
  "name": "customer nama",
  "address": "customer address 2",
  "phone": "0855445787"
}
```

**Response**

```json
{
  "meta": {
    "message": "success",
    "code": 200,
    "status": "success"
  },
  "data": {
    "id": 2,
    "name": "customer nama",
    "email": "customer2@mail.com",
    "role": "customer",
    "address": "customer address 2",
    "phone": "0855445787",
    "updated_at": "2024-10-21T11:49:18.751987+07:00"
  }
}
```

### 4. Get User Details

Get user details endpoint for retrieving user details.

Example URL: http://localhost:8080/api/users/:2

**Response**

```json
{
  "meta": {
    "message": "success",
    "code": 200,
    "status": "success"
  },
  "data": {
    "id": 2,
    "name": "customer nama",
    "email": "customer2@mail.com",
    "role": "customer",
    "address": "customer address 2",
    "phone": "0855445787",
    "updated_at": "2024-10-21T11:49:18.751987+07:00"
  }
}
```

## Notes

### 1. JWT

JWT is a standard for authentication and authorization. It is used for securely transmitting information between parties who are independent of each other.

### 2. Database

PostgreSQL is a relational database management system.
