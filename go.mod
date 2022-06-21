module i-go

go 1.17

// grpcv1.26.0以上版本 和 etcd 冲突
replace google.golang.org/grpc => google.golang.org/grpc v1.26.0

require (
	entgo.io/ent v0.10.0
	github.com/fvbock/endless v0.0.0-20170109170031-447134032cb6
	github.com/mitchellh/go-homedir v1.1.0
	github.com/mosn/holmes v0.0.0-20220125114618-8cb365eb42ac
	github.com/pyroscope-io/client v0.2.0
	github.com/shirou/gopsutil/v3 v3.21.8
	github.com/spf13/cobra v1.3.0
	google.golang.org/genproto v0.0.0-20220107163113-42d7afdf6368
	k8s.io/api v0.20.1
	k8s.io/apimachinery v0.23.5
	k8s.io/client-go v0.20.1
	k8s.io/utils v0.0.0-20211116205334-6203023598ed
)

require (
	ariga.io/atlas v0.3.2-0.20220120225051-c3fac7d636dd // indirect
	github.com/agext/levenshtein v1.2.1 // indirect
	github.com/apparentlymart/go-textseg/v13 v13.0.0 // indirect
	github.com/coreos/etcd v3.3.27+incompatible // indirect
	github.com/coreos/go-semver v0.3.0 // indirect
	github.com/coreos/pkg v0.0.0-20180928190104-399ea9e2e55f // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-openapi/inflect v0.19.0 // indirect
	github.com/google/go-cmp v0.5.6 // indirect
	github.com/google/gofuzz v1.1.0 // indirect
	github.com/googleapis/gnostic v0.5.5 // indirect
	github.com/hashicorp/hcl/v2 v2.10.0 // indirect
	github.com/imdario/mergo v0.3.11 // indirect
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/mitchellh/go-wordwrap v0.0.0-20150314170334-ad45545899c7 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.0.2 // indirect
	github.com/xdg-go/stringprep v1.0.2 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	github.com/zclconf/go-cty v1.8.0 // indirect
	go.etcd.io/bbolt v1.3.6 // indirect
	golang.org/x/mod v0.6.0-dev.0.20220106191415-9b9b3d81d5e3 // indirect
	golang.org/x/oauth2 v0.0.0-20211104180415-d3ed0bb246c8 // indirect
	golang.org/x/term v0.0.0-20210927222741-03fcf44c2211 // indirect
	golang.org/x/xerrors v0.0.0-20220411194840-2f41105eb62f // indirect
	gopkg.in/inf.v0 v0.9.1 // indirect
	k8s.io/klog/v2 v2.60.1 // indirect
	sigs.k8s.io/json v0.0.0-20211020170558-c049b76a60c6 // indirect
	sigs.k8s.io/structured-merge-diff/v4 v4.2.1 // indirect
	sigs.k8s.io/yaml v1.3.0 // indirect
)

