package settings

import (
	"gopkg.in/ini.v1"
	"log"
)

func settings(iniFileName string, sectionName string) *ini.Section {
	cfg, err := ini.Load(iniFileName)
	if err != nil {
		log.Fatalln("Failed to read .ini file: ", err)
	}

	return cfg.Section(sectionName)
}

// GetSubServerInfo how IP or PORT
func GetSubServerInfo(keyName string) string{
	section := settings("settings/settings.ini", "sub_server")
	return section.Key(keyName).String()
}

// GetSubServerInfo how IP or PORT
func GetSumServerInfo(keyName string) string{
	section := settings("settings/settings.ini", "sum_server")
	return section.Key(keyName).String()
}

// GetSubServerInfo how IP or PORT
func GetCentralServerInfo(keyName string) string{
	section := settings("settings/settings.ini", "central_server")
	return section.Key(keyName).String()
}
