# Dokumentasi API

## URL Dasar
```
http://localhost:3000
```

## Endpoint

### 1. Dapatkan Saldo
- **Metode**: GET
- **URL**: `/balance/{userId}`
- **Deskripsi**: Mengambil saldo saat ini untuk pengguna tertentu.
- **Parameter Path**:
  - `userId` (integer): ID unik pengguna.
- **Respons Sukses**:
  - **Kode**: 200 OK
  - **Konten**:
    ```json
    {
        "success": true,
        "balance": 100.50
    }
    ```
- **Respons Error**:
  - **Kode**: 400 Bad Request
  - **Konten**:
    ```json
    {
        "success":"",
        "message": "ID pengguna tidak valid"
    }
    ```
  - **Kode**: 404 Not Found
  - **Konten**:
    ```json
    {
        "success":"",
        "message": "Pengguna tidak ditemukan"
    }
    ```

### 2. Tarik Dana
- **Metode**: POST
- **URL**: `/withdraw`
- **Deskripsi**: Menarik jumlah tertentu dari saldo dompet pengguna.
- **Body Permintaan**:
  ```json
  {
    "user_id": 1,
    "amount": 50.00
  }
  ```
- **Parameter Body**:
  - `user_id` (integer): ID unik pengguna.
  - `amount` (float): Jumlah yang akan ditarik (harus positif dan tidak melebihi saldo saat ini).
- **Respons Sukses**:
  - **Kode**: 200 OK
  - **Konten**:
    ```json
    {
        "success": true,
        "message": "Penarikan berhasil.  Sisa saldo Anda %.2f",
        "new_balance": 50.50
    }
    ```
- **Respons Error**:
  - **Kode**: 400 Bad Request
  - **Konten**:
    ```json
    {
        "success": false,
        "message": "Permintaan tidak valid"
    }
    ```
    atau
    ```json
    {
        "success": false,
        "message": "Saldo tidak cukup. Sisa saldo Anda %.2f"
    }
    ```
  - **Kode**: 404 Not Found
  - **Konten**:
    ```json
    {
        "success": false,
        "message": "Pengguna tidak ditemukan"
    }
    ```
  - **Kode**: 500 Internal Server Error
  - **Konten**:
    ```json
    {
        "success": false,
        "message": "Gagal memperbarui saldo"
    }
    ```
    atau
    ```json
    {
        "success": false,
        "message": "Gagal mencatat transaksi"
    }
    ```

## Catatan
- Semua nilai moneter dalam format desimal (misalnya, 100.50).
- API menggunakan JSON untuk body permintaan dan respons.
- Pastikan database MySQL berjalan dan detail koneksi di `.env` benar.
- Jalankan migrasi untuk membuat tabel yang diperlukan sebelum menggunakan API.
- Aplikasi secara otomatis melakukan seeding 10 pengguna dengan saldo awal saat startup, dan membersihkan semua data transaksi sebelumnya.