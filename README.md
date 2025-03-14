# MyWebArticleAPI

## Deskripsi
**MyWebArticleAPI** adalah sebuah *RESTful API* yang menyediakan layanan backend untuk sebuah website berbasis konten. API ini mendukung fitur autentikasi, manajemen pengguna, artikel, dan komentar dengan peran yang berbeda untuk setiap pengguna.

## ✨ Fitur Utama
- **🔑 Autentikasi & Otorisasi** menggunakan JWT.
- **👥 Manajemen Pengguna** (Admin dapat mengelola semua pengguna, user dapat mengedit profil sendiri).
- **📝 Manajemen Artikel** (Editor/Admin dapat menambah, mengedit, dan menghapus artikel).
- **💬 Manajemen Komentar** (Pengguna dapat menambah, mengedit, dan menghapus komentar pada artikel).

## 🛠 Teknologi yang Digunakan
- **🚀 Bahasa Pemrograman**: Golang
- **⚡ Framework**: Gin
- **💾 Database**: PostgreSQL
- **🔗 ORM**: GORM
- **🔒 Autentikasi**: JWT
- **📖 Dokumentasi API**: Swagger

## 🔧 Instalasi & Konfigurasi
1. **Clone repository**:
   ```sh
   git clone https://github.com/adamfarizi/web-article.git
   cd MyWebArticleAPI
   ```
2. **Install dependencies**:
   ```sh
   go mod tidy
   ```
3. **Setup Database**:
   - Pastikan PostgreSQL telah terinstal.
   - Jalankan skrip database:
     ```sh
     psql -U postgres -d mywebapi -f database/schema.sql
     psql -U postgres -d mywebapi -f database/data.sql
     ```
4. **Jalankan aplikasi**:
   ```sh
   go run .
   ```

## 🌐 Struktur API
### 1️⃣ Authentication
| Method | Endpoint      | Deskripsi                 |
|--------|-------------|---------------------------|
| **POST**   | `/auth/login` | 🔓 Login ke sistem          |
| **POST**   | `/auth/register` | 📝 Registrasi pengguna baru |

### 2️⃣ User Management
| Method | Endpoint          | Deskripsi             |
|--------|-------------------|-----------------------|
| **GET**    | `/users`          | 📋 List semua pengguna  |
| **GET**    | `/users/:id`      | 🔍 Detail pengguna      |
| **PUT**    | `/users/:id`       | ✏️ Update profil        |
| **DELETE** | `/users/:id`      | ❌ Hapus pengguna       |

### 3️⃣ Article Management
| Method | Endpoint         | Deskripsi               |
|--------|-----------------|-------------------------|
| **GET**    | `/article`      | 📄 List artikel           |
| **POST**   | `/article`      | ✍️ Tambah artikel baru    |
| **PUT**    | `/article/:id`  | 🔄 Update artikel         |
| **DELETE** | `/article/:id`  | 🗑 Hapus artikel          |

### 4️⃣ Comment Management
| Method | Endpoint             | Deskripsi              |
|--------|----------------------|------------------------|
| **POST**   | `/comment` | 💬 Tambah komentar  |
| **PUT**    | `/comment/:id`      | ✏️ Update komentar       |
| **DELETE** | `/comment/:id`      | ❌ Hapus komentar        |

## 📜 Dokumentasi API
API ini menggunakan **Swagger** untuk dokumentasi. Setelah aplikasi berjalan, dokumentasi dapat diakses di:
```
http://localhost:8080/docs
```

## 🤝 Kontribusi
Jika ingin berkontribusi dalam proyek ini, silakan *fork* repository, buat *branch* baru, lakukan perubahan, dan kirim *pull request*.

## 📜 Lisensi
Proyek ini dilisensikan di bawah **MIT License**.

