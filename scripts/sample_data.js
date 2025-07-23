// Sample Data Script for LoyalTea Shopping System
// Run this script in MongoDB shell to populate test data

// Connect to your database
// use loyaltea

print("Inserting sample data...");

// Clear existing data (optional - uncomment if you want to start fresh)
/*
db.shops.deleteMany({});
db.shopping_plans.deleteMany({});
db.users.updateMany({}, { $set: { total_points: 0 } });
print("Existing data cleared.");
*/

// Insert sample shops with diverse locations and categories
db.shops.insertMany([
    {
        "_id": "shop_coffee_001",
        "name": "Brew & Bean",
        "brand": "Artisan Coffee Co",
        "address": "123 Coffee Street, Downtown",
        "location": {
            "type": "Point",
            "coordinates": [-74.0060, 40.7128] // NYC - Lower Manhattan
        },
        "points_value": 15,
        "category": "coffee",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop_coffee_002",
        "name": "Morning Grind",
        "brand": "Artisan Coffee Co",
        "address": "456 Espresso Ave, Midtown",
        "location": {
            "type": "Point",
            "coordinates": [-73.9857, 40.7484] // NYC - Times Square area
        },
        "points_value": 12,
        "category": "coffee",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop_clothing_001",
        "name": "Urban Threads",
        "brand": "Fashion Forward",
        "address": "789 Fashion Blvd, SoHo",
        "location": {
            "type": "Point",
            "coordinates": [-74.0020, 40.7209] // NYC - SoHo
        },
        "points_value": 25,
        "category": "clothing",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop_clothing_002",
        "name": "Style Central",
        "brand": "TrendSetters",
        "address": "321 Designer Way, Chelsea",
        "location": {
            "type": "Point",
            "coordinates": [-74.0021, 40.7465] // NYC - Chelsea
        },
        "points_value": 30,
        "category": "clothing",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop_grocery_001",
        "name": "Fresh & Green Market",
        "brand": "Organic Plus",
        "address": "555 Market Street, Brooklyn",
        "location": {
            "type": "Point",
            "coordinates": [-73.9442, 40.6892] // Brooklyn
        },
        "points_value": 20,
        "category": "grocery",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop_grocery_002",
        "name": "Corner Deli",
        "brand": "QuickMart",
        "address": "888 Convenience Lane, Queens",
        "location": {
            "type": "Point",
            "coordinates": [-73.7781, 40.6411] // Queens
        },
        "points_value": 10,
        "category": "grocery",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop_electronics_001",
        "name": "Tech Paradise",
        "brand": "GadgetWorld",
        "address": "999 Electronics Plaza, Manhattan",
        "location": {
            "type": "Point",
            "coordinates": [-73.9776, 40.7831] // Upper West Side
        },
        "points_value": 35,
        "category": "electronics",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop_restaurant_001",
        "name": "Pasta Palace",
        "brand": "Italian Delights",
        "address": "111 Taste Street, Little Italy",
        "location": {
            "type": "Point",
            "coordinates": [-73.9973, 40.7188] // Little Italy
        },
        "points_value": 18,
        "category": "restaurant",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop_bookstore_001",
        "name": "Chapter & Verse",
        "brand": "Independent Books",
        "address": "222 Literary Lane, Village",
        "location": {
            "type": "Point",
            "coordinates": [-74.0059, 40.7282] // Greenwich Village
        },
        "points_value": 8,
        "category": "bookstore",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop_pharmacy_001",
        "name": "Health First Pharmacy",
        "brand": "MediCare Plus",
        "address": "333 Wellness Way, Midtown",
        "location": {
            "type": "Point",
            "coordinates": [-73.9851, 40.7580] // Midtown West
        },
        "points_value": 12,
        "category": "pharmacy",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop_fitness_001",
        "name": "PowerFit Gym",
        "brand": "FitLife",
        "address": "444 Fitness Ave, Brooklyn",
        "location": {
            "type": "Point",
            "coordinates": [-73.9903, 40.6943] // Brooklyn Heights
        },
        "points_value": 22,
        "category": "fitness",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "shop_beauty_001",
        "name": "Glamour Spa",
        "brand": "BeautyWorks",
        "address": "555 Beauty Boulevard, Manhattan",
        "location": {
            "type": "Point",
            "coordinates": [-73.9735, 40.7614] // Upper East Side
        },
        "points_value": 28,
        "category": "beauty",
        "is_active": true,
        "created_at": new Date(),
        "updated_at": new Date()
    }
]);

print("Sample shops inserted successfully! Count: 12");

