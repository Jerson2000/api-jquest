# API Endpoints Documentation

Base URL: `http://localhost:8080` (or your deployed server domain).

---

## Authentication & Headers

Endpoints requiring authentication require a valid JSON Web Token (JWT) in the headers.

### Required Request Headers (Protected Endpoints)
```http
Authorization: Bearer <your_access_token>
```

### CSRF Protection (State-changing Web Requests)
For state-changing web requests (POST, PUT, PATCH, DELETE) without an `Authorization` header, CSRF validation is enforced.
1. Fetch the token from `/api/token` (which sets the `_forgery.anti` cookie and returns a JSON payload).
2. Attach the token in your requests via the header:
```http
X-CSRF-Token: <csrf_token_value>
```

---

## 1. General Config & Utility

### GET `/api/token`
- **Auth Required**: No (Public)
- **Description**: Returns CSRF token and sets the anti-forgery cookie.
- **Sample Request**:
  ```bash
  curl -X GET http://localhost:8080/api/token
  ```

### GET `/api/current`
- **Auth Required**: Yes
- **Description**: Retrieves detailed info of the currently logged-in user.
- **Sample Request**:
  ```bash
  curl -X GET http://localhost:8080/api/current \
       -H "Authorization: Bearer <your_token>"
  ```
- **Response Shape**:
  ```json
  {
    "id": 1,
    "name": "Jane Doe",
    "role": "candidate"
  }
  ```

---

## 2. Authentication (`/api/auth`)

### POST `/api/auth/signup`
- **Auth Required**: No (Public)
- **Request Body (`AuthSignupRequestDto`)**:
  ```json
  {
    "name": "Jane Doe",
    "email": "jane@example.com",
    "password": "password123"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/auth/signup \
       -H "Content-Type: application/json" \
       -d '{"name": "Jane Doe", "email": "jane@example.com", "password": "password123"}'
  ```

### POST `/api/auth/login`
- **Auth Required**: No (Public)
- **Request Body (`AuthLoginRequestDto`)**:
  ```json
  {
    "email": "jane@example.com",
    "password": "password123"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/auth/login \
       -H "Content-Type: application/json" \
       -d '{"email": "jane@example.com", "password": "password123"}'
  ```
- **Response Shape (`AuthResponseDto`)**:
  ```json
  {
    "status": 200,
    "data": {
      "token": "eyJhbGciOi...",
      "refreshToken": "eyJhbGciOi..."
    }
  }
  ```

### POST `/api/auth/refresh`
- **Auth Required**: Yes (JWT bearer token check + refresh token validation)
- **Request Body (`AuthRefreshRequestDto`)**:
  ```json
  {
    "refreshToken": "eyJhbGciOi..."
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/auth/refresh \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"refreshToken": "your_refresh_token"}'
  ```

---

## 3. Users (`/api/users`)

### GET `/api/users`
- **Auth Required**: Yes
- **Description**: Returns all users (cached for 1 minute).
- **Sample Request**:
  ```bash
  curl -X GET http://localhost:8080/api/users \
       -H "Authorization: Bearer <your_token>"
  ```

### GET `/api/users/:id`
- **Auth Required**: Yes
- **Description**: Returns user by ID (cached for 1 minute).
- **Sample Request**:
  ```bash
  curl -X GET http://localhost:8080/api/users/1 \
       -H "Authorization: Bearer <your_token>"
  ```

### POST `/api/users`
- **Auth Required**: Yes
- **Request Body (`UserCreateRequestDto`)**:
  ```json
  {
    "name": "Alex Smith",
    "email": "alex@example.com",
    "password": "securepwd123",
    "phone": "09123456789",
    "gender": "male"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/users \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"name": "Alex", "email": "alex@example.com", "password": "pwd", "gender": "male"}'
  ```

### PUT `/api/users/:id`
- **Auth Required**: Yes
- **Request Body (`UserUpdateRequestDto`)**:
  ```json
  {
    "name": "Alex Updated",
    "email": "alex_new@example.com",
    "phone": "09999999999",
    "gender": "male"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X PUT http://localhost:8080/api/users/1 \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"name": "Alex Updated"}'
  ```

### DELETE `/api/users/:id`
- **Auth Required**: Yes
- **Sample Request**:
  ```bash
  curl -X DELETE http://localhost:8080/api/users/1 \
       -H "Authorization: Bearer <your_token>"
  ```

---

## 4. Companies (`/api/companies`)

