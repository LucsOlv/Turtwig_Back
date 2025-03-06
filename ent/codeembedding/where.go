// Code generated by ent, DO NOT EDIT.

package codeembedding

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/LucsOlv/Turtwing_Back/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLTE(FieldID, id))
}

// FileName applies equality check predicate on the "file_name" field. It's identical to FileNameEQ.
func FileName(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldFileName, v))
}

// FilePath applies equality check predicate on the "file_path" field. It's identical to FilePathEQ.
func FilePath(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldFilePath, v))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldProjectID, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldUpdatedAt, v))
}

// Language applies equality check predicate on the "language" field. It's identical to LanguageEQ.
func Language(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldLanguage, v))
}

// Version applies equality check predicate on the "version" field. It's identical to VersionEQ.
func Version(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldVersion, v))
}

// FileNameEQ applies the EQ predicate on the "file_name" field.
func FileNameEQ(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldFileName, v))
}

// FileNameNEQ applies the NEQ predicate on the "file_name" field.
func FileNameNEQ(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNEQ(FieldFileName, v))
}

// FileNameIn applies the In predicate on the "file_name" field.
func FileNameIn(vs ...string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldIn(FieldFileName, vs...))
}

// FileNameNotIn applies the NotIn predicate on the "file_name" field.
func FileNameNotIn(vs ...string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNotIn(FieldFileName, vs...))
}

// FileNameGT applies the GT predicate on the "file_name" field.
func FileNameGT(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGT(FieldFileName, v))
}

// FileNameGTE applies the GTE predicate on the "file_name" field.
func FileNameGTE(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGTE(FieldFileName, v))
}

// FileNameLT applies the LT predicate on the "file_name" field.
func FileNameLT(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLT(FieldFileName, v))
}

// FileNameLTE applies the LTE predicate on the "file_name" field.
func FileNameLTE(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLTE(FieldFileName, v))
}

// FileNameContains applies the Contains predicate on the "file_name" field.
func FileNameContains(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldContains(FieldFileName, v))
}

// FileNameHasPrefix applies the HasPrefix predicate on the "file_name" field.
func FileNameHasPrefix(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldHasPrefix(FieldFileName, v))
}

// FileNameHasSuffix applies the HasSuffix predicate on the "file_name" field.
func FileNameHasSuffix(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldHasSuffix(FieldFileName, v))
}

// FileNameEqualFold applies the EqualFold predicate on the "file_name" field.
func FileNameEqualFold(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEqualFold(FieldFileName, v))
}

// FileNameContainsFold applies the ContainsFold predicate on the "file_name" field.
func FileNameContainsFold(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldContainsFold(FieldFileName, v))
}

// FilePathEQ applies the EQ predicate on the "file_path" field.
func FilePathEQ(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldFilePath, v))
}

// FilePathNEQ applies the NEQ predicate on the "file_path" field.
func FilePathNEQ(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNEQ(FieldFilePath, v))
}

// FilePathIn applies the In predicate on the "file_path" field.
func FilePathIn(vs ...string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldIn(FieldFilePath, vs...))
}

// FilePathNotIn applies the NotIn predicate on the "file_path" field.
func FilePathNotIn(vs ...string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNotIn(FieldFilePath, vs...))
}

// FilePathGT applies the GT predicate on the "file_path" field.
func FilePathGT(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGT(FieldFilePath, v))
}

// FilePathGTE applies the GTE predicate on the "file_path" field.
func FilePathGTE(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGTE(FieldFilePath, v))
}

// FilePathLT applies the LT predicate on the "file_path" field.
func FilePathLT(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLT(FieldFilePath, v))
}

// FilePathLTE applies the LTE predicate on the "file_path" field.
func FilePathLTE(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLTE(FieldFilePath, v))
}

// FilePathContains applies the Contains predicate on the "file_path" field.
func FilePathContains(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldContains(FieldFilePath, v))
}

// FilePathHasPrefix applies the HasPrefix predicate on the "file_path" field.
func FilePathHasPrefix(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldHasPrefix(FieldFilePath, v))
}

// FilePathHasSuffix applies the HasSuffix predicate on the "file_path" field.
func FilePathHasSuffix(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldHasSuffix(FieldFilePath, v))
}

// FilePathEqualFold applies the EqualFold predicate on the "file_path" field.
func FilePathEqualFold(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEqualFold(FieldFilePath, v))
}

// FilePathContainsFold applies the ContainsFold predicate on the "file_path" field.
func FilePathContainsFold(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldContainsFold(FieldFilePath, v))
}

// EmbeddingIsNil applies the IsNil predicate on the "embedding" field.
func EmbeddingIsNil() predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldIsNull(FieldEmbedding))
}

