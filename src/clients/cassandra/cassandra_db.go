package cassandra

import (
	"github.com/gocql/gocql"
)

var (

	session *gocql.Session
)

func init() {

	//IMP cluster is global (not exported variable) if we put := it is assigned to local instance of the variable.
	// So we in this case to define cluster value at global level we need not use ":="
	cluster := gocql.NewCluster("127.0.0.1")

	cluster.Keyspace = "oauth"
	cluster.Consistency = gocql.Quorum

	var err error
	if session, err = cluster.CreateSession(); err != nil {
		panic(err)
	}
}

func GetSession() (*gocql.Session) {
	return session
}
