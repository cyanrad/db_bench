-- CreateEnum
CREATE TYPE "jobType" AS ENUM ('FULL_TIME', 'PART_TIME', 'INTERNSHIP');

-- CreateEnum
CREATE TYPE "applicationStatus" AS ENUM ('REVOKED', 'REJECTED', 'APPLIED', 'SHORTLISTED', 'INTERVIEW', 'HIRED');

-- CreateEnum
CREATE TYPE "state" AS ENUM ('DUBAI', 'AJMAN', 'SHARJAH', 'ABU_DHABI', 'UMM_AL_QUWAIN', 'FUJAIRAH', 'RAS_AL_KHAIMAH');

-- CreateEnum
CREATE TYPE "locationType" AS ENUM ('ONSITE', 'REMOTE', 'FLEXIBLE');

-- CreateEnum
CREATE TYPE "qualification" AS ENUM ('NONE', 'DIPLOMA', 'ASSOCIATE', 'BACHELORS', 'MASTERS', 'DOCTORATE');

-- CreateEnum
CREATE TYPE "languageFluency" AS ENUM ('NOVICE', 'INTERMEDIATE', 'PROFESSIONAL', 'FLUENT', 'NATIVE');

-- CreateEnum
CREATE TYPE "companyCategory" AS ENUM ('GOVERNMENT', 'SEMI_GOVERNMENT', 'PRIVATE');

-- CreateTable
CREATE TABLE "user" (
    "id" TEXT NOT NULL,
    "username" TEXT NOT NULL,
    "uaePassUuid" TEXT,
    "email" TEXT,
    "phoneNumber" TEXT,
    "isAdmin" BOOLEAN NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "user_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "applicant" (
    "id" TEXT NOT NULL,
    "fname" TEXT NOT NULL,
    "lname" TEXT NOT NULL,
    "gender" TEXT NOT NULL,
    "title" TEXT NOT NULL,
    "headline" TEXT,
    "state" "state" NOT NULL,
    "highestQulification" "qualification" NOT NULL,
    "domainYearsOfExperience" INTEGER NOT NULL,
    "workExperience" JSONB[],
    "education" JSONB[],
    "certifications" JSONB[],
    "lookingFor" JSONB[],
    "awards" JSONB[],
    "languages" JSONB[],
    "publication" JSONB[],

    CONSTRAINT "applicant_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "employer" (
    "id" TEXT NOT NULL,

    CONSTRAINT "employer_pkey" PRIMARY KEY ("id")
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

-- CreateTable
CREATE TABLE "opportunity" (
    "id" TEXT NOT NULL,
    "isActive" BOOLEAN NOT NULL,
    "title" TEXT NOT NULL,
    "description" TEXT NOT NULL,
    "responsibilities" TEXT NOT NULL,
    "qualifications" TEXT NOT NULL,
    "companyId" TEXT NOT NULL,
    "employerId" TEXT NOT NULL,
    "locationType" "locationType" NOT NULL,
    "state" "state" NOT NULL,
    "latitude" DOUBLE PRECISION NOT NULL,
    "longitude" DOUBLE PRECISION NOT NULL,
    "jobType" "jobType" NOT NULL,
    "salaryMin" INTEGER,
    "salaryMax" INTEGER,
    "jobDomain" TEXT NOT NULL,
    "minQualification" "qualification" NOT NULL,
    "yearsOfExperience" INTEGER NOT NULL,
    "vacancyCount" INTEGER NOT NULL,
    "genderPreference" BOOLEAN,
    "requiresGpa" BOOLEAN NOT NULL,
    "startsAt" TIMESTAMP(3) NOT NULL,
    "deadlineAt" TIMESTAMP(3) NOT NULL,
    "createdAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updatedAt" TIMESTAMP(3) NOT NULL,

    CONSTRAINT "opportunity_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "application" (
    "id" SERIAL NOT NULL,
    "progress" "applicationStatus" NOT NULL,
    "appliedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "revokedAt" TIMESTAMP(3) NOT NULL,
    "revokedStage" "applicationStatus" NOT NULL,
    "opportunityId" TEXT NOT NULL,
    "profileId" TEXT NOT NULL,

    CONSTRAINT "application_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "review" (
    "id" SERIAL NOT NULL,
    "rating" INTEGER NOT NULL,
    "body" TEXT NOT NULL,
    "publishedAt" TIMESTAMP(3) NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "opportunityId" TEXT NOT NULL,
    "profileId" TEXT NOT NULL,

    CONSTRAINT "review_pkey" PRIMARY KEY ("id")
);

-- CreateTable
CREATE TABLE "_db_CompanyTodb_Employer" (
    "A" TEXT NOT NULL,
    "B" TEXT NOT NULL
);

-- CreateIndex
CREATE UNIQUE INDEX "user_id_key" ON "user"("id");

-- CreateIndex
CREATE UNIQUE INDEX "user_uaePassUuid_key" ON "user"("uaePassUuid");

-- CreateIndex
CREATE UNIQUE INDEX "user_email_key" ON "user"("email");

-- CreateIndex
CREATE UNIQUE INDEX "user_phoneNumber_key" ON "user"("phoneNumber");

-- CreateIndex
CREATE UNIQUE INDEX "applicant_id_key" ON "applicant"("id");

-- CreateIndex
CREATE UNIQUE INDEX "employer_id_key" ON "employer"("id");

-- CreateIndex
CREATE UNIQUE INDEX "company_id_key" ON "company"("id");

-- CreateIndex
CREATE UNIQUE INDEX "opportunity_id_key" ON "opportunity"("id");

-- CreateIndex
CREATE UNIQUE INDEX "_db_CompanyTodb_Employer_AB_unique" ON "_db_CompanyTodb_Employer"("A", "B");

-- CreateIndex
CREATE INDEX "_db_CompanyTodb_Employer_B_index" ON "_db_CompanyTodb_Employer"("B");

-- AddForeignKey
ALTER TABLE "applicant" ADD CONSTRAINT "applicant_id_fkey" FOREIGN KEY ("id") REFERENCES "user"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "employer" ADD CONSTRAINT "employer_id_fkey" FOREIGN KEY ("id") REFERENCES "user"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "opportunity" ADD CONSTRAINT "opportunity_companyId_fkey" FOREIGN KEY ("companyId") REFERENCES "company"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "opportunity" ADD CONSTRAINT "opportunity_employerId_fkey" FOREIGN KEY ("employerId") REFERENCES "employer"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "application" ADD CONSTRAINT "application_opportunityId_fkey" FOREIGN KEY ("opportunityId") REFERENCES "opportunity"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "application" ADD CONSTRAINT "application_profileId_fkey" FOREIGN KEY ("profileId") REFERENCES "applicant"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "review" ADD CONSTRAINT "review_opportunityId_fkey" FOREIGN KEY ("opportunityId") REFERENCES "opportunity"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "review" ADD CONSTRAINT "review_profileId_fkey" FOREIGN KEY ("profileId") REFERENCES "applicant"("id") ON DELETE RESTRICT ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_db_CompanyTodb_Employer" ADD CONSTRAINT "_db_CompanyTodb_Employer_A_fkey" FOREIGN KEY ("A") REFERENCES "company"("id") ON DELETE CASCADE ON UPDATE CASCADE;

-- AddForeignKey
ALTER TABLE "_db_CompanyTodb_Employer" ADD CONSTRAINT "_db_CompanyTodb_Employer_B_fkey" FOREIGN KEY ("B") REFERENCES "employer"("id") ON DELETE CASCADE ON UPDATE CASCADE;

---- create above / drop below ----
