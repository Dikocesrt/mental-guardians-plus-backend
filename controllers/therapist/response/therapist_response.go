package response

type TherapistResponse struct {
	ID                     uint   `json:"id"`
	Name                   string `json:"name"`
	Age                    int    `json:"age"`
	Specialist             string `json:"specialist"`
	PhotoURL               string `json:"photoURL"`
	PhoneNumber            string `json:"phoneNumber"`
	Gender                 string `json:"gender"`
	Experience             int    `json:"experience"`
	Fee                    int    `json:"fee"`
	PracticeCity           string `json:"practiceCity"`
	PracticeLocation       string `json:"practiceLocation"`
	BachelorAlmamater      string `json:"bachelorAlmamater"`
	BachelorGraduationYear int    `json:"bachelorGraduationYear"`
	ConsultationMode       string `json:"consultationMode"`
}