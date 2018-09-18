package pure_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/andrewhoff/cuckoo/pure"
)

var _ = Describe("Hash", func() {
	var h *pure.Hash

	Describe("NewCuckooHash", func() {
		Context("when calling NewCuckooHash", func() {
			It("returns a new *Hash", func() {
				h = pure.NewCuckooHash(10)
				Expect(h).ToNot(BeNil())
			})
		})
	})

	Describe("Insert", func() {
		BeforeEach(func() {
			h = pure.NewCuckooHash(10)
			Expect(h).ToNot(BeNil())
		})

		Context("when trying to Insert a value into the hash", func() {
			It("inserts the value successfully", func() {
				Expect(h.Insert(5)).To(BeTrue())
			})
		})

		Context("when trying to Insert more values than the hash can hold", func() {
			It("does not fill past the capacity", func() {
				h = pure.NewCuckooHash(3)
				Expect(h.Insert(1)).To(BeTrue())
				Expect(h.Insert(6)).To(BeTrue())
				Expect(h.Insert(2)).To(BeTrue())
				Expect(h.Insert(4)).To(BeFalse())
			})
		})

		XContext("when trying to Insert a -ve value into the hash", func() {
			It("does not insert the value successfully", func() {
				Expect(h.Insert(-1)).Should(BeTrue())
			})
		})

		Context("when trying to Insert a 0 value into the hash", func() {
			It("inserts the value successfully", func() {
				Expect(h.Insert(0)).To(BeTrue())
			})
		})

		XContext("when trying to Insert a large -ve value into the hash", func() {
			It("inserts the value successfully", func() {
				Expect(h.Insert(-1123456789097654)).Should(BeTrue())
			})
		})

		Context("when trying to Insert a large +ve value into the hash", func() {
			It("inserts the value successfully", func() {
				Expect(h.Insert(123456789097654)).To(BeTrue())
			})
		})

		Context("when trying to Insert a value that will cause a cycle into the hash", func() {
			It("does not insert the value, and detects the cycle", func() {
				Expect(h.Insert(5)).To(BeTrue())
				Expect(h.Insert(15)).To(BeTrue())
				Expect(h.Insert(25)).To(BeFalse())
			})
		})
	})

	Describe("Search", func() {
		BeforeEach(func() {
			h = pure.NewCuckooHash(10)
			Expect(h).ToNot(BeNil())
		})

		Context("when trying to Search for a value that exists in the hash", func() {
			It("finds the value", func() {
				Expect(h.Insert(5)).To(BeTrue())

				Expect(h.Search(5)).To(BeTrue())
			})
		})

		Context("when trying to Search for a value that doesn't exist in the hash", func() {
			It("does not find the value", func() {
				Expect(h.Insert(5)).To(BeTrue())

				Expect(h.Search(6)).To(BeFalse())
			})
		})

		Context("when trying to Search for a values that hash to the same indices", func() {
			It("finds both of them", func() {
				Expect(h.Insert(5)).To(BeTrue())
				Expect(h.Insert(15)).To(BeTrue())

				Expect(h.Search(5)).To(BeTrue())
				Expect(h.Search(15)).To(BeTrue())

				Expect(h.Search(15)).To(BeTrue())
				Expect(h.Search(5)).To(BeTrue())
			})
		})
	})

	Describe("Delete", func() {
		BeforeEach(func() {
			h = pure.NewCuckooHash(10)
			Expect(h).ToNot(BeNil())
		})

		Context("when trying to Search for a value that exists in the hash", func() {
			It("finds the value", func() {
				Expect(h.Insert(5)).To(BeTrue())
				Expect(h.Delete(5)).To(BeTrue())
			})
		})

		Context("when trying to Search for a value that doesn't exist in the hash", func() {
			It("does not find the value", func() {
				Expect(h.Insert(5)).To(BeTrue())
				Expect(h.Delete(6)).To(BeFalse())
			})
		})

		Context("when trying to Search for a values that hash to the same indices", func() {
			It("finds both of them", func() {
				Expect(h.Insert(5)).To(BeTrue())
				Expect(h.Insert(15)).To(BeTrue())

				Expect(h.Delete(15)).To(BeTrue())
				Expect(h.Delete(5)).To(BeTrue())

				Expect(h.Search(15)).To(BeFalse())
				Expect(h.Delete(5)).To(BeFalse())
			})
		})
	})
})
