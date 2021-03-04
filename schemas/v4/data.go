// Copyright (c) 2021, Ricard Bejarano. All rights reserved.
// Use of this source code is governed by the BSD 3-Clause
// license, which can be found in the LICENSE file.

package schema

import (
	"time"
)

type Data struct {
	Cost   uint64
	Viewer struct {
		Accounts []struct {
			AccountTag           string
			HttpRequests1dGroups []struct {
				Avg        httpRequestsGroupsAvg
				Dimensions struct {
					Date time.Time
				}
				Sum  httpRequestsGroupsSum
				Uniq httpRequestsGroupsUniq
			}
			HttpRequests1hGroups []struct {
				Avg        httpRequestsGroupsAvg
				Dimensions struct {
					Date     time.Time
					Datetime time.Time
				}
				Sum  httpRequestsGroupsSum
				Uniq httpRequestsGroupsUniq
			}
			HttpRequests1mGroups []struct {
				Avg        httpRequestsGroupsAvg
				Dimensions struct {
					Date                   time.Time
					Datetime               time.Time
					DatetimeDay            time.Time
					DatetimeFifteenMinutes time.Time
					DatetimeFiveMinutes    time.Time
					DatetimeHalfOfHour     time.Time
					DatetimeHour           time.Time
					DatetimeMinute         time.Time
				}
				Sum  httpRequestsGroupsSum
				Uniq httpRequestsGroupsUniq
			}
			Settings struct {
				HttpRequests1dGroups                         settings
				HttpRequests1hGroups                         settings
				HttpRequests1mGroups                         settings
				IpFlows1dGroups                              settings
				IpFlows1hGroups                              settings
				IpFlows1mAttacksGroups                       settings
				IpFlows1mGroups                              settings
				MagicTransitTunnelHealthChecksAdaptiveGroups settings
				RumPageloadEventsAdaptiveGroups              settings
				RumWebVitalsEventsAdaptiveGroups             settings
				VideoBufferEventsAdaptiveGroups              settings
				VideoPlaybackEventsAdaptiveGroups            settings
				VideoQualityEventsAdaptiveGroups             settings
				WorkersInvocationsAdaptive                   settings
				WorkersInvocationsScheduled                  settings
			}
			WorkersInvocationsScheduled []struct {
				CpuTimeUs         uint32
				Cron              string
				Datetime          time.Time
				ScheduledDatetime time.Time
				ScriptName        string
				Status            string
			}
			Zones []zone
		}
		Budget uint64
		Zones  []zone
	}
}

