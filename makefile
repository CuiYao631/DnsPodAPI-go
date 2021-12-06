.PHONY: dev

dev: export COURSE_PLAN_POSTGRESQL_DSN=host=localhost port=5432 user=admin password=password dbname=postgres sslmode=disable

dev: export HOST=https://dnsapi.cn
dev: export NAME=XaioCui_DDNS
dev: export ID=240684
dev: export TOKEN=b61a5717428d23b9122efc6bf4237a67
dev: export SECRETID=AKIDISQAgy2sUs2xd5IG0VmqnQoxxgAPQEqG
dev: export SECRETKEY=NRK0RdcOyOaYIn7B0d2KPHuJTO4e07zj
dev: export LISTEN=:3000


dev:
	@go run main.go
