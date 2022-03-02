package boot

import (
	"github.com/Shopify/sarama"
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"github.com/gogf/gf/os/glog"
	"github.com/gogf/gf/os/gtime"
	"github.com/qiniu/qmgo"
	"sco-acceptor/library/mq/kafka"

	"time"

)

func init() {

	// 设置默认时区
	_, _ = time.LoadLocation("Asia/Shanghai")
	_ = gtime.SetTimeZone("Asia/Shanghai")

	s := g.Server()
	// 开启 pprof 分析
	//s.EnablePProf()

	// 固定的配置文件名称
	c := g.Cfg().SetFileName("storage.config.toml")

	// 初始日志配置
	cfginfo := c.GetMap("logger.default")
	// // glog.Println(cfginfo)
	_ = s.Logger().SetConfigWithMap(cfginfo)
	//s.Logger().SetDebug(true)
	s.Logger().SetAsync(true)
	s.Logger().SetFlags(glog.F_TIME_STD)
	// 开启日志
	s.SetErrorLogEnabled(true)
	s.SetAccessLogEnabled(true)

	// api Server配置 后台不提共 静态目录
	// publicPath := c.GetString("setting.publicPath")
	// s.SetServerRoot(publicPath)

	// TODO 关闭 静态文件服务
	s.SetFileServerEnabled(false)

	// 接口地址的 URI方式
	// s.SetNameToUriType(ghttp.NAME_TO_URI_TYPE_ALLLOWER)
	// s.SetNameToUriType(ghttp.URI_TYPE_DEFAULT)
	s.SetNameToUriType(ghttp.URI_TYPE_CAMEL)

	// 启动端口端口
	apiPort := c.GetInt("system.apiport")
	s.SetPort(apiPort)

	// 链接 mq
	//options := &PubOption{host: c.GetString("mqueue.hosts")}
	//go NewMqCliet(*options)
	//-------------- end
	kCfg := &kConfig{
		KafkaUrls:  c.GetStrings("kafka.hosts"),
		KafkaTopic: c.GetString("kafka.topic"),
	}
	initMq(kCfg)
	// --- end

	// 初始化链接 mongodb
	var timeout int64 = 50
	// Open 成功
	var maxPoolSize uint64 = 100
	var minPoolSize uint64 = 0

	cfg := &qmgo.Config{
		Uri:              c.GetString("mongo.mongoUrl"),
		ConnectTimeoutMS: &timeout,
		Database:         c.GetString("mongo.mgoDbName"),
		MaxPoolSize:      &maxPoolSize,
		MinPoolSize:      &minPoolSize,
	}
	_ = initClient(*cfg)


}

// 初始化 kafka 联接
type kConfig struct {
	KafkaUrls  []string
	KafkaTopic string
}

// initMq 初始化 联接 kafka
func initMq(conf *kConfig) {
	// 生产者配置
	kafkaConfig := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	kafkaConfig.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	kafkaConfig.Producer.Partitioner = sarama.NewReferenceHashPartitioner
	// 是否等待成功和失败后的响应
	kafkaConfig.Producer.Return.Successes = true
	// buffer 每隔多少时间触发flush
	kafkaConfig.Producer.Flush.Frequency = 5 * time.Second
	// buffer 最多装多少条消息
	kafkaConfig.Producer.Flush.MaxMessages = 10000
	// buffer 装多少条消息触发flush
	kafkaConfig.Producer.Flush.Messages = 200
	// 初始化消息队列
	kafka.InitProducer(conf.KafkaUrls, conf.KafkaTopic, kafkaConfig)
}