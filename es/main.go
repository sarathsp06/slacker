package es

import "gopkg.in/olivere/elastic.v2"
import "os"

var client *elastic.Client

//Connection  creates Connection or returns if there is already a Connection
func Connection() (*elastic.Client, error) {
	if client != nil {
		return client, nil
	}
	var err error
	client, err = elastic.NewClient(elastic.SetURL(os.Getenv("ELASTIC_IP")+":"+os.Getenv("ELASTIC_PORT")),
		elastic.SetMaxRetries(10))
	if err != nil {
		client = nil
	}
	return client, err
}

func init() {
	CreateIndex()
	return
}
