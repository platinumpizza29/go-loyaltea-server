# LoyalTea Shopping System

A comprehensive loyalty and shopping management system built with Go, Gin, and MongoDB. The system allows users to discover nearby shops, create shopping plans, track visits, and earn points.

## Features

### 🏪 Shop Management
- **Geolocation-based shop discovery** - Find nearby shops using MongoDB geospatial queries
- **Category filtering** - Browse shops by categories (coffee, clothing, grocery, etc.)
- **Brand-based grouping** - Find all shops from specific brands
- **Search functionality** - Search shops by name, brand, or address
- **Points system** - Each shop has a points value awarded to visitors

### 📋 Shopping Plans
- **Custom shopping plans** - Create personalized shopping itineraries
- **Visit tracking** - Mark shops as visited and automatically earn points
- **Progress monitoring** - Track completion rates and earned points
- **Plan management** - Add/remove shops from existing plans
- **Statistics** - View user shopping statistics and history

### 👤 User Management
- **User registration and authentication** - Secure user accounts with JWT
- **Points accumulation** - Automatic point tracking across all visits
- **Personal dashboard** - View total points, completed plans, and visit history



## Tech Stack

- **Backend**: Go 1.24.3
- **Web Framework**: Gin
- **Database**: MongoDB with geospatial indexing
- **Authentication**: JWT tokens
- **Password Hashing**: bcrypt

## Prerequisites

- Go 1.24.3 or higher
- MongoDB 4.4 or higher
- Git

## Installation

1. **Clone the repository**
   ```bash
   git clone <repository-url>
   cd loyaltea
   ```

2. **Install dependencies**
   ```bash
   go mod download
   ```

3. **Set up environment variables**
   ```bash
   cp .env.example .env
   # Edit .env with your configuration
   ```

4. **Environment variables** (`.env` file):
   ```env
   DATABASE_URL=mongodb://localhost:27017
   DBNAME=loyaltea
   JWT_SECRET=your-super-secret-jwt-key
   PORT=8080
   ```

5. **Initialize MongoDB database**
   ```bash
   # Start MongoDB service
   mongod
   
   # Run database initialization script
   mongo loyaltea < scripts/init_db.js
   
   # Optional: Add sample data for testing
   mongo loyaltea < scripts/sample_data.js
   ```

6. **Run the application**
   ```bash
   go run main.go
   ```

The server will start on `http://localhost:8080`

## Database Setup

The system requires specific MongoDB indexes for optimal performance:

### Required Indexes
- **Geospatial index** on shops.location for nearby searches
- **Compound indexes** on shops for filtering by category and brand
- **Text indexes** on shops for search functionality
- **User indexes** for authentication and points tracking
- **Plan indexes** for efficient plan queries

Run the initialization script to create all required indexes:
```bash
mongo loyaltea < scripts/init_db.js
```

## API Endpoints

### Authentication
- `POST /user/register` - Register new user
- `POST /user/login` - User login
- `GET /user/:id` - Get user profile
- `PUT /user/:id` - Update user profile

### Shop Discovery
- `GET /shops/nearby?lat=40.7128&lng=-74.0060&radius=1000` - Find nearby shops
- `GET /shops?category=coffee&search=query` - Get shops with filters
- `GET /shops/:id` - Get shop details
- `GET /shops/categories` - Get available categories
- `GET /shops/brand/:brand` - Get shops by brand

### Shopping Plans
- `POST /shopping-plans` - Create shopping plan
- `GET /shopping-plans/user/:id` - Get user's plans
- `GET /shopping-plans/user/:id/active` - Get active plans
- `GET /shopping-plans/user/:id/completed` - Get completed plans
- `PUT /shopping-plans/:id/visit/:shopId` - Mark shop as visited
- `POST /shopping-plans/:id/shops` - Add shop to plan
- `DELETE /shopping-plans/:id/shops/:shopId` - Remove shop from plan
- `GET /shopping-plans/:id/progress` - Get plan progress
- `GET /shopping-plans/user/:id/stats` - Get user statistics

### Admin Endpoints
- `POST /shops` - Create new shop
- `PUT /shops/:id` - Update shop
- `DELETE /shops/:id` - Delete shop (soft delete)

For complete API documentation, see [API_DOCUMENTATION.md](./API_DOCUMENTATION.md)

## Usage Examples

