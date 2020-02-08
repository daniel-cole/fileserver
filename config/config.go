package config

import "net"

var FileServer *FileServerConfig

type FileServerConfig struct {
	Port         int          `json:"port"`
	Address      string       `json:"address"`
	SourceRanges []*net.IPNet `json:"sourceRanges"`
	Directory    string       `json:"directory"`
	HTPasswdFile string       `json:"htpasswdFile"`
}
