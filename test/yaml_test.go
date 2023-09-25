package test

import (
	"fmt"
	"jet-ci/internal/pkg/db"
	"jet-ci/internal/pkg/startup"
	"jet-ci/pkg/config"
	"testing"
)

func TestYaml(t *testing.T) {

	d := new(db.Options)
	config.Load("d:/Git/jet-ci/internal/configs/mysql-db.yaml", &d)
	fmt.Println("--->", d)

}

func TestWebYaml(t *testing.T) {

	d := new(startup.Options)
	config.Load("d:/Git/jet-ci/internal/configs/appsettings.yaml", &d)
	fmt.Println("--->", d)

}
