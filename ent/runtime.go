// Code generated by ent, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/LucsOlv/Turtwing_Back/ent/codeembedding"
	"github.com/LucsOlv/Turtwing_Back/ent/history"
	"github.com/LucsOlv/Turtwing_Back/ent/historyproject"
	"github.com/LucsOlv/Turtwing_Back/ent/project"
	"github.com/LucsOlv/Turtwing_Back/ent/projectmember"
	"github.com/LucsOlv/Turtwing_Back/ent/schema"
	"github.com/LucsOlv/Turtwing_Back/ent/user"
	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	codeembeddingFields := schema.CodeEmbedding{}.Fields()
	_ = codeembeddingFields
	// codeembeddingDescFileName is the schema descriptor for file_name field.
	codeembeddingDescFileName := codeembeddingFields[1].Descriptor()
	// codeembedding.FileNameValidator is a validator for the "file_name" field. It is called by the builders before save.
	codeembedding.FileNameValidator = codeembeddingDescFileName.Validators[0].(func(string) error)
	// codeembeddingDescFilePath is the schema descriptor for file_path field.
	codeembeddingDescFilePath := codeembeddingFields[2].Descriptor()
	// codeembedding.FilePathValidator is a validator for the "file_path" field. It is called by the builders before save.
	codeembedding.FilePathValidator = codeembeddingDescFilePath.Validators[0].(func(string) error)
	// codeembeddingDescCreatedAt is the schema descriptor for created_at field.
	codeembeddingDescCreatedAt := codeembeddingFields[5].Descriptor()
	// codeembedding.DefaultCreatedAt holds the default value on creation for the created_at field.
	codeembedding.DefaultCreatedAt = codeembeddingDescCreatedAt.Default.(func() time.Time)
	// codeembeddingDescUpdatedAt is the schema descriptor for updated_at field.
	codeembeddingDescUpdatedAt := codeembeddingFields[6].Descriptor()
	// codeembedding.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	codeembedding.DefaultUpdatedAt = codeembeddingDescUpdatedAt.Default.(func() time.Time)
	// codeembedding.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	codeembedding.UpdateDefaultUpdatedAt = codeembeddingDescUpdatedAt.UpdateDefault.(func() time.Time)
	// codeembeddingDescID is the schema descriptor for id field.
	codeembeddingDescID := codeembeddingFields[0].Descriptor()
	// codeembedding.DefaultID holds the default value on creation for the id field.
	codeembedding.DefaultID = codeembeddingDescID.Default.(func() uuid.UUID)
	historyFields := schema.History{}.Fields()
	_ = historyFields
	// historyDescCreatedAt is the schema descriptor for created_at field.
	historyDescCreatedAt := historyFields[3].Descriptor()
	// history.DefaultCreatedAt holds the default value on creation for the created_at field.
	history.DefaultCreatedAt = historyDescCreatedAt.Default.(func() time.Time)
	// historyDescAction is the schema descriptor for action field.
	historyDescAction := historyFields[4].Descriptor()
	// history.ActionValidator is a validator for the "action" field. It is called by the builders before save.
	history.ActionValidator = historyDescAction.Validators[0].(func(string) error)
	// historyDescDetails is the schema descriptor for details field.
	historyDescDetails := historyFields[5].Descriptor()
	// history.DefaultDetails holds the default value on creation for the details field.
	history.DefaultDetails = historyDescDetails.Default.(map[string]interface{})
	// historyDescID is the schema descriptor for id field.
	historyDescID := historyFields[0].Descriptor()
	// history.DefaultID holds the default value on creation for the id field.
	history.DefaultID = historyDescID.Default.(func() uuid.UUID)
	historyprojectFields := schema.HistoryProject{}.Fields()
	_ = historyprojectFields
	// historyprojectDescCreatedAt is the schema descriptor for created_at field.
	historyprojectDescCreatedAt := historyprojectFields[3].Descriptor()
	// historyproject.DefaultCreatedAt holds the default value on creation for the created_at field.
	historyproject.DefaultCreatedAt = historyprojectDescCreatedAt.Default.(func() time.Time)
	// historyprojectDescChanges is the schema descriptor for changes field.
	historyprojectDescChanges := historyprojectFields[4].Descriptor()
	// historyproject.DefaultChanges holds the default value on creation for the changes field.
	historyproject.DefaultChanges = historyprojectDescChanges.Default.(map[string]interface{})
	// historyprojectDescID is the schema descriptor for id field.
	historyprojectDescID := historyprojectFields[0].Descriptor()
	// historyproject.DefaultID holds the default value on creation for the id field.
	historyproject.DefaultID = historyprojectDescID.Default.(func() uuid.UUID)
	projectFields := schema.Project{}.Fields()
	_ = projectFields
	// projectDescName is the schema descriptor for name field.
	projectDescName := projectFields[1].Descriptor()
	// project.NameValidator is a validator for the "name" field. It is called by the builders before save.
	project.NameValidator = projectDescName.Validators[0].(func(string) error)
	// projectDescCreatedAt is the schema descriptor for created_at field.
	projectDescCreatedAt := projectFields[4].Descriptor()
	// project.DefaultCreatedAt holds the default value on creation for the created_at field.
	project.DefaultCreatedAt = projectDescCreatedAt.Default.(func() time.Time)
	// projectDescUpdatedAt is the schema descriptor for updated_at field.
	projectDescUpdatedAt := projectFields[5].Descriptor()
	// project.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	project.DefaultUpdatedAt = projectDescUpdatedAt.Default.(func() time.Time)
	// project.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	project.UpdateDefaultUpdatedAt = projectDescUpdatedAt.UpdateDefault.(func() time.Time)
	// projectDescSettings is the schema descriptor for settings field.
	projectDescSettings := projectFields[8].Descriptor()
	// project.DefaultSettings holds the default value on creation for the settings field.
	project.DefaultSettings = projectDescSettings.Default.(map[string]interface{})
	// projectDescID is the schema descriptor for id field.
	projectDescID := projectFields[0].Descriptor()
	// project.DefaultID holds the default value on creation for the id field.
	project.DefaultID = projectDescID.Default.(func() uuid.UUID)
	projectmemberFields := schema.ProjectMember{}.Fields()
	_ = projectmemberFields
	// projectmemberDescCreatedAt is the schema descriptor for created_at field.
	projectmemberDescCreatedAt := projectmemberFields[3].Descriptor()
	// projectmember.DefaultCreatedAt holds the default value on creation for the created_at field.
	projectmember.DefaultCreatedAt = projectmemberDescCreatedAt.Default.(func() time.Time)
	// projectmemberDescUpdatedAt is the schema descriptor for updated_at field.
	projectmemberDescUpdatedAt := projectmemberFields[4].Descriptor()
	// projectmember.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	projectmember.DefaultUpdatedAt = projectmemberDescUpdatedAt.Default.(func() time.Time)
	// projectmember.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	projectmember.UpdateDefaultUpdatedAt = projectmemberDescUpdatedAt.UpdateDefault.(func() time.Time)
	// projectmemberDescID is the schema descriptor for id field.
	projectmemberDescID := projectmemberFields[0].Descriptor()
	// projectmember.DefaultID holds the default value on creation for the id field.
	projectmember.DefaultID = projectmemberDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescEmail is the schema descriptor for email field.
	userDescEmail := userFields[1].Descriptor()
	// user.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	user.EmailValidator = userDescEmail.Validators[0].(func(string) error)
	// userDescPassword is the schema descriptor for password field.
	userDescPassword := userFields[2].Descriptor()
	// user.PasswordValidator is a validator for the "password" field. It is called by the builders before save.
	user.PasswordValidator = userDescPassword.Validators[0].(func(string) error)
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[3].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[4].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescUpdatedAt is the schema descriptor for updated_at field.
	userDescUpdatedAt := userFields[5].Descriptor()
	// user.DefaultUpdatedAt holds the default value on creation for the updated_at field.
	user.DefaultUpdatedAt = userDescUpdatedAt.Default.(func() time.Time)
	// user.UpdateDefaultUpdatedAt holds the default value on update for the updated_at field.
	user.UpdateDefaultUpdatedAt = userDescUpdatedAt.UpdateDefault.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}
