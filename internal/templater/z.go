package templater

import "os"

func init() {
	templateFuncs["TASKFILE_DIR"] = func() string {
		if wd, err := os.Getwd(); err == nil {
			return wd
		}
		return ""
	}
}
