# Migrating from version 2.x.x to 3.0.0

## Dependency

The SDK has moved to a new GitHub repository. You need to replace your existing dependency with the new one. For instance:

```bash
go get github.com/Worldline-Global-Collect/connect-sdk-go
```

## Imports

All packages have been renamed, and some  types, functions and constants have moved to different packages. Each API version now has its own package structure that contains all types and functions specific for that version, including structs like `APIError`, errors and webhooks types and functions.

You need to change your imports as follows:

| Previous package                                                        | Type / function / constant    | New package                                                                   | Notes |
|-------------------------------------------------------------------------|-------------------------------|-------------------------------------------------------------------------------|-------|
| github.com/Ingenico-ePayments/connect-sdk-go                            | CallContext                   | github.com/Worldline-Global-Collect/connect-sdk-go/communicator               |
| github.com/Ingenico-ePayments/connect-sdk-go                            | All other types and functions | github.com/Worldline-Global-Collect/connect-sdk-go                  |
| github.com/Ingenico-ePayments/connect-sdk-go/communicator               | Authenticator                 | github.com/Worldline-Global-Collect/connect-sdk-go/authentication             |
| github.com/Ingenico-ePayments/connect-sdk-go/communicator               | Marshaller                    | github.com/Worldline-Global-Collect/connect-sdk-go/json                       |
| github.com/Ingenico-ePayments/connect-sdk-go/communicator               | MultipartFormDataRequest      | github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication |
| github.com/Ingenico-ePayments/connect-sdk-go/communicator               | ParamRequest                  | github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication |
| github.com/Ingenico-ePayments/connect-sdk-go/communicator               | RequestParam                  | github.com/Worldline-Global-Collect/connect-sdk-go/communicator/communication |
| github.com/Ingenico-ePayments/connect-sdk-go/communicator/communication | UploadableFile                | github.com/Worldline-Global-Collect/connect-sdk-go/domain                     |
| github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl                | AuthorizationType             | github.com/Worldline-Global-Collect/connect-sdk-go/authentication             |
| github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl                | DefaultAuthenticator          | github.com/Worldline-Global-Collect/connect-sdk-go/authentication/v1hmac      |
| github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl                | DefaultConnection             | github.com/Worldline-Global-Collect/connect-sdk-go/communication              |
| github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl                | DefaultMarshaller             | github.com/Worldline-Global-Collect/connect-sdk-go/json                       |
| github.com/Ingenico-ePayments/connect-sdk-go/defaultimpl                | V1HMAC                        | github.com/Worldline-Global-Collect/connect-sdk-go/configuration              |
| github.com/Ingenico-ePayments/connect-sdk-go/domain/metadata            | ShoppingCartExtension         | github.com/Worldline-Global-Collect/connect-sdk-go/domain                     |
| github.com/Ingenico-ePayments/connect-sdk-go/domain/*                   | All other domain structs      | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/domain               | All domain strucs for version 1 of the API are now in the same package |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | APIError                      | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors               |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | AuthorizationError            | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors               |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | CommunicationError            | github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors        |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | CreateAPIError                | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors               |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | DeclinedPaymentError          | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors               |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | DeclinedPayoutError           | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors               |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | DeclinedRefundError           | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors               |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | GlobalCollectError            | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors               |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | IdempotenceError              | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors               |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | NotFoundError                 | github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors        |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | ReferenceError                | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors               |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | ResponseError                 | github.com/Worldline-Global-Collect/connect-sdk-go/communicator/errors        |
| github.com/Ingenico-ePayments/connect-sdk-go/errors                     | ValidateError                 | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/errors               |
| github.com/Ingenico-ePayments/connect-sdk-go/logging                    | All types and functions       | github.com/Worldline-Global-Collect/connect-sdk-go/logging                    |
| github.com/Ingenico-ePayments/connect-sdk-go/logging/obfuscation        | All types and functions       | github.com/Worldline-Global-Collect/connect-sdk-go/logging/obfuscation        |
| github.com/Ingenico-ePayments/connect-sdk-go/merchant/*                 | All types                     | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/merchant/*           | The same package structure is used |
| github.com/Ingenico-ePayments/connect-sdk-go/webhooks                   | Helper                        | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/webhooks             |
| github.com/Ingenico-ePayments/connect-sdk-go/webhooks                   | HelperBuilder                 | github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/webhooks             |
| github.com/Ingenico-ePayments/connect-sdk-go/webhooks                   | SecretKeyNotAvailableError    | github.com/Worldline-Global-Collect/connect-sdk-go/webhooks/validation        |
| github.com/Ingenico-ePayments/connect-sdk-go/webhooks                   | All other types and functions | github.com/Worldline-Global-Collect/connect-sdk-go/webhooks                   |
| github.com/Ingenico-ePayments/connect-sdk-go/webhooks/validation        | All types and functions       | github.com/Worldline-Global-Collect/connect-sdk-go/webhooks/validation        |

## Domain structs

Several domain structs had the name of their package removed from their name to prevent stuttering. For instance, `payment.FindResponse` was used instead of `payment.FindPaymentsResponse`. Now that all domain structs have moved to the same package (`domain`), not including the old package name leads to name clashes. Therefore, the following domain structs have been renamed to include the name of the package again:

| Old struct name                                       | New struct name                                             |
|-------------------------------------------------------|-------------------------------------------------------------|
| capture.Output                                        | domain.CaptureOutput                                        |
| capture.Response                                      | domain.CaptureResponse                                      |
| capture.Status                                        | domain.CaptureStatus                                        |
| capture.StatusOutput                                  | domain.CaptureStatusOutput                                  |
| dispute.CreateRequest                                 | domain.CreateDisputeRequest                                 |
| dispute.CreationDetail                                | domain.DisputeCreationDetail                                |
| dispute.Output                                        | domain.DisputeOutput                                        |
| dispute.Reference                                     | domain.DisputeReference                                     |
| dispute.Response                                      | domain.DisputeResponse                                      |
| dispute.StatusOutput                                  | domain.DisputeStatusOutput                                  |
| hostedcheckout.CreateRequest                          | domain.CreateHostedCheckoutRequest                          |
| hostedcheckout.CreateResponse                         | domain.CreateHostedCheckoutResponse                         |
| hostedcheckout.GetResponse                            | domain.GetHostedCheckoutResponse                            |
| hostedcheckout.SpecificInput                          | domain.HostedCheckoutSpecificInput                          |
| hostedcheckout.Status                                 | domain.HostedCheckoutStatus                                 |
| hostedmandatemanagement.CreateRequest                 | domain.CreateHostedMandateManagementRequest                 |
| hostedmandatemanagement.CreateResponse                | domain.CreateHostedMandateManagementResponse                |
| hostedmandatemanagement.GetResponse                   | domain.GetHostedMandateManagementResponse                   |
| hostedmandatemanagement.SpecificInput                 | domain.HostedMandateManagementSpecificInput                 |
| hostedmandatemanagement.Status                        | domain.HostedMandateManagementStatus                        |
| payment.AccountOnFile                                 | domain.PaymentAccountOnFile                                 |
| payment.ApprovalResponse                              | domain.PaymentApprovalResponse                              |
| payment.ApproveRequest                                | domain.ApprovePaymentRequest                                |
| payment.CancelResponse                                | domain.CancelPaymentResponse                                |
| payment.CreateRequest                                 | domain.CreatePaymentRequest                                 |
| payment.CreateResponse                                | domain.CreatePaymentResponse                                |
| payment.CreateResult                                  | domain.CreatePaymentResult                                  |
| payment.CreationOutput                                | domain.PaymentCreationOutput                                |
| payment.CreationReferences                            | domain.PaymentCreationReferences                            |
| payment.ErrorResponse                                 | domain.PaymentErrorResponse                                 |
| payment.FindResponse                                  | domain.FindPaymentsResponse                                 |
| payment.Output                                        | domain.PaymentOutput                                        |
| payment.Product3201SpecificOutput                     | domain.PaymentProduct3201SpecificOutput                     |
| payment.Product771SpecificOutput                      | domain.PaymentProduct771SpecificOutput                      |
| payment.Product806SpecificOutput                      | domain.PaymentProduct806SpecificOutput                      |
| payment.Product836SpecificOutput                      | domain.PaymentProduct836SpecificOutput                      |
| payment.Product840CustomerAccount                     | domain.PaymentProduct840CustomerAccount                     |
| payment.Product840SpecificOutput                      | domain.PaymentProduct840SpecificOutput                      |
| payment.Product863ThirdPartyData                      | domain.PaymentProduct863ThirdPartyData                      |
| payment.References                                    | domain.PaymentReferences                                    |
| payment.Response                                      | domain.PaymentResponse                                      |
| payment.Status                                        | domain.PaymentStatus                                        |
| payment.StatusCategory                                | domain.PaymentStatusCategory                                |
| payment.StatusOutput                                  | domain.PaymentStatusOutput                                  |
| payment.TokenizeRequest                               | domain.TokenizePaymentRequest                               |
| payout.ApproveRequest                                 | domain.ApprovePayoutRequest                                 |
| payout.CreateRequest                                  | domain.CreatePayoutRequest                                  |
| payout.Customer                                       | domain.PayoutCustomer                                       |
| payout.Details                                        | domain.PayoutDetails                                        |
| payout.ErrorResponse                                  | domain.PayoutErrorResponse                                  |
| payout.FindResponse                                   | domain.FindPayoutsResponse                                  |
| payout.Merchant                                       | domain.PayoutMerchant                                       |
| payout.Recipient                                      | domain.PayoutRecipient                                      |
| payout.References                                     | domain.PayoutReferences                                     |
| payout.Response                                       | domain.PayoutResponse                                       |
| payout.Result                                         | domain.PayoutResult                                         |
| payout.Status                                         | domain.PayoutStatus                                         |
| refund.ApproveRequest                                 | domain.ApproveRefundRequest                                 |
| refund.Customer                                       | domain.RefundCustomer                                       |
| refund.ErrorResponse                                  | domain.RefundErrorResponse                                  |
| refund.FindResponse                                   | domain.FindRefundsResponse                                  |
| refund.References                                     | domain.RefundReferences                                     |
| refund.Request                                        | domain.RefundRequest                                        |
| refund.Response                                       | domain.RefundResponse                                       |
| refund.Result                                         | domain.RefundResult                                         |
| refund.Status                                         | domain.RefundStatus                                         |
| token.ApproveRequest                                  | domain.ApproveTokenRequest                                  |
| token.Card                                            | domain.TokenCard                                            |
| token.CardData                                        | domain.TokenCardData                                        |
| token.CreateRequest                                   | domain.CreateTokenRequest                                   |
| token.CreateResponse                                  | domain.CreateTokenResponse                                  |
| token.EWallet                                         | domain.TokenEWallet                                         |
| token.EWalletData                                     | domain.TokenEWalletData                                     |
| token.NonSepaDirectDebit                              | domain.TokenNonSepaDirectDebit                              |
| token.NonSepaDirectDebitPaymentProduct705SpecificData | domain.TokenNonSepaDirectDebitPaymentProduct705SpecificData |
| token.NonSepaDirectDebitPaymentProduct730SpecificData | domain.TokenNonSepaDirectDebitPaymentProduct730SpecificData |
| token.Response                                        | domain.TokenResponse                                        |
| token.SepaDirectDebit                                 | domain.TokenSepaDirectDebit                                 |
| token.SepaDirectDebitWithoutCreditor                  | domain.TokenSepaDirectDebitWithoutCreditor                  |
| token.UpdateRequest                                   | domain.UpdateTokenRequest                                   |
| webhooks.Event                                        | domain.WebhooksEvent                                        |

Matching functions that return pointers to these structs have been renamed accordingly.

## API calls

Method `Merchant` of struct `connectsdk.Client` has moved to new struct `apiv1.Client`. Instances of this struct are available through method `V1` of struct `connectsdk.Client`. You need to replace all occurrences of `.Merchant` with `.V1().Merchant` in your code. For instance:

```go
response, err := client.V1().Merchant(merchantId).Services().Testconnection(nil)
```

## API version

Constant `connectsdk.APIVersion` has been removed. You need to replace all occurrences in your code with string literal `"v1"`.

## CallContext

Struct `connectsdk.CallContext` has moved to package `github.com/Worldline-Global-Collect/connect-sdk-go/communicator`, and interface `communication.CallContext` that had the same methods has been removed. Function `connectsdk.NewCallContext` remains as alias for function `communicator.NewCallContext`. It no longer returns an error, since the returned error was always `nil`. You need to replace all occurrences of `connectsdk.CallContext` in your code with `communicator.CallContext` and all occurrences of `communication.CallContext` with `*communicator.CallContext`. You also need to remove the error argument on the left-hand side of all calls to function `connectsdk.NewCallContext` in your code. For instance:

```go
context := connectsdk.NewCallContext(idempotenceKey)
```

## Factory

Function `connectsdk.CreateConfiguration` has been renamed to `CreateV1HMACConfiguration`. You need to replace all occurrences in your code with the new name.

## Session

Struct `communicator.Session` has been integrated into struct `communicator.Communicator`. Because struct `communicator.Session` no longer exists, struct `communicator.SessionBuilder` has been replaced with struct `connectsdk.CommunicatorBuilder`, and functions `connectsdk.CreateSessionBuilder` and `connectsdk.CreateSessionBuilderFromConfiguration` have been replaced with functions `connectsdk.CreateCommunicatorBuilder` and `connectsdk.CreateCommunicatorBuilderFromConfiguration` respectively.

If you used function `connectsdk.CreateSessionBuilder` or `connectsdk.CreateSessionBuilderFromConfiguration` to instantiate struct `communicator.SessionBuilder` you need to use the new `connectsdk.CreateCommunicatorBuilder` or `connectsdk.CreateCommunicatorBuilderFromConfiguration` function instead. For instance:

```go
// error handling omitted for brevity
builder, _ := connectsdk.CreateCommunicatorBuilderFromConfiguration(conf)
communicator, _ := builder.WithConnection(connection).Build()
client, _ := connectsdk.CreateClientFromCommunicator(communicator)
```

If you used function `communicator.NewSession` you can use the newly added `connectsdk.CreateCommunicatorWithDefaultMarshaller` and `connectsdk.CreateClientWithDefaultMarshaller` functions. These have the same parameters that the `communicator.NewSession` function did. For instance:

```go
communicator, err := connectsdk.CreateCommunicatorWithDefaultMarshaller(apiEndpoint, connection, authenticator, metadataProvider)
```

Alternatively, you can call the `communicator.NewCommunicator` function directly:

```go
communicator, err := communicator.NewCommunicator(apiEndpoint, connection, authenticator, metadataProvider, marshaller)
```

## Configuration

### CommunicatorConfiguration

Properties `APIKeyID` and `SecretAPIKey` of struct `configuration.CommunicatorConfiguration` have been replaced with more generic properties `AuthorizationID` and `AuthorizationSecret`. You need to replace all occurrences in your code with the new names. Alternatively you can replace all occurrences with one of the newly added `GetAPIKeyID`, `SetAPIKeyID`, `GetSecretAPIKey` and `SetSecretAPIKey` methods.

### DefaultConfiguration

Function `configuration.DefaultConfiguration` has been replaced with function `configuration.DefaultV1HMACConfiguration`. Instead of `configuration.CommunicatorConfiguration` this new function returns `(*configuration.CommunicatorConfiguration, error)`. You need to replace all occurrences in your code with calls to the new function. If you provide non-empty values for the parameters you can ignore the error. For instance:

```go
conf, _ := configuration.DefaultV1HMACConfiguration(apiKeyID, secretAPIKey, integrator)`
```

### Default API endpoint host

The default API endpoint host has changed from `world.api-ingenico.com` to `api.connect.worldline-solutions.com`. To use the previous value, you need to set the nested `APIEndpoint.Host` property of the `configuration.CommunicatorConfiguration` instance you use. If you use `connectsdk.CreateCommunicator` or `connectsdk.CreateClient` you need to replace these with `connectsdk.CreateCommunicatorFromConfiguration` or `connect.CreateClientFromConfiguration`. For instance:

```go
conf, _ := connectsdk.CreateV1HMACConfiguration(apiKeyID, secretAPIKey, integrator)`
conf.APIEndpoint.Host = "world.api-ingenico.com"
client, err := connectsdk.CreateClientFromConfiguration(conf)
```

## Errors

### GlobalCollectError

Struct `errors.GlobalCollectError` and functions `errors.NewGlobalCollectError` and `errors.NewGlobalCollectErrorVerbose` have been renamed to `PlatformError`, `NewPlatformError` and `NewPlatformErrorVerbose` respectively. You need to replace all occurrences in your code with the new names.

### ValidateError

Struct `errors.ValidateError` and functions `errors.NewValidateError` and `errors.NewValidateErrorVerbose` have been renamed to `ValidationError`, `NewValidationError` and `NewValidationErrorVerbose` respectively. You need to replace all occurrences in your code with the new names.

### Exported error variables

Several exported error variables like `communicator.ErrNoName` have been removed. These errors could easily be prevented by providing the correct arguments to functions and methods, and therefore should not be caught.

## Authentication

### Authenticator

The `CreateSimpleAuthenticationSignature` method of interface `Authenticator` has been renamed to `GetAuthorization`. You need to replace all occurrences in your code with the new name.

### DefaultAuthenticator

Struct `communicator.DefaultAuthenticator` and function `communicator.NewDefaultAuthenticator` have moved to package `github.com/Worldline-Global-Collect/connect-sdk-go/authentication/v1hmac` and renamed to `v1hmac.Authenticator` and `v1hmac.NewAuthenticator` respectively. You need to replace all occurrences in your code with the new name.

The `authType` parameter has been dropped from the `v1hmac.NewAuthenticator` function, as it should always be `configuration.V1HMAC`. You need to remove the first argument for all calls to the function in your code.

## Communication

### BodyHandler and BodyHandlerFunc

Interface `communicator.BodyHandler` has been removed, and function type definition `communicator.BodyHandlerFunc` has been renamed to `BodyHandler`. You need to remove all conversions from `func(headers []communication.Header, bodyReader io.Reader) error` to `communicator.BodyHandlerFunc` in your code. For instance:

```go
err := client.V1().Merchant(merchantId).Files().GetFile(fileId, nil, func(headers []communication.Header, bodyReader io.Reader) error {
	// your code
})
```

If you pass instances of a struct that implemented `communicator.BodyHandler` instead you need to replace each instance with a reference to the instance's `Handle` method. For instance:

```go
err := client.V1().Merchant(merchantId).Files().GetFile(fileId, nil, handler.Handle)
```

### Connection

Methods `SetBodyObfuscator` and `SetHeaderObfuscator` have been added to interface `Connection`. This effectively makes each struct that implements `communication.Connection` also implement `obfuscation.Capable`. You need to implement these methods in all custom implementations.


### ResponseHandler and ResponseHandlerFunc

Interface `communication.ResponseHandler` has been removed, and function type definition `communication.ResponseHandlerFunc` has been renamed to `ResponseHandler`. You need to remove all conversions from `func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error)` to `communication.ResponseHandlerFunc` in your code. For instance:

```go
result, err = connection.Get(uri, requestHeaders, func(statusCode int, headers []communication.Header, reader io.Reader) (interface{}, error) {
	// your code
})
```

If you pass instances of a struct that implemented `communication.ResponseHandler` instead you need to replace each instance with a reference to the instance's `Handle` method. For instance:

```go
result, err = connection.Get(uri, requestHeaders, handler.Handle)
```

## JSON marshalling

### DefaultMarshaller

Struct `communication.DefaultMarshaller` and function `communication.NewDefaultMarshaller` have been replaced with function `json.DefaultMarshaller`. This returns the same instance of an implementation of interface `json.Marshaller` every time it's called. Unlike `NewDefaultMarshaller` it does not return an error. You need to replace all occurrences of `*communication.DefaultMarshaller` in your code with `json.Marshaller`, and all occurrences of `communication.NewDefaultAuthenticator` with `json.DefaultMarshaller`. You also need to remove the error argument on the left-hand side. For instance:

```go
marshaller := json.DefaultMarshaller()
```

## Metadata

### MetaDataProvider and MetaDataProviderBuilder

Struct `communicator.MetaDataProvider`, its `MetaDataHeaders` method and struct `communicator.MetaDataProviderBuilder` used incorrect capitalization. These have been renamed to `MetadataProvider`, `MetadataHeaders` and `MetadataProviderBuilder` respectively. You need to replace all occurrences in your code with the new names.

Function `communicator.NewMetaDataProviderWithIntegrator` has been renamed to `NewMetadataProvider`. You need to replace all occurrences in your code with the new name.

Function `communicator.NewMetaDataProviderWithBuilder` has been removed. You need to replace all occurrences in your code with calls to `builder.Build()`, where `builder` is the argument to the `communicator.NewMetaDataProviderWithBuilder` function.

### Integrator

The integrator is now required. You need to make sure a non-empty integrator is set on any `configuration.CommunicatorConfiguration` instance you create.

## Logging

### CommunicatorLogger

Methods `LogRequestLogMessage` and `LogResponseLogMessage` of interface `logging.CommunicatorLogger` have been removed. The SDK calls the `Log` method instead with the result of calling `String()` on the `logging.RequestLogMessage` and `logging.ResponseLogMessage` instances. You need to make sure that all custom implementations correctly implement the `Log` method.

### DefaultLogCommunicatorLogger

Methods `LogRequestLogMessage` and `LogResponseLogMessage` of struct `logging.DefaultLogCommunicatorLogger` have been removed. Like the SDK, you need to call the `Log` method with the result of calling `String()` on the `logging.RequestLogMessage` and `logging.ResponseLogMessage` instances instead.

### RequestLogMessageBuilder

Function `logging.NewRequestLogMessageBuilder` has been removed, and function `logging.NewRequestLogMessageBuilderWithObfuscators` has been renamed to `NewRequestLogMessageBuilder`. You need to add arguments of types `obfuscation.BodyObfuscator` and `obfuscation.HeaderObfuscator` to all occurrences of `logging.NewRequestLogMessageBuilder` in your code, and replace all occurrences of `logging.NewRequestLogMessageBuilderWithObfuscators` with the new name.

### ResponseLogMessageBuilder

Function `logging.NewResponseLogMessageBuilder` has been removed, and function `logging.NewResponseLogMessageBuilderWithObfuscators` has been renamed to `NewResponseLogMessageBuilder`. You need to add arguments of types `obfuscation.BodyObfuscator` and `obfuscation.HeaderObfuscator` to all occurrences of `logging.NewResponseLogMessageBuilder` in your code, and replace all occurrences of `logging.NewResponseLogMessageBuilderWithObfuscators` with the new name.

### StdOutCommunicatorLogger

Function `logging.StdOutCommunicatorLogger` now returns `logging.CommunicatorLogger` instead of `*logging.DefaultLogCommunicatorLogger`. You need to make sure that all variables and parameters in your code where this function is used are of type `logging.CommunicatorLogger`.

## Webhooks

### Creating webhooks helpers

Functions `webhooks.CreateHelper` and `webhooks.CreateHelperBuilder` have been replaced with methods `NewHelper` and `NewHelperBuilder` of new struct `webhooks.Factory` in package `github.com/Worldline-Global-Collect/connect-sdk-go/apiv1/webhooks`. Instances of this struct are available through function `webhooks.V1`. Unlike `webhooks.CreateHelperBuilder`, the `NewHelperBuilder` method does not return an error, instead delaying any error to when `Build()` is called. You need to replace all occurrences of `webhooks.CreateHelper` and `webhooks.CreateHelperBuilder` in your code with `webhooks.V1().NewHelper` and `webhooks.V1().NewHelperBuilder` respectively. If you used `webhooks.CreateHelperBuilder` you also need to remove the error argument on the left-hand side. For instance: 

```go
helper, err := webhooks.V1().NewHelperBuilder(webhooks.InMemorySecretKeyStore()).WithMarshaller(marshaller).Build()
```

### APIVersionMismatchError

Struct `webhooks.APIVersionMismatchError` has been removed. You need to replace all occurrences in your code with `validation.APIVersionMismatchError`.

### InMemorySecretKeyStore

Struct `webhooks.InMemorySecretKeyStore` and function `webhooks.NewInMemorySecretKeyStore` have been replaced with functions `webhooks.InMemorySecretKeyStore`, `webhooks.StoreInMemorySecretKey`, `webhooks.RemoveInMemorySecretKey` and `webhooks.ClearInMemorySecretKeys`. The first of these functions always returns the same `validation.SecretKeyStore` instance, the others replace the methods of struct `webhooks.InMemorySecretKeyStore`. You need to replace all occurrences of the following in your code as specified, where `keyStore` is an instance of the old `webhooks.InMemorySecretKeyStore` struct:

| Old type / call                    | New type / call                  |
|------------------------------------|----------------------------------|
| *webhooks.InMemorySecretKeyStore   | validation.SecretKeyStore        |
| webhooks.NewInMemorySecretKeyStore | webhooks.InMemorySecretKeyStore  |
| keyStore.StoreSecretKey            | webhooks.StoreInMemorySecretKey  |
| keyStore.RemoveSecretKey           | webhooks.RemoveInMemorySecretKey |
| keyStore.Clear                     | webhooks.ClearInMemorySecretKeys |

### SecretKeyStore

Type definition `webhooks.SecretKeyStore` has been removed. You need to replace all occurrences in your code with `validation.SecretKeyStore`.

### SignatureValidationError

Struct `webhooks.SignatureValidationError` has been removed. You need to replace all occurrences in your code with `validation.SignatureValidationError`.

### WebhooksEvent

Function `domain.NewWebhooksEvent`, which replaces function `webhooks.NewEvent`, no longer returns an error, since the returned error was always `nil`. You need to remove the error argument on the left-hand side.
