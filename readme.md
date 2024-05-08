## DB MIGRATIONS SETUP
1. Download and install goose
```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```
2. Set up environments variables
```bash
export GOOSE_DBSTRING="user=vasudeogaichor dbname=walletpooldb password=LeaveOutAllTheRest sslmode=disable host=localhost port=5432"
```
```bash
export GOOSE_DRIVER=postgres
```
3. Create a migration
```bash
goose create create_users_table sql
```
4. Run the migration
```bash
goose up
```