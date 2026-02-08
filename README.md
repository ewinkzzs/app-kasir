# 🧾 App Kasir (Point of Sale API)

Aplikasi **Kasir (Point of Sale)** berbasis **Go (Golang)** dengan arsitektur **Clean Architecture** sederhana. Project ini dirancang sebagai **REST API** untuk mengelola data **Produk** dan **Kategori**, serta sudah siap digunakan di environment **cloud** seperti **Railway** dengan database **Supabase (PostgreSQL)**.

---

## ✨ Fitur Utama

* 📦 Manajemen Produk (CRUD)
* 🗂️ Manajemen Kategori (CRUD)
* 🔍 Pencarian Produk berdasarkan Nama (`search by name`)
* 🧾 Transaksi / Checkout (multiple item)
* 📊 Sales Summary (hari ini & range tanggal)
* 🧱 Struktur Clean Architecture (Handler → Service → Repository)
* 🐘 PostgreSQL (Supabase)
* ☁️ Siap deploy ke Railway
* ⚙️ Konfigurasi via Environment Variable

---

## 🛠️ Tech Stack

* **Language**: Go (Golang)
* **Database**: PostgreSQL (Supabase)
* **Driver DB**: pgx
* **Config Management**: Viper
* **Deployment**: Railway

---

## 📁 Struktur Folder

```
app-kasir
├── config
│   ├── config.go                # Load environment config
│   └── database.go              # Database connection
├── handlers                     # HTTP handlers
│   ├── category_handler.go
│   ├── product_handler.go
│   ├── transaction_handler.go   # Checkout / transaksi
│   └── report_handler.go        # Sales report
├── models                       # Entity / Model
│   ├── category.go
│   ├── product.go
│   ├── transaction.go
│   └── report.go
├── repositories                 # Database access layer
│   ├── category_repository.go
│   ├── product_repository.go
│   ├── transaction_repository.go
│   └── report_repository.go
├── services                     # Business logic layer
│   ├── category_service.go
│   ├── product_service.go
│   ├── transaction_service.go
│   └── report_service.go
├── main.go                      # App entry point
├── go.mod
├── go.sum
├── dev.http                     # HTTP request (dev)
├── prod.http                    # HTTP request (prod)
└── README.md
```

---

## ⚙️ Environment Variables

Buat environment variable berikut:

```env
PORT=8080
DB_CONN=postgresql://USER:PASSWORD@HOST:6543/postgres?sslmode=require
```

> ⚠️ **Catatan penting**:
>
> * Gunakan **Supabase pooler (port 6543)**
> * Password **harus di-URL encode** jika ada karakter khusus

---

## ▶️ Menjalankan Project Secara Lokal

### 1️⃣ Clone Repository

```bash
git clone https://github.com/ewinkzzs/app-kasir.git
cd app-kasir
```

### 2️⃣ Install Dependency

```bash
go mod tidy
```

### 3️⃣ Buat file `.env`

```env
PORT=8080
DB_CONN=postgresql://USER:PASSWORD@HOST:6543/postgres?sslmode=require
```

### 4️⃣ Jalankan Aplikasi

```bash
go run main.go
```

Aplikasi akan berjalan di:

```
http://localhost:8080
```

---

## ☁️ Deploy ke Railway

1. Push project ke GitHub
2. Buat project baru di **Railway**
3. Hubungkan ke repository GitHub
4. Set **Environment Variables**:

   * `PORT`
   * `DB_CONN`
5. Deploy 🚀

Railway akan otomatis expose aplikasi ke public URL.

---

## 🔗 Contoh Endpoint

### Health Check

```
GET /
```

### Kategori

```
GET    /categories
POST   /categories
PUT    /categories/{id}
DELETE /categories/{id}
```

### Produk

```
GET    /products
GET    /products?name=indom   # search by name
POST   /products
PUT    /products/{id}
DELETE /products/{id}
```

### Transaksi / Checkout

```
POST /api/checkout
```

Request body:

```json
{
  "items": [
    { "product_id": 1, "quantity": 2 },
    { "product_id": 3, "quantity": 1 }
  ]
}
```

### Sales Report

**Hari ini**

```
GET /api/report
```

**Range tanggal (Optional Challenge)**

```
GET /api/report?start_date=2026-01-01&end_date=2026-02-01
```

---

## 🧪 Testing Endpoint

Gunakan file:

* `dev.http` (local)
* `prod.http` (production)

Atau tool seperti **Postman / Insomnia**.

---

## 📌 Catatan Arsitektur

Project ini menggunakan pola:

```
Handler → Service → Repository → Database
```

Tujuannya:

* Mudah dikembangkan
* Mudah di-maintain
* Mudah di-test

---

## 📄 License

MIT License

---

## 👨‍💻 Author

**Erwin Rianto**
GitHub: [https://github.com/ewinkzzs](https://github.com/ewinkzzs)

---

> ⭐ Jika project ini membantu, jangan lupa beri star di GitHub!
