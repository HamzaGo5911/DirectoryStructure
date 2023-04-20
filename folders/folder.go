package folder

import (
	"fmt"
	"os"
	"path/filepath"
)

func CreateFolders(numFolders int) error {
	for i := 1; i <= numFolders; i++ {
		os.MkdirAll(filepath.Join("level1", fmt.Sprintf("folder%d", i)), 0755)

		for j := 1; j <= numFolders; j++ {

			os.MkdirAll(filepath.Join("level1", fmt.Sprintf("folder%d", i), "level2", fmt.Sprintf("folder%d", j)), 0755)

			for k := 1; k <= numFolders; k++ {

				os.MkdirAll(filepath.Join("level1", fmt.Sprintf("folder%d", i), "level2", fmt.Sprintf("folder%d", j), "level3", fmt.Sprintf("folder%d", k)), 0755)

				os.MkdirAll(filepath.Join("level1", fmt.Sprintf("folder%d", i), "level2", fmt.Sprintf("folder%d", j), "level3", fmt.Sprintf("folder%d", k), "level4"), 0755)

				file, err := os.Create(filepath.Join("level1", fmt.Sprintf("folder%d", i), "level2", fmt.Sprintf("folder%d", j), "level3", fmt.Sprintf("folder%d", k), "level4", "file.txt"))

				if err != nil {
					return err
				}
				defer file.Close()
			}
		}
	}
	return nil
}
