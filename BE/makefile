postgres:
	docker run --name formdatadb -p 1234:5432 -e POSTGRES_USER=kate -e POSTGRES_PASSWORD=secret -d postgres
createdb:
	docker exec -it formdatadb createdb --username=kate --owner=kate Assignment4
dropdb:
	docker exec -it formdatadb dropdb Assignment4
ariga:
	go get ariga.io/entimport/cmd/entimport
entimport:
	go run ariga.io/entimport/cmd/entimport -dsn 'postgres://kate:secret@localhost:1234/Assignment4?sslmode=disable'
generate:
	go generate ./ent
initent:
	go run -mod=mod entgo.io/ent/cmd/ent new 
getgen:
	go get github.com/99designs/gqlgen@v0.17.45
rungen:
	go run github.com/99designs/gqlgen generate
