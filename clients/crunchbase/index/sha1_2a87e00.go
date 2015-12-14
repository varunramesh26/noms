// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/chunks"
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_sha1_2a87e00_CachedRef ref.Ref

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func init() {
	p := types.NewPackage([]types.Type{
		types.MakeStructType("Import",
			[]types.Field{
				types.Field{"FileSHA1", types.MakePrimitiveType(types.StringKind), false},
				types.Field{"Date", types.MakeType(ref.Ref{}, 1), false},
				types.Field{"Companies", types.MakeCompoundType(types.RefKind, types.MakeCompoundType(types.MapKind, types.MakePrimitiveType(types.StringKind), types.MakeCompoundType(types.RefKind, types.MakeType(ref.Parse("sha1-9fb26a609a1d4134935888c9676c79d6ea89acbb"), 0)))), false},
			},
			types.Choices{},
		),
		types.MakeStructType("Date",
			[]types.Field{
				types.Field{"RFC3339", types.MakePrimitiveType(types.StringKind), false},
			},
			types.Choices{},
		),
	}, []ref.Ref{
		ref.Parse("sha1-9fb26a609a1d4134935888c9676c79d6ea89acbb"),
	})
	__mainPackageInFile_sha1_2a87e00_CachedRef = types.RegisterPackage(&p)
}

// Import

type Import struct {
	_FileSHA1  string
	_Date      Date
	_Companies RefOfMapOfStringToRefOfCompany

	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewImport(cs chunks.ChunkStore) Import {
	return Import{
		_FileSHA1:  "",
		_Date:      NewDate(cs),
		_Companies: NewRefOfMapOfStringToRefOfCompany(ref.Ref{}),

		cs:  cs,
		ref: &ref.Ref{},
	}
}

type ImportDef struct {
	FileSHA1  string
	Date      DateDef
	Companies ref.Ref
}

func (def ImportDef) New(cs chunks.ChunkStore) Import {
	return Import{
		_FileSHA1:  def.FileSHA1,
		_Date:      def.Date.New(cs),
		_Companies: NewRefOfMapOfStringToRefOfCompany(def.Companies),
		cs:         cs,
		ref:        &ref.Ref{},
	}
}

func (s Import) Def() (d ImportDef) {
	d.FileSHA1 = s._FileSHA1
	d.Date = s._Date.Def()
	d.Companies = s._Companies.TargetRef()
	return
}

var __typeForImport types.Type

func (m Import) Type() types.Type {
	return __typeForImport
}

func init() {
	__typeForImport = types.MakeType(__mainPackageInFile_sha1_2a87e00_CachedRef, 0)
	types.RegisterStruct(__typeForImport, builderForImport, readerForImport)
}

func builderForImport(cs chunks.ChunkStore, values []types.Value) types.Value {
	i := 0
	s := Import{ref: &ref.Ref{}, cs: cs}
	s._FileSHA1 = values[i].(types.String).String()
	i++
	s._Date = values[i].(Date)
	i++
	s._Companies = values[i].(RefOfMapOfStringToRefOfCompany)
	i++
	return s
}

func readerForImport(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Import)
	values = append(values, types.NewString(s._FileSHA1))
	values = append(values, s._Date)
	values = append(values, s._Companies)
	return values
}

func (s Import) Equals(other types.Value) bool {
	return other != nil && __typeForImport.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Import) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Import) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForImport.Chunks()...)
	chunks = append(chunks, s._Date.Chunks()...)
	chunks = append(chunks, s._Companies.Chunks()...)
	return
}

func (s Import) ChildValues() (ret []types.Value) {
	ret = append(ret, types.NewString(s._FileSHA1))
	ret = append(ret, s._Date)
	ret = append(ret, s._Companies)
	return
}

func (s Import) FileSHA1() string {
	return s._FileSHA1
}

func (s Import) SetFileSHA1(val string) Import {
	s._FileSHA1 = val
	s.ref = &ref.Ref{}
	return s
}

