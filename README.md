# Gin + Go + Redis, Simple Session Auth (Training Project)
A small training project.  
Built with Go, Gin, and Redis.  
It shows a simple session-based authentication flow using JSON APIs.

## The app does:
- POST /login, accept JSON, create a session, and set the session_id cookie.
- GET /panel, a protected route. A middleware checks the session. If it is valid, it reads user_id and role from Redis and returns JSON.
- POST /panel/logout, remove the session in Redis and clear the cookie.

## Features
- Session stored in Redis (key: session:<id>).
- Cookie contains only session_id (HttpOnly).
- Middleware checks Redis and refreshes the TTL (idle timeout).
- Logout deletes the Redis key and removes the cookie.
- JSON API responses make testing and frontend work easier.

## Prerequisites
- Go (1.20+) installed.
- Redis server running (local or docker).
- Git to clone the project.
- Curl or Postman for testing.
