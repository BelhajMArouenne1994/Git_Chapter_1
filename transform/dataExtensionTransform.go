package transform

import (
    "time"

    sfmc_models "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_to_rest_sfmc/models"
    mongo_models "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/mongo_db/models"
)

// Helper function to parse dates
func parseDate(dateStr string) (time.Time, error) {
    layout := "2006-01-02T15:04:05.000Z"
    return time.Parse(layout, dateStr)
}

func TransformToDataExtensionMongoDB(response *sfmc_models.RetrieveDEResponseMsg) ([]*mongo_models.DataExtensionMongoDB, error) {
    var dataExtensions []*mongo_models.DataExtensionMongoDB

    for _, result := range response.Results {
        createdDate, err := parseDate(result.CreatedDate)
        if err != nil {
            return nil, err
        }
        modifiedDate, err := parseDate(result.ModifiedDate)
        if err != nil {
            return nil, err
        }
        //retainUntil, err := parseDate(result.RetainUntil)
        //if err != nil {
        //    return nil, err
        //}

        dataExtension := &mongo_models.DataExtensionMongoDB{
            ObjectID:                     result.ObjectID,
            Name:                         result.Name,
            //Description:                  result.Description,
            IsSendable:                   result.IsSendable,
            //IsTestable:                   result.IsTestable,
            //SendableDataExtensionField:   result.SendableDataExtensionField,
            //SendableSubscriberField:      result.SendableSubscriberField,
            //Template:                     result.Template,
            //DataRetentionPeriodLength:    result.DataRetentionPeriodLength,
            //DataRetentionPeriodUnitOfMeasure: result.DataRetentionPeriodUnitOfMeasure,
            //RowBasedRetention:            result.RowBasedRetention,
            //ResetRetentionPeriodOnImport: result.ResetRetentionPeriodOnImport,
            //DeleteAtEndOfRetentionPeriod: result.DeleteAtEndOfRetentionPeriod,
            //RetainUntil:                  retainUntil,
            //DataRetentionPeriod:          result.DataRetentionPeriod,
            //CategoryID:                   result.CategoryID,
            Status:                       result.Status,
            CreatedDate:                  createdDate,
            ModifiedDate:                 modifiedDate,
            //ClientID:                     result.Client.ID,
        }
        dataExtensions = append(dataExtensions, dataExtension)
    }

    return dataExtensions, nil
}

// Add more transformation functions as needed for other data types
