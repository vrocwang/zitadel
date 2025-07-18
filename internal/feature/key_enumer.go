// Code generated by "enumer -type Key -transform snake -trimprefix Key"; DO NOT EDIT.

package feature

import (
	"fmt"
	"strings"
)

const (
	_KeyName_0      = "unspecifiedlogin_default_org"
	_KeyLowerName_0 = "unspecifiedlogin_default_org"
	_KeyName_1      = "user_schematoken_exchange"
	_KeyLowerName_1 = "user_schematoken_exchange"
	_KeyName_2      = "improved_performance"
	_KeyLowerName_2 = "improved_performance"
	_KeyName_3      = "debug_oidc_parent_erroroidc_single_v1_session_terminationdisable_user_token_eventenable_back_channel_logoutlogin_v2permission_check_v2console_use_v2_user_api"
	_KeyLowerName_3 = "debug_oidc_parent_erroroidc_single_v1_session_terminationdisable_user_token_eventenable_back_channel_logoutlogin_v2permission_check_v2console_use_v2_user_api"
)

var (
	_KeyIndex_0 = [...]uint8{0, 11, 28}
	_KeyIndex_1 = [...]uint8{0, 11, 25}
	_KeyIndex_2 = [...]uint8{0, 20}
	_KeyIndex_3 = [...]uint8{0, 23, 57, 81, 107, 115, 134, 157}
)

func (i Key) String() string {
	switch {
	case 0 <= i && i <= 1:
		return _KeyName_0[_KeyIndex_0[i]:_KeyIndex_0[i+1]]
	case 4 <= i && i <= 5:
		i -= 4
		return _KeyName_1[_KeyIndex_1[i]:_KeyIndex_1[i+1]]
	case i == 7:
		return _KeyName_2
	case 9 <= i && i <= 15:
		i -= 9
		return _KeyName_3[_KeyIndex_3[i]:_KeyIndex_3[i+1]]
	default:
		return fmt.Sprintf("Key(%d)", i)
	}
}

// An "invalid array index" compiler error signifies that the constant values have changed.
// Re-run the stringer command to generate them again.
func _KeyNoOp() {
	var x [1]struct{}
	_ = x[KeyUnspecified-(0)]
	_ = x[KeyLoginDefaultOrg-(1)]
	_ = x[KeyUserSchema-(4)]
	_ = x[KeyTokenExchange-(5)]
	_ = x[KeyImprovedPerformance-(7)]
	_ = x[KeyDebugOIDCParentError-(9)]
	_ = x[KeyOIDCSingleV1SessionTermination-(10)]
	_ = x[KeyDisableUserTokenEvent-(11)]
	_ = x[KeyEnableBackChannelLogout-(12)]
	_ = x[KeyLoginV2-(13)]
	_ = x[KeyPermissionCheckV2-(14)]
	_ = x[KeyConsoleUseV2UserApi-(15)]
}

var _KeyValues = []Key{KeyUnspecified, KeyLoginDefaultOrg, KeyUserSchema, KeyTokenExchange, KeyImprovedPerformance, KeyDebugOIDCParentError, KeyOIDCSingleV1SessionTermination, KeyDisableUserTokenEvent, KeyEnableBackChannelLogout, KeyLoginV2, KeyPermissionCheckV2, KeyConsoleUseV2UserApi}

var _KeyNameToValueMap = map[string]Key{
	_KeyName_0[0:11]:         KeyUnspecified,
	_KeyLowerName_0[0:11]:    KeyUnspecified,
	_KeyName_0[11:28]:        KeyLoginDefaultOrg,
	_KeyLowerName_0[11:28]:   KeyLoginDefaultOrg,
	_KeyName_1[0:11]:         KeyUserSchema,
	_KeyLowerName_1[0:11]:    KeyUserSchema,
	_KeyName_1[11:25]:        KeyTokenExchange,
	_KeyLowerName_1[11:25]:   KeyTokenExchange,
	_KeyName_2[0:20]:         KeyImprovedPerformance,
	_KeyLowerName_2[0:20]:    KeyImprovedPerformance,
	_KeyName_3[0:23]:         KeyDebugOIDCParentError,
	_KeyLowerName_3[0:23]:    KeyDebugOIDCParentError,
	_KeyName_3[23:57]:        KeyOIDCSingleV1SessionTermination,
	_KeyLowerName_3[23:57]:   KeyOIDCSingleV1SessionTermination,
	_KeyName_3[57:81]:        KeyDisableUserTokenEvent,
	_KeyLowerName_3[57:81]:   KeyDisableUserTokenEvent,
	_KeyName_3[81:107]:       KeyEnableBackChannelLogout,
	_KeyLowerName_3[81:107]:  KeyEnableBackChannelLogout,
	_KeyName_3[107:115]:      KeyLoginV2,
	_KeyLowerName_3[107:115]: KeyLoginV2,
	_KeyName_3[115:134]:      KeyPermissionCheckV2,
	_KeyLowerName_3[115:134]: KeyPermissionCheckV2,
	_KeyName_3[134:157]:      KeyConsoleUseV2UserApi,
	_KeyLowerName_3[134:157]: KeyConsoleUseV2UserApi,
}

var _KeyNames = []string{
	_KeyName_0[0:11],
	_KeyName_0[11:28],
	_KeyName_1[0:11],
	_KeyName_1[11:25],
	_KeyName_2[0:20],
	_KeyName_3[0:23],
	_KeyName_3[23:57],
	_KeyName_3[57:81],
	_KeyName_3[81:107],
	_KeyName_3[107:115],
	_KeyName_3[115:134],
	_KeyName_3[134:157],
}

// KeyString retrieves an enum value from the enum constants string name.
// Throws an error if the param is not part of the enum.
func KeyString(s string) (Key, error) {
	if val, ok := _KeyNameToValueMap[s]; ok {
		return val, nil
	}

	if val, ok := _KeyNameToValueMap[strings.ToLower(s)]; ok {
		return val, nil
	}
	return 0, fmt.Errorf("%s does not belong to Key values", s)
}

// KeyValues returns all values of the enum
func KeyValues() []Key {
	return _KeyValues
}

// KeyStrings returns a slice of all String values of the enum
func KeyStrings() []string {
	strs := make([]string, len(_KeyNames))
	copy(strs, _KeyNames)
	return strs
}

// IsAKey returns "true" if the value is listed in the enum definition. "false" otherwise
func (i Key) IsAKey() bool {
	for _, v := range _KeyValues {
		if i == v {
			return true
		}
	}
	return false
}
