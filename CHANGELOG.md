
<a name="Core Contracts Go Mod Changelog"></a>
## Core Contracts Module (in Go)
[Github repository](https://github.com/edgexfoundry/go-mod-core-contracts)

## [2.0.0] - 2021-06-30
### Features ✨
- **v2:** Create Constants for configuration's map key ([#7342969](https://github.com/edgexfoundry/go-mod-core-contracts/commits/7342969))
- **notifications:** Create client library for support-notifications ([#626](https://github.com/edgexfoundry/go-mod-core-contracts/issues/626)) ([#ee4e77d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ee4e77d))
- **SMA:** Prepare new route and DTO for SMA v2 redesign ([#038c30b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/038c30b))
- **v2:** Add secret route path and HTTP client ([#2bee6ef](https://github.com/edgexfoundry/go-mod-core-contracts/commits/2bee6ef))
- **v2:** Enumerate OperationRequest Action value ([#88fc4e1](https://github.com/edgexfoundry/go-mod-core-contracts/commits/88fc4e1))
- **SMA:** Add OperationRequest dto for SMA ([#9c2e52c](https://github.com/edgexfoundry/go-mod-core-contracts/commits/9c2e52c))
- **v2:** Fix the wrong Transmission v2 API route ([#aa0fb6e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/aa0fb6e))
- **v2:** Add const and modify field for transmission v2 API ([#cb56ff0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/cb56ff0))
- **SMA:** Add constants used by sma v2 API ([#fcfb2da](https://github.com/edgexfoundry/go-mod-core-contracts/commits/fcfb2da))
- **SMA:** Update MulitMetricsResponse field type ([#e860b34](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e860b34))
- **SMA:** Add MultiMetricsResponse and MultiConfigsResponse ([#6135065](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6135065))
- **v2:** Implement v2 GeneralClient ([#ec81246](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ec81246))
- **command:** Create v2 client library for core-command ([#e46cefb](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e46cefb))
- **command:** Add parameters field to Core Command ([#939edfe](https://github.com/edgexfoundry/go-mod-core-contracts/commits/939edfe))
- **data:** Add factory methods for AddEventRequest, Event and Reading DTOs ([#00861d0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/00861d0))
- **SMA:** Add HealthResponse for SMA GET health API ([#5a23571](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5a23571))
- **v2:** Added ApiVersion to BaseRequest ([#12e7666](https://github.com/edgexfoundry/go-mod-core-contracts/commits/12e7666))
- **v2:** Check the response header to handle json or cbor response ([#f86e2eb](https://github.com/edgexfoundry/go-mod-core-contracts/commits/f86e2eb))
- **meta:** Update DeviceResource.Attributes type ([#6a7d99b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6a7d99b))
- **v2:** Make ApiVersion property required. ([#63f8e74](https://github.com/edgexfoundry/go-mod-core-contracts/commits/63f8e74))
- **v2:** Refactor HTTP sendRequest function ([#8d6bffb](https://github.com/edgexfoundry/go-mod-core-contracts/commits/8d6bffb))
- **meta:** Add DeviceResourceResponse DTO and API route ([#fc0ca70](https://github.com/edgexfoundry/go-mod-core-contracts/commits/fc0ca70))
- **meta:** Implement the re-designed device profile model ([#540](https://github.com/edgexfoundry/go-mod-core-contracts/issues/540)) ([#1c03d9d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/1c03d9d))
- **meta:** Enhance v2 DeviceServiceCallbackClient ([#df388aa](https://github.com/edgexfoundry/go-mod-core-contracts/commits/df388aa))
- **data:** Add encoding method for AddEventRequest ([#d71b5a3](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d71b5a3))
- **meta:** Rename Resource to SourceName in AutoEvent model ([#0ece284](https://github.com/edgexfoundry/go-mod-core-contracts/commits/0ece284))
- **meta:** Add error type used by device service data transformation ([#076855e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/076855e))
- **notifications:** Return 400 when UpdateSubscription with empty categories, labels ([#1521f71](https://github.com/edgexfoundry/go-mod-core-contracts/commits/1521f71))
- **v2:** Create a common Address struct for v2 API ([#525](https://github.com/edgexfoundry/go-mod-core-contracts/issues/525)) ([#eae89da](https://github.com/edgexfoundry/go-mod-core-contracts/commits/eae89da))
- **data:** Add Origin constant for event and reading v2 API ([#6bf0fec](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6bf0fec))
- **data:** Remove created field from Event and Reading ([#9016fba](https://github.com/edgexfoundry/go-mod-core-contracts/commits/9016fba))
- **v2:** Address add json omitempty and emailAddress struct to DTO ([#6d7c8f1](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6d7c8f1))
- **data:** Add Encode method for EventResponse ([#19b4da6](https://github.com/edgexfoundry/go-mod-core-contracts/commits/19b4da6))
- **data:** Add CBOR support in EventClient for binary reading ([#c13f5d0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/c13f5d0))
- **SMA:** Implement v2 SystemManagementClient ([#7d294dd](https://github.com/edgexfoundry/go-mod-core-contracts/commits/7d294dd))
- **v2:** Add contentType field to Address ([#3adadce](https://github.com/edgexfoundry/go-mod-core-contracts/commits/3adadce))
- **scheduler:** Update IntervalAction to use the common Address ([#81d8f1f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/81d8f1f))
- **data:** Update Add Event route to include SourceName ([#778f72f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/778f72f))
- **v2:** Add factory methods for DeviceRequest DTO ([#6c6dc03](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6c6dc03))
- **v2:** Add factory methods for Request DTO ([#5755a25](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5755a25))
- **v2:** Enhance the DTO's Json and Validate annotation ([#19dc603](https://github.com/edgexfoundry/go-mod-core-contracts/commits/19dc603))
- **v2:** Add factory methods for DeviceProfileRequest DTO ([#ad613fa](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ad613fa))
- **v2:** Add factory methods for DeviceServiceRequest DTO ([#6d785b0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6d785b0))
- **v2:** Create transform func to convert Model to UpdateDTO ([#691f079](https://github.com/edgexfoundry/go-mod-core-contracts/commits/691f079))
- **command:** Revise CommandResponse DTO to MultiCoreCommandsResponse ([#2037e55](https://github.com/edgexfoundry/go-mod-core-contracts/commits/2037e55))
- **command:**  Remove commandName constant ([#673266b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/673266b))
- **command:** Revise CommandResponse DTO to MultiCoreCommandsResponse -- change import to /v2/v2 ([#cce4776](https://github.com/edgexfoundry/go-mod-core-contracts/commits/cce4776))
- **command:** Revise CommandResponse DTO to MultiCoreCommandsResponse ([#a966e11](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a966e11))
- **data:** Implement UnmarshalCBOR for AddEventRequest DTO ([#a85587f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a85587f))
- **meta:** Add resource map for cache ([#c58f4e6](https://github.com/edgexfoundry/go-mod-core-contracts/commits/c58f4e6))
- **meta:** Assign ApiVersion for each ConvertModelToDTO func ([#99a3cf8](https://github.com/edgexfoundry/go-mod-core-contracts/commits/99a3cf8))
- **meta:** Implement Device Resource Client ([#4ed049f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/4ed049f))
- **meta:** Implement ProvisionWatcherClient ([#098d65f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/098d65f))
- **meta:** Rename ProfileResource to DeviceCommand for v2 Model and DTO ([#d0d059f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d0d059f))
- **notifications:** Create Subscription DTO and Model ([#6e7f6a6](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6e7f6a6))
- **notifications:** Create Transmission DTO and Model ([#41d5e11](https://github.com/edgexfoundry/go-mod-core-contracts/commits/41d5e11))
- **notifications:** Update Subscription DTO to adopt common Address ([#50867e2](https://github.com/edgexfoundry/go-mod-core-contracts/commits/50867e2))
- **notifications:** Add factory method for Notification DTO ([#b2eac77](https://github.com/edgexfoundry/go-mod-core-contracts/commits/b2eac77))
- **notifications:** Add required const and field for sending service ([#59eaa4a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/59eaa4a))
- **notifications:** Add Notification DTO and Model ([#24aa260](https://github.com/edgexfoundry/go-mod-core-contracts/commits/24aa260))
- **scheduler:** Rename Address field and add interval constant ([#4bf416a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/4bf416a))
- **scheduler:** Create v2 IntervalAction DTO and Model ([#ca1ab36](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ca1ab36))
- **scheduler:** Modify import path for v2 Go Model changes ([#19021e2](https://github.com/edgexfoundry/go-mod-core-contracts/commits/19021e2))
- **scheduler:** Create Interval DTO and Model ([#6a623b8](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6a623b8))
- **scheduler:** Implement Interval and IntervalAction Client ([#cf5ddad](https://github.com/edgexfoundry/go-mod-core-contracts/commits/cf5ddad))
- **v2:** BaseResponse omit empty RequestId and Message ([#f45e548](https://github.com/edgexfoundry/go-mod-core-contracts/commits/f45e548))
- **v2:** Create new ErrKind for Delete API to return 409 ([#5f77921](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5f77921))
- **v2:** Adjust DeviceServiceCommandClient interface by moving baseURL as func param ([#3393ae7](https://github.com/edgexfoundry/go-mod-core-contracts/commits/3393ae7))
- **v2:** Create Mocking Clients for v2 Client Libraries ([#b47bd74](https://github.com/edgexfoundry/go-mod-core-contracts/commits/b47bd74))
- **v2:** Update DeviceServicCommandClient interface to have queryParams ([#76d71ab](https://github.com/edgexfoundry/go-mod-core-contracts/commits/76d71ab))
- **v2:** Add queryParams as part of SetCommand ([#c9200dd](https://github.com/edgexfoundry/go-mod-core-contracts/commits/c9200dd))
- **v2:** Refactor CoreCommand DTO to DeviceCoreCommand ([#d4786e8](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d4786e8))
- **v2:** Put together the constants of the models package ([#7967573](https://github.com/edgexfoundry/go-mod-core-contracts/commits/7967573))
- **v2:** Implement Device Service Command Client ([#b433f68](https://github.com/edgexfoundry/go-mod-core-contracts/commits/b433f68))
- **v2:** Update the API path to /device/name/{name}/{command} ([#ec69e1e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ec69e1e))
### Bug Fixes 🐛
- **v2:** Check empty value when converting Model to DTO ([#6459cb9](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6459cb9))
- **v2:** Add the missing ending ']' for float arrays ([#7e9c1cc](https://github.com/edgexfoundry/go-mod-core-contracts/commits/7e9c1cc))
- **meta:** Modify DS Get cmd client API to use pointer as return value ([#ac08c44](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ac08c44))
- **v2:** Client lib checks whether the http response is empty ([#233364e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/233364e))
- **meta:** Add missing UpdateDeviceService DTO filed and omitempty tag ([#1f8df6a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/1f8df6a))
- **v2:** Remove unnecessary arrow sign in error message ([#ac12a39](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ac12a39))
- **notifications:** Add validation tag to verify subscription with empty Channels ([#6cd5e03](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6cd5e03))
- **v2:** Add Address model unmarshal func ([#609911a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/609911a))
- **v2:** Ensure ApiVersion is set when converting models to DTOs ([#127a7ca](https://github.com/edgexfoundry/go-mod-core-contracts/commits/127a7ca))
- **v2:** Fix float reading encoding to be simple E-Notation ([#8786adf](https://github.com/edgexfoundry/go-mod-core-contracts/commits/8786adf))
- **meta:** Modify the profile validation to verify both YAML and JSON format ([#7dc4617](https://github.com/edgexfoundry/go-mod-core-contracts/commits/7dc4617))
- **v2:** Modify the client lib error handling ([#034c2e8](https://github.com/edgexfoundry/go-mod-core-contracts/commits/034c2e8))
- **v2:** Fix ambiguous import path for v2 module ([#ee360d7](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ee360d7))
- **meta:** Add yml tag for device profile API ([#2cf30fb](https://github.com/edgexfoundry/go-mod-core-contracts/commits/2cf30fb))
- **notification:** Add channel type check to Validate method ([#3e070a1](https://github.com/edgexfoundry/go-mod-core-contracts/commits/3e070a1))
- **notifications:** Remove name constraints on category value ([#a60e686](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a60e686))
- **notifications:** Modify inappropriate variable names ([#99ccf6c](https://github.com/edgexfoundry/go-mod-core-contracts/commits/99ccf6c))
- **notifications:** Rename type `Category` to `NotificationCategory` ([#ad44d54](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ad44d54))
- **notifications:** Create EmailAddress model ([#1290a85](https://github.com/edgexfoundry/go-mod-core-contracts/commits/1290a85))
- **scheduler:** Fix func name typo ([#f473d52](https://github.com/edgexfoundry/go-mod-core-contracts/commits/f473d52))
- **v2:** Add missing validation tag to Subscription DTO ([#bd94282](https://github.com/edgexfoundry/go-mod-core-contracts/commits/bd94282))
- **v2:** Correct UpdateSubscription DTO validation tags ([#dde2e0e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/dde2e0e))
- **v2:** Fix subscription_test - Versionable now that is required ([#10f81b4](https://github.com/edgexfoundry/go-mod-core-contracts/commits/10f81b4))
- **v2:** Correct ProvisionWatcher dto json tag ([#794f195](https://github.com/edgexfoundry/go-mod-core-contracts/commits/794f195))
- **v2:** Add replacement of ResendLimit and ResendInterval to ReplaceSubscriptionModelFieldsWithDTO func ([#5f87525](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5f87525))
### Code Refactoring ♻
- **scheduler:** Remove runOnce from the Interval ([#782597b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/782597b))
- **SMA:** Remove obsolete SMA response ([#89682ee](https://github.com/edgexfoundry/go-mod-core-contracts/commits/89682ee))
- **scheduler:** Rename Interval.Frequency field to Interval ([#94fa9ab](https://github.com/edgexfoundry/go-mod-core-contracts/commits/94fa9ab))
- **data:** Rename AutoEvent.Frequency field to Interval ([#a14145b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a14145b))
- **v2:** Rename edgex-dto-frequency validator ([#4a06e94](https://github.com/edgexfoundry/go-mod-core-contracts/commits/4a06e94))
- **securit:** Removed 	SecuritySecretsSetupServiceKey ([#38ea8fc](https://github.com/edgexfoundry/go-mod-core-contracts/commits/38ea8fc))
- **v2:** Remove edgex-prefix from all service keys ([#078c96f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/078c96f))
    ```
    BREAKING CHANGE:
    Service key names have changed.
    ```
- **v2:** Refactor DTO's Created and Modified field to DBTimestamp ([#101c1da](https://github.com/edgexfoundry/go-mod-core-contracts/commits/101c1da))
- **v2:** Rename the unmarshal function name ([#b4fb2f3](https://github.com/edgexfoundry/go-mod-core-contracts/commits/b4fb2f3))
- **v2:** Remove ApiVersion from update DTOs ([#5ea0556](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5ea0556))
- **v2:** Remove ApiVersion from normal DTOs except EventDTO ([#581db9a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/581db9a))
- **meta:** Rename PropertyValue struct to ResourceProperties ([#aae2b6e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/aae2b6e))
- **v2:** Move all constants to common package ([#d45ecd7](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d45ecd7))
- **data:** Update NewAddEventRequest factory method ([#321d834](https://github.com/edgexfoundry/go-mod-core-contracts/commits/321d834))
- **meta:** Rename Put Command to Set Command ([#66222d0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/66222d0))
- **meta:** Profile add UnmarshalYAML func and rename Type field ([#2ee068f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/2ee068f))
- **meta:** Factory method use full object instead of name ([#8a3ce9d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/8a3ce9d))
- **v2:** Remove encoding constants ([#bfd2809](https://github.com/edgexfoundry/go-mod-core-contracts/commits/bfd2809))

<a name="v0.1.149"></a>
## [v0.1.149] - 2021-01-19
### Features ✨
- Update device service v2 api route ([#eda60ea](https://github.com/edgexfoundry/go-mod-core-contracts/commits/eda60ea))

<a name="v0.1.147"></a>
## [v0.1.147] - 2021-01-19
### Features ✨
- Enhance v2 validation error message ([#07beb41](https://github.com/edgexfoundry/go-mod-core-contracts/commits/07beb41))

<a name="v0.1.146"></a>
## [v0.1.146] - 2021-01-14
### Features ✨
- **meta:** Get profileName & deviceName from req ([#92e4a24](https://github.com/edgexfoundry/go-mod-core-contracts/commits/92e4a24))
- **meta:** Revert to a single AddEventRequest DTO ([#a0c1ddc](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a0c1ddc))
- **meta:** Update v2 API AddEvent path ([#f3d88ad](https://github.com/edgexfoundry/go-mod-core-contracts/commits/f3d88ad))
- **meta:** Complete the EventClient ([#b476394](https://github.com/edgexfoundry/go-mod-core-contracts/commits/b476394))

<a name="v0.1.145"></a>
## [v0.1.145] - 2021-01-13
### Features ✨
- **data:** Update AddEvent to use single AddEventRequest DTO ([#6160179](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6160179))

<a name="v0.1.144"></a>
## [v0.1.144] - 2021-01-12
### Code Refactoring ♻
- Chnage from plural Secrets to Secret and SecretData ([#0325e96](https://github.com/edgexfoundry/go-mod-core-contracts/commits/0325e96))

<a name="v0.1.142"></a>
## [v0.1.142] - 2021-01-08
### Features ✨
- **meta:** Add ReadingClient ([#d87652e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d87652e))

<a name="v0.1.141"></a>
## [v0.1.141] - 2021-01-08
### Features ✨
- **meta:** Add v2 ProvisionWatcher API route ([#f1e1e8b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/f1e1e8b))

<a name="v0.1.140"></a>
## [v0.1.140] - 2021-01-08
### Features ✨
- **data:** Update v2 API AddEvent path ([#843dcd9](https://github.com/edgexfoundry/go-mod-core-contracts/commits/843dcd9))

<a name="v0.1.139"></a>
## [v0.1.139] - 2021-01-06
### Features ✨
- **meta:** Add RFC3986 validation on name fields ([#b61c297](https://github.com/edgexfoundry/go-mod-core-contracts/commits/b61c297))
- **meta:** Create ProvisionWatcher Model and DTO ([#50cbdd6](https://github.com/edgexfoundry/go-mod-core-contracts/commits/50cbdd6))

<a name="v0.1.137"></a>
## [v0.1.137] - 2021-01-04
### Features ✨
- **meta:** Add DeviceServiceClient ([#d0b5ae4](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d0b5ae4))

<a name="v0.1.136"></a>
## [v0.1.136] - 2020-12-30
### Features ✨
- Add SecretsRequest DTO ([#5dcdd52](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5dcdd52))

<a name="v0.1.135"></a>
## [v0.1.135] - 2020-12-29
### Code Refactoring ♻
- Refactor logging client to remove remote & file options ([#5220a5b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5220a5b))

<a name="v0.1.134"></a>
## [v0.1.134] - 2020-12-29
### Features ✨
- Refactor createRequest method ([#ae795ac](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ae795ac))
- Add device service callback client ([#ed4f318](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ed4f318))

<a name="v0.1.133"></a>
## [v0.1.133] - 2020-12-28
### Features ✨
- **meta:** Add v2 DeviceClient ([#d692bb7](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d692bb7))

<a name="v0.1.132"></a>
## [v0.1.132] - 2020-12-28
### Features ✨
- **v2:** Implement custom validation tag for RFC3986 unreserved chars ([#4e5601f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/4e5601f))
- **v2:** Implement custom validation tag for RFC3986 unreserved characters ([#fb2b1e2](https://github.com/edgexfoundry/go-mod-core-contracts/commits/fb2b1e2))

<a name="v0.1.131"></a>
## [v0.1.131] - 2020-12-16
### Features ✨
- Make valueType case insensitive and covert to camelcase internally ([#7940b1d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/7940b1d))
### Code Refactoring ♻
- Extract the valueType checking func as validation Tag ([#8b4aa8e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/8b4aa8e))

<a name="v0.1.130"></a>
## [v0.1.130] - 2020-12-16
### Features ✨
- **data:** Update core-data v2 API path constants ([#7d41c9d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/7d41c9d))

<a name="v0.1.129"></a>
## [v0.1.129] - 2020-12-15
### Features ✨
- **meta:** Implement validation logic for device profile DTO ([#1f27142](https://github.com/edgexfoundry/go-mod-core-contracts/commits/1f27142))
### Code Refactoring ♻
- **meta:** Use range to iterates element for verifying device profile ([#505e46a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/505e46a))

<a name="v0.1.128"></a>
## [v0.1.128] - 2020-12-15
### Features ✨
- **v2 data:** Remove pushed field from Event DTO/Model ([#7a3afce](https://github.com/edgexfoundry/go-mod-core-contracts/commits/7a3afce))

<a name="v0.1.127"></a>
## [v0.1.127] - 2020-12-14
### Features ✨
- Refactor GetRequest func to accept request path and params ([#6d5c326](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6d5c326))
- **meta:** Add v2 client for querying device profile ([#db07628](https://github.com/edgexfoundry/go-mod-core-contracts/commits/db07628))

<a name="v0.1.126"></a>
## [v0.1.126] - 2020-12-14
### Features ✨
- **meta:** Add v2 client for deleting device profile ([#47fe6c2](https://github.com/edgexfoundry/go-mod-core-contracts/commits/47fe6c2))

<a name="v0.1.125"></a>
## [v0.1.125] - 2020-12-13
### Features ✨
- Add formatted alternatives to log functions ([#580458f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/580458f))
- Add support to get current log level ([#2708f24](https://github.com/edgexfoundry/go-mod-core-contracts/commits/2708f24))

<a name="v0.1.124"></a>
## [v0.1.124] - 2020-12-11
### Features ✨
- Use require pkg to verify test result ([#c47e45b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/c47e45b))
- **meta:** Add v2 client for uploading device profile in YAML file ([#1b4ca27](https://github.com/edgexfoundry/go-mod-core-contracts/commits/1b4ca27))
### Code Refactoring ♻
- **meta:** Rename request func and add negative test case ([#98564d9](https://github.com/edgexfoundry/go-mod-core-contracts/commits/98564d9))

<a name="v0.1.123"></a>
## [v0.1.123] - 2020-12-09
### Features ✨
- **meta:** Add v2 client for adding, updating device profiles ([#26f7976](https://github.com/edgexfoundry/go-mod-core-contracts/commits/26f7976))
- **meta:** Add v2 client for adding, updating device profiles ([#73165ec](https://github.com/edgexfoundry/go-mod-core-contracts/commits/73165ec))
### Code Refactoring ♻
- **meta:** Refactor http client helper method ([#0d0ddc5](https://github.com/edgexfoundry/go-mod-core-contracts/commits/0d0ddc5))
- **meta:** v2 Device OperatingState value change ([#f86ad7d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/f86ad7d))

<a name="v0.1.122"></a>
## [v0.1.122] - 2020-12-08
### Code Refactoring ♻
- **meta:** Remove unnecessary comments ([#85d4d6a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/85d4d6a))
- **meta:** Remove OperatingState field in DeviceService v2 model ([#617ea87](https://github.com/edgexfoundry/go-mod-core-contracts/commits/617ea87))

<a name="v0.1.121"></a>
## [v0.1.121] - 2020-12-04
### Features ✨
- **data:** Remove Labels out of v2 Reading DTO/Model ([#8582742](https://github.com/edgexfoundry/go-mod-core-contracts/commits/8582742))

<a name="v0.1.120"></a>
## [v0.1.120] - 2020-12-03
### Features ✨
- **data:** Add profileName to Event DTO and Model ([#3af72c1](https://github.com/edgexfoundry/go-mod-core-contracts/commits/3af72c1))

<a name="v0.1.119"></a>
## [v0.1.119] - 2020-12-01
### Code Refactoring ♻
- **v2:** Remove base64 encoding for float value ([#26970d7](https://github.com/edgexfoundry/go-mod-core-contracts/commits/26970d7))

<a name="v0.1.118"></a>
## [v0.1.118] - 2020-11-30
### Bug Fixes 🐛
- Modify the fields in Reading DTO and Model ([#4aa5c7d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/4aa5c7d))

<a name="v0.1.116"></a>
## [v0.1.116] - 2020-11-30
### Features ✨
- Add a new CountResponse to replace EventCountResponse and ReadingCountResponse ([#80998b0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/80998b0))

<a name="v0.1.115"></a>
## [v0.1.115] - 2020-11-21
### Features ✨
- **clients:** Implement Add method for v2 EventClient ([#8845bd7](https://github.com/edgexfoundry/go-mod-core-contracts/commits/8845bd7))

<a name="v0.1.114"></a>
## [v0.1.114] - 2020-11-19
### Bug Fixes 🐛
- Remove error log message when logging set to STDOUT ([#a11423a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a11423a))

<a name="v0.1.113"></a>
## [v0.1.113] - 2020-11-19
### Features ✨
- **clients:** Add v2 CommonClient ([#50f8fed](https://github.com/edgexfoundry/go-mod-core-contracts/commits/50f8fed))

<a name="v0.1.112"></a>
## [v0.1.112] - 2020-10-27
### Bug Fixes 🐛
- **metadata:** Add validation tag for UpdateDTO ([#38672b8](https://github.com/edgexfoundry/go-mod-core-contracts/commits/38672b8))

<a name="v0.1.111"></a>
## [v0.1.111] - 2020-10-23
### Code Refactoring ♻
- Rename constant to match edgex-go funce ([#5ae68de](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5ae68de))

<a name="v0.1.110"></a>
## [v0.1.110] - 2020-10-23
### Code Refactoring ♻
- **data:** Modify event and reading v2 API route path ([#f996e27](https://github.com/edgexfoundry/go-mod-core-contracts/commits/f996e27))

<a name="v0.1.109"></a>
## [v0.1.109] - 2020-10-20
### Features ✨
- Add ContentTypeXML to clients.constants ([#3050e5e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/3050e5e))

<a name="v0.1.108"></a>
## [v0.1.108] - 2020-10-20
### Code Refactoring ♻
- Remove all ResponseNoMessage funcs in v2 ([#c01bdb9](https://github.com/edgexfoundry/go-mod-core-contracts/commits/c01bdb9))

<a name="v0.1.106"></a>
## [v0.1.106] - 2020-10-20
### Features ✨
- **v2:** Update DTOs for UpdateEventPushed ([#cbe9a46](https://github.com/edgexfoundry/go-mod-core-contracts/commits/cbe9a46))

<a name="v0.1.105"></a>
## [v0.1.105] - 2020-10-19
### Bug Fixes 🐛
- Replace broken link in pull request template ([#e304680](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e304680))

<a name="v0.1.104"></a>
## [v0.1.104] - 2020-10-19
### Features ✨
- Add constant for Redis bootstrap ([#ba4de2a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ba4de2a))

<a name="v0.1.103"></a>
## [v0.1.103] - 2020-10-19
### Features ✨
- **metadata:** Add label constant for redis key ([#0a8b18e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/0a8b18e))

<a name="v0.1.102"></a>
## [v0.1.102] - 2020-10-19
### Features ✨
- **metadata:** Add API route path for device ([#ac01b7b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ac01b7b))

<a name="v0.1.101"></a>
## [v0.1.101] - 2020-10-19
### Features ✨
- Add more constants core-data and metadatai ([#ba17f5c](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ba17f5c))

<a name="v0.1.100"></a>
## [v0.1.100] - 2020-10-14
### Features ✨
- **v2:** Add new constant for comma separator to split labels ([#d6824fd](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d6824fd))

<a name="v0.1.99"></a>
## [v0.1.99] - 2020-10-14
### Features ✨
- **v2:** Add new constants and default value for offset, limit, and labels ([#9368c2f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/9368c2f))

<a name="v0.1.98"></a>
## [v0.1.98] - 2020-10-14
### Features ✨
- **v2:** Add new error kind for indicating requested range not satisfiable ([#23ef20d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/23ef20d))

<a name="v0.1.97"></a>
## [v0.1.97] - 2020-10-12
### Features ✨
- **v2:** Add new Response DTOs to return an array of objects ([#b203258](https://github.com/edgexfoundry/go-mod-core-contracts/commits/b203258))

<a name="v0.1.96"></a>
## [v0.1.96] - 2020-10-12
### Bug Fixes 🐛
- Rename deviceId to deviceName from Event DTO ([#da55c3b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/da55c3b))

<a name="v0.1.95"></a>
## [v0.1.95] - 2020-10-12
### Bug Fixes 🐛
- Replace ID to Id in v2 ([#24f5622](https://github.com/edgexfoundry/go-mod-core-contracts/commits/24f5622))

<a name="v0.1.94"></a>
## [v0.1.94] - 2020-10-09
### Features ✨
- **metadata:** Modified deviceProfile DTO to support PUT API ([#cadcd80](https://github.com/edgexfoundry/go-mod-core-contracts/commits/cadcd80))

<a name="v0.1.93"></a>
## [v0.1.93] - 2020-10-08
### Features ✨
- **metadata:** Add API route path for device service ([#b736a93](https://github.com/edgexfoundry/go-mod-core-contracts/commits/b736a93))

<a name="v0.1.91"></a>
## [v0.1.91] - 2020-10-07
### Features ✨
- **metadata:** Add API route path for device profile ([#bc90a46](https://github.com/edgexfoundry/go-mod-core-contracts/commits/bc90a46))

<a name="v0.1.89"></a>
## [v0.1.89] - 2020-10-01
### Bug Fixes 🐛
- **metadata:** Add Id field to DTO Model transform func ([#8c13419](https://github.com/edgexfoundry/go-mod-core-contracts/commits/8c13419))

<a name="v0.1.88"></a>
## [v0.1.88] - 2020-10-01
### Bug Fixes 🐛
- **notifications:** Add ContentType to client struct ([#f9360e8](https://github.com/edgexfoundry/go-mod-core-contracts/commits/f9360e8))

<a name="v0.1.86"></a>
## [v0.1.86] - 2020-09-30
### Features ✨
- **metadata:** Add func to transform the deviceProfile model to DTO ([#da904d6](https://github.com/edgexfoundry/go-mod-core-contracts/commits/da904d6))

<a name="v0.1.84"></a>
## [v0.1.84] - 2020-09-25
### Features ✨
- Add new error types for device SDK v2 API ([#eea4301](https://github.com/edgexfoundry/go-mod-core-contracts/commits/eea4301))

<a name="v0.1.82"></a>
## [v0.1.82] - 2020-09-22
### Fix
- Error msg should return first non-empty msg ([#b21f5c8](https://github.com/edgexfoundry/go-mod-core-contracts/commits/b21f5c8))

<a name="v0.1.81"></a>
## [v0.1.81] - 2020-09-22
### Features ✨
- **metadata:** Add API route path for metadata v2 API ([#cbead72](https://github.com/edgexfoundry/go-mod-core-contracts/commits/cbead72))

<a name="v0.1.78"></a>
## [v0.1.78] - 2020-09-11
### Features ✨
- New error mechanism for v2 API ([#35b6e46](https://github.com/edgexfoundry/go-mod-core-contracts/commits/35b6e46))
### Code Refactoring ♻
- Use int for statusCode instead of uint16 ([#5c2b418](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5c2b418))

<a name="v0.1.77"></a>
## [v0.1.77] - 2020-09-09
### Bug Fixes 🐛
- RequestId in v2 API can be empty or uuid ([#d723a17](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d723a17))

<a name="v0.1.76"></a>
## [v0.1.76] - 2020-09-09
### Bug Fixes 🐛
- Don't in-line `Metrics` property on MetricsResponse so that it matches swagger ([#267d3ce](https://github.com/edgexfoundry/go-mod-core-contracts/commits/267d3ce))

<a name="v0.1.75"></a>
## [v0.1.75] - 2020-09-04
### Bug Fixes 🐛
- Provided custom XML marshaling of Event ([#64c3076](https://github.com/edgexfoundry/go-mod-core-contracts/commits/64c3076))

<a name="v0.1.74"></a>
## [v0.1.74] - 2020-09-01
### Features ✨
- Add Tags to V1 Event's UnmarshalJSON ([#570feb4](https://github.com/edgexfoundry/go-mod-core-contracts/commits/570feb4))
- Add `Tags` field to V1 Event model and v2 Event DTO and model ([#f295970](https://github.com/edgexfoundry/go-mod-core-contracts/commits/f295970))

<a name="v0.1.72"></a>
## [v0.1.72] - 2020-08-18
### Features ✨
- Add Reading DTO ValueType value validation ([#9dedf25](https://github.com/edgexfoundry/go-mod-core-contracts/commits/9dedf25))

<a name="v0.1.71"></a>
## [v0.1.71] - 2020-08-17
### Features ✨
- **metadata:** Add Versionable to v2 Response DTO ([#7a8de45](https://github.com/edgexfoundry/go-mod-core-contracts/commits/7a8de45))

<a name="v0.1.70"></a>
## [v0.1.70] - 2020-08-14
### Features ✨
- **metadata:** Create metadata updating DTOs ([#f8d5fa6](https://github.com/edgexfoundry/go-mod-core-contracts/commits/f8d5fa6))

<a name="v0.1.68"></a>
## [v0.1.68] - 2020-08-06
### Refactor
- Improve Reading field accessibility ([#59edc6d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/59edc6d))

<a name="v0.1.67"></a>
## [v0.1.67] - 2020-08-05
### Features ✨
- **metadata:** Create metadata DTOs and models for v2 API ([#37282b2](https://github.com/edgexfoundry/go-mod-core-contracts/commits/37282b2))

<a name="v0.1.66"></a>
## [v0.1.66] - 2020-08-04
### Features ✨
- Update common Response DTOs (Add factory method & remove BaseResponse) ([#98665c9](https://github.com/edgexfoundry/go-mod-core-contracts/commits/98665c9))
### Fixed
- [#256](https://github.com/edgexfoundry/go-mod-core-contracts/issues/256) SMA agent operation api 404 not found ([#fac8a9e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/fac8a9e))

<a name="v0.1.65"></a>
## [v0.1.65] - 2020-07-20
### Features ✨
- Add func to convert event and reading from model to DTO ([#2d40e57](https://github.com/edgexfoundry/go-mod-core-contracts/commits/2d40e57))
- Add API version for v2 DTO ([#3c79dc8](https://github.com/edgexfoundry/go-mod-core-contracts/commits/3c79dc8))

<a name="v0.1.64"></a>
## [v0.1.64] - 2020-07-10
### Code Refactoring ♻
- Rename device field to deviceName in v2 CoreData DTO and Model In Event and Reading, there is a device field. According to Core WG meeting on July 9th 2020, we need to rename to deviceName to make it explicit. Fix: [#251](https://github.com/edgexfoundry/go-mod-core-contracts/issues/251) ([#6fb5b16](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6fb5b16))

<a name="v0.1.63"></a>
## [v0.1.63] - 2020-07-08
### Code Refactoring ♻
- Remove Retry UrlClient ([#1df9e1d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/1df9e1d))

<a name="v0.1.60"></a>
## [v0.1.60] - 2020-06-22
### Features ✨
- Normalize reading's valueType letter case Since Go DS and C DS send reading with different letter cases, we should normalize the valueType to make it consistent. ([#e04bdb8](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e04bdb8))

<a name="v0.1.36"></a>
## [v0.1.36] - 2019-12-13
### Bug
- **modules:** Add missing yaml attribue to PropertyValue.MediaType ([#e9879d5](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e9879d5))

<a name="v0.1.32"></a>
## [v0.1.32] - 2019-10-22
### Reverts
- Add Parameters field to models.Operations struct