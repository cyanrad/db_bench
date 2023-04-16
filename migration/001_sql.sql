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
    "workexperience" JSONB,
    "education" JSONB,
    "certifications" JSONB,
    "lookingfor" JSONB,
    "awards" JSONB,
    "languages" JSONB,
    "publication" JSONB,

    CONSTRAINT "applicant_pkey" PRIMARY KEY ("id")
);


-- CreateTable
CREATE TABLE "company" (
    "id" TEXT NOT NULL,
    "name" TEXT NOT NULL,
    "rating" DOUBLE PRECISION NOT NULL,
    "latitude" DOUBLE PRECISION NOT NULL,
    "longitude" DOUBLE PRECISION NOT NULL,
    "category" "companyCategory" NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "company_pkey" PRIMARY KEY ("id")
);
