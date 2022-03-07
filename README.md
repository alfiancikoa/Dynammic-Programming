# Dynamic-Programming

<p align=justify><b>Program Dinamis</b> (Dynamic Program) merupakan metode pemecahan masalah dengan cara menguraikan solusi menjadi sekumpulan langkah (step) atau tahapan (stage) sedemikian sehingga solusi dari persoalan dapat dipandang dari serangkaian keputusan yang saling berkaitan. Pada program dinamis, rankaian keputusan yang optimal dibuat dengan menggunakan prinsip optimalisasi.</p>
<br>

<b>5 easy steps to DP (Dynamic Program):</b>

<br>
1. Define subproblems
2. Guests (part of solution)
3. Relate subproblems solutions
4. Recursive & Memoize or build DP table bottom-up
5. Solve original problem

#### Dynamic Programming Example

Fibo ke-n = Fibo(n-1) + Fibo(n-2)
<br>

Deklarasi awal:
 - Fibo ke 0 -> 0
 - Fibo ke 1 dan ke 2 -> 1

<br>
Contoh di atas merupakan contoh penerapan dynamic Programming dengan pencarian bilangan fibonacci ke-n, dengan menggunakan pendekatan rekursif 
<br>

1. Ketika bilangan sudah dihitung, maka hasilnya ditaruh ke dalam memoize/chace sehingga kita tidak perlu menghitung dari awal
2. Sebelum melakukan proses perhitungan, nilai fibonacci ke-n dicek dulu apakah sudah disimpan atau beum
3. Jika sudah disimpan, maka langsung return hasilnya
4. Jika belum, maka lakukan proses perhitungan bilangan fibonacci
5. Setelah didapatkan hasil perhitungan, maka akan disimpan ke dalam memoize

<br>
<i>Jangan ragu untuk memberi kritik dan saran, silahkan berikan issue terkait dengan apa yang harus diperbaiki oleh penulis. <a href="https://github.com/alfiancikoa/Dynammic-Programming/issues">here</a></i>
