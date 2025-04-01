# Go Response

## Ringkasan
Go Response adalah aplikasi web sederhana yang dibangun dengan [Fiber](https://gofiber.io/), sebuah framework HTTP yang cepat untuk Go. Aplikasi ini menyediakan endpoint publik dan endpoint yang dilindungi, di mana akses ke endpoint yang dilindungi memerlukan token Bearer yang valid.

## Prasyarat
Sebelum menjalankan aplikasi ini, pastikan Anda telah menginstal:

- [Go](https://golang.org/dl/) (versi 1.22.2 atau lebih baru)
- Jika menggunakan Windows, jalankan aplikasi melalui WSL (Windows Subsystem for Linux) karena Makefile digunakan dalam proyek ini.

## Instalasi & Pengaturan
Klon repositori dan masuk ke direktori proyek:
```sh
git clone <repository-url>
cd go_response
```

Instal dependensi:
```sh
go mod tidy
```

## Menjalankan Aplikasi
Untuk memulai aplikasi, gunakan perintah berikut:
```sh
make run
```
Server akan berjalan di port 3000. Tampilan di terminal akan terlihat seperti ini.

![Screenshot 2025-04-01 165054](https://github.com/user-attachments/assets/3e1b8ac3-755b-4933-afaa-aa84360cbe02)


## Endpoint API

### Endpoint Publik
Endpoint ini dapat diakses tanpa autentikasi.
```sh
curl http://localhost:3000/public
```
#### Respons:
```json
{
  "message": "Ini adalah endpoint publik."
}
```

### Endpoint Terlindungi
Endpoint ini memerlukan token Bearer yang valid untuk diakses.

#### Contoh Token Valid:
```sh
curl -H "Authorization: Bearer custom_token_1234567890abcdefgHIJKLMNOPQRSTUVWXYZ" http://localhost:3000/protected
```
#### Respons:
```json
{
  "message": "Ini adalah endpoint yang dilindungi."
}
```

#### Contoh Token Tidak Valid:
```sh
curl -H "Authorization: Bearer custom_token_123457890abcdefgHIJKLMNOPQRSTUVWXYZ" http://localhost:3000/protected
```
#### Respons:
```json
{
  "error": "Unauthorized"
}
```

## Struktur Proyek
```
go_response/
├── main.go         # Titik masuk aplikasi
├── middleware/
│   └── auth.go    # Middleware autentikasi
├── Makefile       # Perintah untuk menjalankan dan membersihkan proyek
├── go.mod         # File modul Go
```

## Membersihkan Proyek
Untuk membersihkan file hasil kompilasi, jalankan:
```sh
make clean
```

