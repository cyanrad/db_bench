package db

import (
	"context"
	"fmt"
	"log"
	"main/fake"
	"math"
	"os"
	"testing"

	"github.com/brianvoe/gofakeit"
	"github.com/jackc/pgx/v5"
)

// there is a better way of doing this, fuck you
const CONNECTION_STRING = "postgresql://rootuser:rootpass@localhost:5434/testing"

var dbConn *pgx.Conn

func TestMain(m *testing.M) {
	log.Print("connecting")

	var err error
	dbConn, err = pgx.Connect(context.Background(), CONNECTION_STRING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	}

	m.Run()
}

func BenchmarkCreateCompany(b *testing.B) {
	clearCompanyTable(dbConn)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		newCompany := DbCompany{
			uuid:        gofakeit.UUID(),
			companyName: gofakeit.Company(),
			rating:      math.Round(math.Floor(gofakeit.Float64Range(0.0, 5.0)*10) / 10),
			lat:         gofakeit.Float64Range(-90.0, 90.0),
			lon:         gofakeit.Float64Range(-90.0, 90.0),
			companyType: gofakeit.RandString([]string{"PRIVATE", "GOVERNMENT", "SEMI_GOVERNMENT"}),
			createdAt:   gofakeit.Date(),
			updatedAt:   gofakeit.Date(),
		}
		b.StartTimer()

		createCompany(dbConn, newCompany)
	}
}

func BenchmarkCreateApplicant(b *testing.B) {
	clearApplicantTable(dbConn)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		applicant := fake.GenerateFakeApplicant()
		b.StartTimer()

		createApplicant(dbConn, applicant)
	}
}

func BenchmarkCreateApplicantTheHardWay(b *testing.B) {
	clearApplicantTableTheHardWay(dbConn)
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		b.StopTimer()
		applicant := fake.GenerateFakeApplicant()
		b.StartTimer()

		err := createApplicantTheHardWay(dbConn, applicant)
		if err != nil {
			b.Log(err)
		}
	}
}
