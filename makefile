clean:
	find ./ -name "*.smgg.gen*.go" | xargs rm -rf 

install:
	go install -v ./smggcli/...