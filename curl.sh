curl -X POST -H 'Soapaction: Retrieve' -H 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsImtpZCI6IjQiLCJ2ZXIiOiIxIiwidHlwIjoiSldUIn0.eyJhY2Nlc3NfdG9rZW4iOiJYblpNek5IZFVCRE91dGRHbTR0bEVvR2giLCJjbGllbnRfaWQiOiJidW0ybDhpcjhpam5qcWtlMXExaDBmbm8iLCJlaWQiOjUxMDAwOTI2Miwic3RhY2tfa2V5IjoiUzUwIiwicGxhdGZvcm1fdmVyc2lvbiI6MiwiY2xpZW50X3R5cGUiOiJTZXJ2ZXJUb1NlcnZlciIsInBpZCI6OTN9.pZx3pn_AaWYCtw0nukJQ1BMaJizQ9ffSQIo0d90OIZI.5SuN547NGL1iWGurSBF10udOeVdM8IrlYB_VNiFbrtASfOhl8Xtc4v6nzALUUOej5y5TLFGoMvShp4dv-OKbBqqbmMwpEOFs8ZdZpSSHEQ83Mk_zCIJ_aknzxEnz7eocC8ZYOLgaedZN5zq6T4zQGZV4N8aSqP4nHC6DNky' -H 'Content-Type: text/xml;charset=utf-8' -d '<RetrieveRequestMsg xmlns="http://exacttarget.com/wsdl/partnerAPI"><RetrieveRequest><ObjectType>Subscriber</ObjectType><Properties>ID</Properties><Properties>EmailAddress</Properties><Retrieves></Retrieves></RetrieveRequest></RetrieveRequestMsg>' 'https://mc166917qx7lk8ldd5cwvlp6ftzq.soap.marketingcloudapis.com/'



curl --location --globoff 'https://{{et_subdomain}}.soap.marketingcloudapis.com/Service.asmx' \
--header 'Content-Type: application/xml' \
--data-raw '<?xml version="1.0" encoding="UTF-8"?>
<s:Envelope xmlns:s="http://www.w3.org/2003/05/soap-envelope" xmlns:a="http://schemas.xmlsoap.org/ws/2004/08/addressing" xmlns:u="http://docs.oasis-open.org/wss/2004/01/oasis-200401-wss-wssecurity-utility-1.0.xsd">
    <s:Header>
        <a:Action s:mustUnderstand="1">Create</a:Action>
        <a:To s:mustUnderstand="1">https://{{et_subdomain}}.soap.marketingcloudapis.com/Service.asmx</a:To>
        <fueloauth xmlns="http://exacttarget.com">eyJhbGciOiJIUzI1NiIsImtpZCI6IjQiLCJ2ZXIiOiIxIiwidHlwIjoiSldUIn0.eyJhY2Nlc3NfdG9rZW4iOiJYblpNek5IZFVCRE91dGRHbTR0bEVvR2giLCJjbGllbnRfaWQiOiJidW0ybDhpcjhpam5qcWtlMXExaDBmbm8iLCJlaWQiOjUxMDAwOTI2Miwic3RhY2tfa2V5IjoiUzUwIiwicGxhdGZvcm1fdmVyc2lvbiI6MiwiY2xpZW50X3R5cGUiOiJTZXJ2ZXJUb1NlcnZlciIsInBpZCI6OTN9.pZx3pn_AaWYCtw0nukJQ1BMaJizQ9ffSQIo0d90OIZI.5SuN547NGL1iWGurSBF10udOeVdM8IrlYB_VNiFbrtASfOhl8Xtc4v6nzALUUOej5y5TLFGoMvShp4dv-OKbBqqbmMwpEOFs8ZdZpSSHEQ83Mk_zCIJ_aknzxEnz7eocC8ZYOLgaedZN5zq6T4zQGZV4N8aSqP4nHC6DNky</fueloauth>
    </s:Header>
    <s:Body xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xmlns:xsd="http://www.w3.org/2001/XMLSchema">
        <CreateRequest xmlns="http://exacttarget.com/wsdl/partnerAPI">
            <Objects xsi:type="Subscriber">
                <ObjectID xsi:nil="true"></ObjectID>
				<CustomerKey>123456</CustomerKey>
                <EmailAddress>acatessf+postman@gmail.com</EmailAddress>
                <Lists>
                    <ID>71720</ID>
                    <ObjectID xsi:nil="true">
                    </ObjectID>
                </Lists>
                <Attributes>
                    <Name>First Name</Name>
                    <Value>Aaron</Value>
                </Attributes>
                <Attributes>
                    <Name>Last Name</Name>
                    <Value>Cates</Value>
                </Attributes>
                <Attributes>
                    <Name>Company</Name>
                    <Value>Salesforce</Value>
                </Attributes>
            </Objects>
        </CreateRequest>
    </s:Body>
</s:Envelope>'