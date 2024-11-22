# Chirpy

Chirpy is a back end web server for a social network similar to Twitter.

The code in this repository primarily represents an API that can be consumed by a frond end mobile app or web server to manage user accounts and Chirps (_short text messages_).

The back end code is designed to integrate with a PostgreSQL database.

## API Documentation

The Chirpy API follows RESTful conventions.

### Endpoint: Create a New User

#### Description
Creates a new user in the system

#### HTTP Request
**Method:** `POST`

**URL:** `/api/users`

#### Headers
| Name           | Type   | Description                |
|----------------|--------|----------------------------|
| `Content-Type` | string | Must be `application/json` |

#### Request Body
Send a JSON object with the following fields:
| Field Name | Type   | Required | Description                            |
|------------|--------|----------|----------------------------------------|
| `Email`    | string | Yes      | The unique email address for the user. |
| `Password` | string | Yes      | The user's password.                   |

Example:
```json
{
  "email": "user@example.com",
  "password": "my-secure-password-123"
}
```

#### Response

##### Success
**Status Code:** `201`
**Response Body:**
```json
{
  "id": "efe3be18-0116-43cf-b775-f81136f04d41",
  "email": "user@example.com",
  "created_at": "2024-11-22T11:16:37.466073Z",
  "updated_at": "2024-11-22T11:16:37.466073Z",
  "is_chirpy_red": false
}
```

##### Failure
**Status Code:** `500`
**Response Body (server side problems):**
```json
{
    "error": "Couldn't decode parameters",
}
```

```json
{
    "error": "There was a problem hashing your password",
}
```

```json
{
    "error": "Couldn't create user in the database",
}
```

### Endpoint: Update a User's Email and/or Password

#### Description
Updates an existing user account's email address and/or password

#### HTTP Request
**Method:** `PUT`

**URL:** `/api/users`

#### Headers
| Name            | Type   | Description                            |
|-----------------|--------|----------------------------------------|
| `Content-Type`  | string | Must be `application/json`             |
| `Authorization` | string | Bearer token for API access. Required. |

#### Request Body
Send a JSON object with the following fields:
| Field Name | Type   | Required | Description                            |
|------------|--------|----------|----------------------------------------|
| `Email`    | string | Yes      | The unique email address for the user. |
| `Password` | string | Yes      | The user's password.                   |


Example:
```json
{
  "email": "user@example.com",
  "password": "my-new-password-123"
}
```

#### Response

##### Success
**Status Code:** `200`
**Response Body:**
```json
{
  "id": "efe3be18-0116-43cf-b775-f81136f04d41",
  "email": "user@example.com",
  "created_at": "2024-11-22T11:16:37.466073Z",
  "updated_at": "2024-11-24T11:16:37.466073Z",
  "is_chirpy_red": false
}
```

##### Failure
**Status Code:** `401`
**Response Body (Missing or invalid authentication token):**
```json
{
  "error": "Missing authentication token"
}
```

```json
{
  "error": "Invalid authentication token"
}
```

**Status Code:** `500`
**Response Body (server side problems):**
```json
{
    "error": "Couldn't decode parameters",
}
```

```json
{
    "error": "There was a problem hashing your password",
}
```

```json
{
    "error": "Couldn't create user in the database",
}
```

### Endpoint: Login User

#### Description

#### HTTP Request
**Method:** `POST`

**URL:** `/api/login`

#### Headers
| Name           | Type   | Description                |
|----------------|--------|----------------------------|
| `Content-Type` | string | Must be `application/json` |

#### Request Body
Send a JSON object with the following fields:
| Field Name | Type   | Required | Description                            |
|------------|--------|----------|----------------------------------------|
| `Email`    | string | Yes      | The unique email address for the user. |
| `Password` | string | Yes      | The user's password.                   |

Example:
```json
{
  "email": "user@example.com",
  "password": "my-secure-password-123"
}
```

#### Response

##### Success
**Status Code:** `200`
**Response Body:**
```json
{
    "id": "efe3be18-0116-43cf-b775-f81136f04d41",
    "created_at": "2024-11-22T11:16:37.466073Z",
    "updated_at": "2024-11-22T11:16:37.466073Z",
    "email": "walt@breakingbad.com",
    "is_chirpy_red": false,
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHkiLCJzdWIiOiJlZmUzYmUxOC0wMTE2LTQzY2YtYjc3NS1mODExMzZmMDRkNDEiLCJleHAiOjE3MzIyOTU3OTcsImlhdCI6MTczMjI5MjE5N30.FTgu5-0TVTOMXT98eGZTBrUHUUdoflz38i3uoB3Vado",
    "refresh_token": "dabd654dd0e166a9c1d44e17b5cac6d56eb202e7f9a46aafa87ac9e5709b8724"
}
```

##### Failure
**Status Code:** `401`
**Response Body (Incorred email or password):**
```json
{
  "error": "Incorrect email or password"
}
```