type zone struct {
	FirewallEventsAdaptive []struct {
		Action                      string
		BotScore                    uint8
		BotScoreSrcName             string
		ClientASNDescription        string
		ClientAsn                   string
		ClientCountryName           string
		ClientIP                    string
		ClientIPClass               string
		ClientRefererHost           string
		ClientRefererPath           string
		ClientRefererQuery          string
		ClientRefererScheme         string
		ClientRequestHTTPHost       string
		ClientRequestHTTPMethodName string
		ClientRequestHTTPProtocol   string
		ClientRequestPath           string
		ClientRequestQuery          string
		ClientRequestScheme         string
		Date                        time.Time
		Datetime                    time.Time
		DatetimeFifteenMinutes      time.Time
		DatetimeFiveMinutes         time.Time
		DatetimeHour                time.Time
		DatetimeMinute              time.Time
		EdgeColoName                string
		EdgeResponseStatus          uint16
		Kind                        string
		MatchIndex                  uint16
		Metadata                    []struct {
			Key   string
			Value string
		}
		OriginResponseStatus uint16
		OriginatorRayName    string
		RayName              string
		Ref                  string
		RuleId               string
		SampleInterval       uint32
		Source               string
		UserAgent            string
	}
	FirewallEventsAdaptiveByTimeGroups []struct {
		Avg struct {
			SampleInterval float64
		}
		Count      uint64
		Dimensions struct {
			BotScore               uint8
			BotScoreSrcName        string
			Date                   time.Time
			Datetime               time.Time
			DatetimeFifteenMinutes time.Time
			DatetimeFiveMinutes    time.Time
			DatetimeHour           time.Time
			DatetimeMinute         time.Time
		}
	}
	FirewallEventsAdaptiveGroups []struct {
		Avg struct {
			SampleInterval float64
		}
		Count      uint64
		Dimensions struct {
			Action                      string
			BotScore                    uint8
			BotScoreSrcName             string
			ClientASNDescription        string
			ClientAsn                   string
			ClientCountryName           string
			ClientIP                    string
			ClientIPClass               string
			ClientRefererHost           string
			ClientRefererPath           string
			ClientRefererQuery          string
			ClientRefererScheme         string
			ClientRequestHTTPHost       string
			ClientRequestHTTPMethodName string
			ClientRequestHTTPProtocol   string
			ClientRequestPath           string
			ClientRequestQuery          string
			ClientRequestScheme         string
			Date                        time.Time
			Datetime                    time.Time
			DatetimeFifteenMinutes      time.Time
			DatetimeFiveMinutes         time.Time
			DatetimeHour                time.Time
			DatetimeMinute              time.Time
			EdgeColoName                string
			EdgeResponseStatus          uint16
			Kind                        string
			MatchIndex                  uint16
			OriginResponseStatus        uint16
			OriginatorRayName           string
			RayName                     string
			Ref                         string
			RuleId                      string
			SampleInterval              uint32
			Source                      string
			UserAgent                   string
		}
	}
	HealthCheckEventsAdaptive []struct {
		Datetime              time.Time
		EventId               string
		ExpectedResponseCodes string
		FailureReason         string
		Fqdn                  string
		HealthChanged         uint8
		HealthCheckId         string
		HealthCheckName       string
		HealthStatus          string
		OriginIP              string
		OriginResponseStatus  uint16
		Region                string
		RttMs                 uint64
		SampleInterval        uint32
		Scope                 string
		TcpConnMs             uint32
		TimeToFirstByteMs     uint32
		TlsHandshakeMs        uint32
	}
	HealthCheckEventsAdaptiveGroups []struct {
		Avg struct {
			RttMs             uint64
			SampleInterval    float64
			TcpConnMs         uint32
			TimeToFirstByteMs uint32
			TlsHandshakeMs    uint32
		}
		Count      uint64
		Dimensions struct {
			Datetime               time.Time
			DatetimeFifteenMinutes time.Time
			DatetimeFiveMinutes    time.Time
			DatetimeHour           time.Time
			DatetimeMinute         time.Time
			FailureReason          string
			Fqdn                   string
			HealthChanged          uint8
			HealthCheckId          string
			HealthCheckName        string
			HealthStatus           string
			OriginIP               string
			OriginResponseStatus   uint16
			Region                 string
			RttMs                  uint64
			SampleInterval         uint32
			Scope                  string
			TcpConnMs              uint32
			TimeToFirstByteMs      uint32
			TlsHandshakeMs         uint32
		}
		Sum struct {
			HealthStatus uint64
		}
	}
	HttpRequests1dGroups []struct {
		Avg        httpRequestsGroupsAvg
		Dimensions struct {
			Date time.Time
		}
		Sum  httpRequestsGroupsSum
		Uniq httpRequestsGroupsUniq
	}
	HttpRequests1hGroups []struct {
		Avg        httpRequestsGroupsAvg
		Dimensions struct {
			Date     time.Time
			Datetime time.Time
		}
		Sum  httpRequestsGroupsSum
		Uniq httpRequestsGroupsUniq
	}
	HttpRequests1mGroups []struct {
		Avg        httpRequestsGroupsAvg
		Dimensions struct {
			Date                   time.Time
			Datetime               time.Time
			DatetimeDay            time.Time
			DatetimeFifteenMinutes time.Time
			DatetimeFiveMinutes    time.Time
			DatetimeHalfOfHour     time.Time
			DatetimeHour           time.Time
			DatetimeMinute         time.Time
		}
		Sum  httpRequestsGroupsSum
		Uniq httpRequestsGroupsUniq
	}
	HttpRequestsAdaptiveGroups []struct {
		Avg struct {
			SampleInterval float64
		}
		Count      uint64
		Dimensions struct {
			BotManagementDecision       string
			BotScore                    uint8
			BotScoreBucketBy10          uint8
			BotScoreSrcName             string
			CacheStatus                 string
			ClientASNDescription        string
			ClientAsn                   string
			ClientCountryName           string
			ClientDeviceType            string
			ClientIP                    string
			ClientRefererHost           string
			ClientRequestHTTPHost       string
			ClientRequestHTTPMethodName string
			ClientRequestHTTPProtocol   string
			ClientRequestPath           string
			ClientRequestQuery          string
			ClientRequestReferer        string
			ClientRequestScheme         string
			ClientSSLProtocol           string
			ColoCode                    string
			Date                        time.Time
			Datetime                    time.Time
			DatetimeFifteenMinutes      time.Time
			DatetimeFiveMinutes         time.Time
			DatetimeHour                time.Time
			DatetimeMinute              time.Time
			EdgeResponseContentTypeName string
			EdgeResponseStatus          uint16
			OriginASN                   uint32
			OriginASNDescription        string
			OriginIP                    string
			OriginResponseStatus        uint16
			RequestSource               string
			SampleInterval              uint32
			UpperTierColoName           string
			UserAgent                   string
			UserAgentBrowser            string
			UserAgentOS                 string
		}
		Sum struct {
			EdgeResponseBytes uint64
			Visits            uint64
		}
	}
	ImageResizingRequests1mGroups []struct {
		Count      uint64
		Dimensions struct {
			ContentType            string
			Date                   time.Time
			Datetime               time.Time
			DatetimeFifteenMinutes time.Time
			DatetimeFiveMinutes    time.Time
			DatetimeHour           time.Time
			DatetimeMinute         time.Time
			ResizeError            string
		}
		Sum struct {
			OriginalBytes  uint64
			OriginalPixels uint64
			Requests       uint64
			ResizedBytes   uint64
			ResizedPixels  uint64
		}
	}
	LoadBalancingRequestsAdaptive []struct {
		ColoCode              string
		Datetime              time.Time
		ErrorType             string
		LbName                string
		NumberOriginsSelected uint16
		Origins               []struct {
			Fqdn       string
			Health     uint8
			Ipv4       string
			Ipv6       string
			OriginName string
			Selected   uint8
			Weight     float64
		}
		Pools []struct {
			AvgRttMs           uint64
			HealthCheckEnabled uint8
			Healthy            uint8
			Id                 string
			PoolName           string
		}
		Proxied                         uint8
		Region                          string
		SampleInterval                  uint32
		SelectedOriginIndex             uint16
		SelectedPoolAvgRttMs            uint64
		SelectedPoolHealthChecksEnabled uint8
		SelectedPoolHealthy             uint8
		SelectedPoolId                  string
		SelectedPoolIndex               uint32
		SelectedPoolName                string
		SessionAffinity                 string
		SessionAffinityStatus           string
		SteeringPolicy                  string
	}
	LoadBalancingRequestsAdaptiveGroups []struct {
		Avg struct {
			SampleInterval float64
		}
		Count      uint64
		Dimensions struct {
			ColoCode                        string
			Datetime                        time.Time
			DatetimeFifteenMinutes          time.Time
			DatetimeFiveMinutes             time.Time
			DatetimeHalfOfHour              time.Time
			DatetimeHour                    time.Time
			DatetimeMinute                  time.Time
			ErrorType                       string
			LbName                          string
			MultipleOriginsSelected         uint8
			NumberOriginsSelected           uint16
			Proxied                         uint8
			Region                          string
			SampleInterval                  uint32
			SelectedOriginIndex             uint16
			SelectedOriginName              string
			SelectedPoolAvgRttMs            uint64
			SelectedPoolHealthChecksEnabled uint8
			SelectedPoolHealthy             uint8
			SelectedPoolId                  string
			SelectedPoolIndex               uint32
			SelectedPoolName                string
			SessionAffinityEnabled          string
			SessionAffinityStatus           string
			SteeringPolicy                  string
		}
	}
	Settings struct {
		BrowserInsightsAdaptiveGroups         settings
		BrowserInsightsResourceAdaptiveGroups settings
		FirewallEventsAdaptive                settings
		FirewallEventsAdaptiveByTimeGroups    settings
		FirewallEventsAdaptiveGroups          settings
		HealthCheckEvents                     settings
		HealthCheckEventsAdaptive             settings
		HealthCheckEventsAdaptiveGroups       settings
		HealthCheckEventsGroups               settings
		HttpRequests1dByColoGroups            settings
		HttpRequests1dGroups                  settings
		HttpRequests1hGroups                  settings
		HttpRequests1mByColoGroups            settings
		HttpRequests1mGroups                  settings
		HttpRequestsAdaptiveGroups            settings
		ImageResizingRequests1mGroups         settings
		LoadBalancingRequests                 settings
		LoadBalancingRequestsAdaptive         settings
		LoadBalancingRequestsAdaptiveGroups   settings
		LoadBalancingRequestsGroups           settings
		SynAvgPps1mGroups                     settings
		WebVitalsAdaptiveGroups               settings
	}
	SynAvgPps1mGroups []struct {
		Count      uint64
		Dimensions struct {
			Date                   time.Time
			DatetimeDay            time.Time
			DatetimeFifteenMinutes time.Time
			DatetimeFiveMinutes    time.Time
			DatetimeHour           time.Time
			DatetimeMinute         time.Time
		}
		Sum struct {
			AvgPpsPerMinute uint64
		}
	}
	ZoneTag string
}

