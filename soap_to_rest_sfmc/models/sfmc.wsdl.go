package models

import (
	"encoding/xml"
	"time"

	//util "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/util"
	//types "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/types"
	gosoap "github.com/hooklift/gowsdl/soap"
	//gosoap  "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/util" //for debugging
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type Instanceid string

type AsyncResponseType string

const (
	AsyncResponseTypeNone AsyncResponseType = "None"

	AsyncResponseTypeEmail AsyncResponseType = "email"

	AsyncResponseTypeFTP AsyncResponseType = "FTP"

	AsyncResponseTypeHTTPPost AsyncResponseType = "HTTPPost"
)

type Priority string

const (
	PriorityLow Priority = "Low"

	PriorityMedium Priority = "Medium"

	PriorityHigh Priority = "High"
)

type RequestType string

const (
	RequestTypeSynchronous RequestType = "Synchronous"

	RequestTypeAsynchronous RequestType = "Asynchronous"
)

type RespondWhen string

const (
	RespondWhenNever RespondWhen = "Never"

	RespondWhenOnError RespondWhen = "OnError"

	RespondWhenAlways RespondWhen = "Always"

	RespondWhenOnConversationError RespondWhen = "OnConversationError"

	RespondWhenOnConversationComplete RespondWhen = "OnConversationComplete"

	RespondWhenOnCallComplete RespondWhen = "OnCallComplete"

	RespondWhenOnExpired RespondWhen = "OnExpired"
)

type SaveAction string

const (
	SaveActionAddOnly SaveAction = "AddOnly"

	SaveActionDefault SaveAction = "Default"

	SaveActionNothing SaveAction = "Nothing"

	SaveActionUpdateAdd SaveAction = "UpdateAdd"

	SaveActionUpdateOnly SaveAction = "UpdateOnly"

	SaveActionDelete SaveAction = "Delete"
)

type SimpleOperators string

const (
	SimpleOperatorsEquals SimpleOperators = "equals"

	SimpleOperatorsNotEquals SimpleOperators = "notEquals"

	SimpleOperatorsGreaterThan SimpleOperators = "greaterThan"

	SimpleOperatorsLessThan SimpleOperators = "lessThan"

	SimpleOperatorsIsNull SimpleOperators = "isNull"

	SimpleOperatorsIsNotNull SimpleOperators = "isNotNull"

	SimpleOperatorsGreaterThanOrEqual SimpleOperators = "greaterThanOrEqual"

	SimpleOperatorsLessThanOrEqual SimpleOperators = "lessThanOrEqual"

	SimpleOperatorsBetween SimpleOperators = "between"

	SimpleOperatorsIN SimpleOperators = "IN"

	SimpleOperatorsLike SimpleOperators = "like"

	SimpleOperatorsExistsInString SimpleOperators = "existsInString"

	SimpleOperatorsExistsInStringAsAWord SimpleOperators = "existsInStringAsAWord"

	SimpleOperatorsNotExistsInString SimpleOperators = "notExistsInString"

	SimpleOperatorsBeginsWith SimpleOperators = "beginsWith"

	SimpleOperatorsEndsWith SimpleOperators = "endsWith"

	SimpleOperatorsContains SimpleOperators = "contains"

	SimpleOperatorsNotContains SimpleOperators = "notContains"

	SimpleOperatorsIsAnniversary SimpleOperators = "isAnniversary"

	SimpleOperatorsIsNotAnniversary SimpleOperators = "isNotAnniversary"

	SimpleOperatorsGreaterThanAnniversary SimpleOperators = "greaterThanAnniversary"

	SimpleOperatorsLessThanAnniversary SimpleOperators = "lessThanAnniversary"
)

type LogicalOperator string

var (
	LogicalOperatorOR LogicalOperator = "OR"

	LogicalOperatorAND LogicalOperator = "AND"
)

type SoapType string

const (
	SoapTypeXsdstring SoapType = "xsd:string"

	SoapTypeXsdboolean SoapType = "xsd:boolean"

	SoapTypeXsddouble SoapType = "xsd:double"

	SoapTypeXsddateTime SoapType = "xsd:dateTime"
)

type PropertyType string

const (
	PropertyTypeString PropertyType = "string"

	PropertyTypeBoolean PropertyType = "boolean"

	PropertyTypeDouble PropertyType = "double"

	PropertyTypeDatetime PropertyType = "datetime"
)

type RecurrenceTypeEnum string

const (
	RecurrenceTypeEnumSecondly RecurrenceTypeEnum = "Secondly"

	RecurrenceTypeEnumMinutely RecurrenceTypeEnum = "Minutely"

	RecurrenceTypeEnumHourly RecurrenceTypeEnum = "Hourly"

	RecurrenceTypeEnumDaily RecurrenceTypeEnum = "Daily"

	RecurrenceTypeEnumWeekly RecurrenceTypeEnum = "Weekly"

	RecurrenceTypeEnumMonthly RecurrenceTypeEnum = "Monthly"

	RecurrenceTypeEnumYearly RecurrenceTypeEnum = "Yearly"
)

type RecurrenceRangeTypeEnum string

const (
	RecurrenceRangeTypeEnumEndAfter RecurrenceRangeTypeEnum = "EndAfter"

	RecurrenceRangeTypeEnumEndOn RecurrenceRangeTypeEnum = "EndOn"
)

type MinutelyRecurrencePatternTypeEnum string

const (
	MinutelyRecurrencePatternTypeEnumInterval MinutelyRecurrencePatternTypeEnum = "Interval"
)

type HourlyRecurrencePatternTypeEnum string

const (
	HourlyRecurrencePatternTypeEnumInterval HourlyRecurrencePatternTypeEnum = "Interval"
)

type DailyRecurrencePatternTypeEnum string

const (
	DailyRecurrencePatternTypeEnumInterval DailyRecurrencePatternTypeEnum = "Interval"

	DailyRecurrencePatternTypeEnumEveryWeekDay DailyRecurrencePatternTypeEnum = "EveryWeekDay"
)

type WeeklyRecurrencePatternTypeEnum string

const (
	WeeklyRecurrencePatternTypeEnumByDay WeeklyRecurrencePatternTypeEnum = "ByDay"
)

type MonthlyRecurrencePatternTypeEnum string

const (
	MonthlyRecurrencePatternTypeEnumByDay MonthlyRecurrencePatternTypeEnum = "ByDay"

	MonthlyRecurrencePatternTypeEnumByWeek MonthlyRecurrencePatternTypeEnum = "ByWeek"
)

type WeekOfMonthEnum string

const (
	WeekOfMonthEnumFirst WeekOfMonthEnum = "first"

	WeekOfMonthEnumSecond WeekOfMonthEnum = "second"

	WeekOfMonthEnumThird WeekOfMonthEnum = "third"

	WeekOfMonthEnumFourth WeekOfMonthEnum = "fourth"

	WeekOfMonthEnumLast WeekOfMonthEnum = "last"
)

type DayOfWeekEnum string

const (
	DayOfWeekEnumSunday DayOfWeekEnum = "Sunday"

	DayOfWeekEnumMonday DayOfWeekEnum = "Monday"

	DayOfWeekEnumTuesday DayOfWeekEnum = "Tuesday"

	DayOfWeekEnumWednesday DayOfWeekEnum = "Wednesday"

	DayOfWeekEnumThursday DayOfWeekEnum = "Thursday"

	DayOfWeekEnumFriday DayOfWeekEnum = "Friday"

	DayOfWeekEnumSaturday DayOfWeekEnum = "Saturday"
)

type YearlyRecurrencePatternTypeEnum string

const (
	YearlyRecurrencePatternTypeEnumByDay YearlyRecurrencePatternTypeEnum = "ByDay"

	YearlyRecurrencePatternTypeEnumByWeek YearlyRecurrencePatternTypeEnum = "ByWeek"

	YearlyRecurrencePatternTypeEnumByMonth YearlyRecurrencePatternTypeEnum = "ByMonth"
)

type MonthOfYearEnum string

const (
	MonthOfYearEnumJanuary MonthOfYearEnum = "January"

	MonthOfYearEnumFebruary MonthOfYearEnum = "February"

	MonthOfYearEnumMarch MonthOfYearEnum = "March"

	MonthOfYearEnumApril MonthOfYearEnum = "April"

	MonthOfYearEnumMay MonthOfYearEnum = "May"

	MonthOfYearEnumJune MonthOfYearEnum = "June"

	MonthOfYearEnumJuly MonthOfYearEnum = "July"

	MonthOfYearEnumAugust MonthOfYearEnum = "August"

	MonthOfYearEnumSeptember MonthOfYearEnum = "September"

	MonthOfYearEnumOctober MonthOfYearEnum = "October"

	MonthOfYearEnumNovember MonthOfYearEnum = "November"

	MonthOfYearEnumDecember MonthOfYearEnum = "December"
)

type ExtractParameterDataType string

const (
	ExtractParameterDataTypeDatetime ExtractParameterDataType = "datetime"

	ExtractParameterDataTypeBool ExtractParameterDataType = "bool"

	ExtractParameterDataTypeString ExtractParameterDataType = "string"

	ExtractParameterDataTypeInteger ExtractParameterDataType = "integer"

	ExtractParameterDataTypeDropdown ExtractParameterDataType = "dropdown"
)

type UnsubscribeBehaviorEnum string

const (
	UnsubscribeBehaviorEnumENTIRE_ENTERPRISE UnsubscribeBehaviorEnum = "ENTIRE_ENTERPRISE"

	UnsubscribeBehaviorEnumBUSINESS_UNIT_ONLY UnsubscribeBehaviorEnum = "BUSINESS_UNIT_ONLY"
)

type AccountTypeEnum string

const (
	AccountTypeEnumNone AccountTypeEnum = "None"

	AccountTypeEnumEXACTTARGET AccountTypeEnum = "EXACTTARGET"

	AccountTypeEnumPRO_CONNECT AccountTypeEnum = "PRO_CONNECT"

	AccountTypeEnumCHANNEL_CONNECT AccountTypeEnum = "CHANNEL_CONNECT"

	AccountTypeEnumCONNECT AccountTypeEnum = "CONNECT"

	AccountTypeEnumPRO_CONNECT_CLIENT AccountTypeEnum = "PRO_CONNECT_CLIENT"

	AccountTypeEnumLP_MEMBER AccountTypeEnum = "LP_MEMBER"

	AccountTypeEnumDOTO_MEMBER AccountTypeEnum = "DOTO_MEMBER"

	AccountTypeEnumENTERPRISE_2 AccountTypeEnum = "ENTERPRISE_2"

	AccountTypeEnumBUSINESS_UNIT AccountTypeEnum = "BUSINESS_UNIT"
)

type LayoutType string

const (
	LayoutTypeHTMLWrapped LayoutType = "HTMLWrapped"

	LayoutTypeRawText LayoutType = "RawText"

	LayoutTypeSMS LayoutType = "SMS"
)

type EventType string

const (
	EventTypeOpen EventType = "Open"

	EventTypeClick EventType = "Click"

	EventTypeHardBounce EventType = "HardBounce"

	EventTypeSoftBounce EventType = "SoftBounce"

	EventTypeOtherBounce EventType = "OtherBounce"

	EventTypeUnsubscribe EventType = "Unsubscribe"

	EventTypeSent EventType = "Sent"

	EventTypeNotSent EventType = "NotSent"

	EventTypeSurvey EventType = "Survey"

	EventTypeForwardedEmail EventType = "ForwardedEmail"

	EventTypeForwardedEmailOptIn EventType = "ForwardedEmailOptIn"

	EventTypeDeliveredEvent EventType = "DeliveredEvent"
)

type CompressionType string

const (
	CompressionTypeGzip CompressionType = "gzip"
)

type CompressionEncoding string

const (
	CompressionEncodingBase64 CompressionEncoding = "base64"
)

type SubscriberStatus string

const (
	SubscriberStatusActive SubscriberStatus = "Active"

	SubscriberStatusBounced SubscriberStatus = "Bounced"

	SubscriberStatusHeld SubscriberStatus = "Held"

	SubscriberStatusUnsubscribed SubscriberStatus = "Unsubscribed"

	SubscriberStatusDeleted SubscriberStatus = "Deleted"
)

type EmailType string

const (
	EmailTypeText EmailType = "Text"

	EmailTypeHTML EmailType = "HTML"
)

type ListTypeEnum string

const (
	ListTypeEnumPublic ListTypeEnum = "Public"

	ListTypeEnumPrivate ListTypeEnum = "Private"

	ListTypeEnumSalesForce ListTypeEnum = "SalesForce"

	ListTypeEnumGlobalUnsubscribe ListTypeEnum = "GlobalUnsubscribe"

	ListTypeEnumMaster ListTypeEnum = "Master"
)

type ListClassificationEnum string

const (
	ListClassificationEnumExactTargetList ListClassificationEnum = "ExactTargetList"

	ListClassificationEnumPublicationList ListClassificationEnum = "PublicationList"

	ListClassificationEnumSuppressionList ListClassificationEnum = "SuppressionList"
)

type OverrideType string

const (
	OverrideTypeDoNotOverride OverrideType = "DoNotOverride"

	OverrideTypeOverride OverrideType = "Override"

	OverrideTypeOverrideExceptWhenNull OverrideType = "OverrideExceptWhenNull"
)

type ListAttributeFieldType string

const (
	ListAttributeFieldTypeText ListAttributeFieldType = "Text"

	ListAttributeFieldTypeNumber ListAttributeFieldType = "Number"

	ListAttributeFieldTypeDate ListAttributeFieldType = "Date"

	ListAttributeFieldTypeBoolean ListAttributeFieldType = "Boolean"

	ListAttributeFieldTypeDecimal ListAttributeFieldType = "Decimal"
)

type TriggeredSendTypeEnum string

const (
	TriggeredSendTypeEnumContinuous TriggeredSendTypeEnum = "Continuous"

	TriggeredSendTypeEnumBatched TriggeredSendTypeEnum = "Batched"

	TriggeredSendTypeEnumScheduled TriggeredSendTypeEnum = "Scheduled"
)

type TriggeredSendStatusEnum string

const (
	TriggeredSendStatusEnumNew TriggeredSendStatusEnum = "New"

	TriggeredSendStatusEnumInactive TriggeredSendStatusEnum = "Inactive"

	TriggeredSendStatusEnumActive TriggeredSendStatusEnum = "Active"

	TriggeredSendStatusEnumCanceled TriggeredSendStatusEnum = "Canceled"

	TriggeredSendStatusEnumDeleted TriggeredSendStatusEnum = "Deleted"

	TriggeredSendStatusEnumMoved TriggeredSendStatusEnum = "Moved"
)

type TriggeredSendClassEnum string

const (
	TriggeredSendClassEnumStandard TriggeredSendClassEnum = "Standard"

	TriggeredSendClassEnumSMTPRestV1 TriggeredSendClassEnum = "SMTPRestV1"

	TriggeredSendClassEnumSMTPRestV2 TriggeredSendClassEnum = "SMTPRestV2"

	TriggeredSendClassEnumSMTPRestV3 TriggeredSendClassEnum = "SMTPRestV3"
)

type TriggeredSendSubClassEnum string

const (
	TriggeredSendSubClassEnumStandard TriggeredSendSubClassEnum = "Standard"

	TriggeredSendSubClassEnumSenderTemplate TriggeredSendSubClassEnum = "SenderTemplate"

	TriggeredSendSubClassEnumMarketingCloudTemplate TriggeredSendSubClassEnum = "MarketingCloudTemplate"
)

type SendClassificationTypeEnum string

const (
	SendClassificationTypeEnumOperational SendClassificationTypeEnum = "Operational"

	SendClassificationTypeEnumMarketing SendClassificationTypeEnum = "Marketing"
)

type SendPriorityEnum string

const (
	SendPriorityEnumBurst SendPriorityEnum = "Burst"

	SendPriorityEnumNormal SendPriorityEnum = "Normal"

	SendPriorityEnumLow SendPriorityEnum = "Low"
)

type DeliveryProfileSourceAddressTypeEnum string

const (
	DeliveryProfileSourceAddressTypeEnumDefaultPrivateIPAddress DeliveryProfileSourceAddressTypeEnum = "DefaultPrivateIPAddress"

	DeliveryProfileSourceAddressTypeEnumCustomPrivateIPAddress DeliveryProfileSourceAddressTypeEnum = "CustomPrivateIPAddress"
)

type DeliveryProfileDomainTypeEnum string

const (
	DeliveryProfileDomainTypeEnumDefaultDomain DeliveryProfileDomainTypeEnum = "DefaultDomain"

	DeliveryProfileDomainTypeEnumCustomDomain DeliveryProfileDomainTypeEnum = "CustomDomain"
)

type SalutationSourceEnum string

const (
	SalutationSourceEnumDefault SalutationSourceEnum = "Default"

	SalutationSourceEnumContentLibrary SalutationSourceEnum = "ContentLibrary"

	SalutationSourceEnumNone SalutationSourceEnum = "None"
)

type SendDefinitionStatusEnum string

const (
	SendDefinitionStatusEnumActive SendDefinitionStatusEnum = "Active"

	SendDefinitionStatusEnumArchived SendDefinitionStatusEnum = "Archived"

	SendDefinitionStatusEnumDeleted SendDefinitionStatusEnum = "Deleted"
)

type SendDefinitionListTypeEnum string

const (
	SendDefinitionListTypeEnumSourceList SendDefinitionListTypeEnum = "SourceList"

	SendDefinitionListTypeEnumExclusionList SendDefinitionListTypeEnum = "ExclusionList"

	SendDefinitionListTypeEnumDomainExclusion SendDefinitionListTypeEnum = "DomainExclusion"

	SendDefinitionListTypeEnumOptOutList SendDefinitionListTypeEnum = "OptOutList"
)

type DataSourceTypeEnum string

const (
	DataSourceTypeEnumList DataSourceTypeEnum = "List"

	DataSourceTypeEnumCustomObject DataSourceTypeEnum = "CustomObject"

	DataSourceTypeEnumDomainExclusion DataSourceTypeEnum = "DomainExclusion"

	DataSourceTypeEnumSalesForceReport DataSourceTypeEnum = "SalesForceReport"

	DataSourceTypeEnumSalesForceCampaign DataSourceTypeEnum = "SalesForceCampaign"

	DataSourceTypeEnumFilterDefinition DataSourceTypeEnum = "FilterDefinition"

	DataSourceTypeEnumOptOutList DataSourceTypeEnum = "OptOutList"
)

type MessageDeliveryTypeEnum string

const (
	MessageDeliveryTypeEnumStandard MessageDeliveryTypeEnum = "Standard"

	MessageDeliveryTypeEnumDelayedDeliveryByMTAQueue MessageDeliveryTypeEnum = "DelayedDeliveryByMTAQueue"

	MessageDeliveryTypeEnumDelayedDeliveryByOMMQueue MessageDeliveryTypeEnum = "DelayedDeliveryByOMMQueue"
)

type ChatMessagingEventType string

const (
	ChatMessagingEventTypeTemplate ChatMessagingEventType = "Template"

	ChatMessagingEventTypeMO ChatMessagingEventType = "MO"

	ChatMessagingEventTypeWhatsAppTemplate ChatMessagingEventType = "WhatsAppTemplate"

	ChatMessagingEventTypeWhatsAppMO ChatMessagingEventType = "WhatsAppMO"

	ChatMessagingEventTypeWhatsAppPhone ChatMessagingEventType = "WhatsAppPhone"

	ChatMessagingEventTypeWhatsAppAccount ChatMessagingEventType = "WhatsAppAccount"

	ChatMessagingEventTypeE360WhatsApp ChatMessagingEventType = "E360WhatsApp"
)

type DataExtensionFieldType string

const (
	DataExtensionFieldTypeText DataExtensionFieldType = "Text"

	DataExtensionFieldTypeNumber DataExtensionFieldType = "Number"

	DataExtensionFieldTypeDate DataExtensionFieldType = "Date"

	DataExtensionFieldTypeBoolean DataExtensionFieldType = "Boolean"

	DataExtensionFieldTypeEmailAddress DataExtensionFieldType = "EmailAddress"

	DataExtensionFieldTypePhone DataExtensionFieldType = "Phone"

	DataExtensionFieldTypeDecimal DataExtensionFieldType = "Decimal"

	DataExtensionFieldTypeLocale DataExtensionFieldType = "Locale"

	DataExtensionFieldTypeBase16Encrypted DataExtensionFieldType = "Base16Encrypted"

	DataExtensionFieldTypeBase16EncryptedEmail DataExtensionFieldType = "Base16EncryptedEmail"
)

type DataExtensionFieldStorageType string

const (
	DataExtensionFieldStorageTypeUnspecified DataExtensionFieldStorageType = "Unspecified"

	DataExtensionFieldStorageTypePlain DataExtensionFieldStorageType = "Plain"

	DataExtensionFieldStorageTypeObfuscated DataExtensionFieldStorageType = "Obfuscated"

	DataExtensionFieldStorageTypeEncrypted DataExtensionFieldStorageType = "Encrypted"
)

type DateTimeUnitOfMeasure string

const (
	DateTimeUnitOfMeasureDays DateTimeUnitOfMeasure = "Days"

	DateTimeUnitOfMeasureWeeks DateTimeUnitOfMeasure = "Weeks"

	DateTimeUnitOfMeasureMonths DateTimeUnitOfMeasure = "Months"

	DateTimeUnitOfMeasureYears DateTimeUnitOfMeasure = "Years"
)

type FileType string

const (
	FileTypeCSV FileType = "CSV"

	FileTypeTAB FileType = "TAB"

	FileTypeOther FileType = "Other"
)

type ImportDefinitionSubscriberImportType string

const (
	ImportDefinitionSubscriberImportTypeEmail ImportDefinitionSubscriberImportType = "Email"

	ImportDefinitionSubscriberImportTypeSMS ImportDefinitionSubscriberImportType = "SMS"
)

type ImportDefinitionUpdateType string

const (
	ImportDefinitionUpdateTypeAddAndUpdate ImportDefinitionUpdateType = "AddAndUpdate"

	ImportDefinitionUpdateTypeAddAndDoNotUpdate ImportDefinitionUpdateType = "AddAndDoNotUpdate"

	ImportDefinitionUpdateTypeUpdateButDoNotAdd ImportDefinitionUpdateType = "UpdateButDoNotAdd"

	ImportDefinitionUpdateTypeMerge ImportDefinitionUpdateType = "Merge"

	ImportDefinitionUpdateTypeOverwrite ImportDefinitionUpdateType = "Overwrite"

	ImportDefinitionUpdateTypeColumnBased ImportDefinitionUpdateType = "ColumnBased"
)

type ImportDefinitionColumnBasedActionType string

const (
	ImportDefinitionColumnBasedActionTypeAddAndUpdate ImportDefinitionColumnBasedActionType = "AddAndUpdate"

	ImportDefinitionColumnBasedActionTypeAddButDoNotUpdate ImportDefinitionColumnBasedActionType = "AddButDoNotUpdate"

	ImportDefinitionColumnBasedActionTypeDelete ImportDefinitionColumnBasedActionType = "Delete"

	ImportDefinitionColumnBasedActionTypeSkip ImportDefinitionColumnBasedActionType = "Skip"

	ImportDefinitionColumnBasedActionTypeUpdateButDoNotAdd ImportDefinitionColumnBasedActionType = "UpdateButDoNotAdd"
)

type ImportDefinitionFieldMappingType string

const (
	ImportDefinitionFieldMappingTypeInferFromColumnHeadings ImportDefinitionFieldMappingType = "InferFromColumnHeadings"

	ImportDefinitionFieldMappingTypeMapByOrdinal ImportDefinitionFieldMappingType = "MapByOrdinal"

	ImportDefinitionFieldMappingTypeManualMap ImportDefinitionFieldMappingType = "ManualMap"
)

type SystemStatusType string

const (
	SystemStatusTypeOK SystemStatusType = "OK"

	SystemStatusTypeUnplannedOutage SystemStatusType = "UnplannedOutage"

	SystemStatusTypeInMaintenance SystemStatusType = "InMaintenance"
)

type SubscriberAddressStatus string

const (
	SubscriberAddressStatusOptedIn SubscriberAddressStatus = "OptedIn"

	SubscriberAddressStatusOptedOut SubscriberAddressStatus = "OptedOut"

	SubscriberAddressStatusInActive SubscriberAddressStatus = "InActive"
)

type AutomationTestType string

const (
	AutomationTestTypeOK AutomationTestType = "OK"

	AutomationTestTypeUnplannedOutage AutomationTestType = "UnplannedOutage"

	AutomationTestTypeInMaintenance AutomationTestType = "InMaintenance"
)

type AutomationTypes string

const (
	AutomationTypesScheduled AutomationTypes = "scheduled"

	AutomationTypesTriggered AutomationTypes = "triggered"
)

type AutomationSourceTypes string

const (
	AutomationSourceTypesUnknown AutomationSourceTypes = "Unknown"

	AutomationSourceTypesFileTrigger AutomationSourceTypes = "FileTrigger"

	AutomationSourceTypesUserInterface AutomationSourceTypes = "UserInterface"

	AutomationSourceTypesUserAPI AutomationSourceTypes = "UserAPI"

	AutomationSourceTypesRESTApi AutomationSourceTypes = "RESTApi"
)

type AutomationStatus string

const (
	AutomationStatusError AutomationStatus = "Error"

	AutomationStatusBuildingError AutomationStatus = "BuildingError"

	AutomationStatusBuilding AutomationStatus = "Building"

	AutomationStatusReady AutomationStatus = "Ready"

	AutomationStatusRunning AutomationStatus = "Running"

	AutomationStatusPaused AutomationStatus = "Paused"

	AutomationStatusStopped AutomationStatus = "Stopped"

	AutomationStatusScheduled AutomationStatus = "Scheduled"

	AutomationStatusAwaitingTrigger AutomationStatus = "AwaitingTrigger"

	AutomationStatusInactiveTrigger AutomationStatus = "InactiveTrigger"

	AutomationStatusSkipped AutomationStatus = "Skipped"

	AutomationStatusUnknown AutomationStatus = "Unknown"

	AutomationStatusNew AutomationStatus = "New"
)

type SuppressionListContextEnum string

const (
	SuppressionListContextEnumEnterprise SuppressionListContextEnum = "Enterprise"

	SuppressionListContextEnumBusinessUnit SuppressionListContextEnum = "BusinessUnit"

	SuppressionListContextEnumSendClassification SuppressionListContextEnum = "SendClassification"

	SuppressionListContextEnumSend SuppressionListContextEnum = "Send"

	SuppressionListContextEnumGlobal SuppressionListContextEnum = "Global"

	SuppressionListContextEnumSenderProfile SuppressionListContextEnum = "SenderProfile"
)

type CreateRequest struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI CreateRequest"`

	Options *CreateOptions `xml:"Options,omitempty"`

	Objects []*APIObject `xml:"Objects,omitempty"`
}

