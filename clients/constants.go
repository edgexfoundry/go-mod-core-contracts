/*******************************************************************************
 * Copyright 2019 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package clients

// Do not assume that if a constant is identified by your IDE as not being used within this module that it is not being
// used at all. Any application wishing to exchange information with the EdgeX core services will utilize this module,
// so constants located here may be used externally.
//
// Miscellaneous constants
const (
	ClientMonitorDefault = 15000              // Defaults the interval at which a given service client will refresh its endpoint from the Registry, if used
	CorrelationHeader    = "X-Correlation-ID" // Sets the key of the Correlation ID HTTP header
)

// Constants related to how services identify themselves in the Service Registry
const (
	CoreCommandServiceKey               = "core-command"
	CoreDataServiceKey                  = "core-data"
	CoreMetaDataServiceKey              = "core-metadata"
	SupportLoggingServiceKey            = "support-logging"
	SupportNotificationsServiceKey      = "support-notifications"
	SystemManagementAgentServiceKey     = "sys-mgmt-agent"
	SupportSchedulerServiceKey          = "support-scheduler"
	SecuritySecretStoreSetupServiceKey  = "security-secretstore-setup"
	SecurityProxySetupServiceKey        = "security-proxy-setup"
	SecurityFileTokenProviderServiceKey = "security-file-token-provider"
	SecurityBootstrapperKey             = "security-bootstrapper"
	SecurityBootstrapperRedisKey        = "security-bootstrapper-redis"
)

// Constants related to the possible content types supported by the APIs
const (
	ContentType     = "Content-Type"
	ContentLength   = "Content-Length"
	ContentTypeCBOR = "application/cbor"
	ContentTypeJSON = "application/json"
	ContentTypeYAML = "application/x-yaml"
	ContentTypeText = "text/plain"
	ContentTypeXML  = "application/xml"
)
