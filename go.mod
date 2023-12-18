module xibolun/gotest

go 1.16

require (
	github.com/Shopify/sarama v1.30.0
	github.com/c-robinson/iplib v1.0.3
	github.com/coreos/bbolt v0.0.0-00010101000000-000000000000 // indirect
	github.com/coreos/etcd v3.3.25+incompatible
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/damonchen/rubymarshal v0.0.0-20190912082001-c9ea9b37b16f
	github.com/dgrijalva/jwt-go v3.2.0+incompatible // indirect
	github.com/elastic/go-elasticsearch/v7 v7.16.0
	github.com/elliotchance/sshtunnel v1.3.1
	github.com/go-chi/chi v1.5.4
	github.com/google/uuid v1.3.0
	github.com/gorilla/websocket v1.5.0
	github.com/gosnmp/gosnmp v1.30.0
	github.com/grpc-ecosystem/go-grpc-middleware v1.3.0 // indirect
	github.com/jonboulle/clockwork v0.2.2 // indirect
	github.com/mitchellh/mapstructure v1.4.1
	github.com/smartystreets/goconvey v1.7.2
	github.com/soheilhy/cmux v0.1.5 // indirect
	github.com/streadway/amqp v1.0.0
	github.com/stretchr/testify v1.8.1
	github.com/tidwall/gjson v1.14.4
	github.com/tmc/grpc-websocket-proxy v0.0.0-20201229170055-e5319fda7802 // indirect
	github.com/xiang90/probing v0.0.0-20190116061207-43a291ad63a2 // indirect
	github.com/zeromicro/go-zero v1.4.4
	go.mongodb.org/mongo-driver v1.11.1
	golang.org/x/crypto v0.17.0
	gopkg.in/resty.v1 v1.12.0
	sigs.k8s.io/yaml v1.3.0 // indirect
)

replace (
	// it is just sb of bbolt
	// https://github.com/etcd-io/bbolt/issues/211
	// https://github.com/etcd-io/bbolt/issues/224
	github.com/coreos/bbolt => go.etcd.io/bbolt v1.3.5
	google.golang.org/grpc => google.golang.org/grpc v1.26.0
)
