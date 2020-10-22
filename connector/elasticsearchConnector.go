package connector

import (
	"context"
	"github.com/olivere/elastic/v7"
)

var client *elastic.Client

var host = "http://129.211.93.165:9200"

func init() {
	var err error
	client, err = elastic.NewClient(elastic.SetSniff(false), elastic.SetURL(host))
	if err != nil {
		panic(err)
	}
	_, _, err = client.Ping(host).Do(context.Background())
	if err != nil {
		panic(err)
	}
	_, err = client.ElasticsearchVersion(host)
	if err != nil {
		panic(err)
	}
}

func GetElasticConn() *elastic.Client {
	return client
}
