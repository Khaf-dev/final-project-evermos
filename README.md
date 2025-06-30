# FinalDBEcommerce - Backend API

## 📌 Deskripsi

FinalDBEcommerce adalah backend API untuk aplikasi e-commerce sederhana yang menyediakan fitur:

- Registrasi dan login user (dengan JWT)
- Manajemen toko dan produk
- Transaksi pembelian dengan validasi stok
- Hak akses berdasarkan role (user/admin)

Project ini dibangun menggunakan arsitektur clean dan modular.

## 🚀 Teknologi yang Digunakan

- Golang (Fiber Web Framework)
- MySQL
- GORM (ORM untuk Go)
- JWT (JSON Web Token)
- Postman (untuk testing)
- Docker (opsional)
- EMSIFA API (untuk wilayah Indonesia)

---

## 🗂️ Struktur Folder (Clean Architecture)

```
finaldbecommerce/
│
├── config/           # Konfigurasi aplikasi & database
├── domain/           # Interface dan struct utama
├── repository/       # Koneksi dan query database (GORM)
├── service/          # Logika bisnis
├── handler/          # HTTP controller / handler Fiber
├── middleware/       # JWT dan otorisasi
├── routes/           # Definisi routing
├── utils/            # Helper: hashing password, JWT, dsb
├── main.go           # Entry point aplikasi
```

---

## 🔑 Autentikasi & Middleware

Gunakan JWT Token di setiap request protected:

```
Authorization: Bearer <token>
```

Middleware `JWTProtected` akan menyimpan ke dalam context:

- `user_id`
- `role`
- `store_id` (jika bukan admin)

---

## 📌 Contoh Request Body JSON

### 🔐 Register

```
POST /api/v1/register
```

```json
{
  "name": "Rifyat",
  "email": "rifyat@example.com",
  "password": "secret123"
}
```

### 🔐 Login

```
POST /api/v1/login
```

```json
{
  "email": "rifyat@example.com",
  "password": "secret123"
}
```

### 📦 Tambah Produk

```
POST /api/v1/products
```

```json
{
  "name": "Baju Muslim",
  "description": "Baju lengan panjang bahan adem",
  "price": 150000,
  "stock": 20,
  "category_id": 1
}
```

### 🛒 Buat Transaksi

```
POST /api/v1/transactions
```

```json
{
  "address_id": 2,
  "items": [
    {
      "product_id": 1,
      "quantity": 2
    },
    {
      "product_id": 3,
      "quantity": 1
    }
  ]
}
```

### 🏠 Tambah Alamat

```
POST /api/v1/address
```

```json
{
  "label": "Rumah",
  "receiver_name": "Rifyat",
  "phone": "08123456789",
  "province_id": 31,
  "city_id": 3174,
  "detail": "Jl. Mawar No. 123, Kelapa Gading"
}
```

### 📚 Tambah Kategori (Admin Only)

```
POST /api/v1/categories
```

```json
{
  "name": "Fashion Muslim"
}
```

---

## 📊 Hak Akses Berdasarkan Role

| Role   | Akses                                                             |
| ------ | ----------------------------------------------------------------- |
| Admin  | Kelola semua user, toko, kategori, produk, transaksi, log         |
| Seller | Kelola produk toko sendiri, lihat transaksi ke tokonya, lihat log |
| User   | Lihat produk, buat transaksi, kelola alamat                       |

---

## 📮 Contoh Endpoint API

| Method | Endpoint              | Deskripsi                              |
| ------ | --------------------- | -------------------------------------- |
| POST   | `/register`           | Register user baru                     |
| POST   | `/login`              | Login dan ambil JWT                    |
| GET    | `/products`           | Lihat semua produk                     |
| POST   | `/products`           | Tambah produk (hanya seller)           |
| POST   | `/transactions`       | Buat transaksi baru                    |
| GET    | `/transactions/store` | Lihat transaksi ke toko (hanya seller) |
| GET    | `/logs/products`      | Lihat histori log produk               |
| GET    | `/admin/users`        | Lihat semua user (admin)               |
| POST   | `/address`            | Tambah alamat pengiriman               |
| GET    | `/provinces`          | Ambil daftar provinsi dari EMSIFA      |

---

## ⚙️ Konfigurasi Environment

Buat file `.env` di root project:

```
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=yourpassword
DB_NAME=finaldbcommerce
JWT_SECRET=your_jwt_secret
```

---

## 🧪 Postman Collection

Gunakan [Postman Collection Evermos](https://github.com/Fajar-Islami/go-example-cruid/blob/master/Rakamin%20Evermos%20Virtual%20Internship.postman_collection.json) sebagai acuan pengujian API. Import collection ke Postman dan sesuaikan URL dan token.

---

## 📬 Kontak

**Rifyat Chaesa Kaffarozi**  
Backend Developer | Final Project Rakamin x Evermos  
📧 Email: rifyatkaffa@gmail.com
🌐 GitHub: [https://github.com/Khaf-dev](github.com/rifyatkaffarozi)

---

> Terima kasih sudah membaca dokumentasi ini 🙏  
> Semoga project ini bermanfaat dan bisa jadi portofolio yang kuat!
