package tez

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

	err := t.Init(pathConfig)
	if err != nil {
		return err
	}

	return nil
}

func (t *Tez) Init(ip initPaths) error {
	root := ip.rootPath
	for _, p := range ip.folderNames {
		// crate folder if it doesn't exist
		err := t.CreateDirIfNotExist(root + "/" + p)
		if err != nil {
			return err
		}
	}
	return nil
}
