package tez

import "os"

func (t *Tez) CreateDirIfNotExist(p string) error {
	const mode = 0755
	if _, err := os.Stat(p); os.IsNotExist(err) {
		err := os.Mkdir(p, mode)
		if err != nil {
			return err
		}
	}
	return nil
}
