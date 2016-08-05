package util

import (
	"github.com/go-ini/ini"
)

func init() {

}

func NewIni(source interface{}) (*ini.File, error) {
	return ini.Load(source)
}
