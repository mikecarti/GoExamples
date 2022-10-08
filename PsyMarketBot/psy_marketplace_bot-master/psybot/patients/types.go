package patients

type Patient struct {
	userID int
	name   string
	age    int
	city   string

	// Male / Female / Non-Binary
	sex string

	// Проблема с которой хотят обратиться к психологу
	issue string

	// Online / Offline
	prefferedWorkWay string

	// Предпочтительный пол пациента
	prefferedSex string
}

func New(userID int, name string, age int, city string, sex string, issue string, prefferedWorkWay string, prefferedSex string) Patient {
	return Patient{
		userID:           userID,
		name:             name,
		age:              age,
		city:             city,
		sex:              sex,
		issue:            issue,
		prefferedWorkWay: prefferedWorkWay,
		prefferedSex:     prefferedSex,
	}
}
