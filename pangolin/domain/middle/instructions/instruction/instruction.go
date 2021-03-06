package instruction

import (
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/condition"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/match"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/remaining"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/stackframe"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/standard"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/token"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/transform"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/value"
	"github.com/deepvalue-network/software/pangolin/domain/middle/instructions/instruction/variablename"
	"github.com/deepvalue-network/software/pangolin/domain/middle/variables/variable"
)

type instruction struct {
	stackframe   stackframe.Stackframe
	transform    transform.Transform
	variableName variablename.VariableName
	condition    condition.Condition
	standard     standard.Standard
	remaining    remaining.Remaining
	value        value.Value
	insert       variable.Variable
	save         variable.Variable
	del          string
	match        match.Match
	token        token.Token
}

func createInstructionWithStackframe(stackframe stackframe.Stackframe) Instruction {
	return createInstructionInternally(stackframe, nil, nil, nil, nil, nil, nil, nil, nil, "", nil, nil)
}

func createInstructionWithTransform(transform transform.Transform) Instruction {
	return createInstructionInternally(nil, transform, nil, nil, nil, nil, nil, nil, nil, "", nil, nil)
}

func createInstructionWithVariableName(variableName variablename.VariableName) Instruction {
	return createInstructionInternally(nil, nil, variableName, nil, nil, nil, nil, nil, nil, "", nil, nil)
}

func createInstructionWithCondition(condition condition.Condition) Instruction {
	return createInstructionInternally(nil, nil, nil, condition, nil, nil, nil, nil, nil, "", nil, nil)
}

func createInstructionWithStandard(standard standard.Standard) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, standard, nil, nil, nil, nil, "", nil, nil)
}

func createInstructionWithRemaining(remaining remaining.Remaining) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, remaining, nil, nil, nil, "", nil, nil)
}

func createInstructionWithValue(value value.Value) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, value, nil, nil, "", nil, nil)
}

func createInstructionWithInsert(insert variable.Variable) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, insert, nil, "", nil, nil)
}

func createInstructionWithSave(save variable.Variable) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, nil, save, "", nil, nil)
}

func createInstructionWithDelete(del string) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, del, nil, nil)
}

func createInstructionWithMatch(match match.Match) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, "", match, nil)
}

func createInstructionWithToken(token token.Token) Instruction {
	return createInstructionInternally(nil, nil, nil, nil, nil, nil, nil, nil, nil, "", nil, token)
}

func createInstructionInternally(
	stackframe stackframe.Stackframe,
	transform transform.Transform,
	variableName variablename.VariableName,
	condition condition.Condition,
	standard standard.Standard,
	remaining remaining.Remaining,
	value value.Value,
	insert variable.Variable,
	save variable.Variable,
	del string,
	match match.Match,
	token token.Token,
) Instruction {
	out := instruction{
		stackframe:   stackframe,
		transform:    transform,
		variableName: variableName,
		condition:    condition,
		standard:     standard,
		remaining:    remaining,
		value:        value,
		insert:       insert,
		save:         save,
		del:          del,
		match:        match,
		token:        token,
	}

	return &out
}

// IsStackframe returns true if there is a stackframe, false otherwise
func (obj *instruction) IsStackframe() bool {
	return obj.stackframe != nil
}

// Stackframe returns the stackframe, if any
func (obj *instruction) Stackframe() stackframe.Stackframe {
	return obj.stackframe
}

// IsTransform returns true if there is a transform, false otherwise
func (obj *instruction) IsTransform() bool {
	return obj.transform != nil
}

// Transform returns the transform, if any
func (obj *instruction) Transform() transform.Transform {
	return obj.transform
}

// IsVariableName returns true if there is a variableName, false otherwise
func (obj *instruction) IsVariableName() bool {
	return obj.variableName != nil
}

// VariableName returns the variableName, if any
func (obj *instruction) VariableName() variablename.VariableName {
	return obj.variableName
}

// IsCondition returns true if there is a condition, false otherwise
func (obj *instruction) IsCondition() bool {
	return obj.condition != nil
}

// Condition returns the condition, if any
func (obj *instruction) Condition() condition.Condition {
	return obj.condition
}

// IsStandard returns true if there is a standard, false otherwise
func (obj *instruction) IsStandard() bool {
	return obj.standard != nil
}

// Standard returns the standard, if any
func (obj *instruction) Standard() standard.Standard {
	return obj.standard
}

// IsRemaining returns true if there is a remaining, false otherwise
func (obj *instruction) IsRemaining() bool {
	return obj.remaining != nil
}

// Remaining returns the remaining, if any
func (obj *instruction) Remaining() remaining.Remaining {
	return obj.remaining
}

// IsValue returns true if there is a value, false otherwise
func (obj *instruction) IsValue() bool {
	return obj.value != nil
}

// Value returns the value, if any
func (obj *instruction) Value() value.Value {
	return obj.value
}

// IsInsert returns true if there is an insert, false otherwise
func (obj *instruction) IsInsert() bool {
	return obj.insert != nil
}

// Insert returns the insert, if any
func (obj *instruction) Insert() variable.Variable {
	return obj.insert
}

// IsSave returns true if there is a save, false otherwise
func (obj *instruction) IsSave() bool {
	return obj.save != nil
}

// Save returns the save, if any
func (obj *instruction) Save() variable.Variable {
	return obj.save
}

// IsDelete returns true if there is a delete, false otherwise
func (obj *instruction) IsDelete() bool {
	return obj.del != ""
}

// Delete returns the delete, if any
func (obj *instruction) Delete() string {
	return obj.del
}

// IsMatch returns true if there is a match, false otherwise
func (obj *instruction) IsMatch() bool {
	return obj.match != nil
}

// Match returns the match, if any
func (obj *instruction) Match() match.Match {
	return obj.match
}

// IsToken returns true if there is a token, false otherwise
func (obj *instruction) IsToken() bool {
	return obj.token != nil
}

// Token returns the token, if any
func (obj *instruction) Token() token.Token {
	return obj.token
}