**Status Code:** `500`
**Response Body (Problem generating or storing access or refresh token):**
```json
{
  "error": "Unable to complete login at this time. Please try again later."
}
```

### Endpoint: Refresh Access Token

#### Description

#### HTTP Request
**Method:** `POST`

**URL:** `/api/refresh`

#### Headers
| Name            | Type   | Description                            |
|-----------------|--------|----------------------------------------|
| `Authorization` | string | Bearer token for API access. Required. |

#### Request Body
No Request Body processed by this endpoint

#### Response

##### Success
**Status Code:** `200`
**Response Body:**
```json
{
  "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJjaGlycHkiLCJzdWIiOiI2MjE3N2Q3NC0wOTZhLTQ0YzQtYWZmNy04MzdlMGU0YWNmZDciLCJleHAiOjE3MzIzMDQwOTksImlhdCI6MTczMjMwMDQ5OX0.xwT3Y_7cbSTPj_qcUNdErLaMRcTaP0MvEVWtSs3O-DQ"
}
```

##### Failure
**Status Code:** `400`
**Response Body (Bad Request - Missing Authorization Header):**
```json
{
  "error": "Missing 'Authorizaton' header value"
}
```

**Status Code:** `401`
**Response Body (Invalid or expired refresh token):**
```json
{
  "error": "Invalid refresh token"
}
```

```json
{
  "error": "Expired refresh token"
}
```

**Status Code:** `500`
**Response Body:**
```json
{
  "error": "Unable to complete login at this time. Please try again later."
}
```

### Endpoint: Revoke Refresh Token

#### Description

#### HTTP Request
**Method:** `POST`

**URL:** `/api/revoke`

#### Headers
| Name            | Type   | Description                            |
|-----------------|--------|----------------------------------------|
| `Authorization` | string | Bearer token for API access. Required. |

#### Request Body
No Request Body processed by this endpoint

#### Response

##### Success
**Status Code:** `204`
**Response Body:**
No response body sent upon successful request.

##### Failure
**Status Code:** `400`
**Response Body (Bad Request - Missing Authorization Header):**
```json
{
  "error": "Missing 'Authorizaton' header value"
}
```

### Endpoint: Create a New Chirp

#### Description

#### HTTP Request
**Method:** `POST`

**URL:** `/api/chirps`

#### Headers
| Name            | Type   | Description                            |
|-----------------|--------|----------------------------------------|
| `Content-Type`  | string | Must be `application/json`             |
| `Authorization` | string | Bearer token for API access. Required. |

#### Request Body
Send a JSON object with the following fields:
| Field Name | Type   | Required | Description                    |
|------------|--------|----------|--------------------------------|
| Body       | String | Yes      | Text content of the new chirp. |

Example:
```json
{
    "body": "Hello World!"
}
```

#### Response

##### Success
**Status Code:** `201`
**Response Body:**
```json
{
  "id": "03947601-d13e-48f8-bd6c-5eb2864cc595",
  "created_at": "2024-11-22T13:34:59.737205Z",
  "updated_at": "2024-11-22T13:34:59.737205Z",
  "body": "Hello World!",
  "user_id": "62177d74-096a-44c4-aff7-837e0e4acfd7"
}
```

##### Failure
**Status Code:** `400`
**Response Body (Bad Request - Missing Authorization Header):**
```json
{
  "error": "Missing 'Authorizaton' header value"
}
```

**Status Code:** `401`
**Response Body (Invalid access token):**
```json
{
  "error": "Invalid access token"
}
```

**Status Code:** `400`
**Response Body (Bad Request - Chirp too long):**
```json
{
  "error": "Chirp is too long"
}
```

**Status Code:** `500`
**Response Body (Internal server error):**
```json
{
  "error": "Couldn't decode parameters"
}
```

```json
{
  "error": "Couldn't create chirp in the database"
}
```

### Endpoint: Get All Chirps

#### Description

#### HTTP Request
**Method:** `GET`

**URL:** `/api/chirps`

#### Query Parameters
| Name        | Type   | Required | Default | Description                                                    |
|-------------|--------|----------|---------|----------------------------------------------------------------|
| `author_id` | string | No       | `null`  | A user ID to filter chirps by author.                          |
| `sort`      | string | No       | `asc`   | The sort order for chirp listing. Valid values: `asc`, `desc`. |

#### Headers
No Headers processed by this endpoint

#### Request Body
No Request Body processed by this endpoint

#### Response

