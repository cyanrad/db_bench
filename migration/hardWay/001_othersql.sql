
-- CreateEnum
CREATE TYPE "jobType" AS ENUM ('FULL_TIME', 'PART_TIME', 'INTERNSHIP');

-- CreateEnum
CREATE TYPE "state" AS ENUM ('DUBAI', 'AJMAN', 'SHARJAH', 'ABU_DHABI', 'UMM_AL_QUWAIN', 'FUJAIRAH', 'RAS_AL_KHAIMAH');

-- CreateEnum
CREATE TYPE "qualification" AS ENUM ('NONE', 'DIPLOMA', 'ASSOCIATE', 'BACHELORS', 'MASTERS', 'DOCTORATE');

-- CreateEnum
CREATE TYPE "languageFluency" AS ENUM ('NOVICE', 'INTERMEDIATE', 'PROFESSIONAL', 'FLUENT', 'NATIVE');

-- CreateEnum
CREATE TYPE "companyCategory" AS ENUM ('GOVERNMENT', 'SEMI_GOVERNMENT', 'PRIVATE');

-- CreateTable
CREATE TABLE "applicant" (
    "id" TEXT NOT NULL,
    "fname" TEXT NOT NULL,
    "lname" TEXT NOT NULL,
    "gender" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "headline" TEXT,
    "state" "state" NOT NULL,
    "highestqualification" "qualification" NOT NULL,
    "domainyearsofexperience" INTEGER NOT NULL,

    CONSTRAINT "applicant_pkey" PRIMARY KEY ("id")
);

CREATE TABLE "work_experience" (
    "id" SERIAL PRIMARY KEY,
    "applicant_id" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "domain" TEXT NOT NULL,
    "workplace" TEXT NOT NULL,
    "location" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "start_time" TIMESTAMP NOT NULL,
    "end_time" TIMESTAMP NOT NULL
);
ALTER TABLE "work_experience" ADD CONSTRAINT "fk_work_experience_applicant_id" FOREIGN KEY ("applicant_id") REFERENCES "applicant" ("id");

CREATE TABLE "education" (
    "id" SERIAL PRIMARY KEY,
    "applicant_id" TEXT NOT NULL,
    "level" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "gpa" FLOAT NOT NULL,
    "university" TEXT NOT NULL,
    "location" TEXT NOT NULL,
    "start_time" TIMESTAMP NOT NULL,
    "end_time" TIMESTAMP NOT NULL
);
ALTER TABLE "education" ADD CONSTRAINT "fk_education_applicant_id" FOREIGN KEY ("applicant_id") REFERENCES "applicant" ("id");

CREATE TABLE "certification" (
    "id" SERIAL PRIMARY KEY,
    "applicant_id" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "entity" TEXT NOT NULL,
    "location" TEXT,
    "credited_at" TIMESTAMP NOT NULL
);
ALTER TABLE "certification" ADD CONSTRAINT "fk_certification_applicant_id" FOREIGN KEY ("applicant_id") REFERENCES "applicant" ("id");

CREATE TABLE "looking_for" (
    "id" SERIAL PRIMARY KEY,
    "applicant_id" TEXT NOT NULL,
    "job_type" TEXT NOT NULL,
    "domain" TEXT NOT NULL
);
ALTER TABLE "looking_for" ADD CONSTRAINT "fk_looking_for_applicant_id" FOREIGN KEY ("applicant_id") REFERENCES "applicant" ("id");

CREATE TABLE "award" (
    "id" SERIAL PRIMARY KEY,
    "applicant_id" TEXT NOT NULL,
    "awarded_at" TIMESTAMP NOT NULL,
    "awarded_for" TEXT NOT NULL,
    "awarded_by" TEXT NOT NULL
);
ALTER TABLE "award" ADD CONSTRAINT "fk_award_applicant_id" FOREIGN KEY ("applicant_id") REFERENCES "applicant" ("id");

CREATE TABLE "language" (
    "id" SERIAL PRIMARY KEY,
    "applicant_id" TEXT REFERENCES "applicant" ("id"),
    "language" TEXT NOT NULL,
    "fluency" TEXT NOT NULL
);
ALTER TABLE "language" ADD CONSTRAINT "fk_language_applicant_id" FOREIGN KEY ("applicant_id") REFERENCES "applicant" ("id");