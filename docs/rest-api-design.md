# URL Shortener REST API Design

## Base URL

```
/api/v1/
```

---

## Authentication

- Use JWT tokens or API Keys for authentication.
- Some endpoints may be public (e.g., redirect), others require authentication (e.g., managing links).

---

## Endpoints

### 1. Shorten a URL

**POST** `/shorten`

- **Request Body:**
  ```json
  {
    "url": "https://example.com/some/long/path",
    "expires_at": "2025-12-31T23:59:59Z" // (Optional)
  }
  ```
- **Response:**
  ```json
  {
    "short_url": "https://short./abc123",
    "url": "https://example.com/some/long/path",
    "alias": "abc123",
    "expires_at": "2025-12-31T23:59:59Z"
  }
  ```

---

### 2. Redirect to Original URL

**GET** `/{alias}`

- **Redirects to:** Original URL.
- **Response:** HTTP 302/301 Redirect.

---

### 3. Get URL Details (Analytics)

**GET** `/urls/{alias}`

- **Response:**
  ```json
  {
    "short_url": "https://short.ly/abc123",
    "url": "https://example.com/some/long/path",
    "alias": "abc123",
    "created_at": "2025-05-22T03:12:10Z",
    "expires_at": "2025-12-31T23:59:59Z",
    "visit_count": 42,
    "last_visited_at": "2025-05-21T20:00:00Z"
  }
  ```

---

### 4. List User’s Shortened URLs

**GET** `/urls/`

- **Query Parameters:** `page`, `limit`, `sort`
- **Response:**
  ```json
  {
    "urls": [
      {
        /* ...each url object as above... */
      }
    ],
    "total": 100,
    "page": 1,
    "limit": 20
  }
  ```

---

### 5. Update a Shortened URL

**PATCH** `/urls/{alias}`

- **Request Body:**
  ```json
  {
    "expires_at": "2026-01-01T00:00:00Z" // (Optional)
  }
  ```
- **Response:** 204 No Content / 200 Updated Object

---

### 6. Delete a Shortened URL

**DELETE** `/urls/{alias}`

- **Response:** 204 No Content

---

### 7. URL Statistics (Advanced)

**GET** `/urls/{alias}/stats`

- **Response:**
  ```json
  {
    "alias": "abc123",
    "visit_count": 42,
    "unique_visitors": 33,
    "visits": [
      {
        "timestamp": "2025-05-21T20:00:00Z",
        "ip_address": "123.45.67.89",
        "referrer": "https://google.com"
      },
      ...
    ]
  }
  ```

---

## Error Handling

- **400 Bad Request:** Invalid input.
- **401 Unauthorized:** Authentication failure.
- **403 Forbidden:** No permission.
- **404 Not Found:** Resource doesn’t exist.
- **409 Conflict:** Custom alias already exists.
- **429 Too Many Requests:** Throttling.

---

## Optional: User Management (if supporting accounts)

- **POST** `/auth/signup`
- **POST** `/auth/login`
- **GET** `/user/profile`
- **PATCH** `/user/profile`
- **DELETE** `/user/account`

---

## Notes

- Use rate limiting to prevent abuse.
- Consider analytics privacy/GDPR.
- Support CORS for frontend integration.
- Use OpenAPI/Swagger for documentation.

---

## Example OpenAPI/Swagger Spec

> You can generate an OpenAPI spec from the above design for interactive docs and client generation.