// EmbeddingNotNil applies the NotNil predicate on the "embedding" field.
func EmbeddingNotNil() predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNotNull(FieldEmbedding))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...uuid.UUID) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNotIn(FieldProjectID, vs...))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLTE(FieldUpdatedAt, v))
}

// LanguageEQ applies the EQ predicate on the "language" field.
func LanguageEQ(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldLanguage, v))
}

// LanguageNEQ applies the NEQ predicate on the "language" field.
func LanguageNEQ(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNEQ(FieldLanguage, v))
}

// LanguageIn applies the In predicate on the "language" field.
func LanguageIn(vs ...string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldIn(FieldLanguage, vs...))
}

// LanguageNotIn applies the NotIn predicate on the "language" field.
func LanguageNotIn(vs ...string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNotIn(FieldLanguage, vs...))
}

// LanguageGT applies the GT predicate on the "language" field.
func LanguageGT(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGT(FieldLanguage, v))
}

// LanguageGTE applies the GTE predicate on the "language" field.
func LanguageGTE(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGTE(FieldLanguage, v))
}

// LanguageLT applies the LT predicate on the "language" field.
func LanguageLT(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLT(FieldLanguage, v))
}

// LanguageLTE applies the LTE predicate on the "language" field.
func LanguageLTE(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLTE(FieldLanguage, v))
}

// LanguageContains applies the Contains predicate on the "language" field.
func LanguageContains(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldContains(FieldLanguage, v))
}

// LanguageHasPrefix applies the HasPrefix predicate on the "language" field.
func LanguageHasPrefix(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldHasPrefix(FieldLanguage, v))
}

// LanguageHasSuffix applies the HasSuffix predicate on the "language" field.
func LanguageHasSuffix(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldHasSuffix(FieldLanguage, v))
}

// LanguageIsNil applies the IsNil predicate on the "language" field.
func LanguageIsNil() predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldIsNull(FieldLanguage))
}

// LanguageNotNil applies the NotNil predicate on the "language" field.
func LanguageNotNil() predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNotNull(FieldLanguage))
}

// LanguageEqualFold applies the EqualFold predicate on the "language" field.
func LanguageEqualFold(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEqualFold(FieldLanguage, v))
}

// LanguageContainsFold applies the ContainsFold predicate on the "language" field.
func LanguageContainsFold(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldContainsFold(FieldLanguage, v))
}

// VersionEQ applies the EQ predicate on the "version" field.
func VersionEQ(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEQ(FieldVersion, v))
}

// VersionNEQ applies the NEQ predicate on the "version" field.
func VersionNEQ(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNEQ(FieldVersion, v))
}

// VersionIn applies the In predicate on the "version" field.
func VersionIn(vs ...string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldIn(FieldVersion, vs...))
}

// VersionNotIn applies the NotIn predicate on the "version" field.
func VersionNotIn(vs ...string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNotIn(FieldVersion, vs...))
}

// VersionGT applies the GT predicate on the "version" field.
func VersionGT(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGT(FieldVersion, v))
}

// VersionGTE applies the GTE predicate on the "version" field.
func VersionGTE(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldGTE(FieldVersion, v))
}

// VersionLT applies the LT predicate on the "version" field.
func VersionLT(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLT(FieldVersion, v))
}

// VersionLTE applies the LTE predicate on the "version" field.
func VersionLTE(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldLTE(FieldVersion, v))
}

// VersionContains applies the Contains predicate on the "version" field.
func VersionContains(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldContains(FieldVersion, v))
}

// VersionHasPrefix applies the HasPrefix predicate on the "version" field.
func VersionHasPrefix(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldHasPrefix(FieldVersion, v))
}

// VersionHasSuffix applies the HasSuffix predicate on the "version" field.
func VersionHasSuffix(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldHasSuffix(FieldVersion, v))
}

// VersionIsNil applies the IsNil predicate on the "version" field.
func VersionIsNil() predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldIsNull(FieldVersion))
}

// VersionNotNil applies the NotNil predicate on the "version" field.
func VersionNotNil() predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldNotNull(FieldVersion))
}

// VersionEqualFold applies the EqualFold predicate on the "version" field.
func VersionEqualFold(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldEqualFold(FieldVersion, v))
}

// VersionContainsFold applies the ContainsFold predicate on the "version" field.
func VersionContainsFold(v string) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.FieldContainsFold(FieldVersion, v))
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.CodeEmbedding {
	return predicate.CodeEmbedding(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(func(s *sql.Selector) {
		step := newProjectStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CodeEmbedding) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CodeEmbedding) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CodeEmbedding) predicate.CodeEmbedding {
	return predicate.CodeEmbedding(sql.NotPredicates(p))
}
