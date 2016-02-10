package dsl_test

import (
	"github.com/goadesign/gorma"
	gdsl "github.com/goadesign/gorma/dsl"

	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
	. "github.com/goadesign/goa/dslengine"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("RelationalField", func() {
	var sgname, storename, modelname, name string
	var ft gorma.FieldType
	var dsl func()
	var RandomPayload *UserTypeDefinition
	BeforeEach(func() {
		Design = nil
		Errors = nil
		sgname = "production"
		dsl = nil
		storename = "mysql"
		modelname = "Users"
		name = ""
		ft = gorma.String
		gorma.GormaDesign = nil
		InitDesign()
		RandomPayload = Type("RandomPayload", func() {
			Attribute("first_name", String)
			Attribute("last_name", String)
		})

	})

	JustBeforeEach(func() {
		gdsl.StorageGroup(sgname, func() {
			gdsl.Store(storename, gorma.MySQL, func() {
				gdsl.Model(modelname, func() {
					//gdsl.BuildsFrom(RandomPayload)
					gdsl.Field(name, ft, dsl)
					gdsl.Field("id", gorma.PKInteger, dsl) // use lowercase "id" to test sanitizer
					gdsl.Field("MiddleName", gorma.String)
					gdsl.Field("CreatedAt", gorma.Timestamp)
					gdsl.Field("UpdatedAt", gorma.Timestamp)
					gdsl.Field("DeletedAt", gorma.NullableTimestamp)
				})
			})
		})
		Run()

	})

	Context("with no DSL", func() {
		BeforeEach(func() {
			name = "FirstName"
		})

		It("produces a valid Relational Field definition", func() {
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
			sg := gorma.GormaDesign
			rs := sg.RelationalStores[storename]
			rm := rs.RelationalModels[modelname]
			Ω(rm.RelationalFields[name].FieldName).Should(Equal(name))
		})
	})

	Context("with an already defined Relational Field  with the same name", func() {
		BeforeEach(func() {
			name = "FirstName"
		})

		It("does not produce an error", func() {
			gdsl.StorageGroup(sgname, func() {
				gdsl.Store(storename, gorma.MySQL, func() {
					gdsl.Model(modelname, func() {
						gdsl.Field(name, ft, dsl)
					})
				})
			})
			Ω(Errors).Should(HaveOccurred())
		})
	})

	Context("with valid DSL", func() {
		JustBeforeEach(func() {
			Ω(Errors).ShouldNot(HaveOccurred())
			Ω(Design.Validate()).ShouldNot(HaveOccurred())
		})

		Context("sets appropriate fields and flags", func() {
			const description = "description"

			BeforeEach(func() {
				name = "FirstName"
				dsl = func() {
					gdsl.Description(description)
				}
			})

			It("sanitizes the ID field", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				rm := rs.RelationalModels[modelname]
				Ω(rm.RelationalFields).Should(HaveKey("ID"))
			})

			It("sets the relational field description", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				rm := rs.RelationalModels[modelname]
				Ω(rm.RelationalFields[name].Description).Should(Equal(description))
			})

			It("sets the field name", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				rm := rs.RelationalModels[modelname]
				Ω(rm.RelationalFields["ID"].FieldName).Should(Equal("ID"))
			})
			It("sets the field type", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				rm := rs.RelationalModels[modelname]
				Ω(rm.RelationalFields["ID"].Datatype).Should(Equal(gorma.PKInteger))
			})
			It("sets the pk flag", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				rm := rs.RelationalModels[modelname]
				Ω(rm.RelationalFields["ID"].PrimaryKey).Should(Equal(true))
			})
			It("sets has a created at field", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				rm := rs.RelationalModels[modelname]
				Ω(rm.RelationalFields["CreatedAt"].FieldName).Should(Equal("CreatedAt"))
				Ω(rm.RelationalFields["CreatedAt"].Datatype).Should(Equal(gorma.Timestamp))
				Ω(rm.RelationalFields["CreatedAt"].Nullable).Should(Equal(false))
			})
			It("sets has a deleted at field", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				rm := rs.RelationalModels[modelname]
				Ω(rm.RelationalFields["DeletedAt"].FieldName).Should(Equal("DeletedAt"))
				Ω(rm.RelationalFields["DeletedAt"].Datatype).Should(Equal(gorma.NullableTimestamp))
				Ω(rm.RelationalFields["DeletedAt"].Nullable).Should(Equal(true))
			})
			It("has the right number of fields", func() {
				sg := gorma.GormaDesign
				rs := sg.RelationalStores[storename]
				rm := rs.RelationalModels[modelname]
				length := len(rm.RelationalFields)
				Ω(length).Should(Equal(6))
			})
		})

	})

})
