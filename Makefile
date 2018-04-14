all:
	go install
	sudo edu-graph

init-graph:
	cd data-producer && sudo python3 send-initial-data.py
