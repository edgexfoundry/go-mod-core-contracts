//
// Copyright (C) 2020-2021 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2

// Constants related to defined routes in the v2 service APIs
const (
	ApiVersion = "v2"
	ApiBase    = "/api/v2"

	ApiEventRoute                      = ApiBase + "/event"
	ApiEventProfileNameDeviceNameRoute = ApiEventRoute + "/{" + ProfileName + "}" + "/{" + DeviceName + "}"
	ApiAllEventRoute                   = ApiEventRoute + "/" + All
	ApiEventIdRoute                    = ApiEventRoute + "/" + Id + "/{" + Id + "}"
	ApiEventCountRoute                 = ApiEventRoute + "/" + Count
	ApiEventCountByDeviceNameRoute     = ApiEventCountRoute + "/" + Device + "/" + Name + "/{" + Name + "}"
	ApiEventByDeviceNameRoute          = ApiEventRoute + "/" + Device + "/" + Name + "/{" + Name + "}"
	ApiEventByTimeRangeRoute           = ApiEventRoute + "/" + Start + "/{" + Start + "}/" + End + "/{" + End + "}"
	ApiEventByAgeRoute                 = ApiEventRoute + "/" + Age + "/{" + Age + "}"

	ApiReadingRoute                  = ApiBase + "/reading"
	ApiAllReadingRoute               = ApiReadingRoute + "/" + All
	ApiReadingCountRoute             = ApiReadingRoute + "/" + Count
	ApiReadingCountByDeviceNameRoute = ApiReadingCountRoute + "/" + Device + "/" + Name + "/{" + Name + "}"
	ApiReadingByDeviceNameRoute      = ApiReadingRoute + "/" + Device + "/" + Name + "/{" + Name + "}"
	ApiReadingByResourceNameRoute    = ApiReadingRoute + "/" + ResourceName + "/{" + ResourceName + "}"
	ApiReadingByTimeRangeRoute       = ApiReadingRoute + "/" + Start + "/{" + Start + "}/" + End + "/{" + End + "}"

	ApiDeviceProfileRoute                       = ApiBase + "/deviceprofile"
	ApiDeviceProfileUploadFileRoute             = ApiDeviceProfileRoute + "/uploadfile"
	ApiDeviceProfileByNameRoute                 = ApiDeviceProfileRoute + "/" + Name + "/{" + Name + "}"
	ApiDeviceProfileByIdRoute                   = ApiDeviceProfileRoute + "/" + Id + "/{" + Id + "}"
	ApiAllDeviceProfileRoute                    = ApiDeviceProfileRoute + "/" + All
	ApiDeviceProfileByManufacturerRoute         = ApiDeviceProfileRoute + "/" + Manufacturer + "/{" + Manufacturer + "}"
	ApiDeviceProfileByModelRoute                = ApiDeviceProfileRoute + "/" + Model + "/{" + Model + "}"
	ApiDeviceProfileByManufacturerAndModelRoute = ApiDeviceProfileRoute + "/" + Manufacturer + "/{" + Manufacturer + "}" + "/" + Model + "/{" + Model + "}"

	ApiDeviceServiceRoute       = ApiBase + "/deviceservice"
	ApiAllDeviceServiceRoute    = ApiDeviceServiceRoute + "/" + All
	ApiDeviceServiceByNameRoute = ApiDeviceServiceRoute + "/" + Name + "/{" + Name + "}"
	ApiDeviceServiceByIdRoute   = ApiDeviceServiceRoute + "/" + Id + "/{" + Id + "}"

	ApiDeviceRoute                = ApiBase + "/device"
	ApiAllDeviceRoute             = ApiDeviceRoute + "/" + All
	ApiDeviceIdExistsRoute        = ApiDeviceRoute + "/" + Check + "/" + Id + "/{" + Id + "}"
	ApiDeviceNameExistsRoute      = ApiDeviceRoute + "/" + Check + "/" + Name + "/{" + Name + "}"
	ApiDeviceByIdRoute            = ApiDeviceRoute + "/" + Id + "/{" + Id + "}"
	ApiDeviceByNameRoute          = ApiDeviceRoute + "/" + Name + "/{" + Name + "}"
	ApiDeviceByProfileIdRoute     = ApiDeviceRoute + "/" + Profile + "/" + Id + "/{" + Id + "}"
	ApiDeviceByProfileNameRoute   = ApiDeviceRoute + "/" + Profile + "/" + Name + "/{" + Name + "}"
	ApiDeviceByServiceIdRoute     = ApiDeviceRoute + "/" + Service + "/" + Id + "/{" + Id + "}"
	ApiDeviceByServiceNameRoute   = ApiDeviceRoute + "/" + Service + "/" + Name + "/{" + Name + "}"
	ApiDeviceNameCommandNameRoute = ApiDeviceByNameRoute + "/{" + Command + "}"

	ApiProvisionWatcherRoute              = ApiBase + "/provisionwatcher"
	ApiAllProvisionWatcherRoute           = ApiProvisionWatcherRoute + "/" + All
	ApiProvisionWatcherByIdRoute          = ApiProvisionWatcherRoute + "/" + Id + "/{" + Id + "}"
	ApiProvisionWatcherByNameRoute        = ApiProvisionWatcherRoute + "/" + Name + "/{" + Name + "}"
	ApiProvisionWatcherByProfileNameRoute = ApiProvisionWatcherRoute + "/" + Profile + "/" + Name + "/{" + Name + "}"
	ApiProvisionWatcherByServiceNameRoute = ApiProvisionWatcherRoute + "/" + Service + "/" + Name + "/{" + Name + "}"

	ApiSubscriptionRoute           = ApiBase + "/subscription"
	ApiAllSubscriptionRoute        = ApiSubscriptionRoute + "/" + All
	ApiSubscriptionByNameRoute     = ApiSubscriptionRoute + "/" + Name + "/{" + Name + "}"
	ApiSubscriptionByCategoryRoute = ApiSubscriptionRoute + "/" + Category + "/{" + Category + "}"
	ApiSubscriptionByLabelRoute    = ApiSubscriptionRoute + "/" + Label + "/{" + Label + "}"
	ApiSubscriptionByReceiverRoute = ApiSubscriptionRoute + "/" + Receiver + "/{" + Receiver + "}"

	ApiConfigRoute  = ApiBase + "/config"
	ApiMetricsRoute = ApiBase + "/metrics"
	ApiPingRoute    = ApiBase + "/ping"
	ApiVersionRoute = ApiBase + "/version"

	ApiDeviceCallbackRoute      = ApiBase + "/callback/device"
	ApiDeviceCallbackNameRoute  = ApiBase + "/callback/device/name/{name}"
	ApiProfileCallbackRoute     = ApiBase + "/callback/profile"
	ApiProfileCallbackNameRoute = ApiBase + "/callback/profile/name/{name}"
	ApiWatcherCallbackRoute     = ApiBase + "/callback/watcher"
	ApiWatcherCallbackNameRoute = ApiBase + "/callback/watcher/name/{name}"
	ApiServiceCallbackRoute     = ApiBase + "/callback/service"
	ApiDiscoveryRoute           = ApiBase + "/discovery"

	ApiIntervalRoute       = ApiBase + "/interval"
	ApiAllIntervalRoute    = ApiIntervalRoute + "/" + All
	ApiIntervalByNameRoute = ApiIntervalRoute + "/" + Name + "/{" + Name + "}"
)

