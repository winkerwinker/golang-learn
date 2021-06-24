package anonymous_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

//多重继承
//func()
//*reflect.rtype
func TestAnonymouswUse11(t *testing.T) {
	model := &RequestCheckAndLoadModel{BizCode: "model", ModelCode: "model1", ModelType: "xgb", Version: "v2"}

	jsonmsg, _ := json.Marshal(model)
	fmt.Println(string(jsonmsg))

}

type RequestCheckAndLoadModel struct {
	BizCode   string
	ModelCode string `json:"modelCode"`
	ModelType string `json:"modelType"`
	Version   string `json:"version"`
	Timestamp string
	Reason    string `json:"reason"`
}
