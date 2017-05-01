#!/bin/bash
FILE=$(mktemp).go
LINE=$1
[ "$LINE" == "" ] && read LINE

cat > $FILE <<END
package main
import . "github.com/kavehmz/fitgo"
func main() {
    $LINE
}
END

go run $FILE