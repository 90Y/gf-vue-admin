// ============================================================================
// This is auto-generated by gf cli tool only once. Fill this file as you wish.
// ============================================================================

package dao

import (
	"server/app/dao/internal"
)

// dictionaryDetailsDao is the manager for logic model data accessing
// and custom defined data operations functions management. You can define
// methods on it to extend its functionality as you wish.
type dictionaryDetailsDao struct {
	*internal.DictionaryDetailsDao
}

var (
	// DictionaryDetails is globally public accessible object for table dictionary_details operations.
	DictionaryDetails = &dictionaryDetailsDao{
		internal.DictionaryDetails,
	}
)

// Fill with you ideas below.