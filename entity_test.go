package entitystore

import "testing"

//"log"

//"database/sql"
// _ "github.com/mattn/go-sqlite3"
// "gorm.io/driver/sqlite"
// "gorm.io/gorm"

func TestEntityCreate(t *testing.T) {
	db := InitDB("entity_create.db")

	store, _ := NewStore(WithDb(db), WithEntityTableName("cms_entity"), WithAttributeTableName("cms_attribute"), WithAutoMigrate(true))
	//  Init(Config{
	// 	DbInstance: db,
	// })
	entity, _ := store.EntityCreate("post")
	if entity == nil {
		t.Fatalf("Entity could not be created")
	}
}

func TestEntityCreateWithAttributes(t *testing.T) {
	db := InitDB("entity_update.db")

	store, _ := NewStore(WithDb(db), WithEntityTableName("cms_entity"), WithAttributeTableName("cms_attribute"), WithAutoMigrate(true))

	entity := store.EntityCreateWithAttributes("post", map[string]string{
		"name": "Hello world",
	})
	if entity == nil {
		t.Fatalf("Entity could not be created")
	}

	val, _ := entity.GetString("name", "")
	if val != "Hello world" {
		t.Fatalf("Entity attribute mismatch")
	}
}
