package proxy

import (
	"fmt"
	"log"
	"time"
)

// 代理设计模式

//接下来会通过 golang 实现静态代理，
//有 Golang 和 java 的差异性，
//我们无法比较方便的利用反射实现动态代理，
//但是我们可以利用 go generate 实现类似的效果，
//并且这样实现有两个比较大的好处，
//一个是有静态代码检查，
//我们在编译期间就可以及早发现问题，
//第二个是性能会更好。

//todo  没有动态代理

// IUser IUser
type IUser interface {
	Login(username, password string) error
}

// User 用户
type User struct {
}

// Login 用户登录
func (u *User) Login(username, password string) error {
	// 不实现细节
	fmt.Println("logininnnnn ")
	return nil
}

// UserProxy 代理类
//  ->  持有、 继承、 接口实现！！！
type UserProxy struct {
	user *User
}

// NewUserProxy NewUserProxy
func NewUserProxy(user *User) *UserProxy {
	// todo 这里用的是持有
	return &UserProxy{
		user: user,
	}
}

// Login 登录，和 user 实现相同的接口
func (p *UserProxy) Login(username, password string) error {
	// before 这里可能会有一些统计的逻辑
	start := time.Now()

	/// todo 前后加上业务逻辑
	//  -------
	// 这里是原有的业务逻辑
	if err := p.user.Login(username, password); err != nil {
		return err
	}

	// after 这里可能也有一些监控统计的逻辑
	log.Printf("user login cost time: %s", time.Now().Sub(start))

	return nil
}
