package json

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestSay(t *testing.T) {
	str := "[{\"price\":\"0\",\"goods_id\":\"20972\"},{\"price\":\"0\",\"goods_id\":\"23216\"},{\"price\":\"0\",\"goods_id\":\"25032\"},{\"price\":\"0\",\"goods_id\":\"31146\"},{\"price\":\"0\",\"goods_id\":\"39744\"},{\"price\":\"0\",\"goods_id\":\"8793\"},{\"price\":\"0\",\"goods_id\":\"35196\"},{\"price\":\"0\",\"goods_id\":\"23570\"},{\"price\":\"0\",\"goods_id\":\"38022\"},{\"price\":\"0\",\"goods_id\":\"20162\"},{\"price\":\"0\",\"goods_id\":\"11881\"},{\"price\":\"0\",\"goods_id\":\"43598\"},{\"price\":\"0\",\"goods_id\":\"15227\"},{\"price\":\"0\",\"goods_id\":\"9003\"},{\"price\":\"0\",\"goods_id\":\"22856\"},{\"price\":\"0\",\"goods_id\":\"23120\"},{\"price\":\"0\",\"goods_id\":\"43600\"},{\"price\":\"0\",\"goods_id\":\"11871\"},{\"price\":\"0\",\"goods_id\":\"25476\"},{\"price\":\"0\",\"goods_id\":\"8739\"},{\"price\":\"0\",\"goods_id\":\"23732\"},{\"price\":\"0\",\"goods_id\":\"23838\"},{\"price\":\"0\",\"goods_id\":\"11411\"},{\"price\":\"0\",\"goods_id\":\"23704\"},{\"price\":\"0\",\"goods_id\":\"22434\"},{\"price\":\"0\",\"goods_id\":\"31486\"},{\"price\":\"0\",\"goods_id\":\"18336\"},{\"price\":\"0\",\"goods_id\":\"23126\"},{\"price\":\"0\",\"goods_id\":\"19266\"},{\"price\":\"0\",\"goods_id\":\"13733\"},{\"price\":\"0\",\"goods_id\":\"23720\"},{\"price\":\"0\",\"goods_id\":\"26680\"},{\"price\":\"0\",\"goods_id\":\"34738\"},{\"price\":\"0\",\"goods_id\":\"38858\"},{\"price\":\"0\",\"goods_id\":\"19576\"},{\"price\":\"0\",\"goods_id\":\"19012\"},{\"price\":\"0\",\"goods_id\":\"44220\"},{\"price\":\"0\",\"goods_id\":\"22432\"},{\"price\":\"0\",\"goods_id\":\"23812\"},{\"price\":\"0\",\"goods_id\":\"23004\"},{\"price\":\"0\",\"goods_id\":\"37838\"},{\"price\":\"0\",\"goods_id\":\"44676\"},{\"price\":\"0\",\"goods_id\":\"45678\"},{\"price\":\"0\",\"goods_id\":\"40100\"},{\"price\":\"0\",\"goods_id\":\"31206\"},{\"price\":\"0\",\"goods_id\":\"37020\"}]"
	m := make([]map[string]string, 10)
	res := make(map[string]string, 10)
	err := json.Unmarshal([]byte(str), &m)
	if err != nil {
		fmt.Printf("格式化错误 %v \n ", err)
	}
	for _, v := range m {
		res[v["goods_id"]] = v["price"]
	}
	fmt.Println(res)

}
