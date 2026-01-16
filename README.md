# API Dompet Digital

Aplikasi dompet digital sederhana yang dibangun dengan Golang, Fiber, dan MySQL.

## Fitur

- Tarik dana dari dompet
- Periksa saldo dompet
- Data pengguna dan saldo disimpan di database MySQL

## Prasyarat

- Go 1.21 atau lebih baru
- Server MySQL berjalan
- Buat database bernama `wallet`

## Pengaturan

1. Klon repositori
2. Instal dependensi:
   ```bash
   go mod tidy
   ```

3. Konfigurasi pengaturan database di file `.env`.

4. Jalankan migrasi secara manual atau gunakan alat migrasi seperti golang-migrate.

5. Jalankan aplikasi:
   ```bash
   go run main.go
   ```

Aplikasi akan secara otomatis melakukan seeding data pengguna (10 pengguna) dan membersihkan data transaksi sebelumnya saat startup.

Server akan dimulai di port 3000.

## Endpoint API

### Dapatkan Saldo
- **GET** `/balance/:userId`
- Mengembalikan saldo saat ini untuk pengguna.

### Tarik Dana
- **POST** `/withdraw`
- Body: `{"user_id": 1, "amount": 50.00}`
- Menarik jumlah yang ditentukan dari dompet pengguna jika saldo cukup.

Untuk dokumentasi API detail termasuk contoh permintaan/respons dan kode error, lihat [api_documentation.md](api_documentation.md).

## Skema Database

- **users**: id, name, email, balance
- **transactions**: id, user_id, amount, type, created_at

## Migrasi

File migrasi disediakan di folder `migrations/`. Anda dapat menggunakan alat seperti golang-migrate untuk menerapkannya.