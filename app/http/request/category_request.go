package request

import "gfly/app/dto"

// ====================================================================
// ========================== Add Requests ============================
// ====================================================================

// ---------------------- Create Category ------------------------

type CreateCategory struct {
	dto.CreateCategory
}

// ToDto Convert to CreateCategory DTO object.
func (r CreateCategory) ToDto() dto.CreateCategory {
	return r.CreateCategory
}

// ====================================================================
// ========================= Update Requests ==========================
// ====================================================================

// ---------------------- Update Category ------------------------

type UpdateCategory struct {
	dto.UpdateCategory
}

// ToDto Convert to UpdateCategory DTO object.
func (r UpdateCategory) ToDto() dto.UpdateCategory {
	return r.UpdateCategory
}

func (r UpdateCategory) SetID(id int) {
	r.ID = id
}
