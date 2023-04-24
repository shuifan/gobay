// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/shanbay/gobay/testdata/ent/schema"
	"github.com/shanbay/gobay/testdata/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescNickname is the schema descriptor for nickname field.
	userDescNickname := userFields[0].Descriptor()
	// user.DefaultNickname holds the default value on creation for the nickname field.
	user.DefaultNickname = userDescNickname.Default.(string)
}