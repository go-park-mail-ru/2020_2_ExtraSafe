// Code generated by easyjson for marshaling/unmarshaling. DO NOT EDIT.

package models

import (
	json "encoding/json"
	easyjson "github.com/mailru/easyjson"
	jlexer "github.com/mailru/easyjson/jlexer"
	jwriter "github.com/mailru/easyjson/jwriter"
)

// suppress unused package warning
var (
	_ *json.RawMessage
	_ *jlexer.Lexer
	_ *jwriter.Writer
	_ easyjson.Marshaler
)

func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(in *jlexer.Lexer, out *BoardOutsideShort) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "boardID":
			out.BoardID = int64(in.Int64())
		case "boardName":
			out.Name = string(in.String())
		case "boardTheme":
			out.Theme = string(in.String())
		case "boardStar":
			out.Star = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(out *jwriter.Writer, in BoardOutsideShort) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"boardID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.BoardID))
	}
	{
		const prefix string = ",\"boardName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"boardTheme\":"
		out.RawString(prefix)
		out.String(string(in.Theme))
	}
	{
		const prefix string = ",\"boardStar\":"
		out.RawString(prefix)
		out.Bool(bool(in.Star))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardOutsideShort) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardOutsideShort) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardOutsideShort) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardOutsideShort) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(in *jlexer.Lexer, out *BoardOutside) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "boardID":
			out.BoardID = int64(in.Int64())
		case "boardAdmin":
			(out.Admin).UnmarshalEasyJSON(in)
		case "boardName":
			out.Name = string(in.String())
		case "boardTheme":
			out.Theme = string(in.String())
		case "boardStar":
			out.Star = bool(in.Bool())
		case "boardMembers":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]UserOutsideShort, 0, 0)
					} else {
						out.Users = []UserOutsideShort{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v1 UserOutsideShort
					(v1).UnmarshalEasyJSON(in)
					out.Users = append(out.Users, v1)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "boardCards":
			if in.IsNull() {
				in.Skip()
				out.Cards = nil
			} else {
				in.Delim('[')
				if out.Cards == nil {
					if !in.IsDelim(']') {
						out.Cards = make([]CardOutside, 0, 1)
					} else {
						out.Cards = []CardOutside{}
					}
				} else {
					out.Cards = (out.Cards)[:0]
				}
				for !in.IsDelim(']') {
					var v2 CardOutside
					(v2).UnmarshalEasyJSON(in)
					out.Cards = append(out.Cards, v2)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "boardTags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]TagOutside, 0, 1)
					} else {
						out.Tags = []TagOutside{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v3 TagOutside
					(v3).UnmarshalEasyJSON(in)
					out.Tags = append(out.Tags, v3)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(out *jwriter.Writer, in BoardOutside) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"boardID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.BoardID))
	}
	{
		const prefix string = ",\"boardAdmin\":"
		out.RawString(prefix)
		(in.Admin).MarshalEasyJSON(out)
	}
	{
		const prefix string = ",\"boardName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"boardTheme\":"
		out.RawString(prefix)
		out.String(string(in.Theme))
	}
	{
		const prefix string = ",\"boardStar\":"
		out.RawString(prefix)
		out.Bool(bool(in.Star))
	}
	{
		const prefix string = ",\"boardMembers\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v4, v5 := range in.Users {
				if v4 > 0 {
					out.RawByte(',')
				}
				(v5).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"boardCards\":"
		out.RawString(prefix)
		if in.Cards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v6, v7 := range in.Cards {
				if v6 > 0 {
					out.RawByte(',')
				}
				(v7).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"boardTags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v8, v9 := range in.Tags {
				if v8 > 0 {
					out.RawByte(',')
				}
				(v9).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardOutside) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardOutside) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardOutside) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardOutside) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(in *jlexer.Lexer, out *BoardMemberOutside) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "boardName":
			out.BoardName = string(in.String())
		case "initiator":
			out.Initiator = string(in.String())
		case "email":
			out.Email = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "fullName":
			out.FullName = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(out *jwriter.Writer, in BoardMemberOutside) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"boardName\":"
		out.RawString(prefix[1:])
		out.String(string(in.BoardName))
	}
	{
		const prefix string = ",\"initiator\":"
		out.RawString(prefix)
		out.String(string(in.Initiator))
	}
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix)
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"fullName\":"
		out.RawString(prefix)
		out.String(string(in.FullName))
	}
	{
		const prefix string = ",\"avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardMemberOutside) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardMemberOutside) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardMemberOutside) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardMemberOutside) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(in *jlexer.Lexer, out *BoardMemberInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "boardID":
			out.BoardID = int64(in.Int64())
		case "memberUsername":
			out.MemberName = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(out *jwriter.Writer, in BoardMemberInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"boardID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.BoardID))
	}
	{
		const prefix string = ",\"memberUsername\":"
		out.RawString(prefix)
		out.String(string(in.MemberName))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardMemberInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardMemberInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardMemberInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardMemberInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(in *jlexer.Lexer, out *BoardMember) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "UserID":
			out.UserID = int64(in.Int64())
		case "BoardID":
			out.BoardID = int64(in.Int64())
		case "MemberID":
			out.MemberID = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(out *jwriter.Writer, in BoardMember) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"UserID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.UserID))
	}
	{
		const prefix string = ",\"BoardID\":"
		out.RawString(prefix)
		out.Int64(int64(in.BoardID))
	}
	{
		const prefix string = ",\"MemberID\":"
		out.RawString(prefix)
		out.Int64(int64(in.MemberID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardMember) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardMember) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardMember) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardMember) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(in *jlexer.Lexer, out *BoardInviteInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "UserID":
			out.UserID = int64(in.Int64())
		case "BoardID":
			out.BoardID = int64(in.Int64())
		case "UrlHash":
			out.UrlHash = string(in.String())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(out *jwriter.Writer, in BoardInviteInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"UserID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.UserID))
	}
	{
		const prefix string = ",\"BoardID\":"
		out.RawString(prefix)
		out.Int64(int64(in.BoardID))
	}
	{
		const prefix string = ",\"UrlHash\":"
		out.RawString(prefix)
		out.String(string(in.UrlHash))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardInviteInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardInviteInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardInviteInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardInviteInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(in *jlexer.Lexer, out *BoardInternal) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "boardID":
			out.BoardID = int64(in.Int64())
		case "adminID":
			out.AdminID = int64(in.Int64())
		case "name":
			out.Name = string(in.String())
		case "theme":
			out.Theme = string(in.String())
		case "star":
			out.Star = bool(in.Bool())
		case "cards":
			if in.IsNull() {
				in.Skip()
				out.Cards = nil
			} else {
				in.Delim('[')
				if out.Cards == nil {
					if !in.IsDelim(']') {
						out.Cards = make([]CardInternal, 0, 1)
					} else {
						out.Cards = []CardInternal{}
					}
				} else {
					out.Cards = (out.Cards)[:0]
				}
				for !in.IsDelim(']') {
					var v10 CardInternal
					(v10).UnmarshalEasyJSON(in)
					out.Cards = append(out.Cards, v10)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "usersIDs":
			if in.IsNull() {
				in.Skip()
				out.UsersIDs = nil
			} else {
				in.Delim('[')
				if out.UsersIDs == nil {
					if !in.IsDelim(']') {
						out.UsersIDs = make([]int64, 0, 8)
					} else {
						out.UsersIDs = []int64{}
					}
				} else {
					out.UsersIDs = (out.UsersIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v11 int64
					v11 = int64(in.Int64())
					out.UsersIDs = append(out.UsersIDs, v11)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Tags":
			if in.IsNull() {
				in.Skip()
				out.Tags = nil
			} else {
				in.Delim('[')
				if out.Tags == nil {
					if !in.IsDelim(']') {
						out.Tags = make([]TagOutside, 0, 1)
					} else {
						out.Tags = []TagOutside{}
					}
				} else {
					out.Tags = (out.Tags)[:0]
				}
				for !in.IsDelim(']') {
					var v12 TagOutside
					(v12).UnmarshalEasyJSON(in)
					out.Tags = append(out.Tags, v12)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(out *jwriter.Writer, in BoardInternal) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"boardID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.BoardID))
	}
	{
		const prefix string = ",\"adminID\":"
		out.RawString(prefix)
		out.Int64(int64(in.AdminID))
	}
	{
		const prefix string = ",\"name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"theme\":"
		out.RawString(prefix)
		out.String(string(in.Theme))
	}
	{
		const prefix string = ",\"star\":"
		out.RawString(prefix)
		out.Bool(bool(in.Star))
	}
	{
		const prefix string = ",\"cards\":"
		out.RawString(prefix)
		if in.Cards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v13, v14 := range in.Cards {
				if v13 > 0 {
					out.RawByte(',')
				}
				(v14).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"usersIDs\":"
		out.RawString(prefix)
		if in.UsersIDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v15, v16 := range in.UsersIDs {
				if v15 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v16))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v17, v18 := range in.Tags {
				if v17 > 0 {
					out.RawByte(',')
				}
				(v18).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardInternal) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardInternal) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardInternal) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardInternal) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(in *jlexer.Lexer, out *BoardInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "boardID":
			out.BoardID = int64(in.Int64())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(out *jwriter.Writer, in BoardInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"boardID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.BoardID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(in *jlexer.Lexer, out *BoardChangeInput) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "boardID":
			out.BoardID = int64(in.Int64())
		case "boardName":
			out.BoardName = string(in.String())
		case "boardTheme":
			out.Theme = string(in.String())
		case "boardStar":
			out.Star = bool(in.Bool())
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(out *jwriter.Writer, in BoardChangeInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"boardID\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.Int64(int64(in.BoardID))
	}
	{
		const prefix string = ",\"boardName\":"
		out.RawString(prefix)
		out.String(string(in.BoardName))
	}
	{
		const prefix string = ",\"boardTheme\":"
		out.RawString(prefix)
		out.String(string(in.Theme))
	}
	{
		const prefix string = ",\"boardStar\":"
		out.RawString(prefix)
		out.Bool(bool(in.Star))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardChangeInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardChangeInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardChangeInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardChangeInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(in *jlexer.Lexer, out *Board) {
	isTopLevel := in.IsStart()
	if in.IsNull() {
		if isTopLevel {
			in.Consumed()
		}
		in.Skip()
		return
	}
	in.Delim('{')
	for !in.IsDelim('}') {
		key := in.UnsafeFieldName(false)
		in.WantColon()
		if in.IsNull() {
			in.Skip()
			in.WantComma()
			continue
		}
		switch key {
		case "boardID":
			out.BoardID = int64(in.Int64())
		case "adminID":
			out.AdminID = int64(in.Int64())
		case "boardName":
			out.Name = string(in.String())
		case "theme":
			out.Theme = string(in.String())
		case "star":
			out.Star = bool(in.Bool())
		case "usersIDs":
			if in.IsNull() {
				in.Skip()
				out.UsersIDs = nil
			} else {
				in.Delim('[')
				if out.UsersIDs == nil {
					if !in.IsDelim(']') {
						out.UsersIDs = make([]int64, 0, 8)
					} else {
						out.UsersIDs = []int64{}
					}
				} else {
					out.UsersIDs = (out.UsersIDs)[:0]
				}
				for !in.IsDelim(']') {
					var v19 int64
					v19 = int64(in.Int64())
					out.UsersIDs = append(out.UsersIDs, v19)
					in.WantComma()
				}
				in.Delim(']')
			}
		default:
			in.SkipRecursive()
		}
		in.WantComma()
	}
	in.Delim('}')
	if isTopLevel {
		in.Consumed()
	}
}
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(out *jwriter.Writer, in Board) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"boardID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.BoardID))
	}
	{
		const prefix string = ",\"adminID\":"
		out.RawString(prefix)
		out.Int64(int64(in.AdminID))
	}
	{
		const prefix string = ",\"boardName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"theme\":"
		out.RawString(prefix)
		out.String(string(in.Theme))
	}
	{
		const prefix string = ",\"star\":"
		out.RawString(prefix)
		out.Bool(bool(in.Star))
	}
	{
		const prefix string = ",\"usersIDs\":"
		out.RawString(prefix)
		if in.UsersIDs == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.UsersIDs {
				if v20 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v21))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Board) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Board) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Board) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Board) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(l, v)
}
