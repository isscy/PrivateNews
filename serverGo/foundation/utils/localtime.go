package utils

import (
	"encoding/json"
	"fmt"
	"time"
)

type LocalTime time.Time

// MarshalJSON 自定义json序列化
func (localTime LocalTime) MarshalJSON() ([]byte, error) {
	var jsonTimeStr = fmt.Sprintf(`"%s"`, time.Time(localTime).Format("2006-01-02 15:04:05"))
	return []byte(jsonTimeStr), nil
}

// 自定义json反序列化
func (localTime *LocalTime) UnmarshalJSON(data []byte) error {
	t, err := time.Parse(`"2006-01-02 15:04:05"`, string(data))
	if err != nil {
		return err
	}
	*localTime = LocalTime(t)
	return nil
}

// MarshalJSON 自定义json序列化
func (localTime LocalTime) AsStr() string {
	var jsonTimeStr = fmt.Sprintf(`%s`, time.Time(localTime).Format("2006-01-02 15:04:05"))
	return jsonTimeStr
}

type TestLocalTimeObject struct {
	UserName string    `json:"userName"`
	Birthday LocalTime `json:"birthday"`
}

func TestLocalTime() {
	obj := TestLocalTimeObject{
		UserName: "张三",
		Birthday: LocalTime(time.Now()),
	}
	bytes, err := json.Marshal(&obj)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("序列化：", string(bytes))
	o := new(TestLocalTimeObject)
	str := `{"userName":"张三", "birthday": "2020-07-11 12:54:40"}`
	err = json.Unmarshal([]byte(str), o)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("反序列化：name:%s, birthday:%v\n", o.UserName, time.Time(o.Birthday))
}