type CreateResponse struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI CreateResponse"`

	Results []*CreateResult `xml:"Results,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	OverallStatus string `xml:"OverallStatus,omitempty"`
}

type UpdateRequest struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI UpdateRequest"`

	Options *UpdateOptions `xml:"Options,omitempty"`

	Objects []*APIObject `xml:"Objects,omitempty"`
}

type UpdateResponse struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI UpdateResponse"`

	Results []*UpdateResult `xml:"Results,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	OverallStatus string `xml:"OverallStatus,omitempty"`
}

type DeleteRequest struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI DeleteRequest"`

	Options *DeleteOptions `xml:"Options,omitempty"`

	Objects []*APIObject `xml:"Objects,omitempty"`
}

type DeleteResponse struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI DeleteResponse"`

	Results []*DeleteResult `xml:"Results,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	OverallStatus string `xml:"OverallStatus,omitempty"`
}

type RetrieveRequestMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI RetrieveRequestMsg"`

	RetrieveRequest *RetrieveRequest `xml:"RetrieveRequest,omitempty"`
}

type RetrieveResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI RetrieveResponseMsg"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	Results []*APIObject `xml:"Results,omitempty"`
}

type QueryRequestMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI QueryRequestMsg"`

	QueryRequest *QueryRequest `xml:"QueryRequest,omitempty"`
}

type QueryResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI QueryResponseMsg"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	Results []*APIObject `xml:"Results,omitempty"`
}

type DefinitionRequestMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI DefinitionRequestMsg"`

	DescribeRequests *ArrayOfObjectDefinitionRequest `xml:"DescribeRequests,omitempty"`
}

type DefinitionResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI DefinitionResponseMsg"`

	ObjectDefinition []*ObjectDefinition `xml:"ObjectDefinition,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`
}

type ExecuteRequestMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI ExecuteRequestMsg"`

	Requests []*ExecuteRequest `xml:"Requests,omitempty"`
}

type ExecuteResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI ExecuteResponseMsg"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	Results []*ExecuteResponse `xml:"Results,omitempty"`
}

type PerformRequestMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI PerformRequestMsg"`

	Options *PerformOptions `xml:"Options,omitempty"`

	Action string `xml:"Action,omitempty"`

	Definitions struct {
		Definition []*APIObject `xml:"Definition,omitempty"`
	} `xml:"Definitions,omitempty"`
}

type PerformResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI PerformResponseMsg"`

	Results struct {
		Result []*PerformResult `xml:"Result,omitempty"`
	} `xml:"Results,omitempty"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	OverallStatusMessage string `xml:"OverallStatusMessage,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`
}

type ConfigureRequestMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI ConfigureRequestMsg"`

	Options *ConfigureOptions `xml:"Options,omitempty"`

	Action string `xml:"Action,omitempty"`

	Configurations struct {
		Configuration []*APIObject `xml:"Configuration,omitempty"`
	} `xml:"Configurations,omitempty"`
}

type ConfigureResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI ConfigureResponseMsg"`

	Results struct {
		Result []*ConfigureResult `xml:"Result,omitempty"`
	} `xml:"Results,omitempty"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	OverallStatusMessage string `xml:"OverallStatusMessage,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`
}

type ScheduleRequestMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI ScheduleRequestMsg"`

	Options *ScheduleOptions `xml:"Options,omitempty"`

	Action string `xml:"Action,omitempty"`

	Schedule *ScheduleDefinition `xml:"Schedule,omitempty"`

	Interactions struct {
		Interaction []*APIObject `xml:"Interaction,omitempty"`
	} `xml:"Interactions,omitempty"`
}

type ScheduleResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI ScheduleResponseMsg"`

	Results struct {
		Result []*ScheduleResult `xml:"Result,omitempty"`
	} `xml:"Results,omitempty"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	OverallStatusMessage string `xml:"OverallStatusMessage,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`
}

type ExtractRequestMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI ExtractRequestMsg"`

	Requests []*ExtractRequest `xml:"Requests,omitempty"`
}

type ExtractResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI ExtractResponseMsg"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	Results struct {
		ExtractResult []*ExtractResult `xml:"ExtractResult,omitempty"`
	} `xml:"Results,omitempty"`
}

type VersionInfoRequestMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI VersionInfoRequestMsg"`

	IncludeVersionHistory bool `xml:"IncludeVersionHistory,omitempty"`
}

type VersionInfoResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI VersionInfoResponseMsg"`

	VersionInfo *VersionInfoResponse `xml:"VersionInfo,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`
}

type SystemStatusRequestMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI SystemStatusRequestMsg"`

	Options *SystemStatusOptions `xml:"Options,omitempty"`
}

type SystemStatusResponseMsg struct {
	XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI SystemStatusResponseMsg"`

	Results struct {
		Result []*SystemStatusResult `xml:"Result,omitempty"`
	} `xml:"Results,omitempty"`

	OverallStatus string `xml:"OverallStatus,omitempty"`

	OverallStatusMessage string `xml:"OverallStatusMessage,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`
}

/*type APIObject struct {
	Client *ClientID `xml:"Client,omitempty"`

	PartnerKey string `xml:"PartnerKey,omitempty"`

	PartnerProperties []*APIProperty `xml:"PartnerProperties,omitempty"`

	CreatedDate time.Time `xml:"CreatedDate,omitempty"`

	ModifiedDate time.Time `xml:"ModifiedDate,omitempty"`

	ID int32 `xml:"ID,omitempty"`

	ObjectID string `xml:"ObjectID,omitempty"`

	CustomerKey string `xml:"CustomerKey,omitempty"`

	Owner *Owner `xml:"Owner,omitempty"`

	CorrelationID string `xml:"CorrelationID,omitempty"`

	ObjectState string `xml:"ObjectState,omitempty"`
}
*/

type APIObject struct {
	Client            *ClientID      `xml:"Client,omitempty" json:"client,omitempty"`
	PartnerKey        string         `xml:"PartnerKey,omitempty" json:"partnerKey,omitempty"`
	PartnerProperties []*APIProperty `xml:"PartnerProperties,omitempty" json:"partnerProperties,omitempty"`
	CreatedDate       string         `xml:"CreatedDate,omitempty" json:"createdDate,omitempty"`
	ModifiedDate      string         `xml:"ModifiedDate,omitempty" json:"modifiedDate,omitempty"`
	ID                int32          `xml:"ID,omitempty" json:"id,omitempty"`
	ObjectID          string         `xml:"ObjectID,omitempty" json:"objectId,omitempty"`
	CustomerKey       string         `xml:"CustomerKey,omitempty" json:"customerKey,omitempty"`
	Owner             *Owner         `xml:"Owner,omitempty" json:"owner,omitempty"`
	CorrelationID     string         `xml:"CorrelationID,omitempty" json:"correlationId,omitempty"`
	ObjectState       string         `xml:"ObjectState,omitempty" json:"objectState,omitempty"`
}

/*
type ClientID struct {

	// Deprecated.  Use ID.
	ClientID int32 `xml:"ClientID,omitempty"`

	ID int32 `xml:"ID,omitempty"`

	PartnerClientKey string `xml:"PartnerClientKey,omitempty"`

	UserID int32 `xml:"UserID,omitempty"`

	PartnerUserKey string `xml:"PartnerUserKey,omitempty"`

	CreatedBy int32 `xml:"CreatedBy,omitempty"`

	ModifiedBy int32 `xml:"ModifiedBy,omitempty"`

	EnterpriseID int64 `xml:"EnterpriseID,omitempty"`

	CustomerKey string `xml:"CustomerKey,omitempty"`

	CustomerID string `xml:"CustomerID,omitempty"`
}
*/

type ClientID struct {
	ClientID         int32  `xml:"ClientID,omitempty" json:"clientId,omitempty"`
	ID               int32  `xml:"ID,omitempty" json:"id,omitempty"`
	PartnerClientKey string `xml:"PartnerClientKey,omitempty" json:"partnerClientKey,omitempty"`
	UserID           int32  `xml:"UserID,omitempty" json:"userId,omitempty"`
	PartnerUserKey   string `xml:"PartnerUserKey,omitempty" json:"partnerUserKey,omitempty"`
	CreatedBy        int32  `xml:"CreatedBy,omitempty" json:"createdBy,omitempty"`
	ModifiedBy       int32  `xml:"ModifiedBy,omitempty" json:"modifiedBy,omitempty"`
	EnterpriseID     int64  `xml:"EnterpriseID,omitempty" json:"enterpriseId,omitempty"`
	CustomerKey      string `xml:"CustomerKey,omitempty" json:"customerKey,omitempty"`
	CustomerID       string `xml:"CustomerID,omitempty" json:"customerId,omitempty"`
}

/*
type APIProperty struct {
	Name string `xml:"Name,omitempty"`

	Value string `xml:"Value,omitempty"`
}
*/

type APIProperty struct {
	Name  string `xml:"Name,omitempty" json:"name,omitempty"`
	Value string `xml:"Value,omitempty" json:"value,omitempty"`
}

type NullAPIProperty struct {
	*APIProperty
}

