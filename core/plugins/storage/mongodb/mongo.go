// Copyright 2023 The Ryan SU Authors (https://github.com/suyuan32). All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mongodb

import (
	"context"
	"fmt"
	"time"

	"github.com/gzorm/commons/core/logx"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// Conf is the MongoDB configuration structure.
type Conf struct {
	Host                  string `json:",default=localhost"`
	Username              string `json:",optional"`
	Password              string `json:",optional"`
	Port                  int    `json:",default=27017"`
	DBName                string `json:",optional"`
	Option                string `json:",optional"`
	AuthMechanism         string `json:",default=None,options=[SCRAM-SHA-256,SCRAM-SHA-1,MONGODB-X509,MONGODB-AWS,None]"`
	AuthSource            string `json:",optional"`
	TlsCAFile             string `json:",optional"`
	TlsCertificateKeyFile string `json:",optional"`
}

// MustNewClient returns the client if the config is correct.
func (c Conf) MustNewClient() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	var client *mongo.Client
	var err error

	switch c.AuthMechanism {
	case "None":
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(c.GetDSN()))
	case "SCRAM-SHA-256", "SCRAM-SHA-1", "MONGODB-AWS":
		credential := options.Credential{
			AuthMechanism: c.AuthMechanism,
			AuthSource:    c.AuthSource,
			Username:      c.Username,
			Password:      c.Password,
		}
		client, err = mongo.Connect(ctx, options.Client().ApplyURI(c.GetDSN()).SetAuth(credential))
	case "MONGODB-X509":
		if c.Option == "" {
			client, err = mongo.Connect(ctx, options.Client().ApplyURI(
				fmt.Sprintf("%s/?tlsCAFile=%s&tlsCertificateKeyFile=%s", c.GetDSN(),
					c.TlsCAFile, c.TlsCertificateKeyFile)))
		} else {
			client, err = mongo.Connect(ctx, options.Client().ApplyURI(
				fmt.Sprintf("%s&tlsCAFile=%s&tlsCertificateKeyFile=%s", c.GetDSN(),
					c.TlsCAFile, c.TlsCertificateKeyFile)))
		}
	}
	logx.Must(err)

	ctx, cancel = context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()
	err = client.Ping(ctx, readpref.Primary())
	logx.Must(err)

	return client
}

// MustNewDatabase returns the database if the config is correct.
func (c Conf) MustNewDatabase() *mongo.Database {
	return c.MustNewClient().Database(c.DBName)
}

// GetDSN returns the DSN link from the configuration.
func (c Conf) GetDSN() string {
	if c.Option == "" {
		return fmt.Sprintf("mongodb://%s:%d", c.Host, c.Port)
	} else {
		return fmt.Sprintf("mongodb://%s:%d/%s", c.Host, c.Port, c.Option)
	}
}
