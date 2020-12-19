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

func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(in *jlexer.Lexer, out *UserSession) {
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
		case "SessionID":
			out.SessionID = string(in.String())
		case "UserID":
			out.UserID = int64(in.Int64())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(out *jwriter.Writer, in UserSession) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"SessionID\":"
		out.RawString(prefix[1:])
		out.String(string(in.SessionID))
	}
	{
		const prefix string = ",\"UserID\":"
		out.RawString(prefix)
		out.Int64(int64(in.UserID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserSession) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserSession) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserSession) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserSession) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(in *jlexer.Lexer, out *UserOutsideShort) {
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(out *jwriter.Writer, in UserOutsideShort) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
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
func (v UserOutsideShort) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserOutsideShort) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserOutsideShort) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserOutsideShort) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels1(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(in *jlexer.Lexer, out *UserOutside) {
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(out *jwriter.Writer, in UserOutside) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
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
func (v UserOutside) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserOutside) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserOutside) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserOutside) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels2(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(in *jlexer.Lexer, out *UserInternal) {
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(out *jwriter.Writer, in UserInternal) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
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
func (v UserInternal) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInternal) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInternal) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInternal) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels3(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(in *jlexer.Lexer, out *UserInputReg) {
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
		case "email":
			out.Email = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(out *jwriter.Writer, in UserInputReg) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"username\":"
		out.RawString(prefix)
		out.String(string(in.Username))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInputReg) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInputReg) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInputReg) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInputReg) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels4(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(in *jlexer.Lexer, out *UserInputProfile) {
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
		case "email":
			out.Email = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "fullName":
			out.FullName = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(out *jwriter.Writer, in UserInputProfile) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
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
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInputProfile) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInputProfile) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInputProfile) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInputProfile) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels5(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(in *jlexer.Lexer, out *UserInputPassword) {
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
		case "oldpassword":
			out.OldPassword = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(out *jwriter.Writer, in UserInputPassword) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"oldpassword\":"
		if first {
			first = false
			out.RawString(prefix[1:])
		} else {
			out.RawString(prefix)
		}
		out.String(string(in.OldPassword))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInputPassword) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInputPassword) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInputPassword) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInputPassword) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels6(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(in *jlexer.Lexer, out *UserInputLogin) {
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
		case "email":
			out.Email = string(in.String())
		case "password":
			out.Password = string(in.String())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(out *jwriter.Writer, in UserInputLogin) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
		out.String(string(in.Email))
	}
	{
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInputLogin) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInputLogin) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInputLogin) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInputLogin) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels7(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(in *jlexer.Lexer, out *UserInput) {
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
		case "id":
			out.ID = int64(in.Int64())
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(out *jwriter.Writer, in UserInput) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserInput) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserInput) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserInput) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserInput) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels8(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(in *jlexer.Lexer, out *UserBoardsOutside) {
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
		case "email":
			out.Email = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "fullName":
			out.FullName = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		case "boards":
			if in.IsNull() {
				in.Skip()
				out.Boards = nil
			} else {
				in.Delim('[')
				if out.Boards == nil {
					if !in.IsDelim(']') {
						out.Boards = make([]BoardOutsideShort, 0, 1)
					} else {
						out.Boards = []BoardOutsideShort{}
					}
				} else {
					out.Boards = (out.Boards)[:0]
				}
				for !in.IsDelim(']') {
					var v1 BoardOutsideShort
					(v1).UnmarshalEasyJSON(in)
					out.Boards = append(out.Boards, v1)
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(out *jwriter.Writer, in UserBoardsOutside) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"email\":"
		out.RawString(prefix[1:])
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
	{
		const prefix string = ",\"boards\":"
		out.RawString(prefix)
		if in.Boards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v2, v3 := range in.Boards {
				if v2 > 0 {
					out.RawByte(',')
				}
				(v3).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserBoardsOutside) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserBoardsOutside) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserBoardsOutside) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserBoardsOutside) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels9(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels10(in *jlexer.Lexer, out *UserAvatar) {
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
		case "ID":
			out.ID = int64(in.Int64())
		case "Avatar":
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels10(out *jwriter.Writer, in UserAvatar) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"ID\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
	}
	{
		const prefix string = ",\"Avatar\":"
		out.RawString(prefix)
		out.String(string(in.Avatar))
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v UserAvatar) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels10(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v UserAvatar) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels10(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *UserAvatar) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels10(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *UserAvatar) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels10(l, v)
}
func easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(in *jlexer.Lexer, out *User) {
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
		case "id":
			out.ID = int64(in.Int64())
		case "email":
			out.Email = string(in.String())
		case "username":
			out.Username = string(in.String())
		case "password":
			out.Password = string(in.String())
		case "fullName":
			out.FullName = string(in.String())
		case "avatar":
			out.Avatar = string(in.String())
		case "Boards":
			if in.IsNull() {
				in.Skip()
				out.Boards = nil
			} else {
				in.Delim('[')
				if out.Boards == nil {
					if !in.IsDelim(']') {
						out.Boards = make([]BoardInternal, 0, 0)
					} else {
						out.Boards = []BoardInternal{}
					}
				} else {
					out.Boards = (out.Boards)[:0]
				}
				for !in.IsDelim(']') {
					var v4 BoardInternal
					(v4).UnmarshalEasyJSON(in)
					out.Boards = append(out.Boards, v4)
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
func easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(out *jwriter.Writer, in User) {
	out.RawByte('{')
	first := true
	_ = first
	{
		const prefix string = ",\"id\":"
		out.RawString(prefix[1:])
		out.Int64(int64(in.ID))
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
		const prefix string = ",\"password\":"
		out.RawString(prefix)
		out.String(string(in.Password))
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
	{
		const prefix string = ",\"Boards\":"
		out.RawString(prefix)
		if in.Boards == nil && (out.Flags&jwriter.NilSliceAsEmpty) == 0 {
			out.RawString("null")
		} else {
			out.RawByte('[')
			for v5, v6 := range in.Boards {
				if v5 > 0 {
					out.RawByte(',')
				}
				(v6).MarshalEasyJSON(out)
			}
			out.RawByte(']')
		}
	}
	out.RawByte('}')
}

// MarshalJSON supports json.Marshaler interface
func (v User) MarshalJSON() ([]byte, error) {
	w := jwriter.Writer{}
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(&w, v)
	return w.Buffer.BuildBytes(), w.Error
}

// MarshalEasyJSON supports easyjson.Marshaler interface
func (v User) MarshalEasyJSON(w *jwriter.Writer) {
	easyjson9e1087fdEncodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(w, v)
}

// UnmarshalJSON supports json.Unmarshaler interface
func (v *User) UnmarshalJSON(data []byte) error {
	r := jlexer.Lexer{Data: data}
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(&r, v)
	return r.Error()
}

// UnmarshalEasyJSON supports easyjson.Unmarshaler interface
func (v *User) UnmarshalEasyJSON(l *jlexer.Lexer) {
	easyjson9e1087fdDecodeGithubComGoParkMailRu20202ExtraSafeInternalModels11(l, v)
}