// Constants related to defined url path names and parameters in the v2 service APIs
const (
	All          = "all"
	Id           = "id"
	Created      = "created"
	Modified     = "modified"
	Pushed       = "pushed"
	Count        = "count"
	Device       = "device"
	DeviceId     = "deviceId"
	DeviceName   = "deviceName"
	Check        = "check"
	Profile      = "profile"
	Service      = "service"
	Command      = "command"
	ProfileName  = "profileName"
	ServiceName  = "serviceName"
	ResourceName = "resourceName"
	Start        = "start"
	End          = "end"
	Age          = "age"
	Scrub        = "scrub"
	Type         = "type"
	Name         = "name"
	Label        = "label"
	Manufacturer = "manufacturer"
	Model        = "model"
	ValueType    = "valueType"
	Category     = "category"
	Receiver     = "receiver"
	Subscription = "subscription"

	Offset      = "offset"         //query string to specify the number of items to skip before starting to collect the result set.
	Limit       = "limit"          //query string to specify the numbers of items to return
	Labels      = "labels"         //query string to specify associated user-defined labels for querying a given object. More than one label may be specified via a comma-delimited list
	PushEvent   = "ds-pushevent"   //query string to specify if an event should be pushed to the EdgeX system
	ReturnEvent = "ds-returnevent" //query string to specify if an event should be returned from device service
)

// Constants related to the default value of query strings in the v2 service APIs
const (
	DefaultOffset  = 0
	DefaultLimit   = 20
	CommaSeparator = ","
	ValueYes       = "yes"
	ValueNo        = "no"
)

// Constants related to Reading ValueTypes
const (
	ValueTypeBool         = "Bool"
	ValueTypeString       = "String"
	ValueTypeUint8        = "Uint8"
	ValueTypeUint16       = "Uint16"
	ValueTypeUint32       = "Uint32"
	ValueTypeUint64       = "Uint64"
	ValueTypeInt8         = "Int8"
	ValueTypeInt16        = "Int16"
	ValueTypeInt32        = "Int32"
	ValueTypeInt64        = "Int64"
	ValueTypeFloat32      = "Float32"
	ValueTypeFloat64      = "Float64"
	ValueTypeBinary       = "Binary"
	ValueTypeBoolArray    = "BoolArray"
	ValueTypeStringArray  = "StringArray"
	ValueTypeUint8Array   = "Uint8Array"
	ValueTypeUint16Array  = "Uint16Array"
	ValueTypeUint32Array  = "Uint32Array"
	ValueTypeUint64Array  = "Uint64Array"
	ValueTypeInt8Array    = "Int8Array"
	ValueTypeInt16Array   = "Int16Array"
	ValueTypeInt32Array   = "Int32Array"
	ValueTypeInt64Array   = "Int64Array"
	ValueTypeFloat32Array = "Float32Array"
	ValueTypeFloat64Array = "Float64Array"
)
