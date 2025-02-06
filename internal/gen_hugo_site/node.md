+++
date = "{{ .Date }}"
draft = true
title = "{{ .Node.GetName }}"
name = "{{ .Node.GetName }}"
designation = {{ with .Node.Designation }}"{{ . }}"{{ else }}""{{ end }}
type_name = "{{ .Node.TypeName }}"
id = {{ .Node.GetID }}
number = {{ .Node.GetNumber }}
parent_id = {{ .Node.GetParentID }}
description = {{ with .SanitizedDescription }}"{{ . }}"{{ else }}""{{ end }}
parent_type_name = "{{ .Node.ParentTypeName }}"
child_type_name = "{{ .Node.ChildTypeName }}"
+++
