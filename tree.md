.
|-- Caddyfile
|-- Caddyfile.dev
|-- Caddyfile.prod
|-- Dockerfile.caddy
|-- Dockerfile.echo
|-- Dockerfile.echo.dev
|-- Dockerfile.marp
|-- LICENSE
|-- README.md
|-- Taskfile.yml
|-- caddy.json
|-- caddy_config
|   `-- autosave.json
|-- caddy_data
|   |-- certificates  [error opening dir]
|   |-- instance.uuid
|   |-- last_clean.json
|   |-- locks  [error opening dir]
|   `-- pki  [error opening dir]
|-- cmd
|   |-- export_secret
|   |   `-- main.go
|   |-- migrate
|   |   `-- main.go
|   `-- web_app
|       |-- logger.go
|       |-- main.go
|       |-- slide_template.md
|       `-- web_app
|-- compose.dev.yaml
|-- compose.prod.yaml
|-- compose.yaml
|-- docs-input.css
|-- generate.go
|-- go.mod
|-- go.sum
|-- internal
|   |-- assets
|   |   |-- dist
|   |   |   |-- assessments.d.ts
|   |   |   |-- assessments.js
|   |   |   |-- details-hover.d.ts
|   |   |   |-- details-hover.js
|   |   |   |-- favicon.svg
|   |   |   |-- htmx.js
|   |   |   |-- index.d.ts
|   |   |   |-- index.js
|   |   |   |-- lesson-details.d.ts
|   |   |   |-- lesson-details.js
|   |   |   |-- markdown-styles.css
|   |   |   |-- screenshot.png
|   |   |   |-- signin.png
|   |   |   |-- spinner.svg
|   |   |   |-- styles.css
|   |   |   |-- tab.d.ts
|   |   |   |-- tab.js
|   |   |   `-- types
|   |   |       |-- assessments.d.ts
|   |   |       |-- details-hover.d.ts
|   |   |       |-- index.d.ts
|   |   |       `-- lesson-details.d.ts
|   |   |-- embed.go
|   |   `-- ts
|   |       |-- @types
|   |       |   `-- htmx.d.ts
|   |       |-- assessments.ts
|   |       |-- details-hover.ts
|   |       |-- index.ts
|   |       |-- lesson-details.ts
|   |       `-- tab.ts
|   |-- authentication
|   |   |-- authentication.go
|   |   `-- doc.go
|   |-- authorization
|   |   `-- authorization.go
|   |-- components -> components
|   |-- data
|   |   |-- assessment.go
|   |   |-- common.go
|   |   |-- common_test.go
|   |   |-- course.go
|   |   |-- course_test.go
|   |   |-- csv
|   |   |   `-- csv.go
|   |   |-- csv_files
|   |   |   |-- courses_export_Sun,\ 29\ Dec\ 2024\ 11:09:38\ EST
|   |   |   |-- curricula.csv
|   |   |   |-- non_instruct_days.csv
|   |   |   |-- objectives.csv
|   |   |   |-- schedules.csv
|   |   |   |-- schedules_new.csv
|   |   |   |-- standards.csv
|   |   |   |-- terms.csv
|   |   |   `-- wow_test.csv
|   |   |-- database
|   |   |   |-- assignments.sql
|   |   |   |-- assignments.sql.go
|   |   |   |-- course_manager\ copy.db
|   |   |   |-- course_manager.db
|   |   |   |-- courses.sql
|   |   |   |-- courses.sql.go
|   |   |   |-- dates.sql
|   |   |   |-- dates.sql.go
|   |   |   |-- db.go
|   |   |   |-- lessons.sql
|   |   |   |-- lessons.sql.go
|   |   |   |-- migrations
|   |   |   |   |-- archive
|   |   |   |   |   |-- mixed_data_schema
|   |   |   |   |   |   `-- 20250117195114_add_on_delete_cascade.sql
|   |   |   |   |   `-- schema
|   |   |   |   |       |-- 20250118131846_add_term_descr.sql
|   |   |   |   |       |-- 20250118165850_start.sql
|   |   |   |   |       |-- 20250118195415_remove_day_num.sql
|   |   |   |   |       |-- 20250119135405_drop_non_instruct_days_table.sql
|   |   |   |   |       |-- 20250123232353_remove_files.sql
|   |   |   |   |       |-- 20250124183507_drop_images.sql
|   |   |   |   |       |-- 20250125141659_drop_objectives.sql
|   |   |   |   |       |-- 20250125142239_add_parentid_to_stds.sql
|   |   |   |   |       |-- 20250125142632_add_descr_to_stds.sql
|   |   |   |   |       |-- 20250125143748_modify_standards_parentid.sql
|   |   |   |   |       |-- 20250125144607_modify_lesson_stds.sql
|   |   |   |   |       |-- 20250125151625_add_table_std_set.sql
|   |   |   |   |       |-- 20250125152446_make_stdset_cname_unique.sql
|   |   |   |   |       |-- 20250125152921_make_std_courseid_fkey.sql
|   |   |   |   |       |-- 20250126130115_rename_std_set.sql
|   |   |   |   |       |-- 20250126152612_mod_lesson_standards.sql
|   |   |   |   |       |-- 20250126182329_mod_course.sql
|   |   |   |   |       |-- 20250131163931_add_assignments.sql
|   |   |   |   |       |-- 20250131164729_rename_assessments.sql
|   |   |   |   |       |-- 20250131165635_alter_assessments.sql
|   |   |   |   |       |-- 20250131171138_alter_category.sql
|   |   |   |   |       |-- 20250203162346_add_user.sql
|   |   |   |   |       `-- 20250213155530_add_occasion.sql
|   |   |   |   |-- data
|   |   |   |   |   `-- 20250216122111_add_user_fk_terms.sql
|   |   |   |   `-- schema
|   |   |   |       |-- 20250216163022_alter_user_id.sql
|   |   |   |       |-- 20250227141804_add_assgn_filepath.sql
|   |   |   |       |-- 20250227151625_modify_terms_user_id.sql
|   |   |   |       |-- 20250403145551_mod_category.go
|   |   |   |       |-- 20250405170401
|   |   |   |       |   |-- fifth.sql
|   |   |   |       |   |-- first.sql
|   |   |   |       |   |-- fourth.sql
|   |   |   |       |   |-- second.sql
|   |   |   |       |   `-- third.sql
|   |   |   |       `-- 20250405170401_assessment_parent.go
|   |   |   |-- models.go
|   |   |   |-- occasions.sql
|   |   |   |-- occasions.sql.go
|   |   |   |-- schema.go
|   |   |   |-- schema.sql
|   |   |   |-- standard_sets.sql
|   |   |   |-- standard_sets.sql.go
|   |   |   |-- standards.sql
|   |   |   |-- standards.sql.go
|   |   |   |-- terms.sql
|   |   |   |-- terms.sql.go
|   |   |   |-- units.sql
|   |   |   |-- units.sql.go
|   |   |   |-- users.sql
|   |   |   `-- users.sql.go
|   |   |-- lesson.go
|   |   |-- lesson_test.go
|   |   |-- node.go
|   |   |-- node_test.go
|   |   |-- standard_sets.go
|   |   |-- standards.go
|   |   |-- standards_test.go
|   |   |-- term.go
|   |   |-- term_test.go
|   |   |-- unit.go
|   |   |-- unit_test.go
|   |   `-- user.go
|   |-- domain
|   |   |-- assessment.go
|   |   |-- category.go
|   |   |-- course.go
|   |   |-- document.go
|   |   |-- image.go
|   |   |-- lesson.go
|   |   |-- lesson_test.go
|   |   |-- node.go
|   |   |-- occasion.go
|   |   |-- school.go
|   |   |-- standard.go
|   |   |-- term.go
|   |   |-- term_test.go
|   |   |-- unit.go
|   |   `-- user.go
|   |-- gen_site
|   |   |-- generator.go
|   |   `-- markdown.go
|   |-- handlers
|   |   |-- assessments.go
|   |   |-- auth.go
|   |   |-- calendar.go
|   |   |-- common.go
|   |   |-- common_test.go
|   |   |-- courses.go
|   |   |-- home.go
|   |   |-- lessons.go
|   |   |-- node.go
|   |   |-- node_test.go
|   |   |-- terms.go
|   |   |-- units.go
|   |   |-- user_home.go
|   |   `-- util.go
|   |-- service
|   |   |-- assessments.go
|   |   |-- common_test.go
|   |   |-- course_manager.db
|   |   |-- courses.go
|   |   |-- courses_test.go
|   |   |-- files.go
|   |   |-- generate.go
|   |   |-- generator_new.go
|   |   |-- lesson.go
|   |   |-- markdown.go
|   |   |-- markdown_test.go
|   |   |-- node.go
|   |   |-- service.go
|   |   |-- slides.go
|   |   |-- slides_test.go
|   |   |-- standards.go
|   |   |-- terms.go
|   |   |-- unit.go
|   |   `-- user.go
|   |-- templates
|   |   |-- app
|   |   |   |-- assessments.go
|   |   |   |-- assessments.templ
|   |   |   |-- assessments_templ.go
|   |   |   |-- assessments_test.go
|   |   |   |-- authentication.go
|   |   |   |-- authentication.templ
|   |   |   |-- authentication_templ.go
|   |   |   |-- calendar.go
|   |   |   |-- calendar.templ
|   |   |   |-- calendar_templ.go
|   |   |   |-- common.go
|   |   |   |-- common.templ
|   |   |   |-- common_templ.go
|   |   |   |-- courses.go
|   |   |   |-- courses.templ
|   |   |   |-- courses_templ.go
|   |   |   |-- files.go
|   |   |   |-- files.templ
|   |   |   |-- files_templ.go
|   |   |   |-- head.templ
|   |   |   |-- head_templ.go
|   |   |   |-- home.go
|   |   |   |-- home.templ
|   |   |   |-- home_templ.go
|   |   |   |-- layout.go
|   |   |   |-- layout.templ
|   |   |   |-- layout_templ.go
|   |   |   |-- lesson_details.go
|   |   |   |-- lesson_details.templ
|   |   |   |-- lesson_details_templ.go
|   |   |   |-- lessons.go
|   |   |   |-- lessons.templ
|   |   |   |-- lessons_templ.go
|   |   |   |-- markdown.go
|   |   |   |-- markdown.templ
|   |   |   |-- markdown_templ.go
|   |   |   |-- new_lesson.go
|   |   |   |-- new_lesson.templ
|   |   |   |-- new_lesson_templ.go
|   |   |   |-- node.templ
|   |   |   |-- node_create.go
|   |   |   |-- node_create.templ
|   |   |   |-- node_create_templ.go
|   |   |   |-- node_details.go
|   |   |   |-- node_details.templ
|   |   |   |-- node_details_templ.go
|   |   |   |-- node_list.go
|   |   |   |-- node_list.templ
|   |   |   |-- node_list_templ.go
|   |   |   |-- node_templ.go
|   |   |   |-- page.go
|   |   |   |-- standards.go
|   |   |   |-- standards.templ
|   |   |   |-- standards_templ.go
|   |   |   |-- terms.go
|   |   |   |-- terms.templ
|   |   |   |-- terms_templ.go
|   |   |   |-- unauthorized.go
|   |   |   |-- unauthorized.templ
|   |   |   |-- unauthorized_templ.go
|   |   |   |-- unit_details.go
|   |   |   |-- unit_details.templ
|   |   |   |-- unit_details_templ.go
|   |   |   |-- units.go
|   |   |   |-- units.templ
|   |   |   |-- units_templ.go
|   |   |   |-- user_home.go
|   |   |   |-- user_home.templ
|   |   |   `-- user_home_templ.go
|   |   |-- components
|   |   |   |-- app
|   |   |   `-- base
|   |   |       |-- breadcrumbs.go
|   |   |       |-- breadcrumbs.templ
|   |   |       |-- breadcrumbs_templ.go
|   |   |       |-- button.go
|   |   |       |-- button.templ
|   |   |       |-- button_templ.go
|   |   |       |-- description_list.go
|   |   |       |-- description_list.templ
|   |   |       |-- description_list_templ.go
|   |   |       |-- doc.go
|   |   |       |-- element.go
|   |   |       |-- form.go
|   |   |       |-- form.templ
|   |   |       |-- form_templ.go
|   |   |       |-- hx_element.go
|   |   |       |-- icons.templ
|   |   |       |-- icons_templ.go
|   |   |       |-- input.go
|   |   |       |-- input.templ
|   |   |       |-- input_templ.go
|   |   |       |-- label.go
|   |   |       |-- label.templ
|   |   |       |-- label_templ.go
|   |   |       |-- layout.go
|   |   |       |-- layout.templ
|   |   |       |-- layout_templ.go
|   |   |       |-- link.go
|   |   |       |-- link.templ
|   |   |       |-- link_templ.go
|   |   |       |-- mobile.go
|   |   |       |-- mobile.templ
|   |   |       |-- mobile_templ.go
|   |   |       |-- notifications.templ
|   |   |       |-- notifications_templ.go
|   |   |       |-- page_heading.go
|   |   |       |-- page_heading.templ
|   |   |       |-- page_heading_templ.go
|   |   |       |-- select.go
|   |   |       |-- select.templ
|   |   |       |-- select_templ.go
|   |   |       |-- tab.go
|   |   |       |-- tab.templ
|   |   |       |-- tab.ts
|   |   |       |-- tab_templ.go
|   |   |       |-- table.go
|   |   |       |-- table.templ
|   |   |       |-- table_templ.go
|   |   |       |-- textarea.go
|   |   |       |-- user_menu.go
|   |   |       |-- user_menu.templ
|   |   |       `-- user_menu_templ.go
|   |   |-- shared
|   |   |   `-- common.go
|   |   `-- static
|   |       |-- breadcrumbs.go
|   |       |-- course_calendar.go
|   |       |-- course_calendar.templ
|   |       |-- course_calendar_templ.go
|   |       |-- file.go
|   |       |-- file.templ
|   |       |-- file_templ.go
|   |       |-- head.go
|   |       |-- head.templ
|   |       |-- head_templ.go
|   |       |-- home.go
|   |       |-- home.templ
|   |       |-- home_templ.go
|   |       |-- layout.go
|   |       |-- layout.templ
|   |       |-- layout_templ.go
|   |       |-- lesson_details.go
|   |       |-- lesson_details.templ
|   |       |-- lesson_details_templ.go
|   |       |-- markdown.go
|   |       |-- markdown.templ
|   |       |-- markdown_templ.go
|   |       |-- node.templ
|   |       |-- node_details.go
|   |       |-- node_list.go
|   |       `-- node_templ.go
|   `-- util
|       |-- util.go
|       `-- util_test.go
|-- manager-input.css
|-- output.txt
|-- package-lock.json
|-- package.json
|-- scripts
|   |-- backup.sh
|   |-- copy_config.sh
|   |-- copy_data.sh
|   |-- restore_backup.sh
|   |-- test.sh
|   `-- unpack_data.sh
|-- sqlc.yml
|-- tailwind.config.js
|-- test.txt
|-- tree.md
|-- tsconfig.json
|-- tsconfig.server.json
|-- tsconfig.server.tsbuildinfo
|-- tsconfig.static.json
`-- tsconfig.static.tsbuildinfo

41 directories, 365 files
