### setup dev env
setup: 
	go get -u gorm.io/gorm
	go get -u gorm.io/driver/sqlite
	go get -u github.com/gin-gonic/gin

build:
	docker compose build --no-cache

up:
	docker compose up -d

down:
	docker compose down