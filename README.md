## IP Data
<p>Kiritilgan ip manzil bo'yicha uni ma'lumotlari: </p><br>
<ul>
<li>Mamlakat</li>
<li>Shahar</li>
<li>Internet Provayder</li>
<li>Vaqt Zonasi</li>
</ul>
<p>Ushbu manzillarni sizga ko'rsatadi.
</p>

<h2>Albatta ushbu ma'lumotlarni ko'rsatishi uchun siz static ip manzil ma'lumotni ko'rsatishingiz kerak</h2>

<h1>Foydalanish:</h1>
<ul>
<li><i>ip-darwin-amd64</i> - Intel CPU MacOS uchun</li>
<li><i>ip-darwin-arm64</i> - ARM (M1, M2) CPU MacOS uchun</li>
<li><i>ip-linux-amd64</i> - x86_64 CPU Linux uchun</li>
<li><i>ip-linux-arm64</i> - arm64 CPU Linux uchun</li>
<li><i>ip-windows-amd64.exe</i> - 64 bit CPU Windows uchun</li>
<li><i>ip-windows-arm64</i> - arm64 CPU Windows uchun</li>

<h2>* Agar dastur ishlamasa uni qayta kompilyatsiya qilishingiz kerak</h2>

```
GOOS=os GOARCH=arch go build api.go
```
os = bu sizni Operatsion tizimingiz(Windows, Linux va Darwin(MacOS))
arch = amd64 yoki arm64 (Intel/AMD yoki Armga asoslangan CPU)

### Kompilyatsiya qilish uchun sizda Go kompilyatori kerak! Uni <a href="https://go.dev">GO</a>ning rasmiy saytidan olishingiz mumkun
</ul>

# Anvar Alimov (@cosmoasx) - Go ❤️