package es

import (
	"errors"
	"fmt"
	"log"

	"github.com/sarath.sp06/slacker/types"
)

//Index for slack messages
const INDEX = "slackmessages"

//CreateIndex Creates index 'contacts' if does not exists
//This is called to create index .Passed the elastic.Client pointer
//If the index already exists error is returned on success error will be nil
func CreateIndex() error {
	client, err := Connection()
	if err != nil {
		return err
	}

	exists, err := client.IndexExists(INDEX).Do()
	if err != nil {
		return err
	}
	if exists {
		return errors.New("Index [" + INDEX + "] exists already")
	}

	createIndex, err := client.CreateIndex(INDEX).Do()
	if err != nil {
		return err
	}
	if !createIndex.Acknowledged {
		log.Println("Create index is not acknowledged")
		exists, err := client.IndexExists(INDEX).Do()
		if err != nil {
			return err
		}

		if !exists {
			return errors.New("CreateIndex request is not acknowledged")
		}
	}

	return nil
}

//Insert iserts the message into elastic search
func Insert(sm types.SlackMessage) error {
	fmt.Println("Called Insert")
	client, err := Connection()
	if err != nil {
		return err
	}
	_, err = client.
		Index().
		Index(INDEX).
		Type(sm.Channel).
		Id(fmt.Sprintf("%s:%s", sm.User, sm.Timestamp.Format("2006-01-02 15:04:05.999999999"))).
		BodyJson(sm).
		Do()
	if err != nil {
		return err
	}
	return nil
}
