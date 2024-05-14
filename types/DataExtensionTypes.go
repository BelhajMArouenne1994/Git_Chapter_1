package api

//"fmt"
//"net/http"

//soap_sfmc "github.com/BelhajMArouenne1994/GIT_CHAPTER_1/soap_sfmc"

//"github.com/gin-gonic/gin"

type DataExtensionRequest struct {
	//ID          string `uri:"id" binding:"required,uuid"`
	CustomerKey                    string `uri:"customerKey"`                    //binding:"required
	DataExtensionFieldCustomerKey string `uri:"dataExtensionFieldsCustomerKey"` //binding:"required
}




type DataExtensionFieldsRequest struct {
	//ID          string `uri:"id" binding:"required,uuid"`
	CustomerKey                    string `uri:"customerKey"`                    //binding:"required
	DataExtensionFieldCustomerKey string `uri:"dataExtensionFieldsCustomerKey"` //binding:"required
}