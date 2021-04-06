# 💫 Tugas Kecil 3 IF2211 Strategi Algoritma 💫
# 💫 Penyelesaian Shortest Path Dengan Algoritma A* 💫
### Description
Algoritma A* atau biasa disebut dengan A-star merupakan salah satu algoritma yang termasuk dalam kategori
metode pencarian yang memiliki informasi (informed search method). Algoritma ini sangat efektif digunakan
sebagai solusi proses path finding (pencari jalan). Algoritma ini mencari jarak rute terpendek yang
akan ditempuh suatu point awal (source point) hinigga ke objek tujuan  (destination point). Teknik pencarian yang
digunakan dalam aplikasi berbasis web kami ini adalah menggunakan penerapan Algoritma A* dengan fungsi heuristic pada google maps API. ⭐️ ✨ Algoritma A* ✨ ⭐️

### Requirements
go version go1.16.3 windows/amd64
GNU Make 4.3

### How To Start
##### 💻 For Windows 💻
run from the command line:
make dev
open localhost:5000 on the browser

### 💫 Features 💫
1. Load nodes from src/constants
2. Manually draw nodes in the maps API

### 💫 Controls 💫
A. Map
<ul>
    <li>Basic google maps controls</li>
    <li>Left click to draw a marker</li>
    <li>Clear Nodes to remove all markers and polylines</li>
    <li>Find shortest path to get the path and distance between 2 markers</li>
</ul>
B. Marker
<ul>
    <li>Left click two markers to connect them with an edge</li>
    <li>Right click marker to select them as a source/destination marker</li>
    <li>Double Click to remove marker</li>
</ul>
C. Polylines
<ul>
    <li>Double Click to remove polyline</li>
</ul>

🖤 Enjoy the app! 🖤

### Note :
Don't use 127.0.0.1 to access localhost as the app doesn't allow CORS

### Made by
[Jesson Gosal Yo](https://www.linkedin.com/in/jesson-yo/)\
Marcello Faria