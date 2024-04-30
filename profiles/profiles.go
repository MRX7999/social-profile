package profiles

import (
	"math/rand"
	"time"
)

// Profile представляет профиль человека.
type Profile struct {
	Name           string              // Имя и фамилия
	Age            int                 // Возраст
	Occupation     string              // Должность
	City           string              // Город
	Interests      []string            // Интересы
	AdditionalInfo map[string][]string // Дополнительная информация
}

// Наборы данных для генерации профилей на русском языке
var russianNames = []string{"Иван Иванов", "Петр Петров", "Александр Сидоров", "Елена Козлова", "Мария Новикова"}
var russianOccupations = []string{"Учитель", "Врач", "Инженер", "Бухгалтер", "Юрист"}
var russianCities = []string{"Москва", "Санкт-Петербург", "Новосибирск", "Екатеринбург", "Казань"}
var russianInterests = [][]string{
	{"Литература", "Путешествия", "Фотография"},
	{"Кино", "Искусство", "Театр"},
	{"Спорт", "Музыка", "Кулинария"},
	{"Наука", "Технологии", "Природа"},
	{"Мода", "Шоппинг", "Красота"},
}

// GenerateProfiles создает слайс с синтетическими профилями на русском языке.
func GenerateProfiles(numProfiles int) []Profile {
	profiles := make([]Profile, numProfiles)

	// Инициализация генератора случайных чисел для повторяемости (по желанию)
	rand.Seed(time.Now().UnixNano())

	// Генерация профилей
	for i := 0; i < numProfiles; i++ {
		// Заполнение структуры
		profile := Profile{
			Name:           russianNames[rand.Intn(len(russianNames))],
			Age:            rand.Intn(83) + 18, // Возраст от 18 до 100 лет
			Occupation:     russianOccupations[rand.Intn(len(russianOccupations))],
			City:           russianCities[rand.Intn(len(russianCities))],
			Interests:      russianInterests[rand.Intn(len(russianInterests))],
			AdditionalInfo: make(map[string][]string),
		}

		// Добавление профиля в список
		profiles[i] = profile
	}

	return profiles
}

// FindProfileByName находит профиль по имени
func FindProfileByName(name string, profiles []Profile) *Profile {
	for _, profile := range profiles {
		if profile.Name == name {
			return &profile
		}
	}
	return nil
}
