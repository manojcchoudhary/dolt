// This file was generated by nomdl/codegen.

package main

import (
	"github.com/attic-labs/noms/ref"
	"github.com/attic-labs/noms/types"
)

var __mainPackageInFile_types_CachedRef = __mainPackageInFile_types_Ref()

// This function builds up a Noms value that describes the type
// package implemented by this file and registers it with the global
// type package definition cache.
func __mainPackageInFile_types_Ref() types.Ref {
	p := types.PackageDef{
		Types: types.MapOfStringToTypeRefDef{

			"Pitch": __typeRefOfPitch(),
		},
	}.New()
	return types.Ref{R: types.RegisterPackage(&p)}
}

// ListOfMapOfStringToValue

type ListOfMapOfStringToValue struct {
	l types.List
}

func NewListOfMapOfStringToValue() ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{types.NewList()}
}

type ListOfMapOfStringToValueDef []MapOfStringToValueDef

func (def ListOfMapOfStringToValueDef) New() ListOfMapOfStringToValue {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New().NomsValue()
	}
	return ListOfMapOfStringToValue{types.NewList(l...)}
}

func (self ListOfMapOfStringToValue) Def() ListOfMapOfStringToValueDef {
	l := make([]MapOfStringToValueDef, self.Len())
	for i := uint64(0); i < self.Len(); i++ {
		l[i] = MapOfStringToValueFromVal(self.l.Get(i)).Def()
	}
	return l
}

func ListOfMapOfStringToValueFromVal(val types.Value) ListOfMapOfStringToValue {
	// TODO: Validate here
	return ListOfMapOfStringToValue{val.(types.List)}
}

func (self ListOfMapOfStringToValue) NomsValue() types.Value {
	return self.l
}

func (l ListOfMapOfStringToValue) Equals(p ListOfMapOfStringToValue) bool {
	return l.l.Equals(p.l)
}

func (l ListOfMapOfStringToValue) Ref() ref.Ref {
	return l.l.Ref()
}

func (l ListOfMapOfStringToValue) Len() uint64 {
	return l.l.Len()
}

func (l ListOfMapOfStringToValue) Empty() bool {
	return l.Len() == uint64(0)
}

func (self ListOfMapOfStringToValue) Get(i uint64) MapOfStringToValue {
	return MapOfStringToValueFromVal(self.l.Get(i))
}

func (l ListOfMapOfStringToValue) Slice(idx uint64, end uint64) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Slice(idx, end)}
}

func (self ListOfMapOfStringToValue) Set(i uint64, val MapOfStringToValue) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{self.l.Set(i, val.NomsValue())}
}

func (l ListOfMapOfStringToValue) Append(v ...MapOfStringToValue) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Append(l.fromElemSlice(v)...)}
}

func (l ListOfMapOfStringToValue) Insert(idx uint64, v ...MapOfStringToValue) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Insert(idx, l.fromElemSlice(v)...)}
}

func (l ListOfMapOfStringToValue) Remove(idx uint64, end uint64) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{l.l.Remove(idx, end)}
}

func (l ListOfMapOfStringToValue) RemoveAt(idx uint64) ListOfMapOfStringToValue {
	return ListOfMapOfStringToValue{(l.l.RemoveAt(idx))}
}

func (l ListOfMapOfStringToValue) fromElemSlice(p []MapOfStringToValue) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

type ListOfMapOfStringToValueIterCallback func(v MapOfStringToValue) (stop bool)

func (l ListOfMapOfStringToValue) Iter(cb ListOfMapOfStringToValueIterCallback) {
	l.l.Iter(func(v types.Value) bool {
		return cb(MapOfStringToValueFromVal(v))
	})
}

type ListOfMapOfStringToValueIterAllCallback func(v MapOfStringToValue)

func (l ListOfMapOfStringToValue) IterAll(cb ListOfMapOfStringToValueIterAllCallback) {
	l.l.IterAll(func(v types.Value) {
		cb(MapOfStringToValueFromVal(v))
	})
}

type ListOfMapOfStringToValueFilterCallback func(v MapOfStringToValue) (keep bool)

func (l ListOfMapOfStringToValue) Filter(cb ListOfMapOfStringToValueFilterCallback) ListOfMapOfStringToValue {
	nl := NewListOfMapOfStringToValue()
	l.IterAll(func(v MapOfStringToValue) {
		if cb(v) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// MapOfStringToValue

type MapOfStringToValue struct {
	m types.Map
}

func NewMapOfStringToValue() MapOfStringToValue {
	return MapOfStringToValue{types.NewMap()}
}

type MapOfStringToValueDef map[string]types.Value

func (def MapOfStringToValueDef) New() MapOfStringToValue {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v)
	}
	return MapOfStringToValue{types.NewMap(kv...)}
}

func (self MapOfStringToValue) Def() MapOfStringToValueDef {
	def := make(map[string]types.Value)
	self.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v
		return false
	})
	return def
}

