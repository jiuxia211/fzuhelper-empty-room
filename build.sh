
#!/bin/bash
RUN_NAME=empty_room
mkdir -p output/bin
cp script/* output 2>/dev/null
chmod +x output/bootstrap.sh
go build -o output/bin/${RUN_NAME}
