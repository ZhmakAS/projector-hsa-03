package main

import (
	"github.com/caarlos0/env/v6"
)

type Env struct {
	Addr             string   `env:"ADDRESS,required"`
	MongodbURL       string   `env:"MONGODB_URL,required"`
	ElasticSearchURL []string `env:"ELASTICSEARCH_URL,required"`
	InfluxURL        string   `env:"INFLUX_URL,required"`
	InfluxDB         string   `env:"INFLUX_DB,required"`
	InfluxDBUsername string   `env:"INFLUX_DB_USERNAME,required"`
	InfluxDBPassword string   `env:"INFLUX_DB_PASSWORD,required"`
}

func (e *Env) Parse() error {
	if err := env.Parse(e); err != nil {
		return err
	}
	return nil
}
