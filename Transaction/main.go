package main

import (
	"log"
	"time"
	"transaction/db"
	"transaction/model"
	"transaction/settings"
	"transaction/utils"
	"github.com/go-co-op/gocron"
)

func main() {
	// соединение с базой данных
	settings.ReadConfig("settings/config.json")
	log.Println("Открытие соединения с базой данных...")
	if err := db.ConectionToDB(); err != nil {
		log.Fatalf("Ошибка при открытии соединения с базой данных: %v", err)
	}
	log.Println("Соединение с базой данных успешно открыто!")

	//миграция данных
	log.Println("Запуск миграции базы данных...")
	if err := db.DB.AutoMigrate(
		&model.Trnx{},
	); err != nil {
		log.Fatalf("Ошибка при запуске миграции базы данных: %v", err)
	}
	log.Println("Миграция базы данных успешно завершена!")

	//джоба для отправки сообщений
	job := gocron.NewScheduler(time.UTC)
	job.Cron("@daily").Do(utils.CheckDataVendors)
	job.StartAsync()

	log.Println("отчёт по транзакциям происходит каждый день!")

	select {}

}
