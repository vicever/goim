// Copyright © 2014 Terry Mao, LiuDing All rights reserved.
// This file is part of gopush-cluster.

// gopush-cluster is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// gopush-cluster is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with gopush-cluster.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"flag"
	"github.com/Terry-Mao/goconf"
	"runtime"
	"time"
)

var (
	gconf    *goconf.Config
	Conf     *Config
	confFile string
)

func init() {
	flag.StringVar(&confFile, "c", "./comet.conf", " set comet config file path")
}

type Config struct {
	// base section
	PidFile   string   `goconf:"base:pidfile"`
	Dir       string   `goconf:"base:dir"`
	Log       string   `goconf:"base:log"`
	MaxProc   int      `goconf:"base:maxproc"`
	PprofBind []string `goconf:"base:pprof.bind:,"`
	StatBind  []string `goconf:"base:stat.bind:,"`
	// proto section
	TCPBind          []string      `goconf:"proto:tcp.bind:,"`
	Sndbuf           int           `goconf:"proto:sndbuf:memory"`
	Rcvbuf           int           `goconf:"proto:rcvbuf:memory"`
	HandshakeTimeout time.Duration `goconf:"proto:handshake.timeout:time"`
	WriteTimeout     time.Duration `goconf:"proto:write.timeout:time"`
	ReadBuf          int           `goconf:"proto:read.buf"`
	WriteBuf         int           `goconf:"proto:write.buf"`
	// crypto
	RSAPrivate string `goconf:"crypto:rsa.private"`
	// bucket
	Bucket   int `goconf:"bucket:bucket.num"`
	CliProto int `goconf:"bucket:cli.proto.num"`
	SvrProto int `goconf:"bucket:svr.proto.num"`
	Channel  int `goconf:"bucket:channel.num"`
}

func NewConfig() *Config {
	return &Config{
		// base section
		PidFile:   "/tmp/gopush-cluster-comet.pid",
		Dir:       "./",
		Log:       "./log/xml",
		MaxProc:   runtime.NumCPU(),
		PprofBind: []string{"localhost:6971"},
		StatBind:  []string{"localhost:6972"},
		// proto section
		TCPBind:          []string{"localhost:8080"},
		Sndbuf:           2048,
		Rcvbuf:           256,
		HandshakeTimeout: 5 * time.Second,
		WriteTimeout:     5 * time.Second,
		ReadBuf:          1024,
		WriteBuf:         1024,
		// crypto
		RSAPrivate: "./pri.pem",
		// bucket
		Bucket:   1024,
		CliProto: 1024,
		SvrProto: 1024,
		Channel:  1024,
	}
}

// InitConfig init the global config.
func InitConfig() (err error) {
	Conf = NewConfig()
	gconf = goconf.New()
	if err = gconf.Parse(confFile); err != nil {
		return err
	}
	if err := gconf.Unmarshal(Conf); err != nil {
		return err
	}
	return nil
}

func ReloadConfig() (*Config, error) {
	conf := NewConfig()
	ngconf, err := gconf.Reload()
	if err != nil {
		return nil, err
	}
	if err := ngconf.Unmarshal(conf); err != nil {
		return nil, err
	}
	gconf = ngconf
	return conf, nil
}
