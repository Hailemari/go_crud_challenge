Here’s a detailed API documentation for the **Go CRUD Challenge** project. This file describes the endpoints, request/response formats, and error handling.

---

# API Documentation

**Project**: Go-CRUD-Challenge  
**Description**: A simple CRUD API for managing information about persons. The API uses an in-memory database and is accessible to frontend applications from different domains (CORS enabled).

---

## Base URL

All endpoints start with the base URL:

```
/person
```

---

## Endpoints

### 1. **GET /person**

**Description**: Fetch all persons stored in the database.

- **Method**: `GET`
- **URL**: `/person`

**Response**:

- **Status**: 200 OK
- **Body**:
  ```json
  [
    {
      "id": "string",
      "name": "string",
      "age": number,
      "hobbies": ["string"]
    }
  ]
  ```

### 2. **GET /person/{personId}**

**Description**: Fetch a specific person by their unique identifier.

- **Method**: `GET`
- **URL**: `/person/{personId}`

**Path Parameter**:

- `personId`: The unique identifier of the person (UUID).

**Response**:

- **Status**: 200 OK
- **Body**:

  ```json
  {
    "id": "string",
    "name": "string",
    "age": number,
    "hobbies": ["string"]
  }
  ```

- **Status**: 404 Not Found  
  **Body**:
  ```json
  {
    "error": "Person not found"
  }
  ```

### 3. **POST /person**

**Description**: Create a new person and store their information.

- **Method**: `POST`
- **URL**: `/person`

**Request Body**:

```json
{
  "name": "string",
  "age": number,
  "hobbies": ["string"]
}
```

- `name`: Name of the person (required).
- `age`: Age of the person (required).
- `hobbies`: List of hobbies (array, required, can be empty).

**Response**:

- **Status**: 201 Created
- **Body**:

  ```json
  {
    "id": "string",
    "name": "string",
    "age": number,
    "hobbies": ["string"]
  }
  ```

- **Status**: 400 Bad Request  
  **Body**:
  ```json
  {
    "error": "Invalid request data"
  }
  ```

### 4. **PUT /person/{personId}**

**Description**: Update an existing person's information.

- **Method**: `PUT`
- **URL**: `/person/{personId}`

**Path Parameter**:

- `personId`: The unique identifier of the person (UUID).

**Request Body**:

```json
{
  "name": "string",
  "age": number,
  "hobbies": ["string"]
}
```

**Response**:

- **Status**: 200 OK
- **Body**:

  ```json
  {
    "id": "string",
    "name": "string",
    "age": number,
    "hobbies": ["string"]
  }
  ```

- **Status**: 400 Bad Request  
  **Body**:

  ```json
  {
    "error": "Invalid request data"
  }
  ```

- **Status**: 404 Not Found  
  **Body**:
  ```json
  {
    "error": "Person not found"
  }
  ```

### 5. **DELETE /person/{personId}**

**Description**: Delete a specific person from the database.

- **Method**: `DELETE`
- **URL**: `/person/{personId}`

**Path Parameter**:

- `personId`: The unique identifier of the person (UUID).

**Response**:

- **Status**: 204 No Content (Successfully deleted)

- **Status**: 404 Not Found  
  **Body**:
  ```json
  {
    "error": "Person not found"
  }
  ```

---

## Error Handling

- **Invalid Endpoints**: If the client requests an endpoint that doesn’t exist, the API will respond with a `404 Not Found` error and a message indicating that the resource was not found.
- **Internal Server Errors**: Any unexpected server errors will respond with a `500 Internal Server Error` status and a message indicating the server error.

- **Cross-Origin Resource Sharing (CORS)**: CORS is enabled for this API to allow frontend applications on different domains to access it. This is configured to allow all origins (`*`) and methods (GET, POST, PUT, DELETE, OPTIONS).
