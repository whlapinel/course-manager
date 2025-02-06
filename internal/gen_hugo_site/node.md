+++
draft = true
date = {{ .Date }}
title = {{ .Node.GetName }}
name = {{ .Node.GetName }}
designation = {{ .Node.Designation }} 
type_name = {{ .Node.TypeName }}
id = {{.Node.GetID}}
number = {{.Node.GetNumber}}
parent_id = {{.Node.GetParentID}}
description = {{.Node.GetDescription}}
parent_type_name = {{.Node.ParentTypeName}}
child_type_name = {{.Node.ChildTypeName}}
+++