package evaluator

import (
	"../lexer"
	"../object"
	"../parser"
	"testing"
)

func TestEvalIntegerExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected int64
	}{
		{"5", 5},
		{"10", 10},
		//{"-5", -5},
		//{"-10", -10},
		//{"5 + 5 + 5 + 5 - 10", 10},
		//{"2 * 2 * 2 * 2 * 2", 32},
		//{"-50 + 100 + -50", 0},
		//{"5 * 2 + 10", 20},
		//{"5 + 2 * 10", 25},
		//{"20 + 2 * -10", 0},
		//{"50 / 2 * 2 + 10", 60},
		//{"2 * (5 + 10)", 30},
		//{"3 * 3 * 3 + 10", 37},
		//{"3 * (3 * 3) + 10", 37},
		//{"(5 + 10 * 2 + 15 / 3) * 2 + -10", 50},
	}

	for _, tt := range tests {
		evaluated := testEval(tt.input)
		testIntegerObject(t, evaluated, tt.expected)
	}
}

func testEval(input string) object.Object {
	l := lexer.New(input)
	p := parser.New(l)
	program := p.ParseProgram()

	return Eval(program)
}

func testIntegerObject(t *testing.T, obj object.Object, expected int64) bool {
	result, ok := obj.(*object.Integer)
	if !ok {
		t.Errorf("object is not Integer. got=%T (%+v)", obj, obj)
		return false
	}
	if result.Value != expected {
		t.Errorf("object has wrong value. got=%d, want=%d",
			result.Value, expected)
		return false
	}

	return true
}
