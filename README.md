# runsystem-test

## Penjelasan singkat
Test ini menggunakan bahasa Go dan menggunakan database SQLite. Test ini menggunakan struktur direktori "Clean Architecture" yang mana menggunakan direktori ```cmd``` untuk cli atau menjalankan server, lalu direktori ```internal``` core sistem yang dibuat seperti http handler, service, komunikasi dengan database, bussines logic, dll, ```storage``` tempat db path dari sqlite tesebut. Dalam test tersebut ada 2 endpoint, yaitu ```/users``` dengan method ```POST``` dan ```/users/:id``` dengan method ```GET```. Endpoint ```/users``` dengan method ```POST``` digunakan untuk membuat data user dengan bidang ```ID```, ```Nama```, dan ```Hobi```. Dalam endpoint ini ada validasi yang diberikan seperti ```ID``` dengan ```nama``` tertentu jika sudah terbuat akan menerima error ```409``` atau ```conflict error```, jika body json ada yang kosong maka akan menerima error ```422``` atau ```unprocessableentity```. Endpoint ```/users/:id``` dengan method ```GET``` digunakan untuk melihat data user sesuai id yang ingin dilihat. Dalam endpoint ini ada validasi yang diberikan, seperti jika ```ID``` yang dicari tidak ada maka akan menerima error ```404``` atau ```Not Found```.

## Alasan mengapa memilih bidang Hobi
Bagi saya hobi itu sebagai escape untuk melepaskan penat, hobi saya itu berenang dan sering kali pergi berenang jika saya merasa sedang penat atau stress, karena renang bagi saya itu bisa menenangkan pikiran dan bisa menstabilkan nafas. Kenapa bermakna penting bagi saya, karena hobi itu bukan hanya sekedat kegiatan yang dilakukan sesering mungkin tapi bisa menjadi jalan kabur untuk menenangkan pikiran.

## Tantangan - tantangan pada saat mengerjakan technical test
Untuk tantangan pada saat mengerjakan test ini ada beberapa, seperti jaringan internet atau WiFi sering kali tiba tiba mati atau lemot, masalah pada device yang mana ternyata harus install Mingw 64 padahal saya rasa sudah di install software tersebut, dan kucing saya yang sering kali mengganggu saya dalam mengerjakan test ini.

## Cara Menjalankan

### 1. Install Dependencies
```bash
go mod tidy
```

### 2. Setup Environment
Buat file `.env`:
```env
PORT=8080
DB_PATH=storage/database.db
```

### 3. Jalankan Aplikasi
```bash
go run cmd/main.go
```

Server akan berjalan di `http://localhost:8080`
### 1. Create User
**POST** `/users`

**Request Body:**
```json
{
  "name": "Yasid Al Mubarok",
  "hobbies": "swimming"
}
```

**Success Response (201):**
```json
{
  "status": 201,
  "message": "User created successfully",
  "data": {
    "id": 1,
    "name": "Yasid Al Mubarok",
    "hobbies": "swimming"
  }
}
```

### 2. Get User by ID
**GET** `/users/:id`

**Success Response (200):**
```json
{
  "status": 200,
  "message": "User retrieved successfully",
  "data": {
    "id": 1,
    "name": "Yasid Al Mubarok",
    "hobbies": "swimming"
  }
}
```