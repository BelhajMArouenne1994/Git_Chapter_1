CREATE TABLE "recipient" (
  "ocdMasterId" varchar(255) UNIQUE PRIMARY KEY NOT NULL,
  "username" varchar,
  "role" varchar,
  "created_at" timestamp DEFAULT (now())
);

CREATE TABLE "surveyUrl" (
  "brand" varchar(64) NOT NULL,
  "country" varchar(3) NOT NULL,
  "ocdB2cSurveyUrlId" varchar(255) UNIQUE PRIMARY KEY NOT NULL,
  "technicalCreationDate" timestamp NOT NULL,
  "technicalLastUpdateDate" timestamp NOT NULL,
  "sourceName" varchar(255) NOT NULL,
  "sourceSurveyId" varchar(255) NOT NULL,
  "surveyId" varchar(255) NOT NULL,
  "scenario" varchar(255) NOT NULL,
  "channel" varchar(255) NOT NULL,
  "surveyLanguage" varchar(255) NOT NULL,
  "ocdMasterId" varchar(255),
  "relatedObjectName" varchar(255),
  "relatedObjectId" varchar(255) UNIQUE,
  "url" varchar(2048)
);

CREATE TABLE "surveyResponse" (
  "brand" varchar(64) NOT NULL,
  "country" varchar(3) NOT NULL,
  "ocdB2cSurveyResponseId" varchar(255) UNIQUE PRIMARY KEY NOT NULL,
  "ocdB2cSurveyUrlId" varchar(255) UNIQUE NOT NULL,
  "ocdMasterId" varchar(255) NOT NULL,
  "technicalCreationDate" timestamp NOT NULL,
  "technicalLastUpdateDate" timestamp NOT NULL,
  "sourceName" varchar(255) NOT NULL,
  "url" varchar(2048),
  "surveyId" varchar(255) NOT NULL,
  "surveyStatus" varchar(255) NOT NULL,
  "Scenario" varchar(255) NOT NULL,
  "Channel" varchar(255) NOT NULL,
  "relatedObjectName" varchar(255),
  "relatedObjectId" varchar(255) UNIQUE NOT NULL,
  "answerDate" timestamp,
  "endDate" timestamp NOT NULL,
  "npsSegmentation" varchar(255)
);

CREATE TABLE "salesHistory" (
  "brand" varchar,
  "country" varchar,
  "ocdTicketId" varchar UNIQUE PRIMARY KEY NOT NULL,
  "technicalCreationDate" timestamp,
  "technicalLastUpdateDate" timestamp,
  "source" varchar,
  "sourceName" varchar,
  "sourceChannel" varchar,
  "sourcePersonId" varchar,
  "sourceTicketNumber" varchar,
  "sourceStoreType" varchar,
  "sourceStatusOrder" varchar,
  "ocdContactMasterId" varchar,
  "ocdContactVersionId" varchar,
  "ocdStoreId" varchar
);


ALTER TABLE "salesHistory" ADD FOREIGN KEY ("ocdContactMasterId") REFERENCES "recipient" ("ocdMasterId");

ALTER TABLE "surveyUrl" ADD FOREIGN KEY ("ocdMasterId") REFERENCES "recipient" ("ocdMasterId");
ALTER TABLE "surveyUrl" ADD FOREIGN KEY ("relatedObjectId") REFERENCES "salesHistory" ("ocdTicketId");

ALTER TABLE "surveyResponse" ADD FOREIGN KEY ("ocdMasterId") REFERENCES "recipient" ("ocdMasterId");
ALTER TABLE "surveyResponse" ADD FOREIGN KEY ("relatedObjectId") REFERENCES "salesHistory" ("ocdTicketId");
ALTER TABLE "surveyResponse" ADD FOREIGN KEY ("ocdB2cSurveyUrlId") REFERENCES "surveyUrl" ("ocdB2cSurveyUrlId");
