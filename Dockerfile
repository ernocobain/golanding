# Tahap 1: Build - Menggunakan image Golang untuk kompilasi
FROM golang:1.24.4 AS builder

# Menentukan direktori kerja di dalam kontainer
WORKDIR /app

# Salin file manajemen dependensi terlebih dahulu
# Ini memanfaatkan cache Docker agar tidak perlu download ulang setiap kali kode berubah
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server -ldflags "-w -s" .


FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

COPY --from=builder /app/static ./static
COPY --from=builder /app/views ./views

EXPOSE 8080

CMD ["/app/server"]