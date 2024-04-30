package socialengineering

import (
	"fmt"
	"math/rand"
	"time"

	"main.go/profiles"
)

var rnd = rand.New(rand.NewSource(time.Now().UnixNano()))

var emailTemplates = []string{
	"Тема: Важное сообщение о вашей учетной записи\n\nУважаемый %s,\n\nМы заметили подозрительную активность на вашей учетной записи. Пожалуйста, перейдите по следующей ссылке, чтобы подтвердить ваши данные: <злоумышленная-ссылка>\n\nС уважением,\nКоманда безопасности",
	"Тема: Требуется срочное действие: безопасность вашей учетной записи\n\nУважаемый %s,\n\nБезопасность вашей учетной записи нарушена. Перейдите по этой ссылке, чтобы сейчас защитить свою учетную запись: <злоумышленная-ссылка>\n\nС наилучшими пожеланиями,\nКоманда безопасности",
}

var phoneCallTemplates = []string{
	"Здравствуйте, это %s? Я звоню из отдела безопасности вашего банка. Мы обнаружили несанкционированную активность на вашем счете. Пожалуйста, предоставьте свои учетные данные, чтобы подтвердить вашу личность.",
	"Добрый день, %s. Это любезный звонок от команды безопасности вашего финансового учреждения. Мы обнаружили неправильную активность на вашем счете. Для защиты вашей учетной записи, пожалуйста, подтвердите ваши данные, позвонив нам по номеру <фейковый-номер-телефона>.",
}

var smsTemplates = []string{
	"Важно: Мы обнаружили необычную активность на вашей учетной записи. Пожалуйста, немедленно позвоните нам по номеру <фейковый-номер-телефона>, чтобы подтвердить ваши данные.",
	"Предупреждение: Обнаружена подозрительная активность на вашей учетной записи. Ответьте 'ЗАЩИТА', чтобы получить защищенную ссылку для защиты вашей учетной записи.",
}

func GeneratePhishingAttacks(profile profiles.Profile) []string {
	var attacks []string

	attacks = append(attacks, generatePersonalizedEmailPhishingAttack(profile))
	attacks = append(attacks, generatePersonalizedPhoneCallPhishingAttack(profile))
	attacks = append(attacks, generatePersonalizedSMSPhishingAttack(profile))

	return attacks
}

func generatePersonalizedEmailPhishingAttack(profile profiles.Profile) string {
	subject := "Важное сообщение о вашей учетной записи"
	template := chooseRandomTemplate(emailTemplates)
	message := fmt.Sprintf(template, profile.Name)
	return fmt.Sprintf("Тема: %s\n\n%s", subject, message)
}

func generatePersonalizedPhoneCallPhishingAttack(profile profiles.Profile) string {
	template := chooseRandomTemplate(phoneCallTemplates)
	return fmt.Sprintf(template, profile.Name)
}

func generatePersonalizedSMSPhishingAttack(profile profiles.Profile) string {
	template := chooseRandomTemplate(smsTemplates)
	return template
}

func GenerateSegmentedPhishingAttacks(profiles []profiles.Profile) map[string][]string {
	segmentedAttacks := make(map[string][]string)

	for _, profile := range profiles {
		attacks := GeneratePhishingAttacks(profile)
		segment := determineSegment(profile)
		segmentedAttacks[segment] = append(segmentedAttacks[segment], attacks...)
	}

	return segmentedAttacks
}

func determineSegment(profile profiles.Profile) string {
	if profile.Age < 30 {
		return "Молодые взрослые"
	} else {
		return "Средний возраст"
	}
}

func GenerateVarietyOfAttacks(profile profiles.Profile) []string {
	var attacks []string

	for i := 0; i < 5; i++ {
		attacks = append(attacks, generateEmailPhishingAttack(profile))
		attacks = append(attacks, generatePhoneCallPhishingAttack(profile))
		attacks = append(attacks, generateSMSPhishingAttack(profile))
	}

	rand.Shuffle(len(attacks), func(i, j int) {
		attacks[i], attacks[j] = attacks[j], attacks[i]
	})

	return attacks
}

func generateEmailPhishingAttack(profile profiles.Profile) string {
	subject := "Важное сообщение о вашей учетной записи"
	message := fmt.Sprintf("Уважаемый %s,\n\nМы заметили подозрительную активность на вашей учетной записи. Пожалуйста, перейдите по следующей ссылке, чтобы подтвердить ваши данные: <злоумышленная-ссылка>\n\nС уважением,\nКоманда безопасности", profile.Name)
	return fmt.Sprintf("Тема: %s\n\n%s", subject, message)
}

func generatePhoneCallPhishingAttack(profile profiles.Profile) string {
	return fmt.Sprintf("Здравствуйте, это %s? Я звоню из отдела безопасности вашего банка. Мы обнаружили несанкционированную активность на вашем счете. Пожалуйста, предоставьте свои учетные данные, чтобы подтвердить вашу личность.", profile.Name)
}

func generateSMSPhishingAttack(profile profiles.Profile) string {
	return fmt.Sprintf("Важно: Мы обнаружили необычную активность на вашей учетной записи. Пожалуйста, немедленно позвоните нам по номеру <фейковый-номер-телефона>, чтобы подтвердить ваши данные.")
}

func chooseRandomTemplate(templates []string) string {
	index := rnd.Intn(len(templates))
	return templates[index]
}