func MapOfStringToValueFromVal(p types.Value) MapOfStringToValue {
	// TODO: Validate here
	return MapOfStringToValue{p.(types.Map)}
}

func (m MapOfStringToValue) NomsValue() types.Value {
	return m.m
}

func (m MapOfStringToValue) Equals(p MapOfStringToValue) bool {
	return m.m.Equals(p.m)
}

func (m MapOfStringToValue) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToValue) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToValue) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToValue) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToValue) Get(p string) types.Value {
	return m.m.Get(types.NewString(p))
}

func (m MapOfStringToValue) Set(k string, v types.Value) MapOfStringToValue {
	return MapOfStringToValue{m.m.Set(types.NewString(k), v)}
}

// TODO: Implement SetM?

func (m MapOfStringToValue) Remove(p string) MapOfStringToValue {
	return MapOfStringToValue{m.m.Remove(types.NewString(p))}
}

type MapOfStringToValueIterCallback func(k string, v types.Value) (stop bool)

func (m MapOfStringToValue) Iter(cb MapOfStringToValueIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v)
	})
}

type MapOfStringToValueIterAllCallback func(k string, v types.Value)

func (m MapOfStringToValue) IterAll(cb MapOfStringToValueIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v)
	})
}

type MapOfStringToValueFilterCallback func(k string, v types.Value) (keep bool)

func (m MapOfStringToValue) Filter(cb MapOfStringToValueFilterCallback) MapOfStringToValue {
	nm := NewMapOfStringToValue()
	m.IterAll(func(k string, v types.Value) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// MapOfStringToListOfPitch

type MapOfStringToListOfPitch struct {
	m types.Map
}

func NewMapOfStringToListOfPitch() MapOfStringToListOfPitch {
	return MapOfStringToListOfPitch{types.NewMap()}
}

type MapOfStringToListOfPitchDef map[string]ListOfPitchDef

func (def MapOfStringToListOfPitchDef) New() MapOfStringToListOfPitch {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), v.New().NomsValue())
	}
	return MapOfStringToListOfPitch{types.NewMap(kv...)}
}

func (self MapOfStringToListOfPitch) Def() MapOfStringToListOfPitchDef {
	def := make(map[string]ListOfPitchDef)
	self.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = ListOfPitchFromVal(v).Def()
		return false
	})
	return def
}

func MapOfStringToListOfPitchFromVal(p types.Value) MapOfStringToListOfPitch {
	// TODO: Validate here
	return MapOfStringToListOfPitch{p.(types.Map)}
}

func (m MapOfStringToListOfPitch) NomsValue() types.Value {
	return m.m
}

func (m MapOfStringToListOfPitch) Equals(p MapOfStringToListOfPitch) bool {
	return m.m.Equals(p.m)
}

func (m MapOfStringToListOfPitch) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToListOfPitch) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToListOfPitch) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToListOfPitch) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToListOfPitch) Get(p string) ListOfPitch {
	return ListOfPitchFromVal(m.m.Get(types.NewString(p)))
}

func (m MapOfStringToListOfPitch) Set(k string, v ListOfPitch) MapOfStringToListOfPitch {
	return MapOfStringToListOfPitch{m.m.Set(types.NewString(k), v.NomsValue())}
}

// TODO: Implement SetM?

func (m MapOfStringToListOfPitch) Remove(p string) MapOfStringToListOfPitch {
	return MapOfStringToListOfPitch{m.m.Remove(types.NewString(p))}
}

type MapOfStringToListOfPitchIterCallback func(k string, v ListOfPitch) (stop bool)

func (m MapOfStringToListOfPitch) Iter(cb MapOfStringToListOfPitchIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), ListOfPitchFromVal(v))
	})
}

type MapOfStringToListOfPitchIterAllCallback func(k string, v ListOfPitch)

func (m MapOfStringToListOfPitch) IterAll(cb MapOfStringToListOfPitchIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), ListOfPitchFromVal(v))
	})
}

type MapOfStringToListOfPitchFilterCallback func(k string, v ListOfPitch) (keep bool)

