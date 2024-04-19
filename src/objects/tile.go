package objects

import (
	"crypto/md5"
	"strconv"
)

func bytesToAlphaString(b []byte) string {
    var result string
    for _, byteVal := range b {
        alphaNum := byteVal % 26
        char := 'A' + alphaNum
        result += string(char)
    }
    return result
}

func hashId(id string) string{
  hash := md5.New()
  hashDig := hash.Sum([]byte(id))

  return bytesToAlphaString(hashDig)
}

type Tile struct{
  Value string
  Selected bool
  ID string
}

func (tile *Tile) New(x int,y int)  {
  tile.ID = hashId(strconv.Itoa(x) + "--" + strconv.Itoa(y)) 
}