### GET `/api/companies`
- **Auth Required**: No (Public)
- **Description**: Retrieves list of all companies.
- **Sample Request**:
  ```bash
  curl -X GET http://localhost:8080/api/companies
  ```

### POST `/api/companies/apply`
- **Auth Required**: No (Public)
- **Description**: Special path for onboarding a recruiter along with their company.
- **Request Body (`CompanyApplyRequestDto`)**:
  ```json
  {
    "company": {
      "name": "Super Tech Inc",
      "industry": "Software",
      "website": "https://supertech.example.com",
      "location": "New York",
      "companySize": "10-50"
    },
    "user": {
      "name": "Jane Employer",
      "email": "jane@supertech.example.com",
      "password": "password123",
      "gender": "female"
    }
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/companies/apply \
       -H "Content-Type: application/json" \
       -d '{"company": {"name": "Tech Corp", "industry": "IT"}, "user": {"name": "Jane", "email": "jane@corp.com", "password": "pwd", "gender": "female"}}'
  ```

### POST `/api/companies`
- **Auth Required**: Yes
- **Request Body (`CompanyCreateRequestDto`)**:
  ```json
  {
    "name": "Innovative Startups",
    "industry": "Venture Capital",
    "website": "https://innovative.example.com",
    "location": "San Francisco",
    "companySize": "1-10"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/companies \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"name": "Innovative Startups", "industry": "VC"}'
  ```

---

## 5. Jobs (`/api/jobs`)

### GET `/api/jobs`
- **Auth Required**: No (Public)
- **Query Parameters**:
  - `page` (optional, default: `1`)
  - `limit` (optional, default: `10`)
- **Sample Request**:
  ```bash
  curl -X GET "http://localhost:8080/api/jobs?page=1&limit=10"
  ```

### GET `/api/jobs/:id`
- **Auth Required**: No (Public)
- **Sample Request**:
  ```bash
  curl -X GET http://localhost:8080/api/jobs/1
  ```

### POST `/api/jobs`
- **Auth Required**: Yes (Recruiter / Admin)
- **Request Body (`JobCreateJobRequestDto`)**:
  ```json
  {
    "companyId": 1,
    "title": "Backend Go Developer",
    "description": "Develop and maintain robust microservices using Go.",
    "location": "Remote",
    "jobType": "full-time",
    "experience": 3,
    "salaryMin": 80000,
    "salaryMax": 120000,
    "status": "open",
    "deadline": "2026-12-31T23:59:59Z"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/jobs \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"companyId": 1, "title": "Go Developer", "description": "Go lang", "location": "Remote", "jobType": "full-time", "experience": 3}'
  ```

### PUT `/api/jobs/:id`
- **Auth Required**: Yes (Recruiter / Admin)
- **Request Body**: Partial updates of fields specified in `JobCreateJobRequestDto`.
- **Sample Request**:
  ```bash
  curl -X PUT http://localhost:8080/api/jobs/1 \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"title": "Senior Go Developer"}'
  ```

### DELETE `/api/jobs/:id`
- **Auth Required**: Yes (Recruiter / Admin)
- **Sample Request**:
  ```bash
  curl -X DELETE http://localhost:8080/api/jobs/1 \
       -H "Authorization: Bearer <your_token>"
  ```

---

## 6. Applications (`/api/applications`)

### POST `/api/applications`
- **Auth Required**: Yes (Candidate)
- **Request Body (`ApplicationCreateRequestDto`)**:
  ```json
  {
    "jobId": 1
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/applications \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"jobId": 1}'
  ```

### GET `/api/applications/my`
- **Auth Required**: Yes (Candidate)
- **Query Parameters**:
  - `page` (optional, default: `1`)
  - `limit` (optional, default: `10`)
- **Sample Request**:
  ```bash
  curl -X GET "http://localhost:8080/api/applications/my?page=1&limit=10" \
       -H "Authorization: Bearer <your_token>"
  ```

### GET `/api/applications/job/:jobId`
- **Auth Required**: Yes (Recruiter / Admin)
- **Query Parameters**:
  - `page` (optional, default: `1`)
  - `limit` (optional, default: `10`)
- **Sample Request**:
  ```bash
  curl -X GET "http://localhost:8080/api/applications/job/1?page=1&limit=10" \
       -H "Authorization: Bearer <your_token>"
  ```

