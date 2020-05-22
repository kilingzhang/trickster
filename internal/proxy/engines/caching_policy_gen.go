package engines

// Code generated by github.com/tinylib/msgp DO NOT EDIT.

import (
	"github.com/tinylib/msgp/msgp"
)

// DecodeMsg implements msgp.Decodable
func (z *CachingPolicy) DecodeMsg(dc *msgp.Reader) (err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, err = dc.ReadMapHeader()
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, err = dc.ReadMapKeyPtr()
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "is_fresh":
			z.IsFresh, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "nocache":
			z.NoCache, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "notransform":
			z.NoTransform, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "freshness_lifetime":
			z.FreshnessLifetime, err = dc.ReadInt()
			if err != nil {
				return
			}
		case "can_revalidate":
			z.CanRevalidate, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "must_revalidate":
			z.MustRevalidate, err = dc.ReadBool()
			if err != nil {
				return
			}
		case "last_modified":
			z.LastModified, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "expires":
			z.Expires, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "date":
			z.Date, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "local_date":
			z.LocalDate, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "etag":
			z.ETag, err = dc.ReadString()
			if err != nil {
				return
			}
		case "if_none_match_value":
			z.IfNoneMatchValue, err = dc.ReadString()
			if err != nil {
				return
			}
		case "if_modified_since_time":
			z.IfModifiedSinceTime, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "if_unmodified_since_time":
			z.IfUnmodifiedSinceTime, err = dc.ReadTime()
			if err != nil {
				return
			}
		case "is_negative_cache":
			z.IsNegativeCache, err = dc.ReadBool()
			if err != nil {
				return
			}
		default:
			err = dc.Skip()
			if err != nil {
				return
			}
		}
	}
	return
}

