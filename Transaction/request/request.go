package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"transaction/db"
	"transaction/model"
)

func ReqServer(trnx []model.ShowTrn) (err error) {

	var (
		chatID int64
		data   string
		count  int64
		status int
	)

	status = 400
	chatID = -1002192916773

	err = db.DB.Model(&model.Trnx{}).Where("status = ? and created_at >= CURRENT_TIMESTAMP - interval '24 hours'", status).Count(&count).Find(&count).Error
	if err != nil {
		log.Println("ошибка при получении количества из ДБ!")
	}
	data = fmt.Sprintf("Total count:  %v\n", count)

	for _, trnxs := range trnx {
		data += fmt.Sprintf("Vender: %v\n Count: %v\n", trnxs.Vendor, trnxs.Total)
	}

	if count < 10 {
		log.Println("Количество траннзакций не превышает 10-ти!")
		return nil
	}

	messeges := model.Message{
		ChatID: chatID, Text: data,
	}

	jsonData, err := json.Marshal(messeges)
	if err != nil {
		fmt.Println("ошибка при конвертации в JSON:", err)
		return
	}

	req, err := http.NewRequest("POST", "https://api.telegram.org/bot6695075846:AAHqsqHzb-PCGItyX1ijAm9PCLwklkshcUE/sendMessage", bytes.NewBuffer(jsonData))
	if err != nil {
		return
	}

	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Println("err to send msg: ", err)
		return
	}
	defer resp.Body.Close()

	return
}