require (
	github.com/Azure/go-ansiterm v0.0.0-20170929234023-d6e3b3328b78 // indirect
	github.com/HdrHistogram/hdrhistogram-go v1.1.2 // indirect
	github.com/KyleBanks/depth v1.2.1 // indirect
	github.com/Microsoft/go-winio v0.4.17-0.20210211115548-6eac466e5fa3 // indirect
	github.com/Microsoft/hcsshim v0.8.16 // indirect
	github.com/PuerkitoBio/goquery v1.5.1 // indirect
	github.com/Shopify/sarama v1.30.1
	github.com/StackExchange/wmi v1.2.1 // indirect
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5 // indirect
	github.com/alecthomas/template v0.0.0-20190718012654-fb15b899a751 // indirect
	github.com/alecthomas/units v0.0.0-20190924025748-f65c72e2690d // indirect
	github.com/andres-erbsen/clock v0.0.0-20160526145045-9e14626cd129 // indirect
	github.com/andybalholm/cascadia v1.1.0 // indirect
	github.com/antchfx/htmlquery v1.0.0 // indirect
	github.com/antchfx/xmlquery v1.0.0 // indirect
	github.com/antchfx/xpath v0.0.0-20190319080838-ce1d48779e67 // indirect
	github.com/aws/aws-sdk-go v1.42.23 // indirect
	github.com/axiomhq/hyperloglog v0.0.0-20211021164851-7f2dfa314bc7
	github.com/beevik/etree v1.1.0 // indirect
	github.com/beorn7/perks v1.0.1 // indirect
	github.com/bits-and-blooms/bitset v1.2.0 // indirect
	github.com/bits-and-blooms/bloom/v3 v3.0.1
	github.com/bwmarrin/snowflake v0.3.0
	github.com/cenkalti/backoff v2.2.1+incompatible // indirect
	github.com/cespare/xxhash/v2 v2.1.2 // indirect
	github.com/containerd/cgroups v0.0.0-20210114181951-8a68de567b68 // indirect
	github.com/containerd/containerd v1.5.0-beta.4 // indirect
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/dgryski/go-metro v0.0.0-20180109044635-280f6062b5bc // indirect
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/docker/docker v20.10.11+incompatible // indirect
	github.com/docker/go-connections v0.4.0 // indirect
	github.com/docker/go-units v0.4.0 // indirect
	github.com/eapache/go-resiliency v1.2.0 // indirect
	github.com/eapache/go-xerial-snappy v0.0.0-20180814174437-776d5712da21 // indirect
	github.com/eapache/queue v1.1.0 // indirect
	github.com/elastic/go-elasticsearch/v7 v7.5.1-0.20200623144321-26ee26eb5407
	github.com/fatih/structs v1.1.0 // indirect
	github.com/fsnotify/fsnotify v1.5.1
	github.com/gin-contrib/pprof v1.3.0
	github.com/gin-contrib/sse v0.1.0 // indirect
	github.com/gin-gonic/gin v1.7.7
	github.com/go-ego/cedar v0.10.2 // indirect
	github.com/go-ego/gse v0.63.2
	github.com/go-kit/kit v0.10.0
	github.com/go-kratos/aegis v0.1.1
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/go-ole/go-ole v1.2.6 // indirect
	github.com/go-openapi/jsonpointer v0.19.5 // indirect
	github.com/go-openapi/jsonreference v0.20.0 // indirect
	github.com/go-openapi/spec v0.20.4 // indirect
	github.com/go-openapi/swag v0.21.1 // indirect
	github.com/go-pg/pg v8.0.6+incompatible
	github.com/go-playground/locales v0.14.0
	github.com/go-playground/universal-translator v0.18.0
	github.com/go-playground/validator/v10 v10.10.0
	github.com/go-redis/redis v6.15.9+incompatible
	github.com/go-sql-driver/mysql v1.6.0 // indirect
	github.com/go-stack/stack v1.8.0 // indirect
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0
	github.com/gogf/gf v1.13.1
	github.com/gogo/protobuf v1.3.2 // indirect
	github.com/golang-jwt/jwt v3.2.2+incompatible
	github.com/golang/groupcache v0.0.0-20210331224755-41bb18bfe9da // indirect
	github.com/golang/mock v1.6.0
	github.com/golang/protobuf v1.5.2
	github.com/golang/snappy v0.0.4 // indirect
	github.com/google/uuid v1.3.0
	github.com/google/wire v0.5.0
	github.com/gqcn/structs v1.1.1 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.16.0
	github.com/hashicorp/go-uuid v1.0.2 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/ikeikeikeike/go-sitemap-generator/v2 v2.0.2
	github.com/jcmturner/aescts/v2 v2.0.0 // indirect
	github.com/jcmturner/dnsutils/v2 v2.0.0 // indirect
	github.com/jcmturner/gofork v1.0.0 // indirect
	github.com/jcmturner/gokrb5/v8 v8.4.2 // indirect
	github.com/jcmturner/rpc/v2 v2.0.3 // indirect
	github.com/jinzhu/gorm v1.9.16
	github.com/jinzhu/inflection v1.0.0 // indirect
	github.com/jinzhu/now v1.1.4 // indirect
	github.com/jmespath/go-jmespath v0.4.0 // indirect
	github.com/josharian/intern v1.0.0 // indirect
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/klauspost/compress v1.13.6 // indirect
	github.com/leodido/go-urn v1.2.1 // indirect
	github.com/lionsoul2014/ip2region v2.2.0-release+incompatible
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mailru/easyjson v0.7.7 // indirect
	github.com/mattn/go-isatty v0.0.14 // indirect
	github.com/mattn/go-sqlite3 v2.0.1+incompatible
	github.com/matttproud/golang_protobuf_extensions v1.0.2-0.20181231171920-c182affec369 // indirect
	github.com/mfonda/simhash v0.0.0-20151007195837-79f94a1100d6
	github.com/mitchellh/mapstructure v1.4.3
	github.com/moby/sys/mount v0.2.0 // indirect
	github.com/moby/sys/mountinfo v0.5.0 // indirect
	github.com/moby/term v0.0.0-20201216013528-df9cb8a40635 // indirect
	github.com/modern-go/concurrent v0.0.0-20180306012644-bacd9c7ef1dd // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/morikuni/aec v0.0.0-20170113033406-39771216ff4c // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/nats-io/jwt v0.3.2 // indirect
	github.com/nats-io/nats-server/v2 v2.1.4 // indirect
	github.com/nats-io/nats-streaming-server v0.17.0 // indirect
	github.com/nats-io/nats.go v1.10.0
	github.com/nats-io/nkeys v0.1.4 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/nats-io/stan.go v0.6.0
	github.com/olivere/elastic/v7 v7.0.30
	github.com/opencontainers/go-digest v1.0.0 // indirect
	github.com/opencontainers/image-spec v1.0.1 // indirect
	github.com/opencontainers/runc v1.0.2 // indirect
	github.com/opentracing/opentracing-go v1.2.0
	github.com/orcaman/concurrent-map v0.0.0-20190826125027-8c72a8bb44f6
	github.com/oschwald/geoip2-golang v1.5.0
	github.com/oschwald/maxminddb-golang v1.8.0 // indirect
	github.com/pelletier/go-toml v1.9.4 // indirect
	github.com/pierrec/lz4 v2.6.1+incompatible // indirect
	github.com/pkg/errors v0.9.1
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/prometheus/client_golang v1.11.1
	github.com/prometheus/client_model v0.2.0 // indirect
	github.com/prometheus/common v0.26.0
	github.com/prometheus/procfs v0.6.0 // indirect
	github.com/rcrowley/go-metrics v0.0.0-20201227073835-cf1acfcdf475 // indirect
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/robfig/cron v1.1.0
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/satori/go.uuid v1.2.0
	github.com/shirou/gopsutil v3.21.11+incompatible // indirect
	github.com/sirupsen/logrus v1.8.1
	github.com/smartwalle/alipay/v3 v3.1.6
	github.com/smartwalle/crypto4go v1.0.3 // indirect
	github.com/sony/gobreaker v0.4.1
	github.com/sony/sonyflake v1.0.0
	github.com/spaolacci/murmur3 v1.1.0
	github.com/spf13/afero v1.6.0 // indirect
	github.com/spf13/cast v1.4.1 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/spf13/viper v1.10.1
	github.com/streadway/amqp v0.0.0-20200108173154-1c71cc93ed71
	github.com/streadway/handy v0.0.0-20190108123426-d5acb3125c2a // indirect
	github.com/stretchr/testify v1.7.1-0.20210427113832-6241f9ab9942
	github.com/subosito/gotenv v1.2.0 // indirect
	github.com/swaggo/gin-swagger v1.3.3
	github.com/swaggo/swag v1.7.6
	github.com/temoto/robotstxt v0.0.0-20180810133444-97ee4a9ee6ea // indirect
	github.com/testcontainers/testcontainers-go v0.12.0
	github.com/tidwall/pretty v1.0.1 // indirect
	github.com/tklauser/go-sysconf v0.3.9 // indirect
	github.com/tklauser/numcpus v0.4.0 // indirect
	github.com/uber/jaeger-client-go v2.25.0+incompatible
	github.com/uber/jaeger-lib v2.4.0+incompatible // indirect
	github.com/ugorji/go/codec v1.2.7 // indirect
	github.com/unrolled/secure v1.0.8
	github.com/wechatpay-apiv3/wechatpay-go v0.2.9
	github.com/willf/bitset v1.1.11 // indirect
	github.com/willf/bloom v2.0.3+incompatible
	github.com/yusufpapurcu/wmi v1.2.2 // indirect
	go.etcd.io/etcd v3.3.13+incompatible
	go.mongodb.org/mongo-driver v1.5.1
	go.opencensus.io v0.23.0 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/automaxprocs v1.4.0
	go.uber.org/multierr v1.7.0 // indirect
	go.uber.org/ratelimit v0.2.0
	go.uber.org/zap v1.19.1
	golang.org/x/crypto v0.0.0-20220214200702-86341886e292
	golang.org/x/net v0.0.0-20220412020605-290c469a71a5
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	golang.org/x/sys v0.0.0-20220412211240-33da011f77ad // indirect
	golang.org/x/text v0.3.7 // indirect
	golang.org/x/time v0.0.0-20200630173020-3af7569d3a1e
	golang.org/x/tools v0.1.10 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/grpc v1.43.0
	google.golang.org/protobuf v1.28.0
	gopkg.in/alecthomas/kingpin.v2 v2.2.6 // indirect
	gopkg.in/ini.v1 v1.66.2 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.0-20210107192922-496545a6307b // indirect
	gorm.io/driver/mysql v1.2.1
	gorm.io/gorm v1.22.4
	mellium.im/sasl v0.2.1 // indirect
)
