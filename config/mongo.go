package config

import (
	"gopkg.in/mgo.v2"
)

// Mongo type
type Mongo struct {
	Addr     string `json:"addr"`
	Database string `json:"db"`
}

//GetAddr returns Addr property
func (m *Mongo) GetAddr() string {
	return m.Addr
}

//GetDatabase returns Database property
func (m *Mongo) GetDatabase() string {
	return m.Database
}

//NewSession returns MongoConnection
func (m *Mongo) NewSession() (session *mgo.Session, err error) {
	// database := config.Mongo.GetDatabase()
	server := m.GetAddr()
	session, err = mgo.Dial(server)
	if err != nil {
		return
	}
	return
}

//InsertData insert data in a table
func (m *Mongo) InsertData(c *Config, collection string, data []byte) (err error) {
	s, err := m.NewSession()
	if err != nil {
		return err
	}
	con := s.DB(c.Mongo.GetDatabase()).C(collection)
	err = con.Insert(data)
	if err != nil {
		return err
	}
	return
}

//GetData insert data in a table
func (m *Mongo) GetData(c *Config, query []byte, collection string) (result []byte, err error) {
	s, err := m.NewSession()
	if err != nil {
		return
	}
	con := s.DB(c.Mongo.GetDatabase()).C(collection)
	err = con.Find(query).All(&result)
	if err != nil {
		return
	}
	return
}
