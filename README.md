🚀 Manage User By Go
---

## Step การใช้งาน

### 1. Start Services ด้วย Docker
```bash
docker-compose up -d
```
โดย service go จะใช้ port 8181 และ mongoDB port 27017 
mongoDB express 8081 สำหรับดู database แบบ GUI (port สามารถแก้ไขได้ที่ไฟล์ env/config.yaml และใน Dockerfile ,compose.yaml)

### 2. ตรวจสอบว่า Service พร้อมทำงาน
```bash
curl --location 'localhost:8181/health'
```
### 3. เพิ่ม user ผ่านการเรียก API
```bash
curl --location 'localhost:8181/user/register' \
--header 'Content-Type: application/json' \
--data-raw '{
"name": "",
"email": "",
"password": ""
}'
```

### 4. จากนั้น log in ผ่าน API ด้วย email และ password ที่ทำการ register ไว้  
```bash
curl --location 'localhost:8181/user/login' \
--header 'Content-Type: application/json' \
--data-raw '{
"email": "",
"password": ""
}'
```

### จะได้ token สำหรับ authorization

## API ที่ต้องการใช้ token มีดังนี้
### - Get List User
```bash
curl --location 'localhost:8181/user/list' \
--header 'Authorization: Bearer {token}'
```
### - Get User By ID
```bash
curl --location 'localhost:8181/user/list' \
--header 'Authorization: Bearer {token}'
```
### - Update User By ID
```bash
curl --location 'localhost:8181/user/update' \
--header 'Authorization: Bearer {token}' \
--header 'Content-Type: application/json' \
--data-raw '{
    "id": "",
     "name": "",
    "email": ""
}'
```
### - Delete User By ID
```bash
curl --location --request DELETE 'localhost:8181/user/delete/2ea81642-0446-4b2f-9681-035a06fbe864' \
--header 'Authorization: Bearer {token}'
```