type settings struct {
	AvailableFields   []string
	Enabled           bool
	MaxDuration       int64
	MaxNumberOfFields int64
	MaxPageSize       int64
	NotOlderThan      int64
}

// httpRequestsGroupsAvg does not represent the real schema, but it is put into
// its own structure because its contents are the same across multiple parts of
// the schema, thus allowing for code reutilisation.
type httpRequestsGroupsAvg struct {
	Bytes          float64
	SampleInterval float64
}

// httpRequestsGroupsSum does not represent the real schema, but it is put into
// its own structure because its contents are the same across multiple parts of
// the schema, thus allowing for code reutilisation.
type httpRequestsGroupsSum struct {
	BrowserMap []struct {
		PageViews       uint64
		UaBrowserFamily string
	}
	Bytes                uint64
	CachedBytes          uint64
	CachedRequests       uint64
	ClientHTTPVersionMap []struct {
		ClientHTTPProtocol string
		Requests           uint64
	}
	ClientSSLMap []struct {
		ClientSSLProtocol string
		Requests          uint64
	}
	ContentTypeMap []struct {
		Bytes                       uint64
		EdgeResponseContentTypeName string
		Requests                    uint64
	}
	CountryMap []struct {
		Bytes             uint64
		ClientCountryName string
		Requests          uint64
		Threats           uint64
	}
	EncryptedBytes    uint64
	EncryptedRequests uint64
	IpClassMap        []struct {
		IpType   string
		Requests uint64
	}
	PageViews         uint64
	Requests          uint64
	ResponseStatusMap []struct {
		EdgeResponseStatus uint64
		Requests           uint64
	}
	ThreatPathingMap []struct {
		Requests          uint64
		ThreatPathingName string
	}
	Threats uint64
}

// httpRequestsGroupsUniq does not represent the real schema, but it is put into
// its own structure because its contents are the same across multiple parts of
// the schema, thus allowing for code reutilisation.
type httpRequestsGroupsUniq struct {
	Uniques uint64
}
