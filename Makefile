all:
	go install
	sudo edu-graph

send-package:
	cd data-producer && sudo python3 send-initial-data.py
