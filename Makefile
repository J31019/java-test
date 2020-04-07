all:build

build:
	mkdir -p release
	rm -rf release/*
	mkdir release/assets
	cp ./index.html release/
	cp ./answers.html release/
	cp ./task1.html release/
	cp ./task2.html release/
	cp ./task3.html release/
	cp ./task4.html release/
	cp ./task5.html release/
	cp -R ./assets/* release/assets/
