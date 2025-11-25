package widget

import (
	"encoding/json"
	"html/template"
	"strings"

	component "github.com/kurniajigunawan/mikrotik-portal/public/component"
)

type Form struct {
	ActionURL    string
	Method       string
	Body         map[string]interface{}
	Fields       []FieldType
	SubmitButton component.ButtonSolid
}

func (f Form) RenderFormBody() template.JS {
	var bodyBuilder strings.Builder

	// --- Part 1: Generate client-side variable assignments ---
	// This creates: const service_id = document.querySelector('#service_id').value;
	for _, v := range f.Fields {
		id := v.GetID()
		bodyBuilder.WriteString("const " + id + "raw = document.querySelector('#" + id + "').value;\n")
		if v.GetValueType() == Number {
			bodyBuilder.WriteString("const " + id + " = parseInt(" + id + "raw);\n")
		} else {
			bodyBuilder.WriteString("const " + id + " = " + id + "raw;\n")
		}
	}

	// --- Part 2: Generate the final 'body' object literal ---

	// Use a helper function to recursively build the JavaScript object literal string
	// This is much safer than manipulating the marshaled JSON string.
	jsObjectLiteral, err := createJSObjectLiteral(f.Body)
	if err != nil {
		// Log the error in the server logs (not returning it to the client)
		// Returning a safe empty block instead of an error message in the JS
		return "// Error: Could not generate form body."
	}

	bodyBuilder.WriteString("var body = " + jsObjectLiteral + ";\n")

	// Return the raw, executable JavaScript code
	return template.JS(bodyBuilder.String())
}

// 💡 New Helper Function: Recursively creates a JS object literal string
func createJSObjectLiteral(data map[string]interface{}) (string, error) {
	var parts []string

	for key, val := range data {
		switch v := val.(type) {
		case string:
			// CRITICAL FIX: The value is the name of the JS variable created in Part 1.
			// It should be injected directly, NOT as a string literal ("'service_id'")
			parts = append(parts, key+": "+v)

		case map[string]interface{}:
			// Recursively handle nested maps
			nested, err := createJSObjectLiteral(v)
			if err != nil {
				return "", err
			}
			parts = append(parts, key+": "+nested)

		case bool, int, float64:
			// Handle other literal types if they exist in your f.Body, e.g., for config
			parts = append(parts, key+": "+jsonString(v))

		default:
			// Use the standard JSON marshaler for any unexpected types
			// This is safer than manual conversion
			jsonValue, err := json.Marshal(v)
			if err != nil {
				return "", err
			}
			parts = append(parts, key+": "+string(jsonValue))
		}
	}

	// Join parts and wrap in braces to form the final object literal
	return "{ " + strings.Join(parts, ", ") + " }", nil
}

