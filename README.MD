# Online Store Application

Ini adalah aplikasi Web API yang dibangun menggunakan bahasa pemrograman Golang. Aplikasi ini merupakan salah satu syarat yang harus di selesaikan untuk seleksi software engineer backend pada PT. Synapsis Sinergi Digital

## Fitur

- Login and Reqister User
- Create New Category and See Category By ID
- Create New Product and See List Product base on Category
- Add Product to Cart
- See List of Product in Cart
- Delete Product from Cart
- Checkout Product and make payment transaction using midtrans

## Cara menggunakan

### Prasyarat

Sebelum menggunakan aplikasi ini pastikan anda telah memenuhi prasyarat sebagai berikut:

- Go 1.18 atau yang lebih baru
- Docker 23.0.1 atau yang lebih baru
- Akun Midtrans

### Konfigurasi

Untuk Menggunakan Aplikasi ini, dapat melalui dua cara oleh karena itu anda perlu melakukan konfigurasi terlebih dahulu. Berikut ini adalah cara melakukan konfigurasi:

1. Clone repository `git@github.com:darul-eng/synapsis.git` ini ke mesin lokal anda.
2. Buat file .env dari duplikasi file .env_example. Sesuaikan konfigurasi dengan kebutuhan anda.
3. Pastikan docker compose telah terinstall. Kemudian jalankan perintah `docker compose up`
4. Jika terjadi server error. Hentikan container yang berjalan dengan cara `ctrl+c` kemudian jalankan kembali perintah `docker compose up` 

### Endpoints

Aplikasi ini memiliki endpoint sebagai berikut:

- `POST /api/register`: Register New User
- `POST /api/login`: Login User
- `POST /api/category`: Create New Category
- `GET /api/category/:categoryId`: Get Category by ID
- `POST /api/product`: Create New Product base on category
- `GET /api/product/:categoryId`: Get list Product base on category
- `POST /api/cart`: Add product to Cart
- `GET /api/cart`: Get list Product in Cart
- `DELETE /api/cart/:cartId`: Delete Product from Cart
- `POST /api/transaction`: Create New Transaction
- `POST /api/transaction/notificattion`: Receive payment status notification from Midtrans system 

### Dokumentasi API 

Untuk dokumentasi lengkap mengenai API ini, silahkan lihat `https://app.swaggerhub.com/apis/darul-eng/api-documentation/1.0.0`


 

