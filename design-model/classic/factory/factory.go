package factory

// 由于go本身是美誉构造函数的， 一般而言我们采用NewName方式
//创建对象/ 接口 ，
//
// 。。。。 当他返回的是接口的时候，其实就是简单工厂模式

// 有不同的parse , 接口集合
type Irun interface {
	Parse(data []byte)
}

type jsonParase struct {
}

// Parse Parse
func (J jsonParase) Parse(data []byte) {
	panic("implement me")
}

// yamlRuleConfigParser yamlRuleConfigParser
type yamlRuleConfigParser struct {
}

// Parse Parse
func (Y yamlRuleConfigParser) Parse(data []byte) {
	panic("implement me")
}

// switch case 本身就能作为工厂的一种模式
func NewRuleConfigParse(t string) Irun {
	switch t {
	case "json":
		return jsonParase{}
	case "yaml":
		return yamlRuleConfigParser{}
	}
	return nil

}