### Finding Nearby Shops
```bash
curl "http://localhost:8080/shops/nearby?lat=40.7128&lng=-74.0060&radius=1000"
```

### Creating a Shopping Plan
```bash
curl -X POST http://localhost:8080/shopping-plans \
  -H "Content-Type: application/json" \
  -H "user_id: user_123" \
  -d '{
    "name": "Weekend Shopping",
    "description": "Coffee and groceries",
    "stops": [
      {"shop_id": "shop_coffee_001"},
      {"shop_id": "shop_grocery_001"}
    ]
  }'
```

### Marking a Shop as Visited
```bash
curl -X PUT http://localhost:8080/shopping-plans/plan_123/visit/shop_coffee_001 \
  -H "user_id: user_123"
```

## Project Structure

```
loyaltea/
├── main.go                 # Application entry point
├── internal/
│   ├── models/            # Data models
│   │   ├── user.go
│   │   ├── shop.go
│   │   ├── shopping_plan.go

│   │   ├── offer.go
│   │   └── favorite.go
│   ├── db/                # Database layer
│   │   ├── db.go
│   │   ├── userDb.go
│   │   ├── shopDb.go
│   │   ├── shopping_plan_db.go
│   │   └── ...
│   ├── services/          # Business logic
│   │   ├── userService.go
│   │   ├── shopService.go
│   │   ├── shopping_plan_service.go
│   │   ├── errors.go
│   │   └── ...
│   ├── handlers/          # HTTP handlers
│   │   ├── userHandlers.go
│   │   ├── shopHandler.go
│   │   ├── shopping_plan_handler.go
│   │   └── ...
│   └── utils/             # Utility functions
├── scripts/
│   ├── init_db.js         # Database initialization
│   └── sample_data.js     # Sample data for testing
├── go.mod
├── go.sum
├── .gitignore
├── README.md
└── API_DOCUMENTATION.md
```

## Key Features Explained

### Geolocation Search
The system uses MongoDB's geospatial capabilities to find shops within a specified radius:
- Shops store location as GeoJSON Point coordinates `[longitude, latitude]`
- Uses `$near` operator with `$geometry` for proximity searches
- Supports radius filtering in meters

### Points System
- Each shop has a `points_value` indicating rewards for visiting
- Points are automatically awarded when users mark shops as visited
- User's `total_points` are updated in real-time
- Points can be subtracted if visits are undone

### Shopping Plan Lifecycle
1. **Creation**: User selects shops to visit
2. **Active**: Plan is incomplete, shops can be added/removed
3. **Visiting**: User marks shops as visited, earning points
4. **Completion**: All shops visited, plan marked as completed

### Legacy Compatibility
The system provides comprehensive shopping plan functionality with geospatial search capabilities.

## Development

### Running Tests
```bash
go test ./...
```

### Adding Sample Data
```bash
mongo loyaltea < scripts/sample_data.js
```

### Database Queries for Development
```javascript
// Find shops near coordinates
db.shops.find({
  "location": {
    "$near": {
      "$geometry": {"type": "Point", "coordinates": [-74.0060, 40.7128]},
      "$maxDistance": 1000
    }
  }
})

// Get user's active plans
db.shopping_plans.find({"user_id": "user_123", "is_completed": false})

// User points leaderboard
db.users.find({}).sort({"total_points": -1}).limit(10)
```

## Contributing

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## Error Handling

The API uses consistent error response format:
```json
{
  "error": "descriptive error message"
}
```

Common HTTP status codes:
- `200` - Success
- `201` - Created
- `400` - Bad Request
- `401` - Unauthorized
- `403` - Forbidden
- `404` - Not Found
- `500` - Internal Server Error

## Security Considerations

- Passwords are hashed using bcrypt
- JWT tokens for authentication
- User authorization checks on protected endpoints
- Input validation on all endpoints
- MongoDB injection prevention through proper query construction

## Performance Optimizations

- Database indexes for all query patterns
- Geospatial indexing for location searches
- Efficient aggregation pipelines for statistics
- Connection pooling for database access
- Pagination support for large result sets

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Support

For support, email support@loyaltea.com or create an issue in the repository.

## Roadmap

- [ ] Push notifications for nearby shops
- [ ] Social features (share plans, friend leaderboards)
- [ ] Rewards catalog for spending points
- [ ] Analytics dashboard for shop owners
- [ ] Mobile app integration
- [ ] Real-time location tracking
- [ ] Machine learning for shop recommendations