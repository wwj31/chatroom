package dirty

import (
	"bufio"
	"io"
	"os"
	"sync"
)

var trieTree = New()
var once sync.Once

func Init() {
	once.Do(func() {
		file, err := os.Open("list.txt")
		if err != nil {
			panic(err)
		}
		defer file.Close()

		rd := bufio.NewReader(file)
		for {
			dirtyWord, err := rd.ReadString('\n')

			dirtyWord = dirtyWord[:len(dirtyWord)-1]
			if len(dirtyWord) > 0 {
				trieTree.Add(dirtyWord)
			}

			if err != nil || err == io.EOF {
				break
			}
		}
	})

}

func CheckDirty(words string) string {
	reslut := trieTree.Match(words)
	bytes := []byte(words)
	rp := byte('*')
	for _, pair := range reslut {
		for i := pair.begin; i < pair.end; i++ {
			bytes[i] = rp
		}
	}
	return string(bytes)
}