// EncodeMsg implements msgp.Encodable
func (z *CachingPolicy) EncodeMsg(en *msgp.Writer) (err error) {
	// map header, size 15
	// write "is_fresh"
	err = en.Append(0x8f, 0xa8, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x65, 0x73, 0x68)
	if err != nil {
		return
	}
	err = en.WriteBool(z.IsFresh)
	if err != nil {
		return
	}
	// write "nocache"
	err = en.Append(0xa7, 0x6e, 0x6f, 0x63, 0x61, 0x63, 0x68, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.NoCache)
	if err != nil {
		return
	}
	// write "notransform"
	err = en.Append(0xab, 0x6e, 0x6f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	if err != nil {
		return
	}
	err = en.WriteBool(z.NoTransform)
	if err != nil {
		return
	}
	// write "freshness_lifetime"
	err = en.Append(0xb2, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteInt(z.FreshnessLifetime)
	if err != nil {
		return
	}
	// write "can_revalidate"
	err = en.Append(0xae, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.CanRevalidate)
	if err != nil {
		return
	}
	// write "must_revalidate"
	err = en.Append(0xaf, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.MustRevalidate)
	if err != nil {
		return
	}
	// write "last_modified"
	err = en.Append(0xad, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64)
	if err != nil {
		return
	}
	err = en.WriteTime(z.LastModified)
	if err != nil {
		return
	}
	// write "expires"
	err = en.Append(0xa7, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73)
	if err != nil {
		return
	}
	err = en.WriteTime(z.Expires)
	if err != nil {
		return
	}
	// write "date"
	err = en.Append(0xa4, 0x64, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteTime(z.Date)
	if err != nil {
		return
	}
	// write "local_date"
	err = en.Append(0xaa, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65)
	if err != nil {
		return
	}
	err = en.WriteTime(z.LocalDate)
	if err != nil {
		return
	}
	// write "etag"
	err = en.Append(0xa4, 0x65, 0x74, 0x61, 0x67)
	if err != nil {
		return
	}
	err = en.WriteString(z.ETag)
	if err != nil {
		return
	}
	// write "if_none_match_value"
	err = en.Append(0xb3, 0x69, 0x66, 0x5f, 0x6e, 0x6f, 0x6e, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65)
	if err != nil {
		return
	}
	err = en.WriteString(z.IfNoneMatchValue)
	if err != nil {
		return
	}
	// write "if_modified_since_time"
	err = en.Append(0xb6, 0x69, 0x66, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteTime(z.IfModifiedSinceTime)
	if err != nil {
		return
	}
	// write "if_unmodified_since_time"
	err = en.Append(0xb8, 0x69, 0x66, 0x5f, 0x75, 0x6e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	if err != nil {
		return
	}
	err = en.WriteTime(z.IfUnmodifiedSinceTime)
	if err != nil {
		return
	}
	// write "is_negative_cache"
	err = en.Append(0xb1, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65)
	if err != nil {
		return
	}
	err = en.WriteBool(z.IsNegativeCache)
	if err != nil {
		return
	}
	return
}

// MarshalMsg implements msgp.Marshaler
func (z *CachingPolicy) MarshalMsg(b []byte) (o []byte, err error) {
	o = msgp.Require(b, z.Msgsize())
	// map header, size 15
	// string "is_fresh"
	o = append(o, 0x8f, 0xa8, 0x69, 0x73, 0x5f, 0x66, 0x72, 0x65, 0x73, 0x68)
	o = msgp.AppendBool(o, z.IsFresh)
	// string "nocache"
	o = append(o, 0xa7, 0x6e, 0x6f, 0x63, 0x61, 0x63, 0x68, 0x65)
	o = msgp.AppendBool(o, z.NoCache)
	// string "notransform"
	o = append(o, 0xab, 0x6e, 0x6f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x6f, 0x72, 0x6d)
	o = msgp.AppendBool(o, z.NoTransform)
	// string "freshness_lifetime"
	o = append(o, 0xb2, 0x66, 0x72, 0x65, 0x73, 0x68, 0x6e, 0x65, 0x73, 0x73, 0x5f, 0x6c, 0x69, 0x66, 0x65, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendInt(o, z.FreshnessLifetime)
	// string "can_revalidate"
	o = append(o, 0xae, 0x63, 0x61, 0x6e, 0x5f, 0x72, 0x65, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65)
	o = msgp.AppendBool(o, z.CanRevalidate)
	// string "must_revalidate"
	o = append(o, 0xaf, 0x6d, 0x75, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x65)
	o = msgp.AppendBool(o, z.MustRevalidate)
	// string "last_modified"
	o = append(o, 0xad, 0x6c, 0x61, 0x73, 0x74, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64)
	o = msgp.AppendTime(o, z.LastModified)
	// string "expires"
	o = append(o, 0xa7, 0x65, 0x78, 0x70, 0x69, 0x72, 0x65, 0x73)
	o = msgp.AppendTime(o, z.Expires)
	// string "date"
	o = append(o, 0xa4, 0x64, 0x61, 0x74, 0x65)
	o = msgp.AppendTime(o, z.Date)
	// string "local_date"
	o = append(o, 0xaa, 0x6c, 0x6f, 0x63, 0x61, 0x6c, 0x5f, 0x64, 0x61, 0x74, 0x65)
	o = msgp.AppendTime(o, z.LocalDate)
	// string "etag"
	o = append(o, 0xa4, 0x65, 0x74, 0x61, 0x67)
	o = msgp.AppendString(o, z.ETag)
	// string "if_none_match_value"
	o = append(o, 0xb3, 0x69, 0x66, 0x5f, 0x6e, 0x6f, 0x6e, 0x65, 0x5f, 0x6d, 0x61, 0x74, 0x63, 0x68, 0x5f, 0x76, 0x61, 0x6c, 0x75, 0x65)
	o = msgp.AppendString(o, z.IfNoneMatchValue)
	// string "if_modified_since_time"
	o = append(o, 0xb6, 0x69, 0x66, 0x5f, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendTime(o, z.IfModifiedSinceTime)
	// string "if_unmodified_since_time"
	o = append(o, 0xb8, 0x69, 0x66, 0x5f, 0x75, 0x6e, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x64, 0x5f, 0x73, 0x69, 0x6e, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65)
	o = msgp.AppendTime(o, z.IfUnmodifiedSinceTime)
	// string "is_negative_cache"
	o = append(o, 0xb1, 0x69, 0x73, 0x5f, 0x6e, 0x65, 0x67, 0x61, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x63, 0x61, 0x63, 0x68, 0x65)
	o = msgp.AppendBool(o, z.IsNegativeCache)
	return
}

// UnmarshalMsg implements msgp.Unmarshaler
func (z *CachingPolicy) UnmarshalMsg(bts []byte) (o []byte, err error) {
	var field []byte
	_ = field
	var zb0001 uint32
	zb0001, bts, err = msgp.ReadMapHeaderBytes(bts)
	if err != nil {
		return
	}
	for zb0001 > 0 {
		zb0001--
		field, bts, err = msgp.ReadMapKeyZC(bts)
		if err != nil {
			return
		}
		switch msgp.UnsafeString(field) {
		case "is_fresh":
			z.IsFresh, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "nocache":
			z.NoCache, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "notransform":
			z.NoTransform, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "freshness_lifetime":
			z.FreshnessLifetime, bts, err = msgp.ReadIntBytes(bts)
			if err != nil {
				return
			}
		case "can_revalidate":
			z.CanRevalidate, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "must_revalidate":
			z.MustRevalidate, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		case "last_modified":
			z.LastModified, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "expires":
			z.Expires, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "date":
			z.Date, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "local_date":
			z.LocalDate, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "etag":
			z.ETag, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "if_none_match_value":
			z.IfNoneMatchValue, bts, err = msgp.ReadStringBytes(bts)
			if err != nil {
				return
			}
		case "if_modified_since_time":
			z.IfModifiedSinceTime, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "if_unmodified_since_time":
			z.IfUnmodifiedSinceTime, bts, err = msgp.ReadTimeBytes(bts)
			if err != nil {
				return
			}
		case "is_negative_cache":
			z.IsNegativeCache, bts, err = msgp.ReadBoolBytes(bts)
			if err != nil {
				return
			}
		default:
			bts, err = msgp.Skip(bts)
			if err != nil {
				return
			}
		}
	}
	o = bts
	return
}

// Msgsize returns an upper bound estimate of the number of bytes occupied by the serialized message
func (z *CachingPolicy) Msgsize() (s int) {
	s = 1 + 9 + msgp.BoolSize + 8 + msgp.BoolSize + 12 + msgp.BoolSize + 19 + msgp.IntSize + 15 + msgp.BoolSize + 16 + msgp.BoolSize + 14 + msgp.TimeSize + 8 + msgp.TimeSize + 5 + msgp.TimeSize + 11 + msgp.TimeSize + 5 + msgp.StringPrefixSize + len(z.ETag) + 20 + msgp.StringPrefixSize + len(z.IfNoneMatchValue) + 23 + msgp.TimeSize + 25 + msgp.TimeSize + 18 + msgp.BoolSize
	return
}