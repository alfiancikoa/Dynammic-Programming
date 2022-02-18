# Dynamic-Programming
Dynamic Programming Example

Fibo ke-n = Fibo(n-1) + Fibo(n-2)
<br>

syarat:
 - Fibo ke 0 -> 0
 - Fibo ke 1 dan ke 2 -> 1

<br>
Contoh di atas merupakan contoh penerapan dynamic Programming dengan pencrian bilangan fibonacci ke-n
<br>

1. Ketika bilangan sudah dihitung, maka hasilnya ditaruh ke dalam memoize/chace sehingga kita tidak perlu menghitung dari awal
2. Sebelum melakukan proses perhitungan, nilai fibonacci ke-n dicek dulu apakah sudah disimpan atau beum
3. Jika sudah disimpan, maka langsung return hasilnya
4. Jika belum, maka lakukan proses perhitungan bilangan fibonacci
