package web

import (
	"net/http"
	"strconv"

	"main.go/profiles"
	"main.go/socialengineering"

	"github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
)

var profilesList []profiles.Profile // Список профилей

// parseProfile извлекает данные профиля из результатов запроса API.
func parseProfile(result map[string]interface{}) *profiles.Profile {
	profile := &profiles.Profile{
		Name:       result["names"].([]string)[0],
		Age:        result["borns"].([]int)[0],
		Occupation: result["professions"].([]string)[0],
		City:       result["adressess"].([]string)[0],
		Interests:  result["infoAddInfo"].([]string),
		// Дополнительные поля профиля можно добавить по необходимости
	}

	return profile
}

func StartServer() {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html")

	// Генерация профилей при запуске сервера и сохранение в списке profilesList
	profilesList = profiles.GenerateProfiles(10)

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"profiles": profilesList, // Передача списка профилей в шаблон
		})
	})

	router.GET("/export-docx", func(c *gin.Context) {
		// Создание фишинговых атак для каждого профиля
		var phishingAttacks []string
		for _, profile := range profilesList {
			attacks := socialengineering.GeneratePhishingAttacks(profile)
			phishingAttacks = append(phishingAttacks, attacks...)
		}

		// Создание нового документа XLSX
		file := excelize.NewFile()
		index := file.NewSheet("Profiles Report")
		file.SetActiveSheet(index)

		// Заголовок таблицы
		headers := []string{"Имя", "Возраст", "Занятость", "Город проживания", "Фишинговая атака"}
		for col, header := range headers {
			cell := excelize.ToAlphaString(col) + "1"
			file.SetCellValue("Profiles Report", cell, header)
		}

		// Добавление данных в таблицу
		for i, profile := range profilesList {
			row := i + 2
			file.SetCellValue("Profiles Report", "A"+strconv.Itoa(row), profile.Name)
			file.SetCellValue("Profiles Report", "B"+strconv.Itoa(row), profile.Age)
			file.SetCellValue("Profiles Report", "C"+strconv.Itoa(row), profile.Occupation)
			file.SetCellValue("Profiles Report", "D"+strconv.Itoa(row), profile.City)

			attack := ""
			if len(phishingAttacks) > i {
				attack = phishingAttacks[i]
			}
			file.SetCellValue("Profiles Report", "E"+strconv.Itoa(row), attack)
		}

		// Сохранение документа XLSX
		fileName := "profiles.xlsx"
		if err := file.SaveAs(fileName); err != nil {
			c.String(http.StatusInternalServerError, "Error exporting to XLSX: %s", err.Error())
			return
		}

		// Отправка сгенерированного документа XLSX клиенту как вложение
		c.FileAttachment(fileName, fileName)
	})

	router.GET("/profile-and-attacks", func(c *gin.Context) {
		// Получение имени профиля из параметра запроса
		profileName := c.Query("name")

		// Поиск соответствующего профиля в списке профилей
		var profile *profiles.Profile
		for _, p := range profilesList {
			if p.Name == profileName {
				profile = &p
				break
			}
		}

		// Если профиль не найден, отправляем ошибку
		if profile == nil {
			c.String(http.StatusNotFound, "Profile not found")
			return
		}

		// Создание фишинговых атак для найденного профиля
		attacks := socialengineering.GeneratePhishingAttacks(*profile)

		// Отправка информации о профиле и его атаках в формате JSON
		c.JSON(http.StatusOK, gin.H{
			"profile": profile,
			"attacks": attacks,
		})
	})

	router.Run(":8080")
}
