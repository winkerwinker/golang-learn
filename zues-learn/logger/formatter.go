package logger

import (
	"bytes"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
)

var FieldMap = map[string][]string{
	"_com_request_in":     []string{"traceid", "spanid", "uri"},
	"_com_request_out":    []string{"traceid", "spanid", "uri", "proc_time", "errno"},
	"_com_thrift_success": []string{"traceid", "spanid", "cspanid", "uri", "interface", "proc_time", "errno"},
	"_com_thrift_failure": []string{"traceid", "spanid", "cspanid", "uri", "interface", "proc_time", "errno"},
	"_com_http_success":   []string{"traceid", "spanid", "cspanid", "uri", "url", "proc_time", "errno"},
	"_com_http_failure":   []string{"traceid", "spanid", "cspanid", "uri", "url", "proc_time", "errno"},
	"_com_cache_mis":      []string{"traceid", "spanid", "uri", "method_name", "proc_time", "errno"},
	"_com_load_cache":     []string{"traceid", "spanid", "uri", "method_name", "proc_time", "errno"},
	"_com_prediction":     []string{"traceid", "spanid", "uri", "method_name", "proc_time", "errno"},
	"_com_redis_success":  []string{"traceid", "spanid", "uri", "method_name", "proc_time", "errno"},
	"_com_redis_failure":  []string{"traceid", "spanid", "uri", "method_name", "proc_time", "errno"},
	"_undef":              []string{"traceid"},
}

type LogFormatter struct{}

func (f *LogFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	// 获取必填项
	var (
		level     = strings.ToTitle(entry.Level.String())
		timestamp = time.Now().Format("2006-01-02T15:04:05.000-0700")
	)

	var b = &bytes.Buffer{}

	// 获取 DLTAG
	DLTAG := entry.Message

	// 通过 DLTAG 获取 fixedKeys
	fixedKeys, fixedKeysExist := FieldMap[DLTAG]
	if !fixedKeysExist {
		DLTAG = "_undef"
		fixedKeys = FieldMap[DLTAG]
	}

	// 获取所有的 Keys
	var keys = make([]string, 0, len(entry.Data))
	for k := range entry.Data {
		keys = append(keys, k)
	}

	// 合并 fixedKeys 和 Keys 并去重
	var outputKeys = make([]string, 0, len(fixedKeys)+len(keys))
	outputKeys = DropDuplicatesString(append(fixedKeys, keys...))

	data := make(logrus.Fields)
	for k, v := range entry.Data {
		data[k] = v
	}

	msg := fmt.Sprintf("[%s][%s] %s||message=%s", level, timestamp, DLTAG, entry.Message)
	b.WriteString(msg)

	for _, key := range outputKeys {
		var value interface{}
		value = data[key]
		f.appendKeyValue(b, key, value)
	}

	b.WriteByte('\n')

	return b.Bytes(), nil
}

func (f *LogFormatter) appendKeyValue(b *bytes.Buffer, key string, value interface{}) {
	if b.Len() > 0 {
		b.WriteString("||")
	}
	b.WriteString(key)
	b.WriteByte('=')
	f.appendValue(b, value)
}

func (f *LogFormatter) appendValue(b *bytes.Buffer, value interface{}) {
	if value == nil {
		b.WriteString("")
		return
	}

	stringVal, ok := value.(string)
	if !ok {
		stringVal = fmt.Sprint(value)
	}

	b.WriteString(stringVal)
}


func DropDuplicatesString(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}

	return s[:j]
}