package mcpTool

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// parseAuthorityID 现委托 parseUintParam:非整数 float 应报错而非静默截断成错误角色
func TestParseAuthorityIDRejectsNonIntegerFloat(t *testing.T) {
	_, err := parseAuthorityID(float64(2.9))
	require.ErrorContains(t, err, "整数", "2.9 修复前会被截断成角色2,现应拒绝")

	id, err := parseAuthorityID(float64(888))
	require.NoError(t, err)
	require.Equal(t, uint(888), id)

	id, err = parseAuthorityID("888")
	require.NoError(t, err)
	require.Equal(t, uint(888), id)
}

// parseOptionalUint 允许 0、兼容字符串,缺省/非法回落 def
func TestParseOptionalUint(t *testing.T) {
	require.Equal(t, uint(5), parseOptionalUint("5", 0), "字符串数字应被解析(修复 parentId 错挂根)")
	require.Equal(t, uint(5), parseOptionalUint(float64(5), 0))
	require.Equal(t, uint(0), parseOptionalUint(float64(0), 9), "0 是合法值,不应回落 def")
	require.Equal(t, uint(9), parseOptionalUint(nil, 9), "缺省回落 def")
	require.Equal(t, uint(9), parseOptionalUint("abc", 9), "非法字符串回落 def")
	require.Equal(t, uint(9), parseOptionalUint(float64(2.5), 9), "非整数 float 回落 def")
}

// scalarToString 不得把大整数/时间戳打成科学计数法(dynamic_schema path/query 用它)
func TestScalarToStringNoScientificNotation(t *testing.T) {
	require.Equal(t, "1000000", scalarToString(float64(1000000)))
	require.Equal(t, "1700000000", scalarToString(float64(1700000000)))
	require.Equal(t, "123456789", scalarToString(float64(123456789)))
	require.Equal(t, "42", scalarToString(float64(42)))
	require.Equal(t, "true", scalarToString(true))
	require.Equal(t, "abc", scalarToString("abc"))
}
