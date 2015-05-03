DEST=$(HOME)/go/src/

install:
	rm -fr $(DEST)/mol2
	mkdir -p $(DEST)/mol2
	cp -r mol2/ $(DEST)
	go install -x mol2
