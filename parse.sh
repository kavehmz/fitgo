#!/bin/bash
FILE=$(mktemp)
LINE=$1
[ "$LINE" == "" ] && read LINE

cat > $FILE.go <<END
package main
import . "github.com/kavehmz/fitgo"
func main() {
    $LINE
}
END

go build -o $FILE $FILE.go
$FILE