type DataFolder struct {
	*APIObject

	ParentFolder *DataFolder `xml:"ParentFolder,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	ContentType string `xml:"ContentType,omitempty"`

	IsActive bool `xml:"IsActive,omitempty"`

	IsEditable bool `xml:"IsEditable,omitempty"`

	AllowChildren bool `xml:"AllowChildren,omitempty"`
}

type Owner struct {
	Client *ClientID `xml:"Client,omitempty"`

	FromName string `xml:"FromName,omitempty"`

	FromAddress string `xml:"FromAddress,omitempty"`

	User *AccountUser `xml:"User,omitempty"`
}

type AsyncResponse struct {
	ResponseType *AsyncResponseType `xml:"ResponseType,omitempty"`

	ResponseAddress string `xml:"ResponseAddress,omitempty"`

	RespondWhen *RespondWhen `xml:"RespondWhen,omitempty"`

	IncludeResults bool `xml:"IncludeResults,omitempty"`

	IncludeObjects bool `xml:"IncludeObjects,omitempty"`

	OnlyIncludeBase bool `xml:"OnlyIncludeBase,omitempty"`
}

type ContainerID struct {
	APIObject *APIObject `xml:"APIObject,omitempty"`
}

type Request struct {
}

type Result struct {
	StatusCode string `xml:"StatusCode,omitempty"`

	StatusMessage string `xml:"StatusMessage,omitempty"`

	OrdinalID int32 `xml:"OrdinalID,omitempty"`

	ErrorCode int32 `xml:"ErrorCode,omitempty"`

	RequestID string `xml:"RequestID,omitempty"`

	ConversationID string `xml:"ConversationID,omitempty"`

	OverallStatusCode string `xml:"OverallStatusCode,omitempty"`

	RequestType *RequestType `xml:"RequestType,omitempty"`

	ResultType string `xml:"ResultType,omitempty"`

	ResultDetailXML string `xml:"ResultDetailXML,omitempty"`
}

type ResultMessage struct {
	*APIObject

	RequestID string `xml:"RequestID,omitempty"`

	ConversationID string `xml:"ConversationID,omitempty"`

	OverallStatusCode string `xml:"OverallStatusCode,omitempty"`

	StatusCode string `xml:"StatusCode,omitempty"`

	StatusMessage string `xml:"StatusMessage,omitempty"`

	ErrorCode int32 `xml:"ErrorCode,omitempty"`

	RequestType *RequestType `xml:"RequestType,omitempty"`

	ResultType string `xml:"ResultType,omitempty"`

	ResultDetailXML string `xml:"ResultDetailXML,omitempty"`

	SequenceCode int32 `xml:"SequenceCode,omitempty"`

	CallsInConversation int32 `xml:"CallsInConversation,omitempty"`
}

type ResultItem struct {
	*APIObject

	RequestID string `xml:"RequestID,omitempty"`

	ConversationID string `xml:"ConversationID,omitempty"`

	StatusCode string `xml:"StatusCode,omitempty"`

	StatusMessage string `xml:"StatusMessage,omitempty"`

	OrdinalID int32 `xml:"OrdinalID,omitempty"`

	ErrorCode int32 `xml:"ErrorCode,omitempty"`

	RequestType *RequestType `xml:"RequestType,omitempty"`

	RequestObjectType string `xml:"RequestObjectType,omitempty"`
}

type Options struct {
	Client *ClientID `xml:"Client,omitempty"`

	SendResponseTo []*AsyncResponse `xml:"SendResponseTo,omitempty"`

	SaveOptions struct {
		SaveOption []*SaveOption `xml:"SaveOption,omitempty"`
	} `xml:"SaveOptions,omitempty"`

	Priority int8 `xml:"Priority,omitempty"`

	ConversationID string `xml:"ConversationID,omitempty"`

	SequenceCode int32 `xml:"SequenceCode,omitempty"`

	CallsInConversation int32 `xml:"CallsInConversation,omitempty"`

	ScheduledTime time.Time `xml:"ScheduledTime,omitempty"`

	RequestType *RequestType `xml:"RequestType,omitempty"`

	QueuePriority *Priority `xml:"QueuePriority,omitempty"`

	RequestExpirationTime time.Time `xml:"RequestExpirationTime,omitempty"`
}

type TaskResult struct {
	StatusCode string `xml:"StatusCode,omitempty"`

	StatusMessage string `xml:"StatusMessage,omitempty"`

	OrdinalID int32 `xml:"OrdinalID,omitempty"`

	ErrorCode int32 `xml:"ErrorCode,omitempty"`

	ID string `xml:"ID,omitempty"`

	TblAsyncID int32 `xml:"TblAsyncID,omitempty"`

	InteractionObjectID string `xml:"InteractionObjectID,omitempty"`
}

type SaveOption struct {
	PropertyName string `xml:"PropertyName,omitempty"`

	SaveAction *SaveAction `xml:"SaveAction,omitempty"`

	TrackChanges bool `xml:"TrackChanges,omitempty"`
}

type CreateResult struct {
	*Result

	NewID int32 `xml:"NewID,omitempty"`

	NewObjectID string `xml:"NewObjectID,omitempty"`

	PartnerKey string `xml:"PartnerKey,omitempty"`

	Object *APIObject `xml:"Object,omitempty"`

	CreateResults []*CreateResult `xml:"CreateResults,omitempty"`

	ParentPropertyName string `xml:"ParentPropertyName,omitempty"`
}

type CreateOptions struct {
	*Options

	Container *ContainerID `xml:"Container,omitempty"`
}

type UpdateOptions struct {
	*Options

	Container *ContainerID `xml:"Container,omitempty"`

	Action string `xml:"Action,omitempty"`
}

type UpdateResult struct {
	*Result

	Object *APIObject `xml:"Object,omitempty"`

	UpdateResults []*UpdateResult `xml:"UpdateResults,omitempty"`

	ParentPropertyName string `xml:"ParentPropertyName,omitempty"`
}

type DeleteOptions struct {
	*Options
}

type DeleteResult struct {
	*Result

	Object *APIObject `xml:"Object,omitempty"`
}

type RetrieveRequest struct {
	ClientIDs []*ClientID `xml:"ClientIDs,omitempty"`

	ObjectType string `xml:"ObjectType,omitempty"`

	Properties []string `xml:"Properties,omitempty"`

	Filter FilterPart `xml:"Filter,omitempty"`
	//SimpleFilter   *SimpleFilterPart `xml:"http://exacttarget.com/wsdl/partnerAPI Filter,omitempty"`
	//ComplexFilter   *ComplexFilterPart `xml:"http://exacttarget.com/wsdl/partnerAPI Filter,omitempty"`

	RespondTo []*AsyncResponse `xml:"RespondTo,omitempty"`

	PartnerProperties []*APIProperty `xml:"PartnerProperties,omitempty"`

	ContinueRequest string `xml:"ContinueRequest,omitempty"`

	QueryAllAccounts bool `xml:"QueryAllAccounts,omitempty"`

	RetrieveAllSinceLastBatch bool `xml:"RetrieveAllSinceLastBatch,omitempty"`

	RepeatLastResult bool `xml:"RepeatLastResult,omitempty"`

	Retrieves struct {
		Request []*Request `xml:"Request,omitempty"`
	} `xml:"Retrieves,omitempty"`

	Options *RetrieveOptions `xml:"Options,omitempty"`
}

type RetrieveSingleRequest struct {
	*Request

	RequestedObject *APIObject `xml:"RequestedObject,omitempty"`

	RetrieveOption *Options `xml:"RetrieveOption,omitempty"`
}

type Parameters struct {
	Parameter []*APIProperty `xml:"Parameter,omitempty"`
}

type RetrieveSingleOptions struct {
	*Options

	Parameters *Parameters `xml:"Parameters,omitempty"`
}

type RetrieveOptions struct {
	*Options

	BatchSize int32 `xml:"BatchSize,omitempty"`

	IncludeObjects bool `xml:"IncludeObjects,omitempty"`

	OnlyIncludeBase bool `xml:"OnlyIncludeBase,omitempty"`
}

type QueryRequest struct {
	ClientIDs []*ClientID `xml:"ClientIDs,omitempty"`

	Query *Query `xml:"Query,omitempty"`

	RespondTo []*AsyncResponse `xml:"RespondTo,omitempty"`

	PartnerProperties []*APIProperty `xml:"PartnerProperties,omitempty"`

	ContinueRequest string `xml:"ContinueRequest,omitempty"`

	QueryAllAccounts bool `xml:"QueryAllAccounts,omitempty"`

	RetrieveAllSinceLastBatch bool `xml:"RetrieveAllSinceLastBatch,omitempty"`
}

type QueryObject struct {
	ObjectType string `xml:"ObjectType,omitempty"`

	Properties []string `xml:"Properties,omitempty"`

	Objects []*QueryObject `xml:"Objects,omitempty"`
}

type Query struct {
	Object *QueryObject `xml:"Object,omitempty"`

	Filter *FilterPart `xml:"Filter,omitempty"`
}

type FilterPart interface {
	ApplyFilter() string
}

type filterPart struct {
}

type SimpleFilterPart struct {
	*filterPart

	//XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI local"`
	XMLName xml.Name `xml:""` // No namespace or local part here; we'll set it dynamically
	XSIType string   `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`

	Property string `xml:"Property,omitempty"`

	SimpleOperator string `xml:"SimpleOperator,omitempty"`

	Value []string `xml:"Value,omitempty"`

	DateValue []string `xml:"DateValue,omitempty"`
}

func (s *SimpleFilterPart) ApplyFilter() string {
	// Logique de filtrage spécifique
	return "SimpleFilterPart"
}

type TagFilterPart struct {
	*filterPart

	Tags struct {
		Tag []string `xml:"Tag,omitempty"`
	} `xml:"Tags,omitempty"`
}

type ComplexFilterPart struct {
	*filterPart

	//XMLName xml.Name `xml:"http://exacttarget.com/wsdl/partnerAPI local"`
	XMLName xml.Name `xml:""` // No namespace or local part here; we'll set it dynamically

	XSIType string `xml:"http://www.w3.org/2001/XMLSchema-instance type,attr"`

	LeftOperand FilterPart `xml:"LeftOperand,omitempty"`

	LogicalOperator LogicalOperator `xml:"LogicalOperator,omitempty"`

	RightOperand FilterPart `xml:"RightOperand,omitempty"`

	AdditionalOperands struct {
		Operand []FilterPart `xml:"Operand,omitempty"`
	} `xml:"AdditionalOperands,omitempty"`
}

func (c *ComplexFilterPart) ApplyFilter() string {
	// Logique de filtrage spécifique
	return "ComplexFilterPart"
}

type ArrayOfObjectDefinitionRequest struct {
	ObjectDefinitionRequest []*ObjectDefinitionRequest `xml:"ObjectDefinitionRequest,omitempty"`
}

type ObjectDefinitionRequest struct {
	Client *ClientID `xml:"Client,omitempty"`

	ObjectType string `xml:"ObjectType,omitempty"`
}

type PropertyDefinition struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	// Deprecated.  Please use ValueType.
	DataType string `xml:"DataType,omitempty"`

	// Reserved for future use.
	ValueType *SoapType `xml:"ValueType,omitempty"`

	// ExactTarget data type of the property
	PropertyType *PropertyType `xml:"PropertyType,omitempty"`

	// Reserved for future use.
	IsCreatable bool `xml:"IsCreatable,omitempty"`

	// Indicates whether the property can be updated.  If true, then this property value can be set in an update call.
	IsUpdatable bool `xml:"IsUpdatable,omitempty"`

	// Indicates whether the object can be retrieved via the retrieve call.
	IsRetrievable bool `xml:"IsRetrievable,omitempty"`

	// Reserved for future use.
	IsQueryable bool `xml:"IsQueryable,omitempty"`

	// Reserved for future use.
	IsFilterable bool `xml:"IsFilterable,omitempty"`

	// Reserved for future use.
	IsPartnerProperty bool `xml:"IsPartnerProperty,omitempty"`

	// Reserved for future use.
	IsAccountProperty bool `xml:"IsAccountProperty,omitempty"`

	// Deprecated.
	PartnerMap string `xml:"PartnerMap,omitempty"`

	AttributeMaps []*AttributeMap `xml:"AttributeMaps,omitempty"`

	// Deprecated.
	Markups []*APIProperty `xml:"Markups,omitempty"`

	// Reserved for future use.
	Precision int32 `xml:"Precision,omitempty"`

	// Reserved for future use.
	Scale int32 `xml:"Scale,omitempty"`

	// Text label that is displayed next to the field in the user interface.
	Label string `xml:"Label,omitempty"`

	Description string `xml:"Description,omitempty"`

	DefaultValue string `xml:"DefaultValue,omitempty"`

	// Minimum length of the data.
	MinLength int32 `xml:"MinLength,omitempty"`

	// Maximum length of the data.
	MaxLength int32 `xml:"MaxLength,omitempty"`

	// Minimum value that this property can be set to.
	MinValue string `xml:"MinValue,omitempty"`

	// Maximum value that this property can be set to.
	MaxValue string `xml:"MaxValue,omitempty"`

	// Indicates whether the property must have a value specified.
	IsRequired bool `xml:"IsRequired,omitempty"`

	// Indicates whether the property is viewable to the end-user in the profile center.
	IsViewable bool `xml:"IsViewable,omitempty"`

	// Indicates whether the property is editable by the end-user in the profile center.
	IsEditable bool `xml:"IsEditable,omitempty"`

	// Reserved for future use.
	IsNillable bool `xml:"IsNillable,omitempty"`

	// Indicates if the property has a restricted list of valid values.
	IsRestrictedPicklist bool `xml:"IsRestrictedPicklist,omitempty"`

	PicklistItems struct {
		PicklistItem []*PicklistItem `xml:"PicklistItem,omitempty"`
	} `xml:"PicklistItems,omitempty"`

	// Indicates whether the property is a send time attribute.
	IsSendTime bool `xml:"IsSendTime,omitempty"`

	// Indicates the placement of this property within the list of properties.
	DisplayOrder int32 `xml:"DisplayOrder,omitempty"`

	References struct {
		Reference []*APIObject `xml:"Reference,omitempty"`
	} `xml:"References,omitempty"`

	// Reserved for future use.
	RelationshipName string `xml:"RelationshipName,omitempty"`

	// Reserved for future use.
	Status string `xml:"Status,omitempty"`

	// Reserved for future use.
	IsContextSpecific bool `xml:"IsContextSpecific,omitempty"`
}

type ObjectDefinition struct {
	ObjectType string `xml:"ObjectType,omitempty"`

	Name string `xml:"Name,omitempty"`

	IsCreatable bool `xml:"IsCreatable,omitempty"`

	IsUpdatable bool `xml:"IsUpdatable,omitempty"`

	IsRetrievable bool `xml:"IsRetrievable,omitempty"`

	IsQueryable bool `xml:"IsQueryable,omitempty"`

	IsReference bool `xml:"IsReference,omitempty"`

	ReferencedType string `xml:"ReferencedType,omitempty"`

	IsPropertyCollection string `xml:"IsPropertyCollection,omitempty"`

	IsObjectCollection bool `xml:"IsObjectCollection,omitempty"`

	Properties []*PropertyDefinition `xml:"Properties,omitempty"`

	ExtendedProperties struct {
		ExtendedProperty []*PropertyDefinition `xml:"ExtendedProperty,omitempty"`
	} `xml:"ExtendedProperties,omitempty"`

	ChildObjects []*ObjectDefinition `xml:"ChildObjects,omitempty"`
}

type AttributeMap struct {
	*APIProperty

	EntityName string `xml:"EntityName,omitempty"`

	ColumnName string `xml:"ColumnName,omitempty"`

	ColumnNameMappedTo string `xml:"ColumnNameMappedTo,omitempty"`

	EntityNameMappedTo string `xml:"EntityNameMappedTo,omitempty"`

	AdditionalData []*APIProperty `xml:"AdditionalData,omitempty"`
}

type PicklistItem struct {
	IsDefaultValue bool `xml:"IsDefaultValue,omitempty"`

	Label string `xml:"Label,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type ExecuteRequest struct {
	Client *ClientID `xml:"Client,omitempty"`

	Name string `xml:"Name,omitempty"`

	Parameters []*APIProperty `xml:"Parameters,omitempty"`
}

type ExecuteResponse struct {
	StatusCode string `xml:"StatusCode,omitempty"`

	StatusMessage string `xml:"StatusMessage,omitempty"`

	OrdinalID int32 `xml:"OrdinalID,omitempty"`

	Results []*APIProperty `xml:"Results,omitempty"`

	ErrorCode int32 `xml:"ErrorCode,omitempty"`
}

type InteractionDefinition struct {
	*InteractionBaseObject

	InteractionObjectID string `xml:"InteractionObjectID,omitempty"`
}

type InteractionBaseObject struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	Keyword string `xml:"Keyword,omitempty"`
}

type PerformOptions struct {
	*Options

	Explanation string `xml:"Explanation,omitempty"`

	ProgramActivityInstanceID *Instanceid `xml:"ProgramActivityInstanceID,omitempty"`
}

type CampaignPerformOptions struct {
	*PerformOptions

	OccurrenceIDs []string `xml:"OccurrenceIDs,omitempty"`

	OccurrenceIDsIndex int32 `xml:"OccurrenceIDsIndex,omitempty"`
}

type PerformRequest struct {
	Client *ClientID `xml:"Client,omitempty"`

	Action string `xml:"Action,omitempty"`

	Definitions struct {
		Definition []*InteractionBaseObject `xml:"Definition,omitempty"`
	} `xml:"Definitions,omitempty"`
}

type PerformResponse struct {
	StatusCode string `xml:"StatusCode,omitempty"`

	StatusMessage string `xml:"StatusMessage,omitempty"`

	OrdinalID int32 `xml:"OrdinalID,omitempty"`

	Results struct {
		Result []*APIProperty `xml:"Result,omitempty"`
	} `xml:"Results,omitempty"`

	ErrorCode int32 `xml:"ErrorCode,omitempty"`
}

type PerformResult struct {
	*Result

	Object *APIObject `xml:"Object,omitempty"`

	Task *TaskResult `xml:"Task,omitempty"`

	ProgramActivityInstanceId *Instanceid `xml:"ProgramActivityInstanceId,omitempty"`
}

type ValidationAction struct {
	ValidationType string `xml:"ValidationType,omitempty"`

	ValidationOptions struct {
		ValidationOption []*APIProperty `xml:"ValidationOption,omitempty"`
	} `xml:"ValidationOptions,omitempty"`
}

type SpamAssassinValidation struct {
	*ValidationAction
}

type ContentValidation struct {
	*APIObject

	ValidationAction *ValidationAction `xml:"ValidationAction,omitempty"`

	Email *Email `xml:"Email,omitempty"`

	Subscribers struct {
		Subscriber []*Subscriber `xml:"Subscriber,omitempty"`
	} `xml:"Subscribers,omitempty"`
}

type ContentValidationResult struct {
	*PerformResult
}

type ValidationResult struct {
	Subscriber *Subscriber `xml:"Subscriber,omitempty"`

	CheckTime time.Time `xml:"CheckTime,omitempty"`

	CheckTimeUTC time.Time `xml:"CheckTimeUTC,omitempty"`

	IsResultValid bool `xml:"IsResultValid,omitempty"`

	IsSpam bool `xml:"IsSpam,omitempty"`

	Score float64 `xml:"Score,omitempty"`

	Threshold float64 `xml:"Threshold,omitempty"`

	Message string `xml:"Message,omitempty"`
}

type ContentValidationTaskResult struct {
	*TaskResult

	ValidationResults struct {
		ValidationResult []*ValidationResult `xml:"ValidationResult,omitempty"`
	} `xml:"ValidationResults,omitempty"`
}

type ConfigureOptions struct {
	*Options
}

type ConfigureResult struct {
	*Result

	Object *APIObject `xml:"Object,omitempty"`
}

type ScheduleDefinition struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	Recurrence *Recurrence `xml:"Recurrence,omitempty"`

	RecurrenceType *RecurrenceTypeEnum `xml:"RecurrenceType,omitempty"`

	RecurrenceRangeType *RecurrenceRangeTypeEnum `xml:"RecurrenceRangeType,omitempty"`

	StartDateTime time.Time `xml:"StartDateTime,omitempty"`

	EndDateTime time.Time `xml:"EndDateTime,omitempty"`

	Occurrences int32 `xml:"Occurrences,omitempty"`

	Keyword string `xml:"Keyword,omitempty"`

	TimeZone *TimeZone `xml:"TimeZone,omitempty"`
}

type ScheduleOptions struct {
	*Options
}

type ScheduleResponse struct {
	StatusCode string `xml:"StatusCode,omitempty"`

	StatusMessage string `xml:"StatusMessage,omitempty"`

	OrdinalID int32 `xml:"OrdinalID,omitempty"`

	Results struct {
		Result []*APIProperty `xml:"Result,omitempty"`
	} `xml:"Results,omitempty"`

	ErrorCode int32 `xml:"ErrorCode,omitempty"`
}

type ScheduleResult struct {
	*Result

	Object *ScheduleDefinition `xml:"Object,omitempty"`

	Task *TaskResult `xml:"Task,omitempty"`
}

type Recurrence struct {
}

type MinutelyRecurrence struct {
	*Recurrence

	MinutelyRecurrencePatternType *MinutelyRecurrencePatternTypeEnum `xml:"MinutelyRecurrencePatternType,omitempty"`

	MinuteInterval int32 `xml:"MinuteInterval,omitempty"`
}

type HourlyRecurrence struct {
	*Recurrence

	HourlyRecurrencePatternType *HourlyRecurrencePatternTypeEnum `xml:"HourlyRecurrencePatternType,omitempty"`

	HourInterval int32 `xml:"HourInterval,omitempty"`
}

type DailyRecurrence struct {
	*Recurrence

	DailyRecurrencePatternType *DailyRecurrencePatternTypeEnum `xml:"DailyRecurrencePatternType,omitempty"`

	DayInterval int32 `xml:"DayInterval,omitempty"`
}

type WeeklyRecurrence struct {
	*Recurrence

	WeeklyRecurrencePatternType *WeeklyRecurrencePatternTypeEnum `xml:"WeeklyRecurrencePatternType,omitempty"`

	WeekInterval int32 `xml:"WeekInterval,omitempty"`

	Sunday bool `xml:"Sunday,omitempty"`

	Monday bool `xml:"Monday,omitempty"`

	Tuesday bool `xml:"Tuesday,omitempty"`

	Wednesday bool `xml:"Wednesday,omitempty"`

	Thursday bool `xml:"Thursday,omitempty"`

	Friday bool `xml:"Friday,omitempty"`

	Saturday bool `xml:"Saturday,omitempty"`
}

type MonthlyRecurrence struct {
	*Recurrence

	MonthlyRecurrencePatternType *MonthlyRecurrencePatternTypeEnum `xml:"MonthlyRecurrencePatternType,omitempty"`

	MonthlyInterval int32 `xml:"MonthlyInterval,omitempty"`

	ScheduledDay int32 `xml:"ScheduledDay,omitempty"`

	ScheduledWeek *WeekOfMonthEnum `xml:"ScheduledWeek,omitempty"`

	ScheduledDayOfWeek *DayOfWeekEnum `xml:"ScheduledDayOfWeek,omitempty"`
}

type YearlyRecurrence struct {
	*Recurrence

	YearlyRecurrencePatternType *YearlyRecurrencePatternTypeEnum `xml:"YearlyRecurrencePatternType,omitempty"`

	ScheduledDay int32 `xml:"ScheduledDay,omitempty"`

	ScheduledWeek *WeekOfMonthEnum `xml:"ScheduledWeek,omitempty"`

	ScheduledMonth *MonthOfYearEnum `xml:"ScheduledMonth,omitempty"`

	ScheduledDayOfWeek *DayOfWeekEnum `xml:"ScheduledDayOfWeek,omitempty"`
}

type ExtractRequest struct {
	*Request

	Client *ClientID `xml:"Client,omitempty"`

	ID string `xml:"ID,omitempty"`

	Options *ExtractOptions `xml:"Options,omitempty"`

	Parameters struct {
		Parameter []*ExtractParameter `xml:"Parameter,omitempty"`
	} `xml:"Parameters,omitempty"`
}

type ExtractResult struct {
	*Result

	Request *ExtractRequest `xml:"Request,omitempty"`
}

type ExtractOptions struct {
	*Options
}

type ExtractParameter struct {
	*APIProperty
}

type ExtractTemplate struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	ConfigurationPage string `xml:"ConfigurationPage,omitempty"`

	PackageKey string `xml:"PackageKey,omitempty"`
}

type ExtractDescription struct {
	*ExtractTemplate

	Parameters struct {
		Parameter []*ExtractParameterDescription `xml:"Parameter,omitempty"`
	} `xml:"Parameters,omitempty"`
}

type ExtractDefinition struct {
	*ExtractTemplate

	Parameters struct {
		Parameter []*ExtractParameterDescription `xml:"Parameter,omitempty"`
	} `xml:"Parameters,omitempty"`

	Values struct {
		Value []*APIProperty `xml:"Value,omitempty"`
	} `xml:"Values,omitempty"`
}

type ParameterDescription struct {
}

type ExtractParameterDescription struct {
	*ParameterDescription

	Name string `xml:"Name,omitempty"`

	DataType *ExtractParameterDataType `xml:"DataType,omitempty"`

	DefaultValue string `xml:"DefaultValue,omitempty"`

	IsOptional bool `xml:"IsOptional,omitempty"`

	DropDownList string `xml:"DropDownList,omitempty"`
}

type VersionInfoResponse struct {
	Version string `xml:"Version,omitempty"`

	VersionDate time.Time `xml:"VersionDate,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	VersionHistory []*VersionInfoResponse `xml:"VersionHistory,omitempty"`
}

