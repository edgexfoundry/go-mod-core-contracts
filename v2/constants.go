//
// Copyright (C) 2020 IOTech Ltd
//
// SPDX-License-Identifier: Apache-2.0

package v2

// Constants related to defined routes in the v2 service APIs
const (
	ApiVersion = "v2"
	ApiBase    = "/api/v2"

	ApiEventRoute              = ApiBase + "/event"
	ApiAllEventRoute           = ApiEventRoute + "/" + All
	ApiEventIdRoute            = ApiEventRoute + "/" + Id + "/{" + Id + "}"
	ApiEventPushRoute          = ApiEventRoute + "/" + Pushed
	ApiEventCountRoute         = ApiEventRoute + "/" + Count
	ApiEventCountByDeviceRoute = ApiEventCountRoute + "/" + Device + "/{" + DeviceName + "}"
	ApiEventByDeviceNameRoute  = ApiEventRoute + "/" + Device + "/" + Name + "/{" + Name + "}"
	ApiEventByTimeRangeRoute   = ApiEventRoute + "/" + Start + "/{" + Start + "}/" + End + "/{" + End + "}"
	ApiEventByAgeRoute         = ApiEventRoute + "/" + Age + "/{" + Age + "}"
	ApiEventScrubRoute         = ApiEventRoute + "/" + Scrub

	ApiReadingRoute             = ApiBase + "/reading"
	ApiAllReadingRoute          = ApiReadingRoute + "/" + All
	ApiReadingCountRoute        = ApiReadingRoute + "/" + Count
	ApiReadingIdRoute           = ApiReadingRoute + "/" + Id + "/{" + Id + "}"
	ApiReadingByDeviceNameRoute = ApiReadingRoute + "/" + Device + "/" + Name + "/{" + Name + "}"
	ApiReadingByTypeRoute       = ApiReadingRoute + "/" + Type + "/{" + Type + "}"
	ApiReadingByTimeRangeRoute  = ApiReadingRoute + "/" + Start + "/{" + Start + "}/" + End + "/{" + End + "}"

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

	ApiDeviceRoute              = ApiBase + "/device"
	ApiAllDeviceRoute           = ApiDeviceRoute + "/" + All
	ApiDeviceIdExistsRoute      = ApiDeviceRoute + "/" + Check + "/" + Id + "/{" + Id + "}"
	ApiDeviceNameExistsRoute    = ApiDeviceRoute + "/" + Check + "/" + Name + "/{" + Name + "}"
	ApiDeviceByIdRoute          = ApiDeviceRoute + "/" + Id + "/{" + Id + "}"
	ApiDeviceByNameRoute        = ApiDeviceRoute + "/" + Name + "/{" + Name + "}"
	ApiDeviceByProfileIdRoute   = ApiDeviceRoute + "/" + Profile + "/" + Id + "/{" + Id + "}"
	ApiDeviceByProfileNameRoute = ApiDeviceRoute + "/" + Profile + "/" + Name + "/{" + Name + "}"
	ApiDeviceByServiceIdRoute   = ApiDeviceRoute + "/" + Service + "/" + Id + "/{" + Id + "}"
	ApiDeviceByServiceNameRoute = ApiDeviceRoute + "/" + Service + "/" + Name + "/{" + Name + "}"

	ApiConfigRoute  = ApiBase + "/config"
	ApiMetricsRoute = ApiBase + "/metrics"
	ApiPingRoute    = ApiBase + "/ping"
	ApiVersionRoute = ApiBase + "/version"
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
	ProfileName  = "profileName"
	ServiceName  = "serviceName"
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
	Offset       = "offset" //query string to specify the number of items to skip before starting to collect the result set.
	Limit        = "limit"  //query string to specify the numbers of items to return
	Labels       = "labels" //query string to specify associated user-defined labels for querying a given object. More than one label may be specified via a comma-delimited list
)

// Constants related to the default value of query strings in the v2 service APIs
const (
	DefaultOffset  = 0
	DefaultLimit   = 20
	CommaSeparator = ","
)
