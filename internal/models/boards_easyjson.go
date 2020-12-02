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
					easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(in, &v2)
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
					easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(in, &v3)
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
				easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(out, v7)
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
				easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(out, v9)
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
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(in *jlexer.Lexer, out *TagOutside) {
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
		case "tagID":
			out.TagID = int64(in.Int64())
		case "tagColor":
			out.Color = string(in.String())
		case "tagName":
			out.Name = string(in.String())
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(out *jwriter.Writer, in TagOutside) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"tagID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.TagID))
	}
	{
		const prefix string = ",\"tagColor\":"
		out.RawString(prefix)
		out.String(string(in.Color))
	}
	{
		const prefix string = ",\"tagName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	out.RawByte('}')
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(in *jlexer.Lexer, out *CardOutside) {
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
		case "cardID":
			out.CardID = int64(in.Int64())
		case "cardName":
			out.Name = string(in.String())
		case "cardOrder":
			out.Order = int64(in.Int64())
		case "cardTasks":
			if in.IsNull() {
				in.Skip()
				out.Tasks = nil
			} else {
				in.Delim('[')
				if out.Tasks == nil {
					if !in.IsDelim(']') {
						out.Tasks = make([]TaskOutsideShort, 0, 0)
					} else {
						out.Tasks = []TaskOutsideShort{}
					}
				} else {
					out.Tasks = (out.Tasks)[:0]
				}
				for !in.IsDelim(']') {
					var v10 TaskOutsideShort
					easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(in, &v10)
					out.Tasks = append(out.Tasks, v10)
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(out *jwriter.Writer, in CardOutside) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"cardID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CardID))
	}
	{
		const prefix string = ",\"cardName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"cardOrder\":"
		out.RawString(prefix)
		out.Int64(int64(in.Order))
	}
	{
		const prefix string = ",\"cardTasks\":"
		out.RawString(prefix)
		if in.Tasks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v11, v12 := range in.Tasks {
				if v11 > 0 {
					out.RawByte(',')
				}
				easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(out, v12)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(in *jlexer.Lexer, out *TaskOutsideShort) {
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
		case "taskID":
			out.TaskID = int64(in.Int64())
		case "taskName":
			out.Name = string(in.String())
		case "taskDescription":
			out.Description = string(in.String())
		case "taskOrder":
			out.Order = int64(in.Int64())
		case "taskTags":
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
					var v13 TagOutside
					easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(in, &v13)
					out.Tags = append(out.Tags, v13)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "taskAssigners":
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
					var v14 UserOutsideShort
					(v14).UnmarshalEasyJSON(in)
					out.Users = append(out.Users, v14)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "taskChecklists":
			if in.IsNull() {
				in.Skip()
				out.Checklists = nil
			} else {
				in.Delim('[')
				if out.Checklists == nil {
					if !in.IsDelim(']') {
						out.Checklists = make([]ChecklistOutside, 0, 1)
					} else {
						out.Checklists = []ChecklistOutside{}
					}
				} else {
					out.Checklists = (out.Checklists)[:0]
				}
				for !in.IsDelim(']') {
					var v15 ChecklistOutside
					easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(in, &v15)
					out.Checklists = append(out.Checklists, v15)
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(out *jwriter.Writer, in TaskOutsideShort) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"taskID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.TaskID))
	}
	{
		const prefix string = ",\"taskName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"taskDescription\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"taskOrder\":"
		out.RawString(prefix)
		out.Int64(int64(in.Order))
	}
	{
		const prefix string = ",\"taskTags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v16, v17 := range in.Tags {
				if v16 > 0 {
					out.RawByte(',')
				}
				easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(out, v17)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"taskAssigners\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v18, v19 := range in.Users {
				if v18 > 0 {
					out.RawByte(',')
				}
				(v19).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"taskChecklists\":"
		out.RawString(prefix)
		if in.Checklists == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v20, v21 := range in.Checklists {
				if v20 > 0 {
					out.RawByte(',')
				}
				easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(out, v21)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(in *jlexer.Lexer, out *ChecklistOutside) {
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
		case "checklistID":
			out.ChecklistID = int64(in.Int64())
		case "checklistName":
			out.Name = string(in.String())
		case "checklistItems":
			if data := in.Raw(); in.Ok() {
				in.AddError((out.Items).UnmarshalJSON(data))
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(out *jwriter.Writer, in ChecklistOutside) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"checklistID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ChecklistID))
	}
	{
		const prefix string = ",\"checklistName\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"checklistItems\":"
		out.RawString(prefix)
		out.Raw((in.Items).MarshalJSON())
	}
	out.RawByte('}')
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(in *jlexer.Lexer, out *BoardMemberInput) {
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(out *jwriter.Writer, in BoardMemberInput) {
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
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardMemberInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardMemberInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardMemberInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(in *jlexer.Lexer, out *BoardMember) {
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(out *jwriter.Writer, in BoardMember) {
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
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardMember) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardMember) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardMember) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(in *jlexer.Lexer, out *BoardInternal) {
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
					var v22 CardInternal
					easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(in, &v22)
					out.Cards = append(out.Cards, v22)
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
					var v23 int64
					v23 = int64(in.Int64())
					out.UsersIDs = append(out.UsersIDs, v23)
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
					var v24 TagOutside
					easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(in, &v24)
					out.Tags = append(out.Tags, v24)
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(out *jwriter.Writer, in BoardInternal) {
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
			for v25, v26 := range in.Cards {
				if v25 > 0 {
					out.RawByte(',')
				}
				easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(out, v26)
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
			for v27, v28 := range in.UsersIDs {
				if v27 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v28))
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
			for v29, v30 := range in.Tags {
				if v29 > 0 {
					out.RawByte(',')
				}
				easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(out, v30)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v BoardInternal) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardInternal) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardInternal) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardInternal) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(in *jlexer.Lexer, out *CardInternal) {
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
		case "CardID":
			out.CardID = int64(in.Int64())
		case "Name":
			out.Name = string(in.String())
		case "Order":
			out.Order = int64(in.Int64())
		case "Tasks":
			if in.IsNull() {
				in.Skip()
				out.Tasks = nil
			} else {
				in.Delim('[')
				if out.Tasks == nil {
					if !in.IsDelim(']') {
						out.Tasks = make([]TaskInternalShort, 0, 0)
					} else {
						out.Tasks = []TaskInternalShort{}
					}
				} else {
					out.Tasks = (out.Tasks)[:0]
				}
				for !in.IsDelim(']') {
					var v31 TaskInternalShort
					easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels10(in, &v31)
					out.Tasks = append(out.Tasks, v31)
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(out *jwriter.Writer, in CardInternal) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"CardID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.CardID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Order\":"
		out.RawString(prefix)
		out.Int64(int64(in.Order))
	}
	{
		const prefix string = ",\"Tasks\":"
		out.RawString(prefix)
		if in.Tasks == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v32, v33 := range in.Tasks {
				if v32 > 0 {
					out.RawByte(',')
				}
				easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels10(out, v33)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels10(in *jlexer.Lexer, out *TaskInternalShort) {
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
		case "TaskID":
			out.TaskID = int64(in.Int64())
		case "Name":
			out.Name = string(in.String())
		case "Description":
			out.Description = string(in.String())
		case "Order":
			out.Order = int64(in.Int64())
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
					var v34 TagOutside
					easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(in, &v34)
					out.Tags = append(out.Tags, v34)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Users":
			if in.IsNull() {
				in.Skip()
				out.Users = nil
			} else {
				in.Delim('[')
				if out.Users == nil {
					if !in.IsDelim(']') {
						out.Users = make([]int64, 0, 8)
					} else {
						out.Users = []int64{}
					}
				} else {
					out.Users = (out.Users)[:0]
				}
				for !in.IsDelim(']') {
					var v35 int64
					v35 = int64(in.Int64())
					out.Users = append(out.Users, v35)
					in.WantComma()
				}
				in.Delim(']')
			}
		case "Checklists":
			if in.IsNull() {
				in.Skip()
				out.Checklists = nil
			} else {
				in.Delim('[')
				if out.Checklists == nil {
					if !in.IsDelim(']') {
						out.Checklists = make([]ChecklistOutside, 0, 1)
					} else {
						out.Checklists = []ChecklistOutside{}
					}
				} else {
					out.Checklists = (out.Checklists)[:0]
				}
				for !in.IsDelim(']') {
					var v36 ChecklistOutside
					easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(in, &v36)
					out.Checklists = append(out.Checklists, v36)
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels10(out *jwriter.Writer, in TaskInternalShort) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"TaskID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.TaskID))
	}
	{
		const prefix string = ",\"Name\":"
		out.RawString(prefix)
		out.String(string(in.Name))
	}
	{
		const prefix string = ",\"Description\":"
		out.RawString(prefix)
		out.String(string(in.Description))
	}
	{
		const prefix string = ",\"Order\":"
		out.RawString(prefix)
		out.Int64(int64(in.Order))
	}
	{
		const prefix string = ",\"Tags\":"
		out.RawString(prefix)
		if in.Tags == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v37, v38 := range in.Tags {
				if v37 > 0 {
					out.RawByte(',')
				}
				easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(out, v38)
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Users\":"
		out.RawString(prefix)
		if in.Users == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v39, v40 := range in.Users {
				if v39 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v40))
			}
			out.RawByte(']')
		}
	}
	{
		const prefix string = ",\"Checklists\":"
		out.RawString(prefix)
		if in.Checklists == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v41, v42 := range in.Checklists {
				if v41 > 0 {
					out.RawByte(',')
				}
				easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(out, v42)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(in *jlexer.Lexer, out *BoardInput) {
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(out *jwriter.Writer, in BoardInput) {
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
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels12(in *jlexer.Lexer, out *BoardChangeInput) {
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels12(out *jwriter.Writer, in BoardChangeInput) {
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
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels12(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v BoardChangeInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels12(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *BoardChangeInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels12(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *BoardChangeInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels12(l, v)
}
func easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels13(in *jlexer.Lexer, out *Board) {
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
					var v43 int64
					v43 = int64(in.Int64())
					out.UsersIDs = append(out.UsersIDs, v43)
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
func easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels13(out *jwriter.Writer, in Board) {
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
			for v44, v45 := range in.UsersIDs {
				if v44 > 0 {
					out.RawByte(',')
				}
				out.Int64(int64(v45))
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v Board) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels13(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v Board) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson74fcd6ebEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels13(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *Board) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels13(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *Board) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson74fcd6ebDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels13(l, v)
}
