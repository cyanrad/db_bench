package db

import (
	"context"
	"encoding/json"
	"fmt"
	"main/fake"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
)

type DbCompany struct {
	uuid        string
	companyName string
	rating      float64
	lat         float64
	lon         float64
	companyType string
	createdAt   time.Time
	updatedAt   time.Time
}

func createCompany(conn *pgx.Conn, c DbCompany) {
	conn.Exec(
		context.Background(),
		"INSERT INTO company VALUES ($1, $2, $3, $4, $5, $6, $7, $8);",
		c.uuid,
		c.companyName,
		c.rating,
		c.lat,
		c.lon,
		c.companyType,
		c.createdAt,
		c.updatedAt,
	)
}

func createApplicant(conn *pgx.Conn, a fake.Applicant) (pgconn.CommandTag, error) {
	// Convert slices of structs to JSON
	workExperience, err := json.Marshal(a.WorkExperience)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	education, err := json.Marshal(a.Education)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	certifications, err := json.Marshal(a.Certifications)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	lookingFor, err := json.Marshal(a.LookingFor)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	awards, err := json.Marshal(a.Awards)
	if err != nil {
		return pgconn.CommandTag{}, err
	}
	languages, err := json.Marshal(a.Languages)
	if err != nil {
		return pgconn.CommandTag{}, err
	}

	// Execute the statement with the struct fields as arguments
	ret, err := conn.Exec(
		context.Background(),
		`INSERT INTO applicant (
			id,
			fname,
			lname,
			gender,
			title,
			headline,
			state,
			highestqualification,
			domainyearsofexperience,
			workexperience,
			education,
			certifications,
			lookingfor,
			awards,
			languages
		) VALUES (
			$1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15
		)`,
		a.ID,
		a.Fname,
		a.Lname,
		a.Gender,
		a.Title,
		a.Headline,
		a.State,
		a.HighestQualification,
		a.DomainYearsOfExperience,
		workExperience,
		education,
		certifications,
		lookingFor,
		awards,
		languages,
	)

	return ret, err
}

func createApplicantTheHardWay(conn *pgx.Conn, a fake.Applicant) error {
	// Execute the statement with the struct fields as arguments
	_, err := conn.Exec(
		context.Background(),
		`INSERT INTO applicant (
			id,
			fname,
			lname,
			gender,
			title,
			headline,
			state,
			highestqualification,
			domainyearsofexperience
		)  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`,
		a.ID,
		a.Fname,
		a.Lname,
		a.Gender,
		a.Title,
		a.Headline,
		a.State,
		a.HighestQualification,
		a.DomainYearsOfExperience,
	)
	if err != nil {
		return fmt.Errorf("%v: 1", err)
	}

	for _, we := range a.WorkExperience {
		_, err := conn.Exec(context.Background(), `INSERT INTO work_experience (applicant_id, title, domain, workplace, location, description, start_time, end_time)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, a.ID, we.Title, we.Domain, we.Workplace, we.Location, we.Description, we.StartTime, we.EndTime)
		if err != nil {
			return err
		}

	}
	for _, ed := range a.Education {
		_, err := conn.Exec(context.Background(), `INSERT INTO education (applicant_id, level, title, gpa, university, location, start_time, end_time)
        VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`, a.ID, ed.Level, ed.Title, ed.GPA, ed.University, ed.Location, ed.StartTime, ed.EndTime)
		if err != nil {
			return err
		}
	}

	for _, cert := range a.Certifications {
		_, err := conn.Exec(context.Background(), `INSERT INTO certification (applicant_id, title, entity, location, credited_at)
        VALUES ($1, $2, $3, $4, $5)`, a.ID, cert.Title, cert.Entity, cert.Location, cert.CreditedAt)
		if err != nil {
			return err
		}
	}

	for _, lf := range a.LookingFor {
		_, err := conn.Exec(context.Background(), `INSERT INTO looking_for (applicant_id, job_type, domain)
        VALUES ($1, $2, $3)`, a.ID, lf.JobType, lf.Domain)
		if err != nil {
			return err
		}
	}

	for _, aw := range a.Awards {
		_, err := conn.Exec(context.Background(), `INSERT INTO award (applicant_id, awarded_at, awarded_for, awarded_by)
        VALUES ($1, $2, $3, $4)`, a.ID, aw.AwardedAt, aw.AwardedFor, aw.AwardedBy)
		if err != nil {
			return err
		}
	}

	for _, lang := range a.Languages {
		_, err := conn.Exec(context.Background(), `INSERT INTO language (applicant_id, language, fluency)
        VALUES ($1, $2, $3)`, a.ID, lang.Language, lang.Fluency)
		if err != nil {
			return err
		}
	}

	return nil
}

func clearCompanyTable(conn *pgx.Conn) pgconn.CommandTag {
	resp, err := conn.Exec(
		context.Background(),
		"DELETE FROM company;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't do query: %v\n", err)
		os.Exit(1)
	}

	return resp
}

// i know this is fucking dumb....yeah there is no but
func clearApplicantTable(conn *pgx.Conn) pgconn.CommandTag {
	resp, err := conn.Exec(
		context.Background(),
		"DELETE FROM applicant;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't do query: %v\n", err)
		os.Exit(1)
	}

	return resp
}

func clearApplicantTableTheHardWay(conn *pgx.Conn) {
	_, err := conn.Exec(
		context.Background(),
		"DELETE FROM looking_for;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't do query: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(
		context.Background(),
		"DELETE FROM work_experience;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't do query: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(
		context.Background(),
		"DELETE FROM education;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't do query: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(
		context.Background(),
		"DELETE FROM certification;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't do query: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(
		context.Background(),
		"DELETE FROM award;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't do query: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(
		context.Background(),
		"DELETE FROM language;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't do query: %v\n", err)
		os.Exit(1)
	}

	_, err = conn.Exec(
		context.Background(),
		"DELETE FROM applicant;")

	if err != nil {
		fmt.Fprintf(os.Stderr, "can't do query: %v\n", err)
		os.Exit(1)
	}
}
