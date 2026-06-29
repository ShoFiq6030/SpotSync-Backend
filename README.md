# SpotSync Backend

SpotSync is a parking reservation backend service built with Go, Echo, and GORM.

## Live Demo

https://spotsync-backend.onrender.com

## ERD

https://drive.google.com/file/d/1lH5OWYTfIxREmSe76ZBDhqMKQneNda53/view?usp=sharing

## Features

- User registration and authentication
- Create and manage parking reservations
- Driver-only reservation listing
- Admin-only reservation listing
- Reservation cancellation with ownership enforcement
- Parking zone management

## Tech Stack

- Go 1.26
- Echo v5
- GORM v1.31
- PostgreSQL (configured via DSN)

## Quick Start

1. Clone the repository

```bash
git clone <repo-url>
cd SpotSync-backend
```

2. Set environment variables

```bash
PORT=4000
DSN=postgres://user:password@host:port/dbname
JWT_SECRET=your_secret_key
```

3. Run the app

```bash
go run ./cmd/main.go
```

## API Endpoints

- `GET /` - root route
- `GET /health` - basic health endpoint
- `POST /api/v1/users/register` - register user
- `POST /api/v1/users/login` - login user

### Parking Zone Endpoints

- `POST /api/v1/parking-zones` - create parking zone (admin only)
- `GET /api/v1/parking-zones` - list parking zones
- `GET /api/v1/parking-zones/:id` - get parking zone details

### Reservation Endpoints

- `POST /api/v1/reservations` - create reservation (authenticated)
- `GET /api/v1/reservations/my-reservations` - get user reservations (authenticated)
- `DELETE /api/v1/reservations/:id` - cancel reservation (authenticated)
- `GET /api/v1/reservations` - get all reservations (admin only)

## Environment Variables

- `PORT` - server port, default: `4000`
- `DSN` - Postgres connection string
- `JWT_SECRET` - secret for JWT signing

## Notes

- The app auto-migrates the database models on startup.
- Authentication is handled via JWT in `Authorization: Bearer <token>`.
