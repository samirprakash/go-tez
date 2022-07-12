package tez

import "github.com/joho/godotenv"

const version = "1.0.0"

type Tez struct {
	AppName string
	Debug   bool
	Version string
}

func (t *Tez) New(p string) error {
	pathConfig := initPaths{
		rootPath:    p,
		folderNames: []string{"handlers", "migrations", "views", "data", "public", "tmp", "logs", "middleware"},
	}

	// init folders
	err := t.Init(pathConfig)
	if err != nil {
		return err
	}

	// check if .env file exists
	err = t.checkDotEnv(p)
	if err != nil {
		return err
	}

	// read .env file
	err = godotenv.Load(p + "/.env")
	if err != nil {
		return err
	}

	return nil
}

func (t *Tez) Init(ip initPaths) error {
	root := ip.rootPath
	for _, p := range ip.folderNames {
		err := t.CreateDirIfNotExist(root + "/" + p)
		if err != nil {
			return err
		}
	}
	return nil
}

func (t *Tez) checkDotEnv(p string) error {
	err := t.CreateFileIfNotExist(p + "/.env")
	if err != nil {
		return err
	}
	return nil
}
