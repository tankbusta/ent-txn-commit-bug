// Code generated by entc, DO NOT EDIT.

package product

const (
	// Label holds the string label denoting the product type in the database.
	Label = "product"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// EdgeCreatedBy holds the string denoting the created_by edge name in mutations.
	EdgeCreatedBy = "created_by"
	// Table holds the table name of the product in the database.
	Table = "products"
	// CreatedByTable is the table that holds the created_by relation/edge.
	CreatedByTable = "users"
	// CreatedByInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	CreatedByInverseTable = "users"
	// CreatedByColumn is the table column denoting the created_by relation/edge.
	CreatedByColumn = "product_created_by"
)

// Columns holds all SQL columns for product fields.
var Columns = []string{
	FieldID,
	FieldName,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
