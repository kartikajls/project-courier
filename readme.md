# Courier Management System

## Deskripsi

<p style="text-align: justify;">Courier Management System merupakan aplikasi backend berbasis REST API yang dikembangkan menggunakan bahasa pemrograman Go (Golang) dengan pendekatan Monolithic Architecture dan Clean Architecture. Aplikasi ini dirancang untuk mengelola data kurir serta informasi operasional yang berkaitan dengan level pekerjaan dan status kurir. Sistem menyediakan fitur CRUD (Create, Read, Update, Delete) untuk pengelolaan data kurir, serta fitur pencarian, penyaringan, pengurutan, dan pagination untuk mempermudah pengelolaan data.
<br>
Note:
<br>
System ini dibuat untuk memenuhi challange dari Gradin Digital Agency yang beralamatkan di Jl. Taman Internasional B8 no. 33A, Citraland, Surabaya, Indonesia
</p>

## Fitur Utama

### 1. Manajemen Kurir (Courier Management)

- Menambahkan data kurir.
- Menampilkan seluruh data kurir.
- Menampilkan detail kurir berdasarkan ID.
- Memperbarui informasi kurir.
- Menghapus data kurir.
- Mencari kurir berdasarkan nama.
- Memfilter kurir berdasarkan level dan status.
- Mengurutkan data berdasarkan nama atau tanggal registrasi.
- Menampilkan data menggunakan pagination.

### 2. Manajemen Level (Level Management)

- Menampilkan seluruh level pekerjaan kurir.
- Menampilkan detail level berdasarkan ID.
- Data level hanya dapat dibaca (*Read Only*).

Daftar level pekerjaan:

1 - Petugas Sortir Barang
2 - Kurir Sepeda Motor
3 - Kurir Mobil
4 - Koordinator
5 - Supervisor

### 3. Manajemen Status (Status Management)

- Menampilkan seluruh status kurir.
- Menampilkan detail status berdasarkan tipe.
- Data status hanya dapat dibaca (*Read Only*).

Daftar status:

active
inactive

### 4. Dokumentasi API (Swagger)

- Menyediakan dokumentasi API secara interaktif menggunakan Swagger UI.
- Menampilkan informasi endpoint, parameter, request body, dan response.
- Memungkinkan pengujian API secara langsung melalui browser.

## Teknologi yang Digunakan

Teknologi: 
Bahasa Pemrograman : Golang
Framework : Echo v4
Database : MySQL
Dokumentasi API : Swagger

## Arsitektur Aplikasi

Aplikasi ini menerapkan Monolithic Architecture dengan pendekatan Clean Architecture untuk memisahkan tanggung jawab setiap komponen agar kode lebih terstruktur, mudah dipelihara, dan dikembangkan.

Pembagian setiap layer:

- Entity: Mendefinisikan struktur data utama aplikasi.
- DTO (Data Transfer Object): Mengatur struktur data request dan response API.
- Repository: Menangani komunikasi dan operasi data dengan database.
- Usecase: Menangani business logic, validasi, dan proses aplikasi.
- Handler: Menangani HTTP request dan HTTP response.
- Routes: Mengatur dan mendaftarkan endpoint API.
- Config: Mengelola konfigurasi aplikasi dan koneksi database.
- Infrastructure: Menangani kebutuhan infrastruktur, termasuk koneksi MySQL.


Dokumentasi dapat diakses melalui:

```text
http://localhost:8080/swagger/index.html
```
