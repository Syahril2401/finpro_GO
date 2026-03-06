# Dokumentasi API - Fitur Dasar (Basic Task)

Dokumentasi ini mencakup endpoint dasar untuk penambahan dan penampilan data sesuai dengan tugas yang diberikan.

## 1. Simpan Data (Tambah Data)

Endpoint ini memungkinkan penambahan data `Question` ke dalam database. Sesuai permintaan, endpoint ini mendukung metode **GET** (menggunakan query parameters) dan **POST** (menggunakan JSON body).

### A. Menggunakan GET (Quick Test)
**Link:** `/api/test/add`

**Parameter:**
- `text` (String, Required): Isi teks pertanyaan.
- `category` (String, Optional): Kategori pertanyaan.

**Contoh Pemanggilan:**
`http://localhost:8008/api/test/add?text=Apa hobi anda?&category=Umum`

---

### B. Menggunakan POST (Standard)
**Link:** `/api/test/add`

**Method:** `POST`

**Request Body (JSON):**
```json
{
  "question_text": "Apa cita-cita anda?",
  "category": "Karir",
  "scale_min": 1,
  "scale_max": 5
}
```

---

## 2. Lihat Isi Data (Select Data)

Endpoint ini digunakan untuk melihat semua data yang telah tersimpan dalam tabel pertanyaan.

**Link:** `/api/test/list`

**Method:** `GET`

**Response Example:**
```json
{
  "success": true,
  "message": "Success",
  "data": [
    {
      "question_id": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
      "question_text": "Apa hobi anda?",
      "category": "Umum",
      "scale_min": 1,
      "scale_max": 5,
      "is_active": true,
      "created_at": "2024-03-06T13:10:00Z"
    }
  ]
}
```

---

## Contoh Penggunaan Postman

### Screenshot 1: Tambah Data via GET
![Tambah Data GET](https://via.placeholder.com/800x400.png?text=Postman:+GET+/api/test/add?text=Contoh)
*(Gunakan URL: `http://localhost:8008/api/test/add?text=Hello World`)*

### Screenshot 2: Tambah Data via POST
![Tambah Data POST](https://via.placeholder.com/800x400.png?text=Postman:+POST+/api/test/add+with+JSON)
*(Pilih Tab Body -> raw -> JSON -> Masukkan JSON data)*

### Screenshot 3: Lihat Data
![Lihat Data](https://via.placeholder.com/800x400.png?text=Postman:+GET+/api/test/list)
*(Klik Send untuk melihat list data)*
