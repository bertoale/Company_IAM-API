# Refresh Token Implementation - Stateless

## Overview

Implementasi refresh token menggunakan **stateless JWT approach**. Tidak ada penyimpanan token di database, semua informasi tersimpan dalam token JWT itu sendiri. Ini memberikan scalability yang lebih baik dan mengurangi load pada database.

## Architecture

### Key Features

- **Stateless**: Tidak menyimpan token di database
- **JWT-based**: Menggunakan JWT untuk access token dan refresh token
- **Different Secrets**: Access token dan refresh token menggunakan secret key yang berbeda
- **Different Expiry**: Access token short-lived (30m), refresh token long-lived (7d)
- **Secure Cookies**: Token disimpan dalam HTTP-only cookies

### Token Types

#### 1. Access Token

- **Purpose**: Untuk autentikasi request API
- **Expiry**: 30 menit (konfigurasi: `JWT_EXPIRES_IN`)
- **Secret**: `JWT_SECRET`
- **Contains**: User ID, Roles, Permissions, Applications

#### 2. Refresh Token

- **Purpose**: Untuk mendapatkan access token baru
- **Expiry**: 7 hari (konfigurasi: `REFRESH_TOKEN_EXPIRES_IN`)
- **Secret**: `REFRESH_TOKEN_SECRET`
- **Contains**: User ID, Roles, Permissions, Applications

## Configuration

### Environment Variables (.env)

```env
JWT_SECRET=your_super_secret_jwt_key_sosmed_app_2025
JWT_EXPIRES_IN=30m
REFRESH_TOKEN_SECRET=your_super_secret_refresh_token_key_sosmed_app_2025
REFRESH_TOKEN_EXPIRES_IN=7d
```

## API Endpoints

### 1. Login

Generate access token dan refresh token saat login.

**Endpoint**: `POST /api/auth/login`

**Request Body**:

```json
{
  "identifier": "admin@example.com",
  "password": "your_password"
}
```

**Response** (200 OK):

```json
{
  "status": "success",
  "message": "Login successful",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "user": {
      "id": 1,
      "username": "admin",
      "email": "admin@example.com",
      "full_name": "Administrator"
    }
  }
}
```

**Cookies Set**:

- `token`: Access token (HttpOnly, 30 minutes)
- `refresh_token`: Refresh token (HttpOnly, 7 days)

---

### 2. Refresh Token

Generate new access token dan refresh token menggunakan refresh token yang lama.

**Endpoint**: `POST /api/auth/refresh-token`

**Request Body**:

```json
{
  "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
}
```

**Response** (200 OK):

```json
{
  "status": "success",
  "message": "Token refreshed successfully",
  "data": {
    "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9...",
    "refresh_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9..."
  }
}
```

**Response** (401 Unauthorized):

```json
{
  "status": "error",
  "message": "Invalid or expired refresh token: ..."
}
```

**Cookies Set**:

- `token`: New access token (HttpOnly, 30 minutes)
- `refresh_token`: New refresh token (HttpOnly, 7 days)

## Flow Diagram

```
┌─────────┐                 ┌─────────┐
│ Client  │                 │ Server  │
└────┬────┘                 └────┬────┘
     │                           │
     │ 1. POST /login            │
     │ (username + password)     │
     ├──────────────────────────>│
     │                           │
     │ 2. Access Token (30m)     │
     │    Refresh Token (7d)     │
     │<──────────────────────────┤
     │                           │
     │ 3. API Request            │
     │ (with Access Token)       │
     ├──────────────────────────>│
     │                           │
     │ 4. Response               │
     │<──────────────────────────┤
     │                           │
     │ ... (30 minutes later)    │
     │                           │
     │ 5. API Request            │
     │ (Access Token expired)    │
     ├──────────────────────────>│
     │                           │
     │ 6. 401 Unauthorized       │
     │<──────────────────────────┤
     │                           │
     │ 7. POST /refresh-token    │
     │ (with Refresh Token)      │
     ├──────────────────────────>│
     │                           │
     │ 8. New Access Token (30m) │
     │    New Refresh Token (7d) │
     │<──────────────────────────┤
     │                           │
     │ 9. Retry API Request      │
     │ (with new Access Token)   │
     ├──────────────────────────>│
     │                           │
     │ 10. Response              │
     │<──────────────────────────┤
     │                           │
```

## Client Implementation Guide

### JavaScript/TypeScript Example

