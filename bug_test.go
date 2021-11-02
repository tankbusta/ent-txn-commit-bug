package bug

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"testing"

	"entgo.io/bug/ent"
	"entgo.io/bug/ent/enttest"
	"entgo.io/ent/dialect"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
)

func TestBugSQLite(t *testing.T) {
	client := enttest.Open(t, dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	defer client.Close()
	test(t, client)
}

func TestBugMySQL(t *testing.T) {
	for version, port := range map[string]int{"56": 3306, "57": 3307, "8": 3308} {
		addr := net.JoinHostPort("localhost", strconv.Itoa(port))
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugPostrgres(t *testing.T) {
	for version, port := range map[string]int{"10": 5430, "11": 5431, "12": 5432, "13": 5433, "14": 5434} {
		t.Run(version, func(t *testing.T) {
			client := enttest.Open(t, dialect.Postgres, fmt.Sprintf("host=localhost port=%d user=postgres dbname=test password=pass sslmode=disable", port))
			defer client.Close()
			test(t, client)
		})
	}
}

func TestBugMaria(t *testing.T) {
	for version, port := range map[string]int{"10.5": 4306, "10.2": 4307, "10.3": 4308} {
		t.Run(version, func(t *testing.T) {
			addr := net.JoinHostPort("localhost", strconv.Itoa(port))
			client := enttest.Open(t, dialect.MySQL, fmt.Sprintf("root:pass@tcp(%s)/test?parseTime=True", addr))
			defer client.Close()
			test(t, client)
		})
	}
}

func test(t *testing.T, parent *ent.Client) {
	ctx := context.Background()

	client, err := parent.Tx(ctx)
	if err != nil {
		t.Errorf("unexpected error creating txn: %v", err)
	}

	createdUser, err := client.User.Create().SetName("Ariel").SetAge(30).Save(ctx)
	if err != nil {
		t.Errorf("unexpected error creating user: %v", err)
	}

	createdProduct, err := client.Product.Create().AddCreatedBy(createdUser).SetName("Expensive Shoes").Save(ctx)
	if err != nil {
		t.Errorf("unexpected error creating product: %v", err)
	}

	if err := client.Commit(); err != nil {
		t.Errorf("unexpected error committing txn: %v", err)
	}

	// NOTE: At this point the transaction has been committed but the below query will fail
	// as the created entity points to the commited transaction

	users, err := createdProduct.QueryCreatedBy().All(ctx)
	if err != nil || len(users) != 1 {
		t.Errorf("unexpected number of users: %v", err)
	}
}