type Locale struct {
	*APIObject

	LocaleCode string `xml:"LocaleCode,omitempty"`
}

type TimeZone struct {
	*APIObject

	Name string `xml:"Name,omitempty"`
}

type Account struct {
	*APIObject

	AccountType *AccountTypeEnum `xml:"AccountType,omitempty"`

	ParentID int32 `xml:"ParentID,omitempty"`

	BrandID int32 `xml:"BrandID,omitempty"`

	PrivateLabelID int32 `xml:"PrivateLabelID,omitempty"`

	ReportingParentID int32 `xml:"ReportingParentID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Email string `xml:"Email,omitempty"`

	FromName string `xml:"FromName,omitempty"`

	BusinessName string `xml:"BusinessName,omitempty"`

	Phone string `xml:"Phone,omitempty"`

	Address string `xml:"Address,omitempty"`

	Fax string `xml:"Fax,omitempty"`

	City string `xml:"City,omitempty"`

	State string `xml:"State,omitempty"`

	Zip string `xml:"Zip,omitempty"`

	Country string `xml:"Country,omitempty"`

	IsActive int32 `xml:"IsActive,omitempty"`

	IsTestAccount bool `xml:"IsTestAccount,omitempty"`

	OrgID int32 `xml:"OrgID,omitempty"`

	DBID int32 `xml:"DBID,omitempty"`

	ParentName string `xml:"ParentName,omitempty"`

	CustomerID int64 `xml:"CustomerID,omitempty"`

	DeletedDate time.Time `xml:"DeletedDate,omitempty"`

	EditionID int32 `xml:"EditionID,omitempty"`

	Children []*AccountDataItem `xml:"Children,omitempty"`

	Subscription *Subscription `xml:"Subscription,omitempty"`

	PrivateLabels []*PrivateLabel `xml:"PrivateLabels,omitempty"`

	BusinessRules []*BusinessRule `xml:"BusinessRules,omitempty"`

	AccountUsers []*AccountUser `xml:"AccountUsers,omitempty"`

	InheritAddress bool `xml:"InheritAddress,omitempty"`

	IsTrialAccount bool `xml:"IsTrialAccount,omitempty"`

	Locale *Locale `xml:"Locale,omitempty"`

	ParentAccount *Account `xml:"ParentAccount,omitempty"`

	TimeZone *TimeZone `xml:"TimeZone,omitempty"`

	Roles struct {
		Role []*Role `xml:"Role,omitempty"`
	} `xml:"Roles,omitempty"`

	LanguageLocale *Locale `xml:"LanguageLocale,omitempty"`
}

type BusinessUnit struct {
	*Account

	Description string `xml:"Description,omitempty"`

	DefaultSendClassification *SendClassification `xml:"DefaultSendClassification,omitempty"`

	DefaultHomePage *LandingPage `xml:"DefaultHomePage,omitempty"`

	SubscriberFilter *FilterPart `xml:"SubscriberFilter,omitempty"`

	MasterUnsubscribeBehavior *UnsubscribeBehaviorEnum `xml:"MasterUnsubscribeBehavior,omitempty"`
}

type LandingPage struct {
	*APIObject
}

type AccountDataItem struct {
	ChildAccountID int32 `xml:"ChildAccountID,omitempty"`

	BrandID int32 `xml:"BrandID,omitempty"`

	PrivateLabelID int32 `xml:"PrivateLabelID,omitempty"`

	AccountType int32 `xml:"AccountType,omitempty"`
}

type Subscription struct {
	SubscriptionID int32 `xml:"SubscriptionID,omitempty"`

	EmailsPurchased int32 `xml:"EmailsPurchased,omitempty"`

	AccountsPurchased int32 `xml:"AccountsPurchased,omitempty"`

	AdvAccountsPurchased int32 `xml:"AdvAccountsPurchased,omitempty"`

	LPAccountsPurchased int32 `xml:"LPAccountsPurchased,omitempty"`

	DOTOAccountsPurchased int32 `xml:"DOTOAccountsPurchased,omitempty"`

	BUAccountsPurchased int32 `xml:"BUAccountsPurchased,omitempty"`

	BeginDate time.Time `xml:"BeginDate,omitempty"`

	EndDate time.Time `xml:"EndDate,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	Period string `xml:"Period,omitempty"`

	NotificationTitle string `xml:"NotificationTitle,omitempty"`

	NotificationMessage string `xml:"NotificationMessage,omitempty"`

	NotificationFlag string `xml:"NotificationFlag,omitempty"`

	NotificationExpDate time.Time `xml:"NotificationExpDate,omitempty"`

	ForAccounting string `xml:"ForAccounting,omitempty"`

	HasPurchasedEmails bool `xml:"HasPurchasedEmails,omitempty"`

	ContractNumber string `xml:"ContractNumber,omitempty"`

	ContractModifier string `xml:"ContractModifier,omitempty"`

	IsRenewal bool `xml:"IsRenewal,omitempty"`

	NumberofEmails int64 `xml:"NumberofEmails,omitempty"`
}

type PrivateLabel struct {
	ID int32 `xml:"ID,omitempty"`

	Name string `xml:"Name,omitempty"`

	ColorPaletteXML string `xml:"ColorPaletteXML,omitempty"`

	LogoFile string `xml:"LogoFile,omitempty"`

	Delete int32 `xml:"Delete,omitempty"`

	SetActive bool `xml:"SetActive,omitempty"`
}

type AccountPrivateLabel struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	OwnerMemberID int32 `xml:"OwnerMemberID,omitempty"`

	ColorPaletteXML string `xml:"ColorPaletteXML,omitempty"`
}

type BusinessRule struct {
	*APIObject

	MemberBusinessRuleID int32 `xml:"MemberBusinessRuleID,omitempty"`

	BusinessRuleID int32 `xml:"BusinessRuleID,omitempty"`

	Data int32 `xml:"Data,omitempty"`

	Quality string `xml:"Quality,omitempty"`

	Name string `xml:"Name,omitempty"`

	Type string `xml:"Type,omitempty"`

	Description string `xml:"Description,omitempty"`

	IsViewable bool `xml:"IsViewable,omitempty"`

	IsInheritedFromParent bool `xml:"IsInheritedFromParent,omitempty"`

	DisplayName string `xml:"DisplayName,omitempty"`

	ProductCode string `xml:"ProductCode,omitempty"`
}

type AccountUser struct {
	*APIObject

	AccountUserID int32 `xml:"AccountUserID,omitempty"`

	UserID string `xml:"UserID,omitempty"`

	Password string `xml:"Password,omitempty"`

	Name string `xml:"Name,omitempty"`

	Email string `xml:"Email,omitempty"`

	MustChangePassword bool `xml:"MustChangePassword,omitempty"`

	ActiveFlag bool `xml:"ActiveFlag,omitempty"`

	ChallengePhrase string `xml:"ChallengePhrase,omitempty"`

	ChallengeAnswer string `xml:"ChallengeAnswer,omitempty"`

	UserPermissions []*UserAccess `xml:"UserPermissions,omitempty"`

	Delete int32 `xml:"Delete,omitempty"`

	LastSuccessfulLogin time.Time `xml:"LastSuccessfulLogin,omitempty"`

	IsAPIUser bool `xml:"IsAPIUser,omitempty"`

	NotificationEmailAddress string `xml:"NotificationEmailAddress,omitempty"`

	IsLocked bool `xml:"IsLocked,omitempty"`

	Unlock bool `xml:"Unlock,omitempty"`

	BusinessUnit int32 `xml:"BusinessUnit,omitempty"`

	DefaultBusinessUnit int32 `xml:"DefaultBusinessUnit,omitempty"`

	DefaultApplication string `xml:"DefaultApplication,omitempty"`

	Locale *Locale `xml:"Locale,omitempty"`

	TimeZone *TimeZone `xml:"TimeZone,omitempty"`

	DefaultBusinessUnitObject *BusinessUnit `xml:"DefaultBusinessUnitObject,omitempty"`

	AssociatedBusinessUnits struct {
		BusinessUnit []*BusinessUnit `xml:"BusinessUnit,omitempty"`
	} `xml:"AssociatedBusinessUnits,omitempty"`

	Roles struct {
		Role []*Role `xml:"Role,omitempty"`
	} `xml:"Roles,omitempty"`

	LanguageLocale *Locale `xml:"LanguageLocale,omitempty"`

	SsoIdentities struct {
		SsoIdentity []*SsoIdentity `xml:"SsoIdentity,omitempty"`
	} `xml:"SsoIdentities,omitempty"`

	IsSendable bool `xml:"IsSendable,omitempty"`
}

type SsoIdentity struct {
	*APIObject

	FederatedID string `xml:"FederatedID,omitempty"`

	IsActive bool `xml:"IsActive,omitempty"`
}

type UserAccess struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Value string `xml:"Value,omitempty"`

	Description string `xml:"Description,omitempty"`

	Delete int32 `xml:"Delete,omitempty"`
}

type Brand struct {
	*APIObject

	BrandID int32 `xml:"BrandID,omitempty"`

	Label string `xml:"Label,omitempty"`

	Comment string `xml:"Comment,omitempty"`

	BrandTags []*BrandTag `xml:"BrandTags,omitempty"`
}

type BrandTag struct {
	*APIObject

	BrandID int32 `xml:"BrandID,omitempty"`

	Label string `xml:"Label,omitempty"`

	Data string `xml:"Data,omitempty"`
}

type Role struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	IsPrivate bool `xml:"IsPrivate,omitempty"`

	IsSystemDefined bool `xml:"IsSystemDefined,omitempty"`

	ForceInheritance bool `xml:"ForceInheritance,omitempty"`

	PermissionSets struct {
		PermissionSet []*PermissionSet `xml:"PermissionSet,omitempty"`
	} `xml:"PermissionSets,omitempty"`

	Permissions struct {
		Permission []*Permission `xml:"Permission,omitempty"`
	} `xml:"Permissions,omitempty"`
}

type PermissionSet struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	IsAllowed bool `xml:"IsAllowed,omitempty"`

	IsDenied bool `xml:"IsDenied,omitempty"`

	PermissionSets struct {
		PermissionSet []*PermissionSet `xml:"PermissionSet,omitempty"`
	} `xml:"PermissionSets,omitempty"`

	Permissions struct {
		Permission []*Permission `xml:"Permission,omitempty"`
	} `xml:"Permissions,omitempty"`
}

type Permission struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	ObjectType string `xml:"ObjectType,omitempty"`

	Operation string `xml:"Operation,omitempty"`

	IsShareable bool `xml:"IsShareable,omitempty"`

	IsAllowed bool `xml:"IsAllowed,omitempty"`

	IsDenied bool `xml:"IsDenied,omitempty"`
}

type Email struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Folder string `xml:"Folder,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	HTMLBody string `xml:"HTMLBody,omitempty"`

	TextBody string `xml:"TextBody,omitempty"`

	ContentAreas []*ContentArea `xml:"ContentAreas,omitempty"`

	Subject string `xml:"Subject,omitempty"`

	IsActive bool `xml:"IsActive,omitempty"`

	IsHTMLPaste bool `xml:"IsHTMLPaste,omitempty"`

	ClonedFromID int32 `xml:"ClonedFromID,omitempty"`

	Status string `xml:"Status,omitempty"`

	EmailType string `xml:"EmailType,omitempty"`

	CharacterSet string `xml:"CharacterSet,omitempty"`

	HasDynamicSubjectLine bool `xml:"HasDynamicSubjectLine,omitempty"`

	ContentCheckStatus string `xml:"ContentCheckStatus,omitempty"`

	SyncTextWithHTML bool `xml:"SyncTextWithHTML,omitempty"`

	PreHeader string `xml:"PreHeader,omitempty"`

	IsApproved bool `xml:"IsApproved,omitempty"`
}

type ContentArea struct {
	*APIObject

	Key string `xml:"Key,omitempty"`

	Content string `xml:"Content,omitempty"`

	IsBlank bool `xml:"IsBlank,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Layout *LayoutType `xml:"Layout,omitempty"`

	IsDynamicContent bool `xml:"IsDynamicContent,omitempty"`

	IsSurvey bool `xml:"IsSurvey,omitempty"`

	BackgroundColor string `xml:"BackgroundColor,omitempty"`

	BorderColor string `xml:"BorderColor,omitempty"`

	BorderWidth int32 `xml:"BorderWidth,omitempty"`

	Cellpadding int32 `xml:"Cellpadding,omitempty"`

	Cellspacing int32 `xml:"Cellspacing,omitempty"`

	Width string `xml:"Width,omitempty"`

	FontFamily string `xml:"FontFamily,omitempty"`

	HasFontSize bool `xml:"HasFontSize,omitempty"`

	IsLocked bool `xml:"IsLocked,omitempty"`
}

type Message struct {
	*APIObject

	TextBody string `xml:"TextBody,omitempty"`
}

type TrackingEvent struct {
	*APIObject

	SendID int32 `xml:"SendID,omitempty"`

	SubscriberKey string `xml:"SubscriberKey,omitempty"`

	EventDate time.Time `xml:"EventDate,omitempty"`

	EventType *EventType `xml:"EventType,omitempty"`

	TriggeredSendDefinitionObjectID string `xml:"TriggeredSendDefinitionObjectID,omitempty"`

	BatchID int32 `xml:"BatchID,omitempty"`
}

type OpenEvent struct {
	*TrackingEvent
}

type BounceEvent struct {
	*TrackingEvent

	SMTPCode string `xml:"SMTPCode,omitempty"`

	BounceCategory string `xml:"BounceCategory,omitempty"`

	SMTPReason string `xml:"SMTPReason,omitempty"`

	BounceType string `xml:"BounceType,omitempty"`
}

type UnsubEvent struct {
	*TrackingEvent

	List *List `xml:"List,omitempty"`

	IsMasterUnsubscribed bool `xml:"IsMasterUnsubscribed,omitempty"`
}

type ClickEvent struct {
	*TrackingEvent

	URLID int32 `xml:"URLID,omitempty"`

	URL string `xml:"URL,omitempty"`

	URLIDLong int64 `xml:"URLIDLong,omitempty"`
}

