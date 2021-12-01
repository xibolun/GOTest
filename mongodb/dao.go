package mongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/writeconcern"
	"log"
	"sync"
	"time"
)

// failOnError
func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

type DB struct {
	client *mongo.Client
	db     *mongo.Database
}

var globalDB *DB

func InitDB(conf *Config) {
	if conf == nil {
		return
	}
	if globalDB != nil {
		return
	}

	var once sync.Once
	once.Do(func() {
		client := GetClient(conf)
		db := client.Database(conf.Database)

		globalDB = &DB{
			client: client,
			db:     db,
		}
	})
}

// GetClient
func GetClient(conf *Config) *mongo.Client {
	opt := options.Client().ApplyURI(conf.URI)
	timeoutDur := time.Millisecond * 300
	opt.SetConnectTimeout(timeoutDur)
	opt.SetServerSelectionTimeout(timeoutDur)
	opt.SetSocketTimeout(timeoutDur)

	// config auth
	if auth := conf.Auth; auth != nil {
		var cred options.Credential
		cred.AuthMechanism = auth.AuthMechanism
		cred.AuthMechanismProperties = auth.AuthMechanismProperties
		cred.AuthSource = auth.AuthSource
		cred.Username = auth.Username
		if auth.Password != nil {
			cred.Password = *auth.Password
			cred.PasswordSet = true
		}
		opt.SetAuth(cred)
	}

	if wcConf := conf.WriteConcern; wcConf != nil {
		var wcOpts []writeconcern.Option
		if wcConf.Majority {
			wcOpts = append(wcOpts, writeconcern.WMajority())
		}
		if wcConf.TimeoutMs > 0 {
			wcOpts = append(wcOpts, writeconcern.WTimeout(time.Duration(wcConf.TimeoutMs)*time.Millisecond))
		}
		wc := writeconcern.New(wcOpts...)
		opt.SetWriteConcern(wc)
	}

	//if conf.ReadConcern != "" {
	//	if err := checkValidWriteConcern(conf.ReadConcern); err != nil {
	//		return nil, err
	//	}
	//	opt.SetReadConcern(readconcern.New(readconcern.Level(conf.ReadConcern)))
	//}

	client, _ := mongo.NewClient(opt)

	err := client.Connect(context.TODO())
	if err != nil {
		failOnError(err, "fail get connect")
	}

	return client
}

type User struct {
	HowieId     primitive.ObjectID `bson:"_id"`
	Name        string
	Pwd         string
	Age         int64
	CreateTime  int64
	ExpiredTime time.Time
}

func (u *User) Collection() string {
	return "user"
}

func Create(user *User) (id int64, err error) {
	result, err := globalDB.db.Collection(user.Collection()).InsertOne(context.TODO(), user)
	if err != nil {
		return
	}
	return result.InsertedID.(int64), nil
}

func BatchCreate(user []*User) (id []interface{}, err error) {
	result, err := globalDB.db.Collection((&User{}).Collection()).InsertMany(context.TODO(), ArrayConvert(user))
	if err != nil {
		return
	}
	return result.InsertedIDs, nil
}

func GetByID(id int64) (u *User, err error) {
	u = &User{}
	if err := globalDB.db.Collection((&User{}).Collection()).FindOne(context.TODO(), bson.D{{"_id", id}}).Decode(u); err != nil {
		return
	}
	return
}

func GetByCond(u *User) (users []*User, err error) {
	return
}

func ArrayConvert(user []*User) (result []interface{}) {
	for i := range user {
		result = append(result, user[i])
	}
	return
}