// Insert sample shopping plans (assuming some test users exist)
db.shopping_plans.insertMany([
    {
        "_id": "plan_001",
        "user_id": "test_user_001", // Replace with actual user ID
        "name": "Weekend Coffee & Shopping",
        "description": "Saturday morning coffee run and some shopping",
        "stops": [
            {
                "shop_id": "shop_coffee_001",
                "is_visited": true,
                "visited_at": new Date(Date.now() - 86400000), // 1 day ago
                "points_earned": 15
            },
            {
                "shop_id": "shop_clothing_001",
                "is_visited": false,
                "points_earned": 0
            },
            {
                "shop_id": "shop_bookstore_001",
                "is_visited": false,
                "points_earned": 0
            }
        ],
        "is_completed": false,
        "created_at": new Date(Date.now() - 172800000), // 2 days ago
        "updated_at": new Date(Date.now() - 86400000)   // 1 day ago
    },
    {
        "_id": "plan_002",
        "user_id": "test_user_001",
        "name": "Grocery Run",
        "description": "Weekly grocery shopping",
        "stops": [
            {
                "shop_id": "shop_grocery_001",
                "is_visited": true,
                "visited_at": new Date(Date.now() - 604800000), // 1 week ago
                "points_earned": 20
            },
            {
                "shop_id": "shop_pharmacy_001",
                "is_visited": true,
                "visited_at": new Date(Date.now() - 604800000), // 1 week ago
                "points_earned": 12
            }
        ],
        "is_completed": true,
        "created_at": new Date(Date.now() - 864000000), // 10 days ago
        "updated_at": new Date(Date.now() - 604800000)  // 1 week ago
    },
    {
        "_id": "plan_003",
        "user_id": "test_user_002", // Different user
        "name": "Tech & Fitness Day",
        "description": "Buy new gadgets and hit the gym",
        "stops": [
            {
                "shop_id": "shop_electronics_001",
                "is_visited": false,
                "points_earned": 0
            },
            {
                "shop_id": "shop_fitness_001",
                "is_visited": false,
                "points_earned": 0
            },
            {
                "shop_id": "shop_coffee_002",
                "is_visited": true,
                "visited_at": new Date(Date.now() - 7200000), // 2 hours ago
                "points_earned": 12
            }
        ],
        "is_completed": false,
        "created_at": new Date(),
        "updated_at": new Date()
    },
    {
        "_id": "plan_004",
        "user_id": "test_user_002",
        "name": "Pamper Day",
        "description": "Beauty and relaxation day",
        "stops": [
            {
                "shop_id": "shop_beauty_001",
                "is_visited": true,
                "visited_at": new Date(Date.now() - 259200000), // 3 days ago
                "points_earned": 28
            },
            {
                "shop_id": "shop_restaurant_001",
                "is_visited": true,
                "visited_at": new Date(Date.now() - 259200000), // 3 days ago
                "points_earned": 18
            }
        ],
        "is_completed": true,
        "created_at": new Date(Date.now() - 345600000), // 4 days ago
        "updated_at": new Date(Date.now() - 259200000)  // 3 days ago
    }
]);

print("Sample shopping plans inserted successfully! Count: 4");

// Update test users with points (if they exist)
// This assumes you have test users in your system
try {
    db.users.updateOne(
        { "_id": "test_user_001" },
        {
            $set: {
                "total_points": 47, // 15 + 20 + 12 from completed visits
                "updated_at": new Date()
            }
        }
    );
    print("Updated test_user_001 points");
} catch (e) {
    print("Could not update test_user_001 - user may not exist");
}

try {
    db.users.updateOne(
        { "_id": "test_user_002" },
        {
            $set: {
                "total_points": 58, // 12 + 28 + 18 from completed visits
                "updated_at": new Date()
            }
        }
    );
    print("Updated test_user_002 points");
} catch (e) {
    print("Could not update test_user_002 - user may not exist");
}

print("\n=== Sample Data Summary ===");
print("Shops inserted: " + db.shops.countDocuments({}));
print("Shopping plans inserted: " + db.shopping_plans.countDocuments({}));
print("Active shops: " + db.shops.countDocuments({"is_active": true}));
print("Completed plans: " + db.shopping_plans.countDocuments({"is_completed": true}));
print("Active plans: " + db.shopping_plans.countDocuments({"is_completed": false}));

print("\n=== Testing Queries ===");

// Test nearby search (around NYC coordinates)
var nearbyCount = db.shops.find({
    "location": {
        "$near": {
            "$geometry": {
                "type": "Point",
                "coordinates": [-74.0060, 40.7128]
            },
            "$maxDistance": 5000 // 5km radius
        }
    },
    "is_active": true
}).count();
print("Shops within 5km of test location: " + nearbyCount);

// Test category breakdown
var categories = db.shops.aggregate([
    { $match: { "is_active": true } },
    { $group: { _id: "$category", count: { $sum: 1 } } },
    { $sort: { count: -1 } }
]);
print("Shop categories:");
categories.forEach(function(cat) {
    print("  " + cat._id + ": " + cat.count);
});

print("\n=== Sample data insertion completed successfully! ===");
