package fake

import (
	"time"

	"github.com/brianvoe/gofakeit/v6"
)

type Applicant struct {
	ID                      string           `json:"id"`
	User                    User             `json:"user"`
	Fname                   string           `json:"fname"`
	Lname                   string           `json:"lname"`
	Gender                  string           `json:"gender"`
	Title                   string           `json:"title"`
	Headline                *string          `json:"headline"`
	State                   string           `json:"state"`
	HighestQualification    string           `json:"highestQualification"`
	DomainYearsOfExperience int              `json:"domainYearsOfExperience"`
	WorkExperience          []WorkExperience `json:"workExperience"`
	Education               []Education      `json:"education"`
	Certifications          []Certification  `json:"certifications"`
	LookingFor              []LookingFor     `json:"lookingFor"`
	Awards                  []Award          `json:"awards"`
	Languages               []Language       `json:"languages"`
}

type User struct {
	ID string `json:"id"`
}

type WorkExperience struct {
	Title       string    `json:"title"`
	Domain      string    `json:"domain"`
	Workplace   string    `json:"workplace"`
	Location    string    `json:"location"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"startTime"`
	EndTime     time.Time `json:"endTime"`
}

type Education struct {
	Level      string    `json:"level"`
	Title      string    `json:"title"`
	GPA        float64   `json:"GPA"`
	University string    `json:"university"`
	Location   string    `json:"location"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
}

type Certification struct {
	Title      string    `json:"title"`
	Entity     string    `json:"entity"`
	Location   *string   `json:"location"`
	CreditedAt time.Time `json:"creditedAt"`
}

type LookingFor struct {
	JobType string `json:"jobType"`
	Domain  string `json:"domain"`
}

type Award struct {
	AwardedAt  time.Time `json:"awardedAt"`
	AwardedFor string    `json:"awardedFor"`
	AwardedBy  string    `json:"awardedBy"`
}

type Language struct {
	Language     string         `json:"language"`
	Certificates []LanguageCert `json:"certificates"`
	Fluency      string         `json:"fluency"`
}

type LanguageCert struct {
	Title string      `json:"title"`
	Grade interface{} `json:"grade"`
}

// generate fake applicant is one fucking long function
// well it's not really that long (that's what she said)
func GenerateFakeApplicant() Applicant {

	// Generate fake work experience
	var workExperience []WorkExperience
	for i := 0; i < gofakeit.Number(1, 3); i++ {
		startTime := gofakeit.Date()
		endTime := gofakeit.Date()
		if startTime.After(endTime) {
			startTime, endTime = endTime, startTime
		}

		workExperience = append(workExperience, WorkExperience{
			Title:       gofakeit.JobTitle(),
			Domain:      gofakeit.Word(),
			Workplace:   gofakeit.Company(),
			StartTime:   startTime,
			EndTime:     endTime,
			Location:    gofakeit.RandomString(States),
			Description: gofakeit.Sentence(10),
		})
	}

	// Generate fake education
	var education []Education
	for i := 0; i < gofakeit.Number(1, 3); i++ {
		startTime := gofakeit.Date()
		endTime := gofakeit.Date()
		if startTime.After(endTime) {
			startTime, endTime = endTime, startTime
		}
		education = append(education, Education{
			Level:      gofakeit.RandomString(Qualifications),
			Title:      gofakeit.Sentence(3),
			GPA:        gofakeit.Float64Range(0, 4),
			University: gofakeit.Company(),
			Location:   gofakeit.RandomString(States),
			StartTime:  startTime,
			EndTime:    endTime,
		})
	}

	// Generate fake certifications
	var certifications []Certification
	for i := 0; i < gofakeit.Number(1, 3); i++ {
		creditedAt := gofakeit.Date()
		var location string
		if gofakeit.Bool() {
			location = gofakeit.Word()
		}
		certifications = append(certifications, Certification{
			Title:      gofakeit.Sentence(3),
			Entity:     gofakeit.Company(),
			CreditedAt: creditedAt,
			Location:   &location,
		})
	}

	// Generate fake looking for
	var lookingFor []LookingFor
	for i := 0; i < gofakeit.Number(1, 3); i++ {
		lookingFor = append(lookingFor, LookingFor{
			Domain:  gofakeit.Word(),
			JobType: gofakeit.RandomString([]string{"Full-time", "Part-time", "Internship"}),
		})
	}

	// Generate Award
	awards := []Award{}
	numAwards := gofakeit.Number(0, 5)
	for i := 0; i < numAwards; i++ {
		award := Award{
			AwardedAt:  gofakeit.Date(),
			AwardedFor: gofakeit.Sentence(5),
			AwardedBy:  gofakeit.Company(),
		}
		awards = append(awards, award)
	}

	// Generate Language
	languages := []Language{}
	numLanguages := gofakeit.Number(1, 3)
	for i := 0; i < numLanguages; i++ {
		language := Language{
			Language: gofakeit.Language(),
			Fluency:  gofakeit.RandomString(LanguageFluencies),
		}

		// Generate LanguageCert
		numCerts := gofakeit.Number(0, 3)
		for j := 0; j < numCerts; j++ {
			cert := LanguageCert{
				Title: gofakeit.Sentence(2),
				Grade: gofakeit.Number(0, 100),
			}
			language.Certificates = append(language.Certificates, cert)
		}

		languages = append(languages, language)
	}

	var applicant Applicant

	// Generate applicant
	applicant.User.ID = gofakeit.UUID()
	applicant.ID = applicant.User.ID

	var headline string
	if gofakeit.Bool() {
		headline = gofakeit.Sentence(20)
	}
	applicant.Headline = &headline

	applicant.Fname = gofakeit.FirstName()
	applicant.Lname = gofakeit.LastName()
	applicant.Gender = gofakeit.Gender()
	applicant.Title = gofakeit.JobTitle()
	applicant.State = gofakeit.RandomString(States)
	applicant.DomainYearsOfExperience = gofakeit.Number(0, 20)
	applicant.HighestQualification = gofakeit.RandomString(Qualifications)
	applicant.WorkExperience = workExperience
	applicant.Education = education
	applicant.Certifications = certifications
	applicant.LookingFor = lookingFor
	applicant.Awards = awards
	applicant.Languages = languages

	return applicant
}

func GetRandomApplicant(applicants []Applicant) (Applicant, int) {
	i := gofakeit.Number(0, len(applicants)-1)
	return applicants[i], i
}