```typescript
// api.config.ts
import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:5000/api",
  withCredentials: true, // Important: untuk mengirim cookies
});

let isRefreshing = false;
let failedQueue: any[] = [];

const processQueue = (error: any, token: string | null = null) => {
  failedQueue.forEach((prom) => {
    if (error) {
      prom.reject(error);
    } else {
      prom.resolve(token);
    }
  });
  failedQueue = [];
};

// Response interceptor untuk handle token refresh
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    // Jika error 401 dan belum retry
    if (error.response?.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        // Jika sedang refresh, tambahkan ke queue
        return new Promise((resolve, reject) => {
          failedQueue.push({ resolve, reject });
        })
          .then((token) => {
            return api(originalRequest);
          })
          .catch((err) => {
            return Promise.reject(err);
          });
      }

      originalRequest._retry = true;
      isRefreshing = true;

      const refreshToken = localStorage.getItem("refresh_token");

      if (!refreshToken) {
        // Redirect ke login jika tidak ada refresh token
        window.location.href = "/login";
        return Promise.reject(error);
      }

      try {
        // Call refresh token endpoint
        const response = await api.post("/auth/refresh-token", {
          refresh_token: refreshToken,
        });

        const { token, refresh_token } = response.data.data;

        // Simpan token baru
        localStorage.setItem("token", token);
        localStorage.setItem("refresh_token", refresh_token);

        // Process semua request yang menunggu
        processQueue(null, token);

        isRefreshing = false;

        // Retry original request
        return api(originalRequest);
      } catch (err) {
        processQueue(err, null);
        isRefreshing = false;

        // Redirect ke login jika refresh token gagal
        localStorage.removeItem("token");
        localStorage.removeItem("refresh_token");
        window.location.href = "/login";

        return Promise.reject(err);
      }
    }

    return Promise.reject(error);
  }
);

export default api;
```

### Usage Example

```typescript
// auth.service.ts
import api from "./api.config";

export const login = async (identifier: string, password: string) => {
  const response = await api.post("/auth/login", {
    identifier,
    password,
  });

  // Simpan token (cookies sudah di-set otomatis oleh server)
  const { token, refresh_token } = response.data.data;
  localStorage.setItem("token", token);
  localStorage.setItem("refresh_token", refresh_token);

  return response.data;
};

// user.service.ts
import api from "./api.config";

export const getUsers = async () => {
  // Token akan di-refresh otomatis jika expired
  const response = await api.get("/users");
  return response.data;
};
```

## Security Considerations

### 1. **Different Secrets**

Access token dan refresh token menggunakan secret yang berbeda. Jika access token secret ter-compromise, refresh token masih aman.

### 2. **HttpOnly Cookies**

Token disimpan dalam HttpOnly cookies untuk mencegah XSS attacks.

### 3. **Short-lived Access Token**

Access token memiliki lifetime yang pendek (30 menit) untuk mengurangi window of opportunity jika token dicuri.

### 4. **Token Rotation**

Setiap kali refresh token digunakan, token baru akan di-generate. Ini mencegah replay attacks.

### 5. **Fresh User Data**

Saat refresh token, data user, roles, permissions, dan applications diambil fresh dari database untuk memastikan data terbaru.

## Testing

### Test Login

```bash
curl -X POST http://localhost:5000/api/auth/login \
  -H "Content-Type: application/json" \
  -d '{
    "identifier": "admin@example.com",
    "password": "admin123"
  }'
```

### Test Refresh Token

```bash
curl -X POST http://localhost:5000/api/auth/refresh-token \
  -H "Content-Type: application/json" \
  -d '{
    "refresh_token": "YOUR_REFRESH_TOKEN_HERE"
  }'
```

## Migration from Stateful to Stateless

Jika sebelumnya menggunakan stateful refresh token (menyimpan di database), tidak perlu migration khusus:

1. Deploy kode baru
2. User yang sudah login akan tetap bisa menggunakan access token mereka
3. Saat access token expire, mereka akan perlu login ulang
4. Login baru akan menggunakan stateless refresh token

## Troubleshooting

### Issue: "Invalid refresh token"

- **Cause**: Token expired atau invalid
- **Solution**: User perlu login ulang

### Issue: "Failed to get user"

- **Cause**: User ID di token tidak ditemukan di database
- **Solution**: User mungkin sudah dihapus, perlu login ulang

### Issue: Token tidak di-set di cookies

- **Cause**: CORS configuration salah
- **Solution**: Pastikan `withCredentials: true` di client dan CORS diatur dengan benar di server

## Best Practices

1. ✅ Simpan refresh token di secure storage (HttpOnly cookies preferred)
2. ✅ Jangan simpan refresh token di localStorage jika tidak perlu
3. ✅ Implement automatic token refresh di client
4. ✅ Handle refresh token failure dengan redirect ke login
5. ✅ Gunakan HTTPS di production
6. ✅ Rotate refresh token setiap kali digunakan
7. ✅ Set expiry yang sesuai dengan kebutuhan security
