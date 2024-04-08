# Nama target default
.DEFAULT_GOAL := help

# Variabel untuk parameter koneksi database
DB_URL := "mysql://root@tcp(localhost:3306)/hey_notes_api"

# Variabel untuk direktori migrations
MIGRATIONS_DIR := database/migrations/

# Perintah untuk membuat model baru
model:
	@echo "Generating model: $(name)"
	echo -e "package models\n\ntype $(shell echo "$(name)" | sed 's/.*/\u&/') struct {\n    Id int `json:\"id\"`\n}\n" > models/$(name).go

# Perintah untuk menjalankan migrasi
migrate-up:
	migrate -database $(DB_URL) -path $(MIGRATIONS_DIR) up

# Perintah untuk membatalkan migrasi
migrate-down:
	migrate -database $(DB_URL) -path $(MIGRATIONS_DIR) down

# Perintah untuk memaksa migrasi ke versi tertentu
migrate-force:
	migrate -database $(DB_URL) -path $(MIGRATIONS_DIR) force 1

# Perintah untuk membuat file migrasi baru
create-migration:
	migrate create -ext sql -dir $(MIGRATIONS_DIR) -seq create_table_$(name)

# Perintah untuk melakukan migrate down lalu migrate up
migrate-fresh: migrate-down migrate-up

run:
	cd "cmd/app" && go run .
# Target bantuan (help) untuk menampilkan panduan penggunaan makefile
help:
	@echo "Panduan Penggunaan Makefile:"
	@echo "make migrate-up            : Menjalankan migrasi"
	@echo "make migrate-down          : Membatalkan migrasi"
	@echo "make migrate-force         : Memaksa migrasi ke versi tertentu"
	@echo "make create-migration name : Membuat file migrasi baru dengan nama 'name'"