type SentEvent struct {
	*TrackingEvent
}

type NotSentEvent struct {
	*TrackingEvent
}

type SurveyEvent struct {
	*TrackingEvent

	Question string `xml:"Question,omitempty"`

	Answer string `xml:"Answer,omitempty"`
}

type ForwardedEmailEvent struct {
	*TrackingEvent
}

type ForwardedEmailOptInEvent struct {
	*TrackingEvent

	OptInSubscriberKey string `xml:"OptInSubscriberKey,omitempty"`
}

type DeliveredEvent struct {
	*TrackingEvent
}

type Subscriber struct {
	*APIObject

	EmailAddress string `xml:"EmailAddress,omitempty" json:"emailAddress,omitempty"`

	Attributes []*Attribute `xml:"Attributes,omitempty" json:"attributes,omitempty"`

	SubscriberKey string `xml:"SubscriberKey,omitempty" json:"subscriberKey,omitempty"`

	UnsubscribedDate time.Time `xml:"UnsubscribedDate,omitempty" json:"unsubscribedDate,omitempty"`

	Status *SubscriberStatus `xml:"Status,omitempty" json:"status,omitempty"`

	PartnerType string `xml:"PartnerType,omitempty" json:"partnerType,omitempty"`

	EmailTypePreference *EmailType `xml:"EmailTypePreference,omitempty" json:"emailTypePreference,omitempty"`

	Lists []*SubscriberList `xml:"Lists,omitempty" json:"lists,omitempty"`

	GlobalUnsubscribeCategory *GlobalUnsubscribeCategory `xml:"GlobalUnsubscribeCategory,omitempty" json:"globalUnsubscribeCategory,omitempty"`

	SubscriberTypeDefinition *SubscriberTypeDefinition `xml:"SubscriberTypeDefinition,omitempty" json:"subscriberTypeDefinition,omitempty"`

	Addresses struct {
		Address []*SubscriberAddress `xml:"Address,omitempty" json:"address,omitempty"`
	} `xml:"Addresses,omitempty" json:"addresses,omitempty"`

	PrimarySMSAddress *SMSAddress `xml:"PrimarySMSAddress,omitempty" json:"primarySMSAddress,omitempty"`

	PrimarySMSPublicationStatus *SubscriberAddressStatus `xml:"PrimarySMSPublicationStatus,omitempty" json:"primarySMSPublicationStatus,omitempty"`

	PrimaryEmailAddress *EmailAddress `xml:"PrimaryEmailAddress,omitempty" json:"primaryEmailAddress,omitempty"`

	Locale *Locale `xml:"Locale,omitempty" json:"locale,omitempty"`
}

type Attribute struct {
	Name string `xml:"Name,omitempty"`

	Value string `xml:"Value,omitempty"`

	Compression *CompressionConfiguration `xml:"Compression,omitempty"`
}

type CompressionConfiguration struct {
	Type *CompressionType `xml:"Type,omitempty"`

	Encoding *CompressionEncoding `xml:"Encoding,omitempty"`
}

type SubscriberTypeDefinition struct {
	SubscriberType string `xml:"SubscriberType,omitempty"`
}

type ListSubscriber struct {
	*APIObject

	Status *SubscriberStatus `xml:"Status,omitempty"`

	ListID int32 `xml:"ListID,omitempty"`

	SubscriberKey string `xml:"SubscriberKey,omitempty"`
}

type SubscriberList struct {
	*APIObject

	Status *SubscriberStatus `xml:"Status,omitempty"`

	List *List `xml:"List,omitempty"`

	Action string `xml:"Action,omitempty"`

	Subscriber *Subscriber `xml:"Subscriber,omitempty"`
}

type List struct {
	*APIObject

	ListName string `xml:"ListName,omitempty"`

	Category int32 `xml:"Category,omitempty"`

	Type *ListTypeEnum `xml:"Type,omitempty"`

	Description string `xml:"Description,omitempty"`

	Subscribers []*Subscriber `xml:"Subscribers,omitempty"`

	ListClassification *ListClassificationEnum `xml:"ListClassification,omitempty"`

	AutomatedEmail *Email `xml:"AutomatedEmail,omitempty"`

	SendClassification *SendClassification `xml:"SendClassification,omitempty"`
}

type Group struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Category int32 `xml:"Category,omitempty"`

	Description string `xml:"Description,omitempty"`

	Subscribers []*Subscriber `xml:"Subscribers,omitempty"`
}

type ListAttribute struct {
	*APIObject

	List *List `xml:"List,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	FieldType *ListAttributeFieldType `xml:"FieldType,omitempty"`

	FieldLength int32 `xml:"FieldLength,omitempty"`

	Scale int32 `xml:"Scale,omitempty"`

	MinValue string `xml:"MinValue,omitempty"`

	MaxValue string `xml:"MaxValue,omitempty"`

	DefaultValue string `xml:"DefaultValue,omitempty"`

	IsNullable bool `xml:"IsNullable,omitempty"`

	IsHidden bool `xml:"IsHidden,omitempty"`

	IsReadOnly bool `xml:"IsReadOnly,omitempty"`

	Inheritable bool `xml:"Inheritable,omitempty"`

	Overridable bool `xml:"Overridable,omitempty"`

	MustOverride bool `xml:"MustOverride,omitempty"`

	OverrideType *OverrideType `xml:"OverrideType,omitempty"`

	Ordinal int32 `xml:"Ordinal,omitempty"`

	RestrictedValues []*ListAttributeRestrictedValue `xml:"RestrictedValues,omitempty"`

	BaseAttribute *ListAttribute `xml:"BaseAttribute,omitempty"`
}

type ListAttributeRestrictedValue struct {
	*APIObject

	ValueName string `xml:"ValueName,omitempty"`

	IsDefault bool `xml:"IsDefault,omitempty"`

	DisplayOrder int32 `xml:"DisplayOrder,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type GlobalUnsubscribeCategory struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	IgnorableByPartners bool `xml:"IgnorableByPartners,omitempty"`

	Ignore bool `xml:"Ignore,omitempty"`
}

type Campaign struct {
	*InteractionDefinition
}

type Send struct {
	*APIObject

	Email *Email `xml:"Email,omitempty"`

	List []*List `xml:"List,omitempty"`

	SendDate time.Time `xml:"SendDate,omitempty"`

	FromAddress string `xml:"FromAddress,omitempty"`

	FromName string `xml:"FromName,omitempty"`

	Duplicates int32 `xml:"Duplicates,omitempty"`

	InvalidAddresses int32 `xml:"InvalidAddresses,omitempty"`

	ExistingUndeliverables int32 `xml:"ExistingUndeliverables,omitempty"`

	ExistingUnsubscribes int32 `xml:"ExistingUnsubscribes,omitempty"`

	HardBounces int32 `xml:"HardBounces,omitempty"`

	SoftBounces int32 `xml:"SoftBounces,omitempty"`

	OtherBounces int32 `xml:"OtherBounces,omitempty"`

	ForwardedEmails int32 `xml:"ForwardedEmails,omitempty"`

	UniqueClicks int32 `xml:"UniqueClicks,omitempty"`

	UniqueOpens int32 `xml:"UniqueOpens,omitempty"`

	NumberSent int32 `xml:"NumberSent,omitempty"`

	NumberDelivered int32 `xml:"NumberDelivered,omitempty"`

	Unsubscribes int32 `xml:"Unsubscribes,omitempty"`

	MissingAddresses int32 `xml:"MissingAddresses,omitempty"`

	Subject string `xml:"Subject,omitempty"`

	PreviewURL string `xml:"PreviewURL,omitempty"`

	Links []*Link `xml:"Links,omitempty"`

	Events []*TrackingEvent `xml:"Events,omitempty"`

	SentDate time.Time `xml:"SentDate,omitempty"`

	EmailName string `xml:"EmailName,omitempty"`

	Status string `xml:"Status,omitempty"`

	IsMultipart bool `xml:"IsMultipart,omitempty"`

	SendLimit int32 `xml:"SendLimit,omitempty"`

	SendWindowOpen time.Time `xml:"SendWindowOpen,omitempty"`

	SendWindowClose time.Time `xml:"SendWindowClose,omitempty"`

	IsAlwaysOn bool `xml:"IsAlwaysOn,omitempty"`

	Sources struct {
		Source []*APIObject `xml:"Source,omitempty"`
	} `xml:"Sources,omitempty"`

	NumberTargeted int32 `xml:"NumberTargeted,omitempty"`

	NumberErrored int32 `xml:"NumberErrored,omitempty"`

	NumberExcluded int32 `xml:"NumberExcluded,omitempty"`

	Additional string `xml:"Additional,omitempty"`

	BccEmail string `xml:"BccEmail,omitempty"`

	EmailSendDefinition *EmailSendDefinition `xml:"EmailSendDefinition,omitempty"`

	SuppressionLists struct {
		SuppressionList []*AudienceItem `xml:"SuppressionList,omitempty"`
	} `xml:"SuppressionLists,omitempty"`
}

type Link struct {
	*APIObject

	LastClicked time.Time `xml:"LastClicked,omitempty"`

	Alias string `xml:"Alias,omitempty"`

	TotalClicks int32 `xml:"TotalClicks,omitempty"`

	UniqueClicks int32 `xml:"UniqueClicks,omitempty"`

	URL string `xml:"URL,omitempty"`

	Subscribers []*TrackingEvent `xml:"Subscribers,omitempty"`

	IDLong int64 `xml:"IDLong,omitempty"`
}

type SendSummary struct {
	*APIObject

	AccountID int32 `xml:"AccountID,omitempty"`

	AccountName string `xml:"AccountName,omitempty"`

	AccountEmail string `xml:"AccountEmail,omitempty"`

	IsTestAccount bool `xml:"IsTestAccount,omitempty"`

	SendID int32 `xml:"SendID,omitempty"`

	DeliveredTime string `xml:"DeliveredTime,omitempty"`

	TotalSent int32 `xml:"TotalSent,omitempty"`

	Transactional int32 `xml:"Transactional,omitempty"`

	NonTransactional int32 `xml:"NonTransactional,omitempty"`
}

type TriggeredSendDefinition struct {
	*SendDefinition

	// Always will be set to Continuous. For additional fee, TriggeredSendDefinition.Priority can be used to adjust priority.
	TriggeredSendType *TriggeredSendTypeEnum `xml:"TriggeredSendType,omitempty"`

	TriggeredSendStatus *TriggeredSendStatusEnum `xml:"TriggeredSendStatus,omitempty"`

	Email *Email `xml:"Email,omitempty"`

	List *List `xml:"List,omitempty"`

	AutoAddSubscribers bool `xml:"AutoAddSubscribers,omitempty"`

	AutoUpdateSubscribers bool `xml:"AutoUpdateSubscribers,omitempty"`

	// Always will be set to 1.  For additional fee, TriggeredSendDefinition.Priority can be used to adjust priority.
	BatchInterval int32 `xml:"BatchInterval,omitempty"`

	BccEmail string `xml:"BccEmail,omitempty"`

	EmailSubject string `xml:"EmailSubject,omitempty"`

	DynamicEmailSubject string `xml:"DynamicEmailSubject,omitempty"`

	IsMultipart bool `xml:"IsMultipart,omitempty"`

	IsWrapped bool `xml:"IsWrapped,omitempty"`

	AllowedSlots int16 `xml:"AllowedSlots,omitempty"`

	NewSlotTrigger int32 `xml:"NewSlotTrigger,omitempty"`

	SendLimit int32 `xml:"SendLimit,omitempty"`

	SendWindowOpen time.Time `xml:"SendWindowOpen,omitempty"`

	SendWindowClose time.Time `xml:"SendWindowClose,omitempty"`

	SendWindowDelete bool `xml:"SendWindowDelete,omitempty"`

	RefreshContent bool `xml:"RefreshContent,omitempty"`

	ExclusionFilter string `xml:"ExclusionFilter,omitempty"`

	Priority string `xml:"Priority,omitempty"`

	// Deprecated.  Use SendSourceDataExtension instead.
	SendSourceCustomerKey string `xml:"SendSourceCustomerKey,omitempty"`

	ExclusionListCollection []*TriggeredSendExclusionList `xml:"ExclusionListCollection,omitempty"`

	CCEmail string `xml:"CCEmail,omitempty"`

	SendSourceDataExtension *DataExtension `xml:"SendSourceDataExtension,omitempty"`

	IsAlwaysOn bool `xml:"IsAlwaysOn,omitempty"`

	DisableOnEmailBuildError bool `xml:"DisableOnEmailBuildError,omitempty"`

	PreHeader string `xml:"PreHeader,omitempty"`

	ReplyToAddress string `xml:"ReplyToAddress,omitempty"`

	ReplyToDisplayName string `xml:"ReplyToDisplayName,omitempty"`

	TriggeredSendClass *TriggeredSendClassEnum `xml:"TriggeredSendClass,omitempty"`

	TriggeredSendSubClass *TriggeredSendSubClassEnum `xml:"TriggeredSendSubClass,omitempty"`
}

type TriggeredSendExclusionList struct {
	*AudienceItem
}

type TriggeredSend struct {
	*APIObject

	TriggeredSendDefinition *TriggeredSendDefinition `xml:"TriggeredSendDefinition,omitempty"`

	Subscribers []*Subscriber `xml:"Subscribers,omitempty"`

	Attributes []*Attribute `xml:"Attributes,omitempty"`
}

type TriggeredSendCreateResult struct {
	*CreateResult

	SubscriberFailures []*SubscriberResult `xml:"SubscriberFailures,omitempty"`
}

type SubscriberResult struct {
	Subscriber *Subscriber `xml:"Subscriber,omitempty"`

	ErrorCode string `xml:"ErrorCode,omitempty"`

	ErrorDescription string `xml:"ErrorDescription,omitempty"`

	Ordinal int32 `xml:"Ordinal,omitempty"`

	ErrorCodeID int32 `xml:"ErrorCodeID,omitempty"`
}

type SubscriberSendResult struct {
	*APIObject

	Send *Send `xml:"Send,omitempty"`

	Email *Email `xml:"Email,omitempty"`

	Subscriber *Subscriber `xml:"Subscriber,omitempty"`

	ClickDate time.Time `xml:"ClickDate,omitempty"`

	BounceDate time.Time `xml:"BounceDate,omitempty"`

	OpenDate time.Time `xml:"OpenDate,omitempty"`

	SentDate time.Time `xml:"SentDate,omitempty"`

	LastAction string `xml:"LastAction,omitempty"`

	UnsubscribeDate time.Time `xml:"UnsubscribeDate,omitempty"`

	FromAddress string `xml:"FromAddress,omitempty"`

	FromName string `xml:"FromName,omitempty"`

	TotalClicks int32 `xml:"TotalClicks,omitempty"`

	UniqueClicks int32 `xml:"UniqueClicks,omitempty"`

	Subject string `xml:"Subject,omitempty"`

	ViewSentEmailURL string `xml:"ViewSentEmailURL,omitempty"`

	HardBounces int32 `xml:"HardBounces,omitempty"`

	SoftBounces int32 `xml:"SoftBounces,omitempty"`

	OtherBounces int32 `xml:"OtherBounces,omitempty"`
}

type TriggeredSendSummary struct {
	*APIObject

	TriggeredSendDefinition *TriggeredSendDefinition `xml:"TriggeredSendDefinition,omitempty"`

	Sent int64 `xml:"Sent,omitempty"`

	NotSentDueToOptOut int64 `xml:"NotSentDueToOptOut,omitempty"`

	NotSentDueToUndeliverable int64 `xml:"NotSentDueToUndeliverable,omitempty"`

	Bounces int64 `xml:"Bounces,omitempty"`

	Opens int64 `xml:"Opens,omitempty"`

	Clicks int64 `xml:"Clicks,omitempty"`

	UniqueOpens int64 `xml:"UniqueOpens,omitempty"`

	UniqueClicks int64 `xml:"UniqueClicks,omitempty"`

	OptOuts int64 `xml:"OptOuts,omitempty"`

	SurveyResponses int64 `xml:"SurveyResponses,omitempty"`

	FTAFRequests int64 `xml:"FTAFRequests,omitempty"`

	FTAFEmailsSent int64 `xml:"FTAFEmailsSent,omitempty"`

	FTAFOptIns int64 `xml:"FTAFOptIns,omitempty"`

	Conversions int64 `xml:"Conversions,omitempty"`

	UniqueConversions int64 `xml:"UniqueConversions,omitempty"`

	InProcess int64 `xml:"InProcess,omitempty"`

	NotSentDueToError int64 `xml:"NotSentDueToError,omitempty"`

	Queued int64 `xml:"Queued,omitempty"`
}

type AsyncRequestResult struct {
	*APIObject

	Status string `xml:"Status,omitempty"`

	CompleteDate time.Time `xml:"CompleteDate,omitempty"`

	CallStatus string `xml:"CallStatus,omitempty"`

	CallMessage string `xml:"CallMessage,omitempty"`
}

type VoiceTriggeredSend struct {
	*APIObject

	VoiceTriggeredSendDefinition *VoiceTriggeredSendDefinition `xml:"VoiceTriggeredSendDefinition,omitempty"`

	Subscriber *Subscriber `xml:"Subscriber,omitempty"`

	Message string `xml:"Message,omitempty"`

	Number string `xml:"Number,omitempty"`

	TransferMessage string `xml:"TransferMessage,omitempty"`

	TransferNumber string `xml:"TransferNumber,omitempty"`
}

type VoiceTriggeredSendDefinition struct {
	*SendDefinition
}

type SMSTriggeredSend struct {
	*APIObject

	SMSTriggeredSendDefinition *SMSTriggeredSendDefinition `xml:"SMSTriggeredSendDefinition,omitempty"`

	Subscriber *Subscriber `xml:"Subscriber,omitempty"`

	Message string `xml:"Message,omitempty"`

	Number string `xml:"Number,omitempty"`

	FromAddress string `xml:"FromAddress,omitempty"`

	SmsSendId string `xml:"SmsSendId,omitempty"`
}

type SMSTriggeredSendDefinition struct {
	*SendDefinition

	Publication *List `xml:"Publication,omitempty"`

	DataExtension *DataExtension `xml:"DataExtension,omitempty"`

	Content *ContentArea `xml:"Content,omitempty"`

	SendToList bool `xml:"SendToList,omitempty"`
}

