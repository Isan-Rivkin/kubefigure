package common

import (
	"bytes"
	"fmt"
	"sort"
	"strings"

	"github.com/zclconf/go-cty/cty"
	ctyjson "github.com/zclconf/go-cty/cty/json"
)

func ConvertValueToString(val cty.Value) string {
	return convertValueToString(val)
}

func convertValueToString(val cty.Value) string {
	if val.IsNull() {
		return ""
	}
	ty := val.Type()
	switch {
	case ty.IsPrimitiveType():
		switch ty {
		case cty.String:
			{
				// Special behavior for JSON strings containing array or object
				src := []byte(val.AsString())
				ty, err := ctyjson.ImpliedType(src)
				// check for the special case of "null", which decodes to nil,
				// and just allow it to be printed out directly
				if err == nil && !ty.IsPrimitiveType() && strings.TrimSpace(val.AsString()) != "null" {
					jv, err := ctyjson.Unmarshal(src, ty)
					if err != nil {
						return ""
					}
					return convertValueToString(jv)
				}
			}
			return `"` + val.AsString() + `"`
		case cty.Bool:
			if val.True() {
				return "true"
			}
			return "false"
		case cty.Number:
			bf := val.AsBigFloat()
			return bf.Text('f', -1)
		default:
			return fmt.Sprintf("%#v", val)
		}
	case ty.IsListType() || ty.IsSetType() || ty.IsTupleType():
		var b bytes.Buffer
		i := 0
		for it := val.ElementIterator(); it.Next(); {
			_, value := it.Element()
			b.WriteString(convertValueToString(value))
			if i < (val.LengthInt() - 1) {
				b.WriteString(",")
			}
			i++
		}
		if b.Len() == 0 {
			return ""
		}
		return "[" + b.String() + "]"
	case ty.IsMapType():
		var b bytes.Buffer

		i := 0
		for it := val.ElementIterator(); it.Next(); {
			key, value := it.Element()
			k := convertValueToString(key)
			v := convertValueToString(value)
			if k == "" || v == "" {
				continue
			}
			b.WriteString(k)
			b.WriteString(":")
			b.WriteString(v)
			if i < (val.LengthInt() - 1) {
				b.WriteString(",")
			}
			i++
		}
		if b.Len() == 0 {
			return ""
		}
		return "{" + b.String() + "}"
	case ty.IsObjectType():
		atys := ty.AttributeTypes()
		attrNames := make([]string, 0, len(atys))
		nameLen := 0
		for attrName := range atys {
			attrNames = append(attrNames, attrName)
			if len(attrName) > nameLen {
				nameLen = len(attrName)
			}
		}
		sort.Strings(attrNames)

		var b bytes.Buffer
		i := 0
		for _, attr := range attrNames {
			val := val.GetAttr(attr)
			v := convertValueToString(val)
			if v == "" {
				continue
			}

			b.WriteString(`"`)
			b.WriteString(attr)
			b.WriteString(`"`)
			b.WriteString(":")
			b.WriteString(v)
			if i < (len(atys) - 1) {
				b.WriteString(",")
			}
			i++
		}
		if b.Len() == 0 {
			return ""
		}
		return "{" + b.String() + "}"
	}
	return ""
}
