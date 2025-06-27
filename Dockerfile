# Tahap 1: Build - Menggunakan image Golang untuk kompilasi
FROM golang:1.24.4 AS builder

# Menentukan direktori kerja di dalam kontainer
WORKDIR /app

# Salin file manajemen dependensi terlebih dahulu
# Ini memanfaatkan cache Docker agar tidak perlu download ulang setiap kali kode berubah
COPY go.mod go.sum ./
RUN go mod download

# Salin seluruh kode sumber proyek
COPY . .

# Lakukan kompilasi aplikasi Goy
# CGO_ENABLED=0 membuat static binary yang tidak bergantung pada library C
# -ldflags "-w -s" untuk memperkecil ukuran file hasil kompilasi
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server -ldflags "-w -s" .

# ---

# Tahap 2: Final - Menggunakan image minimalis untuk production
FROM alpine:latest

WORKDIR /app

# Salin file hasil kompilasi dari tahap 'builder' ke image final
COPY --from=builder /app/server .

# (Opsional tapi direkomendasikan) Salin aset statis jika ada
# Jika Anda punya folder 'static' dan 'views' untuk template
COPY --from=builder /app/static ./static
COPY --from=builder /app/views ./views

# Memberitahu Docker port mana yang akan diekspos oleh kontainer
# Cloud Run akan menggunakan port ini secara default
EXPOSE 8080

# Perintah yang akan dijalankan saat kontainer dimulai
# Menjalankan file binary 'server' yang sudah kita kompilasi
CMD ["/app/server"]