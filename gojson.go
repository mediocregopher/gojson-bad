package gojson

import (
    "strconv"
    "encoding/json"
    "strings"
)

const (
    STRING = iota
    INT
    FLOAT
    BOOL
    LIST
    MAP
)

type Json interface {
    Type() int
    MarshalJSON() ([]byte,error)
}

/////////////////////////////////////////////////

type JsonString string
func (j JsonString) Type() int { return STRING }
func (j JsonString) MarshalJSON() ([]byte,error) {
    return []byte("\""+j+"\""),nil
}
func ToString(j Json) string {
    return string(j.(JsonString))
}
func IsString(j Json) bool {
    return j.Type() == STRING
}

/////////////////////////////////////////////////

type JsonInt int
func (j JsonInt) Type() int { return INT }
func (j JsonInt) MarshalJSON() ([]byte,error) {
    return []byte(strconv.Itoa(int(j))),nil
}
func ToInt(j Json) int {
    return int(j.(JsonInt))
}
func IsInt(j Json) bool {
    return j.Type() == INT
}

/////////////////////////////////////////////////

type JsonFloat float64
func (j JsonFloat) Type() int { return FLOAT }
func (j JsonFloat) MarshalJSON() ([]byte,error) {
    floatStr := strconv.FormatFloat(float64(j), 'g', -1, 64)
    return []byte(floatStr),nil
}
func ToFloat(j Json) float64 {
    return float64(j.(JsonFloat))
}
func IsFloat(j Json) bool {
    return j.Type() == FLOAT
}

/////////////////////////////////////////////////

type JsonBool bool
func (j JsonBool) Type() int { return BOOL }
func (j JsonBool) MarshalJSON() ([]byte,error) {
    if j {
        return []byte("true"),nil
    }
    return []byte("false"),nil
}
func ToBool(j Json) bool {
    return bool(j.(JsonBool))
}
func IsBool(j Json) bool {
    return j.Type() == BOOL
}

/////////////////////////////////////////////////

type JsonList []Json
func (j JsonList) Type() int { return LIST }
func (j JsonList) MarshalJSON() ([]byte,error) {
    jStr := make([]string,len(j))
    for i := range j {
        jStr[i] = EncodeStr(j[i])
    }
    return []byte("[" + strings.Join(jStr,",") + "]"),nil
}
func ToList(j Json) []Json {
    return []Json(j.(JsonList))
}
func IsList(j Json) bool {
    return j.Type() == LIST
}

/////////////////////////////////////////////////

type JsonMap map[string]Json
func (j JsonMap) Type() int { return MAP }
func (j JsonMap) MarshalJSON() ([]byte,error) {
    jStr := make([]string,0,len(j))
    for k,v := range j {
        jStr = append(jStr, "\""+k+"\":"+EncodeStr(v))
    }
    return []byte("{" + strings.Join(jStr,",") + "}"),nil
}
func ToMap(j Json) map[string]Json {
    return map[string]Json(j.(JsonMap))
}
func IsMap(j Json) bool {
    return j.Type() == MAP
}

/////////////////////////////////////////////////

func Encode(j Json) []byte {
    r,_ := json.Marshal(j)
    return r
}

func EncodeStr(j Json) string {
    return string(Encode(j))
}

func decodeEmptyInterface(f interface{}) Json {
    switch j := f.(type) {
        case string:
            return JsonString(j)

        case int:
            return JsonInt(j)

        case float64:
            return JsonFloat(j)

        case bool:
            return JsonBool(j)

        case nil:
            return nil

        case []interface{}:
            l := make([]Json,len(j))
            for i := range j {
                l[i] = decodeEmptyInterface(j[i])
            }
            return JsonList(l)

        case map[string]interface{}:
            m := map[string]Json{}
            for k,v := range j {
                m[k] = decodeEmptyInterface(v)
            }
            return JsonMap(m)
    }
    return nil
}

func Decode(b []byte) (Json,error) {
    var f interface{}
    err := json.Unmarshal(b,&f)
    if err != nil {
        return nil,err
    }
    return decodeEmptyInterface(f),nil
}

func DecodeStr(s string) (Json,error) {
    return Decode([]byte(s))
}
