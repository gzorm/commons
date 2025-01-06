package zorm

type SchemaConfig struct {
	Host         string `json:",optional,env=DATABASE_HOST"`
	Port         int    `json:",optional,env=DATABASE_PORT"`
	Username     string `json:",optional,default=root,env=DATABASE_USERNAME"`
	Password     string `json:",optional,env=DATABASE_PASSWORD"`
	DBName       string `json:",optional,default=simple_admin,env=DATABASE_DBNAME"`
	SSLMode      string `json:",optional,env=DATABASE_SSL_MODE"`
	Type         string `json:",optional,default=mysql,options=[mysql,postgres,sqlite3],env=DATABASE_TYPE"`
	MaxOpenConn  int    `json:",optional,default=500,env=DATABASE_MAX_OPEN_CONN"`
	MaxIdelConn  int    `json:",optional,default=200,env=DATABASE_MAX_IDEL_CONN"`
	CacheTime    int    `json:",optional,default=10,env=DATABASE_CACHE_TIME"`
	DBPath       string `json:",optional,env=DATABASE_DBPATH"`
	MysqlConfig  string `json:",optional,env=DATABASE_MYSQL_CONFIG"`
	PGConfig     string `json:",optional,env=DATABASE_PG_CONFIG"`
	SqliteConfig string `json:",optional,env=DATABASE_SQLITE_CONFIG"`
}

type DBConfig struct {
	MainSource  Source   `json:",optional"`
	OtherSource []Source `json:",optional"`
	Prometheus  DBPrometheus
	Sharding    []DBSharding
	Mode        string
}

type Source struct {
	SchemaConfig
	Replicas []SchemaConfig `json:",optional"`
	Tables   []string       `json:",optional"`
}

type DBPrometheus struct {
	Name            string `json:",default=zorm"`
	RefreshInterval int    `json:",default=5"`
	HTTPServerPort  int    `json:",default=8084"`
}

type DBSharding struct {
	Name           string
	ShardingKey    string
	NumberOfShards int
	Tables         []string
}
