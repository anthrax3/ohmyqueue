package broker

import (
	"context"
	"math/rand"
	"strconv"
	"time"

	"github.com/astaxie/beego/logs"
	"github.com/coreos/etcd/clientv3"
	"github.com/ohmq/ohmyqueue/config"
)

func (broker *Broker) watchTopics() {
	resp, _ := broker.Client.Get(context.TODO(), "topicname", clientv3.WithPrefix())
	for _, v := range resp.Kvs {
		broker.topics = append(broker.topics, string(v.Key))
	}
	wch := broker.Client.Watch(context.TODO(), "topicname", clientv3.WithPrefix())
	for wresp := range wch {
		for _, ev := range wresp.Events {
			switch ev.Type.String() {
			case "PUT":
				logs.Info("creat topic:", string(ev.Kv.Key))
				broker.topics = append(broker.topics, string(ev.Kv.Key))
				resp, _ := broker.Client.Get(context.TODO(), "brokerleader", clientv3.WithPrefix())
				var sum int
				for _, v := range resp.Kvs {
					tmp, _ := strconv.Atoi(string(v.Value))
					sum += tmp
				}
				if len(broker.leaders) <= sum/int(resp.Count) {
					go broker.voteTopicleader(string(ev.Kv.Key[9:]))
				}
			}
		}
	}
}

func (broker *Broker) watchTopicLeader() {
	wch := broker.Client.Watch(context.TODO(), "topicleader", clientv3.WithPrefix())
	for wresp := range wch {
		for _, ev := range wresp.Events {
			switch ev.Type.String() {
			case "PUT":
				if string(ev.Kv.Value) == broker.ip+":"+broker.clientport {
					broker.leaders = append(broker.leaders, string(ev.Kv.Key[11:]))
				}
			case "DEL":
				go broker.voteTopicleader(string(ev.Kv.Key[11:]))
			}
		}
	}
}

func (broker *Broker) voteTopicleader(name string) {
	logs.Info("I am voting for topic:", name)
	<-time.After(time.Duration(rand.New(rand.NewSource(time.Now().Unix())).Intn(200)) * time.Millisecond)
	resp, err := broker.Client.Grant(context.TODO(), config.Conf.Omq.Timeout)
	if err != nil {
		logs.Error(err)
	}
	if txnresp, _ := broker.Client.Txn(context.TODO()).
		If(clientv3.Compare(clientv3.CreateRevision("topicleader"+name), "=", 0)).
		Then(clientv3.OpPut("topicleader"+name, broker.ip+":"+broker.clientport, clientv3.WithLease(resp.ID))).
		Commit(); txnresp.Succeeded {
		go broker.topicLeaderHeartbeat(resp, name)
	}
}

func (broker *Broker) topicLeaderHeartbeat(resp *clientv3.LeaseGrantResponse, name ...string) {
	for {
		<-time.After(time.Second * time.Duration((config.Conf.Omq.Timeout - 2)))
		logs.Info("topicLeaderHeartbeat for:", name)
		_, err := broker.Client.KeepAliveOnce(context.TODO(), resp.ID)
		if err != nil {
			logs.Error(err)
		}
	}
}

func (broker *Broker) watchBrokers() {
	resp, _ := broker.Client.Get(context.TODO(), "brokerid", clientv3.WithPrefix())
	for _, v := range resp.Kvs {
		if string(v.Key) != "broker"+strconv.Itoa(broker.id) {
			broker.members[string(v.Key)] = string(v.Value)
		}
	}
	wch := broker.Client.Watch(context.TODO(), "brokerid", clientv3.WithPrefix())
	for wresp := range wch {
		for _, ev := range wresp.Events {
			switch ev.Type.String() {
			case "PUT":
				logs.Info("creat broker:", string(ev.Kv.Value))
				broker.members[string(ev.Kv.Key)] = string(ev.Kv.Value)
				//TODO: sync
			case "DELETE":
				delete(broker.members, string(ev.Kv.Key))
			}
		}
	}
}