func (s Import) Date() Date {
	return s._Date
}

func (s Import) SetDate(val Date) Import {
	s._Date = val
	s.ref = &ref.Ref{}
	return s
}

func (s Import) Companies() RefOfMapOfStringToRefOfCompany {
	return s._Companies
}

func (s Import) SetCompanies(val RefOfMapOfStringToRefOfCompany) Import {
	s._Companies = val
	s.ref = &ref.Ref{}
	return s
}

// Date

type Date struct {
	_RFC3339 string

	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewDate(cs chunks.ChunkStore) Date {
	return Date{
		_RFC3339: "",

		cs:  cs,
		ref: &ref.Ref{},
	}
}

type DateDef struct {
	RFC3339 string
}

func (def DateDef) New(cs chunks.ChunkStore) Date {
	return Date{
		_RFC3339: def.RFC3339,
		cs:       cs,
		ref:      &ref.Ref{},
	}
}

func (s Date) Def() (d DateDef) {
	d.RFC3339 = s._RFC3339
	return
}

var __typeForDate types.Type

func (m Date) Type() types.Type {
	return __typeForDate
}

func init() {
	__typeForDate = types.MakeType(__mainPackageInFile_sha1_2a87e00_CachedRef, 1)
	types.RegisterStruct(__typeForDate, builderForDate, readerForDate)
}

func builderForDate(cs chunks.ChunkStore, values []types.Value) types.Value {
	i := 0
	s := Date{ref: &ref.Ref{}, cs: cs}
	s._RFC3339 = values[i].(types.String).String()
	i++
	return s
}

func readerForDate(v types.Value) []types.Value {
	values := []types.Value{}
	s := v.(Date)
	values = append(values, types.NewString(s._RFC3339))
	return values
}

func (s Date) Equals(other types.Value) bool {
	return other != nil && __typeForDate.Equals(other.Type()) && s.Ref() == other.Ref()
}

func (s Date) Ref() ref.Ref {
	return types.EnsureRef(s.ref, s)
}

func (s Date) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, __typeForDate.Chunks()...)
	return
}

func (s Date) ChildValues() (ret []types.Value) {
	ret = append(ret, types.NewString(s._RFC3339))
	return
}

func (s Date) RFC3339() string {
	return s._RFC3339
}

func (s Date) SetRFC3339(val string) Date {
	s._RFC3339 = val
	s.ref = &ref.Ref{}
	return s
}

// RefOfMapOfStringToRefOfCompany

type RefOfMapOfStringToRefOfCompany struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfMapOfStringToRefOfCompany(target ref.Ref) RefOfMapOfStringToRefOfCompany {
	return RefOfMapOfStringToRefOfCompany{target, &ref.Ref{}}
}

func (r RefOfMapOfStringToRefOfCompany) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfMapOfStringToRefOfCompany) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfMapOfStringToRefOfCompany) Equals(other types.Value) bool {
	return other != nil && __typeForRefOfMapOfStringToRefOfCompany.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfMapOfStringToRefOfCompany) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfMapOfStringToRefOfCompany) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfMapOfStringToRefOfCompany.
var __typeForRefOfMapOfStringToRefOfCompany types.Type

func (r RefOfMapOfStringToRefOfCompany) Type() types.Type {
	return __typeForRefOfMapOfStringToRefOfCompany
}

func (r RefOfMapOfStringToRefOfCompany) Less(other types.OrderedValue) bool {
	return r.TargetRef().Less(other.(types.RefBase).TargetRef())
}

func init() {
	__typeForRefOfMapOfStringToRefOfCompany = types.MakeCompoundType(types.RefKind, types.MakeCompoundType(types.MapKind, types.MakePrimitiveType(types.StringKind), types.MakeCompoundType(types.RefKind, types.MakeType(ref.Parse("sha1-9fb26a609a1d4134935888c9676c79d6ea89acbb"), 0))))
	types.RegisterRef(__typeForRefOfMapOfStringToRefOfCompany, builderForRefOfMapOfStringToRefOfCompany)
}

