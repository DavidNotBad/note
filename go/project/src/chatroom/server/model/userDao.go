package model

import (
	"chatroom/common/message"
	"encoding/json"
	"fmt"
	"github.com/garyburd/redigo/redis"
)

//我们在服务器启动后， 就初始化一个userDao实例
//把它做成全局的变量， 在需要和redis操作时， 就直接使用即可
var (
	MyUserDao *UserDao
)

type UserDao struct {
	pool *redis.Pool
}

//使用工厂模式， 创建一个userDao实例
func NewUserDao(pool *redis.Pool) (userDao *UserDao) {
	userDao = &UserDao{
		pool: pool,
	}
	return
}

//根据用户id，返回一个User实例+err
func (this *UserDao) getUserById(conn redis.Conn, id int) (user *message.User, err error) {
	res, err := redis.String(conn.Do("HGet", "users", id))
	if err != nil {
		if err == redis.ErrNil {
			//没有找到对应的id
			err = ERROR_USER_NOTEXISTS
		}
		return
	}

	user = &message.User{}

	//反序列化User实例
	err = json.Unmarshal([]byte(res), &user)
	if err != nil {
		fmt.Println("json.Unmarshal err=", err)
		return
	}
	return
}

//完成登录的校验
func (this *UserDao) Login(userId int, userPwd string) (user *message.User, err error) {
	//从连接池中取出连接
	conn := this.pool.Get()
	defer conn.Close()
	user, err = this.getUserById(conn, userId)
	if err != nil {
		return
	}

	//验证密码
	if user.UserPwd != userPwd {
		err = ERROR_USER_PWD_ERROR
		return
	}
	return
}

//完成注册的校验
func (this *UserDao) Register(user *message.User) (err error) {
	//从连接池中取出连接
	conn := this.pool.Get()
	defer conn.Close()
	_, err = this.getUserById(conn, user.UserId)

	fmt.Println("resiger err", err)
	//用户已存在
	if err != ERROR_USER_NOTEXISTS {
		err = ERROR_USER_EXISTS
		return
	}

	//用户不存在， 入库
	data, err := json.Marshal(user)
	if err != nil {
		return
	}

	_, err = conn.Do("HSet", "users", user.UserId, string(data))
	if err != nil {
		fmt.Println("保存注册用户错误 err=", err)
	}

	return
}
