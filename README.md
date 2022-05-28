## RESTFULL API GIN MYSQL
Repository ini dibuat menggunakan [Golang](https://go.dev/), [Gin](https://gin-gonic.com/) dan [Gorm](https://gorm.io/). API ini akan melakukan CRUD terhadap data mahasiswa yang ada di database. Untuk mencoba API ini, silahkan untuk clone repository ini dengan cara 

#### 1. Clone Repository
```
git clone https://github.com/rezairfanwijaya/Restfull-API-Gin
```

#### 2. Masuk Ke Direktori
```
cd Restfull-API-Gin
```

#### 3. Membuat Database (Mysql)

```
CREATE DATABASE restfull_api_gin_go;
```

### LIST API

#### 1. Register
```
POST http://localhost:8080/register
```

#### 2. Login
```
POST http://localhost:8080/login
```

#### 3. Home
```
GET http://localhost:8080
```

#### 4. Show All Data
```
GET http://localhost:8080/mahasiswa
```

#### 5. Input Data
```
POST http://localhost:8080/mahasiswa
```

#### 6. Update Data
```
PUT http://localhost:8080/mahasiswa/nim
```

#### 6. Delete Data
```
DELETE http://localhost:8080/mahasiswa/nim
```

#### 7. Logout
```
GET http://localhost:8080/logout
```





