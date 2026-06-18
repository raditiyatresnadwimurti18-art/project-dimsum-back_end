# 🥟 Dimsum API Backend

Ini adalah sistem backend berbasis REST API untuk manajemen penjualan dimsum. Dibangun menggunakan arsitektur MVC dengan performa tinggi.

## 🛠️ Teknologi yang Digunakan
* **Bahasa:** Golang (Go)
* **Framework:** Gin Web Framework
* **Database:** Google Cloud Firestore (NoSQL)

## 🚀 Cara Menjalankan Server Lokal
1. Pastikan Anda memiliki file `firebase-key.json` di dalam folder utama.
2. Buka terminal dan jalankan perintah:
   ```bash
   go run main.go

Server akan berjalan di http://localhost:8080

📚 Dokumentasi Endpoint (API)
1. Tambah Menu Dimsum (Create)
URL: /api/dimsum

Method: POST

Body Request (JSON):

JSON
{
    "nama": "Siomay Ayam Premium",
    "harga": 15000,
    "stok": 50
}
2. Lihat Semua Menu (Read)
URL: /api/dimsum

Method: GET

Body Request: Kosong (None)

3. Ubah Data Menu (Update)
URL: /api/dimsum/:id

Method: PUT

Body Request (JSON): (Kirimkan seluruh data terbaru)

4. Hapus Menu (Delete)
URL: /api/dimsum/:id

Method: DELETE

Body Request: Kosong (None)