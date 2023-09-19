// =================================================================================
// This is auto-generated by GoFrame CLI tool only once. Fill this file as you wish.
// =================================================================================

package dao

import (
	"gov2panel/internal/dao/internal"
)

// internalV2UserDao is internal type for wrapping internal DAO implements.
type internalV2UserDao = *internal.V2UserDao

// v2UserDao is the data access object for table v2_user.
// You can define custom methods on it to extend its functionality as you wish.
type v2UserDao struct {
	internalV2UserDao
}

var (
	// V2User is globally public accessible object for table v2_user operations.
	V2User = v2UserDao{
		internal.NewV2UserDao(),
	}
)

// Fill with you ideas below.