package config

import (
	"bytes"
	"testing"
)

var conf1 = `
appname: Terracotta
database:
  connection: "mongodb://admin:admin@localhost/eventsdb"
`

var conf2 = `
externals:
- name: api1
  host: "http://api1.example.com/sales"
  admin: "admin@api1.example.com"
- name: api2
  host: "http://api2.example.com/customers"
  admin: "admin@api2.example.com"
`

type ConfigApp struct {
	AppName  string `yaml:"appname"`
	Database struct {
		Connection string `yaml:"connection"`
	} `yaml:"database"`
}

type ConfigApp2 struct {
	AppName  string `yaml:"appname"`
	Database struct {
		Connection string `yaml:"connection"`
	} `yaml:"database"`
	Externals []struct {
		Name  string `yaml:"name"`
		Host  string `yaml:"host"`
		Admin string `yaml:"admin"`
	} `yaml:"externals"`
}

const (
	success = "\u2713"
	failed  = "\u2717"
)

func TestGetAppConfig(t *testing.T) {
	t.Log("Testing GetAppConfig")
	{
		appConfig := GetAppConfig()
		if appConfig == nil {
			t.Errorf("%s expected appConfig not nil, got nil", failed)
		} else {
			t.Logf("%s expected appConfig not nil, got %v", success, appConfig)
		}

		configApp := ConfigApp{}
		err := appConfig.Load(&configApp, bytes.NewBufferString(conf1))
		if err != nil {
			t.Errorf("%s expected error is nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error is nil, got nil", success)
		}

		if configApp.AppName == "Terracotta" {
			t.Logf("%s expected configApp.AppName = [Terracotta]", success)
		} else {
			t.Errorf("%s expected configApp.AppName = [Terracotta], got [%s]", failed, configApp.AppName)
		}

		if configApp.Database.Connection == "mongodb://admin:admin@localhost/eventsdb" {
			t.Logf("%s expected configApp.Database.Connection = [mongodb://admin:admin@localhost/eventsdb]", success)
		} else {
			t.Errorf("%s expected configApp.Database.Connection = [mongodb://admin:admin@localhost/eventsdb], got [%s]", failed, configApp.Database.Connection)
		}

	}

	t.Log("Testing GetAppConfig ConfigApp2 with single config string")
	{
		appConfig := GetAppConfig()
		if appConfig == nil {
			t.Errorf("%s expected appConfig not nil, got nil", failed)
		} else {
			t.Logf("%s expected appConfig not nil, got %v", success, appConfig)
		}

		configApp := ConfigApp2{}
		err := appConfig.Load(&configApp, bytes.NewBufferString(conf1))
		if err != nil {
			t.Errorf("%s expected error is nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error is nil, got nil", success)
		}

		if configApp.AppName == "Terracotta" {
			t.Logf("%s expected configApp.AppName = [Terracotta]", success)
		} else {
			t.Errorf("%s expected configApp.AppName = [Terracotta], got [%s]", failed, configApp.AppName)
		}

		if configApp.Database.Connection == "mongodb://admin:admin@localhost/eventsdb" {
			t.Logf("%s expected configApp.Database.Connection = [mongodb://admin:admin@localhost/eventsdb]", success)
		} else {
			t.Errorf("%s expected configApp.Database.Connection = [mongodb://admin:admin@localhost/eventsdb], got [%s]", failed, configApp.Database.Connection)
		}
	}

	t.Log("Testing GetAppConfig ConfigApp with multiple config string")
	{
		appConfig := GetAppConfig()
		if appConfig == nil {
			t.Errorf("%s expected appConfig not nil, got nil", failed)
		} else {
			t.Logf("%s expected appConfig not nil, got %v", success, appConfig)
		}

		configApp := ConfigApp{}
		err := appConfig.Load(&configApp, bytes.NewBufferString(conf1), bytes.NewBufferString(conf2))
		if err != nil {
			t.Errorf("%s expected error is nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error is nil, got nil", success)
		}

		if configApp.AppName == "Terracotta" {
			t.Logf("%s expected configApp.AppName = [Terracotta]", success)
		} else {
			t.Errorf("%s expected configApp.AppName = [Terracotta], got [%s]", failed, configApp.AppName)
		}

		if configApp.Database.Connection == "mongodb://admin:admin@localhost/eventsdb" {
			t.Logf("%s expected configApp.Database.Connection = [mongodb://admin:admin@localhost/eventsdb]", success)
		} else {
			t.Errorf("%s expected configApp.Database.Connection = [mongodb://admin:admin@localhost/eventsdb], got [%s]", failed, configApp.Database.Connection)
		}
	}

	t.Log("Testing GetAppConfig ConfigApp2 with multiple config string")
	{
		appConfig := GetAppConfig()
		if appConfig == nil {
			t.Errorf("%s expected appConfig not nil, got nil", failed)
		} else {
			t.Logf("%s expected appConfig not nil, got %v", success, appConfig)
		}

		configApp := ConfigApp2{}
		err := appConfig.Load(&configApp, bytes.NewBufferString(conf1), bytes.NewBufferString(conf2))
		if err != nil {
			t.Errorf("%s expected error is nil, got %s", failed, err.Error())
		} else {
			t.Logf("%s expected error is nil, got nil", success)
		}

		if configApp.AppName == "Terracotta" {
			t.Logf("%s expected configApp.AppName = [Terracotta]", success)
		} else {
			t.Errorf("%s expected configApp.AppName = [Terracotta], got [%s]", failed, configApp.AppName)
		}

		if configApp.Database.Connection == "mongodb://admin:admin@localhost/eventsdb" {
			t.Logf("%s expected configApp.Database.Connection = [mongodb://admin:admin@localhost/eventsdb]", success)
		} else {
			t.Errorf("%s expected configApp.Database.Connection = [mongodb://admin:admin@localhost/eventsdb], got [%s]", failed, configApp.Database.Connection)
		}

		if len(configApp.Externals) == 2 {
			t.Logf("%s expected len(configApp.Externals) == 2", success)

			external1 := configApp.Externals[0]
			external2 := configApp.Externals[1]

			if external1.Name == "api1" {
				t.Logf("%s expected external1.Name == [api1]", success)
			} else {
				t.Errorf("%s expected external1.Name == [api1], got %s", failed, external1.Name)
			}

			if external1.Host == "http://api1.example.com/sales" {
				t.Logf("%s expected external1.Host == [http://api1.example.com/sales]", success)
			} else {
				t.Errorf("%s expected external1.Host == [http://api1.example.com/sales], got %s", failed, external1.Host)
			}

			if external1.Admin == "admin@api1.example.com" {
				t.Logf("%s expected external1.Admin == [admin@api1.example.com]", success)
			} else {
				t.Errorf("%s expected external1.Admin == [admin@api1.example.com], got %s", failed, external1.Admin)
			}

			if external2.Name == "api2" {
				t.Logf("%s expected external2.Name == [api2]", success)
			} else {
				t.Errorf("%s expected external2.Name == [api2], got %s", failed, external2.Name)
			}

			if external2.Host == "http://api2.example.com/customers" {
				t.Logf("%s expected external2.Host == [http://api2.example.com/customers]", success)
			} else {
				t.Errorf("%s expected external2.Host == [http://api2.example.com/customers], got %s", failed, external2.Host)
			}

			if external2.Admin == "admin@api2.example.com" {
				t.Logf("%s expected external2.Admin == [admin@api2.example.com]", success)
			} else {
				t.Errorf("%s expected external2.Admin == [admin@api2.example.com], got %s", failed, external2.Admin)
			}

		} else {
			t.Errorf("%s expected len(configapp.Externals) == 2, got %d", failed, len(configApp.Externals))
		}
	}
}
