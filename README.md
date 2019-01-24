# Custom Auth with golang, combine web app and grpc


* Membuka example app custom auth

![GitHub Logo](/img/1.jpg)



* login dengan custom auth, activity akan mengakses url dengan webview, dan user menginputkan username and password

![GitHub Logo](/img/2.jpg)



* jika login berhasil, akan kembali ke activity sebelumnya dengan membawa data akun

![GitHub Logo](/img/3.jpg)



# Conclusion

tidak mudah untuk mendapakan callback dari hasil login di webview, app android dan webview tidak saling berhubungan satu-sama lain, webview hanya mengakses url dan app android tidak sepenuhnya punya kendali dengan konten yg diakses oleh webview.

bagaimana mendapatkan callback jika login telah di proses?

## dan disitulah proses yg akan dihandle dengan grpc stream 

* data akan dibroadcast oleh backend ke client
* client akan menerima callback data dan user data