func builderForRefOfMapOfStringToRefOfCompany(r ref.Ref) types.Value {
	return NewRefOfMapOfStringToRefOfCompany(r)
}

func (r RefOfMapOfStringToRefOfCompany) TargetValue(cs chunks.ChunkStore) MapOfStringToRefOfCompany {
	return types.ReadValue(r.target, cs).(MapOfStringToRefOfCompany)
}

func (r RefOfMapOfStringToRefOfCompany) SetTargetValue(val MapOfStringToRefOfCompany, cs chunks.ChunkSink) RefOfMapOfStringToRefOfCompany {
	return NewRefOfMapOfStringToRefOfCompany(types.WriteValue(val, cs))
}

// MapOfStringToRefOfCompany

type MapOfStringToRefOfCompany struct {
	m   types.Map
	cs  chunks.ChunkStore
	ref *ref.Ref
}

func NewMapOfStringToRefOfCompany(cs chunks.ChunkStore) MapOfStringToRefOfCompany {
	return MapOfStringToRefOfCompany{types.NewTypedMap(cs, __typeForMapOfStringToRefOfCompany), cs, &ref.Ref{}}
}

type MapOfStringToRefOfCompanyDef map[string]ref.Ref

func (def MapOfStringToRefOfCompanyDef) New(cs chunks.ChunkStore) MapOfStringToRefOfCompany {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), NewRefOfCompany(v))
	}
	return MapOfStringToRefOfCompany{types.NewTypedMap(cs, __typeForMapOfStringToRefOfCompany, kv...), cs, &ref.Ref{}}
}

func (m MapOfStringToRefOfCompany) Def() MapOfStringToRefOfCompanyDef {
	def := make(map[string]ref.Ref)
	m.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(RefOfCompany).TargetRef()
		return false
	})
	return def
}

func (m MapOfStringToRefOfCompany) Equals(other types.Value) bool {
	return other != nil && __typeForMapOfStringToRefOfCompany.Equals(other.Type()) && m.Ref() == other.Ref()
}

func (m MapOfStringToRefOfCompany) Ref() ref.Ref {
	return types.EnsureRef(m.ref, m)
}

func (m MapOfStringToRefOfCompany) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, m.Type().Chunks()...)
	chunks = append(chunks, m.m.Chunks()...)
	return
}

func (m MapOfStringToRefOfCompany) ChildValues() []types.Value {
	return append([]types.Value{}, m.m.ChildValues()...)
}

// A Noms Value that describes MapOfStringToRefOfCompany.
var __typeForMapOfStringToRefOfCompany types.Type

func (m MapOfStringToRefOfCompany) Type() types.Type {
	return __typeForMapOfStringToRefOfCompany
}

func init() {
	__typeForMapOfStringToRefOfCompany = types.MakeCompoundType(types.MapKind, types.MakePrimitiveType(types.StringKind), types.MakeCompoundType(types.RefKind, types.MakeType(ref.Parse("sha1-9fb26a609a1d4134935888c9676c79d6ea89acbb"), 0)))
	types.RegisterValue(__typeForMapOfStringToRefOfCompany, builderForMapOfStringToRefOfCompany, readerForMapOfStringToRefOfCompany)
}

func builderForMapOfStringToRefOfCompany(cs chunks.ChunkStore, v types.Value) types.Value {
	return MapOfStringToRefOfCompany{v.(types.Map), cs, &ref.Ref{}}
}

func readerForMapOfStringToRefOfCompany(v types.Value) types.Value {
	return v.(MapOfStringToRefOfCompany).m
}

func (m MapOfStringToRefOfCompany) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToRefOfCompany) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToRefOfCompany) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToRefOfCompany) Get(p string) RefOfCompany {
	return m.m.Get(types.NewString(p)).(RefOfCompany)
}

