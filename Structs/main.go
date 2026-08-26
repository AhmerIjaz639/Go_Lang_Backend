package main

type CloudInitConfig struct {
	Users []CloudInitUser `yaml:"users"`
}
type CloudInitUser struct {
	Name              string   `yaml:"name"`
	Gecos             string   `yaml:"gecos"`
	Shell             string   `yaml:"shell"`
	Groups            []string `yaml:"groups"`
	sshAuthorizedKeys []string `yaml:"ssh_authorized_keys"`
	Password          string   `yaml:"password"`
	Homedir           string   `yaml:"homedir"`
	Uid               int      `yaml:"uid"`
	Sudo              string   `yaml:"sudo"`
	LockPasswd        bool     `yaml:"lock_passwd"`
}
