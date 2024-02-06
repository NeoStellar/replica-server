package telemetri_gonder

import (
	"encoding/json"
	"slices"
)

func PushTelemetryData(data TelemetryData) []TelemetryData {
	var telemetryDatas []TelemetryData
	if redisDatas := redis.Get("telemetryDatas"); len(redisDatas) > 0 {
		json.Unmarshal([]byte(redisDatas), &telemetryDatas)
	}
	idx := slices.IndexFunc(telemetryDatas, func(telemetryData TelemetryData) bool { return telemetryData.Takim_numarasi == data.Takim_numarasi })
	if idx != -1 {
		telemetryDatas[idx] = data
	} else {
		telemetryDatas = append(telemetryDatas, data)
	}
	out, _ := json.Marshal(telemetryDatas)
	redis.Set("telemetryDatas", string(out))
	return telemetryDatas
}
