package db

import (
    "context"
    "testing"
	"time"

	"github.com/stretchr/testify/require"
    "github.com/google/uuid"
)


func TestCreateSurveyUrl(t *testing.T) {
	surveyUrlParams := CreateSurveyUrlParams{
		Brand:                    uuid.New().String(),
		Country:                 "USA",
		OcdB2cSurveyUrlId:       "SURV123456789",
		TechnicalCreationDate:   time.Now(),
		TechnicalLastUpdateDate: time.Now(),
		SourceName:              "SurveySource",
		SourceSurveyId:          "SRC123456",
		SurveyId:                "SID987654",
		Scenario:                "Product Feedback",
		Channel:                 "Online",
		SurveyLanguage:          "EN",
		OcdMasterId:             "",
		RelatedObjectName:       "Product",
		RelatedObjectId:         "PRD123456",
		Url:                     "http://www.example.com/survey",
	}

	createdSurveyUrl, err := testQueries.CreateSurveyUrl(context.Background(), surveyUrlParams)

	require.NoError(t, err)
	require.NotEmpty(t, createdSurveyUrl)
	require.NotEmpty(t, createdSurveyUrl.OcdB2cSurveyUrlId)
	require.Equal(t, createdSurveyUrl.OcdB2cSurveyUrlId, surveyUrlParams.OcdB2cSurveyUrlId)
}