type SendClassification struct {
	*APIObject

	SendClassificationType *SendClassificationTypeEnum `xml:"SendClassificationType,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	SenderProfile *SenderProfile `xml:"SenderProfile,omitempty"`

	DeliveryProfile *DeliveryProfile `xml:"DeliveryProfile,omitempty"`

	HonorPublicationListOptOutsForTransactionalSends bool `xml:"HonorPublicationListOptOutsForTransactionalSends,omitempty"`

	SendPriority *SendPriorityEnum `xml:"SendPriority,omitempty"`

	ArchiveEmail bool `xml:"ArchiveEmail,omitempty"`
}

type SenderProfile struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	FromName string `xml:"FromName,omitempty"`

	FromAddress string `xml:"FromAddress,omitempty"`

	UseDefaultRMMRules bool `xml:"UseDefaultRMMRules,omitempty"`

	AutoForwardToEmailAddress string `xml:"AutoForwardToEmailAddress,omitempty"`

	AutoForwardToName string `xml:"AutoForwardToName,omitempty"`

	DirectForward bool `xml:"DirectForward,omitempty"`

	AutoForwardTriggeredSend *TriggeredSendDefinition `xml:"AutoForwardTriggeredSend,omitempty"`

	AutoReply bool `xml:"AutoReply,omitempty"`

	AutoReplyTriggeredSend *TriggeredSendDefinition `xml:"AutoReplyTriggeredSend,omitempty"`

	SenderHeaderEmailAddress string `xml:"SenderHeaderEmailAddress,omitempty"`

	SenderHeaderName string `xml:"SenderHeaderName,omitempty"`

	DataRetentionPeriodLength int16 `xml:"DataRetentionPeriodLength,omitempty"`

	DataRetentionPeriodUnitOfMeasure *RecurrenceTypeEnum `xml:"DataRetentionPeriodUnitOfMeasure,omitempty"`

	ReplyManagementRuleSet *APIObject `xml:"ReplyManagementRuleSet,omitempty"`

	ReplyToAddress string `xml:"ReplyToAddress,omitempty"`

	ReplyToDisplayName string `xml:"ReplyToDisplayName,omitempty"`

	FallbackFromAddress string `xml:"FallbackFromAddress,omitempty"`
}

type DeliveryProfile struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	SourceAddressType *DeliveryProfileSourceAddressTypeEnum `xml:"SourceAddressType,omitempty"`

	PrivateIP *PrivateIP `xml:"PrivateIP,omitempty"`

	DomainType *DeliveryProfileDomainTypeEnum `xml:"DomainType,omitempty"`

	PrivateDomain *PrivateDomain `xml:"PrivateDomain,omitempty"`

	HeaderSalutationSource *SalutationSourceEnum `xml:"HeaderSalutationSource,omitempty"`

	HeaderContentArea *ContentArea `xml:"HeaderContentArea,omitempty"`

	FooterSalutationSource *SalutationSourceEnum `xml:"FooterSalutationSource,omitempty"`

	FooterContentArea *ContentArea `xml:"FooterContentArea,omitempty"`

	SubscriberLevelPrivateDomain bool `xml:"SubscriberLevelPrivateDomain,omitempty"`

	SMIMESignatureCertificate *Certificate `xml:"SMIMESignatureCertificate,omitempty"`

	PrivateDomainSet *PrivateDomainSet `xml:"PrivateDomainSet,omitempty"`
}

type PrivateDomain struct {
	*APIObject
}

type PrivateDomainSet struct {
	*APIObject
}

type PrivateIP struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	IsActive bool `xml:"IsActive,omitempty"`

	OrdinalID int16 `xml:"OrdinalID,omitempty"`

	IPAddress string `xml:"IPAddress,omitempty"`
}

type SendDefinition struct {
	*InteractionDefinition

	CategoryID int32 `xml:"CategoryID,omitempty"`

	SendClassification *SendClassification `xml:"SendClassification,omitempty"`

	SenderProfile *SenderProfile `xml:"SenderProfile,omitempty"`

	// As of Fall 2007 SenderProfile.FromName should be used.
	FromName string `xml:"FromName,omitempty"`

	// As of Fall 2007 SenderProfile.FromAddress should be used.
	FromAddress string `xml:"FromAddress,omitempty"`

	DeliveryProfile *DeliveryProfile `xml:"DeliveryProfile,omitempty"`

	// As of Fall 2007 DeliveryProfile.SourceAddressType should be used.
	SourceAddressType *DeliveryProfileSourceAddressTypeEnum `xml:"SourceAddressType,omitempty"`

	// As of Fall 2007 DeliveryProfile.PrivateIP should be used.
	PrivateIP *PrivateIP `xml:"PrivateIP,omitempty"`

	// As of Fall 2007 DeliveryProfile.DomainType should be used.
	DomainType *DeliveryProfileDomainTypeEnum `xml:"DomainType,omitempty"`

	// As of Fall 2007 DeliveryProfile.PrivateDomain should be used.
	PrivateDomain *PrivateDomain `xml:"PrivateDomain,omitempty"`

	// As of Fall 2007 DeliveryProfile.HeaderSalutationSource should be used.
	HeaderSalutationSource *SalutationSourceEnum `xml:"HeaderSalutationSource,omitempty"`

	// As of Fall 2007 DeliveryProfile.HeaderContentArea should be used.
	HeaderContentArea *ContentArea `xml:"HeaderContentArea,omitempty"`

	// As of Fall 2007 DeliveryProfile.FooterSalutationSource should be used.
	FooterSalutationSource *SalutationSourceEnum `xml:"FooterSalutationSource,omitempty"`

	// As of Fall 2007 DeliveryProfile.FooterContentArea should be used.
	FooterContentArea *ContentArea `xml:"FooterContentArea,omitempty"`

	SuppressTracking bool `xml:"SuppressTracking,omitempty"`

	IsSendLogging bool `xml:"IsSendLogging,omitempty"`
}

type AudienceItem struct {
	*APIObject

	List *List `xml:"List,omitempty"`

	SendDefinitionListType *SendDefinitionListTypeEnum `xml:"SendDefinitionListType,omitempty"`

	CustomObjectID string `xml:"CustomObjectID,omitempty"`

	DataSourceTypeID *DataSourceTypeEnum `xml:"DataSourceTypeID,omitempty"`
}

type EmailSendDefinition struct {
	*SendDefinition

	SendDefinitionList []*SendDefinitionList `xml:"SendDefinitionList,omitempty"`

	Email *Email `xml:"Email,omitempty"`

	BccEmail string `xml:"BccEmail,omitempty"`

	AutoBccEmail string `xml:"AutoBccEmail,omitempty"`

	TestEmailAddr string `xml:"TestEmailAddr,omitempty"`

	EmailSubject string `xml:"EmailSubject,omitempty"`

	DynamicEmailSubject string `xml:"DynamicEmailSubject,omitempty"`

	IsMultipart bool `xml:"IsMultipart,omitempty"`

	IsWrapped bool `xml:"IsWrapped,omitempty"`

	SendLimit int32 `xml:"SendLimit,omitempty"`

	SendWindowOpen time.Time `xml:"SendWindowOpen,omitempty"`

	SendWindowClose time.Time `xml:"SendWindowClose,omitempty"`

	SendWindowDelete bool `xml:"SendWindowDelete,omitempty"`

	DeduplicateByEmail bool `xml:"DeduplicateByEmail,omitempty"`

	ExclusionFilter string `xml:"ExclusionFilter,omitempty"`

	TrackingUsers struct {
		TrackingUser []*TrackingUser `xml:"TrackingUser,omitempty"`
	} `xml:"TrackingUsers,omitempty"`

	Additional string `xml:"Additional,omitempty"`

	CCEmail string `xml:"CCEmail,omitempty"`

	DeliveryScheduledTime time.Time `xml:"DeliveryScheduledTime,omitempty"`

	MessageDeliveryType *MessageDeliveryTypeEnum `xml:"MessageDeliveryType,omitempty"`

	IsSeedListSend bool `xml:"IsSeedListSend,omitempty"`

	TimeZone *TimeZone `xml:"TimeZone,omitempty"`

	SeedListOccurance int32 `xml:"SeedListOccurance,omitempty"`

	PreHeader string `xml:"PreHeader,omitempty"`

	ReplyToAddress string `xml:"ReplyToAddress,omitempty"`

	ReplyToDisplayName string `xml:"ReplyToDisplayName,omitempty"`
}

type DeprecatedEmailSendDefinition struct {
	*EmailSendDefinition
}

type SendDefinitionList struct {
	*AudienceItem

	FilterDefinition *FilterDefinition `xml:"FilterDefinition,omitempty"`

	IsTestObject bool `xml:"IsTestObject,omitempty"`

	SalesForceObjectID string `xml:"SalesForceObjectID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Parameters struct {
		Parameter []*APIProperty `xml:"Parameter,omitempty"`
	} `xml:"Parameters,omitempty"`
}

type TrackingUser struct {
	*APIObject

	IsActive bool `xml:"IsActive,omitempty"`

	EmployeeID int32 `xml:"EmployeeID,omitempty"`
}

type MessagingVendorKind struct {
	*APIObject

	Vendor string `xml:"Vendor,omitempty"`

	Kind string `xml:"Kind,omitempty"`

	IsUsernameRequired bool `xml:"IsUsernameRequired,omitempty"`

	IsPasswordRequired bool `xml:"IsPasswordRequired,omitempty"`

	IsProfileRequired bool `xml:"IsProfileRequired,omitempty"`
}

type SMSMTEvent struct {
	*APIObject

	SMSTriggeredSend *SMSTriggeredSend `xml:"SMSTriggeredSend,omitempty"`

	Subscriber *Subscriber `xml:"Subscriber,omitempty"`

	MOCode string `xml:"MOCode,omitempty"`

	EventDate time.Time `xml:"EventDate,omitempty"`

	Carrier string `xml:"Carrier,omitempty"`
}

type SMSMOEvent struct {
	*APIObject

	Keyword *BaseMOKeyword `xml:"Keyword,omitempty"`

	MobileTelephoneNumber string `xml:"MobileTelephoneNumber,omitempty"`

	MOCode string `xml:"MOCode,omitempty"`

	EventDate time.Time `xml:"EventDate,omitempty"`

	MOMessage string `xml:"MOMessage,omitempty"`

	MTMessage string `xml:"MTMessage,omitempty"`

	Carrier string `xml:"Carrier,omitempty"`
}

type BaseMOKeyword struct {
	*APIObject

	IsDefaultKeyword bool `xml:"IsDefaultKeyword,omitempty"`
}

type SendSMSMOKeyword struct {
	*BaseMOKeyword

	NextMOKeyword *BaseMOKeyword `xml:"NextMOKeyword,omitempty"`

	Message string `xml:"Message,omitempty"`

	ScriptErrorMessage string `xml:"ScriptErrorMessage,omitempty"`
}

type UnsubscribeFromSMSPublicationMOKeyword struct {
	*BaseMOKeyword

	NextMOKeyword *BaseMOKeyword `xml:"NextMOKeyword,omitempty"`

	AllUnsubSuccessMessage string `xml:"AllUnsubSuccessMessage,omitempty"`

	InvalidPublicationMessage string `xml:"InvalidPublicationMessage,omitempty"`

	SingleUnsubSuccessMessage string `xml:"SingleUnsubSuccessMessage,omitempty"`
}

type DoubleOptInMOKeyword struct {
	*BaseMOKeyword

	DefaultPublication *List `xml:"DefaultPublication,omitempty"`

	InvalidPublicationMessage string `xml:"InvalidPublicationMessage,omitempty"`

	InvalidResponseMessage string `xml:"InvalidResponseMessage,omitempty"`

	MissingPublicationMessage string `xml:"MissingPublicationMessage,omitempty"`

	NeedPublicationMessage string `xml:"NeedPublicationMessage,omitempty"`

	PromptMessage string `xml:"PromptMessage,omitempty"`

	SuccessMessage string `xml:"SuccessMessage,omitempty"`

	UnexpectedErrorMessage string `xml:"UnexpectedErrorMessage,omitempty"`

	ValidPublications struct {
		ValidPublication []*List `xml:"ValidPublication,omitempty"`
	} `xml:"ValidPublications,omitempty"`

	ValidResponses struct {
		ValidResponse []string `xml:"ValidResponse,omitempty"`
	} `xml:"ValidResponses,omitempty"`
}

type HelpMOKeyword struct {
	*BaseMOKeyword

	FriendlyName string `xml:"FriendlyName,omitempty"`

	DefaultHelpMessage string `xml:"DefaultHelpMessage,omitempty"`

	MenuText string `xml:"MenuText,omitempty"`

	MoreChoicesPrompt string `xml:"MoreChoicesPrompt,omitempty"`
}

type SendEmailMOKeyword struct {
	*BaseMOKeyword

	SuccessMessage string `xml:"SuccessMessage,omitempty"`

	MissingEmailMessage string `xml:"MissingEmailMessage,omitempty"`

	FailureMessage string `xml:"FailureMessage,omitempty"`

	TriggeredSend *TriggeredSendDefinition `xml:"TriggeredSend,omitempty"`

	NextMOKeyword *BaseMOKeyword `xml:"NextMOKeyword,omitempty"`
}

type SMSSharedKeyword struct {
	*APIObject

	ShortCode string `xml:"ShortCode,omitempty"`

	SharedKeyword string `xml:"SharedKeyword,omitempty"`

	RequestDate time.Time `xml:"RequestDate,omitempty"`

	EffectiveDate time.Time `xml:"EffectiveDate,omitempty"`

	ExpireDate time.Time `xml:"ExpireDate,omitempty"`

	ReturnToPoolDate time.Time `xml:"ReturnToPoolDate,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`
}

type UserMap struct {
	*APIProperty

	ETAccountUser *AccountUser `xml:"ETAccountUser,omitempty"`

	AdditionalData []*APIProperty `xml:"AdditionalData,omitempty"`
}

type Folder struct {
	*APIProperty

	ID int32 `xml:"ID,omitempty"`

	ParentID int32 `xml:"ParentID,omitempty"`
}

type SalesforceSendActivity struct {
	*InteractionDefinition
}

type FileTransferLocation struct {
	*APIObject
}

type DataExtractActivity struct {
	*InteractionDefinition
}

type MessageSendActivity struct {
	*InteractionDefinition
}

type SmsSendActivity struct {
	*InteractionDefinition
}

type MobileConnectRefreshListActivity struct {
	*InteractionDefinition
}

type MobileConnectSendSmsActivity struct {
	*InteractionDefinition
}

type MobilePushSendMessageActivity struct {
	*InteractionDefinition
}

type ReportActivity struct {
	*InteractionDefinition
}

type DataExtension struct {
	*APIObject

	ObjectID string `xml:"ObjectID,omitempty" json:"objectId,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	IsSendable bool `xml:"IsSendable,omitempty"`

	IsTestable bool `xml:"IsTestable,omitempty"`

	SendableDataExtensionField *DataExtensionField `xml:"SendableDataExtensionField,omitempty"`

	SendableSubscriberField *Attribute `xml:"SendableSubscriberField,omitempty"`

	Template *DataExtensionTemplate `xml:"Template,omitempty"`

	DataRetentionPeriodLength int32 `xml:"DataRetentionPeriodLength,omitempty"`

	// Deprecated.  Use DataRetentionPeriod instead.
	DataRetentionPeriodUnitOfMeasure int32 `xml:"DataRetentionPeriodUnitOfMeasure,omitempty"`

	RowBasedRetention bool `xml:"RowBasedRetention,omitempty"`

	ResetRetentionPeriodOnImport bool `xml:"ResetRetentionPeriodOnImport,omitempty"`

	DeleteAtEndOfRetentionPeriod bool `xml:"DeleteAtEndOfRetentionPeriod,omitempty"`

	RetainUntil string `xml:"RetainUntil,omitempty"`

	Fields struct {
		Field []*DataExtensionField `xml:"Field,omitempty"`
	} `xml:"Fields,omitempty"`

	DataRetentionPeriod *DateTimeUnitOfMeasure `xml:"DataRetentionPeriod,omitempty"`

	CategoryID int64 `xml:"CategoryID,omitempty"`

	Status string `xml:"Status,omitempty"`
}

type DataExtensionField struct {
	*PropertyDefinition `json:"propertyDefinition,omitempty"`

	Ordinal       int32                          `xml:"Ordinal,omitempty" json:"ordinal,omitempty"`
	IsPrimaryKey  bool                           `xml:"IsPrimaryKey,omitempty" json:"isPrimaryKey,omitempty"`
	FieldType     *DataExtensionFieldType        `xml:"FieldType,omitempty" json:"fieldType,omitempty"`
	DataExtension *DataExtension                 `xml:"DataExtension,omitempty" json:"dataExtension,omitempty"`
	StorageType   *DataExtensionFieldStorageType `xml:"StorageType,omitempty" json:"storageType,omitempty"`
}

type DataExtensionTemplate struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`
}

type DataExtensionObject struct {
	*ObjectExtension

	Name string `xml:"Name,omitempty"`

	Keys struct {
		Key []*APIProperty `xml:"Key,omitempty"`
	} `xml:"Keys,omitempty"`
}

type DataExtensionError struct {
	Name string `xml:"Name,omitempty"`

	ErrorCode int32 `xml:"ErrorCode,omitempty"`

	ErrorMessage string `xml:"ErrorMessage,omitempty"`
}

type DataExtensionCreateResult struct {
	*CreateResult

	ErrorMessage string `xml:"ErrorMessage,omitempty"`

	KeyErrors struct {
		KeyError []*DataExtensionError `xml:"KeyError,omitempty"`
	} `xml:"KeyErrors,omitempty"`

	ValueErrors struct {
		ValueError []*DataExtensionError `xml:"ValueError,omitempty"`
	} `xml:"ValueErrors,omitempty"`
}

type DataExtensionUpdateResult struct {
	*UpdateResult

	ErrorMessage string `xml:"ErrorMessage,omitempty"`

	KeyErrors struct {
		KeyError []*DataExtensionError `xml:"KeyError,omitempty"`
	} `xml:"KeyErrors,omitempty"`

	ValueErrors struct {
		ValueError []*DataExtensionError `xml:"ValueError,omitempty"`
	} `xml:"ValueErrors,omitempty"`
}

type DataExtensionDeleteResult struct {
	*DeleteResult

	ErrorMessage string `xml:"ErrorMessage,omitempty"`

	KeyErrors struct {
		KeyError []*DataExtensionError `xml:"KeyError,omitempty"`
	} `xml:"KeyErrors,omitempty"`
}

type ImportDefinitionColumnBasedAction struct {
	Value string `xml:"Value,omitempty"`

	Action *ImportDefinitionColumnBasedActionType `xml:"Action,omitempty"`
}

type FieldMap struct {
	DestinationName string `xml:"DestinationName,omitempty"`

	SourceName string `xml:"SourceName,omitempty"`

	SourceOrdinal int32 `xml:"SourceOrdinal,omitempty"`
}

type ImportDefinitionAutoGenerateDestination struct {
	DataExtensionTarget *DataExtension `xml:"DataExtensionTarget,omitempty"`

	ErrorIfExists bool `xml:"ErrorIfExists,omitempty"`
}