### PATCH `/api/applications/:id/status`
- **Auth Required**: Yes (Recruiter / Admin)
- **Request Body (`ApplicationUpdateStatusRequestDto`)**:
  - **Status Choices**: `pending`, `reviewing`, `interviewig`, `offered`, `rejected`, `accepted`
  ```json
  {
    "status": "interviewig"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X PATCH http://localhost:8080/api/applications/1/status \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"status": "interviewig"}'
  ```

---

## 7. Recruiters (`/api/recruiters`)

### GET `/api/recruiters`
- **Auth Required**: Yes
- **Sample Request**:
  ```bash
  curl -X GET http://localhost:8080/api/recruiters \
       -H "Authorization: Bearer <your_token>"
  ```

### GET `/api/recruiters/company/:id`
- **Auth Required**: Yes
- **Sample Request**:
  ```bash
  curl -X GET http://localhost:8080/api/recruiters/company/1 \
       -H "Authorization: Bearer <your_token>"
  ```

### GET `/api/recruiters/user/:userId`
- **Auth Required**: Yes
- **Sample Request**:
  ```bash
  curl -X GET http://localhost:8080/api/recruiters/user/1 \
       -H "Authorization: Bearer <your_token>"
  ```

### POST `/api/recruiters`
- **Auth Required**: Yes
- **Request Body (`RecruiterCreateRequestDto`)**:
  ```json
  {
    "companyId": 1,
    "position": "Lead Talent Acquisition"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/recruiters \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"companyId": 1, "position": "Recruitment Specialist"}'
  ```

---

## 8. Skills (`/api/skills`)

### GET `/api/skills`
- **Auth Required**: No (Public)
- **Description**: Returns all registered skills.
- **Sample Request**:
  ```bash
  curl -X GET http://localhost:8080/api/skills
  ```

### POST `/api/skills/candidate`
- **Auth Required**: Yes (Candidate)
- **Description**: Associate a skill to the logged-in candidate's profile.
- **Request Body (`SkillCreateRequestDto`)**:
  ```json
  {
    "name": "Go"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/skills/candidate \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"name": "Go"}'
  ```

### DELETE `/api/skills/candidate`
- **Auth Required**: Yes (Candidate)
- **Description**: Disassociate a skill from the logged-in candidate's profile.
- **Request Body (`SkillCreateRequestDto`)**:
  ```json
  {
    "name": "Go"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X DELETE http://localhost:8080/api/skills/candidate \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"name": "Go"}'
  ```

### POST `/api/skills/job/:jobId`
- **Auth Required**: Yes (Recruiter)
- **Description**: Add a skill requirements to a specific job post.
- **Request Body (`SkillCreateRequestDto`)**:
  ```json
  {
    "name": "Kubernetes"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/skills/job/1 \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"name": "Kubernetes"}'
  ```

### DELETE `/api/skills/job/:jobId`
- **Auth Required**: Yes (Recruiter)
- **Description**: Remove a skill requirement from a specific job post.
- **Request Body (`SkillCreateRequestDto`)**:
  ```json
  {
    "name": "Kubernetes"
  }
  ```
- **Sample Request**:
  ```bash
  curl -X DELETE http://localhost:8080/api/skills/job/1 \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"name": "Kubernetes"}'
  ```

---

## 9. Saved Jobs / Bookmarks (`/api/saved-jobs`)

### POST `/api/saved-jobs`
- **Auth Required**: Yes (Candidate)
- **Description**: Bookmark/save a job posting.
- **Request Body (`SavedJobCreateRequestDto`)**:
  ```json
  {
    "jobId": 1
  }
  ```
- **Sample Request**:
  ```bash
  curl -X POST http://localhost:8080/api/saved-jobs \
       -H "Authorization: Bearer <your_token>" \
       -H "Content-Type: application/json" \
       -d '{"jobId": 1}'
  ```

### DELETE `/api/saved-jobs/:jobId`
- **Auth Required**: Yes (Candidate)
- **Description**: Remove a saved job bookmark.
- **Sample Request**:
  ```bash
  curl -X DELETE http://localhost:8080/api/saved-jobs/1 \
       -H "Authorization: Bearer <your_token>"
  ```

### GET `/api/saved-jobs`
- **Auth Required**: Yes (Candidate)
- **Description**: Retrieve a candidate's saved job posts.
- **Query Parameters**:
  - `page` (optional, default: `1`)
  - `limit` (optional, default: `10`)
- **Sample Request**:
  ```bash
  curl -X GET "http://localhost:8080/api/saved-jobs?page=1&limit=10" \
       -H "Authorization: Bearer <your_token>"
  ```