func (m MapOfStringToListOfPitch) Filter(cb MapOfStringToListOfPitchFilterCallback) MapOfStringToListOfPitch {
	nm := NewMapOfStringToListOfPitch()
	m.IterAll(func(k string, v ListOfPitch) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}

// ListOfPitch

type ListOfPitch struct {
	l types.List
}

func NewListOfPitch() ListOfPitch {
	return ListOfPitch{types.NewList()}
}

type ListOfPitchDef []PitchDef

func (def ListOfPitchDef) New() ListOfPitch {
	l := make([]types.Value, len(def))
	for i, d := range def {
		l[i] = d.New().NomsValue()
	}
	return ListOfPitch{types.NewList(l...)}
}

func (self ListOfPitch) Def() ListOfPitchDef {
	l := make([]PitchDef, self.Len())
	for i := uint64(0); i < self.Len(); i++ {
		l[i] = PitchFromVal(self.l.Get(i)).Def()
	}
	return l
}

func ListOfPitchFromVal(val types.Value) ListOfPitch {
	// TODO: Validate here
	return ListOfPitch{val.(types.List)}
}

func (self ListOfPitch) NomsValue() types.Value {
	return self.l
}

func (l ListOfPitch) Equals(p ListOfPitch) bool {
	return l.l.Equals(p.l)
}

func (l ListOfPitch) Ref() ref.Ref {
	return l.l.Ref()
}

func (l ListOfPitch) Len() uint64 {
	return l.l.Len()
}

func (l ListOfPitch) Empty() bool {
	return l.Len() == uint64(0)
}

func (self ListOfPitch) Get(i uint64) Pitch {
	return PitchFromVal(self.l.Get(i))
}

func (l ListOfPitch) Slice(idx uint64, end uint64) ListOfPitch {
	return ListOfPitch{l.l.Slice(idx, end)}
}

func (self ListOfPitch) Set(i uint64, val Pitch) ListOfPitch {
	return ListOfPitch{self.l.Set(i, val.NomsValue())}
}

func (l ListOfPitch) Append(v ...Pitch) ListOfPitch {
	return ListOfPitch{l.l.Append(l.fromElemSlice(v)...)}
}

func (l ListOfPitch) Insert(idx uint64, v ...Pitch) ListOfPitch {
	return ListOfPitch{l.l.Insert(idx, l.fromElemSlice(v)...)}
}

func (l ListOfPitch) Remove(idx uint64, end uint64) ListOfPitch {
	return ListOfPitch{l.l.Remove(idx, end)}
}

func (l ListOfPitch) RemoveAt(idx uint64) ListOfPitch {
	return ListOfPitch{(l.l.RemoveAt(idx))}
}

func (l ListOfPitch) fromElemSlice(p []Pitch) []types.Value {
	r := make([]types.Value, len(p))
	for i, v := range p {
		r[i] = v.NomsValue()
	}
	return r
}

type ListOfPitchIterCallback func(v Pitch) (stop bool)

func (l ListOfPitch) Iter(cb ListOfPitchIterCallback) {
	l.l.Iter(func(v types.Value) bool {
		return cb(PitchFromVal(v))
	})
}

type ListOfPitchIterAllCallback func(v Pitch)

func (l ListOfPitch) IterAll(cb ListOfPitchIterAllCallback) {
	l.l.IterAll(func(v types.Value) {
		cb(PitchFromVal(v))
	})
}

type ListOfPitchFilterCallback func(v Pitch) (keep bool)

func (l ListOfPitch) Filter(cb ListOfPitchFilterCallback) ListOfPitch {
	nl := NewListOfPitch()
	l.IterAll(func(v Pitch) {
		if cb(v) {
			nl = nl.Append(v)
		}
	})
	return nl
}

// Pitch

type Pitch struct {
	m types.Map
}

func NewPitch() Pitch {
	return Pitch{types.NewMap(
		types.NewString("$name"), types.NewString("Pitch"),
		types.NewString("$type"), types.MakeTypeRef(types.NewString("Pitch"), __mainPackageInFile_types_CachedRef),
		types.NewString("X"), types.Float64(0),
		types.NewString("Z"), types.Float64(0),
	)}
}

type PitchDef struct {
	X float64
	Z float64
}

func (def PitchDef) New() Pitch {
	return Pitch{
		types.NewMap(
			types.NewString("$name"), types.NewString("Pitch"),
			types.NewString("$type"), types.MakeTypeRef(types.NewString("Pitch"), __mainPackageInFile_types_CachedRef),
			types.NewString("X"), types.Float64(def.X),
			types.NewString("Z"), types.Float64(def.Z),
		)}
}

func (self Pitch) Def() PitchDef {
	return PitchDef{
		float64(self.m.Get(types.NewString("X")).(types.Float64)),
		float64(self.m.Get(types.NewString("Z")).(types.Float64)),
	}
}

// Creates and returns a Noms Value that describes Pitch.
func __typeRefOfPitch() types.TypeRef {
	return types.MakeStructTypeRef(types.NewString("Pitch"),
		types.NewList(
			types.NewString("X"), types.MakePrimitiveTypeRef(types.Float64Kind),
			types.NewString("Z"), types.MakePrimitiveTypeRef(types.Float64Kind),
		),
		nil)

}

func PitchFromVal(val types.Value) Pitch {
	// TODO: Validate here
	return Pitch{val.(types.Map)}
}

func (self Pitch) NomsValue() types.Value {
	return self.m
}

func (self Pitch) Equals(other Pitch) bool {
	return self.m.Equals(other.m)
}

func (self Pitch) Ref() ref.Ref {
	return self.m.Ref()
}

func (self Pitch) Type() types.TypeRef {
	return self.m.Get(types.NewString("$type")).(types.TypeRef)
}

func (self Pitch) X() float64 {
	return float64(self.m.Get(types.NewString("X")).(types.Float64))
}

func (self Pitch) SetX(val float64) Pitch {
	return Pitch{self.m.Set(types.NewString("X"), types.Float64(val))}
}

func (self Pitch) Z() float64 {
	return float64(self.m.Get(types.NewString("Z")).(types.Float64))
}

func (self Pitch) SetZ(val float64) Pitch {
	return Pitch{self.m.Set(types.NewString("Z"), types.Float64(val))}
}

// MapOfStringToString

type MapOfStringToString struct {
	m types.Map
}

func NewMapOfStringToString() MapOfStringToString {
	return MapOfStringToString{types.NewMap()}
}

type MapOfStringToStringDef map[string]string

func (def MapOfStringToStringDef) New() MapOfStringToString {
	kv := make([]types.Value, 0, len(def)*2)
	for k, v := range def {
		kv = append(kv, types.NewString(k), types.NewString(v))
	}
	return MapOfStringToString{types.NewMap(kv...)}
}

func (self MapOfStringToString) Def() MapOfStringToStringDef {
	def := make(map[string]string)
	self.m.Iter(func(k, v types.Value) bool {
		def[k.(types.String).String()] = v.(types.String).String()
		return false
	})
	return def
}

func MapOfStringToStringFromVal(p types.Value) MapOfStringToString {
	// TODO: Validate here
	return MapOfStringToString{p.(types.Map)}
}

func (m MapOfStringToString) NomsValue() types.Value {
	return m.m
}

func (m MapOfStringToString) Equals(p MapOfStringToString) bool {
	return m.m.Equals(p.m)
}

func (m MapOfStringToString) Ref() ref.Ref {
	return m.m.Ref()
}

func (m MapOfStringToString) Empty() bool {
	return m.m.Empty()
}

func (m MapOfStringToString) Len() uint64 {
	return m.m.Len()
}

func (m MapOfStringToString) Has(p string) bool {
	return m.m.Has(types.NewString(p))
}

func (m MapOfStringToString) Get(p string) string {
	return m.m.Get(types.NewString(p)).(types.String).String()
}

func (m MapOfStringToString) Set(k string, v string) MapOfStringToString {
	return MapOfStringToString{m.m.Set(types.NewString(k), types.NewString(v))}
}

// TODO: Implement SetM?

func (m MapOfStringToString) Remove(p string) MapOfStringToString {
	return MapOfStringToString{m.m.Remove(types.NewString(p))}
}

type MapOfStringToStringIterCallback func(k string, v string) (stop bool)

func (m MapOfStringToString) Iter(cb MapOfStringToStringIterCallback) {
	m.m.Iter(func(k, v types.Value) bool {
		return cb(k.(types.String).String(), v.(types.String).String())
	})
}

type MapOfStringToStringIterAllCallback func(k string, v string)

func (m MapOfStringToString) IterAll(cb MapOfStringToStringIterAllCallback) {
	m.m.IterAll(func(k, v types.Value) {
		cb(k.(types.String).String(), v.(types.String).String())
	})
}

type MapOfStringToStringFilterCallback func(k string, v string) (keep bool)

func (m MapOfStringToString) Filter(cb MapOfStringToStringFilterCallback) MapOfStringToString {
	nm := NewMapOfStringToString()
	m.IterAll(func(k string, v string) {
		if cb(k, v) {
			nm = nm.Set(k, v)
		}
	})
	return nm
}