type ImportDefinition struct {
	*InteractionDefinition

	AllowErrors bool `xml:"AllowErrors,omitempty"`

	DestinationObject *APIObject `xml:"DestinationObject,omitempty"`

	FieldMappingType *ImportDefinitionFieldMappingType `xml:"FieldMappingType,omitempty"`

	FieldMaps struct {
		FieldMap []*FieldMap `xml:"FieldMap,omitempty"`
	} `xml:"FieldMaps,omitempty"`

	FileSpec string `xml:"FileSpec,omitempty"`

	FileType *FileType `xml:"FileType,omitempty"`

	Notification *AsyncResponse `xml:"Notification,omitempty"`

	RetrieveFileTransferLocation *FileTransferLocation `xml:"RetrieveFileTransferLocation,omitempty"`

	SubscriberImportType *ImportDefinitionSubscriberImportType `xml:"SubscriberImportType,omitempty"`

	UpdateType *ImportDefinitionUpdateType `xml:"UpdateType,omitempty"`

	MaxFileAge int32 `xml:"MaxFileAge,omitempty"`

	MaxFileAgeScheduleOffset int32 `xml:"MaxFileAgeScheduleOffset,omitempty"`

	MaxImportFrequency int32 `xml:"MaxImportFrequency,omitempty"`

	Delimiter string `xml:"Delimiter,omitempty"`

	HeaderLines int32 `xml:"HeaderLines,omitempty"`

	AutoGenerateDestination *ImportDefinitionAutoGenerateDestination `xml:"AutoGenerateDestination,omitempty"`

	ControlColumn string `xml:"ControlColumn,omitempty"`

	ControlColumnDefaultAction *ImportDefinitionColumnBasedActionType `xml:"ControlColumnDefaultAction,omitempty"`

	ControlColumnActions struct {
		ControlColumnAction []*ImportDefinitionColumnBasedAction `xml:"ControlColumnAction,omitempty"`
	} `xml:"ControlColumnActions,omitempty"`

	EndOfLineRepresentation string `xml:"EndOfLineRepresentation,omitempty"`

	NullRepresentation string `xml:"NullRepresentation,omitempty"`

	StandardQuotedStrings bool `xml:"StandardQuotedStrings,omitempty"`

	Filter string `xml:"Filter,omitempty"`

	DateFormattingLocale *Locale `xml:"DateFormattingLocale,omitempty"`

	DeleteFile bool `xml:"DeleteFile,omitempty"`

	SourceObject *APIObject `xml:"SourceObject,omitempty"`

	DestinationType int32 `xml:"DestinationType,omitempty"`

	SubscriptionDefinitionId string `xml:"SubscriptionDefinitionId,omitempty"`

	EncodingCodePage int32 `xml:"EncodingCodePage,omitempty"`

	SmsMemberSharedShortCodeId string `xml:"SmsMemberSharedShortCodeId,omitempty"`

	HasMultipleFiles bool `xml:"HasMultipleFiles,omitempty"`
}

type ImportDefinitionFieldMap struct {
	*APIObject

	DestinationName string `xml:"DestinationName,omitempty"`
}

type ImportResultsSummary struct {
	*APIObject

	ImportDefinitionCustomerKey string `xml:"ImportDefinitionCustomerKey,omitempty"`

	StartDate string `xml:"StartDate,omitempty"`

	EndDate string `xml:"EndDate,omitempty"`

	DestinationID string `xml:"DestinationID,omitempty"`

	NumberSuccessful int32 `xml:"NumberSuccessful,omitempty"`

	NumberDuplicated int32 `xml:"NumberDuplicated,omitempty"`

	NumberRestricted int32 `xml:"NumberRestricted,omitempty"`

	NumberErrors int32 `xml:"NumberErrors,omitempty"`

	TotalRows int32 `xml:"TotalRows,omitempty"`

	ImportType string `xml:"ImportType,omitempty"`

	ImportStatus string `xml:"ImportStatus,omitempty"`

	TaskResultID int32 `xml:"TaskResultID,omitempty"`
}

type FilterActivity struct {
	*InteractionDefinition

	FilterActivityID string `xml:"FilterActivityID,omitempty"`

	FilterDefinitionID string `xml:"FilterDefinitionID,omitempty"`

	DestinationObjectID string `xml:"DestinationObjectID,omitempty"`

	DestinationTypeID int32 `xml:"DestinationTypeID,omitempty"`

	SourceObjectID string `xml:"SourceObjectID,omitempty"`

	SourceTypeID int32 `xml:"SourceTypeID,omitempty"`

	OwnerID string `xml:"OwnerID,omitempty"`

	StatusID string `xml:"StatusID,omitempty"`

	CreatedBy *ClientID `xml:"CreatedBy,omitempty"`

	ModifiedBy *ClientID `xml:"ModifiedBy,omitempty"`
}

type FilterDefinition struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	DataSource *APIObject `xml:"DataSource,omitempty"`

	DataFilter *FilterPart `xml:"DataFilter,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`
}

type GroupDefinition struct {
	*InteractionDefinition
}

type GroupConnectActivity struct {
	*InteractionDefinition
}

type FileTransferActivity struct {
	*InteractionDefinition
}

type ListSend struct {
	*APIObject

	SendID int32 `xml:"SendID,omitempty"`

	List *List `xml:"List,omitempty"`

	Duplicates int32 `xml:"Duplicates,omitempty"`

	InvalidAddresses int32 `xml:"InvalidAddresses,omitempty"`

	ExistingUndeliverables int32 `xml:"ExistingUndeliverables,omitempty"`

	ExistingUnsubscribes int32 `xml:"ExistingUnsubscribes,omitempty"`

	HardBounces int32 `xml:"HardBounces,omitempty"`

	SoftBounces int32 `xml:"SoftBounces,omitempty"`

	OtherBounces int32 `xml:"OtherBounces,omitempty"`

	ForwardedEmails int32 `xml:"ForwardedEmails,omitempty"`

	UniqueClicks int32 `xml:"UniqueClicks,omitempty"`

	UniqueOpens int32 `xml:"UniqueOpens,omitempty"`

	NumberSent int32 `xml:"NumberSent,omitempty"`

	NumberDelivered int32 `xml:"NumberDelivered,omitempty"`

	Unsubscribes int32 `xml:"Unsubscribes,omitempty"`

	MissingAddresses int32 `xml:"MissingAddresses,omitempty"`

	PreviewURL string `xml:"PreviewURL,omitempty"`

	Links []*Link `xml:"Links,omitempty"`

	Events []*TrackingEvent `xml:"Events,omitempty"`
}

type LinkSend struct {
	*APIObject

	SendID int32 `xml:"SendID,omitempty"`

	Link *Link `xml:"Link,omitempty"`

	IDLong int64 `xml:"IDLong,omitempty"`
}

type ObjectExtension struct {
	*APIObject

	Type string `xml:"Type,omitempty"`

	Properties struct {
		Property []*APIProperty `xml:"Property,omitempty"`
	} `xml:"Properties,omitempty"`
}

type PublicKeyManagement struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Key []byte `xml:"Key,omitempty"`
}

type SecurityObject struct {
	*APIObject
}

type Certificate struct {
	*SecurityObject
}

type SystemStatusOptions struct {
	*Options
}

type SystemStatusResult struct {
	*Result

	SystemStatus *SystemStatusType `xml:"SystemStatus,omitempty"`

	Outages struct {
		Outage []*SystemOutage `xml:"Outage,omitempty"`
	} `xml:"Outages,omitempty"`
}

type SystemOutage struct {
}

type Authentication struct {
	*APIObject
}

type UsernameAuthentication struct {
	*Authentication

	UserName string `xml:"UserName,omitempty"`

	PassWord string `xml:"PassWord,omitempty"`
}

type ResourceSpecification struct {
	*APIObject

	URN string `xml:"URN,omitempty"`

	Authentication *Authentication `xml:"Authentication,omitempty"`
}

type Portfolio struct {
	*APIObject

	Source *ResourceSpecification `xml:"Source,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	FileName string `xml:"FileName,omitempty"`

	DisplayName string `xml:"DisplayName,omitempty"`

	Description string `xml:"Description,omitempty"`

	TypeDescription string `xml:"TypeDescription,omitempty"`

	IsUploaded bool `xml:"IsUploaded,omitempty"`

	IsActive bool `xml:"IsActive,omitempty"`

	FileSizeKB int32 `xml:"FileSizeKB,omitempty"`

	ThumbSizeKB int32 `xml:"ThumbSizeKB,omitempty"`

	FileWidthPX int32 `xml:"FileWidthPX,omitempty"`

	FileHeightPX int32 `xml:"FileHeightPX,omitempty"`

	FileURL string `xml:"FileURL,omitempty"`

	ThumbURL string `xml:"ThumbURL,omitempty"`

	CacheClearTime time.Time `xml:"CacheClearTime,omitempty"`

	CategoryType string `xml:"CategoryType,omitempty"`
}

type Template struct {
	*APIObject

	TemplateName string `xml:"TemplateName,omitempty"`

	LayoutHTML string `xml:"LayoutHTML,omitempty"`

	BackgroundColor string `xml:"BackgroundColor,omitempty"`

	BorderColor string `xml:"BorderColor,omitempty"`

	BorderWidth int32 `xml:"BorderWidth,omitempty"`

	Cellpadding int32 `xml:"Cellpadding,omitempty"`

	Cellspacing int32 `xml:"Cellspacing,omitempty"`

	Width int32 `xml:"Width,omitempty"`

	Align string `xml:"Align,omitempty"`

	ActiveFlag int32 `xml:"ActiveFlag,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`

	CategoryType string `xml:"CategoryType,omitempty"`

	OwnerID int32 `xml:"OwnerID,omitempty"`

	HeaderContent *ContentArea `xml:"HeaderContent,omitempty"`

	Layout *Layout `xml:"Layout,omitempty"`

	TemplateSubject string `xml:"TemplateSubject,omitempty"`

	IsTemplateSubjectLocked bool `xml:"IsTemplateSubjectLocked,omitempty"`

	PreHeader string `xml:"PreHeader,omitempty"`
}

type Layout struct {
	*APIObject

	LayoutName string `xml:"LayoutName,omitempty"`
}

type QueryDefinition struct {
	*InteractionDefinition

	QueryText string `xml:"QueryText,omitempty"`

	TargetType string `xml:"TargetType,omitempty"`

	DataExtensionTarget *InteractionBaseObject `xml:"DataExtensionTarget,omitempty"`

	TargetUpdateType string `xml:"TargetUpdateType,omitempty"`

	FileSpec string `xml:"FileSpec,omitempty"`

	FileType string `xml:"FileType,omitempty"`

	Status string `xml:"Status,omitempty"`

	CategoryID int32 `xml:"CategoryID,omitempty"`
}

type JsonWebKey struct {
	KeyType string `xml:"KeyType,omitempty"`

	KeyModulus string `xml:"KeyModulus,omitempty"`

	KeyExponent string `xml:"KeyExponent,omitempty"`

	KeyAlgorithm string `xml:"KeyAlgorithm,omitempty"`

	KeyUse string `xml:"KeyUse,omitempty"`

	KeyId string `xml:"KeyId,omitempty"`
}

type DirectoryTenant struct {
	CustomerId string `xml:"CustomerId,omitempty"`

	TenantOrgId string `xml:"TenantOrgId,omitempty"`

	TenantUri string `xml:"TenantUri,omitempty"`

	TenantType string `xml:"TenantType,omitempty"`

	CustomDomain string `xml:"CustomDomain,omitempty"`

	TenantPublicKeysUri string `xml:"TenantPublicKeysUri,omitempty"`

	TenantName string `xml:"TenantName,omitempty"`

	Description string `xml:"Description,omitempty"`

	PublicKeys []*JsonWebKey `xml:"PublicKeys,omitempty"`
}

type AuditLogUserContext struct {
	InitiatingUserName string `xml:"InitiatingUserName,omitempty"`

	InitiatingUserIpAddress string `xml:"InitiatingUserIpAddress,omitempty"`
}

type IntegrationProfile struct {
	*APIObject

	ProfileID string `xml:"ProfileID,omitempty"`

	SubscriberKey string `xml:"SubscriberKey,omitempty"`

	ExternalID string `xml:"ExternalID,omitempty"`

	ExternalType string `xml:"ExternalType,omitempty"`
}

type IntegrationProfileDefinition struct {
	*APIObject

	ProfileID string `xml:"ProfileID,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	ExternalSystemType int32 `xml:"ExternalSystemType,omitempty"`
}

type ReplyMailManagementConfiguration struct {
	*APIObject

	EmailDisplayName string `xml:"EmailDisplayName,omitempty"`

	ReplySubdomain string `xml:"ReplySubdomain,omitempty"`

	EmailReplyAddress string `xml:"EmailReplyAddress,omitempty"`

	DNSRedirectComplete bool `xml:"DNSRedirectComplete,omitempty"`

	DeleteAutoReplies bool `xml:"DeleteAutoReplies,omitempty"`

	SupportUnsubscribes bool `xml:"SupportUnsubscribes,omitempty"`

	SupportUnsubKeyword bool `xml:"SupportUnsubKeyword,omitempty"`

	SupportUnsubscribeKeyword bool `xml:"SupportUnsubscribeKeyword,omitempty"`

	SupportRemoveKeyword bool `xml:"SupportRemoveKeyword,omitempty"`

	SupportOptOutKeyword bool `xml:"SupportOptOutKeyword,omitempty"`

	SupportLeaveKeyword bool `xml:"SupportLeaveKeyword,omitempty"`

	SupportMisspelledKeywords bool `xml:"SupportMisspelledKeywords,omitempty"`

	SendAutoReplies bool `xml:"SendAutoReplies,omitempty"`

	AutoReplySubject string `xml:"AutoReplySubject,omitempty"`

	AutoReplyBody string `xml:"AutoReplyBody,omitempty"`

	ForwardingAddress string `xml:"ForwardingAddress,omitempty"`
}

type FileTrigger struct {
	*APIObject

	ExternalReference string `xml:"ExternalReference,omitempty"`

	Type string `xml:"Type,omitempty"`

	Status string `xml:"Status,omitempty"`

	StatusMessage string `xml:"StatusMessage,omitempty"`

	RequestParameterDetail string `xml:"RequestParameterDetail,omitempty"`

	ResponseControlManifest string `xml:"ResponseControlManifest,omitempty"`

	FileName string `xml:"FileName,omitempty"`

	Description string `xml:"Description,omitempty"`

	Name string `xml:"Name,omitempty"`

	LastPullDate time.Time `xml:"LastPullDate,omitempty"`

	ScheduledDate time.Time `xml:"ScheduledDate,omitempty"`

	IsActive bool `xml:"IsActive,omitempty"`

	FileTriggerProgramID string `xml:"FileTriggerProgramID,omitempty"`
}

type FileTriggerTypeLastPull struct {
	*APIObject

	ExternalReference string `xml:"ExternalReference,omitempty"`

	Type string `xml:"Type,omitempty"`

	LastPullDate time.Time `xml:"LastPullDate,omitempty"`
}

type ProgramManifestTemplate struct {
	*APIObject

	Type string `xml:"Type,omitempty"`

	OperationType string `xml:"OperationType,omitempty"`

	Content string `xml:"Content,omitempty"`
}

type SubscriberAddress struct {
	AddressType string `xml:"AddressType,omitempty"`

	Address string `xml:"Address,omitempty"`

	Statuses struct {
		Status []*AddressStatus `xml:"Status,omitempty"`
	} `xml:"Statuses,omitempty"`
}

type SMSAddress struct {
	*SubscriberAddress

	Carrier string `xml:"Carrier,omitempty"`
}

type EmailAddress struct {
	*SubscriberAddress

	Type *EmailType `xml:"Type,omitempty"`
}

type AddressStatus struct {
	Status *SubscriberAddressStatus `xml:"Status,omitempty"`
}

type Publication struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	IsActive bool `xml:"IsActive,omitempty"`

	SendClassification *SendClassification `xml:"SendClassification,omitempty"`

	Subscribers struct {
		Subscriber []*Subscriber `xml:"Subscriber,omitempty"`
	} `xml:"Subscribers,omitempty"`

	Category int32 `xml:"Category,omitempty"`
}

type PublicationSubscriber struct {
	*APIObject

	Publication *Publication `xml:"Publication,omitempty"`

	Subscriber *Subscriber `xml:"Subscriber,omitempty"`
}

type Automation struct {
	*InteractionDefinition

	Schedule *ScheduleDefinition `xml:"Schedule,omitempty"`

	AutomationTasks struct {
		AutomationTask []*AutomationTask `xml:"AutomationTask,omitempty"`
	} `xml:"AutomationTasks,omitempty"`

	IsActive bool `xml:"IsActive,omitempty"`

	AutomationSource *AutomationSource `xml:"AutomationSource,omitempty"`

	Status int32 `xml:"Status,omitempty"`

	Notifications struct {
		Notification []*AutomationNotification `xml:"Notification,omitempty"`
	} `xml:"Notifications,omitempty"`

	ScheduledTime time.Time `xml:"ScheduledTime,omitempty"`

	AutomationType string `xml:"AutomationType,omitempty"`

	UpdateModified bool `xml:"UpdateModified,omitempty"`

	LastRunInstanceID string `xml:"LastRunInstanceID,omitempty"`

	CreatedBy int64 `xml:"CreatedBy,omitempty"`

	CategoryID string `xml:"CategoryID,omitempty"`

	LastRunTime time.Time `xml:"LastRunTime,omitempty"`

	LastSaveDate time.Time `xml:"LastSaveDate,omitempty"`

	ModifiedBy int64 `xml:"ModifiedBy,omitempty"`

	RecurrenceID string `xml:"RecurrenceID,omitempty"`

	LastSavedBy int64 `xml:"LastSavedBy,omitempty"`
}

type AutomationSource struct {
	AutomationSourceID string `xml:"AutomationSourceID,omitempty"`

	AutomationSourceType string `xml:"AutomationSourceType,omitempty"`
}

type AutomationInstances struct {
	*APIObject

	InstanceCount int32 `xml:"InstanceCount,omitempty"`

	AutomationInstanceCollection struct {
		AutomationInstance []*AutomationInstance `xml:"AutomationInstance,omitempty"`
	} `xml:"AutomationInstanceCollection,omitempty"`
}

type AutomationInstance struct {
	*Automation

	AutomationID string `xml:"AutomationID,omitempty"`

	StatusMessage string `xml:"StatusMessage,omitempty"`

	StatusLastUpdate time.Time `xml:"StatusLastUpdate,omitempty"`

	TaskInstances struct {
		AutomationTaskInstance []*AutomationTaskInstance `xml:"AutomationTaskInstance,omitempty"`
	} `xml:"TaskInstances,omitempty"`

	StartTime time.Time `xml:"StartTime,omitempty"`

	CompletedTime time.Time `xml:"CompletedTime,omitempty"`
}

type AutomationNotification struct {
	*APIObject

	Address string `xml:"Address,omitempty"`

	Body string `xml:"Body,omitempty"`

	ChannelType string `xml:"ChannelType,omitempty"`

	NotificationType string `xml:"NotificationType,omitempty"`

	AutomationID string `xml:"AutomationID,omitempty"`
}

type AutomationTask struct {
	*APIObject

	AutomationTaskType string `xml:"AutomationTaskType,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	Automation *Automation `xml:"Automation,omitempty"`

	Sequence int32 `xml:"Sequence,omitempty"`

	Activities struct {
		Activity []*AutomationActivity `xml:"Activity,omitempty"`
	} `xml:"Activities,omitempty"`
}

