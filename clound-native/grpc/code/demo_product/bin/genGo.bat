cd .. && cd pbfile && protoc --go_out=plugins=grpc:../service --go_opt=paths=source_relative *.proto
cd ..
cd bin