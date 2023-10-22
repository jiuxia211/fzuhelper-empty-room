
#!/bin/bash
CURDIR=$(cd $(dirname $0); pwd)
BinaryName=empty_room
echo "$CURDIR/bin/${BinaryName}"
exec $CURDIR/bin/${BinaryName}
