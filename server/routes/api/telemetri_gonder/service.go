package telemetri_gonder

import (
	"encoding/json"
	"log"
	"slices"
	"strconv"
	"time"
)

func PushTelemetryData(data TelemetryData) []TelemetryData {
	var telemetryDatas []TelemetryData
	if redisDatas := redis.Get("telemetryDatas"); len(redisDatas) > 0 {
		json.Unmarshal([]byte(redisDatas), &telemetryDatas)
	}
	err := redis.Set("telemetryDatasCache_"+strconv.FormatInt(data.Takim_numarasi, 10), time.Now().Format(time.RFC3339))
	if err != nil {
		log.Println("redisError: " + err.Error())
	}
	idx := slices.IndexFunc(telemetryDatas, func(telemetryData TelemetryData) bool { return telemetryData.Takim_numarasi == data.Takim_numarasi })
	if idx != -1 {
		telemetryDatas[idx] = data
	} else {
		telemetryDatas = append(telemetryDatas, data)
	}
	test := func(d TelemetryData) bool {
		t := redis.Get("telemetryDatasCache_" + strconv.FormatInt(d.Takim_numarasi, 10))
		if len(t) == 0 {
			return false
		}
		oldTime, _ := time.Parse(time.RFC3339, t)
		diff := time.Since(oldTime).Minutes()
		return diff < 5
	}
	telemetryDatas = filter(telemetryDatas, test)
	out, _ := json.Marshal(telemetryDatas)
	redis.Set("telemetryDatas", string(out))
	return telemetryDatas
}

func filter[T any](ss []T, test func(T) bool) (ret []T) {
	for _, s := range ss {
		if test(s) {
			ret = append(ret, s)
		}
	}
	return
}
