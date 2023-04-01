### Book Rest API Final Challenge
---
### Daftar Isi

- [Tentang](#tentang)
- [Panduan Menjalankan](#panduan-menjalankan)
- [Struktur Project](#struktur-project)
- [Dokumentasi API](#dokumentasi-api)
- [Hasil Dijalankan](#hasil-dijalankan)
---

### Tentang
BookLib merupakan Rest API untuk mendapatkan data buku yang telah terdaftar disistem. Pada BookLib telah menggunakan CRUD (Create, Read, Update, Delete) untuk mengolah data buku.

----
### Panduan Menjalankan

Proses menjalankan project dapat dilakukan dengan 3 cara yaitu menjalankan file main secara langsung maupun menggunakan Makefile. Berikut merupakan penjelasan cara menjalankan proyek ini:
+ Cara Pertama
Untuk menjalankan secara langsung file main.go dapat menggunakan command sebagai berikut:
    ```bash
    go run ./app/main.go
    ```
+ Cara Kedua
Untuk menjalankan menggunakan Makefile dapat menggunakan command sebagai berikut:
    ```bash
    Make run
    ```
+ Cara Ketiga
Untuk menjalankan menggunakan nodemon dapat menggunakan command sebagai berikut:
    ```bash
    nodemon --exec go run ./app/main.go
    ```
    Atau dapat menggunakan Makefile yang telah menyediakan command nodemon:

    ```bash
    Make run-nodemon
    ```
    Untuk detail lengkap cara menjalankan menggunakan Makefile dapat dilihat difile Makefile.

----
### Struktur Project
Secara sederhana project ini dibuat dengan menggunakan struktur Clean Architecture yang dipopulerkan oleh Uncle Bob. Berikut merupakan struktur Clean Architecture versi Uncle Bob:


![image](https://user-images.githubusercontent.com/13291041/102681893-84326980-4208-11eb-8f84-2959e03b89d8.png)


Dari Struktur tersebut dilakukan penyesuaian, dikarenakan dalam Rest API masih menggunakan API yang sederhana, maka struktur project akan terlihat seperti berikut:
| Layer                | Directory      |
|----------------------|----------------|
| Frameworks & Drivers | Infrastructures|
| Interface            | Interfaces     |
| Usecases             | Usecases       |
| Entities             | Domain         |

Dengan detail direktori dan file yang menyusun project sebagai berikut:

```bash
├── app
│   └── main.go
├── infrastructures
│   ├── config
│   │   ├── entity.go
│   │   └── config.go
│   └── databases
│       └── postgre.go
├── modules
│   └── v1
│       └── book
│           ├── domain
│           │   └── entity.go
│           ├── interfaces
│           │   ├── controllers
│           │   │   ├── book_adapter.go
│           │   │   └── book_controller.go
│           │   ├── repositories
│           │   │   ├── book_adapter.go
│           │   │   └── book_repository.go
│           ├── routes
│           │   └── book_routes.go
│           └── usecases
│               ├── book_adapater.go
│               └── book_interactor.go
├── pkg
│   ├── api_response
│   │   ├── entity.go
│   │   └── response.go
│   ├── http-error
│   │   └── error.go
├── .env
└── Makefile

```
### Dokumentasi API
Berikut merupakan dokumentasi API yang telah dibuat:

| Method | Endpoint | Description |
|--------|----------|-------------|
| GET    | /books   | Get all books |
| GET    | /books/:id | Get book by id |
| POST   | /books   | Create new book |
| PUT    | /books/:id | Update book by id |
| DELETE | /books/:id | Delete book by id |

### Hasil Dijalankan
Berikut merupakan hasil dari project yang telah dijalankan:

+ Program saat berhasil dijalankan
    ![image](/images/RunningOutput.jpg)

+ Database saat berhasil dijalankan
    ![image](/images/RunningOutputDB.jpg)

+ Get All Books
    +  Berhasil mendapatkan data buku yang telah terdaftar
    ![image](/images/ResultSuccess/Result_Success_GetAllBooks.jpg)
    +  Gagal mendapatkan data buku karena tidak ada data buku yang terdaftar.
    ![image](/images/ResultFailed/Result_Failed_GetAllBook.jpg)
+ Get Book By Id
    + Berhasil mendapatkan data buku berdasarkan id yang telah terdaftar
    ![image](/images/ResultSuccess/Result_Success_GetBookByID.jpg)
    + Gagal mendapatkan data buku karena id yang dimasukkan tidak terdaftar
    ![image](/images/ResultFailed/Result_Failed_GetBookbyID.jpg)
+ Create New Book
    + Berhasil membuat data buku baru
    ![image](/images/ResultSuccess/Result_Success_AddNewBook.jpg)
    + Gagal membuat data buku baru karena data yang dimasukkan tidak sesuai dengan format yang telah ditentukan
    ![image](/images/ResultFailed/Result_Failed_AddBook.jpg)
+ Update Book By Id
    + Berhasil memperbarui data buku berdasarkan id yang telah terdaftar
    ![image](/images/ResultSuccess/Result_Success_UpdateBook.jpg)
    + Gagal memperbarui data buku karena id yang dimasukkan tidak terdaftar
    ![image](/images/ResultFailed/Result_Failed_UpdateBook_BookNotFound.jpg)
    + Gagal memperbarui data buku karena data yang dimasukkan tidak sesuai dengan format yang telah ditentukan
    ![image](/images/ResultFailed/Result_Failed_UpdateBook.jpg)
+ Delete Book By Id
    + Berhasil menghapus data buku berdasarkan id yang telah terdaftar
    ![image](/images/ResultSuccess/Result_Success_DeleteBook.jpg)
    + Gagal menghapus data buku karena id yang dimasukkan tidak terdaftar
    ![image](/images/ResultFailed/Result_Failed_DeleteBook.jpg)
    
+ [Akses dokumentasi Collection Rest API Postman](/BookLib.postman_collection.json)