// Helper to convert non-string primitives to JSON strings
func jsonString(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

type valueType string

const (
	String valueType = "string"
	Number valueType = "number"
)

// Field Type Section
type FieldType interface {
	GetID() string
	GetValueType() valueType
	Render() template.HTML
}

type Input struct {
	ID    string
	Label string
	Name  string
	// text, number, password
	Type string
	// default value
	Value     string
	ValueType valueType
}

func (i Input) GetID() string {
	return i.ID
}

func (i Input) GetValueType() valueType {
	return i.ValueType
}

func (i Input) Render() template.HTML {
	return template.HTML("<input type=\"" + i.Type + "\" id=\"" + i.ID + "\" name=\"" + i.Name + "\" required autocomplete=\"" + i.ID + "\" class=\"block w-full rounded-md bg-white/5 px-3 py-1.5 text-base text-white outline-1 -outline-offset-1 outline-white/10 placeholder:text-gray-500 focus:outline-2 focus:-outline-offset-2 focus:outline-indigo-500 sm:text-sm/6\" value=\"" + i.Value + "\" />")
}

type Select struct {
	ID        string
	Label     string
	Name      string
	ValueType valueType
	// option value as key, option display as value
	Options map[string]string
}

func (s Select) GetID() string {
	return s.ID
}

func (s Select) GetValueType() valueType {
	return s.ValueType
}

func (s Select) Render() template.HTML {
	var option string
	var indexOpt int
	var button string
	var buttonVal string
	for key, value := range s.Options {
		opt := "<el-option value=\"" + key + "\" class=\"group/option relative block cursor-default py-2 pr-9 pl-3 text-white select-none focus:bg-indigo-500 focus:text-white focus:outline-hidden\">"
		opt += "<div class=\"flex items-center\">"
		opt += "<span class=\"block truncate font-normal group-aria-selected/option:font-semibold\">" + value + "</span>"
		opt += "</div>"
		opt += "<span class=\"absolute inset-y-0 right-0 flex items-center pr-4 text-indigo-400 group-not-aria-selected/option:hidden group-focus/option:text-white in-[el-selectedcontent]:hidden\">"
		opt += "<svg viewBox=\"0 0 20 20\" fill=\"currentColor\" data-slot=\"icon\" aria-hidden=\"true\" class=\"size-5\">"
		opt += "<path d=\"M16.704 4.153a.75.75 0 0 1 .143 1.052l-8 10.5a.75.75 0 0 1-1.127.075l-4.5-4.5a.75.75 0 0 1 1.06-1.06l3.894 3.893 7.48-9.817a.75.75 0 0 1 1.05-.143Z\" clip-rule=\"evenodd\" fill-rule=\"evenodd\" />"
		opt += "</svg>"
		opt += "</span>"
		opt += "</el-option>"
		option = option + opt
		indexOpt++
		if indexOpt == len(s.Options)-1 {
			buttonVal = key
			button = "<button type=\"button\" class=\"grid w-full cursor-default grid-cols-1 rounded-md bg-white/5 py-1.5 pr-2 pl-3 text-left text-white outline-1 -outline-offset-1 outline-white/10 focus-visible:outline-2 focus-visible:-outline-offset-2 focus-visible:outline-indigo-500 sm:text-sm/6\">"
			button += "<el-selectedcontent class=\"col-start-1 row-start-1 flex items-center gap-3 pr-6\">"
			button += "<span class=\"block truncate\">" + value + "</span>"
			button += "</el-selectedcontent>"
			button += "<svg viewBox=\"0 0 16 16\" fill=\"currentColor\" data-slot=\"icon\" aria-hidden=\"true\" class=\"col-start-1 row-start-1 size-5 self-center justify-self-end text-gray-400 sm:size-4\">"
			button += "<path d=\"M5.22 10.22a.75.75 0 0 1 1.06 0L8 11.94l1.72-1.72a.75.75 0 1 1 1.06 1.06l-2.25 2.25a.75.75 0 0 1-1.06 0l-2.25-2.25a.75.75 0 0 1 0-1.06ZM10.78 5.78a.75.75 0 0 1-1.06 0L8 4.06 6.28 5.78a.75.75 0 0 1-1.06-1.06l2.25-2.25a.75.75 0 0 1 1.06 0l2.25 2.25a.75.75 0 0 1 0 1.06Z\" clip-rule=\"evenodd\" fill-rule=\"evenodd\" />"
			button += "</svg>"
			button += "</button>"
		}
	}
	selectElement := "<el-select id=\"" + s.ID + "\" name=\"" + s.Name + "\" value=\"" + buttonVal + "\" class=\"mt-2 block\">"
	selectElement += button
	selectElement += "<el-options anchor=\"bottom start\" popover class=\"max-h-56 w-(--button-width) overflow-auto rounded-md bg-gray-800 py-1 text-base outline-1 -outline-offset-1 outline-white/10 [--anchor-gap:--spacing(1)] data-leave:transition data-leave:transition-discrete data-leave:duration-100 data-leave:ease-in data-closed:data-leave:opacity-0 sm:text-sm\">"
	selectElement += option
	selectElement += "</el-options></el-select>"
	selectElement += "<script src=\"https://cdn.jsdelivr.net/npm/@tailwindplus/elements@1\" type=\"module\"></script>"
	return template.HTML(selectElement)
}

type Heading struct {
	Title    string
	Subtitle string
}

type MenuItem struct {
	LinkURL   string
	Title     string
	Subtitle  string
	Icon      string
	IconColor string
}
