
<a name="Core Contracts Go Mod Changelog"></a>
## Core Contracts Module (in Go)
[Github repository](https://github.com/edgexfoundry/go-mod-core-contracts)

## [v4.0.0] - 2025-03-12

### ✨ Features

- Add DeleteById to interface ([49f8627…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/49f86278feb7be638efc0312c2d890135da1788d))
- Add Kind() method to EdgeX error interface ([9172608…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/9172608da6a200124329a18b080222f6fa030ba2))
- Enhance Core Data Event Retention ([c293476…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/c29347609a41d534e07fb76b6089d04855a8188a))
- Add helper function in event DTOs ([#973](https://github.com/edgexfoundry/go-mod-core-contracts/issues/973)) ([658ac86…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/658ac86b41db2d213e72e8edbe30d482a3587384))
- Use `baseUrlFunc` for AuthClient ([86fdbe0…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/86fdbe0ec44bdaf10cc28749820bbaca6e58f71f))
- Add model/DTOs and HTTP client for security-proxy-auth ([8d9ce4b…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/8d9ce4be013658f22f8bc8113ce53db7fa71e4a2))
- Add the post with headers util function for HTTP clients ([#955](https://github.com/edgexfoundry/go-mod-core-contracts/issues/955)) ([9164bc1…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/9164bc1c0b3796a2feabf3521f5acbca65783573))
- Add new API routes for notifications ([ba55863…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/ba5586313443727d187d16cfec8b446e7628fbe5))
- Add more extendable fields to dtos/models for v4 ([b4b29ec…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/b4b29eca62c6eabc141589f69e0eaa5ee2c7d928))
- Change start/end from int to int64 for arm32 ([2e32708…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/2e32708ab81acdef9ec38d9db761411d2c692769))
- Add ZeroMQ channel and update MQTT channel for Address dto/model ([bbd1843…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/bbd1843f5aa1d127b59d23d8e517cff9d4345007))
- Add username/password validations ([4336af0…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/4336af07560a4d02ffb6783f33e18b2f96b5afe4))
- Add Unauthorized `ErrKind` ([a628e4c…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/a628e4c8c657a215a65813f31ce8e62adb89b22e))
- Add `onChangeThreshold` in Device AutoEvent ([f399940…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/f39994077e6603461b5c3fee41b176f810ff9913))
- Add CBOR encode and decode for BaseReading ([8f27c7a…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/8f27c7abae0e16c031c8c7e2ba3c71ec97d07d39))
- Allow reading value to be null ([4a7a10a…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/4a7a10a3a61d76941da71588a5da288692112b8f))
- Add new discovery/profilescan APIs to DeviceServiceCommandClient ([9ff41d2…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/9ff41d212d11429aa0370e900638e98d91d449fa))
- Enhance `WithEmptyPayloadAndId` method ([5095a2a…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/5095a2a290111a3eb981f3cade4dbcfd5959b829))
- Enhance support-scheduler service for advanced job scheduling ([8789206…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/878920651f8f33bcd7d3f24de4f2affa7ae60761))
- Add UnmarshalJSON and WithEmptyPayload methods with more tests ([7ffbf86…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/7ffbf863d5ebd0ce6e0bc811e49f93664f804a9f))
- Add PostgreSQL as the new default persistence layer ([5bc0bbc…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/5bc0bbceaf6546befac2c20fe3f7d9c46a1a1e45))
```text

BREAKING CHANGE: Switched default database to PostgreSQL across all services

```
- Add DS profilescan API DTO and client ([45caab6…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/45caab6e793aec63da2839df359b618f97008afe))
- Add profile basic info http client ([3d51ed7…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/3d51ed7616cc08427e4d538aba11c75e352d1bae))
- Add multiple profile basic info DTO ([2133f19…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/2133f19c414c3a8d1ef62f5484038543baf8ca3d))
- Allow empty profileName in Device ([0782ff0…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/0782ff0d388a9319a830f94ebed6529908651f5b))
- Add ObjectArray valueType support ([#899](https://github.com/edgexfoundry/go-mod-core-contracts/issues/899)) ([467e0a5…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/467e0a5de96c7e27d57f985266f9f068ad2a1d2a))
- Add new Add/update with param device client ([6edc03f…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/6edc03f696a9a074c7115e347077c22ca502b44a))
- Add optional "parent" field to Device objects ([#887](https://github.com/edgexfoundry/go-mod-core-contracts/issues/887)) ([132cf7a…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/132cf7a0d98f69aaa4a72c8537e6ff006b27445a))
- Core Keeper implementation ([e3dd702…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/e3dd7029928a0a6ca0433a8ff353d6de7cdd76df))
```text

BREAKING CHANGE: Introduced Core Keeper as a new service for configuration and registry management

```
- Add openziti to core contracts ([#876](https://github.com/edgexfoundry/go-mod-core-contracts/issues/876)) ([22d5e16…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/22d5e16a8980b24fc88809bb5198de467db1a9fe))
- Add injectEdgeXAuth boolean attribute to RESTAddress types ([66ec24c…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/66ec24c3da583be5fb7edd9ccead2b4b95a54cd3))

### ♻ Code Refactoring

- Generalize LoggingClient interface ([618face…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/618faceb2398c61e8d9464175d38fd493e4878ab))
- Refine the command client library ([ebecdaf…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/ebecdaf26f1810fa691e8939b5c5458ba123432a))
```text

BREAKING CHANGE: Refine the command client library

```
- Update go module to v4 ([3722314…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/37223149a24e8e0216580cbcf98112a0dfd072e8))
```text

BREAKING CHANGE: Update go module to v4

```
- Support null value in reading instead of nil reading value ([43f1e99…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/43f1e99823bcc20f145e229eadd37b2c887ffe9c))

### 🐛 Bug Fixes

- Generate proper error message for 'unique' validation tag ([be22b17…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/be22b176104324ac725c745a5800f937af6b910e))
- Ensure `Properties` field is properly initialized and serialized ([9f205f8…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/9f205f8e9ffbb4e7049917ea8ebf0e95133eb005))
- Replace the error message with BaseResponse.Message ([89e02a4…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/89e02a4eecadce9d454acc7ea4305d261c7934c0))
- Allow binary reading value to be empty ([d524d8c…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/d524d8c204f739d8b5b86ef1ba60ec8e5cd14733))
- Fix inconsistent reading DTO JSON marshal and unmarshal ([c399d60…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/c399d60e2593767685d3b618a083c5a5f2cf0014))
- Fix binary reading validation ([1bf9e0c…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/1bf9e0c7b19d5143de3a4bacf98c81ce5b5d77ee))
- Add missing field "Method" to model and DTO transform functions ([2e1836e…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/2e1836ed65a7b8fa2dc8799e1eaf670623591bec))
- Update DeviceClient and KVSClient mocks ([50a29bf…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/50a29bf8c465817118fd07dc014cdf999ceccbd2))
- Check the response content length before invoking json.Unmarshal ([266762f…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/266762f9c8393c8da7aea63bb418216b78b8829a))


### 👷 Build

- Upgrade to go-1.23, Linter1.61.0 ([1276be6…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/1276be6f117a45c8b13e217674147f4711b849c1))


## [v3.1.0] - 2023-11-15

### ✨  Features

- Allow NameFieldEscape configurable ([b0688fa…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/b0688fa74e6493f4ada8b41961a83322533a5983))
- Use URL escape for DTO name field ([fe869c7…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/fe869c73d3d455c6bcd8a80f44612dba226c759b))
- Add new API route constants for Echo ([ba4bcab…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/ba4bcab7d7f1689a3ec3108d74c56bd728437987))


### 🐛 Bug Fixes

- Use comma without a space as separator to parse array value ([2452e70…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/2452e706a7425134e876c397f528274d17037f0e))
- Add todo comment for the new Echo API const ([e1baa79…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/e1baa79db8796fb641862e2e1fc0f29cb6a78ab5))


### 👷 Build

- Upgrade to go 1.21 and linter 1.54.2 ([aae349d…](https://github.com/edgexfoundry/go-mod-core-contracts/commit/aae349de27787b263ac536cb84a85526ba2e88ac))


## [v3.0.0] - 2023-05-31

### Features ✨

- Using url.PathEscape for http client request path ([#6bcf78a](https://github.com/edgexfoundry/go-mod-core-contracts/commit/6bcf78aee2db63f4e9ae3a03b8b24e3626f753db))
  ```text
  BREAKING CHANGE: To consist with the change for MQTT topic path encoding, use PathEscape to encode the API path because the url.QueryEscape encode the empty space to plus sign, and plus sign is invalid for publishing the MQTT message.
  ```
- Update API version to v3 ([#4f0fdd](https://github.com/edgexfoundry/go-mod-core-contracts/commit/4f0fddbe1238861d7d4ba45b7dfce325304839c5))
  ```text
  BREAKING CHANGE: all EdgeX base API version is changed to v3 for REST endpoints
  ```
- Add 'discoveredDevice' for discovered device fields ([#f889069](https://github.com/edgexfoundry/go-mod-core-contracts/commit/f8890694ba2443089107554beff12780e39f030e))
  ```text
  BREAKING CHANGE: add 'DiscoveredDevice' field which includes profileName,adminState,autoEvents and properties for discovered device
  ```
- Update maximum and minimum to float64 data type ([#1af2d52](https://github.com/edgexfoundry/go-mod-core-contracts/commit/1af2d52b2805dd8b31ab8a74952bd92a1b93e9d7))
  ```text
  BREAKING CHANGE: update maximum and minimum data type from string to float64
  ```
- Update ProtocolProperties to have typed values ([#d9af69a](https://github.com/edgexfoundry/go-mod-core-contracts/commit/d9af69a8b2a38abe867dc9f255879174cd6f5be6))
  ```text
  BREAKING CHANGE: update ProtocolProperties type from 'map[string]string' to 'map[string]any'
  ```
- Update data type in ResourceProperties ([#9353c25](https://github.com/edgexfoundry/go-mod-core-contracts/commit/9353c25cfaefc34f37012e8b0a65cde6e012df1b))
  ```text
  BREAKING CHANGE: update mask,shift,base,scale,offset to numeric data type also rename 'Others' to 'Optional' for the consistency with
  message broker settings.
  ```
- Remove Notify out of Device DTO/Model ([#cb18066](https://github.com/edgexfoundry/go-mod-core-contracts/commit/cb18066d493742bdfaf8f31a0eded08ce169db02))
  ```text
  Since core-metadata will publish system event message when device is created or updated, notify is no longer needed and should be removed.
  ```
- Add interface for injecting authentication data to outgoing requests ([#dc21b6](https://github.com/edgexfoundry/go-mod-core-contracts/commit/dc21b6ba2c81931ffc3ae0ec228443a46b7c8900))
  ```text
  BREAKING CHANGE: This change adds an AuthenticationInjector interface to all http clients. If non-nil, the interface is invoked to add authentication data to the outgoing HTTP request.
  
  While the change does affect the public interface of the HTTP clients, actual requests are only mutated by the authentication injector, which is provided by the consumer.
  ```
- Remove LastConnected and LastReported from Device/DeviceService ([#fbdba8c3e](https://github.com/edgexfoundry/go-mod-core-contracts/commit/fbdba8c3ee7a5841e5299dff2aa29aa4bea36a71))
  ```text
  BREAKING CHANGE: LastConnected and LastReported removed from device and device service DTOs and models
  ```
- Add device service name in Add Event API endpoint ([#61bd91](https://github.com/edgexfoundry/go-mod-core-contracts/commit/61bd91ab058f1c855ec180b3d64afd937c6ababb))
  ```text
  BREAKING CHANGE: Update EventClient 'Add' method signature
  ```
- Remove SimpleReading value required validation for empty string ([#0e331f](https://github.com/edgexfoundry/go-mod-core-contracts/commit/0e331f41e0d4903de81972c1cae15627e716c4a2))
- Allow Device Profile to be empty string in ProvisionWatcher ([#d0b069d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d0b069d))
- Add URLEncode helper function ([#cb6c3dd](https://github.com/edgexfoundry/go-mod-core-contracts/commits/cb6c3dd))
- Add YAML tag for Device and ProvisionWatcher ([#19c4f32](https://github.com/edgexfoundry/go-mod-core-contracts/commits/19c4f32))
- Add 'ds-regexcmd' query parameter ([#6582a37](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6582a37))
- Add ability to authenticate outbound IntervalActions ([#b2c2176](https://github.com/edgexfoundry/go-mod-core-contracts/commits/b2c2176))
- Use URL escape for device command name and resource name ([#1450102](https://github.com/edgexfoundry/go-mod-core-contracts/commits/1450102))
- Add 'Others' field in ResourceProperties ([#6eb51bb](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6eb51bb))
- Add new Properties field to ProvisionWatcher DTO and Model ([#4695302](https://github.com/edgexfoundry/go-mod-core-contracts/commits/4695302))
- Add new Properties field to Device DTO and Model ([#66fa623](https://github.com/edgexfoundry/go-mod-core-contracts/commits/66fa623))
- Added ValidateDeviceSubscribeTopic ([#7d63dd3](https://github.com/edgexfoundry/go-mod-core-contracts/commits/7d63dd3))
- Add UnmarshalObjectValue helper function to Reading DTO ([#baf7df4](https://github.com/edgexfoundry/go-mod-core-contracts/commits/baf7df4))
- Add topic constants that will replace configured topics ([#f8aa83a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/f8aa83a))
- Add service key for new core-common-config-bootstrapper service ([#28efaad](https://github.com/edgexfoundry/go-mod-core-contracts/commits/28efaad))
- Implement xml Marshaller interface for event/reading tags ([#809b610](https://github.com/edgexfoundry/go-mod-core-contracts/commits/809b610))

### Code Refactoring ♻

- Modify the numeric data type in Value Properties to pointer ([#dc7210](https://github.com/edgexfoundry/go-mod-core-contracts/commit/dc72102d704193154e31b181131103b75868c7cc))
  ```text
  BREAKING CHANGE:
  - update mask,shift,base,scale,offset to pointer
  - update maximum and minimum data type from string to float64 pointer
  ```
- Change Config stem constants to flatten config layout ([#f0cd4eda](https://github.com/edgexfoundry/go-mod-core-contracts/commit/f0cd4eda7020d5e86576b480fb0804a1016453ed))
  ```text
  BREAKING CHANGE: Location of service configuartion in Config Provider (Consul) has changed
  ```
- Refactor Secret DTO to rename Path property to SecretName ([#36dc869](https://github.com/edgexfoundry/go-mod-core-contracts/commits/36dc869))
  ```text
    BREAKING CHANGE: Path property renamed to SecretName
  ```
- Remove SystemManagementClient ([#4228851](https://github.com/edgexfoundry/go-mod-core-contracts/commit/42288514cea911c59a0e58050476491dcc7c96d8))
  ```text
  BREAKING CHANGE: SystemManagementClient no longer available
  ```
- Remove old common Metrics DTO and route constant ([#3eccfbd9](https://github.com/edgexfoundry/go-mod-core-contracts/commit/3eccfbd97ef867018de46b5e3b273bb9958dbcdc))
  ```text
  BREAKING CHANGE: /metrics endpoint no longer available for any service
  ```
- Use true/false for command parameters to be more consistent ([#a2a49d7b](https://github.com/edgexfoundry/go-mod-core-contracts/commit/a2a49d7ba2e55cefcdc8c66b99d348e67c20079e))
  ```text
  BREAKING CHANGE: ds-pushevent and ds-returnevent to use true/false instead of yes/no
  ```
- Update module to v3 ([#60a23e1e](https://github.com/edgexfoundry/go-mod-core-contracts/commit/60a23e1e32adaf99c37a33ef7dcfe888a31e5e2f))
  ```text
  BREAKING CHANGE: Import paths will need to change to v3
  ```
- Add new system event type constants and rename action constants ([#d129ffd](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d129ffd))
- Remove obsolete swagger doc reference ([#73cc6cf](https://github.com/edgexfoundry/go-mod-core-contracts/commits/73cc6cf))
- Rename dummy_validator.go to fake_validator.go ([#5a07d2b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5a07d2b))
- Changed config version from 3.0 to v3 ([#78cc929](https://github.com/edgexfoundry/go-mod-core-contracts/commits/78cc929))

### Build 👷

- Update to Go 1.20 and linter v1.51.2 ([#46cc121](https://github.com/edgexfoundry/go-mod-core-contracts/commits/46cc121))

## [v2.3.0] - 2022-11-09

### Features ✨

- Reduce the character restriction for name fields ([#49dcdfb](https://github.com/edgexfoundry/go-mod-core-contracts/commits/49dcdfb))
- Add UoM API route and UoM response DTO ([#aa36825](https://github.com/edgexfoundry/go-mod-core-contracts/commits/aa36825))
- Add new SystemEvent DTO ([#c11c1bf](https://github.com/edgexfoundry/go-mod-core-contracts/commits/c11c1bf))
- Add SimpleReading value validation to prevent type mismatch ([#3cd3c74](https://github.com/edgexfoundry/go-mod-core-contracts/commits/3cd3c74))

### Bug Fixes 🐛

- Keep the existing baseUrl path in helper function ([#abc04af](https://github.com/edgexfoundry/go-mod-core-contracts/commits/abc04af))
- Remove DeviceProfileBasicInfo fieldName from validation error messages ([#8656825](https://github.com/edgexfoundry/go-mod-core-contracts/commits/8656825))

### Code Refactoring ♻

- Adding config stem constants ([#33d1416](https://github.com/edgexfoundry/go-mod-core-contracts/commits/33d1416))
- Improve string concatenate logic for Metric DTO ToLineProtocol() ([#c7c3660](https://github.com/edgexfoundry/go-mod-core-contracts/commits/c7c3660))

### Documentation 📖

- Add deprecated annotation to lastConnected and lastReported ([#ab92af2](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ab92af2))

### Build 👷

- Upgrade to Go 1.18 ([#de188ed](https://github.com/edgexfoundry/go-mod-core-contracts/commits/de188ed))
- Add go build tag to exclude DTO validator ([#af18e4a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/af18e4a))

## [v2.2.0] - 2022-05-11

### Features ✨

- add new method for CommandClient ([#a13ffb3](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a13ffb3))
- Add new device profile API clients ([#3d2ca3b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/3d2ca3b))
- Add Units property to reading DTO and model ([#637851d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/637851d))
- Add Metric DTO to hold data for a single metric ([#6f432a9](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6f432a9))
- Add DTOs for device profile changes ([#ee12520](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ee12520))
- Remove DeviceProfile deviceResources validate gt=0 ([#fd99e17](https://github.com/edgexfoundry/go-mod-core-contracts/commits/fd99e17))
- Add new routes for device profile changes ([#e5aff1a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e5aff1a))
- extend DeviceServiceCallbackClient to support validation API ([#20c4465](https://github.com/edgexfoundry/go-mod-core-contracts/commits/20c4465))
- Add ServiceName to Common DTOs ([#043134c](https://github.com/edgexfoundry/go-mod-core-contracts/commits/043134c))
- update device validation route ([#a30116e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a30116e))
- add validate API route for device service ([#e6de2b0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e6de2b0))
- **security:** Add a new constant for security-spiffe-token-provider ([#306938c](https://github.com/edgexfoundry/go-mod-core-contracts/commits/306938c))

### Test

- Generate mock CommandClient for unit test ([#9b8e31f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/9b8e31f))

### Bug Fixes 🐛

- Remove redundant URL escaping from Client lib ([#c67fdec](https://github.com/edgexfoundry/go-mod-core-contracts/commits/c67fdec))
- Remove JSON omitempty from SimpleReading value ([#d83bf3a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d83bf3a))
- Update device resource/command DTO ([#59aa69d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/59aa69d))
- Add yaml:inline to DeviceProfile DTO DeviceProfileBasicInfo ([#4d80223](https://github.com/edgexfoundry/go-mod-core-contracts/commits/4d80223))
- Update the ReadWrite validation to include WR ([#708](https://github.com/edgexfoundry/go-mod-core-contracts/issues/708)) ([#b3917be](https://github.com/edgexfoundry/go-mod-core-contracts/commits/b3917be))
- update network connectivity problem error kind ([#a52b776](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a52b776))
- Adding the missed PATCH method in REST Address DTO ([#a113ab0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a113ab0))

### Build 👷

- **security:** Enable gosec and default linter set ([#6ac863c](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6ac863c))

## [v2.1.0] - 2021-11-17

### Features ✨

- Add Object Value type in Reading ([#388af6c](https://github.com/edgexfoundry/go-mod-core-contracts/commits/388af6c))
- Add Client API to support Object Value type in Set Command ([#676](https://github.com/edgexfoundry/go-mod-core-contracts/issues/676)) ([#762fd04](https://github.com/edgexfoundry/go-mod-core-contracts/commits/762fd04))
- Add Reading API route constant and client ([#635](https://github.com/edgexfoundry/go-mod-core-contracts/issues/635)) ([#62d0d23](https://github.com/edgexfoundry/go-mod-core-contracts/commits/62d0d23))
- Update routes and ReadingClient for new Reading APIs ([#dcbf024](https://github.com/edgexfoundry/go-mod-core-contracts/commits/dcbf024))
- Remove unclear HTTP status code ([#646](https://github.com/edgexfoundry/go-mod-core-contracts/issues/646)) ([#5e91c92](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5e91c92))
- Add omitempty tag to Reading DTO ([#630bcf1](https://github.com/edgexfoundry/go-mod-core-contracts/commits/630bcf1))
- Update the api version inside all godoc from 2.x to 2.1.0 ([#99ac5f5](https://github.com/edgexfoundry/go-mod-core-contracts/commits/99ac5f5))
- **command:** Add totalCount field into MultiDeviceCoreCommandsResponse DTO ([#eaa77a0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/eaa77a0))
- **data:** Add totalCount field into MultiReadingsResponse DTO ([#94063c0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/94063c0))
- **data:** Use generic interface in the Event Tagging value ([#ad694db](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ad694db))
- **data:** Add totalCount field into MultiEventsResponse DTO ([#e706228](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e706228))
- **data:** Add new core-data reading API route and update ReadingClient ([#2d3bd2a](https://github.com/edgexfoundry/go-mod-core-contracts/commits/2d3bd2a))
- **metadata:** Add totalCount field into core-metadata multi-instance response DTO ([#af86f72](https://github.com/edgexfoundry/go-mod-core-contracts/commits/af86f72))
- **notification:** Add totalCount field into multi-instance response DTOs ([#a61439c](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a61439c))
- **notifications:** Add new notification API route and update TransmissionClient ([#e205b66](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e205b66))
- **scheduler:** Add totalCount field into multi-instance response DTOs ([#e8f11e0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e8f11e0))

### Bug Fixes 🐛

- Add missing DBTimestamp for Model To DTO conversion ([#c361e36](https://github.com/edgexfoundry/go-mod-core-contracts/commits/c361e36))
- Update DTO accept empty Id if the name is provided ([#35f1535](https://github.com/edgexfoundry/go-mod-core-contracts/commits/35f1535))
- Fix error message typo ([#a2d58b6](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a2d58b6))
- **data:** Add reading id mapping during conversion ([#fcb12ca](https://github.com/edgexfoundry/go-mod-core-contracts/commits/fcb12ca))

## [v2.0.0] - 2021-06-30
### General
- **v2:** Implemented V2 DTOs, Model objects and Clients.
- **v1:** Removed v1 APIs and request handling code [e59505e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e59505e)
### Features ✨
- **v2:** Create Constants for configuration's map key ([#7342969](https://github.com/edgexfoundry/go-mod-core-contracts/commits/7342969))
- **notifications:** Create client library for support-notifications ([#626](https://github.com/edgexfoundry/go-mod-core-contracts/issues/626)) ([#ee4e77d](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ee4e77d))
- **SMA:** Prepare new route and DTO for SMA v2 redesign ([#038c30b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/038c30b))
- **SMA:** Add MultiMetricsResponse and MultiConfigsResponse ([#6135065](https://github.com/edgexfoundry/go-mod-core-contracts/commits/6135065))
- **v2:** Implement v2 GeneralClient ([#ec81246](https://github.com/edgexfoundry/go-mod-core-contracts/commits/ec81246))
- **command:** Create v2 client library for core-command ([#e46cefb](https://github.com/edgexfoundry/go-mod-core-contracts/commits/e46cefb))
- **command:** Add parameters field to Core Command ([#939edfe](https://github.com/edgexfoundry/go-mod-core-contracts/commits/939edfe))
- **data:** Add factory methods for AddEventRequest, Event and Reading DTOs ([#00861d0](https://github.com/edgexfoundry/go-mod-core-contracts/commits/00861d0))
- **SMA:** Add HealthResponse for SMA GET health API ([#5a23571](https://github.com/edgexfoundry/go-mod-core-contracts/commits/5a23571))
- **v2:** Added ApiVersion to BaseRequest ([#12e7666](https://github.com/edgexfoundry/go-mod-core-contracts/commits/12e7666))
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
- **command:**  Remove commandName constant ([#673266b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/673266b))
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
### Code Refactoring ♻
- **scheduler:** Remove runOnce from the Interval ([#782597b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/782597b))
- **scheduler:** Rename Interval.Frequency field to Interval ([#94fa9ab](https://github.com/edgexfoundry/go-mod-core-contracts/commits/94fa9ab))
- **data:** Rename AutoEvent.Frequency field to Interval ([#a14145b](https://github.com/edgexfoundry/go-mod-core-contracts/commits/a14145b))
- **security:** Removed	SecuritySecretsSetupServiceKey ([#38ea8fc](https://github.com/edgexfoundry/go-mod-core-contracts/commits/38ea8fc))
- **v2:** Remove edgex-prefix from all service keys ([#078c96f](https://github.com/edgexfoundry/go-mod-core-contracts/commits/078c96f))
    ```
    BREAKING CHANGE:
    Service key names have changed.
    ```
- **meta:** Rename PropertyValue struct to ResourceProperties ([#aae2b6e](https://github.com/edgexfoundry/go-mod-core-contracts/commits/aae2b6e))
- **v2:** Move all constants to common package ([#d45ecd7](https://github.com/edgexfoundry/go-mod-core-contracts/commits/d45ecd7))

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