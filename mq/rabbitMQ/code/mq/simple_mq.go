package main

import (
	"github.com/streadway/amqp"
)

// URL 格式：amqp://账号:密码@rabbitmq服务器地址:端口号/vhost名称
const MQURL = "amq://hblock:hblock@127.0.0.1:5672/hblock"

type RabbitMQ struct {
	conn *amqp.Connection
}
