# LoyalTea Shopping System API Documentation

## Overview

The LoyalTea Shopping System provides a comprehensive API for managing shops, shopping plans, and user points. Users can discover nearby shops, create shopping plans, track visits, and earn points.

## Base URL

```
http://localhost:8080
```

## Authentication

Most endpoints require a `user_id` header for authentication:

```
Headers:
user_id: <user_id_string>
```

---

## User Management

### Register User
- **POST** `/user/register`
- **Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword",
  "name": "John Doe"
}
```
- **Response:**
```json
{
  "id": "user_123",
  "email": "user@example.com",
  "name": "John Doe",
  "total_points": 0,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

### Login User
- **POST** `/user/login`
- **Body:**
```json
{
  "email": "user@example.com",
  "password": "securepassword"
}
```
- **Response:**
```json
{
  "token": "jwt_token_here",
  "user": {
    "id": "user_123",
    "email": "user@example.com",
    "name": "John Doe",
    "total_points": 150
  }
}
```

### Get User
- **GET** `/user/:id`
- **Headers:** `user_id: <user_id>`
- **Response:**
```json
{
  "id": "user_123",
  "email": "user@example.com",
  "name": "John Doe",
  "total_points": 150,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

---

## Shop Management

### Get Nearby Shops
- **GET** `/shops/nearby?lat=40.7128&lng=-74.0060&radius=1000`
- **Query Parameters:**
  - `lat` (required): Latitude
  - `lng` (required): Longitude  
  - `radius` (optional): Search radius in meters (default: 1000)
- **Response:**
```json
{
  "shops": [
    {
      "id": "shop_123",
      "name": "Downtown Coffee",
      "brand": "Coffee Co",
      "address": "123 Main St, Downtown",
      "location": {
        "type": "Point",
        "coordinates": [-74.0060, 40.7128]
      },
      "points_value": 15,
      "category": "coffee",
      "is_active": true,
      "created_at": "2024-01-15T10:30:00Z",
      "updated_at": "2024-01-15T10:30:00Z"
    }
  ],
  "count": 1
}
```

### Get Shops
- **GET** `/shops?category=coffee&search=downtown`
- **Query Parameters:**
  - `category` (optional): Filter by category
  - `search` (optional): Search by name or brand
- **Response:**
```json
{
  "shops": [...],
  "count": 5
}
```

### Get Shop by ID
- **GET** `/shops/:id`
- **Response:**
```json
{
  "id": "shop_123",
  "name": "Downtown Coffee",
  "brand": "Coffee Co",
  "address": "123 Main St, Downtown",
  "location": {
    "type": "Point",
    "coordinates": [-74.0060, 40.7128]
  },
  "points_value": 15,
  "category": "coffee",
  "is_active": true,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

### Get Shop Categories
- **GET** `/shops/categories`
- **Response:**
```json
{
  "categories": [
    "coffee",
    "clothing",
    "grocery",
    "restaurant",
    "electronics",
    "pharmacy",
    "bookstore",
    "fitness",
    "beauty",
    "automotive"
  ]
}
```

### Get Shops by Brand
- **GET** `/shops/brand/:brand`
- **Response:**
```json
{
  "shops": [...],
  "count": 3,
  "brand": "Coffee Co"
}
```

### Create Shop (Admin Only)
- **POST** `/shops`
- **Body:**
```json
{
  "name": "New Coffee Shop",
  "brand": "Coffee Co",
  "address": "456 New St, City",
  "location": {
    "type": "Point",
    "coordinates": [-74.0070, 40.7130]
  },
  "points_value": 20,
  "category": "coffee"
}
```

### Update Shop (Admin Only)
- **PUT** `/shops/:id`
- **Body:** Same as create shop

### Delete Shop (Admin Only)
- **DELETE** `/shops/:id`
- **Response:**
```json
{
  "message": "Shop deleted successfully"
}
```

---

## Shopping Plans

### Create Shopping Plan
- **POST** `/shopping-plans`
- **Headers:** `user_id: <user_id>`
- **Body:**
```json
{
  "name": "Weekend Shopping",
  "description": "Coffee and groceries",
  "stops": [
    {
      "shop_id": "shop_123"
    },
    {
      "shop_id": "shop_456"
    }
  ]
}
```
- **Response:**
```json
{
  "id": "plan_123",
  "user_id": "user_123",
  "name": "Weekend Shopping",
  "description": "Coffee and groceries",
  "stops": [
    {
      "shop_id": "shop_123",
      "is_visited": false,
      "points_earned": 0
    },
    {
      "shop_id": "shop_456",
      "is_visited": false,
      "points_earned": 0
    }
  ],
  "is_completed": false,
  "created_at": "2024-01-15T10:30:00Z",
  "updated_at": "2024-01-15T10:30:00Z"
}
```

### Get Shopping Plan
- **GET** `/shopping-plans/:id`
- **Response:** Same as create response

### Get User's Shopping Plans
- **GET** `/shopping-plans/user/:id`
- **Headers:** `user_id: <user_id>`
- **Response:**
```json
{
  "plans": [...],
  "count": 3
}
```

### Get Active Plans
- **GET** `/shopping-plans/user/:id/active`
- **Headers:** `user_id: <user_id>`
- **Response:**
```json
{
  "plans": [...],
  "count": 2
}
```

### Get Completed Plans
- **GET** `/shopping-plans/user/:id/completed`
- **Headers:** `user_id: <user_id>`
- **Response:**
```json
{
  "plans": [...],
  "count": 5
}
```

### Update Shopping Plan
- **PUT** `/shopping-plans/:id`
- **Headers:** `user_id: <user_id>`
- **Body:** Same as create plan

### Delete Shopping Plan
- **DELETE** `/shopping-plans/:id`
- **Headers:** `user_id: <user_id>`
- **Response:**
```json
{
  "message": "Plan deleted successfully"
}
```

### Mark Shop as Visited
- **PUT** `/shopping-plans/:id/visit/:shopId`
- **Headers:** `user_id: <user_id>`
- **Response:**
```json
{
  "message": "Shop marked as visited and points awarded"
}
```

### Add Shop to Plan
- **POST** `/shopping-plans/:id/shops`
- **Headers:** `user_id: <user_id>`
- **Body:**
```json
{
  "shop_id": "shop_789"
}
```
- **Response:**
```json
{
  "message": "Shop added to plan successfully"
}
```

### Remove Shop from Plan
- **DELETE** `/shopping-plans/:id/shops/:shopId`
- **Headers:** `user_id: <user_id>`
- **Response:**
```json
{
  "message": "Shop removed from plan successfully"
}
```

### Get Plan Progress
- **GET** `/shopping-plans/:id/progress`
- **Response:**
```json
{
  "plan_id": "plan_123",
  "total_stops": 5,
  "visited_stops": 3,
  "total_points": 65,
  "is_completed": false,
  "completion_rate": 60.0
}
```

### Get User Statistics
- **GET** `/shopping-plans/user/:id/stats`
- **Headers:** `user_id: <user_id>`
- **Response:**
```json
{
  "user_id": "user_123",
  "total_plans": 10,
  "completed_plans": 7,
  "total_points": 485,
  "total_visits": 25
}
```

---

## Error Responses

### Common Error Codes

- **400 Bad Request**
```json
{
  "error": "Invalid request data"
}
```

- **401 Unauthorized**
```json
{
  "error": "Unauthorized"
}
```

- **403 Forbidden**
```json
{
  "error": "Forbidden"
}
```

- **404 Not Found**
```json
{
  "error": "Resource not found"
}
```

- **500 Internal Server Error**
```json
{
  "error": "Internal server error"
}
```

### Specific Business Logic Errors

- **Shop not in plan**
```json
{
  "error": "shop is not in the plan"
}
```

- **Shop already visited**
```json
{
  "error": "shop has already been visited"
}
```

- **Plan already completed**
```json
{
  "error": "plan is already completed"
}
```

- **Invalid location coordinates**
```json
{
  "error": "invalid location coordinates"
}
```

---

## Data Models

### User
```json
{
  "id": "string",
  "email": "string",
  "name": "string",
  "total_points": "number",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### Shop
```json
{
  "id": "string",
  "name": "string",
  "brand": "string",
  "address": "string",
  "location": {
    "type": "Point",
    "coordinates": [longitude, latitude]
  },
  "points_value": "number",
  "category": "string",
  "is_active": "boolean",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### Shopping Plan
```json
{
  "id": "string",
  "user_id": "string",
  "name": "string",
  "description": "string",
  "stops": [PlanStop],
  "is_completed": "boolean",
  "created_at": "datetime",
  "updated_at": "datetime"
}
```

### Plan Stop
```json
{
  "shop_id": "string",
  "is_visited": "boolean",
  "visited_at": "datetime",
  "points_earned": "number"
}
```

---

## Usage Examples

### Frontend Integration Example

```javascript
// Get nearby shops
const nearbyShops = async (lat, lng, radius = 1000) => {
  const response = await fetch(
    `/shops/nearby?lat=${lat}&lng=${lng}&radius=${radius}`
  );
  return response.json();
};

// Create shopping plan
const createPlan = async (userId, planData) => {
  const response = await fetch('/shopping-plans', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'user_id': userId
    },
    body: JSON.stringify(planData)
  });
  return response.json();
};

// Mark shop as visited
const markVisited = async (userId, planId, shopId) => {
  const response = await fetch(`/shopping-plans/${planId}/visit/${shopId}`, {
    method: 'PUT',
    headers: {
      'user_id': userId
    }
  });
  return response.json();
};
```

### Database Setup

Before using the API, run the database initialization script:

```bash
# MongoDB shell
mongo loyaltea < scripts/init_db.js

# Or add sample data
mongo loyaltea < scripts/sample_data.js
```

### Environment Variables

Create a `.env` file:

```
DATABASE_URL=mongodb://localhost:27017
DBNAME=loyaltea
JWT_SECRET=your-secret-key
PORT=8080
```

---

## Getting Started

1. **Setup Database:** Run the initialization script to create indexes
2. **Add Sample Data:** Optionally run the sample data script for testing
3. **Register Users:** Create user accounts via `/user/register`
4. **Create Shops:** Add shops via `/shops` (admin endpoint)
5. **Create Plans:** Users can create shopping plans with selected shops
6. **Track Visits:** Mark shops as visited to earn points
7. **View Progress:** Monitor plan completion and user statistics

The system automatically awards points when users visit shops and tracks their progress through shopping plans.