##### Success
**Status Code:** `200`
**Response Body:**
```json
[
    {
        "id": "0ba3a7b9-7386-42aa-94de-b41784c3d3a2",
        "created_at": "2024-11-22T11:16:37.523782Z",
        "updated_at": "2024-11-22T11:16:37.523782Z",
        "body": "Darn that fly, I just wanna cook",
        "user_id": "efe3be18-0116-43cf-b775-f81136f04d41"
    },
    {
        "id": "fe472099-4b92-4bc1-8f70-2a053ac6c80c",
        "created_at": "2024-11-22T11:16:37.52226Z",
        "updated_at": "2024-11-22T11:16:37.52226Z",
        "body": "Cmon Pinkman",
        "user_id": "efe3be18-0116-43cf-b775-f81136f04d41"
    },
    {
        "id": "b6831079-c4fb-4d52-88a6-cbf88c06009f",
        "created_at": "2024-11-22T11:16:37.521006Z",
        "updated_at": "2024-11-22T11:16:37.521006Z",
        "body": "Gale!",
        "user_id": "efe3be18-0116-43cf-b775-f81136f04d41"
    },
    {
        "id": "ad70682f-0adb-45fd-a550-9ab605508bbb",
        "created_at": "2024-11-22T11:16:37.519414Z",
        "updated_at": "2024-11-22T11:16:37.519414Z",
        "body": "I'm the one who knocks!",
        "user_id": "efe3be18-0116-43cf-b775-f81136f04d41"
    }
]
```

##### Failure
**Status Code:** `400`
**Response Body (Bad request when attempting to filter by invalid author ID):**
```json
{
    "error": "Invalid author ID"
}
```

**Status Code:** `500`
**Response Body (server error fetching chirps from the database):**
```json
{
    "error": "Couldn't get chirps from the database"
}
```

### Endpoint: Get Single Chirp

#### Description

#### HTTP Request
**Method:** `GET`

**URL:** `/api/chirps/{chirpID}`

#### Path Value Parameters
| Name      | Type   | Required | Description                   |
|-----------|--------|----------|-------------------------------|
| `chirpID` | string | Yes      | The chirp to get by Chirp ID. |

#### Headers
No Headers processed by this endpoint

#### Request Body
No Request Body processed by this endpoint

#### Response

##### Success
**Status Code:** `200`
**Response Body:**
```json
{
    "id": "0ba3a7b9-7386-42aa-94de-b41784c3d3a2",
    "created_at": "2024-11-22T11:16:37.523782Z",
    "updated_at": "2024-11-22T11:16:37.523782Z",
    "body": "Darn that fly, I just wanna cook",
    "user_id": "efe3be18-0116-43cf-b775-f81136f04d41"
}
```

##### Failure
**Status Code:** `400`
**Response Body (Bad request):**
```json
{
    "error": "Unable to decode requested chirp ID into valid UUID"
}
```

**Status Code:** `404`
**Response Body (Chirp not found in the database):**
```json
{
    "error": "Chirp not found"
}
```

### Endpoint: Delete a Chirp

#### Description

#### HTTP Request
**Method:** `DELETE`

**URL:** `/api/chirps/{chirpID}`

#### Path Value Parameters
| Name      | Type   | Required | Description                      |
|-----------|--------|----------|----------------------------------|
| `chirpID` | string | Yes      | The chirp to delete by Chirp ID. |


#### Headers
| Name            | Type   | Description                            |
|-----------------|--------|----------------------------------------|
| `Authorization` | string | Bearer token for API access. Required. |

#### Request Body
No Request Body processed by this endpoint

#### Response

##### Success
**Status Code:** `204`
**Response Body:**
No response body returned upon successful request.

##### Failure
**Status Code:** `401`
**Response Body (Missing or invalid access token):**
```json
{
  "error": "Missing authentication token"
}
```

```json
{
  "error": "Invalid authentication token"
}
```

**Status Code:** `400`
**Response Body:**
```json
{
    "error": "Unable to decode requested chirp ID into valid UUID"
}
```

**Status Code:** `404`
**Response Body (Chirp not found in the database):**
```json
{
    "error": "Chirp not found"
}
```

**Status Code:** `403`
**Response Body:**
```json
{
    "error": "User is not authorized to delete this chirp"
}
```

**Status Code:** `500`
**Response Body (Internal server error):**
```json
{
    "error": "There was a problem deleting chirp"
}
```

### Endpoint: Get Website Metrics

#### Description

#### HTTP Request
**Method:** `GET`

**URL:** `/admin/metrics`

#### Headers
No Request Headers processed by this endpoint

#### Request Body
No Request Body processed by this endpoint

#### Response

##### Success
**Status Code:** `200`
**Response Body:**
```html
<html>
  <body>
    <h1>Welcome, Chirpy Admin</h1>
    <p>Chirpy has been visited 10 times!</p>
  </body>
</html>
```

##### Failure
No Failure response currently provided by this endpoint

### Endpoint: Reset Database

#### Description

#### HTTP Request
**Method:** `POST`

**URL:** `/admin/reset`

#### Headers
No Request Headers processed by this endpoint

#### Request Body
No Request Body processed by this endpoint

#### Response

##### Success
**Status Code:** `200`
**Response Body:**
```
Page hit metric has been reset to 0. All users have been deleted from the database.
```

##### Failure
**Status Code:** `403`
**Response Body:**
```json
{
    "error": "Accessing this endpoint is forbidden"
}
```

**Status Code:** `500`
**Response Body:**
```json
{
    "error": "Couldn't delete users from database"
}
```