func (m MapOfStringToRefOfCompany) MaybeGet(p string) (RefOfCompany, bool) {
	v, ok := m.m.MaybeGet(types.NewString(p))
	if !ok {
		return NewRefOfCompany(ref.Ref{}), false
	}
	return v.(RefOfCompany), ok
}

func (m MapOfStringToRefOfCompany) Set(k string, v RefOfCompany) MapOfStringToRefOfCompany {
	return MapOfStringToRefOfCompany{m.m.Set(types.NewString(k), v), m.cs, &ref.Ref{}}
}

// TODO: Implement SetM?

func (m MapOfStringToRefOfCompany) Remove(p string) MapOfStringToRefOfCompany {
	return MapOfStringToRefOfCompany{m.m.Remove(types.NewString(p)), m.cs, &ref.Ref{}}
}

type MapOfStringToRefOfCompanyIterCallback func(k string, v RefOfCompany) (stop bool)

func (m MapOfStringToRefOfCompany) Iter(cb MapOfStringToRefOfCompanyIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(RefOfCompany))
	})
}

type MapOfStringToRefOfCompanyIterAllCallback func(k string, v RefOfCompany)

func (m MapOfStringToRefOfCompany) IterAll(cb MapOfStringToRefOfCompanyIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(RefOfCompany))
	})
}

func (m MapOfStringToRefOfCompany) IterAllP(concurrency int, cb MapOfStringToRefOfCompanyIterAllCallback) {
	m.m.IterAllP(concurrency, func(k, v types.Value) {
		cb(k.(types.String).String(), v.(RefOfCompany))
	})
}

type MapOfStringToRefOfCompanyFilterCallback func(k string, v RefOfCompany) (keep bool)

func (m MapOfStringToRefOfCompany) Filter(cb MapOfStringToRefOfCompanyFilterCallback) MapOfStringToRefOfCompany {
	out := m.m.Filter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(RefOfCompany))
	})
	return MapOfStringToRefOfCompany{out, m.cs, &ref.Ref{}}
}

// RefOfCompany

type RefOfCompany struct {
	target ref.Ref
	ref    *ref.Ref
}

func NewRefOfCompany(target ref.Ref) RefOfCompany {
	return RefOfCompany{target, &ref.Ref{}}
}

func (r RefOfCompany) TargetRef() ref.Ref {
	return r.target
}

func (r RefOfCompany) Ref() ref.Ref {
	return types.EnsureRef(r.ref, r)
}

func (r RefOfCompany) Equals(other types.Value) bool {
	return other != nil && __typeForRefOfCompany.Equals(other.Type()) && r.Ref() == other.Ref()
}

func (r RefOfCompany) Chunks() (chunks []ref.Ref) {
	chunks = append(chunks, r.Type().Chunks()...)
	chunks = append(chunks, r.target)
	return
}

func (r RefOfCompany) ChildValues() []types.Value {
	return nil
}

// A Noms Value that describes RefOfCompany.
var __typeForRefOfCompany types.Type

func (r RefOfCompany) Type() types.Type {
	return __typeForRefOfCompany
}

func (r RefOfCompany) Less(other types.OrderedValue) bool {
	return r.TargetRef().Less(other.(types.RefBase).TargetRef())
}

func init() {
	__typeForRefOfCompany = types.MakeCompoundType(types.RefKind, types.MakeType(ref.Parse("sha1-9fb26a609a1d4134935888c9676c79d6ea89acbb"), 0))
	types.RegisterRef(__typeForRefOfCompany, builderForRefOfCompany)
}

func builderForRefOfCompany(r ref.Ref) types.Value {
	return NewRefOfCompany(r)
}

func (r RefOfCompany) TargetValue(cs chunks.ChunkStore) Company {
	return types.ReadValue(r.target, cs).(Company)
}

func (r RefOfCompany) SetTargetValue(val Company, cs chunks.ChunkSink) RefOfCompany {
	return NewRefOfCompany(types.WriteValue(val, cs))
}
