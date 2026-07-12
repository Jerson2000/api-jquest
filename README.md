# JQuest API – Enterprise Job Portal Backend

A high-performance, scalable REST API backend for a modern Job Portal platform. Built with **Golang (Gin)** and designed with **Clean Architecture**, this project aims to provide enterprise-grade features ranging from complex RBAC and ATS (Application Tracking System) capabilities to advanced search and monetization.

## 🚀 Project Status: Active Development

> **Current Focus**: Implementing Core Domain Logic (Jobs, Companies) and Advanced Search.

---

## 🗺️ Enterprise Roadmap & To-Do List

### Phase 1: Core Foundation & Security (✅ Mostly Complete)
Establish the secure bedrock of the application.
- [x] **Project Skeleton**: Standard Go Clean Architecture (Controllers, Services, Repositories, DTOs).
- [x] **Configuration Management**: Environment variables (Godotenv) & Type-safe config.
- [x] **Database Connectivity**: PostgreSQL connection with GORM & Auto-migrations.
- [x] **Authentication**:
    - [x] JWT Access & Refresh Token rotation.
    - [x] Login / Signup / Logout endpoints.
    - [x] Password Hashing (Argon2/Bcrypt) via `crypto` package.
- [x] **Authorization (RBAC)**:
    - [x] Role-Based Access Control using **Casbin**.
    - [x] Middleware for role verification (Admin, Employer, Candidate).
- [x] **Security Hardening**:
    - [x] CSRF Protection (Double Submit Cookie).
    - [x] Rate Limiting (Token Bucket via Redis/Memory).
- [ ] **Advanced Auth** (Planned):
    - [ ] OAuth2 / SSO (Google, LinkedIn, GitHub).
    - [ ] MFA (Multi-Factor Authentication).

### Phase 2: User & Organization Management (🚧 In Progress)
Manage functionality for different user personas.
- [x] **User Management**:
    - [x] CRUD Operations for Users.
    - [x] Profile Management.
- [ ] **Company Profiles (Employer)**:
    - [ ] Create/Manage Company Pages.
    - [ ] Team Member Management (Invite recruiters to company).
    - [ ] Company branding (Logo, Banner, verify status).
- [ ] **Candidate Profiles**:
    - [ ] Resume Parsed Data.
    - [ ] Skills tagging & Portfolio links.
    - [ ] "Open to Work" status.

### Phase 3: Job Board Mechanics
The core domain of the platform.
- [ ] **Job Posting Engine**:
    - [ ] CRUD for Job Posts.
    - [ ] Rich Text Description support.
    - [ ] Job Metadata (Salary range, Remote/On-site, Experience Level).
    - [ ] Expiration & Renewal logic.
- [ ] **Job Applications**:
    - [ ] Apply functionality.
    - [ ] Application lifecycle states (Applied, Screening, Interview, Offer, Rejected).
    - [ ] Resume/CV Uploads (AWS S3 / MinIO integration).

### Phase 4: Search & Discovery (High Performance)
- [ ] **Advanced Search**:
    - [ ] Full-text search (Postgres TSVector or Elasticsearch/Meilisearch).
    - [ ] Filters (Location, Salary, Tech Stack, Date Posted).
- [ ] **Recommendations**:
    - [ ] "Jobs you might like" based on User Skills vs Job Requirements.

### Phase 5: Communication & Notifications
- [ ] **Notification System**:
    - [ ] In-app Realtime Notifications (WebSockets/SSE).
    - [ ] Email Transactional Mails (SendGrid/SES) - Welcome, Password Reset, Application Status.
- [ ] **Messaging System**:
    - [ ] Direct Messaging between Candidate and Recruiter.

### Phase 6: Monetization & Admin
- [ ] **Payments & Subscriptions**:
    - [ ] Stripe/PayPal Integration.
    - [ ] Premium Job Posts (Featured, Pinned).
    - [ ] Employer Subscription Plans (SaaS model).
- [ ] **Admin Dashboard**:
    - [ ] Moderate User/Content.
    - [ ] Platform Analytics (Signups, Applications, Revenue).

### Phase 7: DevOps & Observability
- [x] **Hot Reload**: Configured with `Air`.
- [x] **Docker Support**: Basic Dockerfile (if available, otherwise To-Do).
- [ ] **CI/CD**: GitHub Actions for linting, testing, and building.
- [ ] **Observability**:
    - [ ] Structured Logging (Zap/ZeroLog).
    - [ ] Metrics (Prometheus) & Tracing (OpenTelemetry).
    - [ ] Error Tracking (Sentry).
- [ ] **Documentation**:
    - [ ] Swagger/OpenAPI Auto-generation (`swaggo`).
    - [ ] Setup Guides.

---

## 🛠 Tech Stack

| Component | Technology | Description |
| :--- | :--- | :--- |
| **Language** | Go (Golang) | High-performance, concurrent backend language. |
| **Framework** | Gin | Lightweight and fast Web Framework. |
| **Database** | PostgreSQL | Robust Relational Database. |
| **ORM** | GORM | Developer-friendly ORM for Go. |
| **Caching** | Redis | In-memory store for Caching & Rate Limiting. |
| **Auth** | JWT + Casbin | Secure Stateless Auth & Granular Permissions. |
| **Validation** | go-playground/validator | Struct and field validation. |

---

## ⚡ Getting Started

### Prerequisites
- [Go 1.22+](https://go.dev/dl/)
- [PostgreSQL](https://www.postgresql.org/)
- [Redis](https://redis.io/)

### Installation

1. **Clone the Repo**
   ```bash
   git clone https://github.com/jerson2000/jquest.git
   cd jquest
   ```

2. **Setup Environment**
   ```bash
   cp .env.example .env
   # Edit .env with your DB and Redis credentials
   ```

3. **Install Dependencies**
   ```bash
   go mod download
   ```

4. **Run Locally**
   ```bash
   # Standard Run
   go run main.go

   # Or with Air (Hot Reload)
   air
   ```

## 📡 API Overview

| Method | Endpoint | Access | Description |
| :--- | :--- | :--- | :--- |
| `POST` | `/api/auth/login` | Public | Login and receive Tokens |
| `POST` | `/api/auth/signup` | Public | Register new account |
| `POST` | `/api/auth/refresh` | Public | Refresh Access Token |
| `GET` | `/api/users` | Admin | List all users |
| `GET` | `/api/current` | Public | Get current logged-in user details |

> *Full API documentation will be available via Swagger UI (Coming Soon)*
