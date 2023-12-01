package baseproject

import (
	"os"
	"path/filepath"
)

// https://stackoverflow.com/questions/38218019/how-to-pass-a-function-to-template-through-c-html-in-gin-gonic-framework-gola
// https://chenyitian.gitbooks.io/gin-web-framework/content/docs/29.html
// https://stackoverflow.com/questions/57762069/how-to-iterate-over-a-range-of-numbers-in-golang-template
// https://gist.github.com/suntong/51e3a5c9deadba757ff0a11fd13c5826

func wrapLoadTemplates(root string) (files []string, err error) {
    err = filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
        if err != nil {
            return err
        }
        fileInfo, err := os.Stat(path)
        if err != nil {
            return err
        }
        if fileInfo.IsDir() {
            if (path != root) {
                wrapLoadTemplates(path)
            }
        } else {
            files = append(files, path)
        }
        return err
    })
    return files, err
}