type AutomationTaskInstance struct {
	*AutomationTask

	StepDefinition *AutomationTask `xml:"StepDefinition,omitempty"`

	AutomationInstance *AutomationInstance `xml:"AutomationInstance,omitempty"`

	ActivityInstances struct {
		ActivityInstance []*AutomationActivityInstance `xml:"ActivityInstance,omitempty"`
	} `xml:"ActivityInstances,omitempty"`
}

type AutomationActivity struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	IsActive bool `xml:"IsActive,omitempty"`

	Definition *APIObject `xml:"Definition,omitempty"`

	Automation *Automation `xml:"Automation,omitempty"`

	AutomationTask *AutomationTask `xml:"AutomationTask,omitempty"`

	Sequence int32 `xml:"Sequence,omitempty"`

	ActivityObject *APIObject `xml:"ActivityObject,omitempty"`

	SerializedObject string `xml:"SerializedObject,omitempty"`
}

type AutomationActivityInstance struct {
	*AutomationActivity

	ActivityID string `xml:"ActivityID,omitempty"`

	AutomationID string `xml:"AutomationID,omitempty"`

	SequenceID int32 `xml:"SequenceID,omitempty"`

	Status int32 `xml:"Status,omitempty"`

	StatusLastUpdate time.Time `xml:"StatusLastUpdate,omitempty"`

	StatusMessage string `xml:"StatusMessage,omitempty"`

	ActivityDefinition *AutomationActivity `xml:"ActivityDefinition,omitempty"`

	AutomationInstance *AutomationInstance `xml:"AutomationInstance,omitempty"`

	AutomationTaskInstance *AutomationTaskInstance `xml:"AutomationTaskInstance,omitempty"`

	ScheduledTime time.Time `xml:"ScheduledTime,omitempty"`

	StartTime time.Time `xml:"StartTime,omitempty"`

	CompletedTime time.Time `xml:"CompletedTime,omitempty"`
}

type AutomationChain struct {
	*InteractionDefinition

	AutomationToChainID string `xml:"AutomationToChainID,omitempty"`
}

type PlatformApplication struct {
	*APIObject

	Package *PlatformApplicationPackage `xml:"Package,omitempty"`

	Packages []*PlatformApplicationPackage `xml:"Packages,omitempty"`

	ResourceSpecification *ResourceSpecification `xml:"ResourceSpecification,omitempty"`

	DeveloperVersion string `xml:"DeveloperVersion,omitempty"`
}

type PlatformApplicationPackage struct {
	*APIObject

	ResourceSpecification *ResourceSpecification `xml:"ResourceSpecification,omitempty"`

	SigningKey *PublicKeyManagement `xml:"SigningKey,omitempty"`

	IsUpgrade bool `xml:"IsUpgrade,omitempty"`

	DeveloperVersion string `xml:"DeveloperVersion,omitempty"`
}

type SuppressionListDefinition struct {
	*APIObject

	Name string `xml:"Name,omitempty"`

	Category int64 `xml:"Category,omitempty"`

	Description string `xml:"Description,omitempty"`

	Contexts struct {
		Context []*SuppressionListContext `xml:"Context,omitempty"`
	} `xml:"Contexts,omitempty"`

	Fields struct {
		Field []*DataExtensionField `xml:"Field,omitempty"`
	} `xml:"Fields,omitempty"`

	SubscriberCount int64 `xml:"SubscriberCount,omitempty"`

	NotifyEmail string `xml:"NotifyEmail,omitempty"`
}

type SuppressionListContext struct {
	*APIObject

	Context *SuppressionListContextEnum `xml:"Context,omitempty"`

	SendClassificationType *SendClassificationTypeEnum `xml:"SendClassificationType,omitempty"`

	SendClassification *SendClassification `xml:"SendClassification,omitempty"`

	Send *Send `xml:"Send,omitempty"`

	Definition *SuppressionListDefinition `xml:"Definition,omitempty"`

	AppliesToAllSends bool `xml:"AppliesToAllSends,omitempty"`

	SenderProfile *SenderProfile `xml:"SenderProfile,omitempty"`
}

type SuppressionListData struct {
	*APIObject

	Properties struct {
		Property []*APIProperty `xml:"Property,omitempty"`
	} `xml:"Properties,omitempty"`
}

type SendAdditionalAttribute struct {
	*APIObject

	Email *Email `xml:"Email,omitempty"`

	Name string `xml:"Name,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type AttributeEntityV1 struct {
	Name string `xml:"Name,omitempty"`

	DisplayName string `xml:"DisplayName,omitempty"`

	Value string `xml:"Value,omitempty"`

	Order int32 `xml:"Order,omitempty"`

	Channel string `xml:"Channel,omitempty"`

	AttributeType string `xml:"AttributeType,omitempty"`
}

type Thumbnail struct {
	ThumbnailHtml string `xml:"ThumbnailHtml,omitempty"`

	ThumbnailImage string `xml:"ThumbnailImage,omitempty"`

	ThumbnailUrl string `xml:"ThumbnailUrl,omitempty"`
}

type NameIdReference struct {
	Id int64 `xml:"Id,omitempty"`

	Name string `xml:"Name,omitempty"`

	DisplayName string `xml:"DisplayName,omitempty"`
}

type CategoryNameIdReference struct {
	Id int32 `xml:"Id,omitempty"`

	Name string `xml:"Name,omitempty"`

	ParentId int32 `xml:"ParentId,omitempty"`
}

type UserBasicsEntity struct {
	Id int32 `xml:"Id,omitempty"`

	Email string `xml:"Email,omitempty"`

	Name string `xml:"Name,omitempty"`

	UserId string `xml:"UserId,omitempty"`
}

type AssetAnyProperty struct {
}

type Asset struct {
	*APIObject

	ContentType string `xml:"ContentType,omitempty"`

	Version int32 `xml:"Version,omitempty"`

	Locked bool `xml:"Locked,omitempty"`

	Name string `xml:"Name,omitempty"`

	Description string `xml:"Description,omitempty"`

	ActiveDate time.Time `xml:"ActiveDate,omitempty"`

	ExpirationDate time.Time `xml:"ExpirationDate,omitempty"`

	MemberId int64 `xml:"MemberId,omitempty"`

	EnterpriseId int64 `xml:"EnterpriseId,omitempty"`

	CreatedBy *UserBasicsEntity `xml:"CreatedBy,omitempty"`

	ModifiedBy *UserBasicsEntity `xml:"ModifiedBy,omitempty"`

	Content string `xml:"Content,omitempty"`

	Design string `xml:"Design,omitempty"`

	SuperContent string `xml:"SuperContent,omitempty"`

	MinBlocks int32 `xml:"MinBlocks,omitempty"`

	MaxBlocks int32 `xml:"MaxBlocks,omitempty"`

	File string `xml:"File,omitempty"`

	AssetType *NameIdReference `xml:"AssetType,omitempty"`

	Status *NameIdReference `xml:"Status,omitempty"`

	Thumbnail *Thumbnail `xml:"Thumbnail,omitempty"`

	GenerateFrom string `xml:"GenerateFrom,omitempty"`

	Template *Asset `xml:"Template,omitempty"`

	Category *CategoryNameIdReference `xml:"Category,omitempty"`

	Data *AssetAnyProperty `xml:"Data,omitempty"`

	FileProperties *AssetAnyProperty `xml:"FileProperties,omitempty"`

	Meta *AssetAnyProperty `xml:"Meta,omitempty"`

	CustomFields *AssetAnyProperty `xml:"CustomFields,omitempty"`

	SharingProperties *AssetAnyProperty `xml:"SharingProperties,omitempty"`

	Views *AssetAnyProperty `xml:"Views,omitempty"`

	Blocks *AssetAnyProperty `xml:"Blocks,omitempty"`

	Slots *AssetAnyProperty `xml:"Slots,omitempty"`

	Channels *AssetAnyProperty `xml:"Channels,omitempty"`

	AllowedBlocks struct {
		Block []string `xml:"Block,omitempty"`
	} `xml:"AllowedBlocks,omitempty"`

	Tags struct {
		Tag []string `xml:"Tag,omitempty"`
	} `xml:"Tags,omitempty"`

	Attributes struct {
		Attribute []*AttributeEntityV1 `xml:"Attribute,omitempty"`
	} `xml:"Attributes,omitempty"`
}

type Category struct {
	*APIObject

	Id int32 `xml:"Id,omitempty"`

	Name string `xml:"Name,omitempty"`

	ParentId int32 `xml:"ParentId,omitempty"`

	MemberId int64 `xml:"MemberId,omitempty"`

	EnterpriseId int64 `xml:"EnterpriseId,omitempty"`

	CategoryType string `xml:"CategoryType,omitempty"`

	Description string `xml:"Description,omitempty"`

	SharingProperties *AssetAnyProperty `xml:"SharingProperties,omitempty"`
}

type ImportFileDestination struct {
	*APIObject

	TemplateCustomObject *DataExtension `xml:"TemplateCustomObject,omitempty"`

	FileTransferLocation *FileTransferLocation `xml:"FileTransferLocation,omitempty"`

	FileSpec string `xml:"FileSpec,omitempty"`

	EncodingCodePage int32 `xml:"EncodingCodePage,omitempty"`

	HasColumnHeader bool `xml:"HasColumnHeader,omitempty"`

	FieldDelimiter string `xml:"FieldDelimiter,omitempty"`

	RowDelimiter string `xml:"RowDelimiter,omitempty"`

	NullValue string `xml:"NullValue,omitempty"`

	BooleanFormat string `xml:"BooleanFormat,omitempty"`

	DateTimeFormat string `xml:"DateTimeFormat,omitempty"`

	StringIdentifier string `xml:"StringIdentifier,omitempty"`

	EscapeSequence string `xml:"EscapeSequence,omitempty"`
}

type ContactEvent struct {
	*APIObject

	ContactID int64 `xml:"ContactID,omitempty"`

	ContactKey string `xml:"ContactKey,omitempty"`

	EventDefinitionKey string `xml:"EventDefinitionKey,omitempty"`

	Data struct {
		AttributeSet []*AttributeSet `xml:"AttributeSet,omitempty"`
	} `xml:"Data,omitempty"`
}

type AttributeSet struct {
	*APIObject

	Id string `xml:"Id,omitempty"`

	Key string `xml:"Key,omitempty"`

	Name string `xml:"Name,omitempty"`

	Items struct {
		Item []*AttributeValueContainer `xml:"Item,omitempty"`
	} `xml:"Items,omitempty"`
}

type AttributeValueContainer struct {
	Values struct {
		Value []*AttributeValue `xml:"Value,omitempty"`
	} `xml:"Values,omitempty"`
}

type AttributeValue struct {
	Id string `xml:"Id,omitempty"`

	Key string `xml:"Key,omitempty"`

	Name string `xml:"Name,omitempty"`

	Value string `xml:"Value,omitempty"`
}

type ContactEventCreateResult struct {
	*CreateResult

	EventInstanceID string `xml:"EventInstanceID,omitempty"`

	AsyncRequestID int64 `xml:"AsyncRequestID,omitempty"`
}

type ScheduledRequest struct {
	*APIObject

	RequestID string `xml:"RequestID,omitempty"`
}

type ScheduledConversation struct {
	*APIObject

	ConversationID string `xml:"ConversationID,omitempty"`
}

type APIFault struct {
	XMLName xml.Name `xml:"urn:fault.partner.exacttarget.com apifault"`

	Code int32 `xml:"Code,omitempty"`

	Message string `xml:"Message,omitempty"`

	LogID int64 `xml:"LogID,omitempty"`

	Params struct {
		Param []string `xml:"Param,omitempty"`
	} `xml:"Params,omitempty"`
}

type Soap interface {

	/* Create objects */
	Create(request *CreateRequest) (*CreateResponse, error)

	/* Retrieve objects */
	Retrieve(request *RetrieveRequestMsg) (*RetrieveResponseMsg, error)
	// Subscribers
	RetrieveRecipients(request *RetrieveRequestMsg) (*RetrieveRecipientResponseMsg, error)
	RetrieveRecipient(request *RetrieveRequestMsg) (*RetrieveRecipientResponseMsg, error)

	// Data Extensions
	RetrieveDataExtensions(request *RetrieveRequestMsg) (*RetrieveDEResponseMsg, error)
	//RetrieveDataExtensionByCustomerKey(request *RetrieveRequestMsg, dataExtensionCustomerKey types.DataExtensionRequest) (*RetrieveDEResponseMsg, error)
	RetrieveDataExtensionByCustomerKey(request *RetrieveRequestMsg) (*RetrieveDEResponseMsg, error)
	//RetrieveDataExtensionFieldsByDataExtensionCustomerKey(request *RetrieveRequestMsg, dataExtensionCustomerKey types.DataExtensionRequest) (*RetrieveDEResponseMsg, error)
	RetrieveDataExtensionFieldsByDataExtensionCustomerKey(request *RetrieveRequestMsg) (*RetrieveDEFieldResponseMsg, error)

	RetrieveDataExtensionFields(request *RetrieveRequestMsg) (*RetrieveDEFieldResponseMsg, error)
	RetrieveDataExtensionObject(request *RetrieveRequestMsg) (*RetrieveDataExtensionObjectResponseMsg, error)

	/* Update objects */
	Update(request *UpdateRequest) (*UpdateResponse, error)

	/* Delete objects */
	Delete(request *DeleteRequest) (*DeleteResponse, error)

	Query(request *QueryRequestMsg) (*QueryResponseMsg, error)

	Describe(request *DefinitionRequestMsg) (*DefinitionResponseMsg, error)

	Execute(request *ExecuteRequestMsg) (*ExecuteResponseMsg, error)

	Perform(request *PerformRequestMsg) (*PerformResponseMsg, error)

	Configure(request *ConfigureRequestMsg) (*ConfigureResponseMsg, error)

	Schedule(request *ScheduleRequestMsg) (*ScheduleResponseMsg, error)

	VersionInfo(request *VersionInfoRequestMsg) (*VersionInfoResponseMsg, error)

	/* Perform ad hoc data extracts */
	Extract(request *ExtractRequestMsg) (*ExtractResponseMsg, error)

	/* Get Current System Status */
	GetSystemStatus(request *SystemStatusRequestMsg) (*SystemStatusResponseMsg, error)
}

type soap struct {
	client *gosoap.Client
}

func NewSoap(client *gosoap.Client) Soap {
	return &soap{
		client: client,
	}
}

func (service *soap) Create(request *CreateRequest) (*CreateResponse, error) {
	response := new(CreateResponse)
	err := service.client.Call("Create", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) Retrieve(request *RetrieveRequestMsg) (*RetrieveResponseMsg, error) {
	response := new(RetrieveResponseMsg)
	err := service.client.Call("Retrieve", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// Recipient Functions
func (service *soap) RetrieveRecipients(request *RetrieveRequestMsg) (*RetrieveRecipientResponseMsg, error) {
	response := new(RetrieveRecipientResponseMsg)
	err := service.client.Call("Retrieve", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) RetrieveRecipient(request *RetrieveRequestMsg) (*RetrieveRecipientResponseMsg, error) {
	response := new(RetrieveRecipientResponseMsg)
	err := service.client.Call("Retrieve", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DataExtension Functions
func (service *soap) RetrieveDataExtensions(request *RetrieveRequestMsg) (*RetrieveDEResponseMsg, error) {
	response := new(RetrieveDEResponseMsg)
	err := service.client.Call("Retrieve", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// func (service *soap) RetrieveDataExtensionByCustomerKey(request *RetrieveRequestMsg, dataExtensionCustomerKey types.DataExtensionRequest) (*RetrieveDEResponseMsg, error) {
func (service *soap) RetrieveDataExtensionByCustomerKey(request *RetrieveRequestMsg) (*RetrieveDEResponseMsg, error) {
	response := new(RetrieveDEResponseMsg)
	err := service.client.Call("Retrieve", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) RetrieveDataExtensionFields(request *RetrieveRequestMsg) (*RetrieveDEFieldResponseMsg, error) {
	response := new(RetrieveDEFieldResponseMsg)
	err := service.client.Call("Retrieve", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// func (service *soap) RetrieveDataExtensionFieldsByDataExtensionCustomerKey(request *RetrieveRequestMsg, dataExtensionCustomerKey types.DataExtensionRequest) (*RetrieveDEFieldResponseMsg, error) {
func (service *soap) RetrieveDataExtensionFieldsByDataExtensionCustomerKey(request *RetrieveRequestMsg) (*RetrieveDEFieldResponseMsg, error) {
	response := new(RetrieveDEFieldResponseMsg)
	err := service.client.Call("Retrieve", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) Update(request *UpdateRequest) (*UpdateResponse, error) {
	response := new(UpdateResponse)
	err := service.client.Call("Update", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

// DataExtensionObject Functions
func (service *soap) RetrieveDataExtensionObject(request *RetrieveRequestMsg) (*RetrieveDataExtensionObjectResponseMsg, error) {
	response := new(RetrieveDataExtensionObjectResponseMsg)
	err := service.client.Call("Retrieve", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) Delete(request *DeleteRequest) (*DeleteResponse, error) {
	response := new(DeleteResponse)
	err := service.client.Call("Delete", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) Query(request *QueryRequestMsg) (*QueryResponseMsg, error) {
	response := new(QueryResponseMsg)
	err := service.client.Call("Query", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) Describe(request *DefinitionRequestMsg) (*DefinitionResponseMsg, error) {
	response := new(DefinitionResponseMsg)
	err := service.client.Call("Describe", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) Execute(request *ExecuteRequestMsg) (*ExecuteResponseMsg, error) {
	response := new(ExecuteResponseMsg)
	err := service.client.Call("Execute", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) Perform(request *PerformRequestMsg) (*PerformResponseMsg, error) {
	response := new(PerformResponseMsg)
	err := service.client.Call("Perform", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) Configure(request *ConfigureRequestMsg) (*ConfigureResponseMsg, error) {
	response := new(ConfigureResponseMsg)
	err := service.client.Call("Configure", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) Schedule(request *ScheduleRequestMsg) (*ScheduleResponseMsg, error) {
	response := new(ScheduleResponseMsg)
	err := service.client.Call("Schedule", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) VersionInfo(request *VersionInfoRequestMsg) (*VersionInfoResponseMsg, error) {
	response := new(VersionInfoResponseMsg)
	err := service.client.Call("VersionInfo", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) Extract(request *ExtractRequestMsg) (*ExtractResponseMsg, error) {
	response := new(ExtractResponseMsg)
	err := service.client.Call("Extract", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *soap) GetSystemStatus(request *SystemStatusRequestMsg) (*SystemStatusResponseMsg, error) {
	response := new(SystemStatusResponseMsg)
	err := service.client.Call("GetSystemStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
