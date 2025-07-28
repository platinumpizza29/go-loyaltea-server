// MongoDB Database Initialization Script
// Run this script in MongoDB shell or MongoDB Compass

// Connect to your database (replace 'loyaltea' with your actual database name)
// use loyaltea

// Create geospatial index for shops collection to enable location-based queries
db.shops.createIndex({ location: "2dsphere" });

// Create compound index for shops filtering
db.shops.createIndex({ category: 1, is_active: 1 });
db.shops.createIndex({ brand: 1, is_active: 1 });
db.shops.createIndex({ is_active: 1, created_at: -1 });

// Create text index for shop search functionality
db.shops.createIndex({
  name: "text",
  brand: "text",
  address: "text",
});

// Create indexes for shopping_plans collection
db.shopping_plans.createIndex({ user_id: 1 });
db.shopping_plans.createIndex({ user_id: 1, is_completed: 1 });
db.shopping_plans.createIndex({ user_id: 1, created_at: -1 });

// Create compound index for plan stops
db.shopping_plans.createIndex({ "stops.shop_id": 1 });
db.shopping_plans.createIndex({ "stops.is_visited": 1 });

// Create indexes for users collection
db.users.createIndex({ email: 1 }, { unique: true });
db.users.createIndex({ total_points: -1 });

// Create indexes for existing collections if they don't exist
db.offers.createIndex({ sender_email: 1 });
db.offers.createIndex({ brand: 1 });
db.offers.createIndex({ createdAt: -1 });

db.favorites.createIndex({ user_id: 1 });
db.favorites.createIndex({ offer_id: 1 });
db.favorites.createIndex({ user_id: 1, offer_id: 1 }, { unique: true });

print("Database indexes created successfully!");

// Sample data insertion (optional - uncomment to add sample data)
/*
// Insert sample shops
db.shops.insertMany([
    {
        "_id": "shop1",
        "name": "Downtown Coffee",
        "brand": "Coffee Co",
        "address": "123 Main St, Downtown",
        "location": {
            "type": "Point",
            "coordinates": [-74.0060, 40.7128] // [longitude, latitude] - NYC coordinates
        },
        "points_value": 15,
        "category": "coffee",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop2",
        "name": "Fashion Hub",
        "brand": "StyleMart",
        "address": "456 Fashion Ave, Midtown",
        "location": {
            "type": "Point",
            "coordinates": [-74.0070, 40.7130]
        },
        "points_value": 25,
        "category": "clothing",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop3",
        "name": "Fresh Market",
        "brand": "GreenGrocer",
        "address": "789 Market St, Brooklyn",
        "location": {
            "type": "Point",
            "coordinates": [-73.9442, 40.6892]
        },
        "points_value": 20,
        "category": "grocery",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop4",
        "name": "Tech Store",
        "brand": "ElectroMart",
        "address": "321 Tech Plaza, Queens",
        "location": {
            "type": "Point",
            "coordinates": [-73.7781, 40.6411]
        },
        "points_value": 30,
        "category": "electronics",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop5",
        "name": "Book Corner",
        "brand": "ReadMore",
        "address": "555 Library Ave, Manhattan",
        "location": {
            "type": "Point",
            "coordinates": [-73.9857, 40.7484]
        },
        "points_value": 10,
        "category": "bookstore",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    }
]);

print("Sample shops inserted successfully!");
*/

// Validation queries to test the setup
print("\n=== Testing Indexes ===");

// Test geospatial query
print("Testing nearby shops query...");
var nearbyShops = db.shops
  .find({
    location: {
      $near: {
        $geometry: {
          type: "Point",
          coordinates: [-74.006, 40.7128],
        },
        $maxDistance: 1000,
      },
    },
    is_active: true,
  })
  .limit(5);

print(
  "Nearby shops query executed successfully. Count: " + nearbyShops.count(),
);

// Test category filtering
print("Testing category filtering...");
var coffeeShops = db.shops.find({
  category: "coffee",
  is_active: true,
});
print(
  "Coffee shops query executed successfully. Count: " + coffeeShops.count(),
);

// Test text search
print("Testing text search...");
var searchResults = db.shops.find({
  $text: { $search: "coffee" },
  is_active: true,
});
print(
  "Text search query executed successfully. Count: " + searchResults.count(),
);

print("\n=== Database initialization completed! ===");
