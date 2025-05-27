db = db.getSiblingDB('user');

db.createCollection('users');

print('Database "user" and collection "users" created successfully with sample data!');
